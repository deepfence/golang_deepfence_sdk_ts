# ModelProcessAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | 
**ContainerId** | **string** |  | 
**ContainerImage** | **string** |  | 
**ContainerIp** | **string** |  | 
**ContainerName** | **string** |  | 
**Count** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**EventType** | **string** |  | 
**Filepath** | **string** |  | 
**Fstat** | **string** |  | 
**HostName** | **string** |  | 
**KubernetesClusterId** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**Masked** | **bool** |  | 
**Netstat** | **string** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**Pid** | **int32** |  | 
**PodName** | **string** |  | 
**ProcStatus** | **string** |  | 
**ProcessName** | **string** |  | 
**ResourceType** | **string** |  | 
**RuleId** | **string** |  | 
**Severity** | **string** |  | 
**SeverityScore** | **float32** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Techniques** | **[]string** |  | 
**Top** | **string** |  | 
**UpdatedAt** | **int32** |  | 
**Users** | **string** |  | 
**W** | **int32** |  | 

## Methods

### NewModelProcessAlert

`func NewModelProcessAlert(category string, containerId string, containerImage string, containerIp string, containerName string, count int32, createdAt int32, eventType string, filepath string, fstat string, hostName string, kubernetesClusterId string, kubernetesClusterName string, masked bool, netstat string, nodeId string, nodeType string, pid int32, podName string, procStatus string, processName string, resourceType string, ruleId string, severity string, severityScore float32, summary string, tactics []string, techniques []string, top string, updatedAt int32, users string, w int32, ) *ModelProcessAlert`

NewModelProcessAlert instantiates a new ModelProcessAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelProcessAlertWithDefaults

`func NewModelProcessAlertWithDefaults() *ModelProcessAlert`

NewModelProcessAlertWithDefaults instantiates a new ModelProcessAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ModelProcessAlert) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelProcessAlert) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelProcessAlert) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetContainerId

`func (o *ModelProcessAlert) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelProcessAlert) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelProcessAlert) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetContainerImage

`func (o *ModelProcessAlert) GetContainerImage() string`

GetContainerImage returns the ContainerImage field if non-nil, zero value otherwise.

### GetContainerImageOk

`func (o *ModelProcessAlert) GetContainerImageOk() (*string, bool)`

GetContainerImageOk returns a tuple with the ContainerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImage

`func (o *ModelProcessAlert) SetContainerImage(v string)`

SetContainerImage sets ContainerImage field to given value.


### GetContainerIp

`func (o *ModelProcessAlert) GetContainerIp() string`

GetContainerIp returns the ContainerIp field if non-nil, zero value otherwise.

### GetContainerIpOk

`func (o *ModelProcessAlert) GetContainerIpOk() (*string, bool)`

GetContainerIpOk returns a tuple with the ContainerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIp

`func (o *ModelProcessAlert) SetContainerIp(v string)`

SetContainerIp sets ContainerIp field to given value.


### GetContainerName

`func (o *ModelProcessAlert) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ModelProcessAlert) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ModelProcessAlert) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetCount

`func (o *ModelProcessAlert) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelProcessAlert) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelProcessAlert) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCreatedAt

`func (o *ModelProcessAlert) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelProcessAlert) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelProcessAlert) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetEventType

`func (o *ModelProcessAlert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelProcessAlert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelProcessAlert) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetFilepath

`func (o *ModelProcessAlert) GetFilepath() string`

GetFilepath returns the Filepath field if non-nil, zero value otherwise.

### GetFilepathOk

`func (o *ModelProcessAlert) GetFilepathOk() (*string, bool)`

GetFilepathOk returns a tuple with the Filepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilepath

`func (o *ModelProcessAlert) SetFilepath(v string)`

SetFilepath sets Filepath field to given value.


### GetFstat

`func (o *ModelProcessAlert) GetFstat() string`

GetFstat returns the Fstat field if non-nil, zero value otherwise.

### GetFstatOk

`func (o *ModelProcessAlert) GetFstatOk() (*string, bool)`

GetFstatOk returns a tuple with the Fstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstat

`func (o *ModelProcessAlert) SetFstat(v string)`

SetFstat sets Fstat field to given value.


### GetHostName

`func (o *ModelProcessAlert) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelProcessAlert) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelProcessAlert) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetKubernetesClusterId

`func (o *ModelProcessAlert) GetKubernetesClusterId() string`

GetKubernetesClusterId returns the KubernetesClusterId field if non-nil, zero value otherwise.

### GetKubernetesClusterIdOk

`func (o *ModelProcessAlert) GetKubernetesClusterIdOk() (*string, bool)`

GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterId

`func (o *ModelProcessAlert) SetKubernetesClusterId(v string)`

SetKubernetesClusterId sets KubernetesClusterId field to given value.


### GetKubernetesClusterName

`func (o *ModelProcessAlert) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ModelProcessAlert) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ModelProcessAlert) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.


### GetMasked

`func (o *ModelProcessAlert) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelProcessAlert) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelProcessAlert) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetNetstat

`func (o *ModelProcessAlert) GetNetstat() string`

GetNetstat returns the Netstat field if non-nil, zero value otherwise.

### GetNetstatOk

`func (o *ModelProcessAlert) GetNetstatOk() (*string, bool)`

GetNetstatOk returns a tuple with the Netstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetstat

`func (o *ModelProcessAlert) SetNetstat(v string)`

SetNetstat sets Netstat field to given value.


### GetNodeId

`func (o *ModelProcessAlert) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelProcessAlert) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelProcessAlert) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeType

`func (o *ModelProcessAlert) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelProcessAlert) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelProcessAlert) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetPid

`func (o *ModelProcessAlert) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ModelProcessAlert) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ModelProcessAlert) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetPodName

`func (o *ModelProcessAlert) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ModelProcessAlert) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ModelProcessAlert) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetProcStatus

`func (o *ModelProcessAlert) GetProcStatus() string`

GetProcStatus returns the ProcStatus field if non-nil, zero value otherwise.

### GetProcStatusOk

`func (o *ModelProcessAlert) GetProcStatusOk() (*string, bool)`

GetProcStatusOk returns a tuple with the ProcStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcStatus

`func (o *ModelProcessAlert) SetProcStatus(v string)`

SetProcStatus sets ProcStatus field to given value.


### GetProcessName

`func (o *ModelProcessAlert) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *ModelProcessAlert) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *ModelProcessAlert) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.


### GetResourceType

`func (o *ModelProcessAlert) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelProcessAlert) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelProcessAlert) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetRuleId

`func (o *ModelProcessAlert) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelProcessAlert) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelProcessAlert) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetSeverity

`func (o *ModelProcessAlert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelProcessAlert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelProcessAlert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSeverityScore

`func (o *ModelProcessAlert) GetSeverityScore() float32`

GetSeverityScore returns the SeverityScore field if non-nil, zero value otherwise.

### GetSeverityScoreOk

`func (o *ModelProcessAlert) GetSeverityScoreOk() (*float32, bool)`

GetSeverityScoreOk returns a tuple with the SeverityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityScore

`func (o *ModelProcessAlert) SetSeverityScore(v float32)`

SetSeverityScore sets SeverityScore field to given value.


### GetSummary

`func (o *ModelProcessAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelProcessAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelProcessAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTactics

`func (o *ModelProcessAlert) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelProcessAlert) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelProcessAlert) SetTactics(v []string)`

SetTactics sets Tactics field to given value.


### SetTacticsNil

`func (o *ModelProcessAlert) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelProcessAlert) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
### GetTechniques

`func (o *ModelProcessAlert) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelProcessAlert) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelProcessAlert) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.


### SetTechniquesNil

`func (o *ModelProcessAlert) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelProcessAlert) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil
### GetTop

`func (o *ModelProcessAlert) GetTop() string`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *ModelProcessAlert) GetTopOk() (*string, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *ModelProcessAlert) SetTop(v string)`

SetTop sets Top field to given value.


### GetUpdatedAt

`func (o *ModelProcessAlert) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelProcessAlert) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelProcessAlert) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUsers

`func (o *ModelProcessAlert) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModelProcessAlert) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModelProcessAlert) SetUsers(v string)`

SetUsers sets Users field to given value.


### GetW

`func (o *ModelProcessAlert) GetW() int32`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *ModelProcessAlert) GetWOk() (*int32, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *ModelProcessAlert) SetW(v int32)`

SetW sets W field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


