package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	mysql "github.com/go-sql-driver/mysql"
)

var (
	dbDriver = "mysql"
	user     = "reeloUser"
	password = "password"
	host     = "localhost:3306"
	dbName   = "reelo"
)

func init() {
	if os.Getenv("ENV") == "prod" {
		dbDriver = os.Getenv("DB_DRIVER")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host = os.Getenv("DB_HOST")
		dbName = os.Getenv("DB_NAME")
	}
	log.Println("DB initialized")
	log.Println(dbDriver, user, host, dbName)
}

// DB is a wrapper for the sql.DB
type DB struct {
	db *sql.DB
}

// Close closes the database
func (database *DB) Close() {
	database.db.Close()
}

// NewDB returns the databse used for this program.
// REMEMBER TO CLOSE IT!
func NewDB() *DB {
	dbConfig := mysql.NewConfig()
	dbConfig.User = user
	dbConfig.Passwd = password
	dbConfig.Addr = host
	dbConfig.DBName = dbName
	dbConfig.Net = "tcp"
	dataSourceName := dbConfig.FormatDSN()

	db, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		log.Fatalf("Error opening the database: %s", err)
	}
	database := DB{
		db: db,
	}
	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	return &database
}
