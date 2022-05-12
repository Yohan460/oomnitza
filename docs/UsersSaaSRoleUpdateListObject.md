# UsersSaaSRoleUpdateListObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **string** |  | 
**Roles** | [**[]UsersSaaSRoleUpdateListObjectRolesInner**](UsersSaaSRoleUpdateListObjectRolesInner.md) |  | 

## Methods

### NewUsersSaaSRoleUpdateListObject

`func NewUsersSaaSRoleUpdateListObject(users string, roles []UsersSaaSRoleUpdateListObjectRolesInner, ) *UsersSaaSRoleUpdateListObject`

NewUsersSaaSRoleUpdateListObject instantiates a new UsersSaaSRoleUpdateListObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersSaaSRoleUpdateListObjectWithDefaults

`func NewUsersSaaSRoleUpdateListObjectWithDefaults() *UsersSaaSRoleUpdateListObject`

NewUsersSaaSRoleUpdateListObjectWithDefaults instantiates a new UsersSaaSRoleUpdateListObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersSaaSRoleUpdateListObject) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersSaaSRoleUpdateListObject) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersSaaSRoleUpdateListObject) SetUsers(v string)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *UsersSaaSRoleUpdateListObject) GetRoles() []UsersSaaSRoleUpdateListObjectRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UsersSaaSRoleUpdateListObject) GetRolesOk() (*[]UsersSaaSRoleUpdateListObjectRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UsersSaaSRoleUpdateListObject) SetRoles(v []UsersSaaSRoleUpdateListObjectRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


