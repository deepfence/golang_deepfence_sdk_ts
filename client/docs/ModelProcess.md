# ModelProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCves** | **[]string** |  | 
**ActiveMalwares** | **[]string** |  | 
**ActiveSecrets** | **[]string** |  | 
**Cmdline** | **string** |  | 
**CpuMax** | **float32** |  | 
**CpuUsage** | **float32** |  | 
**HostName** | **string** |  | 
**InboundConnections** | [**[]ModelConnection**](ModelConnection.md) |  | 
**MemoryMax** | **int32** |  | 
**MemoryUsage** | **int32** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**OpenFiles** | **[]string** |  | 
**OpenFilesCount** | **int32** |  | 
**OutboundConnections** | [**[]ModelConnection**](ModelConnection.md) |  | 
**Pid** | **int32** |  | 
**Ppid** | **int32** |  | 
**ShortName** | **string** |  | 
**Threads** | **int32** |  | 

## Methods

### NewModelProcess

`func NewModelProcess(activeCves []string, activeMalwares []string, activeSecrets []string, cmdline string, cpuMax float32, cpuUsage float32, hostName string, inboundConnections []ModelConnection, memoryMax int32, memoryUsage int32, nodeId string, nodeName string, openFiles []string, openFilesCount int32, outboundConnections []ModelConnection, pid int32, ppid int32, shortName string, threads int32, ) *ModelProcess`

NewModelProcess instantiates a new ModelProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelProcessWithDefaults

`func NewModelProcessWithDefaults() *ModelProcess`

NewModelProcessWithDefaults instantiates a new ModelProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCves

`func (o *ModelProcess) GetActiveCves() []string`

GetActiveCves returns the ActiveCves field if non-nil, zero value otherwise.

### GetActiveCvesOk

`func (o *ModelProcess) GetActiveCvesOk() (*[]string, bool)`

GetActiveCvesOk returns a tuple with the ActiveCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCves

`func (o *ModelProcess) SetActiveCves(v []string)`

SetActiveCves sets ActiveCves field to given value.


### SetActiveCvesNil

`func (o *ModelProcess) SetActiveCvesNil(b bool)`

 SetActiveCvesNil sets the value for ActiveCves to be an explicit nil

### UnsetActiveCves
`func (o *ModelProcess) UnsetActiveCves()`

UnsetActiveCves ensures that no value is present for ActiveCves, not even an explicit nil
### GetActiveMalwares

`func (o *ModelProcess) GetActiveMalwares() []string`

GetActiveMalwares returns the ActiveMalwares field if non-nil, zero value otherwise.

### GetActiveMalwaresOk

`func (o *ModelProcess) GetActiveMalwaresOk() (*[]string, bool)`

GetActiveMalwaresOk returns a tuple with the ActiveMalwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMalwares

`func (o *ModelProcess) SetActiveMalwares(v []string)`

SetActiveMalwares sets ActiveMalwares field to given value.


### SetActiveMalwaresNil

`func (o *ModelProcess) SetActiveMalwaresNil(b bool)`

 SetActiveMalwaresNil sets the value for ActiveMalwares to be an explicit nil

### UnsetActiveMalwares
`func (o *ModelProcess) UnsetActiveMalwares()`

UnsetActiveMalwares ensures that no value is present for ActiveMalwares, not even an explicit nil
### GetActiveSecrets

`func (o *ModelProcess) GetActiveSecrets() []string`

GetActiveSecrets returns the ActiveSecrets field if non-nil, zero value otherwise.

### GetActiveSecretsOk

`func (o *ModelProcess) GetActiveSecretsOk() (*[]string, bool)`

GetActiveSecretsOk returns a tuple with the ActiveSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSecrets

`func (o *ModelProcess) SetActiveSecrets(v []string)`

SetActiveSecrets sets ActiveSecrets field to given value.


### SetActiveSecretsNil

`func (o *ModelProcess) SetActiveSecretsNil(b bool)`

 SetActiveSecretsNil sets the value for ActiveSecrets to be an explicit nil

### UnsetActiveSecrets
`func (o *ModelProcess) UnsetActiveSecrets()`

UnsetActiveSecrets ensures that no value is present for ActiveSecrets, not even an explicit nil
### GetCmdline

`func (o *ModelProcess) GetCmdline() string`

GetCmdline returns the Cmdline field if non-nil, zero value otherwise.

### GetCmdlineOk

`func (o *ModelProcess) GetCmdlineOk() (*string, bool)`

GetCmdlineOk returns a tuple with the Cmdline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdline

`func (o *ModelProcess) SetCmdline(v string)`

SetCmdline sets Cmdline field to given value.


### GetCpuMax

`func (o *ModelProcess) GetCpuMax() float32`

GetCpuMax returns the CpuMax field if non-nil, zero value otherwise.

### GetCpuMaxOk

`func (o *ModelProcess) GetCpuMaxOk() (*float32, bool)`

GetCpuMaxOk returns a tuple with the CpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMax

