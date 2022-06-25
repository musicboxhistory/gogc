/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
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
		"/nudm-sdm/v1/",
		Index,
	},

	{
		"GetAmData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/am-data",
		GetAmData,
	},

	{
		"GetMbsData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/5mbs-data",
		GetMbsData,
	},

	{
		"GetEcrData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/am-data/ecr-data",
		GetEcrData,
	},

	{
		"GetSupiOrGpsi",
		http.MethodGet,
		"/nudm-sdm/v1/:ueId/id-translation-result",
		GetSupiOrGpsi,
	},

	{
		"GetGroupIdentifiers",
		http.MethodGet,
		"/nudm-sdm/v1/group-data/group-identifiers",
		GetGroupIdentifiers,
	},

	{
		"GetLcsBcaData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/lcs-bca-data",
		GetLcsBcaData,
	},

	{
		"GetLcsMoData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/lcs-mo-data",
		GetLcsMoData,
	},

	{
		"GetLcsPrivacyData",
		http.MethodGet,
		"/nudm-sdm/v1/:ueId/lcs-privacy-data",
		GetLcsPrivacyData,
	},

	{
		"GetProseData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/prose-data",
		GetProseData,
	},

	{
		"CAGAck",
		http.MethodPut,
		"/nudm-sdm/v1/:supi/am-data/cag-ack",
		CAGAck,
	},

	{
		"SNSSAIsAck",
		http.MethodPut,
		"/nudm-sdm/v1/:supi/am-data/subscribed-snssais-ack",
		SNSSAIsAck,
	},

	{
		"SorAckInfo",
		http.MethodPut,
		"/nudm-sdm/v1/:supi/am-data/sor-ack",
		SorAckInfo,
	},

	{
		"UpuAck",
		http.MethodPut,
		"/nudm-sdm/v1/:supi/am-data/upu-ack",
		UpuAck,
	},

	{
		"GetDataSets",
		http.MethodGet,
		"/nudm-sdm/v1/:supi",
		GetDataSets,
	},

	{
		"GetSharedData",
		http.MethodGet,
		"/nudm-sdm/v1/shared-data",
		GetSharedData,
	},

	{
		"GetIndividualSharedData",
		http.MethodGet,
		"/nudm-sdm/v1/shared-data/:sharedDataId",
		GetIndividualSharedData,
	},

	{
		"GetSmfSelData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/smf-select-data",
		GetSmfSelData,
	},

	{
		"GetSmsMngtData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/sms-mng-data",
		GetSmsMngtData,
	},

	{
		"GetSmsData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/sms-data",
		GetSmsData,
	},

	{
		"GetSmData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/sm-data",
		GetSmData,
	},

	{
		"GetNSSAI",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/nssai",
		GetNSSAI,
	},

	{
		"Subscribe",
		http.MethodPost,
		"/nudm-sdm/v1/:ueId/sdm-subscriptions",
		Subscribe,
	},

	{
		"SubscribeToSharedData",
		http.MethodPost,
		"/nudm-sdm/v1/shared-data-subscriptions",
		SubscribeToSharedData,
	},

	{
		"Unsubscribe",
		http.MethodDelete,
		"/nudm-sdm/v1/:ueId/sdm-subscriptions/:subscriptionId",
		Unsubscribe,
	},

	{
		"UnsubscribeForSharedData",
		http.MethodDelete,
		"/nudm-sdm/v1/shared-data-subscriptions/:subscriptionId",
		UnsubscribeForSharedData,
	},

	{
		"Modify",
		http.MethodPatch,
		"/nudm-sdm/v1/:ueId/sdm-subscriptions/:subscriptionId",
		Modify,
	},

	{
		"ModifySharedDataSubs",
		http.MethodPatch,
		"/nudm-sdm/v1/shared-data-subscriptions/:subscriptionId",
		ModifySharedDataSubs,
	},

	{
		"GetTraceConfigData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/trace-data",
		GetTraceConfigData,
	},

	{
		"UpdateSORInfo",
		http.MethodPost,
		"/nudm-sdm/v1/:supi/am-data/update-sor",
		UpdateSORInfo,
	},

	{
		"GetUeCtxInAmfData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/ue-context-in-amf-data",
		GetUeCtxInAmfData,
	},

	{
		"GetUeCtxInSmfData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/ue-context-in-smf-data",
		GetUeCtxInSmfData,
	},

	{
		"GetUeCtxInSmsfData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/ue-context-in-smsf-data",
		GetUeCtxInSmsfData,
	},

	{
		"GetUcData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/uc-data",
		GetUcData,
	},

	{
		"GetV2xData",
		http.MethodGet,
		"/nudm-sdm/v1/:supi/v2x-data",
		GetV2xData,
	},
}
