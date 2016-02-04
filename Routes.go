package main

import "net/http"

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var apiVersion = "/v1"

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Register",
		"POST",
		"/register",
		Register,
	},
	Route{
		"Read",
		"GET",
		"/read/{token}",
		Read,
	},
	Route{
		"Delete",
		"DELETE",
		"/delete",
		Delete,
	},
	Route{
		"Update",
		"PUT",
		"/update",
		Update,
	},
	Route{
		"ListAll",
		"GET",
		"/list",
		ListAll,
	},
}
