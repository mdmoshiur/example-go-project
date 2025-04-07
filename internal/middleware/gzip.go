package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

// --- gzipResponseWriter wrapper ---
// This wraps the original ResponseWriter to intercept writes and compress them.
type gzipResponseWriter struct {
	io.Writer // The gzip.Writer
	http.ResponseWriter
}

// Implement Write method to write to the gzip writer
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

var gzipPool = sync.Pool{
	New: func() interface{} {
		gw, _ := gzip.NewWriterLevel(nil, gzip.DefaultCompression)
		return gw
	},
}

// Gzip middleware compress response to the gzip format.
func Gzip(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Add("Vary", "Accept-Encoding")

		gz := gzipPool.Get().(*gzip.Writer)
		defer gzipPool.Put(gz) // Put it back when done

		// Reset the writer to wrap the original ResponseWriter
		gz.Reset(w)
		defer gz.Close()

		// Create our response writer wrapper
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}

		next.ServeHTTP(gzw, r)
	}

	return http.HandlerFunc(fn)
}
