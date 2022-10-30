package scenario

// Cause
const (
	MandatoryIeIncorrect         = "MANDATORY_IE_INCORRECT"
	SnssaiNotSupported           = "SNSSAI_NOT_SUPPORTED"
	NotAuthorized                = "NOT_AUTHORIZED"
	ResourceNotFound             = "RESOURCE_NOT_FOUND"
	SubscriptionNotFound         = "SUBSCRIPTION_NOT_FOUND"
	ResourceUriStructureNotFound = "RESOURCE_URI_STRUCTURE_NOT_FOUND"
)

// Detail
const (
	ErrorDetailMandatoryIeIncorrect         = "Mandatory parameters not set"
	ErrorDetailSnssaiNotSupported           = "This cause value shall be set when the requested slice selection information is for SNSSAI(s) not supported."
	ErrorDetailNotAuthorized                = "The request is rejected due to the NF service consumer is not authorized to retrieve the slice selection information."
	ErrorDetailNotAuthorized                = "The request is rejected due to the NF service consumer is not authorized to update the NSSAI availability information, or subscribe for the NSSAI availability information notification."
	ErrorDetailResourceNotFound             = "The request is rejected due to the NF service consumer is authorized, but the resource related to the NF Id for which the NSSAI availability information is updated or deleted is unavailable."
	ErrorDetailSubscriptionNotFound         = "Indicates the modification of subscription has failed due to an application error when the subscription is not found in the NSSF."
	ErrorDetailResourceUriStructureNotFound = "ndicates that the NF Service Consumer (e.g. AMF) received a notification request from NSSF on a callback URI that is not known to the NF Service Consumer."
)
