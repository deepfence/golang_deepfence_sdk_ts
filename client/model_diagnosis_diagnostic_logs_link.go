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

// checks if the DiagnosisDiagnosticLogsLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisDiagnosticLogsLink{}

// DiagnosisDiagnosticLogsLink struct for DiagnosisDiagnosticLogsLink
type DiagnosisDiagnosticLogsLink struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Label *string `json:"label,omitempty"`
	Message *string `json:"message,omitempty"`
	Type *string `json:"type,omitempty"`
	UrlLink *string `json:"url_link,omitempty"`
}

// NewDiagnosisDiagnosticLogsLink instantiates a new DiagnosisDiagnosticLogsLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisDiagnosticLogsLink() *DiagnosisDiagnosticLogsLink {
	this := DiagnosisDiagnosticLogsLink{}
	return &this
}

// NewDiagnosisDiagnosticLogsLinkWithDefaults instantiates a new DiagnosisDiagnosticLogsLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisDiagnosticLogsLinkWithDefaults() *DiagnosisDiagnosticLogsLink {
	this := DiagnosisDiagnosticLogsLink{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DiagnosisDiagnosticLogsLink) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsLink) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DiagnosisDiagnosticLogsLink) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DiagnosisDiagnosticLogsLink) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DiagnosisDiagnosticLogsLink) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsLink) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DiagnosisDiagnosticLogsLink) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DiagnosisDiagnosticLogsLink) SetLabel(v string) {
	o.Label = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DiagnosisDiagnosticLogsLink) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsLink) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DiagnosisDiagnosticLogsLink) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DiagnosisDiagnosticLogsLink) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DiagnosisDiagnosticLogsLink) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsLink) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DiagnosisDiagnosticLogsLink) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DiagnosisDiagnosticLogsLink) SetType(v string) {
	o.Type = &v
}

// GetUrlLink returns the UrlLink field value if set, zero value otherwise.
func (o *DiagnosisDiagnosticLogsLink) GetUrlLink() string {
	if o == nil || IsNil(o.UrlLink) {
		var ret string
		return ret
	}
	return *o.UrlLink
}

// GetUrlLinkOk returns a tuple with the UrlLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsLink) GetUrlLinkOk() (*string, bool) {
	if o == nil || IsNil(o.UrlLink) {
		return nil, false
	}
	return o.UrlLink, true
}

// HasUrlLink returns a boolean if a field has been set.
func (o *DiagnosisDiagnosticLogsLink) HasUrlLink() bool {
	if o != nil && !IsNil(o.UrlLink) {
		return true
	}

	return false
}

// SetUrlLink gets a reference to the given string and assigns it to the UrlLink field.
func (o *DiagnosisDiagnosticLogsLink) SetUrlLink(v string) {
	o.UrlLink = &v
}

func (o DiagnosisDiagnosticLogsLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisDiagnosticLogsLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UrlLink) {
		toSerialize["url_link"] = o.UrlLink
	}
	return toSerialize, nil
}

type NullableDiagnosisDiagnosticLogsLink struct {
	value *DiagnosisDiagnosticLogsLink
	isSet bool
}

func (v NullableDiagnosisDiagnosticLogsLink) Get() *DiagnosisDiagnosticLogsLink {
	return v.value
}

func (v *NullableDiagnosisDiagnosticLogsLink) Set(val *DiagnosisDiagnosticLogsLink) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisDiagnosticLogsLink) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisDiagnosticLogsLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisDiagnosticLogsLink(val *DiagnosisDiagnosticLogsLink) *NullableDiagnosisDiagnosticLogsLink {
	return &NullableDiagnosisDiagnosticLogsLink{value: val, isSet: true}
}

func (v NullableDiagnosisDiagnosticLogsLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisDiagnosticLogsLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


