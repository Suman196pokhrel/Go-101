package router

import (
	"github.com/gorilla/mux"
	"github.com/suman/25golangmongodb/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{movieId}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{movieId}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
