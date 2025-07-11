/*
Deepfence ThreatStryker

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.6
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SinglesignonSSOConfigurationInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SinglesignonSSOConfigurationInstruction{}

// SinglesignonSSOConfigurationInstruction struct for SinglesignonSSOConfigurationInstruction
type SinglesignonSSOConfigurationInstruction struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type _SinglesignonSSOConfigurationInstruction SinglesignonSSOConfigurationInstruction

// NewSinglesignonSSOConfigurationInstruction instantiates a new SinglesignonSSOConfigurationInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinglesignonSSOConfigurationInstruction(key string, value string) *SinglesignonSSOConfigurationInstruction {
	this := SinglesignonSSOConfigurationInstruction{}
	this.Key = key
	this.Value = value
	return &this
}

// NewSinglesignonSSOConfigurationInstructionWithDefaults instantiates a new SinglesignonSSOConfigurationInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinglesignonSSOConfigurationInstructionWithDefaults() *SinglesignonSSOConfigurationInstruction {
	this := SinglesignonSSOConfigurationInstruction{}
	return &this
}

// GetKey returns the Key field value
func (o *SinglesignonSSOConfigurationInstruction) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOConfigurationInstruction) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SinglesignonSSOConfigurationInstruction) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *SinglesignonSSOConfigurationInstruction) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOConfigurationInstruction) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SinglesignonSSOConfigurationInstruction) SetValue(v string) {
	o.Value = v
}

func (o SinglesignonSSOConfigurationInstruction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SinglesignonSSOConfigurationInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *SinglesignonSSOConfigurationInstruction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varSinglesignonSSOConfigurationInstruction := _SinglesignonSSOConfigurationInstruction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSinglesignonSSOConfigurationInstruction)

	if err != nil {
		return err
	}

	*o = SinglesignonSSOConfigurationInstruction(varSinglesignonSSOConfigurationInstruction)

	return err
}

type NullableSinglesignonSSOConfigurationInstruction struct {
	value *SinglesignonSSOConfigurationInstruction
	isSet bool
}

func (v NullableSinglesignonSSOConfigurationInstruction) Get() *SinglesignonSSOConfigurationInstruction {
	return v.value
}

func (v *NullableSinglesignonSSOConfigurationInstruction) Set(val *SinglesignonSSOConfigurationInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableSinglesignonSSOConfigurationInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableSinglesignonSSOConfigurationInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinglesignonSSOConfigurationInstruction(val *SinglesignonSSOConfigurationInstruction) *NullableSinglesignonSSOConfigurationInstruction {
	return &NullableSinglesignonSSOConfigurationInstruction{value: val, isSet: true}
}

func (v NullableSinglesignonSSOConfigurationInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinglesignonSSOConfigurationInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


