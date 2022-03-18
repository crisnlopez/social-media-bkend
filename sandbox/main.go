package main

import (
  "fmt"
  "log"
  db "github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
dbClient, err := db.NewClient("social_media")
if err != nil {
  log.Fatal(err)
}

user, err := dbClient.CreateUser("text@example.com", "1324513", "Pepe", 20)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("User created: %+v\n",user)

user, err = dbClient.UpdateUser("text@example.com", "12424", "Carlos", 34)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("User updated: %+v\n", user)

post, err := dbClient.CreatePost("text@example.com", "Sample Text")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Post created: %+v\n", post)

post, err = dbClient.CreatePost("text@example.com", "Sample two")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Another Post created: %+v\n", post)

fakePost, err := dbClient.CreatePost("fake_email@example.com", "Fake post")
if err != nil {
  fmt.Printf("Error: %+v",err)
}
fmt.Printf("Fake post: %+v\n", fakePost)

posts, err := dbClient.GetPost("text@example.com")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Posts: %+v\n", posts)

for _, p := range posts {
    err = dbClient.DeletePost(p.ID)
    if err != nil {
      log.Fatal(err)
    }
  }

err = dbClient.DeleteUser("text@example.com")
if err != nil {
  log.Fatal(err)
}

}
