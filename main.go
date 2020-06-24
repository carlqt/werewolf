package main

import (
	"log"
	"os"

	"net/http"

	"github.com/carlqt/werewolf/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db := models.InitDB()
	defer db.Close()

	router := mux.NewRouter()
	router.Handle("/play", GamesCreate()).Methods("POST")
	router.Handle("/join", GamesJoin()).Methods("POST")
	router.Use(ResponseHeaderHandler)

	corsOptions := cors.New(cors.Options{
		// AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	log.Println("starting at port 8000")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8000", corsOptions.Handler(loggedRouter))
}
