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
)

// checks if the ModelConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelConnection{}

// ModelConnection struct for ModelConnection
type ModelConnection struct {
	Count *int32 `json:"count,omitempty"`
	Ips []interface{} `json:"ips,omitempty"`
	MaliciousIp []bool `json:"malicious_ip,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeName *string `json:"node_name,omitempty"`
}

// NewModelConnection instantiates a new ModelConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelConnection() *ModelConnection {
	this := ModelConnection{}
	return &this
}

// NewModelConnectionWithDefaults instantiates a new ModelConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelConnectionWithDefaults() *ModelConnection {
	this := ModelConnection{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelConnection) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelConnection) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelConnection) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelConnection) SetCount(v int32) {
	o.Count = &v
}

// GetIps returns the Ips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelConnection) GetIps() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelConnection) GetIpsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *ModelConnection) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []interface{} and assigns it to the Ips field.
func (o *ModelConnection) SetIps(v []interface{}) {
	o.Ips = v
}

// GetMaliciousIp returns the MaliciousIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelConnection) GetMaliciousIp() []bool {
	if o == nil {
		var ret []bool
		return ret
	}
	return o.MaliciousIp
}

// GetMaliciousIpOk returns a tuple with the MaliciousIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelConnection) GetMaliciousIpOk() ([]bool, bool) {
	if o == nil || IsNil(o.MaliciousIp) {
		return nil, false
	}
	return o.MaliciousIp, true
}

// HasMaliciousIp returns a boolean if a field has been set.
func (o *ModelConnection) HasMaliciousIp() bool {
	if o != nil && !IsNil(o.MaliciousIp) {
		return true
	}

	return false
}

// SetMaliciousIp gets a reference to the given []bool and assigns it to the MaliciousIp field.
func (o *ModelConnection) SetMaliciousIp(v []bool) {
	o.MaliciousIp = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ModelConnection) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelConnection) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ModelConnection) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ModelConnection) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *ModelConnection) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelConnection) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *ModelConnection) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *ModelConnection) SetNodeName(v string) {
	o.NodeName = &v
}

func (o ModelConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.MaliciousIp != nil {
		toSerialize["malicious_ip"] = o.MaliciousIp
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.NodeName) {
		toSerialize["node_name"] = o.NodeName
	}
	return toSerialize, nil
}

type NullableModelConnection struct {
	value *ModelConnection
	isSet bool
}

func (v NullableModelConnection) Get() *ModelConnection {
	return v.value
}

func (v *NullableModelConnection) Set(val *ModelConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableModelConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableModelConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelConnection(val *ModelConnection) *NullableModelConnection {
	return &NullableModelConnection{value: val, isSet: true}
}

func (v NullableModelConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


