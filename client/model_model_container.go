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

// checks if the ModelContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContainer{}

// ModelContainer struct for ModelContainer
type ModelContainer struct {
	CpuMax float32 `json:"cpu_max"`
	CpuUsage float32 `json:"cpu_usage"`
	DockerContainerCommand string `json:"docker_container_command"`
	DockerContainerCreated string `json:"docker_container_created"`
	DockerContainerIps []interface{} `json:"docker_container_ips"`
	DockerContainerName string `json:"docker_container_name"`
	DockerContainerNetworkMode string `json:"docker_container_network_mode"`
	DockerContainerNetworks string `json:"docker_container_networks"`
	DockerContainerPorts string `json:"docker_container_ports"`
	DockerContainerState string `json:"docker_container_state"`
	DockerContainerStateHuman string `json:"docker_container_state_human"`
	DockerImageNameWithTag string `json:"docker_image_name_with_tag"`
	DockerLabels map[string]interface{} `json:"docker_labels"`
	HostName string `json:"host_name"`
	Image ModelContainerImage `json:"image"`
	IsDeepfenceSystem bool `json:"is_deepfence_system"`
	KubernetesClusterId string `json:"kubernetes_cluster_id"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	KubernetesNamespace string `json:"kubernetes_namespace"`
	MalwareLatestScanId string `json:"malware_latest_scan_id"`
	MalwareScanStatus string `json:"malware_scan_status"`
	MalwaresCount int32 `json:"malwares_count"`
	MemoryMax int32 `json:"memory_max"`
	MemoryUsage int32 `json:"memory_usage"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	Processes []ModelProcess `json:"processes"`
	SecretLatestScanId string `json:"secret_latest_scan_id"`
	SecretScanStatus string `json:"secret_scan_status"`
	SecretsCount int32 `json:"secrets_count"`
	Tags []string `json:"tags"`
	Uptime int32 `json:"uptime"`
	VulnerabilitiesCount int32 `json:"vulnerabilities_count"`
	VulnerabilityLatestScanId string `json:"vulnerability_latest_scan_id"`
	VulnerabilityScanStatus string `json:"vulnerability_scan_status"`
}

type _ModelContainer ModelContainer

// NewModelContainer instantiates a new ModelContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContainer(cpuMax float32, cpuUsage float32, dockerContainerCommand string, dockerContainerCreated string, dockerContainerIps []interface{}, dockerContainerName string, dockerContainerNetworkMode string, dockerContainerNetworks string, dockerContainerPorts string, dockerContainerState string, dockerContainerStateHuman string, dockerImageNameWithTag string, dockerLabels map[string]interface{}, hostName string, image ModelContainerImage, isDeepfenceSystem bool, kubernetesClusterId string, kubernetesClusterName string, kubernetesNamespace string, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, memoryMax int32, memoryUsage int32, nodeId string, nodeName string, processes []ModelProcess, secretLatestScanId string, secretScanStatus string, secretsCount int32, tags []string, uptime int32, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string) *ModelContainer {
	this := ModelContainer{}
	this.CpuMax = cpuMax
	this.CpuUsage = cpuUsage
	this.DockerContainerCommand = dockerContainerCommand
	this.DockerContainerCreated = dockerContainerCreated
	this.DockerContainerIps = dockerContainerIps
	this.DockerContainerName = dockerContainerName
	this.DockerContainerNetworkMode = dockerContainerNetworkMode
	this.DockerContainerNetworks = dockerContainerNetworks
	this.DockerContainerPorts = dockerContainerPorts
	this.DockerContainerState = dockerContainerState
	this.DockerContainerStateHuman = dockerContainerStateHuman
	this.DockerImageNameWithTag = dockerImageNameWithTag
	this.DockerLabels = dockerLabels
	this.HostName = hostName
	this.Image = image
	this.IsDeepfenceSystem = isDeepfenceSystem
	this.KubernetesClusterId = kubernetesClusterId
	this.KubernetesClusterName = kubernetesClusterName
	this.KubernetesNamespace = kubernetesNamespace
	this.MalwareLatestScanId = malwareLatestScanId
	this.MalwareScanStatus = malwareScanStatus
	this.MalwaresCount = malwaresCount
	this.MemoryMax = memoryMax
	this.MemoryUsage = memoryUsage
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.Processes = processes
	this.SecretLatestScanId = secretLatestScanId
	this.SecretScanStatus = secretScanStatus
	this.SecretsCount = secretsCount
	this.Tags = tags
	this.Uptime = uptime
	this.VulnerabilitiesCount = vulnerabilitiesCount
	this.VulnerabilityLatestScanId = vulnerabilityLatestScanId
	this.VulnerabilityScanStatus = vulnerabilityScanStatus
	return &this
}

