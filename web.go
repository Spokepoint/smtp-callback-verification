package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func newHandler() rest.ResourceHandler {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/verify/", getEmailVerify},
	)
	return handler
}

func main() {
	handler := newHandler()
	http.ListenAndServe(":8080", &handler)
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
