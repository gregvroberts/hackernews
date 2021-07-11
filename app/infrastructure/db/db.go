package db

import (
	"fmt"
	"time"

	"github.com/gregvroberts/hackernews/app/infrastructure/util"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func OpenSqlxDB(dbConfig *util.Config) (*sqlx.DB, error) {
	connConfig := pgx.ConnConfig{
		Host:     dbConfig.DBHost,
		Port:     dbConfig.DBPort,
		Database: dbConfig.DBName,
		User:     dbConfig.DBUser,
		Password: dbConfig.DBPassword,
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 40,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		fmt.Printf("unable to establish pgx Connection Pool %s", err)
		return nil, err
	}

	opts := make([]stdlib.OptionOpenDBFromPool, 0)

	pgxDB := stdlib.OpenDBFromPool(connPool, opts...)

	return sqlx.NewDb(pgxDB, "pgx"), nil
}
