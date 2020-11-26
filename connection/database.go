package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DBCLIENT *sql.DB

func InitPostgresDB() {

	// Load environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to open database: %s", err))
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(fmt.Sprintf("Database connection error: %s", err))
	}

	DBCLIENT = db
}
