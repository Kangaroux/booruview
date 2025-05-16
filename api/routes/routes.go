package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/posts").HandlerFunc(PostsHandler)
	r.Path("/tags").HandlerFunc(TagsHandler)
	r.Path("/tagsearch").HandlerFunc(TagSearchHandler)
	r.Use(RecoverMiddleware, RateLimitMiddleware)

	return r
}
