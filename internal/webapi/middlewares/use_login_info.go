package middlewares

import (
	"context"
	"net/http"
)

func UseUserInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := "email@mail.com" + ":" + r.RemoteAddr
		ctx := context.WithValue(r.Context(), "maya-user-editor", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
