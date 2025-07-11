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

// checks if the ModelNodesInScanResultRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelNodesInScanResultRequest{}

// ModelNodesInScanResultRequest struct for ModelNodesInScanResultRequest
type ModelNodesInScanResultRequest struct {
	ResultIds []string `json:"result_ids"`
	ScanType string `json:"scan_type"`
}

type _ModelNodesInScanResultRequest ModelNodesInScanResultRequest

// NewModelNodesInScanResultRequest instantiates a new ModelNodesInScanResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelNodesInScanResultRequest(resultIds []string, scanType string) *ModelNodesInScanResultRequest {
	this := ModelNodesInScanResultRequest{}
	this.ResultIds = resultIds
	this.ScanType = scanType
	return &this
}

// NewModelNodesInScanResultRequestWithDefaults instantiates a new ModelNodesInScanResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelNodesInScanResultRequestWithDefaults() *ModelNodesInScanResultRequest {
	this := ModelNodesInScanResultRequest{}
	return &this
}

// GetResultIds returns the ResultIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelNodesInScanResultRequest) GetResultIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResultIds
}

// GetResultIdsOk returns a tuple with the ResultIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelNodesInScanResultRequest) GetResultIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultIds) {
		return nil, false
	}
	return o.ResultIds, true
}

// SetResultIds sets field value
func (o *ModelNodesInScanResultRequest) SetResultIds(v []string) {
	o.ResultIds = v
}

// GetScanType returns the ScanType field value
func (o *ModelNodesInScanResultRequest) GetScanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value
// and a boolean to check if the value has been set.
func (o *ModelNodesInScanResultRequest) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanType, true
}

// SetScanType sets field value
func (o *ModelNodesInScanResultRequest) SetScanType(v string) {
	o.ScanType = v
}

func (o ModelNodesInScanResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelNodesInScanResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResultIds != nil {
		toSerialize["result_ids"] = o.ResultIds
	}
	toSerialize["scan_type"] = o.ScanType
	return toSerialize, nil
}

func (o *ModelNodesInScanResultRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result_ids",
		"scan_type",
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

	varModelNodesInScanResultRequest := _ModelNodesInScanResultRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelNodesInScanResultRequest)

	if err != nil {
		return err
	}

	*o = ModelNodesInScanResultRequest(varModelNodesInScanResultRequest)

	return err
}

type NullableModelNodesInScanResultRequest struct {
	value *ModelNodesInScanResultRequest
	isSet bool
}

func (v NullableModelNodesInScanResultRequest) Get() *ModelNodesInScanResultRequest {
	return v.value
}

func (v *NullableModelNodesInScanResultRequest) Set(val *ModelNodesInScanResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelNodesInScanResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelNodesInScanResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelNodesInScanResultRequest(val *ModelNodesInScanResultRequest) *NullableModelNodesInScanResultRequest {
	return &NullableModelNodesInScanResultRequest{value: val, isSet: true}
}

func (v NullableModelNodesInScanResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelNodesInScanResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


