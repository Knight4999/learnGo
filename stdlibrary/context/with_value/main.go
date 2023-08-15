package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// context 包函数测试
// withValue函数测试

const KEY = "trace_id"

// NewRequestID 生成UUID
func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)

}

// NewContextWithTraceID 获取子context,携带Trace_id de 值，UUID
func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	return ctx
}

// GetContextValue 获取context的值
func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string) //根据Key 取值，并进行类型断言
	if !ok {
		return ""
	}
	return v
}

// PrintLog 打印日志信息
func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s \n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

// ProcessEnter 执行打印程序
func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "博得之门3")
}
func main() {
	ProcessEnter(NewContextWithTraceID())
}
