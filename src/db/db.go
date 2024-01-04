package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/volatiletech/sqlboiler/boil"
)

func NewDB() *bun.DB {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))

		sqldb.SetMaxIdleConns(10)
		sqldb.SetMaxOpenConns(50)
		sqldb.SetConnMaxLifetime(300 * time.Second)
		boil.SetDB(sqldb)

		db := bun.NewDB(sqldb, pgdialect.New())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}
