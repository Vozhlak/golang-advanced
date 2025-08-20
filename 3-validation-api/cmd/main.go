package main

import (
	"fmt"
	"github.com/Vozhlak/golang-advanced/3-validation-api/configs"
	"github.com/Vozhlak/golang-advanced/3-validation-api/internal/auth"
	"github.com/Vozhlak/golang-advanced/3-validation-api/internal/hello"
	"github.com/Vozhlak/golang-advanced/3-validation-api/internal/verify"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Happy coding!")

	conf := configs.LoadConfig()

	//Server
	router := http.NewServeMux()
	hello.NewHandlerHello(router)
	auth.NewHandlerAuth(router, auth.DepsAuthHandler{
		Config: conf,
	})
	verify.NewHandlerVerify(router, verify.DepsVerifyHandler{
		Config: conf,
	})

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
