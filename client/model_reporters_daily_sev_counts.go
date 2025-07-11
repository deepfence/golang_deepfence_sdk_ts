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

// checks if the ReportersDailySevCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersDailySevCounts{}

// ReportersDailySevCounts struct for ReportersDailySevCounts
type ReportersDailySevCounts struct {
	DaysToSevCounts map[string]ReportersSevCounts `json:"days_to_sev_counts"`
}

type _ReportersDailySevCounts ReportersDailySevCounts

// NewReportersDailySevCounts instantiates a new ReportersDailySevCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersDailySevCounts(daysToSevCounts map[string]ReportersSevCounts) *ReportersDailySevCounts {
	this := ReportersDailySevCounts{}
	this.DaysToSevCounts = daysToSevCounts
	return &this
}

// NewReportersDailySevCountsWithDefaults instantiates a new ReportersDailySevCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersDailySevCountsWithDefaults() *ReportersDailySevCounts {
	this := ReportersDailySevCounts{}
	return &this
}

// GetDaysToSevCounts returns the DaysToSevCounts field value
// If the value is explicit nil, the zero value for map[string]ReportersSevCounts will be returned
func (o *ReportersDailySevCounts) GetDaysToSevCounts() map[string]ReportersSevCounts {
	if o == nil {
		var ret map[string]ReportersSevCounts
		return ret
	}

	return o.DaysToSevCounts
}

// GetDaysToSevCountsOk returns a tuple with the DaysToSevCounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersDailySevCounts) GetDaysToSevCountsOk() (*map[string]ReportersSevCounts, bool) {
	if o == nil || IsNil(o.DaysToSevCounts) {
		return nil, false
	}
	return &o.DaysToSevCounts, true
}

// SetDaysToSevCounts sets field value
func (o *ReportersDailySevCounts) SetDaysToSevCounts(v map[string]ReportersSevCounts) {
	o.DaysToSevCounts = v
}

func (o ReportersDailySevCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersDailySevCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysToSevCounts != nil {
		toSerialize["days_to_sev_counts"] = o.DaysToSevCounts
	}
	return toSerialize, nil
}

func (o *ReportersDailySevCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"days_to_sev_counts",
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

	varReportersDailySevCounts := _ReportersDailySevCounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportersDailySevCounts)

	if err != nil {
		return err
	}

	*o = ReportersDailySevCounts(varReportersDailySevCounts)

	return err
}

type NullableReportersDailySevCounts struct {
	value *ReportersDailySevCounts
	isSet bool
}

func (v NullableReportersDailySevCounts) Get() *ReportersDailySevCounts {
	return v.value
}

func (v *NullableReportersDailySevCounts) Set(val *ReportersDailySevCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersDailySevCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersDailySevCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersDailySevCounts(val *ReportersDailySevCounts) *NullableReportersDailySevCounts {
	return &NullableReportersDailySevCounts{value: val, isSet: true}
}

func (v NullableReportersDailySevCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersDailySevCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


