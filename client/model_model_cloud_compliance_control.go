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

// checks if the ModelCloudComplianceControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudComplianceControl{}

// ModelCloudComplianceControl struct for ModelCloudComplianceControl
type ModelCloudComplianceControl struct {
	Active *bool `json:"active,omitempty"`
	Category *string `json:"category,omitempty"`
	CategoryHierarchy []string `json:"category_hierarchy,omitempty"`
	CategoryHierarchyShort *string `json:"category_hierarchy_short,omitempty"`
	CloudProvider *string `json:"cloud_provider,omitempty"`
	ComplianceType *string `json:"compliance_type,omitempty"`
	ControlId *string `json:"control_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Executable *bool `json:"executable,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	ParentControlHierarchy []string `json:"parent_control_hierarchy,omitempty"`
	Service *string `json:"service,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewModelCloudComplianceControl instantiates a new ModelCloudComplianceControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudComplianceControl() *ModelCloudComplianceControl {
	this := ModelCloudComplianceControl{}
	return &this
}

// NewModelCloudComplianceControlWithDefaults instantiates a new ModelCloudComplianceControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudComplianceControlWithDefaults() *ModelCloudComplianceControl {
	this := ModelCloudComplianceControl{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ModelCloudComplianceControl) SetActive(v bool) {
	o.Active = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ModelCloudComplianceControl) SetCategory(v string) {
	o.Category = &v
}

// GetCategoryHierarchy returns the CategoryHierarchy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudComplianceControl) GetCategoryHierarchy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CategoryHierarchy
}

// GetCategoryHierarchyOk returns a tuple with the CategoryHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceControl) GetCategoryHierarchyOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoryHierarchy) {
		return nil, false
	}
	return o.CategoryHierarchy, true
}

// HasCategoryHierarchy returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasCategoryHierarchy() bool {
	if o != nil && !IsNil(o.CategoryHierarchy) {
		return true
	}

	return false
}

// SetCategoryHierarchy gets a reference to the given []string and assigns it to the CategoryHierarchy field.
func (o *ModelCloudComplianceControl) SetCategoryHierarchy(v []string) {
	o.CategoryHierarchy = v
}

// GetCategoryHierarchyShort returns the CategoryHierarchyShort field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetCategoryHierarchyShort() string {
	if o == nil || IsNil(o.CategoryHierarchyShort) {
		var ret string
		return ret
	}
	return *o.CategoryHierarchyShort
}

// GetCategoryHierarchyShortOk returns a tuple with the CategoryHierarchyShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetCategoryHierarchyShortOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryHierarchyShort) {
		return nil, false
	}
	return o.CategoryHierarchyShort, true
}

// HasCategoryHierarchyShort returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasCategoryHierarchyShort() bool {
	if o != nil && !IsNil(o.CategoryHierarchyShort) {
		return true
	}

	return false
}

// SetCategoryHierarchyShort gets a reference to the given string and assigns it to the CategoryHierarchyShort field.
func (o *ModelCloudComplianceControl) SetCategoryHierarchyShort(v string) {
	o.CategoryHierarchyShort = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ModelCloudComplianceControl) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetComplianceType returns the ComplianceType field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetComplianceType() string {
	if o == nil || IsNil(o.ComplianceType) {
		var ret string
		return ret
	}
	return *o.ComplianceType
}

// GetComplianceTypeOk returns a tuple with the ComplianceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetComplianceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComplianceType) {
		return nil, false
	}
	return o.ComplianceType, true
}

// HasComplianceType returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasComplianceType() bool {
	if o != nil && !IsNil(o.ComplianceType) {
		return true
	}

	return false
}

// SetComplianceType gets a reference to the given string and assigns it to the ComplianceType field.
func (o *ModelCloudComplianceControl) SetComplianceType(v string) {
	o.ComplianceType = &v
}

// GetControlId returns the ControlId field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetControlId() string {
	if o == nil || IsNil(o.ControlId) {
		var ret string
		return ret
	}
	return *o.ControlId
}

// GetControlIdOk returns a tuple with the ControlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetControlIdOk() (*string, bool) {
	if o == nil || IsNil(o.ControlId) {
		return nil, false
	}
	return o.ControlId, true
}

// HasControlId returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasControlId() bool {
	if o != nil && !IsNil(o.ControlId) {
		return true
	}

	return false
}

// SetControlId gets a reference to the given string and assigns it to the ControlId field.
func (o *ModelCloudComplianceControl) SetControlId(v string) {
	o.ControlId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelCloudComplianceControl) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ModelCloudComplianceControl) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *ModelCloudComplianceControl) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetExecutable returns the Executable field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetExecutable() bool {
	if o == nil || IsNil(o.Executable) {
		var ret bool
		return ret
	}
	return *o.Executable
}

// GetExecutableOk returns a tuple with the Executable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetExecutableOk() (*bool, bool) {
	if o == nil || IsNil(o.Executable) {
		return nil, false
	}
	return o.Executable, true
}

// HasExecutable returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasExecutable() bool {
	if o != nil && !IsNil(o.Executable) {
		return true
	}

	return false
}

// SetExecutable gets a reference to the given bool and assigns it to the Executable field.
func (o *ModelCloudComplianceControl) SetExecutable(v bool) {
	o.Executable = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ModelCloudComplianceControl) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParentControlHierarchy returns the ParentControlHierarchy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudComplianceControl) GetParentControlHierarchy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ParentControlHierarchy
}

// GetParentControlHierarchyOk returns a tuple with the ParentControlHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceControl) GetParentControlHierarchyOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentControlHierarchy) {
		return nil, false
	}
	return o.ParentControlHierarchy, true
}

// HasParentControlHierarchy returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasParentControlHierarchy() bool {
	if o != nil && !IsNil(o.ParentControlHierarchy) {
		return true
	}

	return false
}

// SetParentControlHierarchy gets a reference to the given []string and assigns it to the ParentControlHierarchy field.
func (o *ModelCloudComplianceControl) SetParentControlHierarchy(v []string) {
	o.ParentControlHierarchy = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ModelCloudComplianceControl) SetService(v string) {
	o.Service = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModelCloudComplianceControl) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceControl) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModelCloudComplianceControl) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ModelCloudComplianceControl) SetTitle(v string) {
	o.Title = &v
}

func (o ModelCloudComplianceControl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudComplianceControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if o.CategoryHierarchy != nil {
		toSerialize["category_hierarchy"] = o.CategoryHierarchy
	}
	if !IsNil(o.CategoryHierarchyShort) {
		toSerialize["category_hierarchy_short"] = o.CategoryHierarchyShort
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.ComplianceType) {
		toSerialize["compliance_type"] = o.ComplianceType
	}
	if !IsNil(o.ControlId) {
		toSerialize["control_id"] = o.ControlId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.Executable) {
		toSerialize["executable"] = o.Executable
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if o.ParentControlHierarchy != nil {
		toSerialize["parent_control_hierarchy"] = o.ParentControlHierarchy
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableModelCloudComplianceControl struct {
	value *ModelCloudComplianceControl
	isSet bool
}

func (v NullableModelCloudComplianceControl) Get() *ModelCloudComplianceControl {
	return v.value
}

func (v *NullableModelCloudComplianceControl) Set(val *ModelCloudComplianceControl) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudComplianceControl) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudComplianceControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudComplianceControl(val *ModelCloudComplianceControl) *NullableModelCloudComplianceControl {
	return &NullableModelCloudComplianceControl{value: val, isSet: true}
}

func (v NullableModelCloudComplianceControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudComplianceControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


