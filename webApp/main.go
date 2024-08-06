package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shivangidas/go-to-do-app/webApp/handler"
)

func main() {
	handler.Setup()
	handler.Handlers()
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
