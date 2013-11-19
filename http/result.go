package http

import (
	. "dna"
	"net/http"
)

// Result is the custom type of response from an url
type Result struct {
	StatusCode Int
	Status     String
	Header     http.Header
	Data       String
}

// NewResult initializes new result from inputs
func NewResult(statusCode Int, status String, header http.Header, data String) *Result {
	result := new(Result)
	result.StatusCode = statusCode
	result.Status = status
	result.Header = header
	result.Data = data
	return result
}
