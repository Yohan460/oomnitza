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

// Software struct for Software
type Software struct {
	SoftwareName *string `json:"software_name,omitempty"`
	Category *string `json:"category,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	NumberOfSeats *int32 `json:"number_of_seats,omitempty"`
	NumberOfUsers *int32 `json:"number_of_users,omitempty"`
	NumberOfAssets *int32 `json:"number_of_assets,omitempty"`
	ChangeDate *int32 `json:"change_date,omitempty"`
	ChangedBy *string `json:"changed_by,omitempty"`
	CreationDate *int32 `json:"creation_date,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Url *string `json:"url,omitempty"`
	AnnualSpend *float32 `json:"annual_spend,omitempty"`
	RenewalDate *int32 `json:"renewal_date,omitempty"`
	Hash *string `json:"hash,omitempty"`
	SyncKey *string `json:"sync_key,omitempty"`
	LinkKey *string `json:"link_key,omitempty"`
	SoftwareId *string `json:"software_id,omitempty"`
	Linked *bool `json:"linked,omitempty"`
	Ignored *bool `json:"ignored,omitempty"`
	Version *string `json:"version,omitempty"`
	Deleted *string `json:"deleted,omitempty"`
}

// NewSoftware instantiates a new Software object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftware() *Software {
	this := Software{}
	return &this
}

// NewSoftwareWithDefaults instantiates a new Software object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareWithDefaults() *Software {
	this := Software{}
	return &this
}

// GetSoftwareName returns the SoftwareName field value if set, zero value otherwise.
func (o *Software) GetSoftwareName() string {
	if o == nil || o.SoftwareName == nil {
		var ret string
		return ret
	}
	return *o.SoftwareName
}

// GetSoftwareNameOk returns a tuple with the SoftwareName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetSoftwareNameOk() (*string, bool) {
	if o == nil || o.SoftwareName == nil {
		return nil, false
	}
	return o.SoftwareName, true
}

// HasSoftwareName returns a boolean if a field has been set.
func (o *Software) HasSoftwareName() bool {
	if o != nil && o.SoftwareName != nil {
		return true
	}

	return false
}

// SetSoftwareName gets a reference to the given string and assigns it to the SoftwareName field.
func (o *Software) SetSoftwareName(v string) {
	o.SoftwareName = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Software) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Software) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Software) SetCategory(v string) {
	o.Category = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *Software) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetPublisherOk() (*string, bool) {
	if o == nil || o.Publisher == nil {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *Software) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *Software) SetPublisher(v string) {
	o.Publisher = &v
}

// GetNumberOfSeats returns the NumberOfSeats field value if set, zero value otherwise.
func (o *Software) GetNumberOfSeats() int32 {
	if o == nil || o.NumberOfSeats == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfSeats
}

// GetNumberOfSeatsOk returns a tuple with the NumberOfSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetNumberOfSeatsOk() (*int32, bool) {
	if o == nil || o.NumberOfSeats == nil {
		return nil, false
	}
	return o.NumberOfSeats, true
}

// HasNumberOfSeats returns a boolean if a field has been set.
func (o *Software) HasNumberOfSeats() bool {
	if o != nil && o.NumberOfSeats != nil {
		return true
	}

	return false
}

// SetNumberOfSeats gets a reference to the given int32 and assigns it to the NumberOfSeats field.
func (o *Software) SetNumberOfSeats(v int32) {
	o.NumberOfSeats = &v
}

// GetNumberOfUsers returns the NumberOfUsers field value if set, zero value otherwise.
func (o *Software) GetNumberOfUsers() int32 {
	if o == nil || o.NumberOfUsers == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfUsers
}

// GetNumberOfUsersOk returns a tuple with the NumberOfUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetNumberOfUsersOk() (*int32, bool) {
	if o == nil || o.NumberOfUsers == nil {
		return nil, false
	}
	return o.NumberOfUsers, true
}

// HasNumberOfUsers returns a boolean if a field has been set.
func (o *Software) HasNumberOfUsers() bool {
	if o != nil && o.NumberOfUsers != nil {
		return true
	}

	return false
}

// SetNumberOfUsers gets a reference to the given int32 and assigns it to the NumberOfUsers field.
func (o *Software) SetNumberOfUsers(v int32) {
	o.NumberOfUsers = &v
}

// GetNumberOfAssets returns the NumberOfAssets field value if set, zero value otherwise.
func (o *Software) GetNumberOfAssets() int32 {
	if o == nil || o.NumberOfAssets == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfAssets
}

// GetNumberOfAssetsOk returns a tuple with the NumberOfAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetNumberOfAssetsOk() (*int32, bool) {
	if o == nil || o.NumberOfAssets == nil {
		return nil, false
	}
	return o.NumberOfAssets, true
}

// HasNumberOfAssets returns a boolean if a field has been set.
func (o *Software) HasNumberOfAssets() bool {
	if o != nil && o.NumberOfAssets != nil {
		return true
	}

	return false
}

