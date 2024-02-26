package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tasuke/go-onion/config"
	"github.com/tasuke/go-onion/infrastructure/db/dbgen"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const maxRetries = 5
const delay = 5 * time.Second

var (
	once  sync.Once
	query *dbgen.Queries
	dbcon *sql.DB
)

func GetQuery(ctx context.Context) *dbgen.Queries {
	return query
}

func SetQuery(q *dbgen.Queries) {
	query = q
}

func GetDB() *sql.DB {
	return dbcon
}

func SetDB(d *sql.DB) {
	dbcon = d
}

func NewMainDB(cnf config.DBConfig) {
	once.Do(func() {
		dbcon, err := connect(
			cnf.User,
			cnf.Password,
			cnf.Host,
			cnf.Port,
			cnf.Name,
		)
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})
}

func connect(user string, password string, host string, port string, name string) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			return nil, fmt.Errorf("データベースを開けませんでした: %w", err)
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("データベースに接続できませんでした: %v", err)
		log.Printf("%v秒後に再試行します...", delay/time.Second)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("%d回の試行後にデータベースに接続できませんでした", maxRetries)
}
