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

// SaaSEditRoleObject struct for SaaSEditRoleObject
type SaaSEditRoleObject struct {
	RoleExternalLabel *string `json:"role_external_label,omitempty"`
}

// NewSaaSEditRoleObject instantiates a new SaaSEditRoleObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaaSEditRoleObject() *SaaSEditRoleObject {
	this := SaaSEditRoleObject{}
	return &this
}

// NewSaaSEditRoleObjectWithDefaults instantiates a new SaaSEditRoleObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaaSEditRoleObjectWithDefaults() *SaaSEditRoleObject {
	this := SaaSEditRoleObject{}
	return &this
}

// GetRoleExternalLabel returns the RoleExternalLabel field value if set, zero value otherwise.
func (o *SaaSEditRoleObject) GetRoleExternalLabel() string {
	if o == nil || o.RoleExternalLabel == nil {
		var ret string
		return ret
	}
	return *o.RoleExternalLabel
}

// GetRoleExternalLabelOk returns a tuple with the RoleExternalLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaaSEditRoleObject) GetRoleExternalLabelOk() (*string, bool) {
	if o == nil || o.RoleExternalLabel == nil {
		return nil, false
	}
	return o.RoleExternalLabel, true
}

// HasRoleExternalLabel returns a boolean if a field has been set.
func (o *SaaSEditRoleObject) HasRoleExternalLabel() bool {
	if o != nil && o.RoleExternalLabel != nil {
		return true
	}

	return false
}

// SetRoleExternalLabel gets a reference to the given string and assigns it to the RoleExternalLabel field.
func (o *SaaSEditRoleObject) SetRoleExternalLabel(v string) {
	o.RoleExternalLabel = &v
}

func (o SaaSEditRoleObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleExternalLabel != nil {
		toSerialize["role_external_label"] = o.RoleExternalLabel
	}
	return json.Marshal(toSerialize)
}

type NullableSaaSEditRoleObject struct {
	value *SaaSEditRoleObject
	isSet bool
}

func (v NullableSaaSEditRoleObject) Get() *SaaSEditRoleObject {
	return v.value
}

func (v *NullableSaaSEditRoleObject) Set(val *SaaSEditRoleObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSaaSEditRoleObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSaaSEditRoleObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaaSEditRoleObject(val *SaaSEditRoleObject) *NullableSaaSEditRoleObject {
	return &NullableSaaSEditRoleObject{value: val, isSet: true}
}

func (v NullableSaaSEditRoleObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaaSEditRoleObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


