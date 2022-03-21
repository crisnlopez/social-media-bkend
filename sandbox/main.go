package main

import (
  "fmt"
  "log"
  db "github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
db, err := db.OpenDB("social_media")
if err != nil {
  log.Fatal(err)
}

user, err := db.CreateUser("text@example.com", "1324513", "Pepe", 20)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("User created: %+v\n",user)

user, err = db.UpdateUser("text@example.com", "12424", "Carlos", 34)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("User updated: %+v\n", user)

post, err := db.CreatePost("text@example.com", "Sample Text")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Post created: %+v\n", post)

post, err = db.CreatePost("text@example.com", "Sample two")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Another Post created: %+v\n", post)

fakePost, err := db.CreatePost("fake_email@example.com", "Fake post")
if err != nil {
  fmt.Printf("Error: %+v",err)
}
fmt.Printf("Fake post: %+v\n", fakePost)

posts, err := db.GetPost("text@example.com")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("Posts: %+v\n", posts)

for _, p := range posts {
    err = db.DeletePost(p.ID)
    if err != nil {
      log.Fatal(err)
    }
  }

err = db.DeleteUser("text@example.com")
if err != nil {
  log.Fatal(err)
}

}
