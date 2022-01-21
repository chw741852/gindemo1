package model

import (
	"fmt"
	"test/internal/config"
	"test/internal/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config config.Config) error {
	defer logger.Log.Sync()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
						config.Mysql.Username,
						config.Mysql.Password,
						config.Mysql.Addr,
						config.Mysql.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Error("Failed to open mysql")
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		logger.Log.Error("Failed to get DB")
		return err
	}
	sqlDb.SetMaxIdleConns(config.Mysql.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.Mysql.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(config.Mysql.ConnMaxLifeTime))
	DB = db
	migration()
	return nil
}

type BaseModel struct {
	Id 				uint 		`gorm: "primaryKey"`
	CreatedTime 	time.Time
	LastChangedTime time.Time
	Status 			int
}

func migration() {
	
}