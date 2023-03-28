package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type sqlConfig struct {
	authUrl   string
	sqlDriver string
}

func NewSqlConn() *sqlx.DB {
	db, err := sqlx.Connect("mysql", os.Getenv("PLANETSCALE_DB_DEV"))
	if err != nil {
		panic(err)
		return nil
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Connected to sql DB")
	
	return db

}
