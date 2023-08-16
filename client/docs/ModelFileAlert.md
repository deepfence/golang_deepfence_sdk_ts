# ModelFileAlert

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
**Direction** | **string** |  | 
**EventType** | **string** |  | 
**Filepath** | **string** |  | 
**Fstat** | **string** |  | 
**HostName** | **string** |  | 
**Masked** | **bool** |  | 
**Netstat** | **string** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**Pid** | **int32** |  | 
**ProcStatus** | **string** |  | 
**ProcessName** | **string** |  | 
**ResourceType** | **string** |  | 
**Severity** | **string** |  | 
**SeverityScore** | **float32** |  | 
**SignatureId** | **int32** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Techniques** | **[]string** |  | 
**Top** | **string** |  | 
**UpdatedAt** | **int32** |  | 
**Users** | **string** |  | 
**W** | **int32** |  | 

## Methods

### NewModelFileAlert

`func NewModelFileAlert(category string, containerId string, containerImage string, containerIp string, containerName string, count int32, createdAt int32, direction string, eventType string, filepath string, fstat string, hostName string, masked bool, netstat string, nodeId string, nodeType string, pid int32, procStatus string, processName string, resourceType string, severity string, severityScore float32, signatureId int32, summary string, tactics []string, techniques []string, top string, updatedAt int32, users string, w int32, ) *ModelFileAlert`

NewModelFileAlert instantiates a new ModelFileAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFileAlertWithDefaults

`func NewModelFileAlertWithDefaults() *ModelFileAlert`

NewModelFileAlertWithDefaults instantiates a new ModelFileAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ModelFileAlert) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelFileAlert) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelFileAlert) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetContainerId

`func (o *ModelFileAlert) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelFileAlert) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelFileAlert) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetContainerImage

`func (o *ModelFileAlert) GetContainerImage() string`

GetContainerImage returns the ContainerImage field if non-nil, zero value otherwise.

### GetContainerImageOk

`func (o *ModelFileAlert) GetContainerImageOk() (*string, bool)`

GetContainerImageOk returns a tuple with the ContainerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImage

`func (o *ModelFileAlert) SetContainerImage(v string)`

SetContainerImage sets ContainerImage field to given value.


### GetContainerIp

`func (o *ModelFileAlert) GetContainerIp() string`

GetContainerIp returns the ContainerIp field if non-nil, zero value otherwise.

### GetContainerIpOk

`func (o *ModelFileAlert) GetContainerIpOk() (*string, bool)`

GetContainerIpOk returns a tuple with the ContainerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIp

`func (o *ModelFileAlert) SetContainerIp(v string)`

SetContainerIp sets ContainerIp field to given value.


### GetContainerName

`func (o *ModelFileAlert) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ModelFileAlert) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ModelFileAlert) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetCount

`func (o *ModelFileAlert) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelFileAlert) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelFileAlert) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCreatedAt

`func (o *ModelFileAlert) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelFileAlert) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelFileAlert) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDirection

`func (o *ModelFileAlert) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModelFileAlert) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModelFileAlert) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEventType

`func (o *ModelFileAlert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelFileAlert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelFileAlert) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetFilepath

`func (o *ModelFileAlert) GetFilepath() string`

GetFilepath returns the Filepath field if non-nil, zero value otherwise.

### GetFilepathOk

`func (o *ModelFileAlert) GetFilepathOk() (*string, bool)`

GetFilepathOk returns a tuple with the Filepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilepath

`func (o *ModelFileAlert) SetFilepath(v string)`

SetFilepath sets Filepath field to given value.


### GetFstat

`func (o *ModelFileAlert) GetFstat() string`

GetFstat returns the Fstat field if non-nil, zero value otherwise.

### GetFstatOk

`func (o *ModelFileAlert) GetFstatOk() (*string, bool)`

GetFstatOk returns a tuple with the Fstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstat

