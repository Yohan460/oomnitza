# \SaaSUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3SaasIdentRolesGet**](SaaSUsersApi.md#ApiV3SaasIdentRolesGet) | **Get** /api/v3/saas/{ident}/roles | 
[**ApiV3SaasIdentRolesPost**](SaaSUsersApi.md#ApiV3SaasIdentRolesPost) | **Post** /api/v3/saas/{ident}/roles | 
[**ApiV3SaasIdentRolesSaasRoleIdDelete**](SaaSUsersApi.md#ApiV3SaasIdentRolesSaasRoleIdDelete) | **Delete** /api/v3/saas/{ident}/roles/{saas_role_id} | 
[**ApiV3SaasIdentRolesSaasRoleIdPatch**](SaaSUsersApi.md#ApiV3SaasIdentRolesSaasRoleIdPatch) | **Patch** /api/v3/saas/{ident}/roles/{saas_role_id} | 
[**ApiV3SaasIdentRolesUsersDelete**](SaaSUsersApi.md#ApiV3SaasIdentRolesUsersDelete) | **Delete** /api/v3/saas/{ident}/roles/users | 
[**ApiV3SaasIdentRolesUsersPatch**](SaaSUsersApi.md#ApiV3SaasIdentRolesUsersPatch) | **Patch** /api/v3/saas/{ident}/roles/users | 
[**ApiV3SaasIdentRolesUsersPost**](SaaSUsersApi.md#ApiV3SaasIdentRolesUsersPost) | **Post** /api/v3/saas/{ident}/roles/users | 
[**ApiV3SaasIdentUsersGet**](SaaSUsersApi.md#ApiV3SaasIdentUsersGet) | **Get** /api/v3/saas/{ident}/users | 
[**ApiV3SaasIdentUsersPost**](SaaSUsersApi.md#ApiV3SaasIdentUsersPost) | **Post** /api/v3/saas/{ident}/users | 
[**ApiV3SaasIdentUsersUserIdPost**](SaaSUsersApi.md#ApiV3SaasIdentUsersUserIdPost) | **Post** /api/v3/saas/{ident}/users/{user_id} | 
[**ApiV3SaasUsersActivatePatch**](SaaSUsersApi.md#ApiV3SaasUsersActivatePatch) | **Patch** /api/v3/saas_users/activate | 
[**ApiV3SaasUsersDeactivatePatch**](SaaSUsersApi.md#ApiV3SaasUsersDeactivatePatch) | **Patch** /api/v3/saas_users/deactivate | 
[**ApiV3SaasUsersGet**](SaaSUsersApi.md#ApiV3SaasUsersGet) | **Get** /api/v3/saas_users | 



## ApiV3SaasIdentRolesGet

> ApiV3SaasIdentRolesGet(ctx, ident).Execute()





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
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesGet(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesGetRequest struct via the builder pattern


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


## ApiV3SaasIdentRolesPost

> ApiV3SaasIdentRolesPost(ctx, ident).SaaSCreateRoleObject(saaSCreateRoleObject).Execute()





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
    saaSCreateRoleObject := *openapiclient.NewSaaSCreateRoleObject() // SaaSCreateRoleObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesPost(context.Background(), ident).SaaSCreateRoleObject(saaSCreateRoleObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saaSCreateRoleObject** | [**SaaSCreateRoleObject**](SaaSCreateRoleObject.md) |  | 

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


## ApiV3SaasIdentRolesSaasRoleIdDelete

> ApiV3SaasIdentRolesSaasRoleIdDelete(ctx, ident, saasRoleId).Execute()





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
    saasRoleId := "saasRoleId_example" // string | ID of SaaS Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesSaasRoleIdDelete(context.Background(), ident, saasRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesSaasRoleIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**saasRoleId** | **string** | ID of SaaS Role | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesSaasRoleIdDeleteRequest struct via the builder pattern


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


## ApiV3SaasIdentRolesSaasRoleIdPatch

> ApiV3SaasIdentRolesSaasRoleIdPatch(ctx, ident, saasRoleId).SaaSEditRoleObject(saaSEditRoleObject).Execute()





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
    saasRoleId := "saasRoleId_example" // string | ID of SaaS Role
    saaSEditRoleObject := *openapiclient.NewSaaSEditRoleObject() // SaaSEditRoleObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesSaasRoleIdPatch(context.Background(), ident, saasRoleId).SaaSEditRoleObject(saaSEditRoleObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesSaasRoleIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**saasRoleId** | **string** | ID of SaaS Role | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesSaasRoleIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **saaSEditRoleObject** | [**SaaSEditRoleObject**](SaaSEditRoleObject.md) |  | 

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


## ApiV3SaasIdentRolesUsersDelete

> ApiV3SaasIdentRolesUsersDelete(ctx, ident).Execute()





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
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesUsersDelete(context.Background(), ident).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesUsersDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesUsersDeleteRequest struct via the builder pattern


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


## ApiV3SaasIdentRolesUsersPatch

> ApiV3SaasIdentRolesUsersPatch(ctx, ident).UsersSaaSRoleUpdateListObject(usersSaaSRoleUpdateListObject).Execute()





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
    usersSaaSRoleUpdateListObject := *openapiclient.NewUsersSaaSRoleUpdateListObject("Users_example", []openapiclient.UsersSaaSRoleUpdateListObjectRolesInner{*openapiclient.NewUsersSaaSRoleUpdateListObjectRolesInner("RoleExternalId_example")}) // UsersSaaSRoleUpdateListObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesUsersPatch(context.Background(), ident).UsersSaaSRoleUpdateListObject(usersSaaSRoleUpdateListObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesUsersPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesUsersPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usersSaaSRoleUpdateListObject** | [**UsersSaaSRoleUpdateListObject**](UsersSaaSRoleUpdateListObject.md) |  | 

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


## ApiV3SaasIdentRolesUsersPost

> ApiV3SaasIdentRolesUsersPost(ctx, ident).UsersSaaSRolesListObject(usersSaaSRolesListObject).Execute()





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
    usersSaaSRolesListObject := *openapiclient.NewUsersSaaSRolesListObject([]string{"Users_example"}, []string{"Roles_example"}) // UsersSaaSRolesListObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentRolesUsersPost(context.Background(), ident).UsersSaaSRolesListObject(usersSaaSRolesListObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentRolesUsersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SaasIdentRolesUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usersSaaSRolesListObject** | [**UsersSaaSRolesListObject**](UsersSaaSRolesListObject.md) |  | 

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


## ApiV3SaasIdentUsersGet

> ApiV3SaasIdentUsersGet(ctx, ident, lastVisit, roles).Execute()





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
    lastVisit := "lastVisit_example" // string | Filter by last_visit date (optional)
    roles := "roles_example" // string | Filter by roles. Exclude users where there are no common roles in the list of required roles and list of user roles (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentUsersGet(context.Background(), ident, lastVisit, roles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**lastVisit** | **string** | Filter by last_visit date | 
**roles** | **string** | Filter by roles. Exclude users where there are no common roles in the list of required roles and list of user roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentUsersGetRequest struct via the builder pattern


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


## ApiV3SaasIdentUsersPost

> ApiV3SaasIdentUsersPost(ctx, ident).SaaSUserObjectInner(saaSUserObjectInner).Execute()





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
    saaSUserObjectInner := []openapiclient.SaaSUserObjectInner{*openapiclient.NewSaaSUserObjectInner("UserId_example", false)} // []SaaSUserObjectInner |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentUsersPost(context.Background(), ident).SaaSUserObjectInner(saaSUserObjectInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentUsersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SaasIdentUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saaSUserObjectInner** | [**[]SaaSUserObjectInner**](SaaSUserObjectInner.md) |  | 

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


## ApiV3SaasIdentUsersUserIdPost

> ApiV3SaasIdentUsersUserIdPost(ctx, ident, userId).SaaSUserManageObject(saaSUserManageObject).Execute()





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
    userId := "userId_example" // string | ID of User
    saaSUserManageObject := *openapiclient.NewSaaSUserManageObject("UserExternalId_example") // SaaSUserManageObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasIdentUsersUserIdPost(context.Background(), ident, userId).SaaSUserManageObject(saaSUserManageObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasIdentUsersUserIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** | ID of system object (assets, locations, ...) | 
**userId** | **string** | ID of User | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasIdentUsersUserIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **saaSUserManageObject** | [**SaaSUserManageObject**](SaaSUserManageObject.md) |  | 

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


## ApiV3SaasUsersActivatePatch

> []InlineResponse200 ApiV3SaasUsersActivatePatch(ctx).InlineRequest(inlineRequest).Execute()





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
    inlineRequest := *openapiclient.NewInlineRequest() // InlineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasUsersActivatePatch(context.Background()).InlineRequest(inlineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasUsersActivatePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3SaasUsersActivatePatch`: []InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `SaaSUsersApi.ApiV3SaasUsersActivatePatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasUsersActivatePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineRequest** | [**InlineRequest**](InlineRequest.md) |  | 

### Return type

[**[]InlineResponse200**](InlineResponse200.md)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasUsersDeactivatePatch

> []InlineResponse200 ApiV3SaasUsersDeactivatePatch(ctx).InlineRequest(inlineRequest).Execute()





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
    inlineRequest := *openapiclient.NewInlineRequest() // InlineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasUsersDeactivatePatch(context.Background()).InlineRequest(inlineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasUsersDeactivatePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3SaasUsersDeactivatePatch`: []InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `SaaSUsersApi.ApiV3SaasUsersDeactivatePatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasUsersDeactivatePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineRequest** | [**InlineRequest**](InlineRequest.md) |  | 

### Return type

[**[]InlineResponse200**](InlineResponse200.md)

### Authorization

[api_key](../README.md#api_key), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SaasUsersGet

> ApiV3SaasUsersGet(ctx, filter, text, roles).Execute()





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
    filter := "filter_example" // string | Regular API v3 filter expression (optional)
    text := "text_example" // string | Regular API v3 full text search expression (optional)
    roles := "roles_example" // string | Filter by roles. Exclude users where there are no common roles in the list of required roles and list of user roles (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaaSUsersApi.ApiV3SaasUsersGet(context.Background(), filter, text, roles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaaSUsersApi.ApiV3SaasUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Regular API v3 filter expression | 
**text** | **string** | Regular API v3 full text search expression | 
**roles** | **string** | Filter by roles. Exclude users where there are no common roles in the list of required roles and list of user roles | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SaasUsersGetRequest struct via the builder pattern


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

