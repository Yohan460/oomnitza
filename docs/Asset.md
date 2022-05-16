# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int32** |  | [optional] 
**ChangeDate** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ChangedBy** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Stockroom** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**Tickets** | Pointer to **string** |  | [optional] 
**IsParent** | Pointer to **string** |  | [optional] 
**IsChild** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**AssetType** | Pointer to **string** |  | [optional] 
**EndOfLifeDate** | Pointer to **int32** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**CurrentlyLoggedInUser** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**ImeiMeid** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**PurchaseDate** | Pointer to **int32** |  | [optional] 
**PoNumber** | Pointer to **string** |  | [optional] 
**PurchasePrice** | Pointer to **string** |  | [optional] 
**WarrantyEndDate** | Pointer to **int32** |  | [optional] 
**Cpu** | Pointer to **string** |  | [optional] 
**LoanDueDate** | Pointer to **int32** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentId

`func (o *Asset) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *Asset) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *Asset) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *Asset) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetVersion

`func (o *Asset) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Asset) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Asset) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Asset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationDate

`func (o *Asset) GetCreationDate() int32`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Asset) GetCreationDateOk() (*int32, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Asset) SetCreationDate(v int32)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Asset) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetChangeDate

`func (o *Asset) GetChangeDate() int32`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *Asset) GetChangeDateOk() (*int32, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *Asset) SetChangeDate(v int32)`

SetChangeDate sets ChangeDate field to given value.

### HasChangeDate

`func (o *Asset) HasChangeDate() bool`

HasChangeDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Asset) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Asset) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Asset) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Asset) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetChangedBy

`func (o *Asset) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *Asset) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *Asset) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.

### HasChangedBy

`func (o *Asset) HasChangedBy() bool`

HasChangedBy returns a boolean if a field has been set.

### GetLocation

`func (o *Asset) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Asset) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Asset) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Asset) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStockroom

`func (o *Asset) GetStockroom() string`

GetStockroom returns the Stockroom field if non-nil, zero value otherwise.

### GetStockroomOk

`func (o *Asset) GetStockroomOk() (*string, bool)`

GetStockroomOk returns a tuple with the Stockroom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockroom

`func (o *Asset) SetStockroom(v string)`

SetStockroom sets Stockroom field to given value.

### HasStockroom

`func (o *Asset) HasStockroom() bool`

HasStockroom returns a boolean if a field has been set.

### GetAssignedTo

`func (o *Asset) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *Asset) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *Asset) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *Asset) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetTickets

`func (o *Asset) GetTickets() string`

GetTickets returns the Tickets field if non-nil, zero value otherwise.

### GetTicketsOk

`func (o *Asset) GetTicketsOk() (*string, bool)`

GetTicketsOk returns a tuple with the Tickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickets

`func (o *Asset) SetTickets(v string)`

SetTickets sets Tickets field to given value.

### HasTickets

`func (o *Asset) HasTickets() bool`

HasTickets returns a boolean if a field has been set.

### GetIsParent

`func (o *Asset) GetIsParent() string`

GetIsParent returns the IsParent field if non-nil, zero value otherwise.

### GetIsParentOk

`func (o *Asset) GetIsParentOk() (*string, bool)`

GetIsParentOk returns a tuple with the IsParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParent

`func (o *Asset) SetIsParent(v string)`

SetIsParent sets IsParent field to given value.

### HasIsParent

`func (o *Asset) HasIsParent() bool`

HasIsParent returns a boolean if a field has been set.

### GetIsChild

`func (o *Asset) GetIsChild() string`

GetIsChild returns the IsChild field if non-nil, zero value otherwise.

### GetIsChildOk

`func (o *Asset) GetIsChildOk() (*string, bool)`

GetIsChildOk returns a tuple with the IsChild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChild

`func (o *Asset) SetIsChild(v string)`

SetIsChild sets IsChild field to given value.

### HasIsChild

`func (o *Asset) HasIsChild() bool`

HasIsChild returns a boolean if a field has been set.

### GetBarcode

`func (o *Asset) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Asset) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Asset) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Asset) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetDeleted

`func (o *Asset) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Asset) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Asset) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Asset) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Asset) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Asset) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Asset) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Asset) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetAssetType

