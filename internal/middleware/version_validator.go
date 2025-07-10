package middleware

import (
	"net/http"
	"strings"
)

func VersionValidatorMiddleware(supportedVersions []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path

			valid := false
			for _, version := range supportedVersions {
				if strings.HasPrefix(path, "/"+version) {
					valid = true
					break
				}
			}

			if !valid {
				http.Error(w, "Unsupported API version", http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
