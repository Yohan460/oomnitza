# MetadataField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalId** | Pointer to **NullableString** |  | [optional] 
**Uid** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**BelongsTo** | Pointer to **NullableString** |  | [optional] 
**DataType** | Pointer to **NullableString** |  | [optional] 
**Subtype** | Pointer to **NullableString** |  | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**Editable** | Pointer to **NullableString** |  | [optional] 
**Hardcode** | Pointer to **NullableString** |  | [optional] 
**Optional** | Pointer to **NullableString** |  | [optional] 
**DisplayOrder** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Unique** | **string** |  | 
**Searchhelp** | Pointer to [**[]MetadataFieldSearchhelpInner**](MetadataFieldSearchhelpInner.md) |  | [optional] 
**SearchhelpType** | Pointer to **NullableString** |  | [optional] 
**HelpText** | Pointer to **NullableString** |  | [optional] 
**Dependencies** | Pointer to [**NullableMetadataFieldDependencies**](MetadataFieldDependencies.md) |  | [optional] 
**Flags** | Pointer to **map[string]interface{}** |  | [optional] 
**DropdownHelpTextEnabled** | **string** |  | 

## Methods

### NewMetadataField

`func NewMetadataField(uid string, unique string, dropdownHelpTextEnabled string, ) *MetadataField`

NewMetadataField instantiates a new MetadataField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldWithDefaults

`func NewMetadataFieldWithDefaults() *MetadataField`

NewMetadataFieldWithDefaults instantiates a new MetadataField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalId

`func (o *MetadataField) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *MetadataField) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *MetadataField) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *MetadataField) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *MetadataField) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *MetadataField) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetUid

`func (o *MetadataField) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MetadataField) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MetadataField) SetUid(v string)`

SetUid sets Uid field to given value.


### GetLabel

`func (o *MetadataField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetadataField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetadataField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MetadataField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *MetadataField) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *MetadataField) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetBelongsTo

`func (o *MetadataField) GetBelongsTo() string`

GetBelongsTo returns the BelongsTo field if non-nil, zero value otherwise.

### GetBelongsToOk

`func (o *MetadataField) GetBelongsToOk() (*string, bool)`

GetBelongsToOk returns a tuple with the BelongsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelongsTo

`func (o *MetadataField) SetBelongsTo(v string)`

SetBelongsTo sets BelongsTo field to given value.

### HasBelongsTo

`func (o *MetadataField) HasBelongsTo() bool`

HasBelongsTo returns a boolean if a field has been set.

### SetBelongsToNil

`func (o *MetadataField) SetBelongsToNil(b bool)`

 SetBelongsToNil sets the value for BelongsTo to be an explicit nil

### UnsetBelongsTo
`func (o *MetadataField) UnsetBelongsTo()`

UnsetBelongsTo ensures that no value is present for BelongsTo, not even an explicit nil
### GetDataType

`func (o *MetadataField) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MetadataField) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MetadataField) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *MetadataField) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *MetadataField) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *MetadataField) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil
### GetSubtype

`func (o *MetadataField) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MetadataField) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MetadataField) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *MetadataField) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *MetadataField) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *MetadataField) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetLength

`func (o *MetadataField) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MetadataField) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *MetadataField) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *MetadataField) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *MetadataField) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *MetadataField) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetEditable

`func (o *MetadataField) GetEditable() string`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *MetadataField) GetEditableOk() (*string, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *MetadataField) SetEditable(v string)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *MetadataField) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *MetadataField) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *MetadataField) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetHardcode

`func (o *MetadataField) GetHardcode() string`

GetHardcode returns the Hardcode field if non-nil, zero value otherwise.

### GetHardcodeOk

`func (o *MetadataField) GetHardcodeOk() (*string, bool)`

GetHardcodeOk returns a tuple with the Hardcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardcode

`func (o *MetadataField) SetHardcode(v string)`

SetHardcode sets Hardcode field to given value.

### HasHardcode

`func (o *MetadataField) HasHardcode() bool`

HasHardcode returns a boolean if a field has been set.

### SetHardcodeNil

`func (o *MetadataField) SetHardcodeNil(b bool)`

 SetHardcodeNil sets the value for Hardcode to be an explicit nil

### UnsetHardcode
`func (o *MetadataField) UnsetHardcode()`

UnsetHardcode ensures that no value is present for Hardcode, not even an explicit nil
### GetOptional

`func (o *MetadataField) GetOptional() string`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *MetadataField) GetOptionalOk() (*string, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *MetadataField) SetOptional(v string)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *MetadataField) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *MetadataField) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *MetadataField) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetDisplayOrder

`func (o *MetadataField) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *MetadataField) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *MetadataField) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *MetadataField) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *MetadataField) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *MetadataField) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetDefaultValue

