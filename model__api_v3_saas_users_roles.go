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

// ApiV3SaasUsersRoles struct for ApiV3SaasUsersRoles
type ApiV3SaasUsersRoles struct {
	RoleId string `json:"role_id"`
	RoleExternalId string `json:"role_external_id"`
	RoleExternalLabel string `json:"role_external_label"`
}

// NewApiV3SaasUsersRoles instantiates a new ApiV3SaasUsersRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV3SaasUsersRoles(roleId string, roleExternalId string, roleExternalLabel string) *ApiV3SaasUsersRoles {
	this := ApiV3SaasUsersRoles{}
	this.RoleId = roleId
	this.RoleExternalId = roleExternalId
	this.RoleExternalLabel = roleExternalLabel
	return &this
}

// NewApiV3SaasUsersRolesWithDefaults instantiates a new ApiV3SaasUsersRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV3SaasUsersRolesWithDefaults() *ApiV3SaasUsersRoles {
	this := ApiV3SaasUsersRoles{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *ApiV3SaasUsersRoles) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *ApiV3SaasUsersRoles) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *ApiV3SaasUsersRoles) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleExternalId returns the RoleExternalId field value
func (o *ApiV3SaasUsersRoles) GetRoleExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleExternalId
}

// GetRoleExternalIdOk returns a tuple with the RoleExternalId field value
// and a boolean to check if the value has been set.
func (o *ApiV3SaasUsersRoles) GetRoleExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleExternalId, true
}

// SetRoleExternalId sets field value
func (o *ApiV3SaasUsersRoles) SetRoleExternalId(v string) {
	o.RoleExternalId = v
}

// GetRoleExternalLabel returns the RoleExternalLabel field value
func (o *ApiV3SaasUsersRoles) GetRoleExternalLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleExternalLabel
}

// GetRoleExternalLabelOk returns a tuple with the RoleExternalLabel field value
// and a boolean to check if the value has been set.
func (o *ApiV3SaasUsersRoles) GetRoleExternalLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleExternalLabel, true
}

// SetRoleExternalLabel sets field value
func (o *ApiV3SaasUsersRoles) SetRoleExternalLabel(v string) {
	o.RoleExternalLabel = v
}

func (o ApiV3SaasUsersRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role_id"] = o.RoleId
	}
	if true {
		toSerialize["role_external_id"] = o.RoleExternalId
	}
	if true {
		toSerialize["role_external_label"] = o.RoleExternalLabel
	}
	return json.Marshal(toSerialize)
}

type NullableApiV3SaasUsersRoles struct {
	value *ApiV3SaasUsersRoles
	isSet bool
}

func (v NullableApiV3SaasUsersRoles) Get() *ApiV3SaasUsersRoles {
	return v.value
}

func (v *NullableApiV3SaasUsersRoles) Set(val *ApiV3SaasUsersRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV3SaasUsersRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV3SaasUsersRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV3SaasUsersRoles(val *ApiV3SaasUsersRoles) *NullableApiV3SaasUsersRoles {
	return &NullableApiV3SaasUsersRoles{value: val, isSet: true}
}

func (v NullableApiV3SaasUsersRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV3SaasUsersRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

