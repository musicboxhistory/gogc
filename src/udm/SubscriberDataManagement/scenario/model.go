package scenario

// Cause
const (
	MandatoryIeIncorrect   = "MANDATORY_IE_INCORRECT"   //403
	Unknown5gsSubscription = "UNKNOWN_5GS_SUBSCRIPTION" //403
	UserNotFound           = "USER_NOT_FOUND"           //404
	DnnNotAllowed          = "DNN_NOT_ALLOWED"          //403
	MtcProviderNotAllowed  = "MTC_PROVIDER_NOT_ALLOWED" //403
	AfInstanceNotAllowed   = "AF_INSTANCE_NOT_ALLOWED"  //403
	SnssaiNotAllowed       = "SNSSAI_NOT_ALLOWED"       //403
	ServiceTypeNotAllowed  = "SERVICE_TYPE_NOT_ALLOWED" //403
)

// Detail
const (
	ErrorDetailMandatoryIeIncorrect   = "Mandatory parameters not set."
	ErrorDetailUnknown5gsSubscription = "No 5GS subscription is associated with the user."
	ErrorDetailUserNotFound           = "The user does not exist in the HPLMN."
	ErrorDetailDnnNotAllowed          = "DNN not authorized for the user"
	ErrorDetailMtcProviderNotAllowed  = "MTC Provider not authorized."
	ErrorDetailAfInstanceNotAllowed   = "This AF instance is not authorized."
	ErrorDetailSnssaiNotAllowed       = "This SNSSAI is not authorized to this user."
	ErrorDetailServiceTypeNotAllowed  = "This serviceType is not authorized to this user."
)
