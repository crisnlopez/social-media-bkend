package database

import (
	"encoding/json"
	"os"
)

type Client struct {
	DbPath string
}

type databaseScheme struct {
	Users map[string]User `json:"users"`
	Post  map[string]Post `json:"posts"`
}

// Take the file path to the database and return a Client instance
func NewClient(dbpath string) Client {
	return Client{DbPath: dbpath}
}

// Check if a database already exists. If it does, all is good. Otherwise, create a new database using Client.DBPath.
func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.DbPath)

	if err != nil {
		err = c.createDB()
		return err
	}

	return err
}

// Create a new database using Client.DBPath.
func (c Client) createDB() error {
	dat, err := json.Marshal(databaseScheme{ // Empty database instance
		Users: make(map[string]User),
		Post:  make(map[string]Post),
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(c.DbPath, dat, 0666)
	if err != nil {
		return err
	}

	return err
}
