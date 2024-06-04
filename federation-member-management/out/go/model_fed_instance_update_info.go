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
	"bytes"
	"fmt"
)

// checks if the FedInstanceUpdateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FedInstanceUpdateInfo{}

// FedInstanceUpdateInfo struct for FedInstanceUpdateInfo
type FedInstanceUpdateInfo struct {
	Enabled bool `json:"enabled"`
}

type _FedInstanceUpdateInfo FedInstanceUpdateInfo

// NewFedInstanceUpdateInfo instantiates a new FedInstanceUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFedInstanceUpdateInfo(enabled bool) *FedInstanceUpdateInfo {
	this := FedInstanceUpdateInfo{}
	this.Enabled = enabled
	return &this
}

// NewFedInstanceUpdateInfoWithDefaults instantiates a new FedInstanceUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFedInstanceUpdateInfoWithDefaults() *FedInstanceUpdateInfo {
	this := FedInstanceUpdateInfo{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *FedInstanceUpdateInfo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FedInstanceUpdateInfo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FedInstanceUpdateInfo) SetEnabled(v bool) {
	o.Enabled = v
}

func (o FedInstanceUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FedInstanceUpdateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *FedInstanceUpdateInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varFedInstanceUpdateInfo := _FedInstanceUpdateInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFedInstanceUpdateInfo)

	if err != nil {
		return err
	}

	*o = FedInstanceUpdateInfo(varFedInstanceUpdateInfo)

	return err
}

type NullableFedInstanceUpdateInfo struct {
	value *FedInstanceUpdateInfo
	isSet bool
}

func (v NullableFedInstanceUpdateInfo) Get() *FedInstanceUpdateInfo {
	return v.value
}

func (v *NullableFedInstanceUpdateInfo) Set(val *FedInstanceUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFedInstanceUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFedInstanceUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFedInstanceUpdateInfo(val *FedInstanceUpdateInfo) *NullableFedInstanceUpdateInfo {
	return &NullableFedInstanceUpdateInfo{value: val, isSet: true}
}

func (v NullableFedInstanceUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFedInstanceUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


