package entity

import "github.com/jinzhu/gorm"

func Db() *gorm.DB {
	if dbConn == nil {
		return nil
	}

	return dbConn.Db()
}
