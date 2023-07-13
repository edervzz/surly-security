package middlewares

import (
	"net/http"
	"strings"

	"surly-security/toolkit/localizer"
	"surly-security/toolkit/services"

	"github.com/go-http-utils/headers"
)

func UseLocalizerLanguage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
		 * localizer
		 */
		langu := r.Header.Get(headers.AcceptLanguage)
		if len(langu) == 0 {
			langu = localizer.EN
		}
		l := services.Get[localizer.ILocalizer]()
		l.SetLanguage(strings.ToLower(langu))
		next.ServeHTTP(w, r)
	})
}
