# Go API client for openapi

## Date type fields
API endpoints expected date in UTC±0:00 timezone. Timezones in ISO8601 format will be ignored. API endpoints support date in two formats (one of): ISO8601 ('YYYY-MM-DDTHH:mm:SSZ') or Unix Timestamp (seconds count since January 1st, 1970 at UTC).

## Dropdown fields
Some fields are configured as dropdown fields with a dedicated list of values within Oomnitza. You can review the list of available dropdown values within the customization page in Oomnitza. In case you want to be able to post any data into these fields, you should switch them to dropdown without value within the customization page.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/yohan460/oomnitza"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessoriesApi* | [**ApiV3AccessoriesGet**](docs/AccessoriesApi.md#apiv3accessoriesget) | **Get** /api/v3/accessories | 
*AccessoriesApi* | [**ApiV3AccessoriesIdentChangesHistoryGet**](docs/AccessoriesApi.md#apiv3accessoriesidentchangeshistoryget) | **Get** /api/v3/accessories/{ident}/changes-history | 
*AccessoriesApi* | [**ApiV3AccessoriesIdentDelete**](docs/AccessoriesApi.md#apiv3accessoriesidentdelete) | **Delete** /api/v3/accessories/{ident} | 
*AccessoriesApi* | [**ApiV3AccessoriesIdentGet**](docs/AccessoriesApi.md#apiv3accessoriesidentget) | **Get** /api/v3/accessories/{ident} | 
*AccessoriesApi* | [**ApiV3AccessoriesIdentPatch**](docs/AccessoriesApi.md#apiv3accessoriesidentpatch) | **Patch** /api/v3/accessories/{ident} | 
*AccessoriesApi* | [**ApiV3AccessoriesPost**](docs/AccessoriesApi.md#apiv3accessoriespost) | **Post** /api/v3/accessories | 
*AccessoriesApi* | [**ApiV3AccessoriesSavedsearchesGet**](docs/AccessoriesApi.md#apiv3accessoriessavedsearchesget) | **Get** /api/v3/accessories/savedsearches | 
*AccessoriesApi* | [**ApiV3AccessoriesSavedsearchesIdentGet**](docs/AccessoriesApi.md#apiv3accessoriessavedsearchesidentget) | **Get** /api/v3/accessories/savedsearches/{ident} | 
*ActivitiesApi* | [**ApiV3ActivitiesGet**](docs/ActivitiesApi.md#apiv3activitiesget) | **Get** /api/v3/activities | 
*AssetsApi* | [**ApiV3AssetsGet**](docs/AssetsApi.md#apiv3assetsget) | **Get** /api/v3/assets | 
*AssetsApi* | [**ApiV3AssetsIdentChangesHistoryGet**](docs/AssetsApi.md#apiv3assetsidentchangeshistoryget) | **Get** /api/v3/assets/{ident}/changes-history | 
*AssetsApi* | [**ApiV3AssetsIdentDelete**](docs/AssetsApi.md#apiv3assetsidentdelete) | **Delete** /api/v3/assets/{ident} | 
*AssetsApi* | [**ApiV3AssetsIdentGet**](docs/AssetsApi.md#apiv3assetsidentget) | **Get** /api/v3/assets/{ident} | 
*AssetsApi* | [**ApiV3AssetsIdentPatch**](docs/AssetsApi.md#apiv3assetsidentpatch) | **Patch** /api/v3/assets/{ident} | 
*AssetsApi* | [**ApiV3AssetsIdentSoftwareGet**](docs/AssetsApi.md#apiv3assetsidentsoftwareget) | **Get** /api/v3/assets/{ident}/software | 
*AssetsApi* | [**ApiV3AssetsPost**](docs/AssetsApi.md#apiv3assetspost) | **Post** /api/v3/assets | 
*AssetsApi* | [**ApiV3AssetsSavedsearchesGet**](docs/AssetsApi.md#apiv3assetssavedsearchesget) | **Get** /api/v3/assets/savedsearches | 
*AssetsApi* | [**ApiV3AssetsSavedsearchesIdentGet**](docs/AssetsApi.md#apiv3assetssavedsearchesidentget) | **Get** /api/v3/assets/savedsearches/{ident} | 
*ConnectorApi* | [**ApiV3ConnectorRunLogsConnectorsIdentGet**](docs/ConnectorApi.md#apiv3connectorrunlogsconnectorsidentget) | **Get** /api/v3/connector_run_logs/connectors/{ident} | 
*ConnectorApi* | [**ApiV3ConnectorRunLogsIdentGet**](docs/ConnectorApi.md#apiv3connectorrunlogsidentget) | **Get** /api/v3/connector_run_logs/{ident} | 
*ContractsApi* | [**ApiV3ContractsGet**](docs/ContractsApi.md#apiv3contractsget) | **Get** /api/v3/contracts | 
*ContractsApi* | [**ApiV3ContractsIdentChangesHistoryGet**](docs/ContractsApi.md#apiv3contractsidentchangeshistoryget) | **Get** /api/v3/contracts/{ident}/changes-history | 
*ContractsApi* | [**ApiV3ContractsIdentDelete**](docs/ContractsApi.md#apiv3contractsidentdelete) | **Delete** /api/v3/contracts/{ident} | 
*ContractsApi* | [**ApiV3ContractsIdentGet**](docs/ContractsApi.md#apiv3contractsidentget) | **Get** /api/v3/contracts/{ident} | 
*ContractsApi* | [**ApiV3ContractsIdentPatch**](docs/ContractsApi.md#apiv3contractsidentpatch) | **Patch** /api/v3/contracts/{ident} | 
*ContractsApi* | [**ApiV3ContractsPost**](docs/ContractsApi.md#apiv3contractspost) | **Post** /api/v3/contracts | 
*ContractsApi* | [**ApiV3ContractsSavedsearchesGet**](docs/ContractsApi.md#apiv3contractssavedsearchesget) | **Get** /api/v3/contracts/savedsearches | 
*ContractsApi* | [**ApiV3ContractsSavedsearchesIdentGet**](docs/ContractsApi.md#apiv3contractssavedsearchesidentget) | **Get** /api/v3/contracts/savedsearches/{ident} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsGet**](docs/CustomObjectsApi.md#apiv3customobjectsget) | **Get** /api/v3/custom_objects | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentDelete**](docs/CustomObjectsApi.md#apiv3customobjectsidentdelete) | **Delete** /api/v3/custom_objects/{ident} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentGet**](docs/CustomObjectsApi.md#apiv3customobjectsidentget) | **Get** /api/v3/custom_objects/{ident} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentPatch**](docs/CustomObjectsApi.md#apiv3customobjectsidentpatch) | **Patch** /api/v3/custom_objects/{ident} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdChangesHistoryGet**](docs/CustomObjectsApi.md#apiv3customobjectsidentrecordscustomobjectrecordidchangeshistoryget) | **Get** /api/v3/custom_objects/{ident}/records/{custom_object_record_id}/changes-history | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdDelete**](docs/CustomObjectsApi.md#apiv3customobjectsidentrecordscustomobjectrecordiddelete) | **Delete** /api/v3/custom_objects/{ident}/records/{custom_object_record_id} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdGet**](docs/CustomObjectsApi.md#apiv3customobjectsidentrecordscustomobjectrecordidget) | **Get** /api/v3/custom_objects/{ident}/records/{custom_object_record_id} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentRecordsCustomObjectRecordIdPatch**](docs/CustomObjectsApi.md#apiv3customobjectsidentrecordscustomobjectrecordidpatch) | **Patch** /api/v3/custom_objects/{ident}/records/{custom_object_record_id} | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentRecordsGet**](docs/CustomObjectsApi.md#apiv3customobjectsidentrecordsget) | **Get** /api/v3/custom_objects/{ident}/records | 
*CustomObjectsApi* | [**ApiV3CustomObjectsIdentRecordsPost**](docs/CustomObjectsApi.md#apiv3customobjectsidentrecordspost) | **Post** /api/v3/custom_objects/{ident}/records | 
*CustomObjectsApi* | [**ApiV3CustomObjectsPost**](docs/CustomObjectsApi.md#apiv3customobjectspost) | **Post** /api/v3/custom_objects | 
*LocationsApi* | [**ApiV3LocationsGet**](docs/LocationsApi.md#apiv3locationsget) | **Get** /api/v3/locations | 
*LocationsApi* | [**ApiV3LocationsIdentChangesHistoryGet**](docs/LocationsApi.md#apiv3locationsidentchangeshistoryget) | **Get** /api/v3/locations/{ident}/changes-history | 
*LocationsApi* | [**ApiV3LocationsIdentDelete**](docs/LocationsApi.md#apiv3locationsidentdelete) | **Delete** /api/v3/locations/{ident} | 
*LocationsApi* | [**ApiV3LocationsIdentGet**](docs/LocationsApi.md#apiv3locationsidentget) | **Get** /api/v3/locations/{ident} | 
*LocationsApi* | [**ApiV3LocationsIdentPatch**](docs/LocationsApi.md#apiv3locationsidentpatch) | **Patch** /api/v3/locations/{ident} | 
*LocationsApi* | [**ApiV3LocationsPost**](docs/LocationsApi.md#apiv3locationspost) | **Post** /api/v3/locations | 
*LocationsApi* | [**ApiV3LocationsSavedsearchesGet**](docs/LocationsApi.md#apiv3locationssavedsearchesget) | **Get** /api/v3/locations/savedsearches | 
*LocationsApi* | [**ApiV3LocationsSavedsearchesIdentGet**](docs/LocationsApi.md#apiv3locationssavedsearchesidentget) | **Get** /api/v3/locations/savedsearches/{ident} | 
*SaaSApi* | [**ApiV3SaasGet**](docs/SaaSApi.md#apiv3saasget) | **Get** /api/v3/saas | 
*SaaSApi* | [**ApiV3SaasIdentChangesHistoryGet**](docs/SaaSApi.md#apiv3saasidentchangeshistoryget) | **Get** /api/v3/saas/{ident}/changes-history | 
*SaaSApi* | [**ApiV3SaasIdentDelete**](docs/SaaSApi.md#apiv3saasidentdelete) | **Delete** /api/v3/saas/{ident} | 
*SaaSApi* | [**ApiV3SaasIdentGet**](docs/SaaSApi.md#apiv3saasidentget) | **Get** /api/v3/saas/{ident} | 
*SaaSApi* | [**ApiV3SaasIdentPatch**](docs/SaaSApi.md#apiv3saasidentpatch) | **Patch** /api/v3/saas/{ident} | 
*SaaSApi* | [**ApiV3SaasPost**](docs/SaaSApi.md#apiv3saaspost) | **Post** /api/v3/saas | 
*SaaSApi* | [**ApiV3SaasSavedsearchesGet**](docs/SaaSApi.md#apiv3saassavedsearchesget) | **Get** /api/v3/saas/savedsearches | 
*SaaSApi* | [**ApiV3SaasSavedsearchesIdentGet**](docs/SaaSApi.md#apiv3saassavedsearchesidentget) | **Get** /api/v3/saas/savedsearches/{ident} | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesGet**](docs/SaaSUsersApi.md#apiv3saasidentrolesget) | **Get** /api/v3/saas/{ident}/roles | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesPost**](docs/SaaSUsersApi.md#apiv3saasidentrolespost) | **Post** /api/v3/saas/{ident}/roles | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesSaasRoleIdDelete**](docs/SaaSUsersApi.md#apiv3saasidentrolessaasroleiddelete) | **Delete** /api/v3/saas/{ident}/roles/{saas_role_id} | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesSaasRoleIdPatch**](docs/SaaSUsersApi.md#apiv3saasidentrolessaasroleidpatch) | **Patch** /api/v3/saas/{ident}/roles/{saas_role_id} | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesUsersDelete**](docs/SaaSUsersApi.md#apiv3saasidentrolesusersdelete) | **Delete** /api/v3/saas/{ident}/roles/users | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesUsersPatch**](docs/SaaSUsersApi.md#apiv3saasidentrolesuserspatch) | **Patch** /api/v3/saas/{ident}/roles/users | 
*SaaSUsersApi* | [**ApiV3SaasIdentRolesUsersPost**](docs/SaaSUsersApi.md#apiv3saasidentrolesuserspost) | **Post** /api/v3/saas/{ident}/roles/users | 
*SaaSUsersApi* | [**ApiV3SaasIdentUsersGet**](docs/SaaSUsersApi.md#apiv3saasidentusersget) | **Get** /api/v3/saas/{ident}/users | 
*SaaSUsersApi* | [**ApiV3SaasIdentUsersPost**](docs/SaaSUsersApi.md#apiv3saasidentuserspost) | **Post** /api/v3/saas/{ident}/users | 
*SaaSUsersApi* | [**ApiV3SaasIdentUsersUserIdPost**](docs/SaaSUsersApi.md#apiv3saasidentusersuseridpost) | **Post** /api/v3/saas/{ident}/users/{user_id} | 
*SaaSUsersApi* | [**ApiV3SaasUsersActivatePatch**](docs/SaaSUsersApi.md#apiv3saasusersactivatepatch) | **Patch** /api/v3/saas_users/activate | 
*SaaSUsersApi* | [**ApiV3SaasUsersDeactivatePatch**](docs/SaaSUsersApi.md#apiv3saasusersdeactivatepatch) | **Patch** /api/v3/saas_users/deactivate | 
*SaaSUsersApi* | [**ApiV3SaasUsersGet**](docs/SaaSUsersApi.md#apiv3saasusersget) | **Get** /api/v3/saas_users | 
*SoftwareApi* | [**ApiV3SoftwareGet**](docs/SoftwareApi.md#apiv3softwareget) | **Get** /api/v3/software | 
*SoftwareApi* | [**ApiV3SoftwareIdentAssetsGet**](docs/SoftwareApi.md#apiv3softwareidentassetsget) | **Get** /api/v3/software/{ident}/assets | 
*SoftwareApi* | [**ApiV3SoftwareIdentChangesHistoryGet**](docs/SoftwareApi.md#apiv3softwareidentchangeshistoryget) | **Get** /api/v3/software/{ident}/changes-history | 
*SoftwareApi* | [**ApiV3SoftwareIdentDelete**](docs/SoftwareApi.md#apiv3softwareidentdelete) | **Delete** /api/v3/software/{ident} | 
*SoftwareApi* | [**ApiV3SoftwareIdentGet**](docs/SoftwareApi.md#apiv3softwareidentget) | **Get** /api/v3/software/{ident} | 
*SoftwareApi* | [**ApiV3SoftwareIdentPatch**](docs/SoftwareApi.md#apiv3softwareidentpatch) | **Patch** /api/v3/software/{ident} | 
*SoftwareApi* | [**ApiV3SoftwarePost**](docs/SoftwareApi.md#apiv3softwarepost) | **Post** /api/v3/software | 
*SoftwareApi* | [**ApiV3SoftwareSavedsearchesGet**](docs/SoftwareApi.md#apiv3softwaresavedsearchesget) | **Get** /api/v3/software/savedsearches | 
*SoftwareApi* | [**ApiV3SoftwareSavedsearchesIdentGet**](docs/SoftwareApi.md#apiv3softwaresavedsearchesidentget) | **Get** /api/v3/software/savedsearches/{ident} | 
*StockroomsApi* | [**ApiV3StockroomsGet**](docs/StockroomsApi.md#apiv3stockroomsget) | **Get** /api/v3/stockrooms | 
*StockroomsApi* | [**ApiV3StockroomsIdentChangesHistoryGet**](docs/StockroomsApi.md#apiv3stockroomsidentchangeshistoryget) | **Get** /api/v3/stockrooms/{ident}/changes-history | 
*StockroomsApi* | [**ApiV3StockroomsIdentDelete**](docs/StockroomsApi.md#apiv3stockroomsidentdelete) | **Delete** /api/v3/stockrooms/{ident} | 
*StockroomsApi* | [**ApiV3StockroomsIdentGet**](docs/StockroomsApi.md#apiv3stockroomsidentget) | **Get** /api/v3/stockrooms/{ident} | 
*StockroomsApi* | [**ApiV3StockroomsIdentPatch**](docs/StockroomsApi.md#apiv3stockroomsidentpatch) | **Patch** /api/v3/stockrooms/{ident} | 
*StockroomsApi* | [**ApiV3StockroomsPost**](docs/StockroomsApi.md#apiv3stockroomspost) | **Post** /api/v3/stockrooms | 
*StockroomsApi* | [**ApiV3StockroomsSavedsearchesGet**](docs/StockroomsApi.md#apiv3stockroomssavedsearchesget) | **Get** /api/v3/stockrooms/savedsearches | 
*StockroomsApi* | [**ApiV3StockroomsSavedsearchesIdentGet**](docs/StockroomsApi.md#apiv3stockroomssavedsearchesidentget) | **Get** /api/v3/stockrooms/savedsearches/{ident} | 
*SubscriptionsApi* | [**ApiV3SubscriptionsGet**](docs/SubscriptionsApi.md#apiv3subscriptionsget) | **Get** /api/v3/subscriptions | 
*SubscriptionsApi* | [**ApiV3SubscriptionsIdentDelete**](docs/SubscriptionsApi.md#apiv3subscriptionsidentdelete) | **Delete** /api/v3/subscriptions/{ident} | 
*SubscriptionsApi* | [**ApiV3SubscriptionsIdentGet**](docs/SubscriptionsApi.md#apiv3subscriptionsidentget) | **Get** /api/v3/subscriptions/{ident} | 
*SubscriptionsApi* | [**ApiV3SubscriptionsIdentPatch**](docs/SubscriptionsApi.md#apiv3subscriptionsidentpatch) | **Patch** /api/v3/subscriptions/{ident} | 
*SubscriptionsApi* | [**ApiV3SubscriptionsPost**](docs/SubscriptionsApi.md#apiv3subscriptionspost) | **Post** /api/v3/subscriptions | 
*UsersApi* | [**ApiV3UsersGet**](docs/UsersApi.md#apiv3usersget) | **Get** /api/v3/users | 
*UsersApi* | [**ApiV3UsersIdentAssetsGet**](docs/UsersApi.md#apiv3usersidentassetsget) | **Get** /api/v3/users/{ident}/assets | 
*UsersApi* | [**ApiV3UsersIdentChangesHistoryGet**](docs/UsersApi.md#apiv3usersidentchangeshistoryget) | **Get** /api/v3/users/{ident}/changes-history | 
*UsersApi* | [**ApiV3UsersIdentDelete**](docs/UsersApi.md#apiv3usersidentdelete) | **Delete** /api/v3/users/{ident} | 
*UsersApi* | [**ApiV3UsersIdentGet**](docs/UsersApi.md#apiv3usersidentget) | **Get** /api/v3/users/{ident} | 
*UsersApi* | [**ApiV3UsersIdentPatch**](docs/UsersApi.md#apiv3usersidentpatch) | **Patch** /api/v3/users/{ident} | 
*UsersApi* | [**ApiV3UsersIdentProfileSettingsGet**](docs/UsersApi.md#apiv3usersidentprofilesettingsget) | **Get** /api/v3/users/{ident}/profile_settings | 
*UsersApi* | [**ApiV3UsersIdentSoftwareGet**](docs/UsersApi.md#apiv3usersidentsoftwareget) | **Get** /api/v3/users/{ident}/software | 
*UsersApi* | [**ApiV3UsersIdentSoftwareSaasGet**](docs/UsersApi.md#apiv3usersidentsoftwaresaasget) | **Get** /api/v3/users/{ident}/software/saas | 
*UsersApi* | [**ApiV3UsersPost**](docs/UsersApi.md#apiv3userspost) | **Post** /api/v3/users | 
*UsersApi* | [**ApiV3UsersSavedsearchesGet**](docs/UsersApi.md#apiv3userssavedsearchesget) | **Get** /api/v3/users/savedsearches | 
*UsersApi* | [**ApiV3UsersSavedsearchesIdentGet**](docs/UsersApi.md#apiv3userssavedsearchesidentget) | **Get** /api/v3/users/savedsearches/{ident} | 


## Documentation For Models

 - [Accessory](docs/Accessory.md)
 - [Asset](docs/Asset.md)
 - [ConnectorRunLogs](docs/ConnectorRunLogs.md)
 - [ConnectorRunLogsPortionDetails](docs/ConnectorRunLogsPortionDetails.md)
 - [ConnectorRunLogsPortionDetailsFileAttachment](docs/ConnectorRunLogsPortionDetailsFileAttachment.md)
 - [ConnectorRunLogsPortionDetailsLogs](docs/ConnectorRunLogsPortionDetailsLogs.md)
 - [ConnectorRunLogsPortionDetailsObjects](docs/ConnectorRunLogsPortionDetailsObjects.md)
 - [Contract](docs/Contract.md)
 - [CustomObject](docs/CustomObject.md)
 - [CustomObjectRecord](docs/CustomObjectRecord.md)
 - [InlineRequest](docs/InlineRequest.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Location](docs/Location.md)
 - [MetadataField](docs/MetadataField.md)
 - [MetadataFieldDependencies](docs/MetadataFieldDependencies.md)
 - [MetadataFieldSearchhelpInner](docs/MetadataFieldSearchhelpInner.md)
 - [MetadataFields](docs/MetadataFields.md)
 - [SaaS](docs/SaaS.md)
 - [SaaSAssignedUser](docs/SaaSAssignedUser.md)
 - [SaaSAssignedUserRoles](docs/SaaSAssignedUserRoles.md)
 - [SaaSCreateRoleObject](docs/SaaSCreateRoleObject.md)
 - [SaaSEditRoleObject](docs/SaaSEditRoleObject.md)
 - [SaaSRoles](docs/SaaSRoles.md)
 - [SaaSUserManageObject](docs/SaaSUserManageObject.md)
 - [SaaSUserObjectInner](docs/SaaSUserObjectInner.md)
 - [Software](docs/Software.md)
 - [Stockroom](docs/Stockroom.md)
 - [User](docs/User.md)
 - [UserAssignedSaaS](docs/UserAssignedSaaS.md)
 - [UserAssignedSaaSAvailableContractsInner](docs/UserAssignedSaaSAvailableContractsInner.md)
 - [UsersSaaSRoleUpdateListObject](docs/UsersSaaSRoleUpdateListObject.md)
 - [UsersSaaSRoleUpdateListObjectRolesInner](docs/UsersSaaSRoleUpdateListObjectRolesInner.md)
 - [UsersSaaSRolesListObject](docs/UsersSaaSRolesListObject.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: api_key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api_key and passed in as the auth context for each request.


### basic_auth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



