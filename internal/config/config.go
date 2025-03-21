package config

import "github.com/jinzhu/gorm"

type Config struct {
	db *gorm.DB
}
