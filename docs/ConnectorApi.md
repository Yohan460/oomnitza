# \ConnectorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConnectorRunLogsConnectorsIdentGet**](ConnectorApi.md#ApiV3ConnectorRunLogsConnectorsIdentGet) | **Get** /api/v3/connector_run_logs/connectors/{ident} | 
[**ApiV3ConnectorRunLogsIdentGet**](ConnectorApi.md#ApiV3ConnectorRunLogsIdentGet) | **Get** /api/v3/connector_run_logs/{ident} | 



## ApiV3ConnectorRunLogsConnectorsIdentGet

> ApiV3ConnectorRunLogsConnectorsIdentGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.ConnectorApi.ApiV3ConnectorRunLogsConnectorsIdentGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorApi.ApiV3ConnectorRunLogsConnectorsIdentGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ConnectorRunLogsConnectorsIdentGetRequest struct via the builder pattern


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


## ApiV3ConnectorRunLogsIdentGet

> ApiV3ConnectorRunLogsIdentGet(ctx, filter, ident).Limit(limit).Skip(skip).Sortby(sortby).Execute()





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
    limit := "limit_example" // string | Limit records (optional)
    skip := "skip_example" // string | Skip records (optional)
    sortby := "sortby_example" // string | Order for results (optional)
    filter := "filter_example" // string | Regular API v3 filter expression (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorApi.ApiV3ConnectorRunLogsIdentGet(context.Background(), filter, ident).Limit(limit).Skip(skip).Sortby(sortby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorApi.ApiV3ConnectorRunLogsIdentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Regular API v3 filter expression | 
**ident** | **string** | ID of system object (assets, locations, ...) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConnectorRunLogsIdentGetRequest struct via the builder pattern


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

