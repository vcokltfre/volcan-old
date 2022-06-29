package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() error {
	db_dsn := os.Getenv("DB_DSN")

	var db *gorm.DB
	var err error

	if strings.HasPrefix(db_dsn, "postgresql://") {
		logrus.Info("Connecting to PostgreSQL database.")
		db, err = gorm.Open(postgres.Open(db_dsn), &gorm.Config{})
	} else if strings.HasPrefix(db_dsn, "mysql://") {
		logrus.Info("Connecting to MySQL database.")
		db, err = gorm.Open(mysql.Open(db_dsn), &gorm.Config{})
	} else {
		return fmt.Errorf("DB_DSN mut be a valid Postgres or MySQL connection string.")
	}

	if err != nil {
		return err
	}

	logrus.Info("Connected to the database, running migrations...")

	err = db.AutoMigrate()
	if err != nil {
		return err
	}

	logrus.Info("Migrations complete.")

	DB = db

	return nil
}
