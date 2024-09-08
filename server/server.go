package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SetsuikiHyoryu/intersection-observer/server/handler"
	"github.com/SetsuikiHyoryu/intersection-observer/server/middleware"
)

func main() {
	environment := initEnvironment()
	router := registerRouter()
	serve(&environment, router)
}

func initEnvironment() handler.Environment {
	environment := handler.Environment{}
	environment.Init()
	return environment
}

func registerRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/ping", handler.PingHandler)
	router.HandleFunc("/api/pokemon", handler.GetPokemons)
	return router
}

func serve(environment *handler.Environment, router *http.ServeMux) {
	middlewareHandler := middleware.CorsMiddleware(router)

	message := fmt.Sprintf("Server is listening on http://localhost/%s", environment.Port)
	fmt.Println(message)

	port := fmt.Sprintf(":%s", environment.Port)
	log.Fatal(http.ListenAndServe(port, middlewareHandler))
}
