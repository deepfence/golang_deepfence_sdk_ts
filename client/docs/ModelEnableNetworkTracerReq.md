# ModelEnableNetworkTracerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | [**[]ModelAgentId**](ModelAgentId.md) |  | 
**HttpRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 
**HttpsRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 
**TcpRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 

## Methods

### NewModelEnableNetworkTracerReq

`func NewModelEnableNetworkTracerReq(agentIds []ModelAgentId, httpRules ControlsNetworkRules, httpsRules ControlsNetworkRules, tcpRules ControlsNetworkRules, ) *ModelEnableNetworkTracerReq`

NewModelEnableNetworkTracerReq instantiates a new ModelEnableNetworkTracerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEnableNetworkTracerReqWithDefaults

`func NewModelEnableNetworkTracerReqWithDefaults() *ModelEnableNetworkTracerReq`

NewModelEnableNetworkTracerReqWithDefaults instantiates a new ModelEnableNetworkTracerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *ModelEnableNetworkTracerReq) GetAgentIds() []ModelAgentId`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *ModelEnableNetworkTracerReq) GetAgentIdsOk() (*[]ModelAgentId, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *ModelEnableNetworkTracerReq) SetAgentIds(v []ModelAgentId)`

SetAgentIds sets AgentIds field to given value.


### SetAgentIdsNil

`func (o *ModelEnableNetworkTracerReq) SetAgentIdsNil(b bool)`

 SetAgentIdsNil sets the value for AgentIds to be an explicit nil

### UnsetAgentIds
`func (o *ModelEnableNetworkTracerReq) UnsetAgentIds()`

UnsetAgentIds ensures that no value is present for AgentIds, not even an explicit nil
### GetHttpRules

`func (o *ModelEnableNetworkTracerReq) GetHttpRules() ControlsNetworkRules`

GetHttpRules returns the HttpRules field if non-nil, zero value otherwise.

### GetHttpRulesOk

`func (o *ModelEnableNetworkTracerReq) GetHttpRulesOk() (*ControlsNetworkRules, bool)`

GetHttpRulesOk returns a tuple with the HttpRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRules

`func (o *ModelEnableNetworkTracerReq) SetHttpRules(v ControlsNetworkRules)`

SetHttpRules sets HttpRules field to given value.


### GetHttpsRules

`func (o *ModelEnableNetworkTracerReq) GetHttpsRules() ControlsNetworkRules`

GetHttpsRules returns the HttpsRules field if non-nil, zero value otherwise.

### GetHttpsRulesOk

`func (o *ModelEnableNetworkTracerReq) GetHttpsRulesOk() (*ControlsNetworkRules, bool)`

GetHttpsRulesOk returns a tuple with the HttpsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRules

`func (o *ModelEnableNetworkTracerReq) SetHttpsRules(v ControlsNetworkRules)`

SetHttpsRules sets HttpsRules field to given value.


### GetTcpRules

`func (o *ModelEnableNetworkTracerReq) GetTcpRules() ControlsNetworkRules`

GetTcpRules returns the TcpRules field if non-nil, zero value otherwise.

### GetTcpRulesOk

`func (o *ModelEnableNetworkTracerReq) GetTcpRulesOk() (*ControlsNetworkRules, bool)`

GetTcpRulesOk returns a tuple with the TcpRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpRules

`func (o *ModelEnableNetworkTracerReq) SetTcpRules(v ControlsNetworkRules)`

SetTcpRules sets TcpRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


