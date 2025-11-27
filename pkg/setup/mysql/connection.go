package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN         string
	MaxOpen     int
	MaxIdle     int
	MaxLifetime time.Duration
}

type DB struct {
	Conn *sql.DB
}

var (
	mysqlInstance *DB
	once          sync.Once
)

func NewConnection(cfg Config) (*DB, error) {
	var err error
	once.Do(func() {
		conn, err := sql.Open("mysql", cfg.DSN)
		if err != nil {
			err = fmt.Errorf("[MYSQL] failed to open connection: %w", err)
			return
		}

		conn.SetMaxOpenConns(cfg.MaxOpen)
		conn.SetMaxIdleConns(cfg.MaxIdle)
		conn.SetConnMaxLifetime(cfg.MaxLifetime)

		var version string
		if err := conn.QueryRow("SELECT VERSION()").Scan(&version); err != nil {
			conn.Close()
			err = fmt.Errorf("[MYSQL] failed to verify connection: %w", err)
			return
		}

		log.Printf("[MYSQL] connected")
		mysqlInstance = &DB{Conn: conn}
	})

	if err != nil {
		return nil, err
	}
	return mysqlInstance, nil
}

func (db *DB) Close() error {
	if db.Conn == nil {
		return nil
	}

	if err := db.Conn.Close(); err != nil {
		return fmt.Errorf("[MYSQL] failed to close connection: %w", err)
	}

	log.Println("[MYSQL] connection closed.")
	return nil
}
