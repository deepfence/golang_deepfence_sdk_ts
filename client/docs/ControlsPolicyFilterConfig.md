# ControlsPolicyFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoredRemoteIps** | **[]string** |  | 
**NodeId** | **string** |  | 
**Policies** | [**[]ControlsNetworkPolicy**](ControlsNetworkPolicy.md) |  | 
**UpdatedAt** | **int32** |  | 
**UseWaf** | **bool** |  | 

## Methods

### NewControlsPolicyFilterConfig

`func NewControlsPolicyFilterConfig(ignoredRemoteIps []string, nodeId string, policies []ControlsNetworkPolicy, updatedAt int32, useWaf bool, ) *ControlsPolicyFilterConfig`

NewControlsPolicyFilterConfig instantiates a new ControlsPolicyFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPolicyFilterConfigWithDefaults

`func NewControlsPolicyFilterConfigWithDefaults() *ControlsPolicyFilterConfig`

NewControlsPolicyFilterConfigWithDefaults instantiates a new ControlsPolicyFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoredRemoteIps

`func (o *ControlsPolicyFilterConfig) GetIgnoredRemoteIps() []string`

GetIgnoredRemoteIps returns the IgnoredRemoteIps field if non-nil, zero value otherwise.

### GetIgnoredRemoteIpsOk

`func (o *ControlsPolicyFilterConfig) GetIgnoredRemoteIpsOk() (*[]string, bool)`

GetIgnoredRemoteIpsOk returns a tuple with the IgnoredRemoteIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredRemoteIps

`func (o *ControlsPolicyFilterConfig) SetIgnoredRemoteIps(v []string)`

SetIgnoredRemoteIps sets IgnoredRemoteIps field to given value.


### SetIgnoredRemoteIpsNil

`func (o *ControlsPolicyFilterConfig) SetIgnoredRemoteIpsNil(b bool)`

 SetIgnoredRemoteIpsNil sets the value for IgnoredRemoteIps to be an explicit nil

### UnsetIgnoredRemoteIps
`func (o *ControlsPolicyFilterConfig) UnsetIgnoredRemoteIps()`

UnsetIgnoredRemoteIps ensures that no value is present for IgnoredRemoteIps, not even an explicit nil
### GetNodeId

`func (o *ControlsPolicyFilterConfig) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ControlsPolicyFilterConfig) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ControlsPolicyFilterConfig) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetPolicies

`func (o *ControlsPolicyFilterConfig) GetPolicies() []ControlsNetworkPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ControlsPolicyFilterConfig) GetPoliciesOk() (*[]ControlsNetworkPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ControlsPolicyFilterConfig) SetPolicies(v []ControlsNetworkPolicy)`

SetPolicies sets Policies field to given value.


### SetPoliciesNil

`func (o *ControlsPolicyFilterConfig) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *ControlsPolicyFilterConfig) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetUpdatedAt

`func (o *ControlsPolicyFilterConfig) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsPolicyFilterConfig) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsPolicyFilterConfig) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUseWaf

`func (o *ControlsPolicyFilterConfig) GetUseWaf() bool`

GetUseWaf returns the UseWaf field if non-nil, zero value otherwise.

### GetUseWafOk

`func (o *ControlsPolicyFilterConfig) GetUseWafOk() (*bool, bool)`

GetUseWafOk returns a tuple with the UseWaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWaf

`func (o *ControlsPolicyFilterConfig) SetUseWaf(v bool)`

SetUseWaf sets UseWaf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


