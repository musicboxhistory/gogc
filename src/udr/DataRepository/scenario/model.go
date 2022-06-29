package scenario

import (
	"gogc/src/model"
)

// Cause
const (
	MandatoryIeIncorrect           = "MANDATORY_IE_INCORRECT"
	NyTypeNotAllowed               = "NF_TYPE_NOT_ALLOWED"               //403
	UnsupportedMonitoredUri        = "UNSUPPORTED_MONITORED_URI"         //501
	UserNotFoud                    = "USER_NOT_FOUND"                    //404
	DataNotFound                   = "DATA_NOT_FOUND"                    //404
	IncorrectConditionalGetRequest = "INCORRECT_CONDITIONAL_GET_REQUEST" //412
	UnprocessableRequest           = "UNPROCESSABLE_REQUEST"             //422
	DatabaseInconsistency          = "DATABASE_INCONSISTENCY"            //500
	ResourceTempMoved              = "RESOURCE_TEMP_MOVED"               //307
	ResourceMoved                  = "RESOURCE_MOVED"                    //308
	GroupIdentifierNotFound        = "GROUP_IDENTIFIER_NOT_FOUND"        //404
	ModificationNotAllowed         = "MODIFICATION_NOT_ALLOWED"          //403
	PlmnNotFound                   = "PLMN_NOT_FOUND"                    //404
)

// Detail
const (
	ErrorDetailMandatoryIeIncorrect           = "Mandatory parameters not set."
	ErrorDetailNyTypeNotAllowed               = "The target data set is not permitted to access for the NF type of the NF consumer."
	ErrorDetailUnsupportedMonitoredUri        = "The subscribe service operation is not able to be implemented due to invalid resource URI to be monitored."
	ErrorDetailUserNotFoud                    = "The user indicated in the HTTP/2 request does not exist in the UDR."
	ErrorDetailDataNotFound                   = "The data indicated in the HTTP/2 request is unavailable in the UDR."
	ErrorDetailIncorrectConditionalGetRequest = "One or more conditions given in the request header fields evaluated to false when tested in the UDR."
	ErrorDetailUnprocessableRequest           = "The request cannot be procesed due to semantic errors when trying to process a patch method."
	ErrorDetailDatabaseInconsistency          = "Requested data cannot be returned due to database inconsistency."
	ErrorDetailResourceTempMoved              = "The resource is unavailable in the current target URI but can be temporarily found in an alternative URI."
	ErrorDetailResourceMoved                  = "The resource is unavailable in the current target URI but can be permanently found in an alternative URI."
	ErrorDetailGroupIdentifierNotFound        = "The group identifier does not exist."
	ErrorDetailModificationNotAllowed         = "Modification of the target resource representation is not permitted."
	ErrorDetailPlmnNotFound                   = "The 'servingPlmnId' indicated in the HTTP/2 query is unavailable in the UDR."
)

// Ue Data
type UeDataInfo struct {
	UeIdInfo     model.UeIdentityInfo             `json:"ueidinfo,omitempty" json:"ueidinfo,omitempty"`
	AmfAccessReg *model.Amf3GppAccessRegistration `json:"amfaccessreg,omitempty" json:"amfaccessreg,omitempty"`
}
