# FedInstanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FedId** | **string** |  | 
**FedName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**OwningFedadminId** | **string** |  | 
**ConnectionsUsed** | [**[]ConnectionID**](ConnectionID.md) |  | 

## Methods

### NewFedInstanceInfo

`func NewFedInstanceInfo(fedId string, fedName string, enabled bool, owningFedadminId string, connectionsUsed []ConnectionID, ) *FedInstanceInfo`

NewFedInstanceInfo instantiates a new FedInstanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFedInstanceInfoWithDefaults

`func NewFedInstanceInfoWithDefaults() *FedInstanceInfo`

NewFedInstanceInfoWithDefaults instantiates a new FedInstanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFedId

`func (o *FedInstanceInfo) GetFedId() string`

GetFedId returns the FedId field if non-nil, zero value otherwise.

### GetFedIdOk

`func (o *FedInstanceInfo) GetFedIdOk() (*string, bool)`

GetFedIdOk returns a tuple with the FedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFedId

`func (o *FedInstanceInfo) SetFedId(v string)`

SetFedId sets FedId field to given value.


### GetFedName

`func (o *FedInstanceInfo) GetFedName() string`

GetFedName returns the FedName field if non-nil, zero value otherwise.

### GetFedNameOk

`func (o *FedInstanceInfo) GetFedNameOk() (*string, bool)`

GetFedNameOk returns a tuple with the FedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFedName

`func (o *FedInstanceInfo) SetFedName(v string)`

SetFedName sets FedName field to given value.


### GetDescription

`func (o *FedInstanceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FedInstanceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FedInstanceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FedInstanceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FedInstanceInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FedInstanceInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FedInstanceInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOwningFedadminId

`func (o *FedInstanceInfo) GetOwningFedadminId() string`

GetOwningFedadminId returns the OwningFedadminId field if non-nil, zero value otherwise.

### GetOwningFedadminIdOk

`func (o *FedInstanceInfo) GetOwningFedadminIdOk() (*string, bool)`

GetOwningFedadminIdOk returns a tuple with the OwningFedadminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningFedadminId

`func (o *FedInstanceInfo) SetOwningFedadminId(v string)`

SetOwningFedadminId sets OwningFedadminId field to given value.


### GetConnectionsUsed

`func (o *FedInstanceInfo) GetConnectionsUsed() []ConnectionID`

GetConnectionsUsed returns the ConnectionsUsed field if non-nil, zero value otherwise.

### GetConnectionsUsedOk

`func (o *FedInstanceInfo) GetConnectionsUsedOk() (*[]ConnectionID, bool)`

GetConnectionsUsedOk returns a tuple with the ConnectionsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsUsed

`func (o *FedInstanceInfo) SetConnectionsUsed(v []ConnectionID)`

SetConnectionsUsed sets ConnectionsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


