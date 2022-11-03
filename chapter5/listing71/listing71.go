package main

import (
	"fmt"

	"github.com/jeffcail/go-project-practice/chapter5/listing71/entities"
)

func main() {
	//u := entities.User{Name: "Bill", email: "bill@email.com"}
	u := entities.User{Name: "Bill"}
	fmt.Printf("User: %v\n", u)
}
