package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	mysqldump "github.com/JamesStewy/go-mysqldump"
	mysql "github.com/go-sql-driver/mysql"
)

// DB is a wrapper for the sql.DB
type DB struct {
	db *sql.DB
}

var (
	dbDriver            = "mysql"
	user                = "reeloUser"
	password            = "password"
	host                = "localhost:3306"
	dbName              = "reelo"
	bkpDir              = "bkp"
	maxConnections      = 5
	maxIdleConnections  = 5
	maxConnTries        = 10
	connectionsLifetime = (time.Minute * 5)
	instanceEsists      = false
)

var dbInstance *DB

func init() {
	if os.Getenv("ENV") == "prod" {
		dbDriver = os.Getenv("DB_DRIVER")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host = os.Getenv("DB_HOST")
		dbName = os.Getenv("DB_NAME")
		bkpDir = os.Getenv("DB_BKP_PATH")
	}
	log.Println("DB initialized")
	log.Println(dbDriver, user, host, dbName)
	log.Printf("DB backup dir: %s", bkpDir)

	dbInstance = newDB()
}

// Instance returns the current database instance if it exits.
// Otherwise a new instance will be created.
func Instance() *DB {
	if instanceEsists {
		return dbInstance
	}
	dbInstance = newDB()
	return dbInstance
}

// Close closes the database
func (database *DB) Close() {
	database.db.Close()
	instanceEsists = false
}

// newDB creates a connection to the database, checks for connection integrity
// by pinging the address. If the first N connection tries fails then the call
// fails.
func newDB() *DB {
	dbConfig := mysql.NewConfig()
	dbConfig.User = user
	dbConfig.Passwd = password
	dbConfig.Addr = host
	dbConfig.DBName = dbName
	dbConfig.Net = "tcp"
	dataSourceName := dbConfig.FormatDSN()
	var tries int

	db, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		log.Printf("Error opening the database: %s", err)
	}
	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	// db.SetConnMaxLifetime(connectionsLifetime)
	database := DB{
		db: db,
	}

	for tries <= maxConnTries {
		if err := db.PingContext(context.Background()); err != nil {
			if tries == maxConnTries {
				log.Printf("Error connecting to the database: %s", err)
			}
			tries++
			time.Sleep(time.Duration(tries*100) * time.Millisecond)
		} else {
			break
		}
	}
	instanceEsists = true
	return &database

}

// Backup exports the database into a file, it also closes the connection
// to the database. So you have to reopen it.
func (database *DB) Backup() (resultFilename string) {
	// accepts time layout string and add .sql at the end of file
	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", dbName)

	// Register database with mysqldump
	dumper, err := mysqldump.Register(database.db, bkpDir, dumpFilenameFormat)
	if err != nil {
		log.Printf("Error registering databse: %v", err)
		return resultFilename
	}
	// Close dumper and connected database
	defer dumper.Close()

	// Dump database to file
	resultFilename, err = dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return resultFilename
	}
	log.Printf("File is saved to %s", resultFilename)
	return resultFilename
}
