# ModelAttachAgentConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | [**[]ModelAgentId**](ModelAgentId.md) |  | 
**ConfigId** | **string** |  | 

## Methods

### NewModelAttachAgentConfigReq

`func NewModelAttachAgentConfigReq(agentIds []ModelAgentId, configId string, ) *ModelAttachAgentConfigReq`

NewModelAttachAgentConfigReq instantiates a new ModelAttachAgentConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAttachAgentConfigReqWithDefaults

`func NewModelAttachAgentConfigReqWithDefaults() *ModelAttachAgentConfigReq`

NewModelAttachAgentConfigReqWithDefaults instantiates a new ModelAttachAgentConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *ModelAttachAgentConfigReq) GetAgentIds() []ModelAgentId`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *ModelAttachAgentConfigReq) GetAgentIdsOk() (*[]ModelAgentId, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *ModelAttachAgentConfigReq) SetAgentIds(v []ModelAgentId)`

SetAgentIds sets AgentIds field to given value.


### SetAgentIdsNil

`func (o *ModelAttachAgentConfigReq) SetAgentIdsNil(b bool)`

 SetAgentIdsNil sets the value for AgentIds to be an explicit nil

### UnsetAgentIds
`func (o *ModelAttachAgentConfigReq) UnsetAgentIds()`

UnsetAgentIds ensures that no value is present for AgentIds, not even an explicit nil
### GetConfigId

`func (o *ModelAttachAgentConfigReq) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ModelAttachAgentConfigReq) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ModelAttachAgentConfigReq) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


