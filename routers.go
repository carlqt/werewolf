package main

import (
	"github.com/gorilla/mux"
)

func (a *App) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/games", a.GamesCreate).Methods("POST")

	return router
}
