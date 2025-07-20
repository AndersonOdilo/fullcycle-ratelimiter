package middleware

import (
	"net/http"
	"time"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/database"
	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/usecase"
)

func RateLimiter(tipoBancoDeDados database.TipoBancoDeDados) func(next http.Handler) http.Handler {

	rateLimiterUseCase := usecase.NewRateLimiterUseCase(database.FabricaClienteRepository(tipoBancoDeDados))

	return requestRateLimiter(rateLimiterUseCase);
}

func requestRateLimiter(rateLimiterUseCase *usecase.RateLimiterUseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn :=  func(w http.ResponseWriter, r *http.Request) {

			acessoLiberado, err := rateLimiterUseCase.Execute(r.Context(), r.RemoteAddr, r.Header.Get("API_KEY"), time.Now().UnixMilli());

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


