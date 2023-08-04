# ModelNetworkViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ContainerId** | **string** |  | 
**CreatedAt** | **int32** |  | 
**Direction** | **string** |  | 
**ExecutedAt** | **string** |  | 
**HostName** | **string** |  | 
**NodeId** | **string** |  | 
**PodId** | **string** |  | 
**PolicyId** | **string** |  | 
**RemoteIp** | **string** |  | 
**RemotePort** | **int32** |  | 
**Severity** | **string** |  | 
**SourceIp** | **string** |  | 
**SourcePort** | **int32** |  | 

## Methods

### NewModelNetworkViolation

`func NewModelNetworkViolation(action string, containerId string, createdAt int32, direction string, executedAt string, hostName string, nodeId string, podId string, policyId string, remoteIp string, remotePort int32, severity string, sourceIp string, sourcePort int32, ) *ModelNetworkViolation`

NewModelNetworkViolation instantiates a new ModelNetworkViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkViolationWithDefaults

`func NewModelNetworkViolationWithDefaults() *ModelNetworkViolation`

NewModelNetworkViolationWithDefaults instantiates a new ModelNetworkViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ModelNetworkViolation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModelNetworkViolation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModelNetworkViolation) SetAction(v string)`

SetAction sets Action field to given value.


### GetContainerId

`func (o *ModelNetworkViolation) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelNetworkViolation) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelNetworkViolation) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetCreatedAt

`func (o *ModelNetworkViolation) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelNetworkViolation) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelNetworkViolation) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDirection

`func (o *ModelNetworkViolation) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModelNetworkViolation) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModelNetworkViolation) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetExecutedAt

`func (o *ModelNetworkViolation) GetExecutedAt() string`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *ModelNetworkViolation) GetExecutedAtOk() (*string, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *ModelNetworkViolation) SetExecutedAt(v string)`

SetExecutedAt sets ExecutedAt field to given value.


### GetHostName

`func (o *ModelNetworkViolation) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelNetworkViolation) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelNetworkViolation) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetNodeId

`func (o *ModelNetworkViolation) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelNetworkViolation) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelNetworkViolation) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetPodId

`func (o *ModelNetworkViolation) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *ModelNetworkViolation) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *ModelNetworkViolation) SetPodId(v string)`

SetPodId sets PodId field to given value.


### GetPolicyId

`func (o *ModelNetworkViolation) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ModelNetworkViolation) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ModelNetworkViolation) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetRemoteIp

`func (o *ModelNetworkViolation) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *ModelNetworkViolation) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *ModelNetworkViolation) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemotePort

`func (o *ModelNetworkViolation) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *ModelNetworkViolation) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *ModelNetworkViolation) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.


### GetSeverity

`func (o *ModelNetworkViolation) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelNetworkViolation) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelNetworkViolation) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSourceIp

`func (o *ModelNetworkViolation) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *ModelNetworkViolation) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *ModelNetworkViolation) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.


### GetSourcePort

`func (o *ModelNetworkViolation) GetSourcePort() int32`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *ModelNetworkViolation) GetSourcePortOk() (*int32, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *ModelNetworkViolation) SetSourcePort(v int32)`

SetSourcePort sets SourcePort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


