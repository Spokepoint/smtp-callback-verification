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

///// Benchmarks /////

func BenchmarkVerifyEmail_Found(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerifyEmail("alexw@techcrunch.com")
	}
}

func BenchmarkVerifyEmail_Capitalised(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerifyEmail("AlexW@techcrunch.com")
	}
}

func BenchmarkVerifyEmail_NotFound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerifyEmail("someRandomEmail43432@techcrunch.com")
	}
}

func BenchmarkVerifyEmail_Blank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerifyEmail("")
	}
}
