package echoic

import (
	"bufio"
	"net"
	"net/http"
)

type (
	Response struct {
		writer    http.ResponseWriter
		status    int
		size      int64
		committed bool
		echoic    *Echoic
	}
)

func NewResponse(w http.ResponseWriter, e *Echoic) *Response {
	return &Response{writer: w, echoic: e}
}

func (r *Response) SetWriter(w http.ResponseWriter) {
	r.writer = w
}

func (r *Response) Header() http.Header {
	return r.writer.Header()
}

func (r *Response) Writer() http.ResponseWriter {
	return r.writer
}

func (r *Response) WriteHeader(code int) {
	if r.committed {
		return
	}
	r.status = code
	r.writer.WriteHeader(code)
	r.committed = true
}

func (r *Response) Write(b []byte) (n int, err error) {
	n, err = r.writer.Write(b)
	r.size += int64(n)
	return n, err
}

// Flush wraps response writer's Flush function.
func (r *Response) Flush() {
	r.writer.(http.Flusher).Flush()
}

// Hijack wraps response writer's Hijack function.
func (r *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return r.writer.(http.Hijacker).Hijack()
}

// CloseNotify wraps response writer's CloseNotify function.
func (r *Response) CloseNotify() <-chan bool {
	return r.writer.(http.CloseNotifier).CloseNotify()
}

func (r *Response) Status() int {
	return r.status
}

func (r *Response) Size() int64 {
	return r.size
}

func (r *Response) Committed() bool {
	return r.committed
}

func (r *Response) reset(w http.ResponseWriter, e *Echoic) {
	r.writer = w
	r.size = 0
	r.status = http.StatusOK
	r.committed = false
	r.echoic = e
}
