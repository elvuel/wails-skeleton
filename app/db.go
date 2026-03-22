package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

type AppSetting struct {
	ID        uint   `gorm:"primaryKey"`
	Key       string `gorm:"size:128;uniqueIndex"`
	Value     string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type databaseState struct {
	conn *gorm.DB
	path string
}

func openDatabase(appName string) (*databaseState, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("resolve user config dir: %w", err)
	}

	appDir := filepath.Join(configDir, appName)
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return nil, fmt.Errorf("create app data dir: %w", err)
	}

	dbPath := filepath.Join(appDir, "app.db")
	conn, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open sqlite database: %w", err)
	}

	if err := conn.AutoMigrate(&AppSetting{}); err != nil {
		return nil, fmt.Errorf("auto migrate settings table: %w", err)
	}

	return &databaseState{
		conn: conn,
		path: dbPath,
	}, nil
}

func (d *databaseState) Path() string {
	if d == nil {
		return ""
	}
	return d.path
}

func (d *databaseState) Close() error {
	if d == nil || d.conn == nil {
		return nil
	}

	sqlDB, err := d.conn.DB()
	if err != nil {
		return fmt.Errorf("resolve sql db handle: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("close sqlite database: %w", err)
	}

	return nil
}

func (d *databaseState) GetSetting(key string) (string, error) {
	if d == nil {
		return "", nil
	}

	var setting AppSetting
	err := d.conn.Where("key = ?", key).First(&setting).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("query setting %q: %w", key, err)
	}

	return setting.Value, nil
}

func (d *databaseState) SetSetting(key, value string) error {
	if d == nil {
		return nil
	}

	var setting AppSetting
	err := d.conn.Where("key = ?", key).First(&setting).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return d.conn.Create(&AppSetting{
			Key:   key,
			Value: value,
		}).Error
	case err != nil:
		return fmt.Errorf("query setting %q before update: %w", key, err)
	}

	setting.Value = value
	if err := d.conn.Save(&setting).Error; err != nil {
		return fmt.Errorf("save setting %q: %w", key, err)
	}

	return nil
}
