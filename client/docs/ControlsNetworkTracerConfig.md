# ControlsNetworkTracerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 
**HttpsRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 
**Mode** | **string** |  | 
**NodeId** | **string** |  | 
**ProcessNames** | **[]string** |  | 
**TcpRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 

## Methods

### NewControlsNetworkTracerConfig

`func NewControlsNetworkTracerConfig(httpRules ControlsNetworkRules, httpsRules ControlsNetworkRules, mode string, nodeId string, processNames []string, tcpRules ControlsNetworkRules, ) *ControlsNetworkTracerConfig`

NewControlsNetworkTracerConfig instantiates a new ControlsNetworkTracerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsNetworkTracerConfigWithDefaults

`func NewControlsNetworkTracerConfigWithDefaults() *ControlsNetworkTracerConfig`

NewControlsNetworkTracerConfigWithDefaults instantiates a new ControlsNetworkTracerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpRules

`func (o *ControlsNetworkTracerConfig) GetHttpRules() ControlsNetworkRules`

GetHttpRules returns the HttpRules field if non-nil, zero value otherwise.

### GetHttpRulesOk

`func (o *ControlsNetworkTracerConfig) GetHttpRulesOk() (*ControlsNetworkRules, bool)`

GetHttpRulesOk returns a tuple with the HttpRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRules

`func (o *ControlsNetworkTracerConfig) SetHttpRules(v ControlsNetworkRules)`

SetHttpRules sets HttpRules field to given value.


### GetHttpsRules

`func (o *ControlsNetworkTracerConfig) GetHttpsRules() ControlsNetworkRules`

GetHttpsRules returns the HttpsRules field if non-nil, zero value otherwise.

### GetHttpsRulesOk

`func (o *ControlsNetworkTracerConfig) GetHttpsRulesOk() (*ControlsNetworkRules, bool)`

GetHttpsRulesOk returns a tuple with the HttpsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRules

`func (o *ControlsNetworkTracerConfig) SetHttpsRules(v ControlsNetworkRules)`

SetHttpsRules sets HttpsRules field to given value.


### GetMode

`func (o *ControlsNetworkTracerConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ControlsNetworkTracerConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ControlsNetworkTracerConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetNodeId

`func (o *ControlsNetworkTracerConfig) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ControlsNetworkTracerConfig) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ControlsNetworkTracerConfig) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetProcessNames

`func (o *ControlsNetworkTracerConfig) GetProcessNames() []string`

GetProcessNames returns the ProcessNames field if non-nil, zero value otherwise.

### GetProcessNamesOk

`func (o *ControlsNetworkTracerConfig) GetProcessNamesOk() (*[]string, bool)`

GetProcessNamesOk returns a tuple with the ProcessNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessNames

`func (o *ControlsNetworkTracerConfig) SetProcessNames(v []string)`

SetProcessNames sets ProcessNames field to given value.


### SetProcessNamesNil

`func (o *ControlsNetworkTracerConfig) SetProcessNamesNil(b bool)`

 SetProcessNamesNil sets the value for ProcessNames to be an explicit nil

### UnsetProcessNames
`func (o *ControlsNetworkTracerConfig) UnsetProcessNames()`

UnsetProcessNames ensures that no value is present for ProcessNames, not even an explicit nil
### GetTcpRules

`func (o *ControlsNetworkTracerConfig) GetTcpRules() ControlsNetworkRules`

GetTcpRules returns the TcpRules field if non-nil, zero value otherwise.

### GetTcpRulesOk

`func (o *ControlsNetworkTracerConfig) GetTcpRulesOk() (*ControlsNetworkRules, bool)`

GetTcpRulesOk returns a tuple with the TcpRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpRules

`func (o *ControlsNetworkTracerConfig) SetTcpRules(v ControlsNetworkRules)`

SetTcpRules sets TcpRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


