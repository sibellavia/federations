# ConnectReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FHS_URL** | **string** |  | 
**ConnectTtl** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectReq

`func NewConnectReq(fHSURL string, ) *ConnectReq`

NewConnectReq instantiates a new ConnectReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectReqWithDefaults

`func NewConnectReqWithDefaults() *ConnectReq`

NewConnectReqWithDefaults instantiates a new ConnectReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFHS_URL

`func (o *ConnectReq) GetFHS_URL() string`

GetFHS_URL returns the FHS_URL field if non-nil, zero value otherwise.

### GetFHS_URLOk

`func (o *ConnectReq) GetFHS_URLOk() (*string, bool)`

GetFHS_URLOk returns a tuple with the FHS_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFHS_URL

`func (o *ConnectReq) SetFHS_URL(v string)`

SetFHS_URL sets FHS_URL field to given value.


### GetConnectTtl

`func (o *ConnectReq) GetConnectTtl() time.Time`

GetConnectTtl returns the ConnectTtl field if non-nil, zero value otherwise.

### GetConnectTtlOk

`func (o *ConnectReq) GetConnectTtlOk() (*time.Time, bool)`

GetConnectTtlOk returns a tuple with the ConnectTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTtl

`func (o *ConnectReq) SetConnectTtl(v time.Time)`

SetConnectTtl sets ConnectTtl field to given value.

### HasConnectTtl

`func (o *ConnectReq) HasConnectTtl() bool`

HasConnectTtl returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


