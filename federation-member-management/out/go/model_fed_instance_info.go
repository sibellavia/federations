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

// checks if the FedInstanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FedInstanceInfo{}

// FedInstanceInfo struct for FedInstanceInfo
type FedInstanceInfo struct {
	FedId string `json:"fed_id"`
	FedName string `json:"fed_name"`
	Description *string `json:"description,omitempty"`
	Enabled bool `json:"enabled"`
	OwningFedadminId string `json:"owning_fedadmin_id"`
	ConnectionsUsed []ConnectionID `json:"connections_used"`
}

type _FedInstanceInfo FedInstanceInfo

// NewFedInstanceInfo instantiates a new FedInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFedInstanceInfo(fedId string, fedName string, enabled bool, owningFedadminId string, connectionsUsed []ConnectionID) *FedInstanceInfo {
	this := FedInstanceInfo{}
	this.FedId = fedId
	this.FedName = fedName
	this.Enabled = enabled
	this.OwningFedadminId = owningFedadminId
	this.ConnectionsUsed = connectionsUsed
	return &this
}

// NewFedInstanceInfoWithDefaults instantiates a new FedInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFedInstanceInfoWithDefaults() *FedInstanceInfo {
	this := FedInstanceInfo{}
	return &this
}

// GetFedId returns the FedId field value
func (o *FedInstanceInfo) GetFedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FedId
}

// GetFedIdOk returns a tuple with the FedId field value
// and a boolean to check if the value has been set.
func (o *FedInstanceInfo) GetFedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FedId, true
}

// SetFedId sets field value
func (o *FedInstanceInfo) SetFedId(v string) {
	o.FedId = v
}

// GetFedName returns the FedName field value
func (o *FedInstanceInfo) GetFedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FedName
}

// GetFedNameOk returns a tuple with the FedName field value
// and a boolean to check if the value has been set.
func (o *FedInstanceInfo) GetFedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FedName, true
}

// SetFedName sets field value
func (o *FedInstanceInfo) SetFedName(v string) {
	o.FedName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FedInstanceInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FedInstanceInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FedInstanceInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FedInstanceInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FedInstanceInfo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FedInstanceInfo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FedInstanceInfo) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOwningFedadminId returns the OwningFedadminId field value
func (o *FedInstanceInfo) GetOwningFedadminId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwningFedadminId
}

// GetOwningFedadminIdOk returns a tuple with the OwningFedadminId field value
// and a boolean to check if the value has been set.
func (o *FedInstanceInfo) GetOwningFedadminIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwningFedadminId, true
}

// SetOwningFedadminId sets field value
func (o *FedInstanceInfo) SetOwningFedadminId(v string) {
	o.OwningFedadminId = v
}

// GetConnectionsUsed returns the ConnectionsUsed field value
func (o *FedInstanceInfo) GetConnectionsUsed() []ConnectionID {
	if o == nil {
		var ret []ConnectionID
		return ret
	}

	return o.ConnectionsUsed
}

// GetConnectionsUsedOk returns a tuple with the ConnectionsUsed field value
// and a boolean to check if the value has been set.
func (o *FedInstanceInfo) GetConnectionsUsedOk() ([]ConnectionID, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionsUsed, true
}

// SetConnectionsUsed sets field value
func (o *FedInstanceInfo) SetConnectionsUsed(v []ConnectionID) {
	o.ConnectionsUsed = v
}

func (o FedInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FedInstanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fed_id"] = o.FedId
	toSerialize["fed_name"] = o.FedName
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["owning_fedadmin_id"] = o.OwningFedadminId
	toSerialize["connections_used"] = o.ConnectionsUsed
	return toSerialize, nil
}

func (o *FedInstanceInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fed_id",
		"fed_name",
		"enabled",
		"owning_fedadmin_id",
		"connections_used",
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

	varFedInstanceInfo := _FedInstanceInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFedInstanceInfo)

	if err != nil {
		return err
	}

	*o = FedInstanceInfo(varFedInstanceInfo)

	return err
}

type NullableFedInstanceInfo struct {
	value *FedInstanceInfo
	isSet bool
}

func (v NullableFedInstanceInfo) Get() *FedInstanceInfo {
	return v.value
}

func (v *NullableFedInstanceInfo) Set(val *FedInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFedInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFedInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFedInstanceInfo(val *FedInstanceInfo) *NullableFedInstanceInfo {
	return &NullableFedInstanceInfo{value: val, isSet: true}
}

func (v NullableFedInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFedInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


