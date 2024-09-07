package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SetsuikiHyoryu/intersection-observer/server/handler"
)

func main() {
	environment := initEnvironment()
	registerRouter()
	serve(&environment)
}

func initEnvironment() handler.Environment {
	environment := handler.Environment{}
	environment.Init()
	return environment
}

func registerRouter() {
	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/pokemon", handler.GetPokemons)
}

func serve(environment *handler.Environment) {
	message := fmt.Sprintf("Server is listening on http://localhost/%s", environment.Port)
	fmt.Println(message)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", environment.Port), nil))
}
