# \ContractsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ContractsGet**](ContractsApi.md#ApiV3ContractsGet) | **Get** /api/v3/contracts | 
[**ApiV3ContractsIdentChangesHistoryGet**](ContractsApi.md#ApiV3ContractsIdentChangesHistoryGet) | **Get** /api/v3/contracts/{ident}/changes-history | 
[**ApiV3ContractsIdentDelete**](ContractsApi.md#ApiV3ContractsIdentDelete) | **Delete** /api/v3/contracts/{ident} | 
[**ApiV3ContractsIdentGet**](ContractsApi.md#ApiV3ContractsIdentGet) | **Get** /api/v3/contracts/{ident} | 
[**ApiV3ContractsIdentPatch**](ContractsApi.md#ApiV3ContractsIdentPatch) | **Patch** /api/v3/contracts/{ident} | 
[**ApiV3ContractsPost**](ContractsApi.md#ApiV3ContractsPost) | **Post** /api/v3/contracts | 
[**ApiV3ContractsSavedsearchesGet**](ContractsApi.md#ApiV3ContractsSavedsearchesGet) | **Get** /api/v3/contracts/savedsearches | 
[**ApiV3ContractsSavedsearchesIdentGet**](ContractsApi.md#ApiV3ContractsSavedsearchesIdentGet) | **Get** /api/v3/contracts/savedsearches/{ident} | 



## ApiV3ContractsGet

> ApiV3ContractsGet(ctx, filter).Limit(limit).Skip(skip).Sortby(sortby).Execute()





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
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsGet(context.Background(), filter).Limit(limit).Skip(skip).Sortby(sortby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ContractsGetRequest struct via the builder pattern


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


## ApiV3ContractsIdentChangesHistoryGet

> ApiV3ContractsIdentChangesHistoryGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsIdentChangesHistoryGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsIdentChangesHistoryGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ContractsIdentChangesHistoryGetRequest struct via the builder pattern


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


## ApiV3ContractsIdentDelete

> ApiV3ContractsIdentDelete(ctx, ident).Execute()





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
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsIdentDelete(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsIdentDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ContractsIdentDeleteRequest struct via the builder pattern


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


## ApiV3ContractsIdentGet

> ApiV3ContractsIdentGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsIdentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ContractsIdentGetRequest struct via the builder pattern


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


## ApiV3ContractsIdentPatch

> ApiV3ContractsIdentPatch(ctx, ident).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).Contract(contract).Execute()





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
    contract := *openapiclient.NewContract() // Contract |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsIdentPatch(context.Background(), ident).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).Contract(contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsIdentPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ContractsIdentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oomnitzaIgnoreMetaRestriction** | **string** | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 | 
 **contract** | [**Contract**](Contract.md) |  | 

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


## ApiV3ContractsPost

> ApiV3ContractsPost(ctx).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).Contract(contract).Execute()





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
    contract := *openapiclient.NewContract() // Contract |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsPost(context.Background()).OomnitzaIgnoreMetaRestriction(oomnitzaIgnoreMetaRestriction).Contract(contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ContractsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oomnitzaIgnoreMetaRestriction** | **string** | Used to allow the non-system edit-only field to be updated. Allowed values: 0, 1 | 
 **contract** | [**Contract**](Contract.md) |  | 

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


## ApiV3ContractsSavedsearchesGet

> ApiV3ContractsSavedsearchesGet(ctx).Execute()





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
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsSavedsearchesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsSavedsearchesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ContractsSavedsearchesGetRequest struct via the builder pattern


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


## ApiV3ContractsSavedsearchesIdentGet

> ApiV3ContractsSavedsearchesIdentGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.ContractsApi.ApiV3ContractsSavedsearchesIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ApiV3ContractsSavedsearchesIdentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ContractsSavedsearchesIdentGetRequest struct via the builder pattern


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

