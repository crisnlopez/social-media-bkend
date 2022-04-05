package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
  Db *sql.DB
}

// Take the database name and return a DB reference to data source
func OpenDB(dbName string) (*sql.DB, error) {
  // DB Connection config
  cfg := mysql.Config{
    User: os.Getenv("DBUSER"),
    Passwd: os.Getenv("DBPASS"),
    Net: "tcp",
    Addr: "127.0.0.1:3306",
    DBName: dbName,
    AllowNativePasswords: true,
    ParseTime: true,
  }

  db, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    return nil, err
  }

  pingErr := db.Ping()
  if pingErr != nil {
    log.Fatal(pingErr)
  }

	return db, nil
}
