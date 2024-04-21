package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"

	"github.com/pkg/errors"
)

var db *pgx.Conn

func Connect() error {
	var err error
	fmt.Println(os.Getenv("EVOLVER_MAIN_DATABASE_CONNECTION_STRING"))
	db, err = pgx.Connect(context.Background(), os.Getenv("EVOLVER_MAIN_DATABASE_CONNECTION_STRING"))
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func Close() error {
	if err := db.Close(context.Background()); err != nil {
		return errors.Wrap(err, "Could not close the database connection")
	}
	return nil
}
