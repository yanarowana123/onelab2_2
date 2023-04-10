package main

import (
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/onelab2_1/configs"
	"github.com/yanarowana123/onelab2_1/internal/repositories"
	"github.com/yanarowana123/onelab2_1/internal/services"
	"github.com/yanarowana123/onelab2_1/transport/http"
	"github.com/yanarowana123/onelab2_1/transport/http/handler"
)

func init() {
	gotenv.Load()
}

func main() {
	config, err := configs.New()
	if err != nil {
		panic(err)
	}

	repositoryManager := repositories.NewManager(*config)

	serviceManager := services.NewManager(*repositoryManager)

	handlerManager := handler.NewManager(*serviceManager)

	r := mux.NewRouter()
	router := http.InitRouter(r, *handlerManager)

	http.InitServer(*config, router)
}
