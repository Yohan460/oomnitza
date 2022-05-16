# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractName** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**SoftwareId** | Pointer to **string** |  | [optional] 
**SaasId** | Pointer to **string** |  | [optional] 
**DefaultSaasContract** | Pointer to **bool** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PaymentSchedule** | Pointer to **string** |  | [optional] 
**RenewalDate** | Pointer to **int32** |  | [optional] 
**ResponsiblePerson** | Pointer to **string** |  | [optional] 
**PaymentDate** | Pointer to **int32** |  | [optional] 
**Autorenewal** | Pointer to **bool** |  | [optional] 
**LimitNumberOfLicenses** | Pointer to **bool** |  | [optional] 
**NumberOfLicenses** | Pointer to **int32** |  | [optional] 
**NumberOfSeats** | Pointer to **int32** |  | [optional] 
**OversubscribeLicenses** | Pointer to **bool** |  | [optional] 
**LicenseKey** | Pointer to **string** |  | [optional] 
**CalculatePerUserPrice** | Pointer to **bool** |  | [optional] 
**ChangeDate** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ChangedBy** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LicensesInUse** | Pointer to **int32** |  | [optional] 
**ContractId** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractName

`func (o *Contract) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *Contract) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *Contract) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *Contract) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetContractType

`func (o *Contract) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *Contract) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *Contract) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *Contract) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetSoftwareId

`func (o *Contract) GetSoftwareId() string`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *Contract) GetSoftwareIdOk() (*string, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *Contract) SetSoftwareId(v string)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *Contract) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetSaasId

`func (o *Contract) GetSaasId() string`

GetSaasId returns the SaasId field if non-nil, zero value otherwise.

### GetSaasIdOk

`func (o *Contract) GetSaasIdOk() (*string, bool)`

GetSaasIdOk returns a tuple with the SaasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasId

`func (o *Contract) SetSaasId(v string)`

SetSaasId sets SaasId field to given value.

### HasSaasId

`func (o *Contract) HasSaasId() bool`

HasSaasId returns a boolean if a field has been set.

### GetDefaultSaasContract

`func (o *Contract) GetDefaultSaasContract() bool`

GetDefaultSaasContract returns the DefaultSaasContract field if non-nil, zero value otherwise.

### GetDefaultSaasContractOk

`func (o *Contract) GetDefaultSaasContractOk() (*bool, bool)`

GetDefaultSaasContractOk returns a tuple with the DefaultSaasContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSaasContract

`func (o *Contract) SetDefaultSaasContract(v bool)`

SetDefaultSaasContract sets DefaultSaasContract field to given value.

### HasDefaultSaasContract

`func (o *Contract) HasDefaultSaasContract() bool`

HasDefaultSaasContract returns a boolean if a field has been set.

### GetPrice

