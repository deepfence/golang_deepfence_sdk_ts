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

// checks if the ModelComplinaceScanResultsGroupReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelComplinaceScanResultsGroupReq{}

// ModelComplinaceScanResultsGroupReq struct for ModelComplinaceScanResultsGroupReq
type ModelComplinaceScanResultsGroupReq struct {
	FieldsFilter ReportersFieldsFilters `json:"fields_filter"`
	ScanId string `json:"scan_id"`
}

type _ModelComplinaceScanResultsGroupReq ModelComplinaceScanResultsGroupReq

// NewModelComplinaceScanResultsGroupReq instantiates a new ModelComplinaceScanResultsGroupReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelComplinaceScanResultsGroupReq(fieldsFilter ReportersFieldsFilters, scanId string) *ModelComplinaceScanResultsGroupReq {
	this := ModelComplinaceScanResultsGroupReq{}
	this.FieldsFilter = fieldsFilter
	this.ScanId = scanId
	return &this
}

// NewModelComplinaceScanResultsGroupReqWithDefaults instantiates a new ModelComplinaceScanResultsGroupReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelComplinaceScanResultsGroupReqWithDefaults() *ModelComplinaceScanResultsGroupReq {
	this := ModelComplinaceScanResultsGroupReq{}
	return &this
}

// GetFieldsFilter returns the FieldsFilter field value
func (o *ModelComplinaceScanResultsGroupReq) GetFieldsFilter() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.FieldsFilter
}

// GetFieldsFilterOk returns a tuple with the FieldsFilter field value
// and a boolean to check if the value has been set.
func (o *ModelComplinaceScanResultsGroupReq) GetFieldsFilterOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldsFilter, true
}

// SetFieldsFilter sets field value
func (o *ModelComplinaceScanResultsGroupReq) SetFieldsFilter(v ReportersFieldsFilters) {
	o.FieldsFilter = v
}

// GetScanId returns the ScanId field value
func (o *ModelComplinaceScanResultsGroupReq) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelComplinaceScanResultsGroupReq) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelComplinaceScanResultsGroupReq) SetScanId(v string) {
	o.ScanId = v
}

func (o ModelComplinaceScanResultsGroupReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelComplinaceScanResultsGroupReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields_filter"] = o.FieldsFilter
	toSerialize["scan_id"] = o.ScanId
	return toSerialize, nil
}

func (o *ModelComplinaceScanResultsGroupReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields_filter",
		"scan_id",
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

	varModelComplinaceScanResultsGroupReq := _ModelComplinaceScanResultsGroupReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelComplinaceScanResultsGroupReq)

	if err != nil {
		return err
	}

	*o = ModelComplinaceScanResultsGroupReq(varModelComplinaceScanResultsGroupReq)

	return err
}

type NullableModelComplinaceScanResultsGroupReq struct {
	value *ModelComplinaceScanResultsGroupReq
	isSet bool
}

func (v NullableModelComplinaceScanResultsGroupReq) Get() *ModelComplinaceScanResultsGroupReq {
	return v.value
}

func (v *NullableModelComplinaceScanResultsGroupReq) Set(val *ModelComplinaceScanResultsGroupReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelComplinaceScanResultsGroupReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelComplinaceScanResultsGroupReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelComplinaceScanResultsGroupReq(val *ModelComplinaceScanResultsGroupReq) *NullableModelComplinaceScanResultsGroupReq {
	return &NullableModelComplinaceScanResultsGroupReq{value: val, isSet: true}
}

func (v NullableModelComplinaceScanResultsGroupReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelComplinaceScanResultsGroupReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


