package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
)

func Authorize(enforcer *casbin.Enforcer) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			action := r.Method
			admin := "admin_key"

			authorized, err := enforcer.Enforce(admin, path, action)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !authorized {
				http.Error(w, "Unauthorized", http.StatusForbidden)
				return
			}

			next(w, r)
		}
	}
}
