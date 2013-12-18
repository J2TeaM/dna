package http

import (
	"dna"
	"net/http"
)

// Result is the custom type of response from an url
type Result struct {
	StatusCode dna.Int
	Status     dna.String
	Header     http.Header
	Data       dna.String
}

// NewResult initializes new result from inputs
func NewResult(statusCode dna.Int, status dna.String, header http.Header, data dna.String) *Result {
	result := new(Result)
	result.StatusCode = statusCode
	result.Status = status
	result.Header = header
	result.Data = data
	return result
}
