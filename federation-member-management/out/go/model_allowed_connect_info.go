/*
FHSOperator API

This is the FHSOperator API for the IEEE 2302-2021 Standard for Intercloud Interoperability and Federation (https://standards.ieee.org/ieee/2302/7056).  This standard is based on the NIST Cloud Federation Reference Architecture (CFRA), SP 500-332, (https://doi.org/10.6028/NIST.SP.500-332).

API version: 0.1
Contact: lee@keyvoms.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AllowedConnectInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowedConnectInfo{}

// AllowedConnectInfo struct for AllowedConnectInfo
type AllowedConnectInfo struct {
	FHS_URL string `json:"FHS_URL"`
	ConnectTtl *time.Time `json:"connect_ttl,omitempty"`
	Description *string `json:"description,omitempty"`
}

type _AllowedConnectInfo AllowedConnectInfo

// NewAllowedConnectInfo instantiates a new AllowedConnectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedConnectInfo(fHSURL string) *AllowedConnectInfo {
	this := AllowedConnectInfo{}
	this.FHS_URL = fHSURL
	return &this
}

// NewAllowedConnectInfoWithDefaults instantiates a new AllowedConnectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedConnectInfoWithDefaults() *AllowedConnectInfo {
	this := AllowedConnectInfo{}
	return &this
}

// GetFHS_URL returns the FHS_URL field value
func (o *AllowedConnectInfo) GetFHS_URL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FHS_URL
}

// GetFHS_URLOk returns a tuple with the FHS_URL field value
// and a boolean to check if the value has been set.
func (o *AllowedConnectInfo) GetFHS_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FHS_URL, true
}

// SetFHS_URL sets field value
func (o *AllowedConnectInfo) SetFHS_URL(v string) {
	o.FHS_URL = v
}

// GetConnectTtl returns the ConnectTtl field value if set, zero value otherwise.
func (o *AllowedConnectInfo) GetConnectTtl() time.Time {
	if o == nil || IsNil(o.ConnectTtl) {
		var ret time.Time
		return ret
	}
	return *o.ConnectTtl
}

// GetConnectTtlOk returns a tuple with the ConnectTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedConnectInfo) GetConnectTtlOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConnectTtl) {
		return nil, false
	}
	return o.ConnectTtl, true
}

// HasConnectTtl returns a boolean if a field has been set.
func (o *AllowedConnectInfo) HasConnectTtl() bool {
	if o != nil && !IsNil(o.ConnectTtl) {
		return true
	}

	return false
}

// SetConnectTtl gets a reference to the given time.Time and assigns it to the ConnectTtl field.
func (o *AllowedConnectInfo) SetConnectTtl(v time.Time) {
	o.ConnectTtl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllowedConnectInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedConnectInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllowedConnectInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllowedConnectInfo) SetDescription(v string) {
	o.Description = &v
}

func (o AllowedConnectInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowedConnectInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FHS_URL"] = o.FHS_URL
	if !IsNil(o.ConnectTtl) {
		toSerialize["connect_ttl"] = o.ConnectTtl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *AllowedConnectInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"FHS_URL",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAllowedConnectInfo := _AllowedConnectInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAllowedConnectInfo)

	if err != nil {
		return err
	}

	*o = AllowedConnectInfo(varAllowedConnectInfo)

	return err
}

type NullableAllowedConnectInfo struct {
	value *AllowedConnectInfo
	isSet bool
}

func (v NullableAllowedConnectInfo) Get() *AllowedConnectInfo {
	return v.value
}

func (v *NullableAllowedConnectInfo) Set(val *AllowedConnectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedConnectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedConnectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedConnectInfo(val *AllowedConnectInfo) *NullableAllowedConnectInfo {
	return &NullableAllowedConnectInfo{value: val, isSet: true}
}

func (v NullableAllowedConnectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedConnectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

