package db

import(
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

var Pool *pgxpool.Pool

func Connect(connectionStr string){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	Pool, err = pgxpool.New(ctx,connectionStr)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n, err")
	}

	if err = Pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database : %v\n", err)
	}

	log.Println("Connected to the database successfully")
}