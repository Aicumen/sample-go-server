package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"PostCommentPost",
		strings.ToUpper("Post"),
		"/v1/post/comment",
		PostCommentPost,
	},

	Route{
		"PostCreatePost",
		strings.ToUpper("Post"),
		"/v1/post/create",
		PostCreatePost,
	},

	Route{
		"PostDislikePost",
		strings.ToUpper("Post"),
		"/v1/post/dislike",
		PostDislikePost,
	},

	Route{
		"PostLikePost",
		strings.ToUpper("Post"),
		"/v1/post/like",
		PostLikePost,
	},

	Route{
		"UserCreatePost",
		strings.ToUpper("Post"),
		"/v1/user/create",
		UserCreatePost,
	},

	Route{
		"UserLoginPost",
		strings.ToUpper("Post"),
		"/v1/user/login",
		UserLoginPost,
	},

	Route{
		"UserLogoutGet",
		strings.ToUpper("Get"),
		"/v1/user/logout",
		UserLogoutGet,
	},

	Route{
		"UserMypostsGet",
		strings.ToUpper("Get"),
		"/v1/user/myposts",
		UserMypostsGet,
	},

	Route{
		"LatestGet",
		strings.ToUpper("Get"),
		"/v1/latest",
		LatestGet,
	},
}
