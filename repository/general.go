package repository

import (
	"database/sql"
	"voting/config"

	_ "github.com/go-sql-driver/mysql"
)

func connectDb() (*sql.DB, error) {
	driver := config.GetCServer().Db.Driver
	conString := config.GetCServer().Db.Username + ":@tcp(" + config.GetCServer().Db.Address + ":" +
		config.GetCServer().Db.Port + ")/" + config.GetCServer().Db.Name

	db, err := sql.Open(driver, conString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func prefixDb(name string) string {
	return config.GetCServer().Db.Prefix + name
}
