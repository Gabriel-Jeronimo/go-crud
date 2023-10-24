package tests

import (
	"net/http"
	"testing"
)

type Book struct {
	title       string
	description string
	author      string
}

func TestCreate(t *testing.T) {
	newBook := Book{title: "My book", description: "Great book", author: "Myself"}

	writer := MakeRequest("POST", "/v1/books", newBook)

	if writer.Code != http.StatusCreated {
		t.Fatalf("response status code expected is 201, but instead is: %d", writer.Code)
	}
}
