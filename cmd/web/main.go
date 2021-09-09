package main

import (
	"fmt"
	"net/http"

	"github.com/sudhirphogat/WebDevelopment_2/pkg/handler"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Application stating at portnumber %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
