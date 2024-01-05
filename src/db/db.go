package db

import (
	"database/sql"
	"fmt"
	"homework/config"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/volatiletech/sqlboiler/boil"
)

func NewDB(c config.Config) *bun.DB {	
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
		c.PostgresUser,
		c.PostgresPW,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDB)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))

	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(50)
	sqldb.SetConnMaxLifetime(300 * time.Second)
	boil.SetDB(sqldb)

	db := bun.NewDB(sqldb, pgdialect.New())
	fmt.Println("Connceted")
	return db
}
