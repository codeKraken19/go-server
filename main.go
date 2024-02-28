package main

// imports
import (
	"fmt"
	"log"
	"net/http"

	"go-server/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	err = log.Fatal(http.ListenAndServe(":8080", r))
	if err (
		fmt.Println("Another process is running on 8080/tcp already.")
		)
}
