package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warningLogger *log.Logger
	logFiles      []*os.File
)

func init() {
	// 创建日志目录
	logDir := "log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	// 生成当天的日志文件名
	today := time.Now().Format("2006-01-02")
	allLogPath := filepath.Join(logDir, fmt.Sprintf("%s-all.log", today))
	errLogPath := filepath.Join(logDir, fmt.Sprintf("%s-err.log", today))

	// 打开日志文件
	allFile, err := os.OpenFile(allLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开全部日志文件失败: %v", err)
	}
	errFile, err := os.OpenFile(errLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开错误日志文件失败: %v", err)
	}

	// 保存文件句柄以便后续关闭
	logFiles = append(logFiles, allFile, errFile)

	// 设置日志格式
	flags := log.Ldate | log.Ltime

	// 同时输出到文件和控制台
	infoLogger = log.New(io.MultiWriter(os.Stdout, allFile), "[INFO] ", flags)
	errorLogger = log.New(io.MultiWriter(os.Stderr, allFile, errFile), "[ERROR] ", flags)
	warningLogger = log.New(io.MultiWriter(os.Stdout, allFile), "[WARN] ", flags)
}

// CloseLogFiles 关闭日志文件
func CloseLogFiles() {
	for _, file := range logFiles {
		file.Close()
	}
}

// 其他日志函数保持不变
func Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	infoLogger.Printf("%s | %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	errorLogger.Printf("%s | %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	warningLogger.Printf("%s | %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}