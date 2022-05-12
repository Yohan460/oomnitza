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

// ConnectorRunLogsPortionDetailsLogs struct for ConnectorRunLogsPortionDetailsLogs
type ConnectorRunLogsPortionDetailsLogs struct {
	FatalError []string `json:"Fatal Error,omitempty"`
	Updated []string `json:"updated,omitempty"`
	Added []string `json:"added,omitempty"`
	Failed []string `json:"failed,omitempty"`
	Skipped []string `json:"skipped,omitempty"`
	Unchanged []string `json:"unchanged,omitempty"`
	Canceled []string `json:"canceled,omitempty"`
	Processing []string `json:"processing,omitempty"`
}

// NewConnectorRunLogsPortionDetailsLogs instantiates a new ConnectorRunLogsPortionDetailsLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRunLogsPortionDetailsLogs() *ConnectorRunLogsPortionDetailsLogs {
	this := ConnectorRunLogsPortionDetailsLogs{}
	return &this
}

// NewConnectorRunLogsPortionDetailsLogsWithDefaults instantiates a new ConnectorRunLogsPortionDetailsLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRunLogsPortionDetailsLogsWithDefaults() *ConnectorRunLogsPortionDetailsLogs {
	this := ConnectorRunLogsPortionDetailsLogs{}
	return &this
}

// GetFatalError returns the FatalError field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetFatalError() []string {
	if o == nil || o.FatalError == nil {
		var ret []string
		return ret
	}
	return o.FatalError
}

// GetFatalErrorOk returns a tuple with the FatalError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetFatalErrorOk() ([]string, bool) {
	if o == nil || o.FatalError == nil {
		return nil, false
	}
	return o.FatalError, true
}

// HasFatalError returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasFatalError() bool {
	if o != nil && o.FatalError != nil {
		return true
	}

	return false
}

// SetFatalError gets a reference to the given []string and assigns it to the FatalError field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetFatalError(v []string) {
	o.FatalError = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetUpdated() []string {
	if o == nil || o.Updated == nil {
		var ret []string
		return ret
	}
	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetUpdatedOk() ([]string, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given []string and assigns it to the Updated field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetUpdated(v []string) {
	o.Updated = v
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetAdded() []string {
	if o == nil || o.Added == nil {
		var ret []string
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetAddedOk() ([]string, bool) {
	if o == nil || o.Added == nil {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasAdded() bool {
	if o != nil && o.Added != nil {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []string and assigns it to the Added field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetAdded(v []string) {
	o.Added = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetFailed() []string {
	if o == nil || o.Failed == nil {
		var ret []string
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetFailedOk() ([]string, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []string and assigns it to the Failed field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetFailed(v []string) {
	o.Failed = v
}

// GetSkipped returns the Skipped field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetSkipped() []string {
	if o == nil || o.Skipped == nil {
		var ret []string
		return ret
	}
	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetSkippedOk() ([]string, bool) {
	if o == nil || o.Skipped == nil {
		return nil, false
	}
	return o.Skipped, true
}

// HasSkipped returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasSkipped() bool {
	if o != nil && o.Skipped != nil {
		return true
	}

	return false
}

// SetSkipped gets a reference to the given []string and assigns it to the Skipped field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetSkipped(v []string) {
	o.Skipped = v
}

// GetUnchanged returns the Unchanged field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetUnchanged() []string {
	if o == nil || o.Unchanged == nil {
		var ret []string
		return ret
	}
	return o.Unchanged
}

// GetUnchangedOk returns a tuple with the Unchanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetUnchangedOk() ([]string, bool) {
	if o == nil || o.Unchanged == nil {
		return nil, false
	}
	return o.Unchanged, true
}

// HasUnchanged returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasUnchanged() bool {
	if o != nil && o.Unchanged != nil {
		return true
	}

	return false
}

// SetUnchanged gets a reference to the given []string and assigns it to the Unchanged field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetUnchanged(v []string) {
	o.Unchanged = v
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetCanceled() []string {
	if o == nil || o.Canceled == nil {
		var ret []string
		return ret
	}
	return o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetCanceledOk() ([]string, bool) {
	if o == nil || o.Canceled == nil {
		return nil, false
	}
	return o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasCanceled() bool {
	if o != nil && o.Canceled != nil {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given []string and assigns it to the Canceled field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetCanceled(v []string) {
	o.Canceled = v
}

// GetProcessing returns the Processing field value if set, zero value otherwise.
func (o *ConnectorRunLogsPortionDetailsLogs) GetProcessing() []string {
	if o == nil || o.Processing == nil {
		var ret []string
		return ret
	}
	return o.Processing
}

// GetProcessingOk returns a tuple with the Processing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) GetProcessingOk() ([]string, bool) {
	if o == nil || o.Processing == nil {
		return nil, false
	}
	return o.Processing, true
}

// HasProcessing returns a boolean if a field has been set.
func (o *ConnectorRunLogsPortionDetailsLogs) HasProcessing() bool {
	if o != nil && o.Processing != nil {
		return true
	}

	return false
}

// SetProcessing gets a reference to the given []string and assigns it to the Processing field.
func (o *ConnectorRunLogsPortionDetailsLogs) SetProcessing(v []string) {
	o.Processing = v
}

func (o ConnectorRunLogsPortionDetailsLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FatalError != nil {
		toSerialize["Fatal Error"] = o.FatalError
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Added != nil {
		toSerialize["added"] = o.Added
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	if o.Skipped != nil {
		toSerialize["skipped"] = o.Skipped
	}
	if o.Unchanged != nil {
		toSerialize["unchanged"] = o.Unchanged
	}
	if o.Canceled != nil {
		toSerialize["canceled"] = o.Canceled
	}
	if o.Processing != nil {
		toSerialize["processing"] = o.Processing
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorRunLogsPortionDetailsLogs struct {
	value *ConnectorRunLogsPortionDetailsLogs
	isSet bool
}

func (v NullableConnectorRunLogsPortionDetailsLogs) Get() *ConnectorRunLogsPortionDetailsLogs {
	return v.value
}

func (v *NullableConnectorRunLogsPortionDetailsLogs) Set(val *ConnectorRunLogsPortionDetailsLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRunLogsPortionDetailsLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRunLogsPortionDetailsLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRunLogsPortionDetailsLogs(val *ConnectorRunLogsPortionDetailsLogs) *NullableConnectorRunLogsPortionDetailsLogs {
	return &NullableConnectorRunLogsPortionDetailsLogs{value: val, isSet: true}
}

func (v NullableConnectorRunLogsPortionDetailsLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRunLogsPortionDetailsLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

