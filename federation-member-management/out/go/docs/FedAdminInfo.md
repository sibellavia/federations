# FedAdminInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | **string** |  | 
**MemberName** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**FedsOwned** | [**[]FederationID**](FederationID.md) |  | 

## Methods

### NewFedAdminInfo

`func NewFedAdminInfo(memberId string, memberName string, enabled bool, fedsOwned []FederationID, ) *FedAdminInfo`

NewFedAdminInfo instantiates a new FedAdminInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFedAdminInfoWithDefaults

`func NewFedAdminInfoWithDefaults() *FedAdminInfo`

NewFedAdminInfoWithDefaults instantiates a new FedAdminInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *FedAdminInfo) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *FedAdminInfo) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *FedAdminInfo) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetMemberName

`func (o *FedAdminInfo) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *FedAdminInfo) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *FedAdminInfo) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.


### GetEmail

`func (o *FedAdminInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FedAdminInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FedAdminInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FedAdminInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *FedAdminInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FedAdminInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FedAdminInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FedAdminInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FedAdminInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FedAdminInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FedAdminInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFedsOwned

`func (o *FedAdminInfo) GetFedsOwned() []FederationID`

GetFedsOwned returns the FedsOwned field if non-nil, zero value otherwise.

### GetFedsOwnedOk

`func (o *FedAdminInfo) GetFedsOwnedOk() (*[]FederationID, bool)`

GetFedsOwnedOk returns a tuple with the FedsOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFedsOwned

`func (o *FedAdminInfo) SetFedsOwned(v []FederationID)`

SetFedsOwned sets FedsOwned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


