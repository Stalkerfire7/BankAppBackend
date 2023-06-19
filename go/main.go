package main

import (
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}

//18:32 https://www.youtube.com/watch?v=pwZuNmAzaH8
