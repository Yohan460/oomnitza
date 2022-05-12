# Accessory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessoryId** | Pointer to **string** |  | [optional] 
**AccessoryName** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**LoanOnly** | Pointer to **bool** |  | [optional] 
**CreationDate** | Pointer to **int32** |  | [optional] 
**ChangeDate** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ChangedBy** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessory

`func NewAccessory() *Accessory`

NewAccessory instantiates a new Accessory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessoryWithDefaults

`func NewAccessoryWithDefaults() *Accessory`

NewAccessoryWithDefaults instantiates a new Accessory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessoryId

`func (o *Accessory) GetAccessoryId() string`

GetAccessoryId returns the AccessoryId field if non-nil, zero value otherwise.

### GetAccessoryIdOk

`func (o *Accessory) GetAccessoryIdOk() (*string, bool)`

GetAccessoryIdOk returns a tuple with the AccessoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessoryId

`func (o *Accessory) SetAccessoryId(v string)`

SetAccessoryId sets AccessoryId field to given value.

### HasAccessoryId

`func (o *Accessory) HasAccessoryId() bool`

HasAccessoryId returns a boolean if a field has been set.

### GetAccessoryName

`func (o *Accessory) GetAccessoryName() string`

GetAccessoryName returns the AccessoryName field if non-nil, zero value otherwise.

### GetAccessoryNameOk

`func (o *Accessory) GetAccessoryNameOk() (*string, bool)`

GetAccessoryNameOk returns a tuple with the AccessoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessoryName

`func (o *Accessory) SetAccessoryName(v string)`

SetAccessoryName sets AccessoryName field to given value.

### HasAccessoryName

`func (o *Accessory) HasAccessoryName() bool`

HasAccessoryName returns a boolean if a field has been set.

### GetCategory

`func (o *Accessory) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Accessory) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Accessory) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Accessory) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSku

`func (o *Accessory) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Accessory) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Accessory) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Accessory) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPrice

`func (o *Accessory) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Accessory) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Accessory) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Accessory) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLoanOnly

`func (o *Accessory) GetLoanOnly() bool`

GetLoanOnly returns the LoanOnly field if non-nil, zero value otherwise.

### GetLoanOnlyOk

`func (o *Accessory) GetLoanOnlyOk() (*bool, bool)`

GetLoanOnlyOk returns a tuple with the LoanOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanOnly

`func (o *Accessory) SetLoanOnly(v bool)`

SetLoanOnly sets LoanOnly field to given value.

### HasLoanOnly

`func (o *Accessory) HasLoanOnly() bool`

HasLoanOnly returns a boolean if a field has been set.

### GetCreationDate

`func (o *Accessory) GetCreationDate() int32`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Accessory) GetCreationDateOk() (*int32, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Accessory) SetCreationDate(v int32)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Accessory) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetChangeDate

`func (o *Accessory) GetChangeDate() int32`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *Accessory) GetChangeDateOk() (*int32, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *Accessory) SetChangeDate(v int32)`

SetChangeDate sets ChangeDate field to given value.

### HasChangeDate

`func (o *Accessory) HasChangeDate() bool`

HasChangeDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Accessory) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Accessory) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Accessory) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Accessory) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetChangedBy

`func (o *Accessory) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *Accessory) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *Accessory) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.

### HasChangedBy

`func (o *Accessory) HasChangedBy() bool`

HasChangedBy returns a boolean if a field has been set.

### GetVersion

`func (o *Accessory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Accessory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Accessory) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Accessory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeleted

`func (o *Accessory) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Accessory) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Accessory) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Accessory) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUid

`func (o *Accessory) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Accessory) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Accessory) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Accessory) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


