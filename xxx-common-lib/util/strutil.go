package util

import (
    "github.com/google/uuid"
    "math/rand"
    "regexp"
    "strings"
    "time"
)

var snakeReg = regexp.MustCompile("[A-Z][a-z]")
var ColumnReg = regexp.MustCompile(`^([a-zA-Z_][a-zA-Z0-9_]*\.[a-zA-Z_][a-zA-Z0-9_]*)$|^([a-zA-Z_][a-zA-Z0-9_]*)$`) //字母数字下划线
var DirectReg = regexp.MustCompile("^asc|desc|ASC|DESC$")

const underline = "_"

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var innerRand = rand.New(rand.NewSource(time.Now().UnixMilli()))

// SnakeCase 驼峰转下划线
func SnakeCase(src string) string {
    str := snakeReg.ReplaceAllStringFunc(src, func(s string) string {
        return underline + s
    })
    return strings.ToLower(strings.TrimLeft(str, underline))
}

// UUID 生成UUID
func UUID() string {
    u := uuid.New()
    return u.String()
}

// UUIDShort 生成UUID (去掉-字符)
func UUIDShort() string {
    u := uuid.New()
    return strings.ReplaceAll(u.String(), "-", "")
}

// RandCode 生成随机6位验证码数字
func RandCode() int {
    minValue := 100000
    maxValue := 999999
    return innerRand.Intn(maxValue-minValue+1) + minValue
}

// Random 随机字符串
func Random(length int) string {
    size := len(chars)
    var result = make([]byte, length)
    for i := 0; i < length; i++ {
        idx := innerRand.Intn(size)
        result[i] = chars[idx]
    }
    return string(result)
}
