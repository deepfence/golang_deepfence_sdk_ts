# ModelDisableTracerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | [**[]ModelAgentID**](ModelAgentID.md) |  | 

## Methods

### NewModelDisableTracerReq

`func NewModelDisableTracerReq(agentIds []ModelAgentID, ) *ModelDisableTracerReq`

NewModelDisableTracerReq instantiates a new ModelDisableTracerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDisableTracerReqWithDefaults

`func NewModelDisableTracerReqWithDefaults() *ModelDisableTracerReq`

NewModelDisableTracerReqWithDefaults instantiates a new ModelDisableTracerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *ModelDisableTracerReq) GetAgentIds() []ModelAgentID`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *ModelDisableTracerReq) GetAgentIdsOk() (*[]ModelAgentID, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *ModelDisableTracerReq) SetAgentIds(v []ModelAgentID)`

SetAgentIds sets AgentIds field to given value.


### SetAgentIdsNil

`func (o *ModelDisableTracerReq) SetAgentIdsNil(b bool)`

 SetAgentIdsNil sets the value for AgentIds to be an explicit nil

### UnsetAgentIds
`func (o *ModelDisableTracerReq) UnsetAgentIds()`

UnsetAgentIds ensures that no value is present for AgentIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


