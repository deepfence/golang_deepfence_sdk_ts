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

// checks if the ModelUpdateScheduledTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUpdateScheduledTaskRequest{}

// ModelUpdateScheduledTaskRequest struct for ModelUpdateScheduledTaskRequest
type ModelUpdateScheduledTaskRequest struct {
	IsEnabled bool `json:"is_enabled"`
}

type _ModelUpdateScheduledTaskRequest ModelUpdateScheduledTaskRequest

// NewModelUpdateScheduledTaskRequest instantiates a new ModelUpdateScheduledTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUpdateScheduledTaskRequest(isEnabled bool) *ModelUpdateScheduledTaskRequest {
	this := ModelUpdateScheduledTaskRequest{}
	this.IsEnabled = isEnabled
	return &this
}

// NewModelUpdateScheduledTaskRequestWithDefaults instantiates a new ModelUpdateScheduledTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUpdateScheduledTaskRequestWithDefaults() *ModelUpdateScheduledTaskRequest {
	this := ModelUpdateScheduledTaskRequest{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *ModelUpdateScheduledTaskRequest) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *ModelUpdateScheduledTaskRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *ModelUpdateScheduledTaskRequest) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

func (o ModelUpdateScheduledTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUpdateScheduledTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_enabled"] = o.IsEnabled
	return toSerialize, nil
}

func (o *ModelUpdateScheduledTaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_enabled",
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

	varModelUpdateScheduledTaskRequest := _ModelUpdateScheduledTaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelUpdateScheduledTaskRequest)

	if err != nil {
		return err
	}

	*o = ModelUpdateScheduledTaskRequest(varModelUpdateScheduledTaskRequest)

	return err
}

type NullableModelUpdateScheduledTaskRequest struct {
	value *ModelUpdateScheduledTaskRequest
	isSet bool
}

func (v NullableModelUpdateScheduledTaskRequest) Get() *ModelUpdateScheduledTaskRequest {
	return v.value
}

func (v *NullableModelUpdateScheduledTaskRequest) Set(val *ModelUpdateScheduledTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUpdateScheduledTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUpdateScheduledTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUpdateScheduledTaskRequest(val *ModelUpdateScheduledTaskRequest) *NullableModelUpdateScheduledTaskRequest {
	return &NullableModelUpdateScheduledTaskRequest{value: val, isSet: true}
}

func (v NullableModelUpdateScheduledTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUpdateScheduledTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


