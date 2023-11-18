package books

import "github.com/gin-gonic/gin"

type BookSerializer struct {
	C *gin.Context
	BookModel
}

type BookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func (s *BookSerializer) Response() BookResponse {
	response := BookResponse{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		Author:      s.Author,
	}

	return response
}

type BooksSerializer struct {
	C     *gin.Context
	Books []BookModel
}

func (s *BooksSerializer) Response() []BookResponse {
	response := []BookResponse{}

	for _, book := range s.Books {
		serializer := BookSerializer{s.C, book}
		response = append(response, serializer.Response())
	}

	return response
}
