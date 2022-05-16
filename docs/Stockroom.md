# Stockroom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockroomId** | Pointer to **string** |  | [optional] 
**StockroomName** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**ShippingAddress** | Pointer to **string** |  | [optional] 
**DefaultSpoc** | Pointer to **string** |  | [optional] 
**AccessoriesMinLimit** | Pointer to **int32** |  | [optional] 
**AccessoriesMaxLimit** | Pointer to **int32** |  | [optional] 
**QuantityReplenishBy** | Pointer to **int32** |  | [optional] 
**LoanPeriod** | Pointer to **int32** |  | [optional] 
**ShowNewAndUsed** | Pointer to **bool** |  | [optional] 
**AllowReturnNotAssigned** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int32** |  | [optional] 
**ChangedBy** | Pointer to **string** |  | [optional] 
**ChangeDate** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 

## Methods

### NewStockroom

`func NewStockroom() *Stockroom`

NewStockroom instantiates a new Stockroom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockroomWithDefaults

`func NewStockroomWithDefaults() *Stockroom`

NewStockroomWithDefaults instantiates a new Stockroom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStockroomId

`func (o *Stockroom) GetStockroomId() string`

GetStockroomId returns the StockroomId field if non-nil, zero value otherwise.

### GetStockroomIdOk

`func (o *Stockroom) GetStockroomIdOk() (*string, bool)`

GetStockroomIdOk returns a tuple with the StockroomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockroomId

`func (o *Stockroom) SetStockroomId(v string)`

SetStockroomId sets StockroomId field to given value.

### HasStockroomId

`func (o *Stockroom) HasStockroomId() bool`

HasStockroomId returns a boolean if a field has been set.

### GetStockroomName

`func (o *Stockroom) GetStockroomName() string`

GetStockroomName returns the StockroomName field if non-nil, zero value otherwise.

### GetStockroomNameOk

`func (o *Stockroom) GetStockroomNameOk() (*string, bool)`

GetStockroomNameOk returns a tuple with the StockroomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockroomName

`func (o *Stockroom) SetStockroomName(v string)`

SetStockroomName sets StockroomName field to given value.

### HasStockroomName

`func (o *Stockroom) HasStockroomName() bool`

HasStockroomName returns a boolean if a field has been set.

### GetLocation

`func (o *Stockroom) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Stockroom) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Stockroom) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Stockroom) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetShippingAddress

`func (o *Stockroom) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Stockroom) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Stockroom) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Stockroom) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetDefaultSpoc

`func (o *Stockroom) GetDefaultSpoc() string`

GetDefaultSpoc returns the DefaultSpoc field if non-nil, zero value otherwise.

### GetDefaultSpocOk

`func (o *Stockroom) GetDefaultSpocOk() (*string, bool)`

GetDefaultSpocOk returns a tuple with the DefaultSpoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSpoc

`func (o *Stockroom) SetDefaultSpoc(v string)`

SetDefaultSpoc sets DefaultSpoc field to given value.

### HasDefaultSpoc

`func (o *Stockroom) HasDefaultSpoc() bool`

HasDefaultSpoc returns a boolean if a field has been set.

### GetAccessoriesMinLimit

`func (o *Stockroom) GetAccessoriesMinLimit() int32`

GetAccessoriesMinLimit returns the AccessoriesMinLimit field if non-nil, zero value otherwise.

### GetAccessoriesMinLimitOk

`func (o *Stockroom) GetAccessoriesMinLimitOk() (*int32, bool)`

GetAccessoriesMinLimitOk returns a tuple with the AccessoriesMinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessoriesMinLimit

`func (o *Stockroom) SetAccessoriesMinLimit(v int32)`

SetAccessoriesMinLimit sets AccessoriesMinLimit field to given value.

### HasAccessoriesMinLimit

`func (o *Stockroom) HasAccessoriesMinLimit() bool`

HasAccessoriesMinLimit returns a boolean if a field has been set.

### GetAccessoriesMaxLimit

`func (o *Stockroom) GetAccessoriesMaxLimit() int32`

GetAccessoriesMaxLimit returns the AccessoriesMaxLimit field if non-nil, zero value otherwise.

### GetAccessoriesMaxLimitOk

`func (o *Stockroom) GetAccessoriesMaxLimitOk() (*int32, bool)`

GetAccessoriesMaxLimitOk returns a tuple with the AccessoriesMaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessoriesMaxLimit

`func (o *Stockroom) SetAccessoriesMaxLimit(v int32)`

