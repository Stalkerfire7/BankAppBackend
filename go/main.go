package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
	//server := NewAPIServer(":3000", store)
	//server.Run()
}

//18:32 https://www.youtube.com/watch?v=pwZuNmAzaH8