`func (o *ModelProcess) SetCpuMax(v float32)`

SetCpuMax sets CpuMax field to given value.


### GetCpuUsage

`func (o *ModelProcess) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ModelProcess) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ModelProcess) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.


### GetHostName

`func (o *ModelProcess) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelProcess) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelProcess) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetInboundConnections

`func (o *ModelProcess) GetInboundConnections() []ModelConnection`

GetInboundConnections returns the InboundConnections field if non-nil, zero value otherwise.

### GetInboundConnectionsOk

`func (o *ModelProcess) GetInboundConnectionsOk() (*[]ModelConnection, bool)`

GetInboundConnectionsOk returns a tuple with the InboundConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnections

`func (o *ModelProcess) SetInboundConnections(v []ModelConnection)`

SetInboundConnections sets InboundConnections field to given value.


### SetInboundConnectionsNil

`func (o *ModelProcess) SetInboundConnectionsNil(b bool)`

 SetInboundConnectionsNil sets the value for InboundConnections to be an explicit nil

### UnsetInboundConnections
`func (o *ModelProcess) UnsetInboundConnections()`

UnsetInboundConnections ensures that no value is present for InboundConnections, not even an explicit nil
### GetMemoryMax

`func (o *ModelProcess) GetMemoryMax() int32`

GetMemoryMax returns the MemoryMax field if non-nil, zero value otherwise.

### GetMemoryMaxOk

`func (o *ModelProcess) GetMemoryMaxOk() (*int32, bool)`

GetMemoryMaxOk returns a tuple with the MemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMax

`func (o *ModelProcess) SetMemoryMax(v int32)`

SetMemoryMax sets MemoryMax field to given value.


### GetMemoryUsage

`func (o *ModelProcess) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ModelProcess) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ModelProcess) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.


### GetNodeId

`func (o *ModelProcess) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelProcess) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelProcess) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelProcess) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelProcess) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelProcess) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetOpenFiles

`func (o *ModelProcess) GetOpenFiles() []string`

GetOpenFiles returns the OpenFiles field if non-nil, zero value otherwise.

### GetOpenFilesOk

`func (o *ModelProcess) GetOpenFilesOk() (*[]string, bool)`

GetOpenFilesOk returns a tuple with the OpenFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenFiles

`func (o *ModelProcess) SetOpenFiles(v []string)`

SetOpenFiles sets OpenFiles field to given value.


### SetOpenFilesNil

`func (o *ModelProcess) SetOpenFilesNil(b bool)`

 SetOpenFilesNil sets the value for OpenFiles to be an explicit nil

### UnsetOpenFiles
`func (o *ModelProcess) UnsetOpenFiles()`

UnsetOpenFiles ensures that no value is present for OpenFiles, not even an explicit nil
### GetOpenFilesCount

`func (o *ModelProcess) GetOpenFilesCount() int32`

GetOpenFilesCount returns the OpenFilesCount field if non-nil, zero value otherwise.

### GetOpenFilesCountOk

`func (o *ModelProcess) GetOpenFilesCountOk() (*int32, bool)`

GetOpenFilesCountOk returns a tuple with the OpenFilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenFilesCount

`func (o *ModelProcess) SetOpenFilesCount(v int32)`

SetOpenFilesCount sets OpenFilesCount field to given value.


### GetOutboundConnections

`func (o *ModelProcess) GetOutboundConnections() []ModelConnection`

GetOutboundConnections returns the OutboundConnections field if non-nil, zero value otherwise.

### GetOutboundConnectionsOk

`func (o *ModelProcess) GetOutboundConnectionsOk() (*[]ModelConnection, bool)`

GetOutboundConnectionsOk returns a tuple with the OutboundConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundConnections

`func (o *ModelProcess) SetOutboundConnections(v []ModelConnection)`

SetOutboundConnections sets OutboundConnections field to given value.


### SetOutboundConnectionsNil

`func (o *ModelProcess) SetOutboundConnectionsNil(b bool)`

 SetOutboundConnectionsNil sets the value for OutboundConnections to be an explicit nil

### UnsetOutboundConnections
`func (o *ModelProcess) UnsetOutboundConnections()`

UnsetOutboundConnections ensures that no value is present for OutboundConnections, not even an explicit nil
### GetPid

`func (o *ModelProcess) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ModelProcess) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ModelProcess) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetPpid

`func (o *ModelProcess) GetPpid() int32`

GetPpid returns the Ppid field if non-nil, zero value otherwise.

### GetPpidOk

`func (o *ModelProcess) GetPpidOk() (*int32, bool)`

GetPpidOk returns a tuple with the Ppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpid

`func (o *ModelProcess) SetPpid(v int32)`

SetPpid sets Ppid field to given value.


### GetShortName

`func (o *ModelProcess) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ModelProcess) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ModelProcess) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetThreads

`func (o *ModelProcess) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ModelProcess) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ModelProcess) SetThreads(v int32)`

SetThreads sets Threads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


