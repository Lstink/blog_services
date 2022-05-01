package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// String 转化为字符串
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

// NewLogger 实例化日志
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	// 实例化日志
	i := log.New(w, prefix, flag)
	// 返回自定义日志结构体
	return &Logger{newLogger: i}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller 堆栈信息
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}

	return ll
}

func (l *Logger) WithCallerFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers

	return ll
}

func (l *Logger) JOSNFormat(message string) map[string]interface{} {
	// 实例化map，长度为字段长度+4
	data := make(Fields, len(l.fields)+4)
	// 实例化map的 time 字段赋值为纳秒的时间戳
	data["time"] = time.Now().Local().UnixNano()
	// 实例化map的 message 字段赋值为传过来的字符串
	data["message"] = message
	// 实例化map的 callers 赋值
	data["callers"] = l.callers
	// 如果这个字段存在值，则遍历这个字符串切片，把值复制到当前实例化的map中
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	// 最后返回这个map的内容
	return data
}

// Output 输出日志内容 接收 日志等级 和 内容
func (l *Logger) Output(level Level, message string) {
	// 将map转化为 json 格式的字符串
	body, _ := json.Marshal(l.JOSNFormat(message))
	// 格式化日志字符串内容为想要的自定义格式
	content := fmt.Sprintf("[%s] %s", level, string(body))
	// 判断日志等级-写入日志
	switch level {
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// Info info等级的日志
func (l *Logger) Info(v ...any) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

// Infof info等级的日志，接收格式化参数
func (l *Logger) Infof(format string, v ...any) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

// Fatal fatal等级的日志
func (l *Logger) Fatal(v ...any) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

// Fatalf fatal等级的日志，接收格式化参数
func (l *Logger) Fatalf(format string, v ...any) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

// Errorf error等级的日志，接收格式化参数
func (l *Logger) Errorf(format string, v ...any) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

// Error error等级的日志
func (l *Logger) Error(v ...any) {
	l.Output(LevelError, fmt.Sprint(v...))
}
