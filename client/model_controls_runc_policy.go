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

// checks if the ControlsRuncPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsRuncPolicy{}

// ControlsRuncPolicy struct for ControlsRuncPolicy
type ControlsRuncPolicy struct {
	Action string `json:"action"`
	CountLimit int32 `json:"count_limit"`
	DurationCountLimitSec int32 `json:"duration_count_limit_sec"`
	Matcher ControlsPolicyAlertMatcher `json:"matcher"`
	NodeType string `json:"node_type"`
	PolicyId string `json:"policy_id"`
	UpdatedAt int32 `json:"updated_at"`
}

type _ControlsRuncPolicy ControlsRuncPolicy

// NewControlsRuncPolicy instantiates a new ControlsRuncPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsRuncPolicy(action string, countLimit int32, durationCountLimitSec int32, matcher ControlsPolicyAlertMatcher, nodeType string, policyId string, updatedAt int32) *ControlsRuncPolicy {
	this := ControlsRuncPolicy{}
	this.Action = action
	this.CountLimit = countLimit
	this.DurationCountLimitSec = durationCountLimitSec
	this.Matcher = matcher
	this.NodeType = nodeType
	this.PolicyId = policyId
	this.UpdatedAt = updatedAt
	return &this
}

// NewControlsRuncPolicyWithDefaults instantiates a new ControlsRuncPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsRuncPolicyWithDefaults() *ControlsRuncPolicy {
	this := ControlsRuncPolicy{}
	return &this
}

// GetAction returns the Action field value
func (o *ControlsRuncPolicy) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ControlsRuncPolicy) SetAction(v string) {
	o.Action = v
}

// GetCountLimit returns the CountLimit field value
func (o *ControlsRuncPolicy) GetCountLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CountLimit
}

// GetCountLimitOk returns a tuple with the CountLimit field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetCountLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountLimit, true
}

// SetCountLimit sets field value
func (o *ControlsRuncPolicy) SetCountLimit(v int32) {
	o.CountLimit = v
}

// GetDurationCountLimitSec returns the DurationCountLimitSec field value
func (o *ControlsRuncPolicy) GetDurationCountLimitSec() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationCountLimitSec
}

// GetDurationCountLimitSecOk returns a tuple with the DurationCountLimitSec field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetDurationCountLimitSecOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationCountLimitSec, true
}

// SetDurationCountLimitSec sets field value
func (o *ControlsRuncPolicy) SetDurationCountLimitSec(v int32) {
	o.DurationCountLimitSec = v
}

// GetMatcher returns the Matcher field value
func (o *ControlsRuncPolicy) GetMatcher() ControlsPolicyAlertMatcher {
	if o == nil {
		var ret ControlsPolicyAlertMatcher
		return ret
	}

	return o.Matcher
}

// GetMatcherOk returns a tuple with the Matcher field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetMatcherOk() (*ControlsPolicyAlertMatcher, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matcher, true
}

// SetMatcher sets field value
func (o *ControlsRuncPolicy) SetMatcher(v ControlsPolicyAlertMatcher) {
	o.Matcher = v
}

// GetNodeType returns the NodeType field value
func (o *ControlsRuncPolicy) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ControlsRuncPolicy) SetNodeType(v string) {
	o.NodeType = v
}

// GetPolicyId returns the PolicyId field value
func (o *ControlsRuncPolicy) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *ControlsRuncPolicy) SetPolicyId(v string) {
	o.PolicyId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ControlsRuncPolicy) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ControlsRuncPolicy) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ControlsRuncPolicy) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ControlsRuncPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsRuncPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["count_limit"] = o.CountLimit
	toSerialize["duration_count_limit_sec"] = o.DurationCountLimitSec
	toSerialize["matcher"] = o.Matcher
	toSerialize["node_type"] = o.NodeType
	toSerialize["policy_id"] = o.PolicyId
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ControlsRuncPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"count_limit",
		"duration_count_limit_sec",
		"matcher",
		"node_type",
		"policy_id",
		"updated_at",
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

	varControlsRuncPolicy := _ControlsRuncPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varControlsRuncPolicy)

	if err != nil {
		return err
	}

	*o = ControlsRuncPolicy(varControlsRuncPolicy)

	return err
}

type NullableControlsRuncPolicy struct {
	value *ControlsRuncPolicy
	isSet bool
}

func (v NullableControlsRuncPolicy) Get() *ControlsRuncPolicy {
	return v.value
}

func (v *NullableControlsRuncPolicy) Set(val *ControlsRuncPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsRuncPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsRuncPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsRuncPolicy(val *ControlsRuncPolicy) *NullableControlsRuncPolicy {
	return &NullableControlsRuncPolicy{value: val, isSet: true}
}

func (v NullableControlsRuncPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsRuncPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


