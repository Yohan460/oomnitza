/*
External API

## Date type fields API endpoints expected date in UTC±0:00 timezone. Timezones in ISO8601 format will be ignored. API endpoints support date in two formats (one of): ISO8601 ('YYYY-MM-DDTHH:mm:SSZ') or Unix Timestamp (seconds count since January 1st, 1970 at UTC).  ## Dropdown fields Some fields are configured as dropdown fields with a dedicated list of values within Oomnitza. You can review the list of available dropdown values within the customization page in Oomnitza. In case you want to be able to post any data into these fields, you should switch them to dropdown without value within the customization page. 

API version: 3.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oomnitza

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// SoftwareApiService SoftwareApi service
type SoftwareApiService service

type ApiApiV3SoftwareGetRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	limit *string
	skip *string
	sortby *string
	filter string
}

// Limit records
func (r ApiApiV3SoftwareGetRequest) Limit(limit string) ApiApiV3SoftwareGetRequest {
	r.limit = &limit
	return r
}

// Skip records
func (r ApiApiV3SoftwareGetRequest) Skip(skip string) ApiApiV3SoftwareGetRequest {
	r.skip = &skip
	return r
}

// Order for results
func (r ApiApiV3SoftwareGetRequest) Sortby(sortby string) ApiApiV3SoftwareGetRequest {
	r.sortby = &sortby
	return r
}

func (r ApiApiV3SoftwareGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareGetExecute(r)
}

/*
ApiV3SoftwareGet Method for ApiV3SoftwareGet

Get records list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param filter Regular API v3 filter expression
 @return ApiApiV3SoftwareGetRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareGet(ctx context.Context, filter string) ApiApiV3SoftwareGetRequest {
	return ApiApiV3SoftwareGetRequest{
		ApiService: a,
		ctx: ctx,
		filter: filter,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareGetExecute(r ApiApiV3SoftwareGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software"
	localVarPath = strings.Replace(localVarPath, "{"+"filter"+"}", url.PathEscape(parameterToString(r.filter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.skip != nil {
		localVarQueryParams.Add("skip", parameterToString(*r.skip, ""))
	}
	if r.sortby != nil {
		localVarQueryParams.Add("sortby", parameterToString(*r.sortby, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareIdentAssetsGetRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	ident string
}

func (r ApiApiV3SoftwareIdentAssetsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareIdentAssetsGetExecute(r)
}

/*
ApiV3SoftwareIdentAssetsGet Method for ApiV3SoftwareIdentAssetsGet

Get history changes list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SoftwareIdentAssetsGetRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareIdentAssetsGet(ctx context.Context, ident string) ApiApiV3SoftwareIdentAssetsGetRequest {
	return ApiApiV3SoftwareIdentAssetsGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareIdentAssetsGetExecute(r ApiApiV3SoftwareIdentAssetsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareIdentAssetsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/{ident}/assets"
	localVarPath = strings.Replace(localVarPath, "{"+"ident"+"}", url.PathEscape(parameterToString(r.ident, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareIdentChangesHistoryGetRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	ident string
}

func (r ApiApiV3SoftwareIdentChangesHistoryGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareIdentChangesHistoryGetExecute(r)
}

/*
ApiV3SoftwareIdentChangesHistoryGet Method for ApiV3SoftwareIdentChangesHistoryGet

Get history changes list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SoftwareIdentChangesHistoryGetRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareIdentChangesHistoryGet(ctx context.Context, ident string) ApiApiV3SoftwareIdentChangesHistoryGetRequest {
	return ApiApiV3SoftwareIdentChangesHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareIdentChangesHistoryGetExecute(r ApiApiV3SoftwareIdentChangesHistoryGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareIdentChangesHistoryGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/{ident}/changes-history"
	localVarPath = strings.Replace(localVarPath, "{"+"ident"+"}", url.PathEscape(parameterToString(r.ident, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareIdentDeleteRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	ident string
}

func (r ApiApiV3SoftwareIdentDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareIdentDeleteExecute(r)
}

/*
ApiV3SoftwareIdentDelete Method for ApiV3SoftwareIdentDelete

Delete record

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SoftwareIdentDeleteRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareIdentDelete(ctx context.Context, ident string) ApiApiV3SoftwareIdentDeleteRequest {
	return ApiApiV3SoftwareIdentDeleteRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareIdentDeleteExecute(r ApiApiV3SoftwareIdentDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareIdentDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/{ident}"
	localVarPath = strings.Replace(localVarPath, "{"+"ident"+"}", url.PathEscape(parameterToString(r.ident, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareIdentGetRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	ident string
}

func (r ApiApiV3SoftwareIdentGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareIdentGetExecute(r)
}

/*
ApiV3SoftwareIdentGet Method for ApiV3SoftwareIdentGet

Get record details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SoftwareIdentGetRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareIdentGet(ctx context.Context, ident string) ApiApiV3SoftwareIdentGetRequest {
	return ApiApiV3SoftwareIdentGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareIdentGetExecute(r ApiApiV3SoftwareIdentGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareIdentGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/{ident}"
	localVarPath = strings.Replace(localVarPath, "{"+"ident"+"}", url.PathEscape(parameterToString(r.ident, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareIdentPatchRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	ident string
	oomnitzaIgnoreMetaRestriction *string
	software *Software
}

// Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1
func (r ApiApiV3SoftwareIdentPatchRequest) OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction string) ApiApiV3SoftwareIdentPatchRequest {
	r.oomnitzaIgnoreMetaRestriction = &oomnitzaIgnoreMetaRestriction
	return r
}

func (r ApiApiV3SoftwareIdentPatchRequest) Software(software Software) ApiApiV3SoftwareIdentPatchRequest {
	r.software = &software
	return r
}

func (r ApiApiV3SoftwareIdentPatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareIdentPatchExecute(r)
}

/*
ApiV3SoftwareIdentPatch Method for ApiV3SoftwareIdentPatch

Edit record details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SoftwareIdentPatchRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareIdentPatch(ctx context.Context, ident string) ApiApiV3SoftwareIdentPatchRequest {
	return ApiApiV3SoftwareIdentPatchRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareIdentPatchExecute(r ApiApiV3SoftwareIdentPatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareIdentPatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/{ident}"
	localVarPath = strings.Replace(localVarPath, "{"+"ident"+"}", url.PathEscape(parameterToString(r.ident, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.oomnitzaIgnoreMetaRestriction != nil {
		localVarHeaderParams["Oomnitza-Ignore-Meta-Restriction"] = parameterToString(*r.oomnitzaIgnoreMetaRestriction, "")
	}
	// body params
	localVarPostBody = r.software
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwarePostRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	oomnitzaIgnoreMetaRestriction *string
	software *Software
}

// Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1
func (r ApiApiV3SoftwarePostRequest) OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction string) ApiApiV3SoftwarePostRequest {
	r.oomnitzaIgnoreMetaRestriction = &oomnitzaIgnoreMetaRestriction
	return r
}

func (r ApiApiV3SoftwarePostRequest) Software(software Software) ApiApiV3SoftwarePostRequest {
	r.software = &software
	return r
}

func (r ApiApiV3SoftwarePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwarePostExecute(r)
}

/*
ApiV3SoftwarePost Method for ApiV3SoftwarePost

Create record

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3SoftwarePostRequest
*/
func (a *SoftwareApiService) ApiV3SoftwarePost(ctx context.Context) ApiApiV3SoftwarePostRequest {
	return ApiApiV3SoftwarePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwarePostExecute(r ApiApiV3SoftwarePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwarePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.oomnitzaIgnoreMetaRestriction != nil {
		localVarHeaderParams["Oomnitza-Ignore-Meta-Restriction"] = parameterToString(*r.oomnitzaIgnoreMetaRestriction, "")
	}
	// body params
	localVarPostBody = r.software
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareSavedsearchesGetRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
}

func (r ApiApiV3SoftwareSavedsearchesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareSavedsearchesGetExecute(r)
}

/*
ApiV3SoftwareSavedsearchesGet Method for ApiV3SoftwareSavedsearchesGet

Get saved searches list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3SoftwareSavedsearchesGetRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareSavedsearchesGet(ctx context.Context) ApiApiV3SoftwareSavedsearchesGetRequest {
	return ApiApiV3SoftwareSavedsearchesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareSavedsearchesGetExecute(r ApiApiV3SoftwareSavedsearchesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareSavedsearchesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/savedsearches"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3SoftwareSavedsearchesIdentGetRequest struct {
	ctx context.Context
	ApiService *SoftwareApiService
	ident string
}

func (r ApiApiV3SoftwareSavedsearchesIdentGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SoftwareSavedsearchesIdentGetExecute(r)
}

/*
ApiV3SoftwareSavedsearchesIdentGet Method for ApiV3SoftwareSavedsearchesIdentGet

Get saved search details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SoftwareSavedsearchesIdentGetRequest
*/
func (a *SoftwareApiService) ApiV3SoftwareSavedsearchesIdentGet(ctx context.Context, ident string) ApiApiV3SoftwareSavedsearchesIdentGetRequest {
	return ApiApiV3SoftwareSavedsearchesIdentGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SoftwareApiService) ApiV3SoftwareSavedsearchesIdentGetExecute(r ApiApiV3SoftwareSavedsearchesIdentGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareApiService.ApiV3SoftwareSavedsearchesIdentGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/software/savedsearches/{ident}"
	localVarPath = strings.Replace(localVarPath, "{"+"ident"+"}", url.PathEscape(parameterToString(r.ident, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization2"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
