package scenario

// Cause
const (
	MandatoryIeIncorrect = "MANDATORY_IE_INCORRECT"
	InvalidRequest       = "INVALID_REQUEST"
	InvalidClient        = "INVALID_CLIENT"
	UnauthorizedClient   = "UNAUTHORIZED_CLIENT"
	UnsupportedGrantType = "UNSUPPORTED_GRANT_TYPE"
	InvalidScope         = "INVALID_SCOPE"
)

// Detail
const (
	ErrorDetailMandatoryIeIncorrect = "mandatory parameters not set"
	ErrorDetailInvalidRequest       = "invalid request"
	ErrorDetailInvalidClient        = "invalid client"
	ErrorDetailUnauthorizedClient   = "unauthorized client"
	ErrorDetailUnsupportedGrantType = "unsupported grant type"
	ErrorDetailInvalidScope         = "invalid scope"
)
