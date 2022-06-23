/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/nudm-pp/v1/",
		Index,
	},

	{
		"Create5GVNGroup",
		http.MethodPut,
		"/nudm-pp/v1/5g-vn-groups/:extGroupId",
		Create5GVNGroup,
	},

	{
		"Delete5GVNGroup",
		http.MethodDelete,
		"/nudm-pp/v1/5g-vn-groups/:extGroupId",
		Delete5GVNGroup,
	},

	{
		"Get5GVNGroup",
		http.MethodGet,
		"/nudm-pp/v1/5g-vn-groups/:extGroupId",
		Get5GVNGroup,
	},

	{
		"Modify5GVNGroup",
		http.MethodPatch,
		"/nudm-pp/v1/5g-vn-groups/:extGroupId",
		Modify5GVNGroup,
	},

	{
		"CreatePPDataEntry",
		http.MethodPut,
		"/nudm-pp/v1/:ueId/pp-data-store/:afInstanceId",
		CreatePPDataEntry,
	},

	{
		"DeletePPDataEntry",
		http.MethodDelete,
		"/nudm-pp/v1/:ueId/pp-data-store/:afInstanceId",
		DeletePPDataEntry,
	},

	{
		"GetPPDataEntry",
		http.MethodGet,
		"/nudm-pp/v1/:ueId/pp-data-store/:afInstanceId",
		GetPPDataEntry,
	},

	{
		"Update",
		http.MethodPatch,
		"/nudm-pp/v1/:ueId/pp-data",
		Update,
	},
}
