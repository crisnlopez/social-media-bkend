package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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

	return nil
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

	return nil
}

func (c Client) updateDB(db databaseScheme) error {
	err := c.EnsureDB()
	if err != nil {
		return err
	}

	dat, err := json.Marshal(databaseScheme{
		Users: db.Users,
		Post:  db.Post,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(c.DbPath, dat, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) readDB() (databaseScheme, error) {
	newDb := databaseScheme{}

	data, err := os.ReadFile(c.DbPath)
	if err != nil {
		return databaseScheme{}, err
	}

	err = json.Unmarshal(data, &newDb)
	if err != nil {
		return databaseScheme{}, err
	}

	return newDb, nil
}

// Users

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	// New user instance
	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}

	db.Users[email] = user // User email as "Primery Key" cause we can´t have two users with the same email
	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	// Check if this user exist
	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("User doesn´t exist")
	}

	// Update fields
	user.Password = password
	user.Name = name
	user.Age = age
	db.Users[email] = user

	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	_, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("User doesn´t exist")
	}

	return db.Users[email], nil
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	delete(db.Users, email)

	err = c.updateDB(db)
	if err != nil {
		return err
	}

	return nil
}
