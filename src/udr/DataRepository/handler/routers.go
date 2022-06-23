/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
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
		"/nudr-dr/v2/",
		Index,
	},

	{
		"AmfContext3gpp",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/amf-3gpp-access",
		AmfContext3gpp,
	},

	{
		"CreateAmfContext3gpp",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/amf-3gpp-access",
		CreateAmfContext3gpp,
	},

	{
		"QueryAmfContext3gpp",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/amf-3gpp-access",
		QueryAmfContext3gpp,
	},

	{
		"CreateAmfGroupSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		CreateAmfGroupSubscriptions,
	},

	{
		"AmfContextNon3gpp",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/amf-non-3gpp-access",
		AmfContextNon3gpp,
	},

	{
		"CreateAmfContextNon3gpp",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/amf-non-3gpp-access",
		CreateAmfContextNon3gpp,
	},

	{
		"QueryAmfContextNon3gpp",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/amf-non-3gpp-access",
		QueryAmfContextNon3gpp,
	},

	{
		"CreateAMFSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		CreateAMFSubscriptions,
	},

	{
		"CreateOrReplaceAccessAndMobilityData",
		http.MethodPut,
		"/nudr-dr/v2/exposure-data/:ueId/access-and-mobility-data",
		CreateOrReplaceAccessAndMobilityData,
	},

	{
		"DeleteAccessAndMobilityData",
		http.MethodDelete,
		"/nudr-dr/v2/exposure-data/:ueId/access-and-mobility-data",
		DeleteAccessAndMobilityData,
	},

	{
		"QueryAccessAndMobilityData",
		http.MethodGet,
		"/nudr-dr/v2/exposure-data/:ueId/access-and-mobility-data",
		QueryAccessAndMobilityData,
	},

	{
		"UpdateAccessAndMobilityData",
		http.MethodPatch,
		"/nudr-dr/v2/exposure-data/:ueId/access-and-mobility-data",
		UpdateAccessAndMobilityData,
	},

	{
		"ReadAccessAndMobilityPolicyData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/ues/:ueId/am-data",
		ReadAccessAndMobilityPolicyData,
	},

	{
		"QueryAmData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/am-data",
		QueryAmData,
	},

	{
		"ModifyAmfGroupSubscriptions",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		ModifyAmfGroupSubscriptions,
	},

	{
		"ModifyAmfSubscriptionInfo",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		ModifyAmfSubscriptionInfo,
	},

	{
		"CreateIndividualApplicationDataSubscription",
		http.MethodPost,
		"/nudr-dr/v2/application-data/subs-to-notify",
		CreateIndividualApplicationDataSubscription,
	},

	{
		"ReadApplicationDataChangeSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/application-data/subs-to-notify",
		ReadApplicationDataChangeSubscriptions,
	},

	{
		"DeleteAuthenticationStatus",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-status",
		DeleteAuthenticationStatus,
	},

	{
		"QueryAuthenticationStatus",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-status",
		QueryAuthenticationStatus,
	},

	{
		"QueryAuthSubsData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-subscription",
		QueryAuthSubsData,
	},

	{
		"CreateAuthenticationSoR",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/sor-data",
		CreateAuthenticationSoR,
	},

	{
		"QueryAuthSoR",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/sor-data",
		QueryAuthSoR,
	},

	{
		"UpdateAuthenticationSoR",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/sor-data",
		UpdateAuthenticationSoR,
	},

	{
		"CreateAuthenticationStatus",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-status",
		CreateAuthenticationStatus,
	},

	{
		"ModifyAuthenticationSubscription",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-subscription",
		ModifyAuthenticationSubscription,
	},

	{
		"CreateAuthenticationUPU",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/upu-data",
		CreateAuthenticationUPU,
	},

	{
		"QueryAuthUPU",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/upu-data",
		QueryAuthUPU,
	},

	{
		"ReadBdtData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/bdt-data",
		ReadBdtData,
	},

	{
		"ReadBdtPolicyData",
		http.MethodGet,
		"/nudr-dr/v2/application-data/bdtPolicyData",
		ReadBdtPolicyData,
	},

	{
		"QueryCagAck",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/subscribed-cag",
		QueryCagAck,
	},

	{
		"CreateCagUpdateAck",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/subscribed-cag",
		CreateCagUpdateAck,
	},

	{
		"Query5GVnGroupInternal",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/internal",
		Query5GVnGroupInternal,
	},

	{
		"Query5GVnGroup",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups",
		Query5GVnGroup,
	},

	{
		"Create5GVnGroup",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Create5GVnGroup,
	},

	{
		"Query5mbsData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/5mbs-data",
		Query5mbsData,
	},

	{
		"QueryContextData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data",
		QueryContextData,
	},

	{
		"Delete5GVnGroup",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Delete5GVnGroup,
	},

	{
		"QueryCoverageRestrictionData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/coverage-restriction-data",
		QueryCoverageRestrictionData,
	},

	{
		"RemoveAmfGroupSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		RemoveAmfGroupSubscriptions,
	},

	{
		"RemoveAmfSubscriptionsInfo",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		RemoveAmfSubscriptionsInfo,
	},

	{
		"QueryEEData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/ee-profile-data",
		QueryEEData,
	},

	{
		"QueryGroupEEData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-profile-data",
		QueryGroupEEData,
	},

	{
		"ModifyEeGroupSubscription",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		ModifyEeGroupSubscription,
	},

	{
		"QueryEeGroupSubscription",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		QueryEeGroupSubscription,
	},

	{
		"RemoveEeGroupSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		RemoveEeGroupSubscriptions,
	},

	{
		"UpdateEeGroupSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId",
		UpdateEeGroupSubscriptions,
	},

	{
		"CreateEeGroupSubscriptions",
		http.MethodPost,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions",
		CreateEeGroupSubscriptions,
	},

	{
		"QueryEeGroupSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions",
		QueryEeGroupSubscriptions,
	},

	{
		"ModifyEesubscription",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		ModifyEesubscription,
	},

	{
		"QueryeeSubscription",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		QueryeeSubscription,
	},

	{
		"RemoveeeSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		RemoveeeSubscriptions,
	},

	{
		"UpdateEesubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId",
		UpdateEesubscriptions,
	},

	{
		"CreateEeSubscriptions",
		http.MethodPost,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions",
		CreateEeSubscriptions,
	},

	{
		"Queryeesubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions",
		Queryeesubscriptions,
	},

	{
		"CreateIndividualExposureDataSubscription",
		http.MethodPost,
		"/nudr-dr/v2/exposure-data/subs-to-notify",
		CreateIndividualExposureDataSubscription,
	},

	{
		"GetGroupIdentifiers",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/group-identifiers",
		GetGroupIdentifiers,
	},

	{
		"CreateHssGroupSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/hss-subscriptions",
		CreateHssGroupSubscriptions,
	},

	{
		"CreateHSSSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		CreateHSSSubscriptions,
	},

	{
		"GetHssGroupSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/hss-subscriptions",
		GetHssGroupSubscriptions,
	},

	{
		"GetHssSubscriptionInfo",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		GetHssSubscriptionInfo,
	},

	{
		"ModifyHssGroupSubscriptions",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/hss-subscriptions",
		ModifyHssGroupSubscriptions,
	},

	{
		"ModifyHssSubscriptionInfo",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		ModifyHssSubscriptionInfo,
	},

	{
		"RemoveHssGroupSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/hss-subscriptions",
		RemoveHssGroupSubscriptions,
	},

	{
		"RemoveHssSubscriptionsInfo",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/hss-subscriptions",
		RemoveHssSubscriptionsInfo,
	},

	{
		"CreateHSSSDMSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		CreateHSSSDMSubscriptions,
	},

	{
		"GetHssSDMSubscriptionInfo",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		GetHssSDMSubscriptionInfo,
	},

	{
		"ModifyHssSDMSubscriptionInfo",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		ModifyHssSDMSubscriptionInfo,
	},

	{
		"RemoveHssSDMSubscriptionsInfo",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId/hss-sdm-subscriptions",
		RemoveHssSDMSubscriptionsInfo,
	},

	{
		"CreateIpSmGwContext",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ip-sm-gw",
		CreateIpSmGwContext,
	},

	{
		"DeleteIpSmGwContext",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ip-sm-gw",
		DeleteIpSmGwContext,
	},

	{
		"ModifyIpSmGwContext",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ip-sm-gw",
		ModifyIpSmGwContext,
	},

	{
		"QueryIpSmGwContext",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ip-sm-gw",
		QueryIpSmGwContext,
	},

	{
		"ReadIPTVCongifurationData",
		http.MethodGet,
		"/nudr-dr/v2/application-data/iptvConfigData",
		ReadIPTVCongifurationData,
	},

	{
		"DeleteIndividualApplicationDataSubscription",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/subs-to-notify/:subsId",
		DeleteIndividualApplicationDataSubscription,
	},

	{
		"ReadIndividualApplicationDataSubscription",
		http.MethodGet,
		"/nudr-dr/v2/application-data/subs-to-notify/:subsId",
		ReadIndividualApplicationDataSubscription,
	},

	{
		"ReplaceIndividualApplicationDataSubscription",
		http.MethodPut,
		"/nudr-dr/v2/application-data/subs-to-notify/:subsId",
		ReplaceIndividualApplicationDataSubscription,
	},

	{
		"CreateIndividualAppliedBdtPolicyData",
		http.MethodPut,
		"/nudr-dr/v2/application-data/bdtPolicyData/:bdtPolicyId",
		CreateIndividualAppliedBdtPolicyData,
	},

	{
		"DeleteIndividualAppliedBdtPolicyData",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/bdtPolicyData/:bdtPolicyId",
		DeleteIndividualAppliedBdtPolicyData,
	},

	{
		"UpdateIndividualAppliedBdtPolicyData",
		http.MethodPatch,
		"/nudr-dr/v2/application-data/bdtPolicyData/:bdtPolicyId",
		UpdateIndividualAppliedBdtPolicyData,
	},

	{
		"DeleteIndividualAuthenticationStatus",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName",
		DeleteIndividualAuthenticationStatus,
	},

	{
		"QueryIndividualAuthenticationStatus",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName",
		QueryIndividualAuthenticationStatus,
	},

	{
		"CreateIndividualAuthenticationStatus",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/authentication-data/authentication-status/:servingNetworkName",
		CreateIndividualAuthenticationStatus,
	},

	{
		"CreateIndividualBdtData",
		http.MethodPut,
		"/nudr-dr/v2/policy-data/bdt-data/:bdtReferenceId",
		CreateIndividualBdtData,
	},

	{
		"DeleteIndividualBdtData",
		http.MethodDelete,
		"/nudr-dr/v2/policy-data/bdt-data/:bdtReferenceId",
		DeleteIndividualBdtData,
	},

	{
		"ReadIndividualBdtData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/bdt-data/:bdtReferenceId",
		ReadIndividualBdtData,
	},

	{
		"UpdateIndividualBdtData",
		http.MethodPatch,
		"/nudr-dr/v2/policy-data/bdt-data/:bdtReferenceId",
		UpdateIndividualBdtData,
	},

	{
		"DeleteIndividualExposureDataSubscription",
		http.MethodDelete,
		"/nudr-dr/v2/exposure-data/subs-to-notify/:subId",
		DeleteIndividualExposureDataSubscription,
	},

	{
		"ReplaceIndividualExposureDataSubscription",
		http.MethodPut,
		"/nudr-dr/v2/exposure-data/subs-to-notify/:subId",
		ReplaceIndividualExposureDataSubscription,
	},

	{
		"CreateOrReplaceIndividualIPTVConfigurationData",
		http.MethodPut,
		"/nudr-dr/v2/application-data/iptvConfigData/:configurationId",
		CreateOrReplaceIndividualIPTVConfigurationData,
	},

	{
		"DeleteIndividualIPTVConfigurationData",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/iptvConfigData/:configurationId",
		DeleteIndividualIPTVConfigurationData,
	},

	{
		"PartialReplaceIndividualIPTVConfigurationData",
		http.MethodPatch,
		"/nudr-dr/v2/application-data/iptvConfigData/:configurationId",
		PartialReplaceIndividualIPTVConfigurationData,
	},

	{
		"CreateOrReplaceIndividualInfluenceData",
		http.MethodPut,
		"/nudr-dr/v2/application-data/influenceData/:influenceId",
		CreateOrReplaceIndividualInfluenceData,
	},

	{
		"DeleteIndividualInfluenceData",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/influenceData/:influenceId",
		DeleteIndividualInfluenceData,
	},

	{
		"UpdateIndividualInfluenceData",
		http.MethodPatch,
		"/nudr-dr/v2/application-data/influenceData/:influenceId",
		UpdateIndividualInfluenceData,
	},

	{
		"DeleteIndividualInfluenceDataSubscription",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/influenceData/subs-to-notify/:subscriptionId",
		DeleteIndividualInfluenceDataSubscription,
	},

	{
		"ReadIndividualInfluenceDataSubscription",
		http.MethodGet,
		"/nudr-dr/v2/application-data/influenceData/subs-to-notify/:subscriptionId",
		ReadIndividualInfluenceDataSubscription,
	},

	{
		"ReplaceIndividualInfluenceDataSubscription",
		http.MethodPut,
		"/nudr-dr/v2/application-data/influenceData/subs-to-notify/:subscriptionId",
		ReplaceIndividualInfluenceDataSubscription,
	},

	{
		"CreateOrReplaceIndividualPFDData",
		http.MethodPut,
		"/nudr-dr/v2/application-data/pfds/:appId",
		CreateOrReplaceIndividualPFDData,
	},

	{
		"DeleteIndividualPFDData",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/pfds/:appId",
		DeleteIndividualPFDData,
	},

	{
		"ReadIndividualPFDData",
		http.MethodGet,
		"/nudr-dr/v2/application-data/pfds/:appId",
		ReadIndividualPFDData,
	},

	{
		"DeleteIndividualPolicyDataSubscription",
		http.MethodDelete,
		"/nudr-dr/v2/policy-data/subs-to-notify/:subsId",
		DeleteIndividualPolicyDataSubscription,
	},

	{
		"ReplaceIndividualPolicyDataSubscription",
		http.MethodPut,
		"/nudr-dr/v2/policy-data/subs-to-notify/:subsId",
		ReplaceIndividualPolicyDataSubscription,
	},

	{
		"CreateOrReplaceServiceParameterData",
		http.MethodPut,
		"/nudr-dr/v2/application-data/serviceParamData/:serviceParamId",
		CreateOrReplaceServiceParameterData,
	},

	{
		"DeleteIndividualServiceParameterData",
		http.MethodDelete,
		"/nudr-dr/v2/application-data/serviceParamData/:serviceParamId",
		DeleteIndividualServiceParameterData,
	},

	{
		"UpdateIndividualServiceParameterData",
		http.MethodPatch,
		"/nudr-dr/v2/application-data/serviceParamData/:serviceParamId",
		UpdateIndividualServiceParameterData,
	},

	{
		"ReadInfluenceData",
		http.MethodGet,
		"/nudr-dr/v2/application-data/influenceData",
		ReadInfluenceData,
	},

	{
		"CreateIndividualInfluenceDataSubscription",
		http.MethodPost,
		"/nudr-dr/v2/application-data/influenceData/subs-to-notify",
		CreateIndividualInfluenceDataSubscription,
	},

	{
		"ReadInfluenceDataSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/application-data/influenceData/subs-to-notify",
		ReadInfluenceDataSubscriptions,
	},

	{
		"QueryLcsBcaData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/lcs-bca-data",
		QueryLcsBcaData,
	},

	{
		"QueryLcsMoData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/lcs-mo-data",
		QueryLcsMoData,
	},

	{
		"QueryLcsPrivacyData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/lcs-privacy-data",
		QueryLcsPrivacyData,
	},

	{
		"CreateMessageWaitingData",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/mwd",
		CreateMessageWaitingData,
	},

	{
		"DeleteMessageWaitingData",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/mwd",
		DeleteMessageWaitingData,
	},

	{
		"ModifyMessageWaitingData",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/mwd",
		ModifyMessageWaitingData,
	},

	{
		"QueryMessageWaitingData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/mwd",
		QueryMessageWaitingData,
	},

	{
		"Modify5GVnGroup",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Modify5GVnGroup,
	},

	{
		"CreateNIDDAuthorizationInfo",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/nidd-authorizations",
		CreateNIDDAuthorizationInfo,
	},

	{
		"GetNiddAuthorizationInfo",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/nidd-authorizations",
		GetNiddAuthorizationInfo,
	},

	{
		"ModifyNiddAuthorizationInfo",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/nidd-authorizations",
		ModifyNiddAuthorizationInfo,
	},

	{
		"RemoveNiddAuthorizationInfo",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/nidd-authorizations",
		RemoveNiddAuthorizationInfo,
	},

	{
		"QueryNssaiAck",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/subscribed-snssais",
		QueryNssaiAck,
	},

	{
		"CreateOrUpdateNssaiAck",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/ue-update-confirmation-data/subscribed-snssais",
		CreateOrUpdateNssaiAck,
	},

	{
		"CreateOperSpecData",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/operator-specific-data",
		CreateOperSpecData,
	},

	{
		"DeleteOperSpecData",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/operator-specific-data",
		DeleteOperSpecData,
	},

	{
		"ModifyOperSpecData",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/operator-specific-data",
		ModifyOperSpecData,
	},

	{
		"QueryOperSpecData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/operator-specific-data",
		QueryOperSpecData,
	},

	{
		"DeleteOperatorSpecificData",
		http.MethodDelete,
		"/nudr-dr/v2/policy-data/ues/:ueId/operator-specific-data",
		DeleteOperatorSpecificData,
	},

	{
		"ReadOperatorSpecificData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/ues/:ueId/operator-specific-data",
		ReadOperatorSpecificData,
	},

	{
		"ReplaceOperatorSpecificData",
		http.MethodPut,
		"/nudr-dr/v2/policy-data/ues/:ueId/operator-specific-data",
		ReplaceOperatorSpecificData,
	},

	{
		"UpdateOperatorSpecificData",
		http.MethodPatch,
		"/nudr-dr/v2/policy-data/ues/:ueId/operator-specific-data",
		UpdateOperatorSpecificData,
	},

	{
		"ReadPFDData",
		http.MethodGet,
		"/nudr-dr/v2/application-data/pfds",
		ReadPFDData,
	},

	{
		"GetppData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/pp-data",
		GetppData,
	},

	{
		"QueryPPData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/pp-profile-data",
		QueryPPData,
	},

	{
		"Query5GVNGroupPPData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/pp-profile-data",
		Query5GVNGroupPPData,
	},

	{
		"CreateOrReplaceSessionManagementData",
		http.MethodPut,
		"/nudr-dr/v2/exposure-data/:ueId/session-management-data/:pduSessionId",
		CreateOrReplaceSessionManagementData,
	},

	{
		"DeleteSessionManagementData",
		http.MethodDelete,
		"/nudr-dr/v2/exposure-data/:ueId/session-management-data/:pduSessionId",
		DeleteSessionManagementData,
	},

	{
		"QuerySessionManagementData",
		http.MethodGet,
		"/nudr-dr/v2/exposure-data/:ueId/session-management-data/:pduSessionId",
		QuerySessionManagementData,
	},

	{
		"ReadPlmnUePolicySet",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/plmns/:plmnId/ue-policy-set",
		ReadPlmnUePolicySet,
	},

	{
		"ReadPolicyData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/ues/:ueId",
		ReadPolicyData,
	},

	{
		"CreateIndividualPolicyDataSubscription",
		http.MethodPost,
		"/nudr-dr/v2/policy-data/subs-to-notify",
		CreateIndividualPolicyDataSubscription,
	},

	{
		"QueryPorseData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/prose-data",
		QueryPorseData,
	},

	{
		"QueryProvisionedData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data",
		QueryProvisionedData,
	},

	{
		"ModifyPpData",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/pp-data",
		ModifyPpData,
	},

	{
		"GetMultiplePPDataEntries",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/pp-data-store",
		GetMultiplePPDataEntries,
	},

	{
		"Get5GVnGroupConfiguration",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/:externalGroupId",
		Get5GVnGroupConfiguration,
	},

	{
		"GetAmfGroupSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/amf-subscriptions",
		GetAmfGroupSubscriptions,
	},

	{
		"GetAmfSubscriptionInfo",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/amf-subscriptions",
		GetAmfSubscriptionInfo,
	},

	{
		"GetIdentityData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/identity-data",
		GetIdentityData,
	},

	{
		"GetNiddAuData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/nidd-authorization-data",
		GetNiddAuData,
	},

	{
		"GetOdbData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/operator-determined-barring-data",
		GetOdbData,
	},

	{
		"GetSSAuData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/service-specific-authorization-data/:serviceType",
		GetSSAuData,
	},

	{
		"GetIndividualSharedData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/shared-data/:sharedDataId",
		GetIndividualSharedData,
	},

	{
		"GetSharedData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/shared-data",
		GetSharedData,
	},

	{
		"QueryRoamingInformation",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/roaming-information",
		QueryRoamingInformation,
	},

	{
		"ModifysdmSubscription",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		ModifysdmSubscription,
	},

	{
		"QuerysdmSubscription",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		QuerysdmSubscription,
	},

	{
		"RemovesdmSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		RemovesdmSubscriptions,
	},

	{
		"Updatesdmsubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId",
		Updatesdmsubscriptions,
	},

	{
		"CreateSdmSubscriptions",
		http.MethodPost,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions",
		CreateSdmSubscriptions,
	},

	{
		"Querysdmsubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/sdm-subscriptions",
		Querysdmsubscriptions,
	},

	{
		"CreateSmfGroupSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		CreateSmfGroupSubscriptions,
	},

	{
		"CreateSMFSubscriptions",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		CreateSMFSubscriptions,
	},

	{
		"GetSmfGroupSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		GetSmfGroupSubscriptions,
	},

	{
		"GetSmfSubscriptionInfo",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		GetSmfSubscriptionInfo,
	},

	{
		"ModifySmfGroupSubscriptions",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		ModifySmfGroupSubscriptions,
	},

	{
		"ModifySmfSubscriptionInfo",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		ModifySmfSubscriptionInfo,
	},

	{
		"RemoveSmfGroupSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/group-data/:ueGroupId/ee-subscriptions/:subsId/smf-subscriptions",
		RemoveSmfGroupSubscriptions,
	},

	{
		"RemoveSmfSubscriptionsInfo",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/ee-subscriptions/:subsId/smf-subscriptions",
		RemoveSmfSubscriptionsInfo,
	},

	{
		"CreateOrUpdateSmfRegistration",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		CreateOrUpdateSmfRegistration,
	},

	{
		"DeleteSmfRegistration",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		DeleteSmfRegistration,
	},

	{
		"QuerySmfRegistration",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		QuerySmfRegistration,
	},

	{
		"UpdateSmfContext",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smf-registrations/:pduSessionId",
		UpdateSmfContext,
	},

	{
		"QuerySmfRegList",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smf-registrations",
		QuerySmfRegList,
	},

	{
		"QuerySmfSelectData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/smf-selection-subscription-data",
		QuerySmfSelectData,
	},

	{
		"CreateSmsfContext3gpp",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smsf-3gpp-access",
		CreateSmsfContext3gpp,
	},

	{
		"DeleteSmsfContext3gpp",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smsf-3gpp-access",
		DeleteSmsfContext3gpp,
	},

	{
		"QuerySmsfContext3gpp",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smsf-3gpp-access",
		QuerySmsfContext3gpp,
	},

	{
		"CreateSmsfContextNon3gpp",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smsf-non-3gpp-access",
		CreateSmsfContextNon3gpp,
	},

	{
		"DeleteSmsfContextNon3gpp",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smsf-non-3gpp-access",
		DeleteSmsfContextNon3gpp,
	},

	{
		"QuerySmsfContextNon3gpp",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/smsf-non-3gpp-access",
		QuerySmsfContextNon3gpp,
	},

	{
		"QuerySmsMngData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/sms-mng-data",
		QuerySmsMngData,
	},

	{
		"QuerySmsData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/sms-data",
		QuerySmsData,
	},

	{
		"ReadServiceParameterData",
		http.MethodGet,
		"/nudr-dr/v2/application-data/serviceParamData",
		ReadServiceParameterData,
	},

	{
		"CreateServiceSpecificAuthorizationInfo",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		CreateServiceSpecificAuthorizationInfo,
	},

	{
		"GetServiceSpecificAuthorizationInfo",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		GetServiceSpecificAuthorizationInfo,
	},

	{
		"ModifyServiceSpecificAuthorizationInfo",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		ModifyServiceSpecificAuthorizationInfo,
	},

	{
		"RemoveServiceSpecificAuthorizationInfo",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/service-specific-authorizations/:serviceType",
		RemoveServiceSpecificAuthorizationInfo,
	},

	{
		"ReadSessionManagementPolicyData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/ues/:ueId/sm-data",
		ReadSessionManagementPolicyData,
	},

	{
		"UpdateSessionManagementPolicyData",
		http.MethodPatch,
		"/nudr-dr/v2/policy-data/ues/:ueId/sm-data",
		UpdateSessionManagementPolicyData,
	},

	{
		"QuerySmData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/sm-data",
		QuerySmData,
	},

	{
		"ReadSponsorConnectivityData",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/sponsor-connectivity-data/:sponsorId",
		ReadSponsorConnectivityData,
	},

	{
		"QuerySubsToNotify",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/subs-to-notify",
		QuerySubsToNotify,
	},

	{
		"RemoveMultipleSubscriptionDataSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/subs-to-notify",
		RemoveMultipleSubscriptionDataSubscriptions,
	},

	{
		"SubscriptionDataSubscriptions",
		http.MethodPost,
		"/nudr-dr/v2/subscription-data/subs-to-notify",
		SubscriptionDataSubscriptions,
	},

	{
		"ModifysubscriptionDataSubscription",
		http.MethodPatch,
		"/nudr-dr/v2/subscription-data/subs-to-notify/:subsId",
		ModifysubscriptionDataSubscription,
	},

	{
		"QuerySubscriptionDataSubscriptions",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/subs-to-notify/:subsId",
		QuerySubscriptionDataSubscriptions,
	},

	{
		"RemovesubscriptionDataSubscriptions",
		http.MethodDelete,
		"/nudr-dr/v2/subscription-data/subs-to-notify/:subsId",
		RemovesubscriptionDataSubscriptions,
	},

	{
		"CreateIndividualSubcription",
		http.MethodPost,
		"/nudr-dr/v2/data-restoration-events",
		CreateIndividualSubcription,
	},

	{
		"QueryTraceData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/:servingPlmnId/provisioned-data/trace-data",
		QueryTraceData,
	},

	{
		"CreateOrReplaceUEPolicySet",
		http.MethodPut,
		"/nudr-dr/v2/policy-data/ues/:ueId/ue-policy-set",
		CreateOrReplaceUEPolicySet,
	},

	{
		"ReadUEPolicySet",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/ues/:ueId/ue-policy-set",
		ReadUEPolicySet,
	},

	{
		"UpdateUEPolicySet",
		http.MethodPatch,
		"/nudr-dr/v2/policy-data/ues/:ueId/ue-policy-set",
		UpdateUEPolicySet,
	},

	{
		"QueryUeLocation",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/location",
		QueryUeLocation,
	},

	{
		"UpdateRoamingInformation",
		http.MethodPut,
		"/nudr-dr/v2/subscription-data/:ueId/context-data/roaming-information",
		UpdateRoamingInformation,
	},

	{
		"CreateUsageMonitoringResource",
		http.MethodPut,
		"/nudr-dr/v2/policy-data/ues/:ueId/sm-data/:usageMonId",
		CreateUsageMonitoringResource,
	},

	{
		"DeleteUsageMonitoringInformation",
		http.MethodDelete,
		"/nudr-dr/v2/policy-data/ues/:ueId/sm-data/:usageMonId",
		DeleteUsageMonitoringInformation,
	},

	{
		"ReadUsageMonitoringInformation",
		http.MethodGet,
		"/nudr-dr/v2/policy-data/ues/:ueId/sm-data/:usageMonId",
		ReadUsageMonitoringInformation,
	},

	{
		"QueryUserConsentData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/uc-data",
		QueryUserConsentData,
	},

	{
		"QueryV2xData",
		http.MethodGet,
		"/nudr-dr/v2/subscription-data/:ueId/v2x-data",
		QueryV2xData,
	},
}
