package db

import (
	"demo/src/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Instance *gorm.DB
}

func (db *DB) Init(c *config.DbConfig) {
	// tpuser:1234@tcp(127.0.0.1:4000)/TPDB?charset=utf8mb4&parseTime=True&loc=Local
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/TPDB?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port)
	instance, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Instance = instance
}
