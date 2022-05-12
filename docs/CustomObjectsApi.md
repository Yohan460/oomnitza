# \CustomObjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3CustomObjectsGet**](CustomObjectsApi.md#ApiV3CustomObjectsGet) | **Get** /api/v3/custom_objects | 
[**ApiV3CustomObjectsIdentDelete**](CustomObjectsApi.md#ApiV3CustomObjectsIdentDelete) | **Delete** /api/v3/custom_objects/{ident} | 
[**ApiV3CustomObjectsIdentGet**](CustomObjectsApi.md#ApiV3CustomObjectsIdentGet) | **Get** /api/v3/custom_objects/{ident} | 
[**ApiV3CustomObjectsIdentPatch**](CustomObjectsApi.md#ApiV3CustomObjectsIdentPatch) | **Patch** /api/v3/custom_objects/{ident} | 
[**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet**](CustomObjectsApi.md#ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet) | **Get** /api/v3/custom_objects/{ident}/records/{custom_object_record_id}/changes-history | 
[**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete**](CustomObjectsApi.md#ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete) | **Delete** /api/v3/custom_objects/{ident}/records/{custom_object_record_id} | 
[**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet**](CustomObjectsApi.md#ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet) | **Get** /api/v3/custom_objects/{ident}/records/{custom_object_record_id} | 
[**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch**](CustomObjectsApi.md#ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch) | **Patch** /api/v3/custom_objects/{ident}/records/{custom_object_record_id} | 
[**ApiV3CustomObjectsIdentRecordsGet**](CustomObjectsApi.md#ApiV3CustomObjectsIdentRecordsGet) | **Get** /api/v3/custom_objects/{ident}/records | 
[**ApiV3CustomObjectsIdentRecordsPost**](CustomObjectsApi.md#ApiV3CustomObjectsIdentRecordsPost) | **Post** /api/v3/custom_objects/{ident}/records | 
[**ApiV3CustomObjectsPost**](CustomObjectsApi.md#ApiV3CustomObjectsPost) | **Post** /api/v3/custom_objects | 



## ApiV3CustomObjectsGet

> ApiV3CustomObjectsGet(ctx).Execute()





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
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomObjectsGetRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentDelete

> ApiV3CustomObjectsIdentDelete(ctx, ident).Execute()





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
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentDelete(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentDeleteRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentGet

> ApiV3CustomObjectsIdentGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentGetRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentPatch

> ApiV3CustomObjectsIdentPatch(ctx, ident).CustomObject(customObject).Execute()





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
    customObject := *openapiclient.NewCustomObject("CustomObjectName_example", "Type_example") // CustomObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentPatch(context.Background(), ident).CustomObject(customObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customObject** | [**CustomObject**](CustomObject.md) |  | 

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


## ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet

> ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet(ctx, ident, customObjectRecordId).Execute()





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
    customObjectRecordId := "customObjectRecordId_example" // string | ID of Custom Object Record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet(context.Background(), ident, customObjectRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**customObjectRecordId** | **string** | ID of Custom Object Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGetRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete

> ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete(ctx, ident, customObjectRecordId).Execute()





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
    customObjectRecordId := "customObjectRecordId_example" // string | ID of Custom Object Record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete(context.Background(), ident, customObjectRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**customObjectRecordId** | **string** | ID of Custom Object Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDeleteRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet

> ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet(ctx, ident, customObjectRecordId).Execute()





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
    customObjectRecordId := "customObjectRecordId_example" // string | ID of Custom Object Record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet(context.Background(), ident, customObjectRecordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**customObjectRecordId** | **string** | ID of Custom Object Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGetRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch

> ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch(ctx, ident, customObjectRecordId).CustomObjectRecord(customObjectRecord).Execute()





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
    customObjectRecordId := "customObjectRecordId_example" // string | ID of Custom Object Record
    customObjectRecord := *openapiclient.NewCustomObjectRecord("CustomObjectRecordName_example") // CustomObjectRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch(context.Background(), ident, customObjectRecordId).CustomObjectRecord(customObjectRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**customObjectRecordId** | **string** | ID of Custom Object Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customObjectRecord** | [**CustomObjectRecord**](CustomObjectRecord.md) |  | 

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


## ApiV3CustomObjectsIdentRecordsGet

> ApiV3CustomObjectsIdentRecordsGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentRecordsGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentRecordsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentRecordsGetRequest struct via the builder pattern


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


## ApiV3CustomObjectsIdentRecordsPost

> ApiV3CustomObjectsIdentRecordsPost(ctx, ident).CustomObjectRecord(customObjectRecord).Execute()





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
    customObjectRecord := *openapiclient.NewCustomObjectRecord("CustomObjectRecordName_example") // CustomObjectRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsIdentRecordsPost(context.Background(), ident).CustomObjectRecord(customObjectRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsIdentRecordsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomObjectsIdentRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customObjectRecord** | [**CustomObjectRecord**](CustomObjectRecord.md) |  | 

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


## ApiV3CustomObjectsPost

> ApiV3CustomObjectsPost(ctx).CustomObject(customObject).Execute()





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
    customObject := *openapiclient.NewCustomObject("CustomObjectName_example", "Type_example") // CustomObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomObjectsApi.ApiV3CustomObjectsPost(context.Background()).CustomObject(customObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomObjectsApi.ApiV3CustomObjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomObjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customObject** | [**CustomObject**](CustomObject.md) |  | 

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

