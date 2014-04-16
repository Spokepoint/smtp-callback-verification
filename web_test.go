package main

import (
	"github.com/ant0ine/go-json-rest/rest/test"
	"testing"
)

func TestGetEmailVerify(t *testing.T) {
	handler := newHandler()

	type testpair struct {
		email string
		valid bool
	}

	var tests = []testpair{
		{"paul@gmail.com", true},
		{"hugo.martin@latimes.com", true},
		{"someBadEmail33ff92999g99f9f9f@example.net", false},
		{"not%20an%20email", false},
	}

	for _, pair := range tests {
		recorded := test.RunRequest(t, &handler,
			test.MakeSimpleRequest("GET",
				"http://1.2.3.4/verify/?email="+pair.email, nil))
		recorded.CodeIs(200)
		recorded.ContentTypeIsJson()

		// check actual response content
		msg := Message{}
		recorded.DecodeJsonPayload(&msg)
		if msg.IsValid != pair.valid {
			t.Errorf("For %v. Expected %v: Actual %v", pair.email, pair.valid, msg.IsValid)
		}
	}
}
