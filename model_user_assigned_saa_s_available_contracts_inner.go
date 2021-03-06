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

// UserAssignedSaaSAvailableContractsInner struct for UserAssignedSaaSAvailableContractsInner
type UserAssignedSaaSAvailableContractsInner struct {
	ContractId string `json:"contract_id"`
	ContractName string `json:"contract_name"`
	ContractType string `json:"contract_type"`
	HasFreeLicenseKeys bool `json:"has_free_license_keys"`
}

// NewUserAssignedSaaSAvailableContractsInner instantiates a new UserAssignedSaaSAvailableContractsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAssignedSaaSAvailableContractsInner(contractId string, contractName string, contractType string, hasFreeLicenseKeys bool) *UserAssignedSaaSAvailableContractsInner {
	this := UserAssignedSaaSAvailableContractsInner{}
	this.ContractId = contractId
	this.ContractName = contractName
	this.ContractType = contractType
	this.HasFreeLicenseKeys = hasFreeLicenseKeys
	return &this
}

// NewUserAssignedSaaSAvailableContractsInnerWithDefaults instantiates a new UserAssignedSaaSAvailableContractsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAssignedSaaSAvailableContractsInnerWithDefaults() *UserAssignedSaaSAvailableContractsInner {
	this := UserAssignedSaaSAvailableContractsInner{}
	return &this
}

// GetContractId returns the ContractId field value
func (o *UserAssignedSaaSAvailableContractsInner) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *UserAssignedSaaSAvailableContractsInner) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *UserAssignedSaaSAvailableContractsInner) SetContractId(v string) {
	o.ContractId = v
}

// GetContractName returns the ContractName field value
func (o *UserAssignedSaaSAvailableContractsInner) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *UserAssignedSaaSAvailableContractsInner) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *UserAssignedSaaSAvailableContractsInner) SetContractName(v string) {
	o.ContractName = v
}

// GetContractType returns the ContractType field value
func (o *UserAssignedSaaSAvailableContractsInner) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *UserAssignedSaaSAvailableContractsInner) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *UserAssignedSaaSAvailableContractsInner) SetContractType(v string) {
	o.ContractType = v
}

// GetHasFreeLicenseKeys returns the HasFreeLicenseKeys field value
func (o *UserAssignedSaaSAvailableContractsInner) GetHasFreeLicenseKeys() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasFreeLicenseKeys
}

// GetHasFreeLicenseKeysOk returns a tuple with the HasFreeLicenseKeys field value
// and a boolean to check if the value has been set.
func (o *UserAssignedSaaSAvailableContractsInner) GetHasFreeLicenseKeysOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasFreeLicenseKeys, true
}

// SetHasFreeLicenseKeys sets field value
func (o *UserAssignedSaaSAvailableContractsInner) SetHasFreeLicenseKeys(v bool) {
	o.HasFreeLicenseKeys = v
}

func (o UserAssignedSaaSAvailableContractsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contract_id"] = o.ContractId
	}
	if true {
		toSerialize["contract_name"] = o.ContractName
	}
	if true {
		toSerialize["contract_type"] = o.ContractType
	}
	if true {
		toSerialize["has_free_license_keys"] = o.HasFreeLicenseKeys
	}
	return json.Marshal(toSerialize)
}

type NullableUserAssignedSaaSAvailableContractsInner struct {
	value *UserAssignedSaaSAvailableContractsInner
	isSet bool
}

func (v NullableUserAssignedSaaSAvailableContractsInner) Get() *UserAssignedSaaSAvailableContractsInner {
	return v.value
}

func (v *NullableUserAssignedSaaSAvailableContractsInner) Set(val *UserAssignedSaaSAvailableContractsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAssignedSaaSAvailableContractsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAssignedSaaSAvailableContractsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAssignedSaaSAvailableContractsInner(val *UserAssignedSaaSAvailableContractsInner) *NullableUserAssignedSaaSAvailableContractsInner {
	return &NullableUserAssignedSaaSAvailableContractsInner{value: val, isSet: true}
}

func (v NullableUserAssignedSaaSAvailableContractsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAssignedSaaSAvailableContractsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


