package server

import (
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (rw *responseWriter) Status() int {
	if rw.status == 0 {
		rw.status = 200
	}
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

type response struct {
	Data     interface{}    `json:"data,omitempty"`
	Error    ErrorInterface `json:"error,omitempty"`
	Metadata *Metadata      `json:"metadata,omitempty"`
}
