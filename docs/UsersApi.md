# \UsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3UsersGet**](UsersApi.md#ApiV3UsersGet) | **Get** /api/v3/users | 
[**ApiV3UsersIdentAssetsGet**](UsersApi.md#ApiV3UsersIdentAssetsGet) | **Get** /api/v3/users/{ident}/assets | 
[**ApiV3UsersIdentChangesHistoryGet**](UsersApi.md#ApiV3UsersIdentChangesHistoryGet) | **Get** /api/v3/users/{ident}/changes-history | 
[**ApiV3UsersIdentDelete**](UsersApi.md#ApiV3UsersIdentDelete) | **Delete** /api/v3/users/{ident} | 
[**ApiV3UsersIdentGet**](UsersApi.md#ApiV3UsersIdentGet) | **Get** /api/v3/users/{ident} | 
[**ApiV3UsersIdentPatch**](UsersApi.md#ApiV3UsersIdentPatch) | **Patch** /api/v3/users/{ident} | 
[**ApiV3UsersIdentProfileSettingsGet**](UsersApi.md#ApiV3UsersIdentProfileSettingsGet) | **Get** /api/v3/users/{ident}/profile_settings | 
[**ApiV3UsersIdentSoftwareGet**](UsersApi.md#ApiV3UsersIdentSoftwareGet) | **Get** /api/v3/users/{ident}/software | 
[**ApiV3UsersIdentSoftwareSaasGet**](UsersApi.md#ApiV3UsersIdentSoftwareSaasGet) | **Get** /api/v3/users/{ident}/software/saas | 
[**ApiV3UsersPost**](UsersApi.md#ApiV3UsersPost) | **Post** /api/v3/users | 
[**ApiV3UsersSavedsearchesGet**](UsersApi.md#ApiV3UsersSavedsearchesGet) | **Get** /api/v3/users/savedsearches | 
[**ApiV3UsersSavedsearchesIdentGet**](UsersApi.md#ApiV3UsersSavedsearchesIdentGet) | **Get** /api/v3/users/savedsearches/{ident} | 



## ApiV3UsersGet

> ApiV3UsersGet(ctx, filter).Limit(limit).Skip(skip).Sortby(sortby).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersGet(context.Background(), filter).Limit(limit).Skip(skip).Sortby(sortby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersGetRequest struct via the builder pattern


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


## ApiV3UsersIdentAssetsGet

> ApiV3UsersIdentAssetsGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentAssetsGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentAssetsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentAssetsGetRequest struct via the builder pattern


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


## ApiV3UsersIdentChangesHistoryGet

> ApiV3UsersIdentChangesHistoryGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentChangesHistoryGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentChangesHistoryGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentChangesHistoryGetRequest struct via the builder pattern


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


## ApiV3UsersIdentDelete

> ApiV3UsersIdentDelete(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentDelete(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentDeleteRequest struct via the builder pattern


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


## ApiV3UsersIdentGet

> ApiV3UsersIdentGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentGetRequest struct via the builder pattern


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


## ApiV3UsersIdentPatch

> ApiV3UsersIdentPatch(ctx, ident).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).User(user).Execute()





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
    user := *openapiclient.NewUser() // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentPatch(context.Background(), ident).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oomnitzaIgnoreMetaRestriction** | **string** | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 | 
 **user** | [**User**](User.md) |  | 

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


## ApiV3UsersIdentProfileSettingsGet

> ApiV3UsersIdentProfileSettingsGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentProfileSettingsGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentProfileSettingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentProfileSettingsGetRequest struct via the builder pattern


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


## ApiV3UsersIdentSoftwareGet

> ApiV3UsersIdentSoftwareGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentSoftwareGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentSoftwareGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentSoftwareGetRequest struct via the builder pattern


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


## ApiV3UsersIdentSoftwareSaasGet

> ApiV3UsersIdentSoftwareSaasGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersIdentSoftwareSaasGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersIdentSoftwareSaasGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersIdentSoftwareSaasGetRequest struct via the builder pattern


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


## ApiV3UsersPost

> ApiV3UsersPost(ctx).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).User(user).Execute()





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
    user := *openapiclient.NewUser() // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiV3UsersPost(context.Background()).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3UsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oomnitzaIgnoreMetaRestriction** | **string** | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 | 
 **user** | [**User**](User.md) |  | 

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


## ApiV3UsersSavedsearchesGet

> ApiV3UsersSavedsearchesGet(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersSavedsearchesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersSavedsearchesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3UsersSavedsearchesGetRequest struct via the builder pattern


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


## ApiV3UsersSavedsearchesIdentGet

> ApiV3UsersSavedsearchesIdentGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.UsersApi.ApiV3UsersSavedsearchesIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiV3UsersSavedsearchesIdentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3UsersSavedsearchesIdentGetRequest struct via the builder pattern


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

