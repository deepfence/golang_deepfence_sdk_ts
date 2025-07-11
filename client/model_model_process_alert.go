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

// checks if the ModelProcessAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelProcessAlert{}

// ModelProcessAlert struct for ModelProcessAlert
type ModelProcessAlert struct {
	Category string `json:"category"`
	Command string `json:"command"`
	ContainerId string `json:"container_id"`
	ContainerImage string `json:"container_image"`
	ContainerIp string `json:"container_ip"`
	ContainerName string `json:"container_name"`
	Count int32 `json:"count"`
	CpuTime float32 `json:"cpu_time"`
	CreatedAt int32 `json:"created_at"`
	EventType string `json:"event_type"`
	ExecPath string `json:"exec_path"`
	Group string `json:"group"`
	KubernetesClusterId string `json:"kubernetes_cluster_id"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	Masked bool `json:"masked"`
	Netstat string `json:"netstat"`
	NodeId string `json:"node_id"`
	NodeType string `json:"node_type"`
	NumThreads int32 `json:"num_threads"`
	Pid int32 `json:"pid"`
	PodName string `json:"pod_name"`
	Priority int32 `json:"priority"`
	ProcStatus string `json:"proc_status"`
	Return int32 `json:"return"`
	Rss int32 `json:"rss"`
	RuleId string `json:"rule_id"`
	Session int32 `json:"session"`
	Severity string `json:"severity"`
	State string `json:"state"`
	Summary string `json:"summary"`
	Tactics []string `json:"tactics"`
	Techniques []string `json:"techniques"`
	Top string `json:"top"`
	UpdatedAt int32 `json:"updated_at"`
	User string `json:"user"`
	UserName string `json:"user_name"`
	Users string `json:"users"`
	Vsize int32 `json:"vsize"`
}

type _ModelProcessAlert ModelProcessAlert

// NewModelProcessAlert instantiates a new ModelProcessAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelProcessAlert(category string, command string, containerId string, containerImage string, containerIp string, containerName string, count int32, cpuTime float32, createdAt int32, eventType string, execPath string, group string, kubernetesClusterId string, kubernetesClusterName string, masked bool, netstat string, nodeId string, nodeType string, numThreads int32, pid int32, podName string, priority int32, procStatus string, return_ int32, rss int32, ruleId string, session int32, severity string, state string, summary string, tactics []string, techniques []string, top string, updatedAt int32, user string, userName string, users string, vsize int32) *ModelProcessAlert {
	this := ModelProcessAlert{}
	this.Category = category
	this.Command = command
	this.ContainerId = containerId
	this.ContainerImage = containerImage
	this.ContainerIp = containerIp
	this.ContainerName = containerName
	this.Count = count
	this.CpuTime = cpuTime
	this.CreatedAt = createdAt
	this.EventType = eventType
	this.ExecPath = execPath
	this.Group = group
	this.KubernetesClusterId = kubernetesClusterId
	this.KubernetesClusterName = kubernetesClusterName
	this.Masked = masked
	this.Netstat = netstat
	this.NodeId = nodeId
	this.NodeType = nodeType
	this.NumThreads = numThreads
	this.Pid = pid
	this.PodName = podName
	this.Priority = priority
	this.ProcStatus = procStatus
	this.Return = return_
	this.Rss = rss
	this.RuleId = ruleId
	this.Session = session
	this.Severity = severity
	this.State = state
	this.Summary = summary
	this.Tactics = tactics
	this.Techniques = techniques
	this.Top = top
	this.UpdatedAt = updatedAt
	this.User = user
	this.UserName = userName
	this.Users = users
	this.Vsize = vsize
	return &this
}

// NewModelProcessAlertWithDefaults instantiates a new ModelProcessAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelProcessAlertWithDefaults() *ModelProcessAlert {
	this := ModelProcessAlert{}
	return &this
}

// GetCategory returns the Category field value
func (o *ModelProcessAlert) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ModelProcessAlert) SetCategory(v string) {
	o.Category = v
}

// GetCommand returns the Command field value
func (o *ModelProcessAlert) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *ModelProcessAlert) SetCommand(v string) {
	o.Command = v
}

// GetContainerId returns the ContainerId field value
func (o *ModelProcessAlert) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *ModelProcessAlert) SetContainerId(v string) {
	o.ContainerId = v
}

// GetContainerImage returns the ContainerImage field value
func (o *ModelProcessAlert) GetContainerImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerImage
}

// GetContainerImageOk returns a tuple with the ContainerImage field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetContainerImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerImage, true
}

// SetContainerImage sets field value
func (o *ModelProcessAlert) SetContainerImage(v string) {
	o.ContainerImage = v
}

// GetContainerIp returns the ContainerIp field value
func (o *ModelProcessAlert) GetContainerIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerIp
}

// GetContainerIpOk returns a tuple with the ContainerIp field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetContainerIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerIp, true
}

// SetContainerIp sets field value
func (o *ModelProcessAlert) SetContainerIp(v string) {
	o.ContainerIp = v
}

// GetContainerName returns the ContainerName field value
func (o *ModelProcessAlert) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value
func (o *ModelProcessAlert) SetContainerName(v string) {
	o.ContainerName = v
}

// GetCount returns the Count field value
func (o *ModelProcessAlert) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ModelProcessAlert) SetCount(v int32) {
	o.Count = v
}

// GetCpuTime returns the CpuTime field value
func (o *ModelProcessAlert) GetCpuTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuTime
}

// GetCpuTimeOk returns a tuple with the CpuTime field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetCpuTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuTime, true
}

// SetCpuTime sets field value
func (o *ModelProcessAlert) SetCpuTime(v float32) {
	o.CpuTime = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelProcessAlert) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelProcessAlert) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetEventType returns the EventType field value
func (o *ModelProcessAlert) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ModelProcessAlert) SetEventType(v string) {
	o.EventType = v
}

// GetExecPath returns the ExecPath field value
func (o *ModelProcessAlert) GetExecPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecPath
}

// GetExecPathOk returns a tuple with the ExecPath field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetExecPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecPath, true
}

// SetExecPath sets field value
func (o *ModelProcessAlert) SetExecPath(v string) {
	o.ExecPath = v
}

// GetGroup returns the Group field value
func (o *ModelProcessAlert) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ModelProcessAlert) SetGroup(v string) {
	o.Group = v
}

// GetKubernetesClusterId returns the KubernetesClusterId field value
func (o *ModelProcessAlert) GetKubernetesClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterId
}

// GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetKubernetesClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterId, true
}

// SetKubernetesClusterId sets field value
func (o *ModelProcessAlert) SetKubernetesClusterId(v string) {
	o.KubernetesClusterId = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelProcessAlert) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelProcessAlert) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetMasked returns the Masked field value
func (o *ModelProcessAlert) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelProcessAlert) SetMasked(v bool) {
	o.Masked = v
}

// GetNetstat returns the Netstat field value
func (o *ModelProcessAlert) GetNetstat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Netstat
}

// GetNetstatOk returns a tuple with the Netstat field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetNetstatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Netstat, true
}

// SetNetstat sets field value
func (o *ModelProcessAlert) SetNetstat(v string) {
	o.Netstat = v
}

// GetNodeId returns the NodeId field value
func (o *ModelProcessAlert) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelProcessAlert) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeType returns the NodeType field value
func (o *ModelProcessAlert) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelProcessAlert) SetNodeType(v string) {
	o.NodeType = v
}

// GetNumThreads returns the NumThreads field value
func (o *ModelProcessAlert) GetNumThreads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumThreads
}

// GetNumThreadsOk returns a tuple with the NumThreads field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetNumThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumThreads, true
}

// SetNumThreads sets field value
func (o *ModelProcessAlert) SetNumThreads(v int32) {
	o.NumThreads = v
}

// GetPid returns the Pid field value
func (o *ModelProcessAlert) GetPid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetPidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value
func (o *ModelProcessAlert) SetPid(v int32) {
	o.Pid = v
}

// GetPodName returns the PodName field value
func (o *ModelProcessAlert) GetPodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetPodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodName, true
}

// SetPodName sets field value
func (o *ModelProcessAlert) SetPodName(v string) {
	o.PodName = v
}

// GetPriority returns the Priority field value
func (o *ModelProcessAlert) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ModelProcessAlert) SetPriority(v int32) {
	o.Priority = v
}

// GetProcStatus returns the ProcStatus field value
func (o *ModelProcessAlert) GetProcStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcStatus
}

// GetProcStatusOk returns a tuple with the ProcStatus field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetProcStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcStatus, true
}

// SetProcStatus sets field value
func (o *ModelProcessAlert) SetProcStatus(v string) {
	o.ProcStatus = v
}

// GetReturn returns the Return field value
func (o *ModelProcessAlert) GetReturn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Return
}

// GetReturnOk returns a tuple with the Return field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetReturnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Return, true
}

// SetReturn sets field value
func (o *ModelProcessAlert) SetReturn(v int32) {
	o.Return = v
}

// GetRss returns the Rss field value
func (o *ModelProcessAlert) GetRss() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rss
}

// GetRssOk returns a tuple with the Rss field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetRssOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rss, true
}

// SetRss sets field value
func (o *ModelProcessAlert) SetRss(v int32) {
	o.Rss = v
}

// GetRuleId returns the RuleId field value
func (o *ModelProcessAlert) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *ModelProcessAlert) SetRuleId(v string) {
	o.RuleId = v
}

// GetSession returns the Session field value
func (o *ModelProcessAlert) GetSession() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetSessionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *ModelProcessAlert) SetSession(v int32) {
	o.Session = v
}

// GetSeverity returns the Severity field value
func (o *ModelProcessAlert) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *ModelProcessAlert) SetSeverity(v string) {
	o.Severity = v
}

// GetState returns the State field value
func (o *ModelProcessAlert) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ModelProcessAlert) SetState(v string) {
	o.State = v
}

// GetSummary returns the Summary field value
func (o *ModelProcessAlert) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ModelProcessAlert) SetSummary(v string) {
	o.Summary = v
}

// GetTactics returns the Tactics field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelProcessAlert) GetTactics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tactics
}

// GetTacticsOk returns a tuple with the Tactics field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelProcessAlert) GetTacticsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tactics) {
		return nil, false
	}
	return o.Tactics, true
}

// SetTactics sets field value
func (o *ModelProcessAlert) SetTactics(v []string) {
	o.Tactics = v
}

// GetTechniques returns the Techniques field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelProcessAlert) GetTechniques() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Techniques
}

// GetTechniquesOk returns a tuple with the Techniques field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelProcessAlert) GetTechniquesOk() ([]string, bool) {
	if o == nil || IsNil(o.Techniques) {
		return nil, false
	}
	return o.Techniques, true
}

// SetTechniques sets field value
func (o *ModelProcessAlert) SetTechniques(v []string) {
	o.Techniques = v
}

// GetTop returns the Top field value
func (o *ModelProcessAlert) GetTop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Top
}

// GetTopOk returns a tuple with the Top field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetTopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Top, true
}

// SetTop sets field value
func (o *ModelProcessAlert) SetTop(v string) {
	o.Top = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelProcessAlert) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelProcessAlert) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetUser returns the User field value
func (o *ModelProcessAlert) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ModelProcessAlert) SetUser(v string) {
	o.User = v
}

// GetUserName returns the UserName field value
func (o *ModelProcessAlert) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *ModelProcessAlert) SetUserName(v string) {
	o.UserName = v
}

// GetUsers returns the Users field value
func (o *ModelProcessAlert) GetUsers() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetUsersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *ModelProcessAlert) SetUsers(v string) {
	o.Users = v
}

// GetVsize returns the Vsize field value
func (o *ModelProcessAlert) GetVsize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vsize
}

// GetVsizeOk returns a tuple with the Vsize field value
// and a boolean to check if the value has been set.
func (o *ModelProcessAlert) GetVsizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vsize, true
}

// SetVsize sets field value
func (o *ModelProcessAlert) SetVsize(v int32) {
	o.Vsize = v
}

func (o ModelProcessAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelProcessAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	toSerialize["command"] = o.Command
	toSerialize["container_id"] = o.ContainerId
	toSerialize["container_image"] = o.ContainerImage
	toSerialize["container_ip"] = o.ContainerIp
	toSerialize["container_name"] = o.ContainerName
	toSerialize["count"] = o.Count
	toSerialize["cpu_time"] = o.CpuTime
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["event_type"] = o.EventType
	toSerialize["exec_path"] = o.ExecPath
	toSerialize["group"] = o.Group
	toSerialize["kubernetes_cluster_id"] = o.KubernetesClusterId
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["masked"] = o.Masked
	toSerialize["netstat"] = o.Netstat
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_type"] = o.NodeType
	toSerialize["num_threads"] = o.NumThreads
	toSerialize["pid"] = o.Pid
	toSerialize["pod_name"] = o.PodName
	toSerialize["priority"] = o.Priority
	toSerialize["proc_status"] = o.ProcStatus
	toSerialize["return"] = o.Return
	toSerialize["rss"] = o.Rss
	toSerialize["rule_id"] = o.RuleId
	toSerialize["session"] = o.Session
	toSerialize["severity"] = o.Severity
	toSerialize["state"] = o.State
	toSerialize["summary"] = o.Summary
	if o.Tactics != nil {
		toSerialize["tactics"] = o.Tactics
	}
	if o.Techniques != nil {
		toSerialize["techniques"] = o.Techniques
	}
	toSerialize["top"] = o.Top
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["user"] = o.User
	toSerialize["user_name"] = o.UserName
	toSerialize["users"] = o.Users
	toSerialize["vsize"] = o.Vsize
	return toSerialize, nil
}

func (o *ModelProcessAlert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
		"command",
		"container_id",
		"container_image",
		"container_ip",
		"container_name",
		"count",
		"cpu_time",
		"created_at",
		"event_type",
		"exec_path",
		"group",
		"kubernetes_cluster_id",
		"kubernetes_cluster_name",
		"masked",
		"netstat",
		"node_id",
		"node_type",
		"num_threads",
		"pid",
		"pod_name",
		"priority",
		"proc_status",
		"return",
		"rss",
		"rule_id",
		"session",
		"severity",
		"state",
		"summary",
		"tactics",
		"techniques",
		"top",
		"updated_at",
		"user",
		"user_name",
		"users",
		"vsize",
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

	varModelProcessAlert := _ModelProcessAlert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelProcessAlert)

	if err != nil {
		return err
	}

	*o = ModelProcessAlert(varModelProcessAlert)

	return err
}

type NullableModelProcessAlert struct {
	value *ModelProcessAlert
	isSet bool
}

func (v NullableModelProcessAlert) Get() *ModelProcessAlert {
	return v.value
}

func (v *NullableModelProcessAlert) Set(val *ModelProcessAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableModelProcessAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableModelProcessAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelProcessAlert(val *ModelProcessAlert) *NullableModelProcessAlert {
	return &NullableModelProcessAlert{value: val, isSet: true}
}

func (v NullableModelProcessAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelProcessAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


