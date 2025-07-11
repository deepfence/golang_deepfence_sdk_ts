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

// checks if the ReportersSevCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersSevCounts{}

// ReportersSevCounts struct for ReportersSevCounts
type ReportersSevCounts struct {
	Counts map[string]int32 `json:"counts"`
}

type _ReportersSevCounts ReportersSevCounts

// NewReportersSevCounts instantiates a new ReportersSevCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersSevCounts(counts map[string]int32) *ReportersSevCounts {
	this := ReportersSevCounts{}
	this.Counts = counts
	return &this
}

// NewReportersSevCountsWithDefaults instantiates a new ReportersSevCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersSevCountsWithDefaults() *ReportersSevCounts {
	this := ReportersSevCounts{}
	return &this
}

// GetCounts returns the Counts field value
// If the value is explicit nil, the zero value for map[string]int32 will be returned
func (o *ReportersSevCounts) GetCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersSevCounts) GetCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *ReportersSevCounts) SetCounts(v map[string]int32) {
	o.Counts = v
}

func (o ReportersSevCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersSevCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

func (o *ReportersSevCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"counts",
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

	varReportersSevCounts := _ReportersSevCounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportersSevCounts)

	if err != nil {
		return err
	}

	*o = ReportersSevCounts(varReportersSevCounts)

	return err
}

type NullableReportersSevCounts struct {
	value *ReportersSevCounts
	isSet bool
}

func (v NullableReportersSevCounts) Get() *ReportersSevCounts {
	return v.value
}

func (v *NullableReportersSevCounts) Set(val *ReportersSevCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersSevCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersSevCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersSevCounts(val *ReportersSevCounts) *NullableReportersSevCounts {
	return &NullableReportersSevCounts{value: val, isSet: true}
}

func (v NullableReportersSevCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersSevCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


