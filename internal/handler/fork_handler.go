package handler

import (
	"github.com/dragoscojocaru/forxy/internal/service"
	"net/http"
)

type ForkHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type forkHandler struct {
	forkService service.ForkService
}

func NewForkHandler() ForkHandler {
	return &forkHandler{
		forkService: service.NewForkService(),
	}
}

func (h *forkHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h.forkService.Fork(w, r)
}