SetAccessoriesMaxLimit sets AccessoriesMaxLimit field to given value.

### HasAccessoriesMaxLimit

`func (o *Stockroom) HasAccessoriesMaxLimit() bool`

HasAccessoriesMaxLimit returns a boolean if a field has been set.

### GetQuantityReplenishBy

`func (o *Stockroom) GetQuantityReplenishBy() int32`

GetQuantityReplenishBy returns the QuantityReplenishBy field if non-nil, zero value otherwise.

### GetQuantityReplenishByOk

`func (o *Stockroom) GetQuantityReplenishByOk() (*int32, bool)`

GetQuantityReplenishByOk returns a tuple with the QuantityReplenishBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityReplenishBy

`func (o *Stockroom) SetQuantityReplenishBy(v int32)`

SetQuantityReplenishBy sets QuantityReplenishBy field to given value.

### HasQuantityReplenishBy

`func (o *Stockroom) HasQuantityReplenishBy() bool`

HasQuantityReplenishBy returns a boolean if a field has been set.

### GetLoanPeriod

`func (o *Stockroom) GetLoanPeriod() int32`

GetLoanPeriod returns the LoanPeriod field if non-nil, zero value otherwise.

### GetLoanPeriodOk

`func (o *Stockroom) GetLoanPeriodOk() (*int32, bool)`

GetLoanPeriodOk returns a tuple with the LoanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanPeriod

`func (o *Stockroom) SetLoanPeriod(v int32)`

SetLoanPeriod sets LoanPeriod field to given value.

### HasLoanPeriod

`func (o *Stockroom) HasLoanPeriod() bool`

HasLoanPeriod returns a boolean if a field has been set.

### GetShowNewAndUsed

`func (o *Stockroom) GetShowNewAndUsed() bool`

GetShowNewAndUsed returns the ShowNewAndUsed field if non-nil, zero value otherwise.

### GetShowNewAndUsedOk

`func (o *Stockroom) GetShowNewAndUsedOk() (*bool, bool)`

GetShowNewAndUsedOk returns a tuple with the ShowNewAndUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNewAndUsed

`func (o *Stockroom) SetShowNewAndUsed(v bool)`

SetShowNewAndUsed sets ShowNewAndUsed field to given value.

### HasShowNewAndUsed

`func (o *Stockroom) HasShowNewAndUsed() bool`

HasShowNewAndUsed returns a boolean if a field has been set.

### GetAllowReturnNotAssigned

`func (o *Stockroom) GetAllowReturnNotAssigned() bool`

GetAllowReturnNotAssigned returns the AllowReturnNotAssigned field if non-nil, zero value otherwise.

### GetAllowReturnNotAssignedOk

`func (o *Stockroom) GetAllowReturnNotAssignedOk() (*bool, bool)`

GetAllowReturnNotAssignedOk returns a tuple with the AllowReturnNotAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReturnNotAssigned

`func (o *Stockroom) SetAllowReturnNotAssigned(v bool)`

SetAllowReturnNotAssigned sets AllowReturnNotAssigned field to given value.

### HasAllowReturnNotAssigned

`func (o *Stockroom) HasAllowReturnNotAssigned() bool`

HasAllowReturnNotAssigned returns a boolean if a field has been set.

### GetLocked

`func (o *Stockroom) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Stockroom) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Stockroom) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Stockroom) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Stockroom) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Stockroom) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Stockroom) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Stockroom) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreationDate

`func (o *Stockroom) GetCreationDate() int32`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Stockroom) GetCreationDateOk() (*int32, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Stockroom) SetCreationDate(v int32)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Stockroom) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetChangedBy

`func (o *Stockroom) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *Stockroom) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *Stockroom) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.

### HasChangedBy

`func (o *Stockroom) HasChangedBy() bool`

HasChangedBy returns a boolean if a field has been set.

### GetChangeDate

`func (o *Stockroom) GetChangeDate() int32`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *Stockroom) GetChangeDateOk() (*int32, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *Stockroom) SetChangeDate(v int32)`

SetChangeDate sets ChangeDate field to given value.

### HasChangeDate

`func (o *Stockroom) HasChangeDate() bool`

HasChangeDate returns a boolean if a field has been set.

### GetVersion

`func (o *Stockroom) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Stockroom) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Stockroom) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Stockroom) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeleted

`func (o *Stockroom) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Stockroom) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Stockroom) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Stockroom) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUid

`func (o *Stockroom) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Stockroom) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Stockroom) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Stockroom) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


