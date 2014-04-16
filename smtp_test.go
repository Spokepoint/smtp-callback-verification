package main

import (
	"testing"
)

type testpair struct {
	email string
	valid bool
}

var tests = []testpair{
	{"paul@gmail.com", true},
	{"hugo.martin@latimes.com", true},
	{"someBadEmail33ff92999g99f9f9f@example.net", false},
	{"not an email", false},
}

func TestVerifyEmail(t *testing.T) {
	for _, pair := range tests {
		v, _ := VerifyEmail(pair.email)
		if v != pair.valid {
			t.Error(
				"For", pair.email,
				"expected", pair.valid,
				"got", v,
			)
		}
	}
}
