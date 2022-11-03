package main

import (
	"fmt"

	"github.com/jeffcail/go-project-practice/chapter5/listing64/counters"
)

func main() {
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
