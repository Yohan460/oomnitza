# SaaS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaasName** | **string** |  | 
**Linked** | **string** |  | 
**Category** | Pointer to **string** |  | [optional] 
**UsageStatistic** | Pointer to **string** |  | [optional] 
**UsageHistory** | Pointer to **string** |  | [optional] 
**LinkKey** | **string** |  | 
**Publisher** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Source** | **string** |  | 
**Ignored** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Visited** | Pointer to **int32** |  | [optional] 
**AvailableLicenses** | Pointer to **int32** |  | [optional] 
**AnnualSpend** | Pointer to **float32** |  | [optional] 
**RenewalDate** | Pointer to **string** |  | [optional] 
**Changedate** | **string** |  | 
**Changedby** | **string** |  | 
**Creationdate** | **string** |  | 
**Createdby** | **string** |  | 
**SaasId** | Pointer to **int32** |  | [optional] 
**Hash** | **bool** |  | 
**TotalLicenses** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 
**IsInf** | Pointer to **bool** |  | [optional] 

## Methods

### NewSaaS

`func NewSaaS(saasName string, linked string, linkKey string, source string, ignored string, changedate string, changedby string, creationdate string, createdby string, hash bool, ) *SaaS`

NewSaaS instantiates a new SaaS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaaSWithDefaults

`func NewSaaSWithDefaults() *SaaS`

NewSaaSWithDefaults instantiates a new SaaS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaasName

`func (o *SaaS) GetSaasName() string`

GetSaasName returns the SaasName field if non-nil, zero value otherwise.

### GetSaasNameOk

`func (o *SaaS) GetSaasNameOk() (*string, bool)`

GetSaasNameOk returns a tuple with the SaasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasName

`func (o *SaaS) SetSaasName(v string)`

SetSaasName sets SaasName field to given value.


### GetLinked

`func (o *SaaS) GetLinked() string`

GetLinked returns the Linked field if non-nil, zero value otherwise.

### GetLinkedOk

`func (o *SaaS) GetLinkedOk() (*string, bool)`

GetLinkedOk returns a tuple with the Linked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinked

`func (o *SaaS) SetLinked(v string)`

SetLinked sets Linked field to given value.


### GetCategory

`func (o *SaaS) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SaaS) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SaaS) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SaaS) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetUsageStatistic

`func (o *SaaS) GetUsageStatistic() string`

GetUsageStatistic returns the UsageStatistic field if non-nil, zero value otherwise.

### GetUsageStatisticOk

`func (o *SaaS) GetUsageStatisticOk() (*string, bool)`

GetUsageStatisticOk returns a tuple with the UsageStatistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStatistic

`func (o *SaaS) SetUsageStatistic(v string)`

SetUsageStatistic sets UsageStatistic field to given value.

### HasUsageStatistic

`func (o *SaaS) HasUsageStatistic() bool`

HasUsageStatistic returns a boolean if a field has been set.

### GetUsageHistory

`func (o *SaaS) GetUsageHistory() string`

GetUsageHistory returns the UsageHistory field if non-nil, zero value otherwise.

### GetUsageHistoryOk

`func (o *SaaS) GetUsageHistoryOk() (*string, bool)`

GetUsageHistoryOk returns a tuple with the UsageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHistory

`func (o *SaaS) SetUsageHistory(v string)`

SetUsageHistory sets UsageHistory field to given value.

### HasUsageHistory

`func (o *SaaS) HasUsageHistory() bool`

HasUsageHistory returns a boolean if a field has been set.

### GetLinkKey

`func (o *SaaS) GetLinkKey() string`

GetLinkKey returns the LinkKey field if non-nil, zero value otherwise.

### GetLinkKeyOk

`func (o *SaaS) GetLinkKeyOk() (*string, bool)`

GetLinkKeyOk returns a tuple with the LinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkKey

`func (o *SaaS) SetLinkKey(v string)`

SetLinkKey sets LinkKey field to given value.


### GetPublisher

`func (o *SaaS) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *SaaS) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *SaaS) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *SaaS) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetUrl

`func (o *SaaS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SaaS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SaaS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SaaS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSource

`func (o *SaaS) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SaaS) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SaaS) SetSource(v string)`

SetSource sets Source field to given value.


### GetIgnored

`func (o *SaaS) GetIgnored() string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *SaaS) GetIgnoredOk() (*string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *SaaS) SetIgnored(v string)`

SetIgnored sets Ignored field to given value.


### GetVersion

`func (o *SaaS) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SaaS) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SaaS) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SaaS) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVisited

`func (o *SaaS) GetVisited() int32`

GetVisited returns the Visited field if non-nil, zero value otherwise.

### GetVisitedOk

`func (o *SaaS) GetVisitedOk() (*int32, bool)`

GetVisitedOk returns a tuple with the Visited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisited

`func (o *SaaS) SetVisited(v int32)`

SetVisited sets Visited field to given value.

### HasVisited

`func (o *SaaS) HasVisited() bool`

HasVisited returns a boolean if a field has been set.

### GetAvailableLicenses

`func (o *SaaS) GetAvailableLicenses() int32`

GetAvailableLicenses returns the AvailableLicenses field if non-nil, zero value otherwise.

### GetAvailableLicensesOk

`func (o *SaaS) GetAvailableLicensesOk() (*int32, bool)`

