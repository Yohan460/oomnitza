# CustomObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomObjectName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IsMainObject** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewCustomObject

`func NewCustomObject(customObjectName string, type_ string, ) *CustomObject`

NewCustomObject instantiates a new CustomObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectWithDefaults

`func NewCustomObjectWithDefaults() *CustomObject`

NewCustomObjectWithDefaults instantiates a new CustomObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomObjectName

`func (o *CustomObject) GetCustomObjectName() string`

GetCustomObjectName returns the CustomObjectName field if non-nil, zero value otherwise.

### GetCustomObjectNameOk

`func (o *CustomObject) GetCustomObjectNameOk() (*string, bool)`

GetCustomObjectNameOk returns a tuple with the CustomObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomObjectName

`func (o *CustomObject) SetCustomObjectName(v string)`

SetCustomObjectName sets CustomObjectName field to given value.


### GetDescription

`func (o *CustomObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsMainObject

`func (o *CustomObject) GetIsMainObject() bool`

GetIsMainObject returns the IsMainObject field if non-nil, zero value otherwise.

### GetIsMainObjectOk

`func (o *CustomObject) GetIsMainObjectOk() (*bool, bool)`

GetIsMainObjectOk returns a tuple with the IsMainObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMainObject

`func (o *CustomObject) SetIsMainObject(v bool)`

SetIsMainObject sets IsMainObject field to given value.

### HasIsMainObject

`func (o *CustomObject) HasIsMainObject() bool`

HasIsMainObject returns a boolean if a field has been set.

### GetType

`func (o *CustomObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomObject) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


