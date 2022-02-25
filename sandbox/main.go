package main

import (
	"fmt"
	"log"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
	c := db.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	user, err := c.CreateUser("text@example.com", "password", "Pepe Cuenca", 34)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created", user)

	updateUser, err := c.UpdateUser("text@example.com", "passsssword", "Pepe", 34)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User update", updateUser)

	gotUser, err := c.GetUser("text@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gotUser)

	err = c.DeleteUser("text@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted")

	_, err = c.GetUser("text@example.com")
	if err == nil {
		log.Fatal("Shouldn´t be able to get user")
	}
	fmt.Println("user confirmed deleted")

	user, err = c.CreateUser("test@example.com", "password", "Pepe Cuenca", 34)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user recreated", user)

	post, err := c.CreatePost("test@example.com", "This is a post test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post created", post)

	secondPost, err := c.CreatePost("test@example.com", "Another test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Another post created", secondPost)

	posts, err := c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Posts:", posts)

	err = c.DeletePost(post.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted first post")

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Posts:", posts)

	err = c.DeletePost(secondPost.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted second post")

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Posts:", posts)

	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User redeleted")
}
