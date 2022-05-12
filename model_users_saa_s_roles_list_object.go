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

// UsersSaaSRolesListObject struct for UsersSaaSRolesListObject
type UsersSaaSRolesListObject struct {
	Users []string `json:"users"`
	Roles []string `json:"roles"`
}

// NewUsersSaaSRolesListObject instantiates a new UsersSaaSRolesListObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersSaaSRolesListObject(users []string, roles []string) *UsersSaaSRolesListObject {
	this := UsersSaaSRolesListObject{}
	this.Users = users
	this.Roles = roles
	return &this
}

// NewUsersSaaSRolesListObjectWithDefaults instantiates a new UsersSaaSRolesListObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersSaaSRolesListObjectWithDefaults() *UsersSaaSRolesListObject {
	this := UsersSaaSRolesListObject{}
	return &this
}

// GetUsers returns the Users field value
func (o *UsersSaaSRolesListObject) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UsersSaaSRolesListObject) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UsersSaaSRolesListObject) SetUsers(v []string) {
	o.Users = v
}

// GetRoles returns the Roles field value
func (o *UsersSaaSRolesListObject) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UsersSaaSRolesListObject) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *UsersSaaSRolesListObject) SetRoles(v []string) {
	o.Roles = v
}

func (o UsersSaaSRolesListObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableUsersSaaSRolesListObject struct {
	value *UsersSaaSRolesListObject
	isSet bool
}

func (v NullableUsersSaaSRolesListObject) Get() *UsersSaaSRolesListObject {
	return v.value
}

func (v *NullableUsersSaaSRolesListObject) Set(val *UsersSaaSRolesListObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersSaaSRolesListObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersSaaSRolesListObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersSaaSRolesListObject(val *UsersSaaSRolesListObject) *NullableUsersSaaSRolesListObject {
	return &NullableUsersSaaSRolesListObject{value: val, isSet: true}
}

func (v NullableUsersSaaSRolesListObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersSaaSRolesListObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


