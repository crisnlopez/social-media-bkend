package database

import (
	"fmt"

	"github.com/google/uuid"
)

type Post struct {
  ID string `json:"id"`
  UserEmail string `json:"userEmail"`
  Text string `json:"text"`
}

// CreatePost needs userEmail to check if user exists and text for post
func (c Client) CreatePost(userEmail, text string) (Post, error) {
  // Check if user exist 
  var emptyUser User
  user, err := c.GetUser(userEmail)
  if err != nil {
    return Post{}, err
  }
  if user == emptyUser {
    return Post{}, fmt.Errorf("User doesn't exists")
  }

  // NewPost instance
  post := Post{
    ID: uuid.New().String(),
    UserEmail: userEmail,
    Text: text,
  }

  // Create Post in Database
  result, err := c.db.Exec("insert into posts (id, userEmail, postText) values (?, ?, ?)", post.ID, post.UserEmail, post.Text)
  if err != nil {
    return Post{}, err
  }
  result.RowsAffected()

  return post, nil
}

// Returns all Posts from the given userEmail
func (c Client) GetPost(userEmail string) ([]Post, error){
  var posts []Post

  rows, err := c.db.Query("select * from posts where userEmail = ?", userEmail)
  if err != nil {
    return nil, fmt.Errorf("Get posts by userEmail %q: %v", userEmail, err)
  }

  defer rows.Close()
  // Loop through rows, using Scan to assing column data to struct fields
  for rows.Next() {
    var post Post
    if err := rows.Scan(&post.ID, &post.UserEmail, &post.Text); err != nil {
      return nil, fmt.Errorf("Get posts by id %q: %v", userEmail, err)
    }
    posts = append(posts, post)
  }

  // Checks for errors
  if err := rows.Err(); err != nil {
    return nil, fmt.Errorf("Get posts by id %q: %v", userEmail, err)
  }

  return posts, nil
}

// Delete Post with the given id
func (c Client) DeletePost(id string) error {
  result, err := c.db.Exec("delete from posts where id = ?", id)
  if err != nil {
    return err
  }

  result.RowsAffected()
  return nil
}
