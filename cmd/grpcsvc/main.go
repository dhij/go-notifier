package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/dhij/go-notifier/internal/grpcsvc"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "root:password@tcp(localhost:33060)/notifier_db"
)

func main() {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("connecting to MySQL: %s", err)
	}
	defer db.Close()

	err = grpcsvc.Main(context.Background(), db)
	if err != nil {
		log.Fatalf("encountered unexpected error: %s", err)
	}
}
