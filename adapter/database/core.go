package database

import (
	"fmt"
	"math"
	"time"

	"music-app/adapter/database/model"
	"music-app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB() (*gorm.DB, error) {
	dsn := config.DSN()

	gormConfig := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)

	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetConnMaxIdleTime(100)
	sqlDB.SetMaxOpenConns(100)

	// Check connection
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retryMax-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		fmt.Errorf("failed to connect to database retrying")
		time.Sleep(duration)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.BuiltinBoard{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Message{}); err != nil {
		return err
	}
	return nil
}


