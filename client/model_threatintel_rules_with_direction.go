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

// checks if the ThreatintelRulesWithDirection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatintelRulesWithDirection{}

// ThreatintelRulesWithDirection struct for ThreatintelRulesWithDirection
type ThreatintelRulesWithDirection struct {
	Inbound map[string]string `json:"inbound,omitempty"`
	Outbound map[string]string `json:"outbound,omitempty"`
}

// NewThreatintelRulesWithDirection instantiates a new ThreatintelRulesWithDirection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatintelRulesWithDirection() *ThreatintelRulesWithDirection {
	this := ThreatintelRulesWithDirection{}
	return &this
}

// NewThreatintelRulesWithDirectionWithDefaults instantiates a new ThreatintelRulesWithDirection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatintelRulesWithDirectionWithDefaults() *ThreatintelRulesWithDirection {
	this := ThreatintelRulesWithDirection{}
	return &this
}

// GetInbound returns the Inbound field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreatintelRulesWithDirection) GetInbound() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreatintelRulesWithDirection) GetInboundOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Inbound) {
		return nil, false
	}
	return &o.Inbound, true
}

// HasInbound returns a boolean if a field has been set.
func (o *ThreatintelRulesWithDirection) HasInbound() bool {
	if o != nil && !IsNil(o.Inbound) {
		return true
	}

	return false
}

// SetInbound gets a reference to the given map[string]string and assigns it to the Inbound field.
func (o *ThreatintelRulesWithDirection) SetInbound(v map[string]string) {
	o.Inbound = v
}

// GetOutbound returns the Outbound field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreatintelRulesWithDirection) GetOutbound() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreatintelRulesWithDirection) GetOutboundOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Outbound) {
		return nil, false
	}
	return &o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *ThreatintelRulesWithDirection) HasOutbound() bool {
	if o != nil && !IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given map[string]string and assigns it to the Outbound field.
func (o *ThreatintelRulesWithDirection) SetOutbound(v map[string]string) {
	o.Outbound = v
}

func (o ThreatintelRulesWithDirection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatintelRulesWithDirection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Inbound != nil {
		toSerialize["inbound"] = o.Inbound
	}
	if o.Outbound != nil {
		toSerialize["outbound"] = o.Outbound
	}
	return toSerialize, nil
}

type NullableThreatintelRulesWithDirection struct {
	value *ThreatintelRulesWithDirection
	isSet bool
}

func (v NullableThreatintelRulesWithDirection) Get() *ThreatintelRulesWithDirection {
	return v.value
}

func (v *NullableThreatintelRulesWithDirection) Set(val *ThreatintelRulesWithDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatintelRulesWithDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatintelRulesWithDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatintelRulesWithDirection(val *ThreatintelRulesWithDirection) *NullableThreatintelRulesWithDirection {
	return &NullableThreatintelRulesWithDirection{value: val, isSet: true}
}

func (v NullableThreatintelRulesWithDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatintelRulesWithDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


