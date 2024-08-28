package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connect() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file %v", err)
	}
	connStr := os.Getenv("DATABASE_URL")

	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("unable to connect to database :%v\n", err)
	}
	_, err = DB.Exec(context.Background(), "SET search_path TO new_billing_test")
	if err != nil {
		return fmt.Errorf("failled to search path :%v", err)
	}

	log.Println("connected to the database ")
	return nil

}

func RunTables() {
	conn, err := DB.Acquire(context.Background())
	if err != nil {
		log.Fatalf("unable to connect %v\n", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS new_billing_test.users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS new_billing_test.transactions (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		transaction_id TEXT NOT NULL UNIQUE,
		amount NUMERIC(10,2) NOT NULL,
		currency TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS new_billing_test.payment_methods (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		card_number TEXT NOT NULL,
		card_expiry TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)

	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	} else {

		log.Println("migration was successful")
	}

}

func Close() {
	DB.Close()
	log.Println("closed the database connection")
}
