/*
External API

## Date type fields API endpoints expected date in UTC±0:00 timezone. Timezones in ISO8601 format will be ignored. API endpoints support date in two formats (one of): ISO8601 ('YYYY-MM-DDTHH:mm:SSZ') or Unix Timestamp (seconds count since January 1st, 1970 at UTC).  ## Dropdown fields Some fields are configured as dropdown fields with a dedicated list of values within Oomnitza. You can review the list of available dropdown values within the customization page in Oomnitza. In case you want to be able to post any data into these fields, you should switch them to dropdown without value within the customization page. 

API version: 3.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	SaasUserId *string `json:"saas_user_id,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetSaasUserId returns the SaasUserId field value if set, zero value otherwise.
func (o *InlineResponse200) GetSaasUserId() string {
	if o == nil || o.SaasUserId == nil {
		var ret string
		return ret
	}
	return *o.SaasUserId
}

// GetSaasUserIdOk returns a tuple with the SaasUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetSaasUserIdOk() (*string, bool) {
	if o == nil || o.SaasUserId == nil {
		return nil, false
	}
	return o.SaasUserId, true
}

// HasSaasUserId returns a boolean if a field has been set.
func (o *InlineResponse200) HasSaasUserId() bool {
	if o != nil && o.SaasUserId != nil {
		return true
	}

	return false
}

// SetSaasUserId gets a reference to the given string and assigns it to the SaasUserId field.
func (o *InlineResponse200) SetSaasUserId(v string) {
	o.SaasUserId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse200) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse200) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse200) SetError(v string) {
	o.Error = &v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SaasUserId != nil {
		toSerialize["saas_user_id"] = o.SaasUserId
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