// SetNumberOfAssets gets a reference to the given int32 and assigns it to the NumberOfAssets field.
func (o *Software) SetNumberOfAssets(v int32) {
	o.NumberOfAssets = &v
}

// GetChangeDate returns the ChangeDate field value if set, zero value otherwise.
func (o *Software) GetChangeDate() int32 {
	if o == nil || o.ChangeDate == nil {
		var ret int32
		return ret
	}
	return *o.ChangeDate
}

// GetChangeDateOk returns a tuple with the ChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetChangeDateOk() (*int32, bool) {
	if o == nil || o.ChangeDate == nil {
		return nil, false
	}
	return o.ChangeDate, true
}

// HasChangeDate returns a boolean if a field has been set.
func (o *Software) HasChangeDate() bool {
	if o != nil && o.ChangeDate != nil {
		return true
	}

	return false
}

// SetChangeDate gets a reference to the given int32 and assigns it to the ChangeDate field.
func (o *Software) SetChangeDate(v int32) {
	o.ChangeDate = &v
}

// GetChangedBy returns the ChangedBy field value if set, zero value otherwise.
func (o *Software) GetChangedBy() string {
	if o == nil || o.ChangedBy == nil {
		var ret string
		return ret
	}
	return *o.ChangedBy
}

// GetChangedByOk returns a tuple with the ChangedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetChangedByOk() (*string, bool) {
	if o == nil || o.ChangedBy == nil {
		return nil, false
	}
	return o.ChangedBy, true
}

// HasChangedBy returns a boolean if a field has been set.
func (o *Software) HasChangedBy() bool {
	if o != nil && o.ChangedBy != nil {
		return true
	}

	return false
}

// SetChangedBy gets a reference to the given string and assigns it to the ChangedBy field.
func (o *Software) SetChangedBy(v string) {
	o.ChangedBy = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Software) GetCreationDate() int32 {
	if o == nil || o.CreationDate == nil {
		var ret int32
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetCreationDateOk() (*int32, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Software) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int32 and assigns it to the CreationDate field.
func (o *Software) SetCreationDate(v int32) {
	o.CreationDate = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Software) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Software) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Software) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Software) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Software) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Software) SetUrl(v string) {
	o.Url = &v
}

// GetAnnualSpend returns the AnnualSpend field value if set, zero value otherwise.
func (o *Software) GetAnnualSpend() float32 {
	if o == nil || o.AnnualSpend == nil {
		var ret float32
		return ret
	}
	return *o.AnnualSpend
}

// GetAnnualSpendOk returns a tuple with the AnnualSpend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetAnnualSpendOk() (*float32, bool) {
	if o == nil || o.AnnualSpend == nil {
		return nil, false
	}
	return o.AnnualSpend, true
}

// HasAnnualSpend returns a boolean if a field has been set.
func (o *Software) HasAnnualSpend() bool {
	if o != nil && o.AnnualSpend != nil {
		return true
	}

	return false
}

// SetAnnualSpend gets a reference to the given float32 and assigns it to the AnnualSpend field.
func (o *Software) SetAnnualSpend(v float32) {
	o.AnnualSpend = &v
}

// GetRenewalDate returns the RenewalDate field value if set, zero value otherwise.
func (o *Software) GetRenewalDate() int32 {
	if o == nil || o.RenewalDate == nil {
		var ret int32
		return ret
	}
	return *o.RenewalDate
}

// GetRenewalDateOk returns a tuple with the RenewalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetRenewalDateOk() (*int32, bool) {
	if o == nil || o.RenewalDate == nil {
		return nil, false
	}
	return o.RenewalDate, true
}

// HasRenewalDate returns a boolean if a field has been set.
func (o *Software) HasRenewalDate() bool {
	if o != nil && o.RenewalDate != nil {
		return true
	}

	return false
}

// SetRenewalDate gets a reference to the given int32 and assigns it to the RenewalDate field.
func (o *Software) SetRenewalDate(v int32) {
	o.RenewalDate = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Software) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Software) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Software) SetHash(v string) {
	o.Hash = &v
}

// GetSyncKey returns the SyncKey field value if set, zero value otherwise.
func (o *Software) GetSyncKey() string {
	if o == nil || o.SyncKey == nil {
		var ret string
		return ret
	}
	return *o.SyncKey
}

// GetSyncKeyOk returns a tuple with the SyncKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetSyncKeyOk() (*string, bool) {
	if o == nil || o.SyncKey == nil {
		return nil, false
	}
	return o.SyncKey, true
}

// HasSyncKey returns a boolean if a field has been set.
func (o *Software) HasSyncKey() bool {
	if o != nil && o.SyncKey != nil {
		return true
	}

	return false
}

// SetSyncKey gets a reference to the given string and assigns it to the SyncKey field.
func (o *Software) SetSyncKey(v string) {
	o.SyncKey = &v
}

