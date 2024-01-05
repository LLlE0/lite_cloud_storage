package dbworker

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewDBInstance() *sql.DB {
	sqlDB, err := sql.Open("sqlite3", viper.GetString("db"))
	if err != nil {
		log.Fatal(err)
	}

	dbinit, err := os.ReadFile("../configs/db_init")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = sqlDB.Exec(string(dbinit)); err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connxted to DB")
	return sqlDB

}
