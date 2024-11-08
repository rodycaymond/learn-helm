package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DbConnect() *sql.DB {
	if db != nil {
		return db
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, name, password, host, port)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to database")
		fmt.Println(fmt.Errorf("error during database connection: %+v", err))
		db = nil
		os.Exit(1)
	}
	fmt.Println("#############")
	fmt.Printf("This is the error returned from connection: %+v \n", err)
	fmt.Println("#############")

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS test(name VARCHAR(255))")
	if err != nil {
		fmt.Println("Failed to create table 'test'")
		fmt.Println(fmt.Sprintf("error: %+v", err))
	}
	_, err = conn.Exec("INSERT INTO test(name) VALUES('cody')")
	if err != nil {
		fmt.Println("failed to insert into 'test'")
		fmt.Println(fmt.Sprintf("error: %+v", err))
	}

	db = conn
	return db
}
