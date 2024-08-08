package main

import (
	"fmt"
	"log"
	"net/http"

	//api "github.com/shivangidas/go-to-do-app/webApp/apiwithconcurrency"
	"github.com/shivangidas/go-to-do-app/webApp/handler"
)

func startRegularServer() {
	handler.InjectData()
	handler.Handlers()
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	startRegularServer()
	//api.StartServer()
}
