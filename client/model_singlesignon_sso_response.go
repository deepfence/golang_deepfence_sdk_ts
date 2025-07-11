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
	"time"
	"bytes"
	"fmt"
)

// checks if the SinglesignonSSOResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SinglesignonSSOResponse{}

// SinglesignonSSOResponse struct for SinglesignonSSOResponse
type SinglesignonSSOResponse struct {
	ClientId string `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
	DisablePasswordLogin bool `json:"disable_password_login"`
	Id int32 `json:"id"`
	IssuerAliasUrl string `json:"issuer_alias_url"`
	IssuerUrl string `json:"issuer_url"`
	Label string `json:"label"`
	SsoProviderType string `json:"sso_provider_type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _SinglesignonSSOResponse SinglesignonSSOResponse

// NewSinglesignonSSOResponse instantiates a new SinglesignonSSOResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinglesignonSSOResponse(clientId string, createdAt time.Time, disablePasswordLogin bool, id int32, issuerAliasUrl string, issuerUrl string, label string, ssoProviderType string, updatedAt time.Time) *SinglesignonSSOResponse {
	this := SinglesignonSSOResponse{}
	this.ClientId = clientId
	this.CreatedAt = createdAt
	this.DisablePasswordLogin = disablePasswordLogin
	this.Id = id
	this.IssuerAliasUrl = issuerAliasUrl
	this.IssuerUrl = issuerUrl
	this.Label = label
	this.SsoProviderType = ssoProviderType
	this.UpdatedAt = updatedAt
	return &this
}

// NewSinglesignonSSOResponseWithDefaults instantiates a new SinglesignonSSOResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinglesignonSSOResponseWithDefaults() *SinglesignonSSOResponse {
	this := SinglesignonSSOResponse{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *SinglesignonSSOResponse) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *SinglesignonSSOResponse) SetClientId(v string) {
	o.ClientId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SinglesignonSSOResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SinglesignonSSOResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDisablePasswordLogin returns the DisablePasswordLogin field value
func (o *SinglesignonSSOResponse) GetDisablePasswordLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisablePasswordLogin
}

// GetDisablePasswordLoginOk returns a tuple with the DisablePasswordLogin field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetDisablePasswordLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisablePasswordLogin, true
}

// SetDisablePasswordLogin sets field value
func (o *SinglesignonSSOResponse) SetDisablePasswordLogin(v bool) {
	o.DisablePasswordLogin = v
}

// GetId returns the Id field value
func (o *SinglesignonSSOResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SinglesignonSSOResponse) SetId(v int32) {
	o.Id = v
}

// GetIssuerAliasUrl returns the IssuerAliasUrl field value
func (o *SinglesignonSSOResponse) GetIssuerAliasUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerAliasUrl
}

// GetIssuerAliasUrlOk returns a tuple with the IssuerAliasUrl field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetIssuerAliasUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerAliasUrl, true
}

// SetIssuerAliasUrl sets field value
func (o *SinglesignonSSOResponse) SetIssuerAliasUrl(v string) {
	o.IssuerAliasUrl = v
}

// GetIssuerUrl returns the IssuerUrl field value
func (o *SinglesignonSSOResponse) GetIssuerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetIssuerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerUrl, true
}

// SetIssuerUrl sets field value
func (o *SinglesignonSSOResponse) SetIssuerUrl(v string) {
	o.IssuerUrl = v
}

// GetLabel returns the Label field value
func (o *SinglesignonSSOResponse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SinglesignonSSOResponse) SetLabel(v string) {
	o.Label = v
}

// GetSsoProviderType returns the SsoProviderType field value
func (o *SinglesignonSSOResponse) GetSsoProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoProviderType
}

// GetSsoProviderTypeOk returns a tuple with the SsoProviderType field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetSsoProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoProviderType, true
}

// SetSsoProviderType sets field value
func (o *SinglesignonSSOResponse) SetSsoProviderType(v string) {
	o.SsoProviderType = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SinglesignonSSOResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SinglesignonSSOResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SinglesignonSSOResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SinglesignonSSOResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SinglesignonSSOResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_id"] = o.ClientId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["disable_password_login"] = o.DisablePasswordLogin
	toSerialize["id"] = o.Id
	toSerialize["issuer_alias_url"] = o.IssuerAliasUrl
	toSerialize["issuer_url"] = o.IssuerUrl
	toSerialize["label"] = o.Label
	toSerialize["sso_provider_type"] = o.SsoProviderType
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *SinglesignonSSOResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
		"created_at",
		"disable_password_login",
		"id",
		"issuer_alias_url",
		"issuer_url",
		"label",
		"sso_provider_type",
		"updated_at",
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

	varSinglesignonSSOResponse := _SinglesignonSSOResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSinglesignonSSOResponse)

	if err != nil {
		return err
	}

	*o = SinglesignonSSOResponse(varSinglesignonSSOResponse)

	return err
}

type NullableSinglesignonSSOResponse struct {
	value *SinglesignonSSOResponse
	isSet bool
}

func (v NullableSinglesignonSSOResponse) Get() *SinglesignonSSOResponse {
	return v.value
}

func (v *NullableSinglesignonSSOResponse) Set(val *SinglesignonSSOResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSinglesignonSSOResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSinglesignonSSOResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinglesignonSSOResponse(val *SinglesignonSSOResponse) *NullableSinglesignonSSOResponse {
	return &NullableSinglesignonSSOResponse{value: val, isSet: true}
}

func (v NullableSinglesignonSSOResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinglesignonSSOResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


