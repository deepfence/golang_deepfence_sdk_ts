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

// checks if the ModelInviteUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInviteUserResponse{}

// ModelInviteUserResponse struct for ModelInviteUserResponse
type ModelInviteUserResponse struct {
	InviteExpiryHours *int32 `json:"invite_expiry_hours,omitempty"`
	InviteUrl *string `json:"invite_url,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewModelInviteUserResponse instantiates a new ModelInviteUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInviteUserResponse() *ModelInviteUserResponse {
	this := ModelInviteUserResponse{}
	return &this
}

// NewModelInviteUserResponseWithDefaults instantiates a new ModelInviteUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInviteUserResponseWithDefaults() *ModelInviteUserResponse {
	this := ModelInviteUserResponse{}
	return &this
}

// GetInviteExpiryHours returns the InviteExpiryHours field value if set, zero value otherwise.
func (o *ModelInviteUserResponse) GetInviteExpiryHours() int32 {
	if o == nil || IsNil(o.InviteExpiryHours) {
		var ret int32
		return ret
	}
	return *o.InviteExpiryHours
}

// GetInviteExpiryHoursOk returns a tuple with the InviteExpiryHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInviteUserResponse) GetInviteExpiryHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.InviteExpiryHours) {
		return nil, false
	}
	return o.InviteExpiryHours, true
}

// HasInviteExpiryHours returns a boolean if a field has been set.
func (o *ModelInviteUserResponse) HasInviteExpiryHours() bool {
	if o != nil && !IsNil(o.InviteExpiryHours) {
		return true
	}

	return false
}

// SetInviteExpiryHours gets a reference to the given int32 and assigns it to the InviteExpiryHours field.
func (o *ModelInviteUserResponse) SetInviteExpiryHours(v int32) {
	o.InviteExpiryHours = &v
}

// GetInviteUrl returns the InviteUrl field value if set, zero value otherwise.
func (o *ModelInviteUserResponse) GetInviteUrl() string {
	if o == nil || IsNil(o.InviteUrl) {
		var ret string
		return ret
	}
	return *o.InviteUrl
}

// GetInviteUrlOk returns a tuple with the InviteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInviteUserResponse) GetInviteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InviteUrl) {
		return nil, false
	}
	return o.InviteUrl, true
}

// HasInviteUrl returns a boolean if a field has been set.
func (o *ModelInviteUserResponse) HasInviteUrl() bool {
	if o != nil && !IsNil(o.InviteUrl) {
		return true
	}

	return false
}

// SetInviteUrl gets a reference to the given string and assigns it to the InviteUrl field.
func (o *ModelInviteUserResponse) SetInviteUrl(v string) {
	o.InviteUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ModelInviteUserResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInviteUserResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ModelInviteUserResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ModelInviteUserResponse) SetMessage(v string) {
	o.Message = &v
}

func (o ModelInviteUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInviteUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InviteExpiryHours) {
		toSerialize["invite_expiry_hours"] = o.InviteExpiryHours
	}
	if !IsNil(o.InviteUrl) {
		toSerialize["invite_url"] = o.InviteUrl
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableModelInviteUserResponse struct {
	value *ModelInviteUserResponse
	isSet bool
}

func (v NullableModelInviteUserResponse) Get() *ModelInviteUserResponse {
	return v.value
}

func (v *NullableModelInviteUserResponse) Set(val *ModelInviteUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInviteUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInviteUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInviteUserResponse(val *ModelInviteUserResponse) *NullableModelInviteUserResponse {
	return &NullableModelInviteUserResponse{value: val, isSet: true}
}

func (v NullableModelInviteUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInviteUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


