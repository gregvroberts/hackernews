package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	MigrateDB(pgxDB) // Migrate our db

	return sqlx.NewDb(pgxDB, "pgx"), nil
}

// MigrateDB migrates the database to the most recent migration up file(s)
func MigrateDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Print(err)
	}

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://app/infrastructure/db/migrations", "postgres", driver)
	if err != nil {
		fmt.Print(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
