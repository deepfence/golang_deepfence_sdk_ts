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

// checks if the ModelEnableTracerReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelEnableTracerReq{}

// ModelEnableTracerReq struct for ModelEnableTracerReq
type ModelEnableTracerReq struct {
	AgentIds []ModelAgentID `json:"agent_ids"`
}

type _ModelEnableTracerReq ModelEnableTracerReq

// NewModelEnableTracerReq instantiates a new ModelEnableTracerReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelEnableTracerReq(agentIds []ModelAgentID) *ModelEnableTracerReq {
	this := ModelEnableTracerReq{}
	this.AgentIds = agentIds
	return &this
}

// NewModelEnableTracerReqWithDefaults instantiates a new ModelEnableTracerReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelEnableTracerReqWithDefaults() *ModelEnableTracerReq {
	this := ModelEnableTracerReq{}
	return &this
}

// GetAgentIds returns the AgentIds field value
// If the value is explicit nil, the zero value for []ModelAgentID will be returned
func (o *ModelEnableTracerReq) GetAgentIds() []ModelAgentID {
	if o == nil {
		var ret []ModelAgentID
		return ret
	}

	return o.AgentIds
}

// GetAgentIdsOk returns a tuple with the AgentIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelEnableTracerReq) GetAgentIdsOk() ([]ModelAgentID, bool) {
	if o == nil || IsNil(o.AgentIds) {
		return nil, false
	}
	return o.AgentIds, true
}

// SetAgentIds sets field value
func (o *ModelEnableTracerReq) SetAgentIds(v []ModelAgentID) {
	o.AgentIds = v
}

func (o ModelEnableTracerReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelEnableTracerReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentIds != nil {
		toSerialize["agent_ids"] = o.AgentIds
	}
	return toSerialize, nil
}

func (o *ModelEnableTracerReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_ids",
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

	varModelEnableTracerReq := _ModelEnableTracerReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelEnableTracerReq)

	if err != nil {
		return err
	}

	*o = ModelEnableTracerReq(varModelEnableTracerReq)

	return err
}

type NullableModelEnableTracerReq struct {
	value *ModelEnableTracerReq
	isSet bool
}

func (v NullableModelEnableTracerReq) Get() *ModelEnableTracerReq {
	return v.value
}

func (v *NullableModelEnableTracerReq) Set(val *ModelEnableTracerReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelEnableTracerReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelEnableTracerReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelEnableTracerReq(val *ModelEnableTracerReq) *NullableModelEnableTracerReq {
	return &NullableModelEnableTracerReq{value: val, isSet: true}
}

func (v NullableModelEnableTracerReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelEnableTracerReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


