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

// checks if the ModelAgentPluginDisable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAgentPluginDisable{}

// ModelAgentPluginDisable struct for ModelAgentPluginDisable
type ModelAgentPluginDisable struct {
	NodeId string `json:"node_id"`
	PluginName string `json:"plugin_name"`
}

type _ModelAgentPluginDisable ModelAgentPluginDisable

// NewModelAgentPluginDisable instantiates a new ModelAgentPluginDisable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAgentPluginDisable(nodeId string, pluginName string) *ModelAgentPluginDisable {
	this := ModelAgentPluginDisable{}
	this.NodeId = nodeId
	this.PluginName = pluginName
	return &this
}

// NewModelAgentPluginDisableWithDefaults instantiates a new ModelAgentPluginDisable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAgentPluginDisableWithDefaults() *ModelAgentPluginDisable {
	this := ModelAgentPluginDisable{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *ModelAgentPluginDisable) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelAgentPluginDisable) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelAgentPluginDisable) SetNodeId(v string) {
	o.NodeId = v
}

// GetPluginName returns the PluginName field value
func (o *ModelAgentPluginDisable) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *ModelAgentPluginDisable) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *ModelAgentPluginDisable) SetPluginName(v string) {
	o.PluginName = v
}

func (o ModelAgentPluginDisable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAgentPluginDisable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_id"] = o.NodeId
	toSerialize["plugin_name"] = o.PluginName
	return toSerialize, nil
}

func (o *ModelAgentPluginDisable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_id",
		"plugin_name",
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

	varModelAgentPluginDisable := _ModelAgentPluginDisable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelAgentPluginDisable)

	if err != nil {
		return err
	}

	*o = ModelAgentPluginDisable(varModelAgentPluginDisable)

	return err
}

type NullableModelAgentPluginDisable struct {
	value *ModelAgentPluginDisable
	isSet bool
}

func (v NullableModelAgentPluginDisable) Get() *ModelAgentPluginDisable {
	return v.value
}

func (v *NullableModelAgentPluginDisable) Set(val *ModelAgentPluginDisable) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAgentPluginDisable) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAgentPluginDisable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAgentPluginDisable(val *ModelAgentPluginDisable) *NullableModelAgentPluginDisable {
	return &NullableModelAgentPluginDisable{value: val, isSet: true}
}

func (v NullableModelAgentPluginDisable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAgentPluginDisable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


