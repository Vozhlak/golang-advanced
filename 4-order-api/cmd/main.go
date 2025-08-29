package main

import (
	"fmt"
	"github.com/Vozhlak/golang-advanced/4-order-api/configs"
	"github.com/Vozhlak/golang-advanced/4-order-api/internal/product"
	"github.com/Vozhlak/golang-advanced/4-order-api/middleware"
	"github.com/Vozhlak/golang-advanced/4-order-api/pkg/db"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Happy coding!")

	conf := configs.LoadConfig()

	//Server
	router := http.NewServeMux()

	//DataBase
	database := db.NewDb(conf)

	//Repository
	repositoryProduct := product.NewRepositoryProduct(database)

	//Handlers
	product.NewHandlerProduct(router, product.HandlerProductDeps{
		ProductRepository: repositoryProduct,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listen on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
