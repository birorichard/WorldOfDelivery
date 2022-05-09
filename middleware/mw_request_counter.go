package middleware

import (
	"net/http"

	"github.com/birorichard/WorldOfDelivery/counter"
)

func CreateRequestCounterMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter.Increment()
		next.ServeHTTP(w, r)
	})
}
