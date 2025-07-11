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

// checks if the ModelIntegrationAddReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelIntegrationAddReq{}

// ModelIntegrationAddReq struct for ModelIntegrationAddReq
type ModelIntegrationAddReq struct {
	Config map[string]interface{} `json:"config,omitempty"`
	Filters *ModelIntegrationFilters `json:"filters,omitempty"`
	IntegrationType string `json:"integration_type"`
	NotificationType string `json:"notification_type"`
	SendSummary *bool `json:"send_summary,omitempty"`
}

type _ModelIntegrationAddReq ModelIntegrationAddReq

// NewModelIntegrationAddReq instantiates a new ModelIntegrationAddReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelIntegrationAddReq(integrationType string, notificationType string) *ModelIntegrationAddReq {
	this := ModelIntegrationAddReq{}
	this.IntegrationType = integrationType
	this.NotificationType = notificationType
	return &this
}

// NewModelIntegrationAddReqWithDefaults instantiates a new ModelIntegrationAddReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelIntegrationAddReqWithDefaults() *ModelIntegrationAddReq {
	this := ModelIntegrationAddReq{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelIntegrationAddReq) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelIntegrationAddReq) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ModelIntegrationAddReq) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ModelIntegrationAddReq) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ModelIntegrationAddReq) GetFilters() ModelIntegrationFilters {
	if o == nil || IsNil(o.Filters) {
		var ret ModelIntegrationFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationAddReq) GetFiltersOk() (*ModelIntegrationFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ModelIntegrationAddReq) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ModelIntegrationFilters and assigns it to the Filters field.
func (o *ModelIntegrationAddReq) SetFilters(v ModelIntegrationFilters) {
	o.Filters = &v
}

// GetIntegrationType returns the IntegrationType field value
func (o *ModelIntegrationAddReq) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *ModelIntegrationAddReq) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *ModelIntegrationAddReq) SetIntegrationType(v string) {
	o.IntegrationType = v
}

// GetNotificationType returns the NotificationType field value
func (o *ModelIntegrationAddReq) GetNotificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *ModelIntegrationAddReq) GetNotificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *ModelIntegrationAddReq) SetNotificationType(v string) {
	o.NotificationType = v
}

// GetSendSummary returns the SendSummary field value if set, zero value otherwise.
func (o *ModelIntegrationAddReq) GetSendSummary() bool {
	if o == nil || IsNil(o.SendSummary) {
		var ret bool
		return ret
	}
	return *o.SendSummary
}

// GetSendSummaryOk returns a tuple with the SendSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationAddReq) GetSendSummaryOk() (*bool, bool) {
	if o == nil || IsNil(o.SendSummary) {
		return nil, false
	}
	return o.SendSummary, true
}

// HasSendSummary returns a boolean if a field has been set.
func (o *ModelIntegrationAddReq) HasSendSummary() bool {
	if o != nil && !IsNil(o.SendSummary) {
		return true
	}

	return false
}

// SetSendSummary gets a reference to the given bool and assigns it to the SendSummary field.
func (o *ModelIntegrationAddReq) SetSendSummary(v bool) {
	o.SendSummary = &v
}

func (o ModelIntegrationAddReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelIntegrationAddReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["notification_type"] = o.NotificationType
	if !IsNil(o.SendSummary) {
		toSerialize["send_summary"] = o.SendSummary
	}
	return toSerialize, nil
}

func (o *ModelIntegrationAddReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"integration_type",
		"notification_type",
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

	varModelIntegrationAddReq := _ModelIntegrationAddReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelIntegrationAddReq)

	if err != nil {
		return err
	}

	*o = ModelIntegrationAddReq(varModelIntegrationAddReq)

	return err
}

type NullableModelIntegrationAddReq struct {
	value *ModelIntegrationAddReq
	isSet bool
}

func (v NullableModelIntegrationAddReq) Get() *ModelIntegrationAddReq {
	return v.value
}

func (v *NullableModelIntegrationAddReq) Set(val *ModelIntegrationAddReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelIntegrationAddReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelIntegrationAddReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelIntegrationAddReq(val *ModelIntegrationAddReq) *NullableModelIntegrationAddReq {
	return &NullableModelIntegrationAddReq{value: val, isSet: true}
}

func (v NullableModelIntegrationAddReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelIntegrationAddReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


