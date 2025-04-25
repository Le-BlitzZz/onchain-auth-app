package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Le-BlitzZz/onchain-auth-app/internal/entity"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/mutex"
)

// SQL defaults to MySQL.
const (
	MySQL          = "mysql"
	DatabaseDriver = MySQL
)

func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("Config: database not connected")
	}

	return c.db
}

func (c *Config) InitDb() {
	entity.SetDbProvider(c)
	c.MigrateDb()
}

func (c *Config) MigrateDb() {
	c.db.AutoMigrate(entity.User{})
	entity.CreateUser(c.options.DefaultUser, c.options.DefaultPassword)
}

func (c *Config) connectDb() error {
	// Make sure this is not running twice.
	mutex.Db.Lock()
	defer mutex.Db.Unlock()

	// Get database data source name.
	dbDsn := c.DatabaseDsn()

	if dbDsn == "" {
		return errors.New("Config: database DSN not specified")
	}

	// Open database connection.
	db, err := gorm.Open(DatabaseDriver, dbDsn)
	if err != nil || db == nil {
		log.Info("Config: waiting for the database to become available")

		for i := 1; i <= 12; i++ {
			db, err = gorm.Open(DatabaseDriver, dbDsn)

			if db != nil && err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			return err
		}
	}

	log.Debug("Database: established connection")

	c.db = db

	return nil
}

// DatabaseDsn returns the database data source name (DSN).
func (c *Config) DatabaseDsn() string {
	if c.options.DatabaseDsn == "" {
		databaseServer := fmt.Sprintf("tcp(%s)", c.DatabaseServer())

		return fmt.Sprintf(
			"%s:%s@%s/%s?charset=utf8mb4,utf8&collation=utf8mb4_unicode_ci&parseTime=true&timeout=%ds",
			c.DatabaseUser(),
			c.DatabasePassword(),
			databaseServer,
			c.DatabaseName(),
			c.DatabaseTimeout(),
		)
	}

	return c.options.DatabaseDsn
}

// DatabaseServer returns the database server.
func (c *Config) DatabaseServer() string {
	if c.options.DatabaseServer == "" {
		return "localhost"
	}

	return c.options.DatabaseServer
}

// DatabaseName returns the database schema name.
func (c *Config) DatabaseName() string {
	if c.options.DatabaseName == "" {
		return "authonchain"
	}

	return c.options.DatabaseName
}

// DatabaseUser returns the database user name.
func (c *Config) DatabaseUser() string {
	if c.options.DatabaseUser == "" {
		return "authonchain"
	}

	return c.options.DatabaseUser
}

// DatabasePassword returns the database user password.
func (c *Config) DatabasePassword() string {
	return c.options.DatabasePassword
}

// DatabaseTimeout returns the TCP timeout in seconds for establishing a database connection:
func (c *Config) DatabaseTimeout() int {
	// Ensure that the timeout is between 1 and a maximum
	// of 60 seconds, with a default of 15 seconds.
	if c.options.DatabaseTimeout <= 0 {
		return 15
	} else if c.options.DatabaseTimeout > 60 {
		return 60
	}

	return c.options.DatabaseTimeout
}
