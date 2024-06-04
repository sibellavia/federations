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

// checks if the FedAdminID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FedAdminID{}

// FedAdminID struct for FedAdminID
type FedAdminID struct {
	MemberId string `json:"member_id"`
}

type _FedAdminID FedAdminID

// NewFedAdminID instantiates a new FedAdminID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFedAdminID(memberId string) *FedAdminID {
	this := FedAdminID{}
	this.MemberId = memberId
	return &this
}

// NewFedAdminIDWithDefaults instantiates a new FedAdminID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFedAdminIDWithDefaults() *FedAdminID {
	this := FedAdminID{}
	return &this
}

// GetMemberId returns the MemberId field value
func (o *FedAdminID) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *FedAdminID) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *FedAdminID) SetMemberId(v string) {
	o.MemberId = v
}

func (o FedAdminID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FedAdminID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["member_id"] = o.MemberId
	return toSerialize, nil
}

func (o *FedAdminID) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"member_id",
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

	varFedAdminID := _FedAdminID{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFedAdminID)

	if err != nil {
		return err
	}

	*o = FedAdminID(varFedAdminID)

	return err
}

type NullableFedAdminID struct {
	value *FedAdminID
	isSet bool
}

func (v NullableFedAdminID) Get() *FedAdminID {
	return v.value
}

func (v *NullableFedAdminID) Set(val *FedAdminID) {
	v.value = val
	v.isSet = true
}

func (v NullableFedAdminID) IsSet() bool {
	return v.isSet
}

func (v *NullableFedAdminID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFedAdminID(val *FedAdminID) *NullableFedAdminID {
	return &NullableFedAdminID{value: val, isSet: true}
}

func (v NullableFedAdminID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFedAdminID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

