package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config cấu hình cho ứng dụng
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig cấu hình cho server
type ServerConfig struct {
	Port int
}

// DatabaseConfig cấu hình cho database
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// NewConfig tạo config mới từ biến môi trường
func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("lỗi khi tải file .env: %w", err)
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("lỗi khi parse SERVER_PORT: %w", err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("lỗi khi parse DB_PORT: %w", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: serverPort,
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}, nil
}
