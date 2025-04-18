package main

import (
	"fmt"
	"log"

	"github.com/mrjxtr-dev/mr-shush/internal/generator"
)

func main() {
	gen := generator.New()

	password, err := gen.GeneratePassword(12, "strong")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password:", password)
}
