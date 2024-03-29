/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CompleteStoredSearchDocumentApiService CompleteStoredSearchDocumentApi service
type CompleteStoredSearchDocumentApiService service

type ApiRetrieveCompleteSearchRequest struct {
	ctx context.Context
	ApiService *CompleteStoredSearchDocumentApiService
	searchId string
	acceptEncoding *string
}

// Accept-Encoding, described in IETF RFC 7231
func (r ApiRetrieveCompleteSearchRequest) AcceptEncoding(acceptEncoding string) ApiRetrieveCompleteSearchRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

func (r ApiRetrieveCompleteSearchRequest) Execute() (*StoredSearchResult, *http.Response, error) {
	return r.ApiService.RetrieveCompleteSearchExecute(r)
}

/*
RetrieveCompleteSearch Method for RetrieveCompleteSearch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param searchId Id of a stored search
 @return ApiRetrieveCompleteSearchRequest
*/
func (a *CompleteStoredSearchDocumentApiService) RetrieveCompleteSearch(ctx context.Context, searchId string) ApiRetrieveCompleteSearchRequest {
	return ApiRetrieveCompleteSearchRequest{
		ApiService: a,
		ctx: ctx,
		searchId: searchId,
	}
}

// Execute executes the request
//  @return StoredSearchResult
func (a *CompleteStoredSearchDocumentApiService) RetrieveCompleteSearchExecute(r ApiRetrieveCompleteSearchRequest) (*StoredSearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StoredSearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompleteStoredSearchDocumentApiService.RetrieveCompleteSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/searches/{searchId}/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"searchId"+"}", url.PathEscape(parameterToString(r.searchId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = parameterToString(*r.acceptEncoding, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 307 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 308 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
