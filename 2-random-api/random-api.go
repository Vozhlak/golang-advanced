package random_api

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type HandlerRandomApi struct{}

func NewHandlerApiRandom(router *http.ServeMux) {
	handler := &HandlerRandomApi{}
	router.HandleFunc("/api/random", handler.getNumber)
}

func (handler *HandlerRandomApi) getNumber(w http.ResponseWriter, r *http.Request) {
	num := getRandomNumber()

	w.Header().Set("Content-Type", "text/plain")

	_, err := fmt.Fprintf(w, "%d", num)
	if err != nil {
		log.Printf("Ошибка при отправке ответа: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func getRandomNumber() int {
	return rng.Intn(6) + 1
}
