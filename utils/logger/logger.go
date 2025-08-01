package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// 定义上下文中 logger 的键
type loggerKey struct{}

// 定义上下文中 TraceID 的键
type traceIDKey struct{}

var logFile *os.File

// Init 初始化日志系统
func Init() {
	logDir := "log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	today := time.Now().Format("2006-01-02")
	logPath := filepath.Join(logDir, fmt.Sprintf("%s.log", today))

	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件失败: %v", err))
	}

	// 创建一个同时写入文件和标准输出的 writer
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 创建 slog 的 JSON Handler
	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		// 开启源文件和行号
		AddSource: true,
		// 定义日志级别
		Level: slog.LevelDebug,
		// 自定义 ReplaceAttr, 用于从 context 中提取 trace_id
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return a
		},
	})

	// 创建 logger
	logger := slog.New(handler)

	// 设置为默认 logger
	slog.SetDefault(logger)
}

// CloseLogFile 关闭日志文件
func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}

// WithLogger 将 logger 存入 context
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext 从 context 中获取 logger
// 如果没有，则返回默认 logger
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

// WithTraceID 将 TraceID 存入 context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// GetTraceID 从 context 中获取 TraceID
func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(traceIDKey{}).(string); ok {
		return traceID
	}
	return ""
}

// CtxErrorf 是一个辅助函数，用于从 context 记录错误日志
func CtxErrorf(ctx context.Context, format string, args ...interface{}) {
	log := FromContext(ctx).With("trace_id", GetTraceID(ctx))
	log.Error(fmt.Sprintf(format, args...))
}

// CtxInfof 是一个辅助函数，用于从 context 记录信息日志
func CtxInfof(ctx context.Context, format string, args ...interface{}) {
	log := FromContext(ctx).With("trace_id", GetTraceID(ctx))
	log.Info(fmt.Sprintf(format, args...))
}

// CtxWarnf 是一个辅助函数，用于从 context 记录警告日志
func CtxWarnf(ctx context.Context, format string, args ...interface{}) {
	log := FromContext(ctx).With("trace_id", GetTraceID(ctx))
	log.Warn(fmt.Sprintf(format, args...))
}