GetAvailableLicensesOk returns a tuple with the AvailableLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLicenses

`func (o *SaaS) SetAvailableLicenses(v int32)`

SetAvailableLicenses sets AvailableLicenses field to given value.

### HasAvailableLicenses

`func (o *SaaS) HasAvailableLicenses() bool`

HasAvailableLicenses returns a boolean if a field has been set.

### GetAnnualSpend

`func (o *SaaS) GetAnnualSpend() float32`

GetAnnualSpend returns the AnnualSpend field if non-nil, zero value otherwise.

### GetAnnualSpendOk

`func (o *SaaS) GetAnnualSpendOk() (*float32, bool)`

GetAnnualSpendOk returns a tuple with the AnnualSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualSpend

`func (o *SaaS) SetAnnualSpend(v float32)`

SetAnnualSpend sets AnnualSpend field to given value.

### HasAnnualSpend

`func (o *SaaS) HasAnnualSpend() bool`

HasAnnualSpend returns a boolean if a field has been set.

### GetRenewalDate

`func (o *SaaS) GetRenewalDate() string`

GetRenewalDate returns the RenewalDate field if non-nil, zero value otherwise.

### GetRenewalDateOk

`func (o *SaaS) GetRenewalDateOk() (*string, bool)`

GetRenewalDateOk returns a tuple with the RenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalDate

`func (o *SaaS) SetRenewalDate(v string)`

SetRenewalDate sets RenewalDate field to given value.

### HasRenewalDate

`func (o *SaaS) HasRenewalDate() bool`

HasRenewalDate returns a boolean if a field has been set.

### GetChangedate

`func (o *SaaS) GetChangedate() string`

GetChangedate returns the Changedate field if non-nil, zero value otherwise.

### GetChangedateOk

`func (o *SaaS) GetChangedateOk() (*string, bool)`

GetChangedateOk returns a tuple with the Changedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedate

`func (o *SaaS) SetChangedate(v string)`

SetChangedate sets Changedate field to given value.


### GetChangedby

`func (o *SaaS) GetChangedby() string`

GetChangedby returns the Changedby field if non-nil, zero value otherwise.

### GetChangedbyOk

`func (o *SaaS) GetChangedbyOk() (*string, bool)`

GetChangedbyOk returns a tuple with the Changedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedby

`func (o *SaaS) SetChangedby(v string)`

SetChangedby sets Changedby field to given value.


### GetCreationdate

`func (o *SaaS) GetCreationdate() string`

GetCreationdate returns the Creationdate field if non-nil, zero value otherwise.

### GetCreationdateOk

`func (o *SaaS) GetCreationdateOk() (*string, bool)`

GetCreationdateOk returns a tuple with the Creationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationdate

`func (o *SaaS) SetCreationdate(v string)`

SetCreationdate sets Creationdate field to given value.


### GetCreatedby

`func (o *SaaS) GetCreatedby() string`

GetCreatedby returns the Createdby field if non-nil, zero value otherwise.

### GetCreatedbyOk

`func (o *SaaS) GetCreatedbyOk() (*string, bool)`

GetCreatedbyOk returns a tuple with the Createdby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedby

`func (o *SaaS) SetCreatedby(v string)`

SetCreatedby sets Createdby field to given value.


### GetSaasId

`func (o *SaaS) GetSaasId() int32`

GetSaasId returns the SaasId field if non-nil, zero value otherwise.

### GetSaasIdOk

`func (o *SaaS) GetSaasIdOk() (*int32, bool)`

GetSaasIdOk returns a tuple with the SaasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasId

`func (o *SaaS) SetSaasId(v int32)`

SetSaasId sets SaasId field to given value.

### HasSaasId

`func (o *SaaS) HasSaasId() bool`

HasSaasId returns a boolean if a field has been set.

### GetHash

`func (o *SaaS) GetHash() bool`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SaaS) GetHashOk() (*bool, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SaaS) SetHash(v bool)`

SetHash sets Hash field to given value.


### GetTotalLicenses

`func (o *SaaS) GetTotalLicenses() int32`

GetTotalLicenses returns the TotalLicenses field if non-nil, zero value otherwise.

### GetTotalLicensesOk

`func (o *SaaS) GetTotalLicensesOk() (*int32, bool)`

GetTotalLicensesOk returns a tuple with the TotalLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLicenses

`func (o *SaaS) SetTotalLicenses(v int32)`

SetTotalLicenses sets TotalLicenses field to given value.

### HasTotalLicenses

`func (o *SaaS) HasTotalLicenses() bool`

HasTotalLicenses returns a boolean if a field has been set.

### GetDeleted

`func (o *SaaS) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SaaS) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SaaS) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SaaS) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUid

`func (o *SaaS) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SaaS) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SaaS) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SaaS) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetIsInf

`func (o *SaaS) GetIsInf() bool`

GetIsInf returns the IsInf field if non-nil, zero value otherwise.

### GetIsInfOk

`func (o *SaaS) GetIsInfOk() (*bool, bool)`

GetIsInfOk returns a tuple with the IsInf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInf

`func (o *SaaS) SetIsInf(v bool)`

SetIsInf sets IsInf field to given value.

### HasIsInf

`func (o *SaaS) HasIsInf() bool`

HasIsInf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


