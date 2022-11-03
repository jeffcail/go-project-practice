package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	d := duration(42)
	d.pretty()
}
