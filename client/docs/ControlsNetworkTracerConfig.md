# ControlsNetworkTracerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**HttpRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 
**HttpsRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 
**TcpRules** | [**ControlsNetworkRules**](ControlsNetworkRules.md) |  | 

## Methods

### NewControlsNetworkTracerConfig

`func NewControlsNetworkTracerConfig(hash string, httpRules ControlsNetworkRules, httpsRules ControlsNetworkRules, tcpRules ControlsNetworkRules, ) *ControlsNetworkTracerConfig`

NewControlsNetworkTracerConfig instantiates a new ControlsNetworkTracerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsNetworkTracerConfigWithDefaults

`func NewControlsNetworkTracerConfigWithDefaults() *ControlsNetworkTracerConfig`

NewControlsNetworkTracerConfigWithDefaults instantiates a new ControlsNetworkTracerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ControlsNetworkTracerConfig) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ControlsNetworkTracerConfig) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ControlsNetworkTracerConfig) SetHash(v string)`

SetHash sets Hash field to given value.


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