// NewModelContainerWithDefaults instantiates a new ModelContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContainerWithDefaults() *ModelContainer {
	this := ModelContainer{}
	return &this
}

// GetCpuMax returns the CpuMax field value
func (o *ModelContainer) GetCpuMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuMax
}

// GetCpuMaxOk returns a tuple with the CpuMax field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetCpuMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMax, true
}

// SetCpuMax sets field value
func (o *ModelContainer) SetCpuMax(v float32) {
	o.CpuMax = v
}

// GetCpuUsage returns the CpuUsage field value
func (o *ModelContainer) GetCpuUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetCpuUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUsage, true
}

// SetCpuUsage sets field value
func (o *ModelContainer) SetCpuUsage(v float32) {
	o.CpuUsage = v
}

// GetDockerContainerCommand returns the DockerContainerCommand field value
func (o *ModelContainer) GetDockerContainerCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerCommand
}

// GetDockerContainerCommandOk returns a tuple with the DockerContainerCommand field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerCommand, true
}

// SetDockerContainerCommand sets field value
func (o *ModelContainer) SetDockerContainerCommand(v string) {
	o.DockerContainerCommand = v
}

// GetDockerContainerCreated returns the DockerContainerCreated field value
func (o *ModelContainer) GetDockerContainerCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerCreated
}

// GetDockerContainerCreatedOk returns a tuple with the DockerContainerCreated field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerCreated, true
}

// SetDockerContainerCreated sets field value
func (o *ModelContainer) SetDockerContainerCreated(v string) {
	o.DockerContainerCreated = v
}

// GetDockerContainerIps returns the DockerContainerIps field value
// If the value is explicit nil, the zero value for []interface{} will be returned
func (o *ModelContainer) GetDockerContainerIps() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.DockerContainerIps
}

// GetDockerContainerIpsOk returns a tuple with the DockerContainerIps field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainer) GetDockerContainerIpsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.DockerContainerIps) {
		return nil, false
	}
	return o.DockerContainerIps, true
}

// SetDockerContainerIps sets field value
func (o *ModelContainer) SetDockerContainerIps(v []interface{}) {
	o.DockerContainerIps = v
}

// GetDockerContainerName returns the DockerContainerName field value
func (o *ModelContainer) GetDockerContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerName
}

// GetDockerContainerNameOk returns a tuple with the DockerContainerName field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerName, true
}

// SetDockerContainerName sets field value
func (o *ModelContainer) SetDockerContainerName(v string) {
	o.DockerContainerName = v
}

// GetDockerContainerNetworkMode returns the DockerContainerNetworkMode field value
func (o *ModelContainer) GetDockerContainerNetworkMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerNetworkMode
}

// GetDockerContainerNetworkModeOk returns a tuple with the DockerContainerNetworkMode field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerNetworkModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerNetworkMode, true
}

// SetDockerContainerNetworkMode sets field value
func (o *ModelContainer) SetDockerContainerNetworkMode(v string) {
	o.DockerContainerNetworkMode = v
}

// GetDockerContainerNetworks returns the DockerContainerNetworks field value
func (o *ModelContainer) GetDockerContainerNetworks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerNetworks
}

// GetDockerContainerNetworksOk returns a tuple with the DockerContainerNetworks field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerNetworksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerNetworks, true
}

// SetDockerContainerNetworks sets field value
func (o *ModelContainer) SetDockerContainerNetworks(v string) {
	o.DockerContainerNetworks = v
}

