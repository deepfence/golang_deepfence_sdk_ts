# ModelQuarantineViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**AlertId** | **string** |  | 
**ConfigId** | **string** |  | 
**ContainerId** | **string** |  | 
**CreatedAt** | **int32** |  | 
**ExecutedAt** | **int32** |  | 
**HostName** | **string** |  | 
**NodeId** | **string** |  | 
**PodId** | **string** |  | 
**PolicyIndex** | **int32** |  | 
**Severity** | **string** |  | 
**Ttl** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewModelQuarantineViolation

`func NewModelQuarantineViolation(action string, alertId string, configId string, containerId string, createdAt int32, executedAt int32, hostName string, nodeId string, podId string, policyIndex int32, severity string, ttl int32, type_ string, ) *ModelQuarantineViolation`

NewModelQuarantineViolation instantiates a new ModelQuarantineViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelQuarantineViolationWithDefaults

`func NewModelQuarantineViolationWithDefaults() *ModelQuarantineViolation`

NewModelQuarantineViolationWithDefaults instantiates a new ModelQuarantineViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ModelQuarantineViolation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModelQuarantineViolation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModelQuarantineViolation) SetAction(v string)`

SetAction sets Action field to given value.


### GetAlertId

`func (o *ModelQuarantineViolation) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *ModelQuarantineViolation) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *ModelQuarantineViolation) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.


### GetConfigId

`func (o *ModelQuarantineViolation) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ModelQuarantineViolation) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ModelQuarantineViolation) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetContainerId

`func (o *ModelQuarantineViolation) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelQuarantineViolation) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelQuarantineViolation) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetCreatedAt

`func (o *ModelQuarantineViolation) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelQuarantineViolation) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelQuarantineViolation) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExecutedAt

`func (o *ModelQuarantineViolation) GetExecutedAt() int32`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *ModelQuarantineViolation) GetExecutedAtOk() (*int32, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *ModelQuarantineViolation) SetExecutedAt(v int32)`

SetExecutedAt sets ExecutedAt field to given value.


### GetHostName

`func (o *ModelQuarantineViolation) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelQuarantineViolation) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelQuarantineViolation) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetNodeId

`func (o *ModelQuarantineViolation) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelQuarantineViolation) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelQuarantineViolation) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetPodId

`func (o *ModelQuarantineViolation) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *ModelQuarantineViolation) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *ModelQuarantineViolation) SetPodId(v string)`

SetPodId sets PodId field to given value.


### GetPolicyIndex

`func (o *ModelQuarantineViolation) GetPolicyIndex() int32`

GetPolicyIndex returns the PolicyIndex field if non-nil, zero value otherwise.

### GetPolicyIndexOk

`func (o *ModelQuarantineViolation) GetPolicyIndexOk() (*int32, bool)`

GetPolicyIndexOk returns a tuple with the PolicyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIndex

`func (o *ModelQuarantineViolation) SetPolicyIndex(v int32)`

SetPolicyIndex sets PolicyIndex field to given value.


### GetSeverity

`func (o *ModelQuarantineViolation) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelQuarantineViolation) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelQuarantineViolation) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetTtl

`func (o *ModelQuarantineViolation) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ModelQuarantineViolation) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ModelQuarantineViolation) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetType

`func (o *ModelQuarantineViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelQuarantineViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelQuarantineViolation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


