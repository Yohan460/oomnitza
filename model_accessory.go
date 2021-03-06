/*
External API

## Date type fields API endpoints expected date in UTC±0:00 timezone. Timezones in ISO8601 format will be ignored. API endpoints support date in two formats (one of): ISO8601 ('YYYY-MM-DDTHH:mm:SSZ') or Unix Timestamp (seconds count since January 1st, 1970 at UTC).  ## Dropdown fields Some fields are configured as dropdown fields with a dedicated list of values within Oomnitza. You can review the list of available dropdown values within the customization page in Oomnitza. In case you want to be able to post any data into these fields, you should switch them to dropdown without value within the customization page. 

API version: 3.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oomnitza

import (
	"encoding/json"
)

// Accessory struct for Accessory
type Accessory struct {
	AccessoryId *string `json:"accessory_id,omitempty"`
	AccessoryName *string `json:"accessory_name,omitempty"`
	Category *string `json:"category,omitempty"`
	Sku *string `json:"sku,omitempty"`
	Price *float32 `json:"price,omitempty"`
	LoanOnly *bool `json:"loan_only,omitempty"`
	CreationDate *int32 `json:"creation_date,omitempty"`
	ChangeDate *int32 `json:"change_date,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	ChangedBy *string `json:"changed_by,omitempty"`
	Version *string `json:"version,omitempty"`
	Deleted *string `json:"deleted,omitempty"`
	Uid *int32 `json:"uid,omitempty"`
}

// NewAccessory instantiates a new Accessory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessory() *Accessory {
	this := Accessory{}
	return &this
}

// NewAccessoryWithDefaults instantiates a new Accessory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessoryWithDefaults() *Accessory {
	this := Accessory{}
	return &this
}

// GetAccessoryId returns the AccessoryId field value if set, zero value otherwise.
func (o *Accessory) GetAccessoryId() string {
	if o == nil || o.AccessoryId == nil {
		var ret string
		return ret
	}
	return *o.AccessoryId
}

// GetAccessoryIdOk returns a tuple with the AccessoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetAccessoryIdOk() (*string, bool) {
	if o == nil || o.AccessoryId == nil {
		return nil, false
	}
	return o.AccessoryId, true
}

// HasAccessoryId returns a boolean if a field has been set.
func (o *Accessory) HasAccessoryId() bool {
	if o != nil && o.AccessoryId != nil {
		return true
	}

	return false
}

// SetAccessoryId gets a reference to the given string and assigns it to the AccessoryId field.
func (o *Accessory) SetAccessoryId(v string) {
	o.AccessoryId = &v
}

// GetAccessoryName returns the AccessoryName field value if set, zero value otherwise.
func (o *Accessory) GetAccessoryName() string {
	if o == nil || o.AccessoryName == nil {
		var ret string
		return ret
	}
	return *o.AccessoryName
}

// GetAccessoryNameOk returns a tuple with the AccessoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetAccessoryNameOk() (*string, bool) {
	if o == nil || o.AccessoryName == nil {
		return nil, false
	}
	return o.AccessoryName, true
}

// HasAccessoryName returns a boolean if a field has been set.
func (o *Accessory) HasAccessoryName() bool {
	if o != nil && o.AccessoryName != nil {
		return true
	}

	return false
}

// SetAccessoryName gets a reference to the given string and assigns it to the AccessoryName field.
func (o *Accessory) SetAccessoryName(v string) {
	o.AccessoryName = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Accessory) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Accessory) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Accessory) SetCategory(v string) {
	o.Category = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Accessory) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Accessory) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Accessory) SetSku(v string) {
	o.Sku = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Accessory) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Accessory) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *Accessory) SetPrice(v float32) {
	o.Price = &v
}

// GetLoanOnly returns the LoanOnly field value if set, zero value otherwise.
func (o *Accessory) GetLoanOnly() bool {
	if o == nil || o.LoanOnly == nil {
		var ret bool
		return ret
	}
	return *o.LoanOnly
}

// GetLoanOnlyOk returns a tuple with the LoanOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetLoanOnlyOk() (*bool, bool) {
	if o == nil || o.LoanOnly == nil {
		return nil, false
	}
	return o.LoanOnly, true
}

// HasLoanOnly returns a boolean if a field has been set.
func (o *Accessory) HasLoanOnly() bool {
	if o != nil && o.LoanOnly != nil {
		return true
	}

	return false
}

// SetLoanOnly gets a reference to the given bool and assigns it to the LoanOnly field.
func (o *Accessory) SetLoanOnly(v bool) {
	o.LoanOnly = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Accessory) GetCreationDate() int32 {
	if o == nil || o.CreationDate == nil {
		var ret int32
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetCreationDateOk() (*int32, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Accessory) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int32 and assigns it to the CreationDate field.
func (o *Accessory) SetCreationDate(v int32) {
	o.CreationDate = &v
}

// GetChangeDate returns the ChangeDate field value if set, zero value otherwise.
func (o *Accessory) GetChangeDate() int32 {
	if o == nil || o.ChangeDate == nil {
		var ret int32
		return ret
	}
	return *o.ChangeDate
}

// GetChangeDateOk returns a tuple with the ChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetChangeDateOk() (*int32, bool) {
	if o == nil || o.ChangeDate == nil {
		return nil, false
	}
	return o.ChangeDate, true
}

// HasChangeDate returns a boolean if a field has been set.
func (o *Accessory) HasChangeDate() bool {
	if o != nil && o.ChangeDate != nil {
		return true
	}

	return false
}

// SetChangeDate gets a reference to the given int32 and assigns it to the ChangeDate field.
func (o *Accessory) SetChangeDate(v int32) {
	o.ChangeDate = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Accessory) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Accessory) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Accessory) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetChangedBy returns the ChangedBy field value if set, zero value otherwise.
func (o *Accessory) GetChangedBy() string {
	if o == nil || o.ChangedBy == nil {
		var ret string
		return ret
	}
	return *o.ChangedBy
}

// GetChangedByOk returns a tuple with the ChangedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetChangedByOk() (*string, bool) {
	if o == nil || o.ChangedBy == nil {
		return nil, false
	}
	return o.ChangedBy, true
}

// HasChangedBy returns a boolean if a field has been set.
func (o *Accessory) HasChangedBy() bool {
	if o != nil && o.ChangedBy != nil {
		return true
	}

	return false
}

// SetChangedBy gets a reference to the given string and assigns it to the ChangedBy field.
func (o *Accessory) SetChangedBy(v string) {
	o.ChangedBy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Accessory) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Accessory) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Accessory) SetVersion(v string) {
	o.Version = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Accessory) GetDeleted() string {
	if o == nil || o.Deleted == nil {
		var ret string
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetDeletedOk() (*string, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Accessory) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given string and assigns it to the Deleted field.
func (o *Accessory) SetDeleted(v string) {
	o.Deleted = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *Accessory) GetUid() int32 {
	if o == nil || o.Uid == nil {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accessory) GetUidOk() (*int32, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *Accessory) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *Accessory) SetUid(v int32) {
	o.Uid = &v
}

func (o Accessory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessoryId != nil {
		toSerialize["accessory_id"] = o.AccessoryId
	}
	if o.AccessoryName != nil {
		toSerialize["accessory_name"] = o.AccessoryName
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.LoanOnly != nil {
		toSerialize["loan_only"] = o.LoanOnly
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.ChangeDate != nil {
		toSerialize["change_date"] = o.ChangeDate
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.ChangedBy != nil {
		toSerialize["changed_by"] = o.ChangedBy
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableAccessory struct {
	value *Accessory
	isSet bool
}

func (v NullableAccessory) Get() *Accessory {
	return v.value
}

func (v *NullableAccessory) Set(val *Accessory) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessory) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessory(val *Accessory) *NullableAccessory {
	return &NullableAccessory{value: val, isSet: true}
}

func (v NullableAccessory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


