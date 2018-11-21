package main

import (
	api "./api"
	"github.com/gorilla/mux"
)

// NewRouter creates all request routing
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").HandlerFunc(api.GetHome)
	router.Methods("GET").Path("/blog").HandlerFunc(api.ListBlog)
	router.Methods("POST").Path("/blog").HandlerFunc(api.InsertBlog)
	router.Methods("PUT").Path("/blog").HandlerFunc(api.UpdateBlog)
	router.Methods("DELETE").Path("/blog").HandlerFunc(api.DeleteBlog)
	router.Methods("GET").Path("/blog/{id}").HandlerFunc(api.GetBlog)
	router.Methods("GET").Path("/blog/{length}/{index}").HandlerFunc(api.ListPagedBlog)

	return router
}
