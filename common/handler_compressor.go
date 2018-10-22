package common

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

func HandlerCompressor(originalHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			originalHandler(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gZipResponseWriter{Writer: gz, ResponseWriter: w}
		originalHandler(gzr, r)
	}
}

type gZipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gZipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
