package router

import (
	"github.com/bikashsaud/mongodb_api/controller"
	"github.com/gorilla/mux"
)

// import "controller"

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
