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

// checks if the ModelCloudComplianceScanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudComplianceScanResult{}

// ModelCloudComplianceScanResult struct for ModelCloudComplianceScanResult
type ModelCloudComplianceScanResult struct {
	BenchmarkType []string `json:"benchmark_type"`
	CloudAccountId string `json:"cloud_account_id"`
	CompliancePercentage float32 `json:"compliance_percentage"`
	Compliances []ModelCloudCompliance `json:"compliances"`
	CreatedAt int64 `json:"created_at"`
	DockerContainerName string `json:"docker_container_name"`
	DockerImageName string `json:"docker_image_name"`
	HostName string `json:"host_name"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	NodeType string `json:"node_type"`
	ScanId string `json:"scan_id"`
	StatusCounts map[string]int32 `json:"status_counts"`
	UpdatedAt int64 `json:"updated_at"`
}

type _ModelCloudComplianceScanResult ModelCloudComplianceScanResult

// NewModelCloudComplianceScanResult instantiates a new ModelCloudComplianceScanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudComplianceScanResult(benchmarkType []string, cloudAccountId string, compliancePercentage float32, compliances []ModelCloudCompliance, createdAt int64, dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, scanId string, statusCounts map[string]int32, updatedAt int64) *ModelCloudComplianceScanResult {
	this := ModelCloudComplianceScanResult{}
	this.BenchmarkType = benchmarkType
	this.CloudAccountId = cloudAccountId
	this.CompliancePercentage = compliancePercentage
	this.Compliances = compliances
	this.CreatedAt = createdAt
	this.DockerContainerName = dockerContainerName
	this.DockerImageName = dockerImageName
	this.HostName = hostName
	this.KubernetesClusterName = kubernetesClusterName
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.NodeType = nodeType
	this.ScanId = scanId
	this.StatusCounts = statusCounts
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelCloudComplianceScanResultWithDefaults instantiates a new ModelCloudComplianceScanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudComplianceScanResultWithDefaults() *ModelCloudComplianceScanResult {
	this := ModelCloudComplianceScanResult{}
	return &this
}

// GetBenchmarkType returns the BenchmarkType field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelCloudComplianceScanResult) GetBenchmarkType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BenchmarkType
}

// GetBenchmarkTypeOk returns a tuple with the BenchmarkType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceScanResult) GetBenchmarkTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.BenchmarkType) {
		return nil, false
	}
	return o.BenchmarkType, true
}

// SetBenchmarkType sets field value
func (o *ModelCloudComplianceScanResult) SetBenchmarkType(v []string) {
	o.BenchmarkType = v
}

// GetCloudAccountId returns the CloudAccountId field value
func (o *ModelCloudComplianceScanResult) GetCloudAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetCloudAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudAccountId, true
}

// SetCloudAccountId sets field value
func (o *ModelCloudComplianceScanResult) SetCloudAccountId(v string) {
	o.CloudAccountId = v
}

// GetCompliancePercentage returns the CompliancePercentage field value
func (o *ModelCloudComplianceScanResult) GetCompliancePercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CompliancePercentage
}

// GetCompliancePercentageOk returns a tuple with the CompliancePercentage field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetCompliancePercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompliancePercentage, true
}

// SetCompliancePercentage sets field value
func (o *ModelCloudComplianceScanResult) SetCompliancePercentage(v float32) {
	o.CompliancePercentage = v
}

// GetCompliances returns the Compliances field value
// If the value is explicit nil, the zero value for []ModelCloudCompliance will be returned
func (o *ModelCloudComplianceScanResult) GetCompliances() []ModelCloudCompliance {
	if o == nil {
		var ret []ModelCloudCompliance
		return ret
	}

	return o.Compliances
}

// GetCompliancesOk returns a tuple with the Compliances field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceScanResult) GetCompliancesOk() ([]ModelCloudCompliance, bool) {
	if o == nil || IsNil(o.Compliances) {
		return nil, false
	}
	return o.Compliances, true
}

// SetCompliances sets field value
func (o *ModelCloudComplianceScanResult) SetCompliances(v []ModelCloudCompliance) {
	o.Compliances = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelCloudComplianceScanResult) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelCloudComplianceScanResult) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetDockerContainerName returns the DockerContainerName field value
func (o *ModelCloudComplianceScanResult) GetDockerContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerName
}

// GetDockerContainerNameOk returns a tuple with the DockerContainerName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetDockerContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerName, true
}

// SetDockerContainerName sets field value
func (o *ModelCloudComplianceScanResult) SetDockerContainerName(v string) {
	o.DockerContainerName = v
}

// GetDockerImageName returns the DockerImageName field value
func (o *ModelCloudComplianceScanResult) GetDockerImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageName
}

// GetDockerImageNameOk returns a tuple with the DockerImageName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetDockerImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageName, true
}

// SetDockerImageName sets field value
func (o *ModelCloudComplianceScanResult) SetDockerImageName(v string) {
	o.DockerImageName = v
}

// GetHostName returns the HostName field value
func (o *ModelCloudComplianceScanResult) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelCloudComplianceScanResult) SetHostName(v string) {
	o.HostName = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelCloudComplianceScanResult) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelCloudComplianceScanResult) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetNodeId returns the NodeId field value
func (o *ModelCloudComplianceScanResult) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelCloudComplianceScanResult) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelCloudComplianceScanResult) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelCloudComplianceScanResult) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelCloudComplianceScanResult) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelCloudComplianceScanResult) SetNodeType(v string) {
	o.NodeType = v
}

// GetScanId returns the ScanId field value
func (o *ModelCloudComplianceScanResult) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelCloudComplianceScanResult) SetScanId(v string) {
	o.ScanId = v
}

// GetStatusCounts returns the StatusCounts field value
// If the value is explicit nil, the zero value for map[string]int32 will be returned
func (o *ModelCloudComplianceScanResult) GetStatusCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.StatusCounts
}

// GetStatusCountsOk returns a tuple with the StatusCounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceScanResult) GetStatusCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.StatusCounts) {
		return nil, false
	}
	return &o.StatusCounts, true
}

// SetStatusCounts sets field value
func (o *ModelCloudComplianceScanResult) SetStatusCounts(v map[string]int32) {
	o.StatusCounts = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelCloudComplianceScanResult) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanResult) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelCloudComplianceScanResult) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o ModelCloudComplianceScanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudComplianceScanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BenchmarkType != nil {
		toSerialize["benchmark_type"] = o.BenchmarkType
	}
	toSerialize["cloud_account_id"] = o.CloudAccountId
	toSerialize["compliance_percentage"] = o.CompliancePercentage
	if o.Compliances != nil {
		toSerialize["compliances"] = o.Compliances
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["docker_container_name"] = o.DockerContainerName
	toSerialize["docker_image_name"] = o.DockerImageName
	toSerialize["host_name"] = o.HostName
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["node_type"] = o.NodeType
	toSerialize["scan_id"] = o.ScanId
	if o.StatusCounts != nil {
		toSerialize["status_counts"] = o.StatusCounts
	}
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelCloudComplianceScanResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"benchmark_type",
		"cloud_account_id",
		"compliance_percentage",
		"compliances",
		"created_at",
		"docker_container_name",
		"docker_image_name",
		"host_name",
		"kubernetes_cluster_name",
		"node_id",
		"node_name",
		"node_type",
		"scan_id",
		"status_counts",
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

	varModelCloudComplianceScanResult := _ModelCloudComplianceScanResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelCloudComplianceScanResult)

	if err != nil {
		return err
	}

	*o = ModelCloudComplianceScanResult(varModelCloudComplianceScanResult)

	return err
}

type NullableModelCloudComplianceScanResult struct {
	value *ModelCloudComplianceScanResult
	isSet bool
}

func (v NullableModelCloudComplianceScanResult) Get() *ModelCloudComplianceScanResult {
	return v.value
}

func (v *NullableModelCloudComplianceScanResult) Set(val *ModelCloudComplianceScanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudComplianceScanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudComplianceScanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudComplianceScanResult(val *ModelCloudComplianceScanResult) *NullableModelCloudComplianceScanResult {
	return &NullableModelCloudComplianceScanResult{value: val, isSet: true}
}

func (v NullableModelCloudComplianceScanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudComplianceScanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


