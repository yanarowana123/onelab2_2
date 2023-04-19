package http

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/yanarowana123/onelab2_1/transport/http/handler"
	"net/http"
)

func InitRouter(r *mux.Router, h handler.Manager) *mux.Router {
	//TODO add auth to swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8082/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.HandleFunc("/transaction", h.SecretKeyValidateMiddleware(h.CreateTransaction())).Methods("POST")
	r.HandleFunc("/sum/book/{bookID}", h.SecretKeyValidateMiddleware(h.GetMoneySumByBookID())).Methods("GET")
	return r
}
