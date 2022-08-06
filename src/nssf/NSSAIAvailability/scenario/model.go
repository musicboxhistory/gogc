package scenario

// Cause
const (
	MandatoryIeIncorrect = "MANDATORY_IE_INCORRECT"
	SnssaiNotSupported   = "SNSSAI_NOT_SUPPORTED"
	NotAuthorized        = "NOT_AUTHORIZED"
)

// Detail
const (
	ErrorDetailMandatoryIeIncorrect = "Mandatory parameters not set"
	ErrorDetailSnssaiNotSupported   = "This cause value shall be set when the requested slice selection information is for SNSSAI(s) not supported."
	ErrorDetailNotAuthorized        = "The request is rejected due to the NF service consumer is not authorized to retrieve the slice selection information."
)
