package db

import (
	"database/sql"
	"fmt"
	"homework/config"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
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

	db := bun.NewDB(sqldb, pgdialect.New())
	// db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	fmt.Println("DB Connceted")
	return db
}