// GetLinkKey returns the LinkKey field value if set, zero value otherwise.
func (o *Software) GetLinkKey() string {
	if o == nil || o.LinkKey == nil {
		var ret string
		return ret
	}
	return *o.LinkKey
}

// GetLinkKeyOk returns a tuple with the LinkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetLinkKeyOk() (*string, bool) {
	if o == nil || o.LinkKey == nil {
		return nil, false
	}
	return o.LinkKey, true
}

// HasLinkKey returns a boolean if a field has been set.
func (o *Software) HasLinkKey() bool {
	if o != nil && o.LinkKey != nil {
		return true
	}

	return false
}

// SetLinkKey gets a reference to the given string and assigns it to the LinkKey field.
func (o *Software) SetLinkKey(v string) {
	o.LinkKey = &v
}

// GetSoftwareId returns the SoftwareId field value if set, zero value otherwise.
func (o *Software) GetSoftwareId() string {
	if o == nil || o.SoftwareId == nil {
		var ret string
		return ret
	}
	return *o.SoftwareId
}

// GetSoftwareIdOk returns a tuple with the SoftwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetSoftwareIdOk() (*string, bool) {
	if o == nil || o.SoftwareId == nil {
		return nil, false
	}
	return o.SoftwareId, true
}

// HasSoftwareId returns a boolean if a field has been set.
func (o *Software) HasSoftwareId() bool {
	if o != nil && o.SoftwareId != nil {
		return true
	}

	return false
}

// SetSoftwareId gets a reference to the given string and assigns it to the SoftwareId field.
func (o *Software) SetSoftwareId(v string) {
	o.SoftwareId = &v
}

// GetLinked returns the Linked field value if set, zero value otherwise.
func (o *Software) GetLinked() bool {
	if o == nil || o.Linked == nil {
		var ret bool
		return ret
	}
	return *o.Linked
}

// GetLinkedOk returns a tuple with the Linked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetLinkedOk() (*bool, bool) {
	if o == nil || o.Linked == nil {
		return nil, false
	}
	return o.Linked, true
}

// HasLinked returns a boolean if a field has been set.
func (o *Software) HasLinked() bool {
	if o != nil && o.Linked != nil {
		return true
	}

	return false
}

// SetLinked gets a reference to the given bool and assigns it to the Linked field.
func (o *Software) SetLinked(v bool) {
	o.Linked = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *Software) GetIgnored() bool {
	if o == nil || o.Ignored == nil {
		var ret bool
		return ret
	}
	return *o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetIgnoredOk() (*bool, bool) {
	if o == nil || o.Ignored == nil {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *Software) HasIgnored() bool {
	if o != nil && o.Ignored != nil {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given bool and assigns it to the Ignored field.
func (o *Software) SetIgnored(v bool) {
	o.Ignored = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Software) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Software) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Software) SetVersion(v string) {
	o.Version = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Software) GetDeleted() string {
	if o == nil || o.Deleted == nil {
		var ret string
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Software) GetDeletedOk() (*string, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Software) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given string and assigns it to the Deleted field.
func (o *Software) SetDeleted(v string) {
	o.Deleted = &v
}

func (o Software) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SoftwareName != nil {
		toSerialize["software_name"] = o.SoftwareName
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Publisher != nil {
		toSerialize["publisher"] = o.Publisher
	}
	if o.NumberOfSeats != nil {
		toSerialize["number_of_seats"] = o.NumberOfSeats
	}
	if o.NumberOfUsers != nil {
		toSerialize["number_of_users"] = o.NumberOfUsers
	}
	if o.NumberOfAssets != nil {
		toSerialize["number_of_assets"] = o.NumberOfAssets
	}
	if o.ChangeDate != nil {
		toSerialize["change_date"] = o.ChangeDate
	}
	if o.ChangedBy != nil {
		toSerialize["changed_by"] = o.ChangedBy
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.AnnualSpend != nil {
		toSerialize["annual_spend"] = o.AnnualSpend
	}
	if o.RenewalDate != nil {
		toSerialize["renewal_date"] = o.RenewalDate
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.SyncKey != nil {
		toSerialize["sync_key"] = o.SyncKey
	}
	if o.LinkKey != nil {
		toSerialize["link_key"] = o.LinkKey
	}
	if o.SoftwareId != nil {
		toSerialize["software_id"] = o.SoftwareId
	}
	if o.Linked != nil {
		toSerialize["linked"] = o.Linked
	}
	if o.Ignored != nil {
		toSerialize["ignored"] = o.Ignored
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	return json.Marshal(toSerialize)
}

type NullableSoftware struct {
	value *Software
	isSet bool
}

func (v NullableSoftware) Get() *Software {
	return v.value
}

func (v *NullableSoftware) Set(val *Software) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftware) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftware(val *Software) *NullableSoftware {
	return &NullableSoftware{value: val, isSet: true}
}

func (v NullableSoftware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

