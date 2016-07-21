package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Main function called")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8086", router))
}
