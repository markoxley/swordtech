package swordtech

import "net/http"

// Response allows the status of the response to be exposed for logging
type Response struct {
	http.ResponseWriter
	status int
}

// NewResponse creates a new HiveResponse
func NewResponse(r http.ResponseWriter) *Response {
	w := &Response{}
	w.ResponseWriter = r
	w.status = 200
	return w
}

// WriteHeader overrites the original method to allow retrieval of the status code
func (r *Response) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// Status retrieves the status code of the response
func (r *Response) Status() int {
	return r.status
}
