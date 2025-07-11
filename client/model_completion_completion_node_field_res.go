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

// checks if the CompletionCompletionNodeFieldRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletionCompletionNodeFieldRes{}

// CompletionCompletionNodeFieldRes struct for CompletionCompletionNodeFieldRes
type CompletionCompletionNodeFieldRes struct {
	PossibleValues []string `json:"possible_values"`
}

type _CompletionCompletionNodeFieldRes CompletionCompletionNodeFieldRes

// NewCompletionCompletionNodeFieldRes instantiates a new CompletionCompletionNodeFieldRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletionCompletionNodeFieldRes(possibleValues []string) *CompletionCompletionNodeFieldRes {
	this := CompletionCompletionNodeFieldRes{}
	this.PossibleValues = possibleValues
	return &this
}

// NewCompletionCompletionNodeFieldResWithDefaults instantiates a new CompletionCompletionNodeFieldRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletionCompletionNodeFieldResWithDefaults() *CompletionCompletionNodeFieldRes {
	this := CompletionCompletionNodeFieldRes{}
	return &this
}

// GetPossibleValues returns the PossibleValues field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CompletionCompletionNodeFieldRes) GetPossibleValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PossibleValues
}

// GetPossibleValuesOk returns a tuple with the PossibleValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompletionCompletionNodeFieldRes) GetPossibleValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.PossibleValues) {
		return nil, false
	}
	return o.PossibleValues, true
}

// SetPossibleValues sets field value
func (o *CompletionCompletionNodeFieldRes) SetPossibleValues(v []string) {
	o.PossibleValues = v
}

func (o CompletionCompletionNodeFieldRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletionCompletionNodeFieldRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PossibleValues != nil {
		toSerialize["possible_values"] = o.PossibleValues
	}
	return toSerialize, nil
}

func (o *CompletionCompletionNodeFieldRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"possible_values",
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

	varCompletionCompletionNodeFieldRes := _CompletionCompletionNodeFieldRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompletionCompletionNodeFieldRes)

	if err != nil {
		return err
	}

	*o = CompletionCompletionNodeFieldRes(varCompletionCompletionNodeFieldRes)

	return err
}

type NullableCompletionCompletionNodeFieldRes struct {
	value *CompletionCompletionNodeFieldRes
	isSet bool
}

func (v NullableCompletionCompletionNodeFieldRes) Get() *CompletionCompletionNodeFieldRes {
	return v.value
}

func (v *NullableCompletionCompletionNodeFieldRes) Set(val *CompletionCompletionNodeFieldRes) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionCompletionNodeFieldRes) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionCompletionNodeFieldRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionCompletionNodeFieldRes(val *CompletionCompletionNodeFieldRes) *NullableCompletionCompletionNodeFieldRes {
	return &NullableCompletionCompletionNodeFieldRes{value: val, isSet: true}
}

func (v NullableCompletionCompletionNodeFieldRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionCompletionNodeFieldRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


