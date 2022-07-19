package controllerHelper

import (
	"encoding/json"
	"net/http"

	"github.com/Anishadahal/mongoAPI/helper"
	"github.com/Anishadahal/mongoAPI/model"
	"github.com/gorilla/mux"
)

//controllers

//get all movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := helper.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

//get one movie
func GetMyOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movie := helper.GetOneMovie(params["id"])
	json.NewEncoder(w).Encode(movie)

}


//create movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	helper.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

//mark movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	helper.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

//delete one movie
func DeleteMyOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	helper.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//delete all movies
func DeleteMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deletedMovieCount := helper.DeleteAllMovies()
	json.NewEncoder(w).Encode(deletedMovieCount)

}
