# ConnectorRunLogsPortionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** |  | 
**Deployment** | **string** |  | 
**StartTime** | **int32** |  | 
**EndTime** | **int32** |  | 
**Status** | **string** |  | 
**Ipv4Address** | **string** |  | 
**Objects** | [**ConnectorRunLogsPortionDetailsObjects**](ConnectorRunLogsPortionDetailsObjects.md) |  | 
**Logs** | [**ConnectorRunLogsPortionDetailsLogs**](ConnectorRunLogsPortionDetailsLogs.md) |  | 
**FileAttachment** | [**NullableConnectorRunLogsPortionDetailsFileAttachment**](ConnectorRunLogsPortionDetailsFileAttachment.md) |  | 
**TestRun** | **bool** |  | 
**TestRunPayload** | **[]map[string]interface{}** |  | 

## Methods

### NewConnectorRunLogsPortionDetails

`func NewConnectorRunLogsPortionDetails(connector string, deployment string, startTime int32, endTime int32, status string, ipv4Address string, objects ConnectorRunLogsPortionDetailsObjects, logs ConnectorRunLogsPortionDetailsLogs, fileAttachment NullableConnectorRunLogsPortionDetailsFileAttachment, testRun bool, testRunPayload []map[string]interface{}, ) *ConnectorRunLogsPortionDetails`

NewConnectorRunLogsPortionDetails instantiates a new ConnectorRunLogsPortionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorRunLogsPortionDetailsWithDefaults

`func NewConnectorRunLogsPortionDetailsWithDefaults() *ConnectorRunLogsPortionDetails`

NewConnectorRunLogsPortionDetailsWithDefaults instantiates a new ConnectorRunLogsPortionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ConnectorRunLogsPortionDetails) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectorRunLogsPortionDetails) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectorRunLogsPortionDetails) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetDeployment

`func (o *ConnectorRunLogsPortionDetails) GetDeployment() string`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ConnectorRunLogsPortionDetails) GetDeploymentOk() (*string, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ConnectorRunLogsPortionDetails) SetDeployment(v string)`

SetDeployment sets Deployment field to given value.


### GetStartTime

`func (o *ConnectorRunLogsPortionDetails) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ConnectorRunLogsPortionDetails) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ConnectorRunLogsPortionDetails) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ConnectorRunLogsPortionDetails) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ConnectorRunLogsPortionDetails) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ConnectorRunLogsPortionDetails) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetStatus

`func (o *ConnectorRunLogsPortionDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorRunLogsPortionDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorRunLogsPortionDetails) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIpv4Address

`func (o *ConnectorRunLogsPortionDetails) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *ConnectorRunLogsPortionDetails) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *ConnectorRunLogsPortionDetails) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.


### GetObjects

`func (o *ConnectorRunLogsPortionDetails) GetObjects() ConnectorRunLogsPortionDetailsObjects`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ConnectorRunLogsPortionDetails) GetObjectsOk() (*ConnectorRunLogsPortionDetailsObjects, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ConnectorRunLogsPortionDetails) SetObjects(v ConnectorRunLogsPortionDetailsObjects)`

SetObjects sets Objects field to given value.


### GetLogs

`func (o *ConnectorRunLogsPortionDetails) GetLogs() ConnectorRunLogsPortionDetailsLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ConnectorRunLogsPortionDetails) GetLogsOk() (*ConnectorRunLogsPortionDetailsLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ConnectorRunLogsPortionDetails) SetLogs(v ConnectorRunLogsPortionDetailsLogs)`

SetLogs sets Logs field to given value.


### GetFileAttachment

`func (o *ConnectorRunLogsPortionDetails) GetFileAttachment() ConnectorRunLogsPortionDetailsFileAttachment`

GetFileAttachment returns the FileAttachment field if non-nil, zero value otherwise.

### GetFileAttachmentOk

`func (o *ConnectorRunLogsPortionDetails) GetFileAttachmentOk() (*ConnectorRunLogsPortionDetailsFileAttachment, bool)`

GetFileAttachmentOk returns a tuple with the FileAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAttachment

`func (o *ConnectorRunLogsPortionDetails) SetFileAttachment(v ConnectorRunLogsPortionDetailsFileAttachment)`

SetFileAttachment sets FileAttachment field to given value.


### SetFileAttachmentNil

`func (o *ConnectorRunLogsPortionDetails) SetFileAttachmentNil(b bool)`

 SetFileAttachmentNil sets the value for FileAttachment to be an explicit nil

### UnsetFileAttachment
`func (o *ConnectorRunLogsPortionDetails) UnsetFileAttachment()`

UnsetFileAttachment ensures that no value is present for FileAttachment, not even an explicit nil
### GetTestRun

`func (o *ConnectorRunLogsPortionDetails) GetTestRun() bool`

GetTestRun returns the TestRun field if non-nil, zero value otherwise.

### GetTestRunOk

`func (o *ConnectorRunLogsPortionDetails) GetTestRunOk() (*bool, bool)`

GetTestRunOk returns a tuple with the TestRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRun

`func (o *ConnectorRunLogsPortionDetails) SetTestRun(v bool)`

SetTestRun sets TestRun field to given value.


### GetTestRunPayload

`func (o *ConnectorRunLogsPortionDetails) GetTestRunPayload() []map[string]interface{}`

GetTestRunPayload returns the TestRunPayload field if non-nil, zero value otherwise.

### GetTestRunPayloadOk

`func (o *ConnectorRunLogsPortionDetails) GetTestRunPayloadOk() (*[]map[string]interface{}, bool)`

GetTestRunPayloadOk returns a tuple with the TestRunPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunPayload

`func (o *ConnectorRunLogsPortionDetails) SetTestRunPayload(v []map[string]interface{})`

SetTestRunPayload sets TestRunPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


