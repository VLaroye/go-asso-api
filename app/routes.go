package app

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var apiRoutes = routes{
	route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: homeHandler,
	},
	route{
		Name:        "GetMembers",
		Method:      "GET",
		Pattern:     "/members",
		HandlerFunc: getMembersHandler,
	},
	route{
		Name:        "CreateMember",
		Method:      "POST",
		Pattern:     "/members",
		HandlerFunc: createMemberHandler,
	},
}