// GetDockerContainerPorts returns the DockerContainerPorts field value
func (o *ModelContainer) GetDockerContainerPorts() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerPorts
}

// GetDockerContainerPortsOk returns a tuple with the DockerContainerPorts field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerPortsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerPorts, true
}

// SetDockerContainerPorts sets field value
func (o *ModelContainer) SetDockerContainerPorts(v string) {
	o.DockerContainerPorts = v
}

// GetDockerContainerState returns the DockerContainerState field value
func (o *ModelContainer) GetDockerContainerState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerState
}

// GetDockerContainerStateOk returns a tuple with the DockerContainerState field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerState, true
}

// SetDockerContainerState sets field value
func (o *ModelContainer) SetDockerContainerState(v string) {
	o.DockerContainerState = v
}

// GetDockerContainerStateHuman returns the DockerContainerStateHuman field value
func (o *ModelContainer) GetDockerContainerStateHuman() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerStateHuman
}

// GetDockerContainerStateHumanOk returns a tuple with the DockerContainerStateHuman field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerContainerStateHumanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerStateHuman, true
}

// SetDockerContainerStateHuman sets field value
func (o *ModelContainer) SetDockerContainerStateHuman(v string) {
	o.DockerContainerStateHuman = v
}

// GetDockerImageNameWithTag returns the DockerImageNameWithTag field value
func (o *ModelContainer) GetDockerImageNameWithTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageNameWithTag
}

// GetDockerImageNameWithTagOk returns a tuple with the DockerImageNameWithTag field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerImageNameWithTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageNameWithTag, true
}

// SetDockerImageNameWithTag sets field value
func (o *ModelContainer) SetDockerImageNameWithTag(v string) {
	o.DockerImageNameWithTag = v
}

// GetDockerLabels returns the DockerLabels field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ModelContainer) GetDockerLabels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DockerLabels
}

// GetDockerLabelsOk returns a tuple with the DockerLabels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainer) GetDockerLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DockerLabels) {
		return map[string]interface{}{}, false
	}
	return o.DockerLabels, true
}

// SetDockerLabels sets field value
func (o *ModelContainer) SetDockerLabels(v map[string]interface{}) {
	o.DockerLabels = v
}

// GetHostName returns the HostName field value
func (o *ModelContainer) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelContainer) SetHostName(v string) {
	o.HostName = v
}

// GetImage returns the Image field value
func (o *ModelContainer) GetImage() ModelContainerImage {
	if o == nil {
		var ret ModelContainerImage
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetImageOk() (*ModelContainerImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ModelContainer) SetImage(v ModelContainerImage) {
	o.Image = v
}

// GetIsDeepfenceSystem returns the IsDeepfenceSystem field value
func (o *ModelContainer) GetIsDeepfenceSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeepfenceSystem
}

// GetIsDeepfenceSystemOk returns a tuple with the IsDeepfenceSystem field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetIsDeepfenceSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeepfenceSystem, true
}

// SetIsDeepfenceSystem sets field value
func (o *ModelContainer) SetIsDeepfenceSystem(v bool) {
	o.IsDeepfenceSystem = v
}

// GetKubernetesClusterId returns the KubernetesClusterId field value
func (o *ModelContainer) GetKubernetesClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterId
}

// GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetKubernetesClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterId, true
}

// SetKubernetesClusterId sets field value
func (o *ModelContainer) SetKubernetesClusterId(v string) {
	o.KubernetesClusterId = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelContainer) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelContainer) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetKubernetesNamespace returns the KubernetesNamespace field value
func (o *ModelContainer) GetKubernetesNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesNamespace
}

// GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetKubernetesNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesNamespace, true
}

// SetKubernetesNamespace sets field value
func (o *ModelContainer) SetKubernetesNamespace(v string) {
	o.KubernetesNamespace = v
}

// GetMalwareLatestScanId returns the MalwareLatestScanId field value
func (o *ModelContainer) GetMalwareLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareLatestScanId
}

// GetMalwareLatestScanIdOk returns a tuple with the MalwareLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMalwareLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareLatestScanId, true
}

