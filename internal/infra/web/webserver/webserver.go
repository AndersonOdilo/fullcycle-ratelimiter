package webserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/database"
	internal "github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/web/middleware"
)

type WebServer struct {
	Router        		chi.Router
	Handlers      		map[string]http.HandlerFunc
	WebServerPort 		string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        		chi.NewRouter(),
		Handlers:      		make(map[string]http.HandlerFunc),
		WebServerPort: 		serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(internal.RateLimiter(database.REDIS))

	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}