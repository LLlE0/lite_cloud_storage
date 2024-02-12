package dbworker

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func NewDBInstance() *sql.DB {
	sqlDB, err := sql.Open("sqlite", viper.GetString("db"))
	if err != nil {
		log.Print(err)
	}

	dbinit, err := os.ReadFile("../configs/db_init")
	if err != nil {
		log.Print(err)
	}
	if _, err = sqlDB.Exec(string(dbinit)); err != nil {
		log.Print(err)
	}
	log.Print("Successfully connxted to DB")
	return sqlDB

}
