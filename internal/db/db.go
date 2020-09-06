package db

import (
	"errors"
	"strconv"
	"time"

	"github.com/diegogomesaraujo/fund-manager-api/internal/config"
	"github.com/diegogomesaraujo/fund-manager-api/internal/crypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const cryptPasswordPassphrase = "database"

var dbGorm *gorm.DB

// Open the connection with database
func Open(config *config.Config) {
	password := decryptDbUserPassword(config.Database.Password)

	host := config.Database.Host + ":" + strconv.FormatInt(config.Database.Port, 10)
	url := config.Database.Username + ":" + password + "@tcp(" + host + ")/" + config.Database.Database

	var err error
	dbGorm, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	db, err := dbGorm.DB()
	if err != nil {
		panic(err.Error)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
}

// Close the connection with database
func Close() {
	db, err := dbGorm.DB()
	if err != nil {
		panic(err.Error)
	}

	db.Close()
}

// Ping database
func Ping() {
	db, err := dbGorm.DB()
	if err != nil {
		panic(err.Error)
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error)
	}
}

// GetDb return a database connection
func GetDb() (*gorm.DB, error) {
	if dbGorm == nil {
		return dbGorm, errors.New("No database connections open")
	}

	return dbGorm, nil
}

// EncryptDbUserPassword encrypt the database user password
func EncryptDbUserPassword(password string) string {
	return crypt.EncryptAsString([]byte(password), cryptPasswordPassphrase)
}

func decryptDbUserPassword(passwordEncrypted string) string {
	passwordBytes := crypt.DencryptFromString(passwordEncrypted, cryptPasswordPassphrase)

	return string(passwordBytes)
}
