package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/diegogomesaraujo/fund-manager-api/internal/config"
	"github.com/diegogomesaraujo/fund-manager-api/internal/crypt"

	// import mysql drive to create a connection with mysql database
	_ "github.com/go-sql-driver/mysql"
)

const cryptPasswordPassphrase = "database"

var db *sql.DB

// Open the connection with database
func Open(config *config.Config) {
	password := decryptDbUserPassword(config.Database.Password)

	host := config.Database.Host + ":" + strconv.FormatInt(config.Database.Port, 10)
	url := config.Database.Username + ":" + password + "@tcp(" + host + ")/" + config.Database.Database

	var err error
	if db, err = sql.Open("mysql", url); err != nil {
		log.Fatalln(err)
		panic(err.Error)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

// Close the connection with database
func Close() {
	if err := db.Close(); err != nil {
		panic(err.Error)
	}
}

// GetDb return a database connection
func GetDb() (*sql.DB, error) {
	if db == nil {
		return db, errors.New("No database connections open")
	}

	return db, nil
}

// EncryptDbUserPassword encrypt the database user password
func EncryptDbUserPassword(password string) string {
	return crypt.EncryptAsString([]byte(password), cryptPasswordPassphrase)
}

func decryptDbUserPassword(passwordEncrypted string) string {
	passwordBytes := crypt.DencryptFromString(passwordEncrypted, cryptPasswordPassphrase)

	return string(passwordBytes)
}
