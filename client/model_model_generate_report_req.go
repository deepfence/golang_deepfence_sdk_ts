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

// checks if the ModelGenerateReportReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerateReportReq{}

// ModelGenerateReportReq struct for ModelGenerateReportReq
type ModelGenerateReportReq struct {
	Filters *UtilsReportFilters `json:"filters,omitempty"`
	FromTimestamp *int32 `json:"from_timestamp,omitempty"`
	Options *UtilsReportOptions `json:"options,omitempty"`
	ReportType string `json:"report_type"`
	ToTimestamp *int32 `json:"to_timestamp,omitempty"`
	ZippedReport *bool `json:"zipped_report,omitempty"`
}

type _ModelGenerateReportReq ModelGenerateReportReq

// NewModelGenerateReportReq instantiates a new ModelGenerateReportReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerateReportReq(reportType string) *ModelGenerateReportReq {
	this := ModelGenerateReportReq{}
	this.ReportType = reportType
	return &this
}

// NewModelGenerateReportReqWithDefaults instantiates a new ModelGenerateReportReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerateReportReqWithDefaults() *ModelGenerateReportReq {
	this := ModelGenerateReportReq{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetFilters() UtilsReportFilters {
	if o == nil || IsNil(o.Filters) {
		var ret UtilsReportFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetFiltersOk() (*UtilsReportFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given UtilsReportFilters and assigns it to the Filters field.
func (o *ModelGenerateReportReq) SetFilters(v UtilsReportFilters) {
	o.Filters = &v
}

// GetFromTimestamp returns the FromTimestamp field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetFromTimestamp() int32 {
	if o == nil || IsNil(o.FromTimestamp) {
		var ret int32
		return ret
	}
	return *o.FromTimestamp
}

// GetFromTimestampOk returns a tuple with the FromTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetFromTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.FromTimestamp) {
		return nil, false
	}
	return o.FromTimestamp, true
}

// HasFromTimestamp returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasFromTimestamp() bool {
	if o != nil && !IsNil(o.FromTimestamp) {
		return true
	}

	return false
}

// SetFromTimestamp gets a reference to the given int32 and assigns it to the FromTimestamp field.
func (o *ModelGenerateReportReq) SetFromTimestamp(v int32) {
	o.FromTimestamp = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetOptions() UtilsReportOptions {
	if o == nil || IsNil(o.Options) {
		var ret UtilsReportOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetOptionsOk() (*UtilsReportOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given UtilsReportOptions and assigns it to the Options field.
func (o *ModelGenerateReportReq) SetOptions(v UtilsReportOptions) {
	o.Options = &v
}

// GetReportType returns the ReportType field value
func (o *ModelGenerateReportReq) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ModelGenerateReportReq) SetReportType(v string) {
	o.ReportType = v
}

// GetToTimestamp returns the ToTimestamp field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetToTimestamp() int32 {
	if o == nil || IsNil(o.ToTimestamp) {
		var ret int32
		return ret
	}
	return *o.ToTimestamp
}

// GetToTimestampOk returns a tuple with the ToTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetToTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.ToTimestamp) {
		return nil, false
	}
	return o.ToTimestamp, true
}

// HasToTimestamp returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasToTimestamp() bool {
	if o != nil && !IsNil(o.ToTimestamp) {
		return true
	}

	return false
}

// SetToTimestamp gets a reference to the given int32 and assigns it to the ToTimestamp field.
func (o *ModelGenerateReportReq) SetToTimestamp(v int32) {
	o.ToTimestamp = &v
}

// GetZippedReport returns the ZippedReport field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetZippedReport() bool {
	if o == nil || IsNil(o.ZippedReport) {
		var ret bool
		return ret
	}
	return *o.ZippedReport
}

// GetZippedReportOk returns a tuple with the ZippedReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetZippedReportOk() (*bool, bool) {
	if o == nil || IsNil(o.ZippedReport) {
		return nil, false
	}
	return o.ZippedReport, true
}

// HasZippedReport returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasZippedReport() bool {
	if o != nil && !IsNil(o.ZippedReport) {
		return true
	}

	return false
}

// SetZippedReport gets a reference to the given bool and assigns it to the ZippedReport field.
func (o *ModelGenerateReportReq) SetZippedReport(v bool) {
	o.ZippedReport = &v
}

func (o ModelGenerateReportReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerateReportReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.FromTimestamp) {
		toSerialize["from_timestamp"] = o.FromTimestamp
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["report_type"] = o.ReportType
	if !IsNil(o.ToTimestamp) {
		toSerialize["to_timestamp"] = o.ToTimestamp
	}
	if !IsNil(o.ZippedReport) {
		toSerialize["zipped_report"] = o.ZippedReport
	}
	return toSerialize, nil
}

func (o *ModelGenerateReportReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"report_type",
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

	varModelGenerateReportReq := _ModelGenerateReportReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelGenerateReportReq)

	if err != nil {
		return err
	}

	*o = ModelGenerateReportReq(varModelGenerateReportReq)

	return err
}

type NullableModelGenerateReportReq struct {
	value *ModelGenerateReportReq
	isSet bool
}

func (v NullableModelGenerateReportReq) Get() *ModelGenerateReportReq {
	return v.value
}

func (v *NullableModelGenerateReportReq) Set(val *ModelGenerateReportReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerateReportReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerateReportReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerateReportReq(val *ModelGenerateReportReq) *NullableModelGenerateReportReq {
	return &NullableModelGenerateReportReq{value: val, isSet: true}
}

func (v NullableModelGenerateReportReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerateReportReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


