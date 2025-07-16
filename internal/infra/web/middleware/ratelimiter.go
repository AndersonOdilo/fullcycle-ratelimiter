package middleware

import (
	"net/http"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/usecase"
)

func RateLimiter(rateLimiterUseCase *usecase.RateLimiterUseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn :=  func(w http.ResponseWriter, r *http.Request) {

			acessoLiberado, err := rateLimiterUseCase.Execute(r);

			if (err != nil){

				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			if (acessoLiberado){
			
				next.ServeHTTP(w, r)

			} else {

				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			
			}

		}

		return http.HandlerFunc(fn)
	}
}


