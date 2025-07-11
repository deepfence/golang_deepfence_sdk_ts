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

// checks if the ModelLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLoginRequest{}

// ModelLoginRequest struct for ModelLoginRequest
type ModelLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type _ModelLoginRequest ModelLoginRequest

// NewModelLoginRequest instantiates a new ModelLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLoginRequest(email string, password string) *ModelLoginRequest {
	this := ModelLoginRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewModelLoginRequestWithDefaults instantiates a new ModelLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLoginRequestWithDefaults() *ModelLoginRequest {
	this := ModelLoginRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ModelLoginRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelLoginRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelLoginRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *ModelLoginRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ModelLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ModelLoginRequest) SetPassword(v string) {
	o.Password = v
}

func (o ModelLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *ModelLoginRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
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

	varModelLoginRequest := _ModelLoginRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelLoginRequest)

	if err != nil {
		return err
	}

	*o = ModelLoginRequest(varModelLoginRequest)

	return err
}

type NullableModelLoginRequest struct {
	value *ModelLoginRequest
	isSet bool
}

func (v NullableModelLoginRequest) Get() *ModelLoginRequest {
	return v.value
}

func (v *NullableModelLoginRequest) Set(val *ModelLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLoginRequest(val *ModelLoginRequest) *NullableModelLoginRequest {
	return &NullableModelLoginRequest{value: val, isSet: true}
}

func (v NullableModelLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


