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

// checks if the ReportRawReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportRawReport{}

// ReportRawReport struct for ReportRawReport
type ReportRawReport struct {
	Payload string `json:"payload"`
}

type _ReportRawReport ReportRawReport

// NewReportRawReport instantiates a new ReportRawReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportRawReport(payload string) *ReportRawReport {
	this := ReportRawReport{}
	this.Payload = payload
	return &this
}

// NewReportRawReportWithDefaults instantiates a new ReportRawReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportRawReportWithDefaults() *ReportRawReport {
	this := ReportRawReport{}
	return &this
}

// GetPayload returns the Payload field value
func (o *ReportRawReport) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *ReportRawReport) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *ReportRawReport) SetPayload(v string) {
	o.Payload = v
}

func (o ReportRawReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportRawReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

func (o *ReportRawReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payload",
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

	varReportRawReport := _ReportRawReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportRawReport)

	if err != nil {
		return err
	}

	*o = ReportRawReport(varReportRawReport)

	return err
}

type NullableReportRawReport struct {
	value *ReportRawReport
	isSet bool
}

func (v NullableReportRawReport) Get() *ReportRawReport {
	return v.value
}

func (v *NullableReportRawReport) Set(val *ReportRawReport) {
	v.value = val
	v.isSet = true
}

func (v NullableReportRawReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReportRawReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportRawReport(val *ReportRawReport) *NullableReportRawReport {
	return &NullableReportRawReport{value: val, isSet: true}
}

func (v NullableReportRawReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportRawReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


