# ModelProcessAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | 
**Command** | **string** |  | 
**ContainerId** | **string** |  | 
**ContainerImage** | **string** |  | 
**ContainerIp** | **string** |  | 
**ContainerName** | **string** |  | 
**Count** | **int32** |  | 
**CpuTime** | **float32** |  | 
**CreatedAt** | **int32** |  | 
**EventType** | **string** |  | 
**ExecPath** | **string** |  | 
**Failure** | **string** |  | 
**Group** | **string** |  | 
**KubernetesClusterId** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**Masked** | **bool** |  | 
**Netstat** | **string** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**NumThreads** | **int32** |  | 
**Pid** | **int32** |  | 
**PodName** | **string** |  | 
**Priority** | **int32** |  | 
**ProcStatus** | **string** |  | 
**Return** | **int32** |  | 
**Rss** | **int32** |  | 
**RuleId** | **string** |  | 
**Session** | **int32** |  | 
**Severity** | **string** |  | 
**State** | **string** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Techniques** | **[]string** |  | 
**UpdatedAt** | **int32** |  | 
**User** | **string** |  | 
**UserName** | **string** |  | 
**Vsize** | **int32** |  | 

## Methods

### NewModelProcessAlert

`func NewModelProcessAlert(category string, command string, containerId string, containerImage string, containerIp string, containerName string, count int32, cpuTime float32, createdAt int32, eventType string, execPath string, failure string, group string, kubernetesClusterId string, kubernetesClusterName string, masked bool, netstat string, nodeId string, nodeType string, numThreads int32, pid int32, podName string, priority int32, procStatus string, return_ int32, rss int32, ruleId string, session int32, severity string, state string, summary string, tactics []string, techniques []string, updatedAt int32, user string, userName string, vsize int32, ) *ModelProcessAlert`

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


### GetCommand

`func (o *ModelProcessAlert) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ModelProcessAlert) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ModelProcessAlert) SetCommand(v string)`

SetCommand sets Command field to given value.


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


### GetCpuTime

`func (o *ModelProcessAlert) GetCpuTime() float32`

GetCpuTime returns the CpuTime field if non-nil, zero value otherwise.

### GetCpuTimeOk

`func (o *ModelProcessAlert) GetCpuTimeOk() (*float32, bool)`

GetCpuTimeOk returns a tuple with the CpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTime

`func (o *ModelProcessAlert) SetCpuTime(v float32)`

SetCpuTime sets CpuTime field to given value.


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


### GetExecPath

`func (o *ModelProcessAlert) GetExecPath() string`

GetExecPath returns the ExecPath field if non-nil, zero value otherwise.

### GetExecPathOk

`func (o *ModelProcessAlert) GetExecPathOk() (*string, bool)`

GetExecPathOk returns a tuple with the ExecPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecPath

`func (o *ModelProcessAlert) SetExecPath(v string)`

SetExecPath sets ExecPath field to given value.


### GetFailure

`func (o *ModelProcessAlert) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *ModelProcessAlert) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *ModelProcessAlert) SetFailure(v string)`

SetFailure sets Failure field to given value.


### GetGroup

`func (o *ModelProcessAlert) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ModelProcessAlert) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ModelProcessAlert) SetGroup(v string)`

SetGroup sets Group field to given value.


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


### GetNumThreads

`func (o *ModelProcessAlert) GetNumThreads() int32`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ModelProcessAlert) GetNumThreadsOk() (*int32, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ModelProcessAlert) SetNumThreads(v int32)`

SetNumThreads sets NumThreads field to given value.


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


### GetPriority

`func (o *ModelProcessAlert) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelProcessAlert) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelProcessAlert) SetPriority(v int32)`

SetPriority sets Priority field to given value.


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


### GetReturn

`func (o *ModelProcessAlert) GetReturn() int32`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *ModelProcessAlert) GetReturnOk() (*int32, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *ModelProcessAlert) SetReturn(v int32)`

SetReturn sets Return field to given value.


### GetRss

`func (o *ModelProcessAlert) GetRss() int32`

GetRss returns the Rss field if non-nil, zero value otherwise.

### GetRssOk

`func (o *ModelProcessAlert) GetRssOk() (*int32, bool)`

GetRssOk returns a tuple with the Rss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRss

`func (o *ModelProcessAlert) SetRss(v int32)`

SetRss sets Rss field to given value.


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


### GetSession

`func (o *ModelProcessAlert) GetSession() int32`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *ModelProcessAlert) GetSessionOk() (*int32, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *ModelProcessAlert) SetSession(v int32)`

SetSession sets Session field to given value.


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


### GetState

`func (o *ModelProcessAlert) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelProcessAlert) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelProcessAlert) SetState(v string)`

SetState sets State field to given value.


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


### GetUser

`func (o *ModelProcessAlert) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelProcessAlert) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelProcessAlert) SetUser(v string)`

SetUser sets User field to given value.


### GetUserName

`func (o *ModelProcessAlert) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ModelProcessAlert) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ModelProcessAlert) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetVsize

`func (o *ModelProcessAlert) GetVsize() int32`

GetVsize returns the Vsize field if non-nil, zero value otherwise.

### GetVsizeOk

`func (o *ModelProcessAlert) GetVsizeOk() (*int32, bool)`

GetVsizeOk returns a tuple with the Vsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsize

`func (o *ModelProcessAlert) SetVsize(v int32)`

SetVsize sets Vsize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


