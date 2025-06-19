package glog

import (
    "context"
    "errors"
    "fmt"
    "log/slog"
    "path"
    "runtime"
    "time"
    
    "gorm.io/gorm"
    gormlogger "gorm.io/gorm/logger"
    "gorm.io/gorm/utils"
)

// New creates a new logger with default config
func New() *Logger {
    return NewWithConfig(NewConfig(slog.Default().Handler()))
}

// NewWithConfig creates a new logger with given config
func NewWithConfig(config *Config) *Logger {
    return &Logger{
        config,
    }
}

type Logger struct {
    *Config
}

// ensure our logger implements gorm logger.Interface
var _ gormlogger.Interface = (*Logger)(nil)

// LogMode log mode
func (l *Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
    // This function will switch to logging all queries, whenever the level is set to Info.
    // It's to support the Debug() function of gorm which sets the log level to info for subsequent queries, see:
    //   https://gorm.io/docs/session.html#Debug
    
    // Note: Error and Warn levels are ignored as the log level is managed by slog already.
    if level == gormlogger.Error || level == gormlogger.Warn {
        return l
    }
    
    // clone logger for session mode
    nc := l.Config.clone()
    nl := NewWithConfig(nc.WithTraceAll(level == gormlogger.Info).WithSilent(level == gormlogger.Silent))
    return nl
}

// Info logs info message
func (l *Logger) Info(ctx context.Context, format string, args ...any) {
    if l.enabled(ctx, slog.LevelInfo) {
        l.log(ctx, slog.LevelInfo, fmt.Sprintf(format, args...), l.contextAttrs(ctx)...)
    }
}

// Warn logs warn message
func (l *Logger) Warn(ctx context.Context, format string, args ...any) {
    if l.enabled(ctx, slog.LevelWarn) {
        l.log(ctx, slog.LevelWarn, fmt.Sprintf(format, args...), l.contextAttrs(ctx)...)
    }
}

// Error logs error message
func (l *Logger) Error(ctx context.Context, format string, args ...any) {
    if l.enabled(ctx, slog.LevelError) {
        l.log(ctx, slog.LevelError, fmt.Sprintf(format, args...), l.contextAttrs(ctx)...)
    }
}

// Trace logs sql message
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    if l.silent {
        return
    }
    
    elapsed := time.Since(begin)
    switch {
    case err != nil && l.enabled(ctx, slog.LevelError) && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.ignoreRecordNotFoundError):
        attrs := l.traceAttrs(ctx, elapsed, fc, utils.FileWithLineNum(), err, false)
        l.log(ctx, slog.LevelError, l.errorMsg, attrs...)
    case l.slowThreshold != 0 && elapsed > l.slowThreshold && l.enabled(ctx, slog.LevelWarn):
        attrs := l.traceAttrs(ctx, elapsed, fc, utils.FileWithLineNum(), nil, true)
        l.log(ctx, slog.LevelWarn, l.slowMsg, attrs...)
    case l.traceAll && l.enabled(ctx, slog.LevelInfo):
        attrs := l.traceAttrs(ctx, elapsed, fc, utils.FileWithLineNum(), nil, false)
        l.log(ctx, slog.LevelInfo, l.okMsg, attrs...)
    }
}

// ParamsFilter filter params
func (l *Logger) ParamsFilter(_ context.Context, sql string, params ...interface{}) (string, []interface{}) {
    if l.parameterizedQueries {
        return sql, nil
    }
    return sql, params
}

// log adds context attributes and logs a message with the given slog level
func (l *Logger) log(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
    var pcs [1]uintptr
    // skip [runtime.Callers, this function, this function's caller]
    runtime.Callers(3, pcs[:])
    r := slog.NewRecord(time.Now(), level, msg, pcs[0])
    r.AddAttrs(attrs...)
    
    if ctx == nil {
        ctx = context.Background()
    }
    _ = l.slogHandler.Handle(ctx, r)
}

func (l *Logger) traceAttrs(ctx context.Context, elapsed time.Duration, fc func() (string, int64), file string, err error, slow bool) []slog.Attr {
    sql, rows := fc()
    
    attrs := make([]slog.Attr, 0, 5)
    
    if l.durationKey != "" {
        attrs = append(attrs, slog.Duration(l.durationKey, elapsed))
    }
    if rows >= 0 && l.rowsKey != "" { // rows could be -1
        attrs = append(attrs, slog.Int64(l.rowsKey, rows))
    }
    if l.sourceKey != "" {
        if l.fullSourcePath {
            attrs = append(attrs, slog.String(l.sourceKey, file))
        } else {
            attrs = append(attrs, slog.String(l.sourceKey, path.Base(file)))
        }
    }
    if err != nil && l.errorKey != "" {
        attrs = append(attrs, slog.Any(l.errorKey, err))
    } else if slow && l.slowThresholdKey != "" {
        attrs = append(attrs, slog.Duration(l.slowThresholdKey, l.slowThreshold))
    }
    if l.queryKey != "" {
        attrs = append(attrs, slog.String(l.queryKey, sql))
    }
    
    if l.groupKey != "" {
        return append(l.contextAttrs(ctx), slog.Attr{Key: l.groupKey, Value: slog.GroupValue(attrs...)})
    }
    
    return append(l.contextAttrs(ctx), attrs...)
}

// contextAttrs extracts attributes from context
func (l *Logger) contextAttrs(ctx context.Context) []slog.Attr {
    if ctx == nil {
        ctx = context.Background()
    }
    
    var attrs []slog.Attr
    if l.contextExtractor != nil {
        attrs = l.contextExtractor(ctx)
    }
    for ak, ck := range l.contextKeys {
        if val := ctx.Value(ck); val != nil {
            attrs = append(attrs, slog.Any(ak, val))
        }
    }
    return attrs
}

// enabled reports whether the logger is enabled at the given level
func (l *Logger) enabled(ctx context.Context, lvl slog.Level) bool {
    return !l.silent && l.slogHandler.Enabled(ctx, lvl)
}
