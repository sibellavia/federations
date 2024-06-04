# ConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** |  | 
**FHS_URL** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**TimeEstablished** | Pointer to **time.Time** |  | [optional] 
**TimeTerminated** | Pointer to **time.Time** |  | [optional] 
**KnownFederations** | [**[]FederationID**](FederationID.md) |  | 

## Methods

### NewConnectionInfo

`func NewConnectionInfo(connectionId string, fHSURL string, status string, knownFederations []FederationID, ) *ConnectionInfo`

NewConnectionInfo instantiates a new ConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInfoWithDefaults

`func NewConnectionInfoWithDefaults() *ConnectionInfo`

NewConnectionInfoWithDefaults instantiates a new ConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionInfo) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionInfo) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionInfo) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFHS_URL

`func (o *ConnectionInfo) GetFHS_URL() string`

GetFHS_URL returns the FHS_URL field if non-nil, zero value otherwise.

### GetFHS_URLOk

`func (o *ConnectionInfo) GetFHS_URLOk() (*string, bool)`

GetFHS_URLOk returns a tuple with the FHS_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFHS_URL

`func (o *ConnectionInfo) SetFHS_URL(v string)`

SetFHS_URL sets FHS_URL field to given value.


### GetDescription

`func (o *ConnectionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeEstablished

`func (o *ConnectionInfo) GetTimeEstablished() time.Time`

GetTimeEstablished returns the TimeEstablished field if non-nil, zero value otherwise.

### GetTimeEstablishedOk

`func (o *ConnectionInfo) GetTimeEstablishedOk() (*time.Time, bool)`

GetTimeEstablishedOk returns a tuple with the TimeEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEstablished

`func (o *ConnectionInfo) SetTimeEstablished(v time.Time)`

SetTimeEstablished sets TimeEstablished field to given value.

### HasTimeEstablished

`func (o *ConnectionInfo) HasTimeEstablished() bool`

HasTimeEstablished returns a boolean if a field has been set.

### GetTimeTerminated

`func (o *ConnectionInfo) GetTimeTerminated() time.Time`

GetTimeTerminated returns the TimeTerminated field if non-nil, zero value otherwise.

### GetTimeTerminatedOk

`func (o *ConnectionInfo) GetTimeTerminatedOk() (*time.Time, bool)`

GetTimeTerminatedOk returns a tuple with the TimeTerminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTerminated

`func (o *ConnectionInfo) SetTimeTerminated(v time.Time)`

SetTimeTerminated sets TimeTerminated field to given value.

### HasTimeTerminated

`func (o *ConnectionInfo) HasTimeTerminated() bool`

HasTimeTerminated returns a boolean if a field has been set.

### GetKnownFederations

`func (o *ConnectionInfo) GetKnownFederations() []FederationID`

GetKnownFederations returns the KnownFederations field if non-nil, zero value otherwise.

### GetKnownFederationsOk

`func (o *ConnectionInfo) GetKnownFederationsOk() (*[]FederationID, bool)`

GetKnownFederationsOk returns a tuple with the KnownFederations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownFederations

`func (o *ConnectionInfo) SetKnownFederations(v []FederationID)`

SetKnownFederations sets KnownFederations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


