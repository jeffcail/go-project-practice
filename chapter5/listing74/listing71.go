package main

import (
	"fmt"

	"github.com/jeffcail/go-project-practice/chapter5/listing74/entities"
)

func main() {
	a := entities.Admin{Rights: 10}
	a.Name = "bill"
	a.Email = "bill@email.com"
	fmt.Printf("User: %v\n", a)
}
