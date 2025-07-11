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

// checks if the ModelAddGenerativeAiBedrockIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAddGenerativeAiBedrockIntegration{}

// ModelAddGenerativeAiBedrockIntegration struct for ModelAddGenerativeAiBedrockIntegration
type ModelAddGenerativeAiBedrockIntegration struct {
	AwsAccessKey *string `json:"aws_access_key,omitempty"`
	AwsRegion string `json:"aws_region"`
	AwsSecretKey *string `json:"aws_secret_key,omitempty"`
	ModelId string `json:"model_id"`
	UseIamRole *bool `json:"use_iam_role,omitempty"`
}

type _ModelAddGenerativeAiBedrockIntegration ModelAddGenerativeAiBedrockIntegration

// NewModelAddGenerativeAiBedrockIntegration instantiates a new ModelAddGenerativeAiBedrockIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAddGenerativeAiBedrockIntegration(awsRegion string, modelId string) *ModelAddGenerativeAiBedrockIntegration {
	this := ModelAddGenerativeAiBedrockIntegration{}
	this.AwsRegion = awsRegion
	this.ModelId = modelId
	return &this
}

// NewModelAddGenerativeAiBedrockIntegrationWithDefaults instantiates a new ModelAddGenerativeAiBedrockIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAddGenerativeAiBedrockIntegrationWithDefaults() *ModelAddGenerativeAiBedrockIntegration {
	this := ModelAddGenerativeAiBedrockIntegration{}
	return &this
}

// GetAwsAccessKey returns the AwsAccessKey field value if set, zero value otherwise.
func (o *ModelAddGenerativeAiBedrockIntegration) GetAwsAccessKey() string {
	if o == nil || IsNil(o.AwsAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsAccessKey
}

// GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) GetAwsAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccessKey) {
		return nil, false
	}
	return o.AwsAccessKey, true
}

// HasAwsAccessKey returns a boolean if a field has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) HasAwsAccessKey() bool {
	if o != nil && !IsNil(o.AwsAccessKey) {
		return true
	}

	return false
}

// SetAwsAccessKey gets a reference to the given string and assigns it to the AwsAccessKey field.
func (o *ModelAddGenerativeAiBedrockIntegration) SetAwsAccessKey(v string) {
	o.AwsAccessKey = &v
}

// GetAwsRegion returns the AwsRegion field value
func (o *ModelAddGenerativeAiBedrockIntegration) GetAwsRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) GetAwsRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRegion, true
}

// SetAwsRegion sets field value
func (o *ModelAddGenerativeAiBedrockIntegration) SetAwsRegion(v string) {
	o.AwsRegion = v
}

// GetAwsSecretKey returns the AwsSecretKey field value if set, zero value otherwise.
func (o *ModelAddGenerativeAiBedrockIntegration) GetAwsSecretKey() string {
	if o == nil || IsNil(o.AwsSecretKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretKey
}

// GetAwsSecretKeyOk returns a tuple with the AwsSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) GetAwsSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsSecretKey) {
		return nil, false
	}
	return o.AwsSecretKey, true
}

// HasAwsSecretKey returns a boolean if a field has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) HasAwsSecretKey() bool {
	if o != nil && !IsNil(o.AwsSecretKey) {
		return true
	}

	return false
}

// SetAwsSecretKey gets a reference to the given string and assigns it to the AwsSecretKey field.
func (o *ModelAddGenerativeAiBedrockIntegration) SetAwsSecretKey(v string) {
	o.AwsSecretKey = &v
}

// GetModelId returns the ModelId field value
func (o *ModelAddGenerativeAiBedrockIntegration) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *ModelAddGenerativeAiBedrockIntegration) SetModelId(v string) {
	o.ModelId = v
}

// GetUseIamRole returns the UseIamRole field value if set, zero value otherwise.
func (o *ModelAddGenerativeAiBedrockIntegration) GetUseIamRole() bool {
	if o == nil || IsNil(o.UseIamRole) {
		var ret bool
		return ret
	}
	return *o.UseIamRole
}

// GetUseIamRoleOk returns a tuple with the UseIamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) GetUseIamRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.UseIamRole) {
		return nil, false
	}
	return o.UseIamRole, true
}

// HasUseIamRole returns a boolean if a field has been set.
func (o *ModelAddGenerativeAiBedrockIntegration) HasUseIamRole() bool {
	if o != nil && !IsNil(o.UseIamRole) {
		return true
	}

	return false
}

// SetUseIamRole gets a reference to the given bool and assigns it to the UseIamRole field.
func (o *ModelAddGenerativeAiBedrockIntegration) SetUseIamRole(v bool) {
	o.UseIamRole = &v
}

func (o ModelAddGenerativeAiBedrockIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAddGenerativeAiBedrockIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccessKey) {
		toSerialize["aws_access_key"] = o.AwsAccessKey
	}
	toSerialize["aws_region"] = o.AwsRegion
	if !IsNil(o.AwsSecretKey) {
		toSerialize["aws_secret_key"] = o.AwsSecretKey
	}
	toSerialize["model_id"] = o.ModelId
	if !IsNil(o.UseIamRole) {
		toSerialize["use_iam_role"] = o.UseIamRole
	}
	return toSerialize, nil
}

func (o *ModelAddGenerativeAiBedrockIntegration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aws_region",
		"model_id",
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

	varModelAddGenerativeAiBedrockIntegration := _ModelAddGenerativeAiBedrockIntegration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelAddGenerativeAiBedrockIntegration)

	if err != nil {
		return err
	}

	*o = ModelAddGenerativeAiBedrockIntegration(varModelAddGenerativeAiBedrockIntegration)

	return err
}

type NullableModelAddGenerativeAiBedrockIntegration struct {
	value *ModelAddGenerativeAiBedrockIntegration
	isSet bool
}

func (v NullableModelAddGenerativeAiBedrockIntegration) Get() *ModelAddGenerativeAiBedrockIntegration {
	return v.value
}

func (v *NullableModelAddGenerativeAiBedrockIntegration) Set(val *ModelAddGenerativeAiBedrockIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAddGenerativeAiBedrockIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAddGenerativeAiBedrockIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAddGenerativeAiBedrockIntegration(val *ModelAddGenerativeAiBedrockIntegration) *NullableModelAddGenerativeAiBedrockIntegration {
	return &NullableModelAddGenerativeAiBedrockIntegration{value: val, isSet: true}
}

func (v NullableModelAddGenerativeAiBedrockIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAddGenerativeAiBedrockIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


