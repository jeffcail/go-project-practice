package main

import (
	"log"
	"os"

	_ "github.com/jeffcail/go-project-practice/chapter2/matchers"
	"github.com/jeffcail/go-project-practice/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
