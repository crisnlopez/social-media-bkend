package main

import (
	"fmt"
	"log"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
	client := db.NewClient("db.json")
	err := client.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	user, err := client.CreateUser("text@example.com", "password", "Pepe Cuenca", 34)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created", user)

	updateUser, err := client.UpdateUser("text@example.com", "passsssword", "Pepe", 34)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User update", updateUser)

	gotUser, err := client.GetUser("text@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gotUser)

	err = client.DeleteUser("text@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted")

	_, err = client.GetUser("text@example.com")
	if err == nil {
		log.Fatal("Shouldn´t be able to get user")
	}
	fmt.Println("user confirmed deleted")
}
