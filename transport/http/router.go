package http

import (
	"github.com/gorilla/mux"
	"github.com/yanarowana123/onelab2_1/transport/http/handler"
)

func InitRouter(r *mux.Router, h handler.Manager) *mux.Router {
	r.HandleFunc("/transaction", h.CreateUserTransaction()).Methods("POST")
	return r
}
