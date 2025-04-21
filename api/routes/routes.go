package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/search").HandlerFunc(SearchHandler)

	return r
}
