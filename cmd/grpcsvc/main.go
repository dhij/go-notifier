package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/dhij/go-notifier/internal/grpcsvc"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		dbDriver = "mysql"
		// if you are running the mysql container locally without docker-compose, use dbSource = "root:password@tcp(localhost:33060)/notifier_db"
		dbSource = "root:password@tcp(mysql:3306)/notifier_db"
	)

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
