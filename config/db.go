package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gowiki/models"
)

func InitDB(cfg Config) (*gorm.DB, error) {
	dsn := "example:example@tcp(localhost:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = cfg.DBInfo

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Page{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
