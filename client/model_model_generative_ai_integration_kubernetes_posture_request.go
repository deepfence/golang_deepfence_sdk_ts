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

// checks if the ModelGenerativeAiIntegrationKubernetesPostureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerativeAiIntegrationKubernetesPostureRequest{}

// ModelGenerativeAiIntegrationKubernetesPostureRequest struct for ModelGenerativeAiIntegrationKubernetesPostureRequest
type ModelGenerativeAiIntegrationKubernetesPostureRequest struct {
	ComplianceCheckType string `json:"compliance_check_type"`
	Description string `json:"description"`
	IntegrationId *int32 `json:"integration_id,omitempty"`
	QueryType string `json:"query_type"`
	RemediationFormat string `json:"remediation_format"`
}

type _ModelGenerativeAiIntegrationKubernetesPostureRequest ModelGenerativeAiIntegrationKubernetesPostureRequest

// NewModelGenerativeAiIntegrationKubernetesPostureRequest instantiates a new ModelGenerativeAiIntegrationKubernetesPostureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerativeAiIntegrationKubernetesPostureRequest(complianceCheckType string, description string, queryType string, remediationFormat string) *ModelGenerativeAiIntegrationKubernetesPostureRequest {
	this := ModelGenerativeAiIntegrationKubernetesPostureRequest{}
	this.ComplianceCheckType = complianceCheckType
	this.Description = description
	this.QueryType = queryType
	this.RemediationFormat = remediationFormat
	return &this
}

// NewModelGenerativeAiIntegrationKubernetesPostureRequestWithDefaults instantiates a new ModelGenerativeAiIntegrationKubernetesPostureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerativeAiIntegrationKubernetesPostureRequestWithDefaults() *ModelGenerativeAiIntegrationKubernetesPostureRequest {
	this := ModelGenerativeAiIntegrationKubernetesPostureRequest{}
	return &this
}

// GetComplianceCheckType returns the ComplianceCheckType field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetComplianceCheckType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceCheckType
}

// GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field value
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetComplianceCheckTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCheckType, true
}

// SetComplianceCheckType sets field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) SetComplianceCheckType(v string) {
	o.ComplianceCheckType = v
}

// GetDescription returns the Description field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) SetDescription(v string) {
	o.Description = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetIntegrationId() int32 {
	if o == nil || IsNil(o.IntegrationId) {
		var ret int32
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetIntegrationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given int32 and assigns it to the IntegrationId field.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) SetIntegrationId(v int32) {
	o.IntegrationId = &v
}

// GetQueryType returns the QueryType field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetRemediationFormat returns the RemediationFormat field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetRemediationFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemediationFormat
}

// GetRemediationFormatOk returns a tuple with the RemediationFormat field value
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) GetRemediationFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationFormat, true
}

// SetRemediationFormat sets field value
func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) SetRemediationFormat(v string) {
	o.RemediationFormat = v
}

func (o ModelGenerativeAiIntegrationKubernetesPostureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerativeAiIntegrationKubernetesPostureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["compliance_check_type"] = o.ComplianceCheckType
	toSerialize["description"] = o.Description
	if !IsNil(o.IntegrationId) {
		toSerialize["integration_id"] = o.IntegrationId
	}
	toSerialize["query_type"] = o.QueryType
	toSerialize["remediation_format"] = o.RemediationFormat
	return toSerialize, nil
}

func (o *ModelGenerativeAiIntegrationKubernetesPostureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"compliance_check_type",
		"description",
		"query_type",
		"remediation_format",
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

	varModelGenerativeAiIntegrationKubernetesPostureRequest := _ModelGenerativeAiIntegrationKubernetesPostureRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelGenerativeAiIntegrationKubernetesPostureRequest)

	if err != nil {
		return err
	}

	*o = ModelGenerativeAiIntegrationKubernetesPostureRequest(varModelGenerativeAiIntegrationKubernetesPostureRequest)

	return err
}

type NullableModelGenerativeAiIntegrationKubernetesPostureRequest struct {
	value *ModelGenerativeAiIntegrationKubernetesPostureRequest
	isSet bool
}

func (v NullableModelGenerativeAiIntegrationKubernetesPostureRequest) Get() *ModelGenerativeAiIntegrationKubernetesPostureRequest {
	return v.value
}

func (v *NullableModelGenerativeAiIntegrationKubernetesPostureRequest) Set(val *ModelGenerativeAiIntegrationKubernetesPostureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerativeAiIntegrationKubernetesPostureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerativeAiIntegrationKubernetesPostureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerativeAiIntegrationKubernetesPostureRequest(val *ModelGenerativeAiIntegrationKubernetesPostureRequest) *NullableModelGenerativeAiIntegrationKubernetesPostureRequest {
	return &NullableModelGenerativeAiIntegrationKubernetesPostureRequest{value: val, isSet: true}
}

func (v NullableModelGenerativeAiIntegrationKubernetesPostureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerativeAiIntegrationKubernetesPostureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


