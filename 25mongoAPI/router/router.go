package router

import (
	"github.com/Anishadahal/mongoAPI/controllerHelper"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllerHelper.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controllerHelper.GetMyOneMovie).Methods("GET")
	router.HandleFunc("/api/movie", controllerHelper.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllerHelper.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllerHelper.DeleteMyOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controllerHelper.DeleteMyAllMovies).Methods("DELETE")
	return router
}
