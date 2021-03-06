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


// SaaSApiService SaaSApi service
type SaaSApiService service

type ApiApiV3SaasGetRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	limit *string
	skip *string
	sortby *string
	filter *string
}

// Limit records
func (r ApiApiV3SaasGetRequest) Limit(limit string) ApiApiV3SaasGetRequest {
	r.limit = &limit
	return r
}

// Skip records
func (r ApiApiV3SaasGetRequest) Skip(skip string) ApiApiV3SaasGetRequest {
	r.skip = &skip
	return r
}

// Order for results
func (r ApiApiV3SaasGetRequest) Sortby(sortby string) ApiApiV3SaasGetRequest {
	r.sortby = &sortby
	return r
}

// Regular API v3 filter expression
func (r ApiApiV3SaasGetRequest) Filter(filter string) ApiApiV3SaasGetRequest {
	r.filter = &filter
	return r
}

func (r ApiApiV3SaasGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasGetExecute(r)
}

/*
ApiV3SaasGet Method for ApiV3SaasGet

Get SaaS list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3SaasGetRequest
*/
func (a *SaaSApiService) ApiV3SaasGet(ctx context.Context) ApiApiV3SaasGetRequest {
	return ApiApiV3SaasGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasGetExecute(r ApiApiV3SaasGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas"

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
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
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

type ApiApiV3SaasIdentChangesHistoryGetRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	ident string
}

func (r ApiApiV3SaasIdentChangesHistoryGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasIdentChangesHistoryGetExecute(r)
}

/*
ApiV3SaasIdentChangesHistoryGet Method for ApiV3SaasIdentChangesHistoryGet

Get history changes list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SaasIdentChangesHistoryGetRequest
*/
func (a *SaaSApiService) ApiV3SaasIdentChangesHistoryGet(ctx context.Context, ident string) ApiApiV3SaasIdentChangesHistoryGetRequest {
	return ApiApiV3SaasIdentChangesHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasIdentChangesHistoryGetExecute(r ApiApiV3SaasIdentChangesHistoryGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasIdentChangesHistoryGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas/{ident}/changes-history"
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

type ApiApiV3SaasIdentDeleteRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	ident string
}

func (r ApiApiV3SaasIdentDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasIdentDeleteExecute(r)
}

/*
ApiV3SaasIdentDelete Method for ApiV3SaasIdentDelete

Delete SaaS

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SaasIdentDeleteRequest
*/
func (a *SaaSApiService) ApiV3SaasIdentDelete(ctx context.Context, ident string) ApiApiV3SaasIdentDeleteRequest {
	return ApiApiV3SaasIdentDeleteRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasIdentDeleteExecute(r ApiApiV3SaasIdentDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasIdentDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas/{ident}"
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

type ApiApiV3SaasIdentGetRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	ident string
}

func (r ApiApiV3SaasIdentGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasIdentGetExecute(r)
}

/*
ApiV3SaasIdentGet Method for ApiV3SaasIdentGet

Get SaaS details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SaasIdentGetRequest
*/
func (a *SaaSApiService) ApiV3SaasIdentGet(ctx context.Context, ident string) ApiApiV3SaasIdentGetRequest {
	return ApiApiV3SaasIdentGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasIdentGetExecute(r ApiApiV3SaasIdentGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasIdentGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas/{ident}"
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

type ApiApiV3SaasIdentPatchRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	ident string
	oomnitzaIgnoreMetaRestriction *string
	saaS *SaaS
}

// Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1
func (r ApiApiV3SaasIdentPatchRequest) OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction string) ApiApiV3SaasIdentPatchRequest {
	r.oomnitzaIgnoreMetaRestriction = &oomnitzaIgnoreMetaRestriction
	return r
}

func (r ApiApiV3SaasIdentPatchRequest) SaaS(saaS SaaS) ApiApiV3SaasIdentPatchRequest {
	r.saaS = &saaS
	return r
}

func (r ApiApiV3SaasIdentPatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasIdentPatchExecute(r)
}

/*
ApiV3SaasIdentPatch Method for ApiV3SaasIdentPatch

Edit SaaS details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SaasIdentPatchRequest
*/
func (a *SaaSApiService) ApiV3SaasIdentPatch(ctx context.Context, ident string) ApiApiV3SaasIdentPatchRequest {
	return ApiApiV3SaasIdentPatchRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasIdentPatchExecute(r ApiApiV3SaasIdentPatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasIdentPatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas/{ident}"
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
	localVarPostBody = r.saaS
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

type ApiApiV3SaasPostRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	oomnitzaIgnoreMetaRestriction *string
	saaS *SaaS
}

// Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1
func (r ApiApiV3SaasPostRequest) OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction string) ApiApiV3SaasPostRequest {
	r.oomnitzaIgnoreMetaRestriction = &oomnitzaIgnoreMetaRestriction
	return r
}

func (r ApiApiV3SaasPostRequest) SaaS(saaS SaaS) ApiApiV3SaasPostRequest {
	r.saaS = &saaS
	return r
}

func (r ApiApiV3SaasPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasPostExecute(r)
}

/*
ApiV3SaasPost Method for ApiV3SaasPost

Create SaaS

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3SaasPostRequest
*/
func (a *SaaSApiService) ApiV3SaasPost(ctx context.Context) ApiApiV3SaasPostRequest {
	return ApiApiV3SaasPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasPostExecute(r ApiApiV3SaasPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas"

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
	localVarPostBody = r.saaS
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

type ApiApiV3SaasSavedsearchesGetRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
}

func (r ApiApiV3SaasSavedsearchesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasSavedsearchesGetExecute(r)
}

/*
ApiV3SaasSavedsearchesGet Method for ApiV3SaasSavedsearchesGet

Get saved searches list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3SaasSavedsearchesGetRequest
*/
func (a *SaaSApiService) ApiV3SaasSavedsearchesGet(ctx context.Context) ApiApiV3SaasSavedsearchesGetRequest {
	return ApiApiV3SaasSavedsearchesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasSavedsearchesGetExecute(r ApiApiV3SaasSavedsearchesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasSavedsearchesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas/savedsearches"

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

type ApiApiV3SaasSavedsearchesIdentGetRequest struct {
	ctx context.Context
	ApiService *SaaSApiService
	ident string
}

func (r ApiApiV3SaasSavedsearchesIdentGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3SaasSavedsearchesIdentGetExecute(r)
}

/*
ApiV3SaasSavedsearchesIdentGet Method for ApiV3SaasSavedsearchesIdentGet

Get saved search details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ident ID of system object (assets, locations, ...)
 @return ApiApiV3SaasSavedsearchesIdentGetRequest
*/
func (a *SaaSApiService) ApiV3SaasSavedsearchesIdentGet(ctx context.Context, ident string) ApiApiV3SaasSavedsearchesIdentGetRequest {
	return ApiApiV3SaasSavedsearchesIdentGetRequest{
		ApiService: a,
		ctx: ctx,
		ident: ident,
	}
}

// Execute executes the request
func (a *SaaSApiService) ApiV3SaasSavedsearchesIdentGetExecute(r ApiApiV3SaasSavedsearchesIdentGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SaaSApiService.ApiV3SaasSavedsearchesIdentGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/saas/savedsearches/{ident}"
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
