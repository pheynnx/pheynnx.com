package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pheynnx/pheynnx.com/internal/database"
)

func Metrics(DB *database.DBPool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			go func() {
				time.Sleep(5 * time.Second)
				fmt.Println(r.URL)
			}()

			next.ServeHTTP(w, r)
		})
	}
}
