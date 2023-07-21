# ModelEnableFilesystemTracerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | [**[]ModelAgentId**](ModelAgentId.md) |  | 
**Path** | **string** |  | 

## Methods

### NewModelEnableFilesystemTracerReq

`func NewModelEnableFilesystemTracerReq(agentIds []ModelAgentId, path string, ) *ModelEnableFilesystemTracerReq`

NewModelEnableFilesystemTracerReq instantiates a new ModelEnableFilesystemTracerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEnableFilesystemTracerReqWithDefaults

`func NewModelEnableFilesystemTracerReqWithDefaults() *ModelEnableFilesystemTracerReq`

NewModelEnableFilesystemTracerReqWithDefaults instantiates a new ModelEnableFilesystemTracerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *ModelEnableFilesystemTracerReq) GetAgentIds() []ModelAgentId`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *ModelEnableFilesystemTracerReq) GetAgentIdsOk() (*[]ModelAgentId, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *ModelEnableFilesystemTracerReq) SetAgentIds(v []ModelAgentId)`

SetAgentIds sets AgentIds field to given value.


### SetAgentIdsNil

`func (o *ModelEnableFilesystemTracerReq) SetAgentIdsNil(b bool)`

 SetAgentIdsNil sets the value for AgentIds to be an explicit nil

### UnsetAgentIds
`func (o *ModelEnableFilesystemTracerReq) UnsetAgentIds()`

UnsetAgentIds ensures that no value is present for AgentIds, not even an explicit nil
### GetPath

`func (o *ModelEnableFilesystemTracerReq) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelEnableFilesystemTracerReq) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelEnableFilesystemTracerReq) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


