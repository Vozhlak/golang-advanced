package main

import (
	"fmt"
	random_api "github.com/Vozhlak/golang-advanced/2-random-api"
	concurency "github.com/Vozhlak/golang-advanced/concurency_1"
	"log"
	"net/http"
)

func main() {
	concurency.Run()

	router := http.NewServeMux()
	random_api.NewHandlerApiRandom(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listen on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
