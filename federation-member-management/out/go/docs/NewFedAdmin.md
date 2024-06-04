# NewFedAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | [default to false]

## Methods

### NewNewFedAdmin

`func NewNewFedAdmin(name string, enabled bool, ) *NewFedAdmin`

NewNewFedAdmin instantiates a new NewFedAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewFedAdminWithDefaults

`func NewNewFedAdminWithDefaults() *NewFedAdmin`

NewNewFedAdminWithDefaults instantiates a new NewFedAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewFedAdmin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewFedAdmin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewFedAdmin) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *NewFedAdmin) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewFedAdmin) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewFedAdmin) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewFedAdmin) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *NewFedAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewFedAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewFedAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewFedAdmin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NewFedAdmin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewFedAdmin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewFedAdmin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


