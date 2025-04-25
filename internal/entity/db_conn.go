package entity

import (
	"github.com/jinzhu/gorm"
	"os"
	"sync"
	"time"
)

// dbConn is the global gorm.DB connection provider.
var dbConn Gorm

// Gorm is a gorm.DB connection provider interface.
type Gorm interface {
	Db() *gorm.DB
}

// DbConn is a gorm.DB connection provider
type DbConn struct {
	Driver string
	Dsn    string

	once sync.Once
	db   *gorm.DB
}

// Db returns the gorm db connection.
func (d *DbConn) Db() *gorm.DB {
	d.once.Do(d.Open)

	if d.db == nil {
		log.Fatal("Migrate: database not connected.")
	}

	return d.db
}

// Open create a new gorm db connection
func (d *DbConn) Open() {
	db, err := gorm.Open(d.Driver, d.Dsn)

	if err != nil || db == nil {
		for i := 1; i <= 12; i++ {
			log.Printf("gorm.Open(%s, %s) %d\n", d.Driver, d.Dsn, i)
			db, err = gorm.Open(d.Driver, d.Dsn)

			if db != nil && err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			log.Info(err)
			os.Exit(0)
		}
	}

	db.LogMode(false)
	db.SetLogger(log)
	db.DB().SetMaxIdleConns(4)
	db.DB().SetMaxOpenConns(256)

	d.db = db
}

func (d *DbConn) Close() {
	if d.db != nil {
		if err := d.db.Close(); err != nil {
			log.Fatal(err)

			d.db = nil
		}
	}
}

func SetDbProvider(conn Gorm) {
	dbConn = conn
}
