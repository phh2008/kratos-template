package logger

import (
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/exp/zapslog"
    "go.uber.org/zap/zapcore"
    "io"
    "log/slog"
    "os"
    "strings"
    "time"
)

func init() {
    // 做为slog的后端
    zapLog := zap.NewNop()
    sl := slog.New(zapslog.NewHandler(zapLog.Core(), zapslog.WithCaller(true)))
    slog.SetDefault(sl)
}

var levelMap = map[string]zapcore.Level{
    "debug": zapcore.DebugLevel,
    "info":  zapcore.InfoLevel,
    "warn":  zapcore.WarnLevel,
    "error": zapcore.ErrorLevel,
}

type Config struct {
    Level      string // 日志级别: debug, info, warn, error
    Filename   string // 日志文件路径
    MaxSize    int32  // 每个日志文件保存的最大尺寸 单位：M
    MaxBackups int32  // 保留旧日志文件数量
    MaxAge     int32  // 日志保留时间（天）
    Compress   bool   // 是否压缩
    LocalTime  bool   // 是否使用本地时间
}

// NewLogger 创建 logger
func NewLogger(config *Config) *zap.Logger {
    zapLog := newZapLog(config)
    // 替换 zap 全局 zapLog
    zap.ReplaceGlobals(zapLog)
    // 做为slog的后端
    sl := slog.New(zapslog.NewHandler(zapLog.Core(), zapslog.WithCaller(true)))
    slog.SetDefault(sl)
    return zapLog
}

// getWriter
func getWriter(config *Config) io.Writer {
    return &lumberjack.Logger{
        Filename:   config.Filename,
        MaxSize:    int(config.MaxSize), // megabytes
        MaxBackups: int(config.MaxBackups),
        MaxAge:     int(config.MaxAge), //days
        LocalTime:  config.LocalTime,
        Compress:   config.Compress, // disabled by default
    }
}

func newZapLog(config *Config) *zap.Logger {
    // 设置日志格式
    //encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
    encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
        MessageKey:  "msg",
        LevelKey:    "level",
        EncodeLevel: zapcore.CapitalLevelEncoder,
        TimeKey:     "ts",
        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(t.Format("2006-01-02 15:04:05"))
        },
        CallerKey:    "file",
        EncodeCaller: zapcore.ShortCallerEncoder,
        EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendInt64(int64(d) / 1000000)
        },
    })

    // 记录什么级别的日志
    level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= levelMap[strings.ToLower(config.Level)]
    })

    // 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
    writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter(config)))
    // 如果info、debug、error分文件记录，就创建多个 writer
    // 最后创建具体的Logger
    core := zapcore.NewTee(
        zapcore.NewCore(encoder, writer, level), // 可添加多个
    )
    // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
    return zap.New(core, zap.AddCaller())
}
