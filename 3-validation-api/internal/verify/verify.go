package verify

import (
	"fmt"
	"github.com/Vozhlak/golang-advanced/3-validation-api/configs"
	"net/http"
)

type DepsVerifyHandler struct {
	Config *configs.Config
}

type HandlerVerify struct {
	Config *configs.Config
}

func NewHandlerVerify(router *http.ServeMux, deps DepsVerifyHandler) {
	handler := &HandlerVerify{
		Config: deps.Config,
	}
	router.HandleFunc("POST /api/send", handler.Send())
	router.HandleFunc("GET /api/verify/{hash}", handler.Verify())
}

func (handler *HandlerVerify) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send")
	}
}

func (handler *HandlerVerify) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify", r)
	}
}
