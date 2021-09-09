package main

import (
	"WebDevelopment/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application stating at portnumber %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
