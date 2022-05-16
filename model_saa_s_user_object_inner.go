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

// SaaSUserObjectInner struct for SaaSUserObjectInner
type SaaSUserObjectInner struct {
	UserId string `json:"user_id"`
	Assign bool `json:"assign"`
}

// NewSaaSUserObjectInner instantiates a new SaaSUserObjectInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaaSUserObjectInner(userId string, assign bool) *SaaSUserObjectInner {
	this := SaaSUserObjectInner{}
	this.UserId = userId
	this.Assign = assign
	return &this
}

// NewSaaSUserObjectInnerWithDefaults instantiates a new SaaSUserObjectInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaaSUserObjectInnerWithDefaults() *SaaSUserObjectInner {
	this := SaaSUserObjectInner{}
	return &this
}

// GetUserId returns the UserId field value
func (o *SaaSUserObjectInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SaaSUserObjectInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SaaSUserObjectInner) SetUserId(v string) {
	o.UserId = v
}

// GetAssign returns the Assign field value
func (o *SaaSUserObjectInner) GetAssign() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Assign
}

// GetAssignOk returns a tuple with the Assign field value
// and a boolean to check if the value has been set.
func (o *SaaSUserObjectInner) GetAssignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assign, true
}

// SetAssign sets field value
func (o *SaaSUserObjectInner) SetAssign(v bool) {
	o.Assign = v
}

func (o SaaSUserObjectInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["assign"] = o.Assign
	}
	return json.Marshal(toSerialize)
}

type NullableSaaSUserObjectInner struct {
	value *SaaSUserObjectInner
	isSet bool
}

func (v NullableSaaSUserObjectInner) Get() *SaaSUserObjectInner {
	return v.value
}

func (v *NullableSaaSUserObjectInner) Set(val *SaaSUserObjectInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSaaSUserObjectInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSaaSUserObjectInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaaSUserObjectInner(val *SaaSUserObjectInner) *NullableSaaSUserObjectInner {
	return &NullableSaaSUserObjectInner{value: val, isSet: true}
}

func (v NullableSaaSUserObjectInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaaSUserObjectInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


