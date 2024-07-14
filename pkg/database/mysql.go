package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type DB interface {
	PingContext(ctx context.Context) error
	Close() error
}

type mysqlDB struct {
	db *sql.DB
}

var (
	dbInstance *mysqlDB
)

func NewMySQLDB() (DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbInstance = &mysqlDB{db: db}
	return dbInstance, nil
}

func (m *mysqlDB) PingContext(ctx context.Context) error {
	return m.db.PingContext(ctx)
}

func (m *mysqlDB) Close() error {
	return m.db.Close()
}