`func (o *ModelFileAlert) SetFstat(v string)`

SetFstat sets Fstat field to given value.


### GetHostName

`func (o *ModelFileAlert) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelFileAlert) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelFileAlert) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetMasked

`func (o *ModelFileAlert) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelFileAlert) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelFileAlert) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetNetstat

`func (o *ModelFileAlert) GetNetstat() string`

GetNetstat returns the Netstat field if non-nil, zero value otherwise.

### GetNetstatOk

`func (o *ModelFileAlert) GetNetstatOk() (*string, bool)`

GetNetstatOk returns a tuple with the Netstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetstat

`func (o *ModelFileAlert) SetNetstat(v string)`

SetNetstat sets Netstat field to given value.


### GetNodeId

`func (o *ModelFileAlert) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelFileAlert) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelFileAlert) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeType

`func (o *ModelFileAlert) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelFileAlert) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelFileAlert) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetPid

`func (o *ModelFileAlert) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ModelFileAlert) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ModelFileAlert) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetProcStatus

`func (o *ModelFileAlert) GetProcStatus() string`

GetProcStatus returns the ProcStatus field if non-nil, zero value otherwise.

### GetProcStatusOk

`func (o *ModelFileAlert) GetProcStatusOk() (*string, bool)`

GetProcStatusOk returns a tuple with the ProcStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcStatus

`func (o *ModelFileAlert) SetProcStatus(v string)`

SetProcStatus sets ProcStatus field to given value.


### GetProcessName

`func (o *ModelFileAlert) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *ModelFileAlert) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *ModelFileAlert) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.


### GetResourceType

`func (o *ModelFileAlert) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelFileAlert) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelFileAlert) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetSeverity

`func (o *ModelFileAlert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelFileAlert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelFileAlert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSeverityScore

`func (o *ModelFileAlert) GetSeverityScore() float32`

GetSeverityScore returns the SeverityScore field if non-nil, zero value otherwise.

### GetSeverityScoreOk

`func (o *ModelFileAlert) GetSeverityScoreOk() (*float32, bool)`

GetSeverityScoreOk returns a tuple with the SeverityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityScore

`func (o *ModelFileAlert) SetSeverityScore(v float32)`

SetSeverityScore sets SeverityScore field to given value.


### GetSignatureId

`func (o *ModelFileAlert) GetSignatureId() int32`

GetSignatureId returns the SignatureId field if non-nil, zero value otherwise.

### GetSignatureIdOk

`func (o *ModelFileAlert) GetSignatureIdOk() (*int32, bool)`

GetSignatureIdOk returns a tuple with the SignatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureId

`func (o *ModelFileAlert) SetSignatureId(v int32)`

SetSignatureId sets SignatureId field to given value.


### GetSummary

`func (o *ModelFileAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelFileAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelFileAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTactics

`func (o *ModelFileAlert) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelFileAlert) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelFileAlert) SetTactics(v []string)`

SetTactics sets Tactics field to given value.


### SetTacticsNil

`func (o *ModelFileAlert) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelFileAlert) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
### GetTechniques

`func (o *ModelFileAlert) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelFileAlert) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelFileAlert) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.


### SetTechniquesNil

`func (o *ModelFileAlert) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelFileAlert) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil
### GetTop

`func (o *ModelFileAlert) GetTop() string`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *ModelFileAlert) GetTopOk() (*string, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *ModelFileAlert) SetTop(v string)`

SetTop sets Top field to given value.


### GetUpdatedAt

`func (o *ModelFileAlert) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelFileAlert) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelFileAlert) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUsers

`func (o *ModelFileAlert) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModelFileAlert) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModelFileAlert) SetUsers(v string)`

SetUsers sets Users field to given value.


### GetW

`func (o *ModelFileAlert) GetW() int32`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *ModelFileAlert) GetWOk() (*int32, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *ModelFileAlert) SetW(v int32)`

SetW sets W field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


