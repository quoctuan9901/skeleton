package mylogger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger tạo một logger mới với output ra file
func NewLogger(filename string) (*zap.Logger, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)

	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("lỗi khi mở file log: %w", err)
	}

	writer := zapcore.AddSync(logFile)
	core := zapcore.NewCore(fileEncoder, writer, zap.ErrorLevel)
	logger := zap.New(core)

	return logger, nil
}