// SetMalwareLatestScanId sets field value
func (o *ModelContainer) SetMalwareLatestScanId(v string) {
	o.MalwareLatestScanId = v
}

// GetMalwareScanStatus returns the MalwareScanStatus field value
func (o *ModelContainer) GetMalwareScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareScanStatus
}

// GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMalwareScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareScanStatus, true
}

// SetMalwareScanStatus sets field value
func (o *ModelContainer) SetMalwareScanStatus(v string) {
	o.MalwareScanStatus = v
}

// GetMalwaresCount returns the MalwaresCount field value
func (o *ModelContainer) GetMalwaresCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MalwaresCount
}

// GetMalwaresCountOk returns a tuple with the MalwaresCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMalwaresCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwaresCount, true
}

// SetMalwaresCount sets field value
func (o *ModelContainer) SetMalwaresCount(v int32) {
	o.MalwaresCount = v
}

// GetMemoryMax returns the MemoryMax field value
func (o *ModelContainer) GetMemoryMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryMax
}

// GetMemoryMaxOk returns a tuple with the MemoryMax field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMemoryMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryMax, true
}

// SetMemoryMax sets field value
func (o *ModelContainer) SetMemoryMax(v int32) {
	o.MemoryMax = v
}

// GetMemoryUsage returns the MemoryUsage field value
func (o *ModelContainer) GetMemoryUsage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMemoryUsageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsage, true
}

// SetMemoryUsage sets field value
func (o *ModelContainer) SetMemoryUsage(v int32) {
	o.MemoryUsage = v
}

// GetNodeId returns the NodeId field value
func (o *ModelContainer) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelContainer) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelContainer) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelContainer) SetNodeName(v string) {
	o.NodeName = v
}

// GetProcesses returns the Processes field value
// If the value is explicit nil, the zero value for []ModelProcess will be returned
func (o *ModelContainer) GetProcesses() []ModelProcess {
	if o == nil {
		var ret []ModelProcess
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainer) GetProcessesOk() ([]ModelProcess, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *ModelContainer) SetProcesses(v []ModelProcess) {
	o.Processes = v
}

// GetSecretLatestScanId returns the SecretLatestScanId field value
func (o *ModelContainer) GetSecretLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretLatestScanId
}

// GetSecretLatestScanIdOk returns a tuple with the SecretLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetSecretLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretLatestScanId, true
}

// SetSecretLatestScanId sets field value
func (o *ModelContainer) SetSecretLatestScanId(v string) {
	o.SecretLatestScanId = v
}

// GetSecretScanStatus returns the SecretScanStatus field value
func (o *ModelContainer) GetSecretScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretScanStatus
}

// GetSecretScanStatusOk returns a tuple with the SecretScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetSecretScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretScanStatus, true
}

// SetSecretScanStatus sets field value
func (o *ModelContainer) SetSecretScanStatus(v string) {
	o.SecretScanStatus = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *ModelContainer) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *ModelContainer) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetTags returns the Tags field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelContainer) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainer) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ModelContainer) SetTags(v []string) {
	o.Tags = v
}

// GetUptime returns the Uptime field value
func (o *ModelContainer) GetUptime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetUptimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uptime, true
}

// SetUptime sets field value
func (o *ModelContainer) SetUptime(v int32) {
	o.Uptime = v
}

// GetVulnerabilitiesCount returns the VulnerabilitiesCount field value
func (o *ModelContainer) GetVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilitiesCount
}

// GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilitiesCount, true
}

// SetVulnerabilitiesCount sets field value
func (o *ModelContainer) SetVulnerabilitiesCount(v int32) {
	o.VulnerabilitiesCount = v
}

// GetVulnerabilityLatestScanId returns the VulnerabilityLatestScanId field value
func (o *ModelContainer) GetVulnerabilityLatestScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityLatestScanId
}

// GetVulnerabilityLatestScanIdOk returns a tuple with the VulnerabilityLatestScanId field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetVulnerabilityLatestScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityLatestScanId, true
}