`func (o *MetadataField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetadataField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetadataField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetadataField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *MetadataField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *MetadataField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetUnique

`func (o *MetadataField) GetUnique() string`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *MetadataField) GetUniqueOk() (*string, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *MetadataField) SetUnique(v string)`

SetUnique sets Unique field to given value.


### GetSearchhelp

`func (o *MetadataField) GetSearchhelp() []MetadataFieldSearchhelpInner`

GetSearchhelp returns the Searchhelp field if non-nil, zero value otherwise.

### GetSearchhelpOk

`func (o *MetadataField) GetSearchhelpOk() (*[]MetadataFieldSearchhelpInner, bool)`

GetSearchhelpOk returns a tuple with the Searchhelp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchhelp

`func (o *MetadataField) SetSearchhelp(v []MetadataFieldSearchhelpInner)`

SetSearchhelp sets Searchhelp field to given value.

### HasSearchhelp

`func (o *MetadataField) HasSearchhelp() bool`

HasSearchhelp returns a boolean if a field has been set.

### SetSearchhelpNil

`func (o *MetadataField) SetSearchhelpNil(b bool)`

 SetSearchhelpNil sets the value for Searchhelp to be an explicit nil

### UnsetSearchhelp
`func (o *MetadataField) UnsetSearchhelp()`

UnsetSearchhelp ensures that no value is present for Searchhelp, not even an explicit nil
### GetSearchhelpType

`func (o *MetadataField) GetSearchhelpType() string`

GetSearchhelpType returns the SearchhelpType field if non-nil, zero value otherwise.

### GetSearchhelpTypeOk

`func (o *MetadataField) GetSearchhelpTypeOk() (*string, bool)`

GetSearchhelpTypeOk returns a tuple with the SearchhelpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchhelpType

`func (o *MetadataField) SetSearchhelpType(v string)`

SetSearchhelpType sets SearchhelpType field to given value.

### HasSearchhelpType

`func (o *MetadataField) HasSearchhelpType() bool`

HasSearchhelpType returns a boolean if a field has been set.

### SetSearchhelpTypeNil

`func (o *MetadataField) SetSearchhelpTypeNil(b bool)`

 SetSearchhelpTypeNil sets the value for SearchhelpType to be an explicit nil

### UnsetSearchhelpType
`func (o *MetadataField) UnsetSearchhelpType()`

UnsetSearchhelpType ensures that no value is present for SearchhelpType, not even an explicit nil
### GetHelpText

`func (o *MetadataField) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *MetadataField) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *MetadataField) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *MetadataField) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### SetHelpTextNil

`func (o *MetadataField) SetHelpTextNil(b bool)`

 SetHelpTextNil sets the value for HelpText to be an explicit nil

### UnsetHelpText
`func (o *MetadataField) UnsetHelpText()`

UnsetHelpText ensures that no value is present for HelpText, not even an explicit nil
### GetDependencies

`func (o *MetadataField) GetDependencies() MetadataFieldDependencies`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *MetadataField) GetDependenciesOk() (*MetadataFieldDependencies, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *MetadataField) SetDependencies(v MetadataFieldDependencies)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *MetadataField) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *MetadataField) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *MetadataField) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetFlags

`func (o *MetadataField) GetFlags() map[string]interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MetadataField) GetFlagsOk() (*map[string]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MetadataField) SetFlags(v map[string]interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MetadataField) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *MetadataField) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *MetadataField) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetDropdownHelpTextEnabled

`func (o *MetadataField) GetDropdownHelpTextEnabled() string`

GetDropdownHelpTextEnabled returns the DropdownHelpTextEnabled field if non-nil, zero value otherwise.

### GetDropdownHelpTextEnabledOk

`func (o *MetadataField) GetDropdownHelpTextEnabledOk() (*string, bool)`

GetDropdownHelpTextEnabledOk returns a tuple with the DropdownHelpTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropdownHelpTextEnabled

`func (o *MetadataField) SetDropdownHelpTextEnabled(v string)`

SetDropdownHelpTextEnabled sets DropdownHelpTextEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


