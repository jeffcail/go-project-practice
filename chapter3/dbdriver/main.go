package main

import (
	"database/sql"

	_ "github.com/jeffcail/go-project-practice/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
