package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool
var dbErr error

func Init() {

	// Init Db connection
	dbPool, dbErr = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", dbErr)
		os.Exit(1)
	}

	pingErr := dbPool.Ping(context.Background())
	if pingErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect on db: %v\n", pingErr)
		os.Exit(1)
	}

	log.Println("Connected to the database successfully!")
}

func Close() {
	dbPool.Close()
}

func CreateTables(tablesDdl []string) {
	for i := range tablesDdl {
		_, err := dbPool.Exec(context.Background(), tablesDdl[i])

		if err != nil {
			log.Fatal(err)
		}
	}
}
