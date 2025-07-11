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

// checks if the IngestersCloudWafConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersCloudWafConfig{}

// IngestersCloudWafConfig struct for IngestersCloudWafConfig
type IngestersCloudWafConfig struct {
	AwsWafArn []IngestersAWSWafARN `json:"aws_waf_arn"`
	CloudProvider string `json:"cloud_provider"`
}

type _IngestersCloudWafConfig IngestersCloudWafConfig

// NewIngestersCloudWafConfig instantiates a new IngestersCloudWafConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersCloudWafConfig(awsWafArn []IngestersAWSWafARN, cloudProvider string) *IngestersCloudWafConfig {
	this := IngestersCloudWafConfig{}
	this.AwsWafArn = awsWafArn
	this.CloudProvider = cloudProvider
	return &this
}

// NewIngestersCloudWafConfigWithDefaults instantiates a new IngestersCloudWafConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersCloudWafConfigWithDefaults() *IngestersCloudWafConfig {
	this := IngestersCloudWafConfig{}
	return &this
}

// GetAwsWafArn returns the AwsWafArn field value
// If the value is explicit nil, the zero value for []IngestersAWSWafARN will be returned
func (o *IngestersCloudWafConfig) GetAwsWafArn() []IngestersAWSWafARN {
	if o == nil {
		var ret []IngestersAWSWafARN
		return ret
	}

	return o.AwsWafArn
}

// GetAwsWafArnOk returns a tuple with the AwsWafArn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersCloudWafConfig) GetAwsWafArnOk() ([]IngestersAWSWafARN, bool) {
	if o == nil || IsNil(o.AwsWafArn) {
		return nil, false
	}
	return o.AwsWafArn, true
}

// SetAwsWafArn sets field value
func (o *IngestersCloudWafConfig) SetAwsWafArn(v []IngestersAWSWafARN) {
	o.AwsWafArn = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *IngestersCloudWafConfig) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *IngestersCloudWafConfig) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *IngestersCloudWafConfig) SetCloudProvider(v string) {
	o.CloudProvider = v
}

func (o IngestersCloudWafConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersCloudWafConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsWafArn != nil {
		toSerialize["aws_waf_arn"] = o.AwsWafArn
	}
	toSerialize["cloud_provider"] = o.CloudProvider
	return toSerialize, nil
}

func (o *IngestersCloudWafConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aws_waf_arn",
		"cloud_provider",
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

	varIngestersCloudWafConfig := _IngestersCloudWafConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestersCloudWafConfig)

	if err != nil {
		return err
	}

	*o = IngestersCloudWafConfig(varIngestersCloudWafConfig)

	return err
}

type NullableIngestersCloudWafConfig struct {
	value *IngestersCloudWafConfig
	isSet bool
}

func (v NullableIngestersCloudWafConfig) Get() *IngestersCloudWafConfig {
	return v.value
}

func (v *NullableIngestersCloudWafConfig) Set(val *IngestersCloudWafConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersCloudWafConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersCloudWafConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersCloudWafConfig(val *IngestersCloudWafConfig) *NullableIngestersCloudWafConfig {
	return &NullableIngestersCloudWafConfig{value: val, isSet: true}
}

func (v NullableIngestersCloudWafConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersCloudWafConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


