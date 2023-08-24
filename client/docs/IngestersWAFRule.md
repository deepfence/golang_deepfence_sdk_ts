# IngestersWAFRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ExecutedAt** | Pointer to **int32** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**RemoteIp** | **string** |  | 
**RemotePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewIngestersWAFRule

`func NewIngestersWAFRule(action string, remoteIp string, ) *IngestersWAFRule`

NewIngestersWAFRule instantiates a new IngestersWAFRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestersWAFRuleWithDefaults

`func NewIngestersWAFRuleWithDefaults() *IngestersWAFRule`

NewIngestersWAFRuleWithDefaults instantiates a new IngestersWAFRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IngestersWAFRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IngestersWAFRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IngestersWAFRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetExecutedAt

`func (o *IngestersWAFRule) GetExecutedAt() int32`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *IngestersWAFRule) GetExecutedAtOk() (*int32, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *IngestersWAFRule) SetExecutedAt(v int32)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *IngestersWAFRule) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### GetHostName

`func (o *IngestersWAFRule) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *IngestersWAFRule) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *IngestersWAFRule) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *IngestersWAFRule) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetRemoteIp

`func (o *IngestersWAFRule) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *IngestersWAFRule) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *IngestersWAFRule) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemotePort

`func (o *IngestersWAFRule) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *IngestersWAFRule) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *IngestersWAFRule) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *IngestersWAFRule) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


