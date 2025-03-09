package main

import (
	"github.com/lo0kat/kerebro-api/internal/api"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"fmt"
	"os"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err!= nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool : %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var words string
	err = dbpool.QueryRow(context.Background(), "select kr_word from Words where en_word=$1","Hello").Scan(&words)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed %v\n", err)
		os.Exit(1)
	}

	fmt.Println(words)

	server := api.NewAPIServer(":8080", nil)
	server.Run()
}
