# ControlsQuarantineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** |  | 
**Policies** | [**[]ControlsPolicy**](ControlsPolicy.md) |  | 

## Methods

### NewControlsQuarantineConfig

`func NewControlsQuarantineConfig(nodeId string, policies []ControlsPolicy, ) *ControlsQuarantineConfig`

NewControlsQuarantineConfig instantiates a new ControlsQuarantineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsQuarantineConfigWithDefaults

`func NewControlsQuarantineConfigWithDefaults() *ControlsQuarantineConfig`

NewControlsQuarantineConfigWithDefaults instantiates a new ControlsQuarantineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *ControlsQuarantineConfig) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ControlsQuarantineConfig) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ControlsQuarantineConfig) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetPolicies

`func (o *ControlsQuarantineConfig) GetPolicies() []ControlsPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ControlsQuarantineConfig) GetPoliciesOk() (*[]ControlsPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ControlsQuarantineConfig) SetPolicies(v []ControlsPolicy)`

SetPolicies sets Policies field to given value.


### SetPoliciesNil

`func (o *ControlsQuarantineConfig) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *ControlsQuarantineConfig) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