// SetVulnerabilityLatestScanId sets field value
func (o *ModelContainer) SetVulnerabilityLatestScanId(v string) {
	o.VulnerabilityLatestScanId = v
}

// GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field value
func (o *ModelContainer) GetVulnerabilityScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityScanStatus
}

// GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetVulnerabilityScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityScanStatus, true
}

// SetVulnerabilityScanStatus sets field value
func (o *ModelContainer) SetVulnerabilityScanStatus(v string) {
	o.VulnerabilityScanStatus = v
}

func (o ModelContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu_max"] = o.CpuMax
	toSerialize["cpu_usage"] = o.CpuUsage
	toSerialize["docker_container_command"] = o.DockerContainerCommand
	toSerialize["docker_container_created"] = o.DockerContainerCreated
	if o.DockerContainerIps != nil {
		toSerialize["docker_container_ips"] = o.DockerContainerIps
	}
	toSerialize["docker_container_name"] = o.DockerContainerName
	toSerialize["docker_container_network_mode"] = o.DockerContainerNetworkMode
	toSerialize["docker_container_networks"] = o.DockerContainerNetworks
	toSerialize["docker_container_ports"] = o.DockerContainerPorts
	toSerialize["docker_container_state"] = o.DockerContainerState
	toSerialize["docker_container_state_human"] = o.DockerContainerStateHuman
	toSerialize["docker_image_name_with_tag"] = o.DockerImageNameWithTag
	if o.DockerLabels != nil {
		toSerialize["docker_labels"] = o.DockerLabels
	}
	toSerialize["host_name"] = o.HostName
	toSerialize["image"] = o.Image
	toSerialize["is_deepfence_system"] = o.IsDeepfenceSystem
	toSerialize["kubernetes_cluster_id"] = o.KubernetesClusterId
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["kubernetes_namespace"] = o.KubernetesNamespace
	toSerialize["malware_latest_scan_id"] = o.MalwareLatestScanId
	toSerialize["malware_scan_status"] = o.MalwareScanStatus
	toSerialize["malwares_count"] = o.MalwaresCount
	toSerialize["memory_max"] = o.MemoryMax
	toSerialize["memory_usage"] = o.MemoryUsage
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	toSerialize["secret_latest_scan_id"] = o.SecretLatestScanId
	toSerialize["secret_scan_status"] = o.SecretScanStatus
	toSerialize["secrets_count"] = o.SecretsCount
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["uptime"] = o.Uptime
	toSerialize["vulnerabilities_count"] = o.VulnerabilitiesCount
	toSerialize["vulnerability_latest_scan_id"] = o.VulnerabilityLatestScanId
	toSerialize["vulnerability_scan_status"] = o.VulnerabilityScanStatus
	return toSerialize, nil
}

func (o *ModelContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cpu_max",
		"cpu_usage",
		"docker_container_command",
		"docker_container_created",
		"docker_container_ips",
		"docker_container_name",
		"docker_container_network_mode",
		"docker_container_networks",
		"docker_container_ports",
		"docker_container_state",
		"docker_container_state_human",
		"docker_image_name_with_tag",
		"docker_labels",
		"host_name",
		"image",
		"is_deepfence_system",
		"kubernetes_cluster_id",
		"kubernetes_cluster_name",
		"kubernetes_namespace",
		"malware_latest_scan_id",
		"malware_scan_status",
		"malwares_count",
		"memory_max",
		"memory_usage",
		"node_id",
		"node_name",
		"processes",
		"secret_latest_scan_id",
		"secret_scan_status",
		"secrets_count",
		"tags",
		"uptime",
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

	varModelContainer := _ModelContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelContainer)

	if err != nil {
		return err
	}

	*o = ModelContainer(varModelContainer)

	return err
}

type NullableModelContainer struct {
	value *ModelContainer
	isSet bool
}

func (v NullableModelContainer) Get() *ModelContainer {
	return v.value
}

func (v *NullableModelContainer) Set(val *ModelContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContainer(val *ModelContainer) *NullableModelContainer {
	return &NullableModelContainer{value: val, isSet: true}
}

func (v NullableModelContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


