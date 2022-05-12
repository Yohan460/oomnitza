# UserAssignedSaaS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedContractId** | **string** |  | 
**AvailableContracts** | [**[]UserAssignedSaaSAvailableContractsInner**](UserAssignedSaaSAvailableContractsInner.md) |  | 
**Roles** | [**[]SaaSAssignedUserRoles**](SaaSAssignedUserRoles.md) |  | 
**EstimatedSpend** | **int32** |  | 
**SaasId** | **string** |  | 
**SaasName** | **string** |  | 
**Uid** | **int32** |  | 

## Methods

### NewUserAssignedSaaS

`func NewUserAssignedSaaS(assignedContractId string, availableContracts []UserAssignedSaaSAvailableContractsInner, roles []SaaSAssignedUserRoles, estimatedSpend int32, saasId string, saasName string, uid int32, ) *UserAssignedSaaS`

NewUserAssignedSaaS instantiates a new UserAssignedSaaS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAssignedSaaSWithDefaults

`func NewUserAssignedSaaSWithDefaults() *UserAssignedSaaS`

NewUserAssignedSaaSWithDefaults instantiates a new UserAssignedSaaS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedContractId

`func (o *UserAssignedSaaS) GetAssignedContractId() string`

GetAssignedContractId returns the AssignedContractId field if non-nil, zero value otherwise.

### GetAssignedContractIdOk

`func (o *UserAssignedSaaS) GetAssignedContractIdOk() (*string, bool)`

GetAssignedContractIdOk returns a tuple with the AssignedContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedContractId

`func (o *UserAssignedSaaS) SetAssignedContractId(v string)`

SetAssignedContractId sets AssignedContractId field to given value.


### GetAvailableContracts

`func (o *UserAssignedSaaS) GetAvailableContracts() []UserAssignedSaaSAvailableContractsInner`

GetAvailableContracts returns the AvailableContracts field if non-nil, zero value otherwise.

### GetAvailableContractsOk

`func (o *UserAssignedSaaS) GetAvailableContractsOk() (*[]UserAssignedSaaSAvailableContractsInner, bool)`

GetAvailableContractsOk returns a tuple with the AvailableContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableContracts

`func (o *UserAssignedSaaS) SetAvailableContracts(v []UserAssignedSaaSAvailableContractsInner)`

SetAvailableContracts sets AvailableContracts field to given value.


### GetRoles

`func (o *UserAssignedSaaS) GetRoles() []SaaSAssignedUserRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAssignedSaaS) GetRolesOk() (*[]SaaSAssignedUserRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAssignedSaaS) SetRoles(v []SaaSAssignedUserRoles)`

SetRoles sets Roles field to given value.


### GetEstimatedSpend

`func (o *UserAssignedSaaS) GetEstimatedSpend() int32`

GetEstimatedSpend returns the EstimatedSpend field if non-nil, zero value otherwise.

### GetEstimatedSpendOk

`func (o *UserAssignedSaaS) GetEstimatedSpendOk() (*int32, bool)`

GetEstimatedSpendOk returns a tuple with the EstimatedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSpend

`func (o *UserAssignedSaaS) SetEstimatedSpend(v int32)`

SetEstimatedSpend sets EstimatedSpend field to given value.


### GetSaasId

`func (o *UserAssignedSaaS) GetSaasId() string`

GetSaasId returns the SaasId field if non-nil, zero value otherwise.

### GetSaasIdOk

`func (o *UserAssignedSaaS) GetSaasIdOk() (*string, bool)`

GetSaasIdOk returns a tuple with the SaasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasId

`func (o *UserAssignedSaaS) SetSaasId(v string)`

SetSaasId sets SaasId field to given value.


### GetSaasName

`func (o *UserAssignedSaaS) GetSaasName() string`

GetSaasName returns the SaasName field if non-nil, zero value otherwise.

### GetSaasNameOk

`func (o *UserAssignedSaaS) GetSaasNameOk() (*string, bool)`

GetSaasNameOk returns a tuple with the SaasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasName

`func (o *UserAssignedSaaS) SetSaasName(v string)`

SetSaasName sets SaasName field to given value.


### GetUid

`func (o *UserAssignedSaaS) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserAssignedSaaS) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserAssignedSaaS) SetUid(v int32)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


