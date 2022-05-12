# ConnectorRunLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortionId** | **int32** |  | 
**PortionCorrelationId** | **string** |  | 
**StartTime** | **float32** |  | 
**EndTime** | **float32** |  | 
**Connector** | **string** |  | 
**IsMediaStorageService** | **bool** |  | 
**Status** | **string** |  | 
**TestRun** | **bool** |  | 
**Deployment** | **string** |  | 
**Ipv4Address** | **string** |  | 
**AddedRecordsCount** | **int32** |  | 
**UpdatedRecordsCount** | **int32** |  | 
**SkippedRecordsCount** | **int32** |  | 
**ErrorLogsCount** | **int32** |  | 

## Methods

### NewConnectorRunLogs

`func NewConnectorRunLogs(portionId int32, portionCorrelationId string, startTime float32, endTime float32, connector string, isMediaStorageService bool, status string, testRun bool, deployment string, ipv4Address string, addedRecordsCount int32, updatedRecordsCount int32, skippedRecordsCount int32, errorLogsCount int32, ) *ConnectorRunLogs`

NewConnectorRunLogs instantiates a new ConnectorRunLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorRunLogsWithDefaults

`func NewConnectorRunLogsWithDefaults() *ConnectorRunLogs`

NewConnectorRunLogsWithDefaults instantiates a new ConnectorRunLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortionId

`func (o *ConnectorRunLogs) GetPortionId() int32`

GetPortionId returns the PortionId field if non-nil, zero value otherwise.

### GetPortionIdOk

`func (o *ConnectorRunLogs) GetPortionIdOk() (*int32, bool)`

GetPortionIdOk returns a tuple with the PortionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortionId

`func (o *ConnectorRunLogs) SetPortionId(v int32)`

SetPortionId sets PortionId field to given value.


### GetPortionCorrelationId

`func (o *ConnectorRunLogs) GetPortionCorrelationId() string`

GetPortionCorrelationId returns the PortionCorrelationId field if non-nil, zero value otherwise.

### GetPortionCorrelationIdOk

`func (o *ConnectorRunLogs) GetPortionCorrelationIdOk() (*string, bool)`

GetPortionCorrelationIdOk returns a tuple with the PortionCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortionCorrelationId

`func (o *ConnectorRunLogs) SetPortionCorrelationId(v string)`

SetPortionCorrelationId sets PortionCorrelationId field to given value.


### GetStartTime

`func (o *ConnectorRunLogs) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ConnectorRunLogs) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ConnectorRunLogs) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ConnectorRunLogs) GetEndTime() float32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ConnectorRunLogs) GetEndTimeOk() (*float32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ConnectorRunLogs) SetEndTime(v float32)`

SetEndTime sets EndTime field to given value.


### GetConnector

`func (o *ConnectorRunLogs) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectorRunLogs) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectorRunLogs) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetIsMediaStorageService

`func (o *ConnectorRunLogs) GetIsMediaStorageService() bool`

GetIsMediaStorageService returns the IsMediaStorageService field if non-nil, zero value otherwise.

### GetIsMediaStorageServiceOk

`func (o *ConnectorRunLogs) GetIsMediaStorageServiceOk() (*bool, bool)`

GetIsMediaStorageServiceOk returns a tuple with the IsMediaStorageService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMediaStorageService

`func (o *ConnectorRunLogs) SetIsMediaStorageService(v bool)`

SetIsMediaStorageService sets IsMediaStorageService field to given value.


### GetStatus

`func (o *ConnectorRunLogs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorRunLogs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorRunLogs) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTestRun

`func (o *ConnectorRunLogs) GetTestRun() bool`

GetTestRun returns the TestRun field if non-nil, zero value otherwise.

### GetTestRunOk

`func (o *ConnectorRunLogs) GetTestRunOk() (*bool, bool)`

GetTestRunOk returns a tuple with the TestRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRun

`func (o *ConnectorRunLogs) SetTestRun(v bool)`

SetTestRun sets TestRun field to given value.


### GetDeployment

`func (o *ConnectorRunLogs) GetDeployment() string`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ConnectorRunLogs) GetDeploymentOk() (*string, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ConnectorRunLogs) SetDeployment(v string)`

SetDeployment sets Deployment field to given value.


### GetIpv4Address

`func (o *ConnectorRunLogs) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *ConnectorRunLogs) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *ConnectorRunLogs) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.


### GetAddedRecordsCount

`func (o *ConnectorRunLogs) GetAddedRecordsCount() int32`

GetAddedRecordsCount returns the AddedRecordsCount field if non-nil, zero value otherwise.

### GetAddedRecordsCountOk

`func (o *ConnectorRunLogs) GetAddedRecordsCountOk() (*int32, bool)`

GetAddedRecordsCountOk returns a tuple with the AddedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRecordsCount

`func (o *ConnectorRunLogs) SetAddedRecordsCount(v int32)`

SetAddedRecordsCount sets AddedRecordsCount field to given value.


### GetUpdatedRecordsCount

`func (o *ConnectorRunLogs) GetUpdatedRecordsCount() int32`

GetUpdatedRecordsCount returns the UpdatedRecordsCount field if non-nil, zero value otherwise.

### GetUpdatedRecordsCountOk

`func (o *ConnectorRunLogs) GetUpdatedRecordsCountOk() (*int32, bool)`

GetUpdatedRecordsCountOk returns a tuple with the UpdatedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedRecordsCount

`func (o *ConnectorRunLogs) SetUpdatedRecordsCount(v int32)`

SetUpdatedRecordsCount sets UpdatedRecordsCount field to given value.


### GetSkippedRecordsCount

`func (o *ConnectorRunLogs) GetSkippedRecordsCount() int32`

GetSkippedRecordsCount returns the SkippedRecordsCount field if non-nil, zero value otherwise.

### GetSkippedRecordsCountOk

`func (o *ConnectorRunLogs) GetSkippedRecordsCountOk() (*int32, bool)`

GetSkippedRecordsCountOk returns a tuple with the SkippedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedRecordsCount

`func (o *ConnectorRunLogs) SetSkippedRecordsCount(v int32)`

SetSkippedRecordsCount sets SkippedRecordsCount field to given value.


### GetErrorLogsCount

`func (o *ConnectorRunLogs) GetErrorLogsCount() int32`

GetErrorLogsCount returns the ErrorLogsCount field if non-nil, zero value otherwise.

### GetErrorLogsCountOk

`func (o *ConnectorRunLogs) GetErrorLogsCountOk() (*int32, bool)`

GetErrorLogsCountOk returns a tuple with the ErrorLogsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLogsCount

`func (o *ConnectorRunLogs) SetErrorLogsCount(v int32)`

SetErrorLogsCount sets ErrorLogsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


