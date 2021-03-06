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

// InlineRequest struct for InlineRequest
type InlineRequest struct {
	SaasUserIds []string `json:"saas_user_ids,omitempty"`
}

// NewInlineRequest instantiates a new InlineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineRequest() *InlineRequest {
	this := InlineRequest{}
	return &this
}

// NewInlineRequestWithDefaults instantiates a new InlineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineRequestWithDefaults() *InlineRequest {
	this := InlineRequest{}
	return &this
}

// GetSaasUserIds returns the SaasUserIds field value if set, zero value otherwise.
func (o *InlineRequest) GetSaasUserIds() []string {
	if o == nil || o.SaasUserIds == nil {
		var ret []string
		return ret
	}
	return o.SaasUserIds
}

// GetSaasUserIdsOk returns a tuple with the SaasUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineRequest) GetSaasUserIdsOk() ([]string, bool) {
	if o == nil || o.SaasUserIds == nil {
		return nil, false
	}
	return o.SaasUserIds, true
}

// HasSaasUserIds returns a boolean if a field has been set.
func (o *InlineRequest) HasSaasUserIds() bool {
	if o != nil && o.SaasUserIds != nil {
		return true
	}

	return false
}

// SetSaasUserIds gets a reference to the given []string and assigns it to the SaasUserIds field.
func (o *InlineRequest) SetSaasUserIds(v []string) {
	o.SaasUserIds = v
}

func (o InlineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SaasUserIds != nil {
		toSerialize["saas_user_ids"] = o.SaasUserIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineRequest struct {
	value *InlineRequest
	isSet bool
}

func (v NullableInlineRequest) Get() *InlineRequest {
	return v.value
}

func (v *NullableInlineRequest) Set(val *InlineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineRequest(val *InlineRequest) *NullableInlineRequest {
	return &NullableInlineRequest{value: val, isSet: true}
}

func (v NullableInlineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


