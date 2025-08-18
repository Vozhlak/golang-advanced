package auth

import (
	"fmt"
	"github.com/Vozhlak/golang-advanced/3-validation-api/configs"
	"github.com/Vozhlak/golang-advanced/3-validation-api/pkg/res"
	"log"
	"net/http"
)

type DepsAuthHandler struct {
	*configs.Config
}

type HandlerAuth struct {
	*configs.Config
}

func NewHandlerAuth(router *http.ServeMux, deps DepsAuthHandler) {
	handler := &HandlerAuth{
		Config: deps.Config,
	}
	router.HandleFunc("POST /api/login", handler.Login())
	router.HandleFunc("POST /api/register", handler.Register())
}

func (handler *HandlerAuth) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LOGIN")
		fmt.Printf("SECRET: %s", handler.Config.Auth.Secret)
		response := LoginResponse{
			Token: "Token_123",
		}

		err := res.Json(w, response, 200)
		if err != nil {
			log.Println("error: ", err.Error())
		}
	}
}

func (handler *HandlerAuth) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("REGISTER")
	}
}
