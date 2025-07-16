package web

import (
	"encoding/json"
	"net/http"
)

type WebHelloHandler struct {

}

func NewWebHellopHandler() *WebHelloHandler {
	return &WebHelloHandler{};
}

func (h *WebHelloHandler) Get(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode("Hello Word")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
