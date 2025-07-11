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

// checks if the ModelContainerImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContainerImage{}

// ModelContainerImage struct for ModelContainerImage
type ModelContainerImage struct {
	Containers []ModelContainer `json:"containers"`
	DockerImageCreatedAt string `json:"docker_image_created_at"`
	DockerImageId string `json:"docker_image_id"`
	DockerImageName string `json:"docker_image_name"`
	DockerImageSize string `json:"docker_image_size"`
	DockerImageTag string `json:"docker_image_tag"`
	DockerImageTagList []string `json:"docker_image_tag_list"`
	DockerImageVirtualSize string `json:"docker_image_virtual_size"`
	ImageNodeId string `json:"image_node_id"`
	IsDeepfenceSystem bool `json:"is_deepfence_system"`
	MalwareLatestScanId string `json:"malware_latest_scan_id"`
	MalwareScanStatus string `json:"malware_scan_status"`
	MalwaresCount int32 `json:"malwares_count"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	SecretLatestScanId string `json:"secret_latest_scan_id"`
	SecretScanStatus string `json:"secret_scan_status"`
	SecretsCount int32 `json:"secrets_count"`
	Tags []string `json:"tags"`
	VulnerabilitiesCount int32 `json:"vulnerabilities_count"`
	VulnerabilityLatestScanId string `json:"vulnerability_latest_scan_id"`
	VulnerabilityScanStatus string `json:"vulnerability_scan_status"`
}

type _ModelContainerImage ModelContainerImage

// NewModelContainerImage instantiates a new ModelContainerImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContainerImage(containers []ModelContainer, dockerImageCreatedAt string, dockerImageId string, dockerImageName string, dockerImageSize string, dockerImageTag string, dockerImageTagList []string, dockerImageVirtualSize string, imageNodeId string, isDeepfenceSystem bool, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, nodeId string, nodeName string, secretLatestScanId string, secretScanStatus string, secretsCount int32, tags []string, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string) *ModelContainerImage {
	this := ModelContainerImage{}
	this.Containers = containers
	this.DockerImageCreatedAt = dockerImageCreatedAt
	this.DockerImageId = dockerImageId
	this.DockerImageName = dockerImageName
	this.DockerImageSize = dockerImageSize
	this.DockerImageTag = dockerImageTag
	this.DockerImageTagList = dockerImageTagList
	this.DockerImageVirtualSize = dockerImageVirtualSize
	this.ImageNodeId = imageNodeId
	this.IsDeepfenceSystem = isDeepfenceSystem
	this.MalwareLatestScanId = malwareLatestScanId
	this.MalwareScanStatus = malwareScanStatus
	this.MalwaresCount = malwaresCount
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.SecretLatestScanId = secretLatestScanId
	this.SecretScanStatus = secretScanStatus
	this.SecretsCount = secretsCount
	this.Tags = tags
	this.VulnerabilitiesCount = vulnerabilitiesCount
	this.VulnerabilityLatestScanId = vulnerabilityLatestScanId
	this.VulnerabilityScanStatus = vulnerabilityScanStatus
	return &this
}

// NewModelContainerImageWithDefaults instantiates a new ModelContainerImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContainerImageWithDefaults() *ModelContainerImage {
	this := ModelContainerImage{}
	return &this
}

// GetContainers returns the Containers field value
// If the value is explicit nil, the zero value for []ModelContainer will be returned
func (o *ModelContainerImage) GetContainers() []ModelContainer {
	if o == nil {
		var ret []ModelContainer
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainerImage) GetContainersOk() ([]ModelContainer, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ModelContainerImage) SetContainers(v []ModelContainer) {
	o.Containers = v
}

// GetDockerImageCreatedAt returns the DockerImageCreatedAt field value
func (o *ModelContainerImage) GetDockerImageCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageCreatedAt
}

// GetDockerImageCreatedAtOk returns a tuple with the DockerImageCreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageCreatedAt, true
}

// SetDockerImageCreatedAt sets field value
func (o *ModelContainerImage) SetDockerImageCreatedAt(v string) {
	o.DockerImageCreatedAt = v
}

// GetDockerImageId returns the DockerImageId field value
func (o *ModelContainerImage) GetDockerImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageId
}

// GetDockerImageIdOk returns a tuple with the DockerImageId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageId, true
}

// SetDockerImageId sets field value
func (o *ModelContainerImage) SetDockerImageId(v string) {
	o.DockerImageId = v
}

// GetDockerImageName returns the DockerImageName field value
func (o *ModelContainerImage) GetDockerImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageName
}

// GetDockerImageNameOk returns a tuple with the DockerImageName field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageName, true
}

// SetDockerImageName sets field value
func (o *ModelContainerImage) SetDockerImageName(v string) {
	o.DockerImageName = v
}

// GetDockerImageSize returns the DockerImageSize field value
func (o *ModelContainerImage) GetDockerImageSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageSize
}

// GetDockerImageSizeOk returns a tuple with the DockerImageSize field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageSize, true
}

// SetDockerImageSize sets field value
func (o *ModelContainerImage) SetDockerImageSize(v string) {
	o.DockerImageSize = v
}

// GetDockerImageTag returns the DockerImageTag field value
func (o *ModelContainerImage) GetDockerImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageTag, true
}

// SetDockerImageTag sets field value
func (o *ModelContainerImage) SetDockerImageTag(v string) {
	o.DockerImageTag = v
}

// GetDockerImageTagList returns the DockerImageTagList field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelContainerImage) GetDockerImageTagList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DockerImageTagList
}

// GetDockerImageTagListOk returns a tuple with the DockerImageTagList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainerImage) GetDockerImageTagListOk() ([]string, bool) {
	if o == nil || IsNil(o.DockerImageTagList) {
		return nil, false
	}
	return o.DockerImageTagList, true
}

// SetDockerImageTagList sets field value
func (o *ModelContainerImage) SetDockerImageTagList(v []string) {
	o.DockerImageTagList = v
}

// GetDockerImageVirtualSize returns the DockerImageVirtualSize field value
func (o *ModelContainerImage) GetDockerImageVirtualSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageVirtualSize
}

// GetDockerImageVirtualSizeOk returns a tuple with the DockerImageVirtualSize field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageVirtualSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageVirtualSize, true
}

// SetDockerImageVirtualSize sets field value
func (o *ModelContainerImage) SetDockerImageVirtualSize(v string) {
	o.DockerImageVirtualSize = v
}

// GetImageNodeId returns the ImageNodeId field value
func (o *ModelContainerImage) GetImageNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageNodeId
}

// GetImageNodeIdOk returns a tuple with the ImageNodeId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetImageNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageNodeId, true
}

// SetImageNodeId sets field value
func (o *ModelContainerImage) SetImageNodeId(v string) {
	o.ImageNodeId = v
}

// GetIsDeepfenceSystem returns the IsDeepfenceSystem field value
func (o *ModelContainerImage) GetIsDeepfenceSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeepfenceSystem
}

// GetIsDeepfenceSystemOk returns a tuple with the IsDeepfenceSystem field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetIsDeepfenceSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeepfenceSystem, true
}

// SetIsDeepfenceSystem sets field value
func (o *ModelContainerImage) SetIsDeepfenceSystem(v bool) {
	o.IsDeepfenceSystem = v
}

// GetMalwareLatestScanId returns the MalwareLatestScanId field value
func (o *ModelContainerImage) GetMalwareLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareLatestScanId
}

// GetMalwareLatestScanIdOk returns a tuple with the MalwareLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMalwareLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareLatestScanId, true
}

// SetMalwareLatestScanId sets field value
func (o *ModelContainerImage) SetMalwareLatestScanId(v string) {
	o.MalwareLatestScanId = v
}

// GetMalwareScanStatus returns the MalwareScanStatus field value
func (o *ModelContainerImage) GetMalwareScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareScanStatus
}

// GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMalwareScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareScanStatus, true
}

// SetMalwareScanStatus sets field value
func (o *ModelContainerImage) SetMalwareScanStatus(v string) {
	o.MalwareScanStatus = v
}

// GetMalwaresCount returns the MalwaresCount field value
func (o *ModelContainerImage) GetMalwaresCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MalwaresCount
}

// GetMalwaresCountOk returns a tuple with the MalwaresCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMalwaresCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwaresCount, true
}

// SetMalwaresCount sets field value
func (o *ModelContainerImage) SetMalwaresCount(v int32) {
	o.MalwaresCount = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContainerImage) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainerImage) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModelContainerImage) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ModelContainerImage) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNodeId returns the NodeId field value
func (o *ModelContainerImage) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelContainerImage) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelContainerImage) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelContainerImage) SetNodeName(v string) {
	o.NodeName = v
}

// GetSecretLatestScanId returns the SecretLatestScanId field value
func (o *ModelContainerImage) GetSecretLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretLatestScanId
}

// GetSecretLatestScanIdOk returns a tuple with the SecretLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetSecretLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretLatestScanId, true
}

// SetSecretLatestScanId sets field value
func (o *ModelContainerImage) SetSecretLatestScanId(v string) {
	o.SecretLatestScanId = v
}

// GetSecretScanStatus returns the SecretScanStatus field value
func (o *ModelContainerImage) GetSecretScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretScanStatus
}

// GetSecretScanStatusOk returns a tuple with the SecretScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetSecretScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretScanStatus, true
}

// SetSecretScanStatus sets field value
func (o *ModelContainerImage) SetSecretScanStatus(v string) {
	o.SecretScanStatus = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *ModelContainerImage) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *ModelContainerImage) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetTags returns the Tags field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelContainerImage) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainerImage) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ModelContainerImage) SetTags(v []string) {
	o.Tags = v
}

// GetVulnerabilitiesCount returns the VulnerabilitiesCount field value
func (o *ModelContainerImage) GetVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilitiesCount
}

// GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilitiesCount, true
}

// SetVulnerabilitiesCount sets field value
func (o *ModelContainerImage) SetVulnerabilitiesCount(v int32) {
	o.VulnerabilitiesCount = v
}

// GetVulnerabilityLatestScanId returns the VulnerabilityLatestScanId field value
func (o *ModelContainerImage) GetVulnerabilityLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityLatestScanId
}

// GetVulnerabilityLatestScanIdOk returns a tuple with the VulnerabilityLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetVulnerabilityLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityLatestScanId, true
}

// SetVulnerabilityLatestScanId sets field value
func (o *ModelContainerImage) SetVulnerabilityLatestScanId(v string) {
	o.VulnerabilityLatestScanId = v
}

// GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field value
func (o *ModelContainerImage) GetVulnerabilityScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityScanStatus
}

// GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetVulnerabilityScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityScanStatus, true
}

// SetVulnerabilityScanStatus sets field value
func (o *ModelContainerImage) SetVulnerabilityScanStatus(v string) {
	o.VulnerabilityScanStatus = v
}

func (o ModelContainerImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContainerImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	toSerialize["docker_image_created_at"] = o.DockerImageCreatedAt
	toSerialize["docker_image_id"] = o.DockerImageId
	toSerialize["docker_image_name"] = o.DockerImageName
	toSerialize["docker_image_size"] = o.DockerImageSize
	toSerialize["docker_image_tag"] = o.DockerImageTag
	if o.DockerImageTagList != nil {
		toSerialize["docker_image_tag_list"] = o.DockerImageTagList
	}
	toSerialize["docker_image_virtual_size"] = o.DockerImageVirtualSize
	toSerialize["image_node_id"] = o.ImageNodeId
	toSerialize["is_deepfence_system"] = o.IsDeepfenceSystem
	toSerialize["malware_latest_scan_id"] = o.MalwareLatestScanId
	toSerialize["malware_scan_status"] = o.MalwareScanStatus
	toSerialize["malwares_count"] = o.MalwaresCount
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["secret_latest_scan_id"] = o.SecretLatestScanId
	toSerialize["secret_scan_status"] = o.SecretScanStatus
	toSerialize["secrets_count"] = o.SecretsCount
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["vulnerabilities_count"] = o.VulnerabilitiesCount
	toSerialize["vulnerability_latest_scan_id"] = o.VulnerabilityLatestScanId
	toSerialize["vulnerability_scan_status"] = o.VulnerabilityScanStatus
	return toSerialize, nil
}

func (o *ModelContainerImage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"containers",
		"docker_image_created_at",
		"docker_image_id",
		"docker_image_name",
		"docker_image_size",
		"docker_image_tag",
		"docker_image_tag_list",
		"docker_image_virtual_size",
		"image_node_id",
		"is_deepfence_system",
		"malware_latest_scan_id",
		"malware_scan_status",
		"malwares_count",
		"node_id",
		"node_name",
		"secret_latest_scan_id",
		"secret_scan_status",
		"secrets_count",
		"tags",
		"vulnerabilities_count",
		"vulnerability_latest_scan_id",
		"vulnerability_scan_status",
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

	varModelContainerImage := _ModelContainerImage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelContainerImage)

	if err != nil {
		return err
	}

	*o = ModelContainerImage(varModelContainerImage)

	return err
}

type NullableModelContainerImage struct {
	value *ModelContainerImage
	isSet bool
}

func (v NullableModelContainerImage) Get() *ModelContainerImage {
	return v.value
}

func (v *NullableModelContainerImage) Set(val *ModelContainerImage) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContainerImage) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContainerImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContainerImage(val *ModelContainerImage) *NullableModelContainerImage {
	return &NullableModelContainerImage{value: val, isSet: true}
}

func (v NullableModelContainerImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContainerImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


