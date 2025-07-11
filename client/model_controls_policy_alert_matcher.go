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

// checks if the ControlsPolicyAlertMatcher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsPolicyAlertMatcher{}

// ControlsPolicyAlertMatcher struct for ControlsPolicyAlertMatcher
type ControlsPolicyAlertMatcher struct {
	FieldsMatcher map[string][]string `json:"fields_matcher"`
}

type _ControlsPolicyAlertMatcher ControlsPolicyAlertMatcher

// NewControlsPolicyAlertMatcher instantiates a new ControlsPolicyAlertMatcher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsPolicyAlertMatcher(fieldsMatcher map[string][]string) *ControlsPolicyAlertMatcher {
	this := ControlsPolicyAlertMatcher{}
	this.FieldsMatcher = fieldsMatcher
	return &this
}

// NewControlsPolicyAlertMatcherWithDefaults instantiates a new ControlsPolicyAlertMatcher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsPolicyAlertMatcherWithDefaults() *ControlsPolicyAlertMatcher {
	this := ControlsPolicyAlertMatcher{}
	return &this
}

// GetFieldsMatcher returns the FieldsMatcher field value
// If the value is explicit nil, the zero value for map[string][]string will be returned
func (o *ControlsPolicyAlertMatcher) GetFieldsMatcher() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.FieldsMatcher
}

// GetFieldsMatcherOk returns a tuple with the FieldsMatcher field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlsPolicyAlertMatcher) GetFieldsMatcherOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.FieldsMatcher) {
		return nil, false
	}
	return &o.FieldsMatcher, true
}

// SetFieldsMatcher sets field value
func (o *ControlsPolicyAlertMatcher) SetFieldsMatcher(v map[string][]string) {
	o.FieldsMatcher = v
}

func (o ControlsPolicyAlertMatcher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsPolicyAlertMatcher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldsMatcher != nil {
		toSerialize["fields_matcher"] = o.FieldsMatcher
	}
	return toSerialize, nil
}

func (o *ControlsPolicyAlertMatcher) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields_matcher",
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

	varControlsPolicyAlertMatcher := _ControlsPolicyAlertMatcher{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varControlsPolicyAlertMatcher)

	if err != nil {
		return err
	}

	*o = ControlsPolicyAlertMatcher(varControlsPolicyAlertMatcher)

	return err
}

type NullableControlsPolicyAlertMatcher struct {
	value *ControlsPolicyAlertMatcher
	isSet bool
}

func (v NullableControlsPolicyAlertMatcher) Get() *ControlsPolicyAlertMatcher {
	return v.value
}

func (v *NullableControlsPolicyAlertMatcher) Set(val *ControlsPolicyAlertMatcher) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsPolicyAlertMatcher) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsPolicyAlertMatcher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsPolicyAlertMatcher(val *ControlsPolicyAlertMatcher) *NullableControlsPolicyAlertMatcher {
	return &NullableControlsPolicyAlertMatcher{value: val, isSet: true}
}

func (v NullableControlsPolicyAlertMatcher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsPolicyAlertMatcher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


