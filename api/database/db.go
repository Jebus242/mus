package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jebus24/mus/config"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
