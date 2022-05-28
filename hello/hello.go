package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	//message := greetings.Hello("Gladys")
	message, err := greetings.Hello("jonny")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
