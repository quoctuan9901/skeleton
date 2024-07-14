package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"skeleton/internal/config"
)

// PostgresDB struct cho database PostgreSQL
type PostgresDB struct {
	DB *gorm.DB
}

// NewPostgresDB tạo kết nối database PostgreSQL
func NewPostgresDB(cfg config.DatabaseConfig) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("lỗi khi kết nối database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("lỗi khi lấy sql.DB: %w", err)
	}

	// Cấu hình connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return &PostgresDB{
		DB: db,
	}, nil
}

// Close đóng kết nối database
func (p *PostgresDB) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return fmt.Errorf("lỗi khi lấy sql.DB: %w", err)
	}
	return sqlDB.Close()
}
