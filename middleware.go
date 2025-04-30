package securitytxt

import "net/http"
import "go.uber.org/zap"

func Middleware() func(l *zap.Logger, name string, next http.Handler) http.Handler {
	h := Handler()
	return func(l *zap.Logger, name string, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/.well-known/security.txt" {
				h.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
