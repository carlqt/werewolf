package main

import (
	"github.com/gorilla/mux"
)

func (a *App) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(ResponseHeaderHandler)

	router.HandleFunc("/games", a.GamesCreate).Methods("POST")
	router.HandleFunc("/games/join", a.GamesJoin).Methods("POST")

	return router
}
