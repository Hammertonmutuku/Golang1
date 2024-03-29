package router

import (
	"github.com/gorilla/mux"
	controller "github.com/hammertonmutuku/mongodbapi/Controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
    
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
   
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	router.HandleFunc("/api/deleteall", controller.DeleteAllMovies).Methods("DELETE")
	
	return router
}
