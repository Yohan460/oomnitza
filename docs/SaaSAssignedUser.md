# SaaSAssignedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**FullName** | **string** |  | 
**LastVisit** | **int32** |  | 
**OldUsage** | **bool** |  | 
**UserId** | **string** |  | 
**Username** | **string** |  | 
**MediaUrl** | **string** |  | 
**ContractId** | **string** |  | 
**Roles** | [**SaaSAssignedUserRoles**](SaaSAssignedUserRoles.md) |  | 
**UserExternalId** | **string** |  | 

## Methods

### NewSaaSAssignedUser

`func NewSaaSAssignedUser(email string, fullName string, lastVisit int32, oldUsage bool, userId string, username string, mediaUrl string, contractId string, roles SaaSAssignedUserRoles, userExternalId string, ) *SaaSAssignedUser`

NewSaaSAssignedUser instantiates a new SaaSAssignedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaaSAssignedUserWithDefaults

`func NewSaaSAssignedUserWithDefaults() *SaaSAssignedUser`

NewSaaSAssignedUserWithDefaults instantiates a new SaaSAssignedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SaaSAssignedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SaaSAssignedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SaaSAssignedUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *SaaSAssignedUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SaaSAssignedUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SaaSAssignedUser) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLastVisit

`func (o *SaaSAssignedUser) GetLastVisit() int32`

GetLastVisit returns the LastVisit field if non-nil, zero value otherwise.

### GetLastVisitOk

`func (o *SaaSAssignedUser) GetLastVisitOk() (*int32, bool)`

GetLastVisitOk returns a tuple with the LastVisit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVisit

`func (o *SaaSAssignedUser) SetLastVisit(v int32)`

SetLastVisit sets LastVisit field to given value.


### GetOldUsage

`func (o *SaaSAssignedUser) GetOldUsage() bool`

GetOldUsage returns the OldUsage field if non-nil, zero value otherwise.

### GetOldUsageOk

`func (o *SaaSAssignedUser) GetOldUsageOk() (*bool, bool)`

GetOldUsageOk returns a tuple with the OldUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldUsage

`func (o *SaaSAssignedUser) SetOldUsage(v bool)`

SetOldUsage sets OldUsage field to given value.


### GetUserId

`func (o *SaaSAssignedUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SaaSAssignedUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SaaSAssignedUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *SaaSAssignedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SaaSAssignedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SaaSAssignedUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetMediaUrl

`func (o *SaaSAssignedUser) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *SaaSAssignedUser) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *SaaSAssignedUser) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.


### GetContractId

`func (o *SaaSAssignedUser) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *SaaSAssignedUser) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *SaaSAssignedUser) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetRoles

`func (o *SaaSAssignedUser) GetRoles() SaaSAssignedUserRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SaaSAssignedUser) GetRolesOk() (*SaaSAssignedUserRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SaaSAssignedUser) SetRoles(v SaaSAssignedUserRoles)`

SetRoles sets Roles field to given value.


### GetUserExternalId

`func (o *SaaSAssignedUser) GetUserExternalId() string`

GetUserExternalId returns the UserExternalId field if non-nil, zero value otherwise.

### GetUserExternalIdOk

`func (o *SaaSAssignedUser) GetUserExternalIdOk() (*string, bool)`

GetUserExternalIdOk returns a tuple with the UserExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExternalId

`func (o *SaaSAssignedUser) SetUserExternalId(v string)`

SetUserExternalId sets UserExternalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


