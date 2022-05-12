# \SaaSApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3SaasGet**](SaaSApi.md#ApiV3SaasGet) | **Get** /api/v3/saas | 
[**ApiV3SaasIdentChangesHistoryGet**](SaaSApi.md#ApiV3SaasIdentChangesHistoryGet) | **Get** /api/v3/saas/{ident}/changes-history | 
[**ApiV3SaasIdentDelete**](SaaSApi.md#ApiV3SaasIdentDelete) | **Delete** /api/v3/saas/{ident} | 
[**ApiV3SaasIdentGet**](SaaSApi.md#ApiV3SaasIdentGet) | **Get** /api/v3/saas/{ident} | 
[**ApiV3SaasIdentPatch**](SaaSApi.md#ApiV3SaasIdentPatch) | **Patch** /api/v3/saas/{ident} | 
[**ApiV3SaasPost**](SaaSApi.md#ApiV3SaasPost) | **Post** /api/v3/saas | 
[**ApiV3SaasSavedsearchesGet**](SaaSApi.md#ApiV3SaasSavedsearchesGet) | **Get** /api/v3/saas/savedsearches | 
[**ApiV3SaasSavedsearchesIdentGet**](SaaSApi.md#ApiV3SaasSavedsearchesIdentGet) | **Get** /api/v3/saas/savedsearches/{ident} | 



## ApiV3SaasGet

> ApiV3SaasGet(ctx, filter).Limit(limit).Skip(skip).Sortby(sortby).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := "limit_example" // string | Limit records (optional)
    skip := "skip_example" // string | Skip records (optional)
    sortby := "sortby_example" // string | Order for results (optional)
    filter := "filter_example" // string | Regular API v3 filter expression (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasGet(context.Background(), filter).Limit(limit).Skip(skip).Sortby(sortby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Regular API v3 filter expression | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Limit records | 
 **skip** | **string** | Skip records | 
 **sortby** | **string** | Order for results | 


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasIdentChangesHistoryGet

> ApiV3SaasIdentChangesHistoryGet(ctx, ident).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ident := "ident_example" // string | ID of system object (assets, locations, ...)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasIdentChangesHistoryGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasIdentChangesHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentChangesHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasIdentDelete

> ApiV3SaasIdentDelete(ctx, ident).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ident := "ident_example" // string | ID of system object (assets, locations, ...)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasIdentDelete(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasIdentDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasIdentGet

> ApiV3SaasIdentGet(ctx, ident).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ident := "ident_example" // string | ID of system object (assets, locations, ...)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasIdentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasIdentPatch

> ApiV3SaasIdentPatch(ctx, ident).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).SaaS(saaS).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ident := "ident_example" // string | ID of system object (assets, locations, ...)
    oomnitzaIgnoreMetaRestriction := "oomnitzaIgnoreMetaRestriction_example" // string | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 (optional)
    saaS := *openapiclient.NewSaaS("SaasName_example", "Linked_example", "LinkKey_example", "Source_example", "Ignored_example", "Changedate_example", "Changedby_example", "Creationdate_example", "Createdby_example", false) // SaaS |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasIdentPatch(context.Background(), ident).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).SaaS(saaS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasIdentPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oomnitzaIgnoreMetaRestriction** | **string** | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 | 
 **saaS** | [**SaaS**](SaaS.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasPost

> ApiV3SaasPost(ctx).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).SaaS(saaS).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oomnitzaIgnoreMetaRestriction := "oomnitzaIgnoreMetaRestriction_example" // string | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 (optional)
    saaS := *openapiclient.NewSaaS("SaasName_example", "Linked_example", "LinkKey_example", "Source_example", "Ignored_example", "Changedate_example", "Changedby_example", "Creationdate_example", "Createdby_example", false) // SaaS |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasPost(context.Background()).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).SaaS(saaS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oomnitzaIgnoreMetaRestriction** | **string** | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 | 
 **saaS** | [**SaaS**](SaaS.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasSavedsearchesGet

> ApiV3SaasSavedsearchesGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasSavedsearchesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasSavedsearchesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasSavedsearchesGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasSavedsearchesIdentGet

> ApiV3SaasSavedsearchesIdentGet(ctx, ident).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ident := "ident_example" // string | ID of system object (assets, locations, ...)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSApi.ApiV3SaasSavedsearchesIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSApi.ApiV3SaasSavedsearchesIdentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasSavedsearchesIdentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

