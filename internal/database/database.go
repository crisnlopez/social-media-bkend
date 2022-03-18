package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Client struct {
  db *sql.DB
}

// Take the database name and return a Client instance
func NewClient(dbName string) (Client, error) {
  // DB Connection config
  cfg := mysql.Config{
    User: os.Getenv("DBUSER"),
    Passwd: os.Getenv("DBPASS"),
    Net: "tcp",
    Addr: "127.0.0.1:3306",
    DBName: dbName,
    AllowNativePasswords: true,
  }

  DbConnection, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    return Client{&sql.DB{}}, err
  }

  pingErr := DbConnection.Ping()
  if pingErr != nil {
    log.Fatal(pingErr)
  }

	return Client{DbConnection}, nil
}
