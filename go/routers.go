package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"OrdersGet",
		"GET",
		"/api/orders",
		OrdersGet,
	},

	Route{
		"OrdersPost",
		"POST",
		"/api/orders",
		OrdersPost,
	},

	Route{
		"ProductsGet",
		"GET",
		"/api/products",
		ProductsGet,
	},

	Route{
		"ProductsPost",
		"POST",
		"/api/products",
		ProductsPost,
	},

}
