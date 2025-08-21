package main

import (
	"fmt"
	"github.com/Vozhlak/golang-advanced/4-order-api/configs"
	"github.com/Vozhlak/golang-advanced/4-order-api/pkg/db"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Happy coding!")

	conf := configs.LoadConfig()

	//Server
	router := http.NewServeMux()
	_ = db.NewDb(conf)

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
