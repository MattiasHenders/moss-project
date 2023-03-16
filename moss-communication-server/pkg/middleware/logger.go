package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	chiMiddleware "github.com/go-chi/chi/middleware"
)

func Logger(filters ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger := &chiMiddleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: false}

		fn := func(w http.ResponseWriter, r *http.Request) {
			for _, filter := range filters {
				if r.URL.Path == filter {
					next.ServeHTTP(w, r)
					return
				}
			}

			entry := logger.NewLogEntry(r)
			ww := chiMiddleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), nil)
			}()

			next.ServeHTTP(ww, chiMiddleware.WithLogEntry(r, entry))
		}
		return http.HandlerFunc(fn)
	}
}
