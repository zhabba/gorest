package main

import "net/http"

type Route struct {
	Name       string
	Method     string
	Pattern    string
	ApiVersion string
	Handler    http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		"/v1",
		Index,
	},
	Route{
		"Register",
		"POST",
		"/register",
		"/v1",
		Register,
	},
	Route{
		"Read",
		"GET",
		"/read/{token}",
		"/v1",
		Read,
	},
	Route{
		"Delete",
		"DELETE",
		"/delete",
		"/v1",
		Delete,
	},
	Route{
		"Update",
		"PUT",
		"/update",
		"/v1",
		Update,
	},
}