`func (o *Contract) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Contract) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Contract) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Contract) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPaymentSchedule

`func (o *Contract) GetPaymentSchedule() string`

GetPaymentSchedule returns the PaymentSchedule field if non-nil, zero value otherwise.

### GetPaymentScheduleOk

`func (o *Contract) GetPaymentScheduleOk() (*string, bool)`

GetPaymentScheduleOk returns a tuple with the PaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedule

`func (o *Contract) SetPaymentSchedule(v string)`

SetPaymentSchedule sets PaymentSchedule field to given value.

### HasPaymentSchedule

`func (o *Contract) HasPaymentSchedule() bool`

HasPaymentSchedule returns a boolean if a field has been set.

### GetRenewalDate

`func (o *Contract) GetRenewalDate() int32`

GetRenewalDate returns the RenewalDate field if non-nil, zero value otherwise.

### GetRenewalDateOk

`func (o *Contract) GetRenewalDateOk() (*int32, bool)`

GetRenewalDateOk returns a tuple with the RenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalDate

`func (o *Contract) SetRenewalDate(v int32)`

SetRenewalDate sets RenewalDate field to given value.

### HasRenewalDate

`func (o *Contract) HasRenewalDate() bool`

HasRenewalDate returns a boolean if a field has been set.

### GetResponsiblePerson

`func (o *Contract) GetResponsiblePerson() string`

GetResponsiblePerson returns the ResponsiblePerson field if non-nil, zero value otherwise.

### GetResponsiblePersonOk

`func (o *Contract) GetResponsiblePersonOk() (*string, bool)`

GetResponsiblePersonOk returns a tuple with the ResponsiblePerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePerson

`func (o *Contract) SetResponsiblePerson(v string)`

SetResponsiblePerson sets ResponsiblePerson field to given value.

### HasResponsiblePerson

`func (o *Contract) HasResponsiblePerson() bool`

HasResponsiblePerson returns a boolean if a field has been set.

### GetPaymentDate

`func (o *Contract) GetPaymentDate() int32`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *Contract) GetPaymentDateOk() (*int32, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *Contract) SetPaymentDate(v int32)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *Contract) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetAutorenewal

`func (o *Contract) GetAutorenewal() bool`

GetAutorenewal returns the Autorenewal field if non-nil, zero value otherwise.

### GetAutorenewalOk

`func (o *Contract) GetAutorenewalOk() (*bool, bool)`

GetAutorenewalOk returns a tuple with the Autorenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorenewal

`func (o *Contract) SetAutorenewal(v bool)`

SetAutorenewal sets Autorenewal field to given value.

### HasAutorenewal

`func (o *Contract) HasAutorenewal() bool`

HasAutorenewal returns a boolean if a field has been set.

### GetLimitNumberOfLicenses

`func (o *Contract) GetLimitNumberOfLicenses() bool`

GetLimitNumberOfLicenses returns the LimitNumberOfLicenses field if non-nil, zero value otherwise.

### GetLimitNumberOfLicensesOk

`func (o *Contract) GetLimitNumberOfLicensesOk() (*bool, bool)`

GetLimitNumberOfLicensesOk returns a tuple with the LimitNumberOfLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumberOfLicenses

`func (o *Contract) SetLimitNumberOfLicenses(v bool)`

SetLimitNumberOfLicenses sets LimitNumberOfLicenses field to given value.

### HasLimitNumberOfLicenses

`func (o *Contract) HasLimitNumberOfLicenses() bool`

HasLimitNumberOfLicenses returns a boolean if a field has been set.

### GetNumberOfLicenses

`func (o *Contract) GetNumberOfLicenses() int32`

GetNumberOfLicenses returns the NumberOfLicenses field if non-nil, zero value otherwise.

### GetNumberOfLicensesOk

`func (o *Contract) GetNumberOfLicensesOk() (*int32, bool)`

GetNumberOfLicensesOk returns a tuple with the NumberOfLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLicenses

`func (o *Contract) SetNumberOfLicenses(v int32)`

SetNumberOfLicenses sets NumberOfLicenses field to given value.

### HasNumberOfLicenses

`func (o *Contract) HasNumberOfLicenses() bool`

HasNumberOfLicenses returns a boolean if a field has been set.

### GetNumberOfSeats

`func (o *Contract) GetNumberOfSeats() int32`

GetNumberOfSeats returns the NumberOfSeats field if non-nil, zero value otherwise.

### GetNumberOfSeatsOk

`func (o *Contract) GetNumberOfSeatsOk() (*int32, bool)`

GetNumberOfSeatsOk returns a tuple with the NumberOfSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSeats

`func (o *Contract) SetNumberOfSeats(v int32)`

SetNumberOfSeats sets NumberOfSeats field to given value.

### HasNumberOfSeats

`func (o *Contract) HasNumberOfSeats() bool`

HasNumberOfSeats returns a boolean if a field has been set.

### GetOversubscribeLicenses

`func (o *Contract) GetOversubscribeLicenses() bool`

GetOversubscribeLicenses returns the OversubscribeLicenses field if non-nil, zero value otherwise.

### GetOversubscribeLicensesOk

`func (o *Contract) GetOversubscribeLicensesOk() (*bool, bool)`

GetOversubscribeLicensesOk returns a tuple with the OversubscribeLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribeLicenses

`func (o *Contract) SetOversubscribeLicenses(v bool)`

SetOversubscribeLicenses sets OversubscribeLicenses field to given value.

### HasOversubscribeLicenses

`func (o *Contract) HasOversubscribeLicenses() bool`

HasOversubscribeLicenses returns a boolean if a field has been set.

### GetLicenseKey

`func (o *Contract) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *Contract) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *Contract) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *Contract) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetCalculatePerUserPrice

`func (o *Contract) GetCalculatePerUserPrice() bool`

GetCalculatePerUserPrice returns the CalculatePerUserPrice field if non-nil, zero value otherwise.

### GetCalculatePerUserPriceOk

`func (o *Contract) GetCalculatePerUserPriceOk() (*bool, bool)`

GetCalculatePerUserPriceOk returns a tuple with the CalculatePerUserPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatePerUserPrice

`func (o *Contract) SetCalculatePerUserPrice(v bool)`

SetCalculatePerUserPrice sets CalculatePerUserPrice field to given value.

### HasCalculatePerUserPrice

`func (o *Contract) HasCalculatePerUserPrice() bool`

HasCalculatePerUserPrice returns a boolean if a field has been set.

### GetChangeDate

`func (o *Contract) GetChangeDate() int32`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *Contract) GetChangeDateOk() (*int32, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *Contract) SetChangeDate(v int32)`

SetChangeDate sets ChangeDate field to given value.

### HasChangeDate

`func (o *Contract) HasChangeDate() bool`

HasChangeDate returns a boolean if a field has been set.

### GetVersion

`func (o *Contract) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Contract) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Contract) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Contract) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetChangedBy

`func (o *Contract) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *Contract) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *Contract) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.

### HasChangedBy

`func (o *Contract) HasChangedBy() bool`

HasChangedBy returns a boolean if a field has been set.

### GetCreationDate

`func (o *Contract) GetCreationDate() int32`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Contract) GetCreationDateOk() (*int32, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Contract) SetCreationDate(v int32)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Contract) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Contract) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Contract) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Contract) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Contract) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLicensesInUse

`func (o *Contract) GetLicensesInUse() int32`

GetLicensesInUse returns the LicensesInUse field if non-nil, zero value otherwise.

### GetLicensesInUseOk

`func (o *Contract) GetLicensesInUseOk() (*int32, bool)`

GetLicensesInUseOk returns a tuple with the LicensesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensesInUse

`func (o *Contract) SetLicensesInUse(v int32)`

SetLicensesInUse sets LicensesInUse field to given value.

### HasLicensesInUse

`func (o *Contract) HasLicensesInUse() bool`

HasLicensesInUse returns a boolean if a field has been set.

### GetContractId

`func (o *Contract) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *Contract) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *Contract) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *Contract) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetDeleted

`func (o *Contract) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Contract) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Contract) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Contract) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUid

`func (o *Contract) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Contract) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Contract) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Contract) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


