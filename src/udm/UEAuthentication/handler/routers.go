/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
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
		"/nudm-ueau/v1/",
		Index,
	},

	{
		"ConfirmAuth",
		http.MethodPost,
		"/nudm-ueau/v1/:supi/auth-events",
		ConfirmAuth,
	},

	{
		"DeleteAuth",
		http.MethodPut,
		"/nudm-ueau/v1/:supi/auth-events/:authEventId",
		DeleteAuth,
	},

	{
		"GenerateAuthData",
		http.MethodPost,
		"/nudm-ueau/v1/:supiOrSuci/security-information/generate-auth-data",
		GenerateAuthData,
	},

	{
		"GenerateGbaAv",
		http.MethodPost,
		"/nudm-ueau/v1/:supi/gba-security-information/generate-av",
		GenerateGbaAv,
	},

	{
		"GenerateAv",
		http.MethodPost,
		"/nudm-ueau/v1/:supi/hss-security-information/:hssAuthType/generate-av",
		GenerateAv,
	},

	{
		"GetRgAuthData",
		http.MethodGet,
		"/nudm-ueau/v1/:supiOrSuci/security-information-rg",
		GetRgAuthData,
	},
}