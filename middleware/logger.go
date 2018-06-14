package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	// AccessLogWriter is io.Writer for access logging.
	AccessLogWriter io.Writer = os.Stderr

	// LogRequestBody denotes if request body needs to be logged.
	LogRequestBody bool

	// BodyCompacter appends to dst the appropriately encoded src with
	// insignificant characters elided.
	BodyCompacter func(dst *bytes.Buffer, src []byte) error
)

// Some common headers.
const (
	FkForwardedForHTTPHeader = "X-Forwarded-For"
)


// LogRequest logs the request details.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		remoteIP := remoteHost(r)
		forwardedFor := forwardedForHost(r)
		logger := loggers.Get().(*log.Logger)
		logger.Printf(
			"msg=started method=%s remote_ip=%s forwarded_for=%s uri=\"%s\" user_agent=\"%s\" \n",
			r.Method, remoteIP, forwardedFor, r.RequestURI, r.UserAgent(),
		)
		if LogRequestBody {
			body := copyOfBody(r)
			logger.Printf(
				"msg=requestBody method=%s uri=\"%s\" body='%s'\n",
				r.Method, r.RequestURI, body,
			)
		}
		rw := &responseWriter{w, 200, 0}
		next.ServeHTTP(rw, r)
		logger.Printf(
			"msg=finished method=%s remote_ip=%s forwarded_for=%s uri=\"%s\" latency=%s status=%d bytes=%d user_agent=\"%s\"\n",
			r.Method, remoteIP, forwardedFor, r.RequestURI, time.Since(start).String(), rw.status, rw.bytes, r.UserAgent(),
		)
		loggers.Put(logger)
	})
}

func copyOfBody(r *http.Request) *bytes.Buffer {
	if r.Body == nil {
		return nil
	}
	b, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if len(b) > 0 && BodyCompacter != nil {
		copy := new(bytes.Buffer)
		if err := BodyCompacter(copy, b); err == nil {
			return copy
		}
	}
	return bytes.NewBuffer(b)
}

var loggers = sync.Pool{
	New: newLogger,
}

func newLogger() interface{} {
	return log.New(AccessLogWriter, "", log.LstdFlags)
}

type responseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytes += n
	return n, err
}

func remoteHost(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

func forwardedForHost(r *http.Request) string {
	remoteAddr := r.Header.Get(FkForwardedForHTTPHeader)
	host, _, _ := net.SplitHostPort(remoteAddr)
	return host
}
