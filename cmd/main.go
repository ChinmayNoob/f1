package main

import (
	"context"
	"log"

	"github.com/ChinmayNoob/f1/pkg/db"
)

func main() {
	ctx := context.Background()
	db.InitDB(ctx)

	var greeting string
	err := db.Conn.QueryRow(ctx, "SELECT 'Hello from Postgres via Go!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println(greeting)
}