`func (o *Asset) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Asset) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Asset) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *Asset) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetEndOfLifeDate

`func (o *Asset) GetEndOfLifeDate() int32`

GetEndOfLifeDate returns the EndOfLifeDate field if non-nil, zero value otherwise.

### GetEndOfLifeDateOk

`func (o *Asset) GetEndOfLifeDateOk() (*int32, bool)`

GetEndOfLifeDateOk returns a tuple with the EndOfLifeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeDate

`func (o *Asset) SetEndOfLifeDate(v int32)`

SetEndOfLifeDate sets EndOfLifeDate field to given value.

### HasEndOfLifeDate

`func (o *Asset) HasEndOfLifeDate() bool`

HasEndOfLifeDate returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Asset) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Asset) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Asset) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Asset) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetDepartment

`func (o *Asset) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Asset) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Asset) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Asset) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetStatus

`func (o *Asset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Asset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetManufacturer

`func (o *Asset) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Asset) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Asset) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *Asset) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *Asset) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Asset) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Asset) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Asset) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDeviceName

`func (o *Asset) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *Asset) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *Asset) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *Asset) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetCurrentlyLoggedInUser

`func (o *Asset) GetCurrentlyLoggedInUser() string`

GetCurrentlyLoggedInUser returns the CurrentlyLoggedInUser field if non-nil, zero value otherwise.

### GetCurrentlyLoggedInUserOk

`func (o *Asset) GetCurrentlyLoggedInUserOk() (*string, bool)`

GetCurrentlyLoggedInUserOk returns a tuple with the CurrentlyLoggedInUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyLoggedInUser

`func (o *Asset) SetCurrentlyLoggedInUser(v string)`

SetCurrentlyLoggedInUser sets CurrentlyLoggedInUser field to given value.

### HasCurrentlyLoggedInUser

`func (o *Asset) HasCurrentlyLoggedInUser() bool`

HasCurrentlyLoggedInUser returns a boolean if a field has been set.

### GetIpAddress

`func (o *Asset) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Asset) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Asset) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Asset) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *Asset) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Asset) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Asset) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Asset) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetImeiMeid

`func (o *Asset) GetImeiMeid() string`

GetImeiMeid returns the ImeiMeid field if non-nil, zero value otherwise.

### GetImeiMeidOk

`func (o *Asset) GetImeiMeidOk() (*string, bool)`

GetImeiMeidOk returns a tuple with the ImeiMeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImeiMeid

`func (o *Asset) SetImeiMeid(v string)`

SetImeiMeid sets ImeiMeid field to given value.

### HasImeiMeid

`func (o *Asset) HasImeiMeid() bool`

HasImeiMeid returns a boolean if a field has been set.

### GetVendor

`func (o *Asset) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Asset) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Asset) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Asset) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *Asset) GetPurchaseDate() int32`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *Asset) GetPurchaseDateOk() (*int32, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *Asset) SetPurchaseDate(v int32)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *Asset) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetPoNumber

`func (o *Asset) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *Asset) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *Asset) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *Asset) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *Asset) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *Asset) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *Asset) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *Asset) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetWarrantyEndDate

`func (o *Asset) GetWarrantyEndDate() int32`

GetWarrantyEndDate returns the WarrantyEndDate field if non-nil, zero value otherwise.

### GetWarrantyEndDateOk

`func (o *Asset) GetWarrantyEndDateOk() (*int32, bool)`

GetWarrantyEndDateOk returns a tuple with the WarrantyEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEndDate

`func (o *Asset) SetWarrantyEndDate(v int32)`

SetWarrantyEndDate sets WarrantyEndDate field to given value.

### HasWarrantyEndDate

`func (o *Asset) HasWarrantyEndDate() bool`

HasWarrantyEndDate returns a boolean if a field has been set.

### GetCpu

`func (o *Asset) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Asset) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Asset) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Asset) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetLoanDueDate

`func (o *Asset) GetLoanDueDate() int32`

GetLoanDueDate returns the LoanDueDate field if non-nil, zero value otherwise.

### GetLoanDueDateOk

`func (o *Asset) GetLoanDueDateOk() (*int32, bool)`

GetLoanDueDateOk returns a tuple with the LoanDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanDueDate

`func (o *Asset) SetLoanDueDate(v int32)`

SetLoanDueDate sets LoanDueDate field to given value.

### HasLoanDueDate

`func (o *Asset) HasLoanDueDate() bool`

HasLoanDueDate returns a boolean if a field has been set.

### GetNotes

`func (o *Asset) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Asset) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Asset) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Asset) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCondition

`func (o *Asset) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Asset) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Asset) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *Asset) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetUid

`func (o *Asset) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Asset) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Asset) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Asset) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


