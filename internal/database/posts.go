package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDB()
	if err != nil {
		return Post{}, err
	}

	_, ok := db.Users[userEmail]
	if !ok {
		return Post{}, errors.New("the user doesn´t exist")
	}

	newPost := Post{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC().Local(),
		UserEmail: userEmail,
		Text:      text,
	}
	db.Post[newPost.ID] = newPost

	err = c.updateDB(db)
	if err != nil {
		return Post{}, err
	}

	return newPost, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		return []Post{}, err
	}

	_, ok := db.Users[userEmail]
	if !ok {
		return []Post{}, errors.New("the user doesn´t exist")
	}

	var posts []Post
	for _, v := range db.Post {
		if v.UserEmail == userEmail {
			posts = append(posts, v)
		}
	}

	return posts, nil
}

func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	delete(db.Post, id)

	err = c.updateDB(db)
	if err != nil {
		return err
	}

	return nil
}
