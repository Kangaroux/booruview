package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/tagsearch").HandlerFunc(TagSearchHandler)

	return r
}
