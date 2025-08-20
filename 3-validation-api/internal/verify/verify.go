package verify

import (
	"fmt"
	"github.com/Vozhlak/golang-advanced/3-validation-api/configs"
	"github.com/Vozhlak/golang-advanced/3-validation-api/pkg/req"
	"github.com/Vozhlak/golang-advanced/3-validation-api/pkg/res"
	"github.com/jordan-wright/email"
	"log"
	"net/http"
	"net/smtp"
	"net/textproto"
	"time"
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
		body, err := req.HandleBody[RequestVerify](&w, r)
		if err != nil {
			return
		}

		hash, errGenerateHash := generateRandomHash()
		if errGenerateHash != nil {
			http.Error(w, fmt.Sprintf("internal error: %s", errGenerateHash.Error()), http.StatusInternalServerError)
			return
		}

		verifyURL := fmt.Sprintf("%s/api/verify/%s", handler.Config.Server.BaseUrl, hash)
		htmlBody, err := renderVerifyEmail(verifyURL)
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		e := &email.Email{
			To:      []string{body.Email},
			From:    "Mr. Nikto <vozlaki@gmail.com>",
			Subject: "Awesome Subject",
			Text:    []byte("Text Body is, of course, supported!"),
			HTML:    []byte(htmlBody),
			Headers: textproto.MIMEHeader{},
		}

		err = e.Send(handler.Config.Email.SMTPHost+":"+handler.Config.Email.SMTPPort, smtp.PlainAuth("", handler.Config.Email.Address, handler.Config.Email.Password, handler.Config.Email.SMTPHost))
		if err != nil {
			log.Printf("Failed to send email: %v", err)
			res.Json(w, map[string]bool{"success": false}, http.StatusInternalServerError)
			return
		}

		storeVerificationHash(hash, body.Email, time.Hour*24)

		res.Json(w, map[string]bool{"success": true}, http.StatusOK)
	}
}

func (handler *HandlerVerify) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		_, isExists := getEmailByHash(hash)

		if !isExists {
			res.Json(w, map[string]bool{"success": false}, http.StatusOK)
			return
		}

		deleteHash(hash)
		res.Json(w, map[string]bool{"success": true}, http.StatusOK)
	}
}
