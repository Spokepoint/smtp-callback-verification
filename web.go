package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os"
)

func newHandler() rest.ResourceHandler {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/verify/", getEmailVerify},
	)
	return handler
}

func main() {
	PORT := os.Getenv("PORT")

	handler := newHandler()

	log.Println("Listening on port " + PORT)
	err := http.ListenAndServe(":"+PORT, &handler)
	if err != nil {
		log.Panic(err)
	}
}

type Message struct {
	Email   string `json:"email"`
	IsValid bool   `json:"isValid"`
}

func getEmailVerify(w rest.ResponseWriter, req *rest.Request) {
	email := req.URL.Query().Get("email")
	valid, _ := VerifyEmail(email)

	w.WriteJson(&Message{
		Email:   email,
		IsValid: valid,
	})
}
