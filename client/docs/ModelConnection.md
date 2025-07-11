# ModelConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Ips** | Pointer to **[]interface{}** |  | [optional] 
**MaliciousIp** | Pointer to **[]bool** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 

## Methods

### NewModelConnection

`func NewModelConnection() *ModelConnection`

NewModelConnection instantiates a new ModelConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConnectionWithDefaults

`func NewModelConnectionWithDefaults() *ModelConnection`

NewModelConnectionWithDefaults instantiates a new ModelConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ModelConnection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelConnection) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelConnection) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ModelConnection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIps

`func (o *ModelConnection) GetIps() []interface{}`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ModelConnection) GetIpsOk() (*[]interface{}, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ModelConnection) SetIps(v []interface{})`

SetIps sets Ips field to given value.

### HasIps

`func (o *ModelConnection) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *ModelConnection) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *ModelConnection) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetMaliciousIp

`func (o *ModelConnection) GetMaliciousIp() []bool`

GetMaliciousIp returns the MaliciousIp field if non-nil, zero value otherwise.

### GetMaliciousIpOk

`func (o *ModelConnection) GetMaliciousIpOk() (*[]bool, bool)`

GetMaliciousIpOk returns a tuple with the MaliciousIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousIp

`func (o *ModelConnection) SetMaliciousIp(v []bool)`

SetMaliciousIp sets MaliciousIp field to given value.

### HasMaliciousIp

`func (o *ModelConnection) HasMaliciousIp() bool`

HasMaliciousIp returns a boolean if a field has been set.

### SetMaliciousIpNil

`func (o *ModelConnection) SetMaliciousIpNil(b bool)`

 SetMaliciousIpNil sets the value for MaliciousIp to be an explicit nil

### UnsetMaliciousIp
`func (o *ModelConnection) UnsetMaliciousIp()`

UnsetMaliciousIp ensures that no value is present for MaliciousIp, not even an explicit nil
### GetNodeId

`func (o *ModelConnection) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelConnection) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelConnection) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelConnection) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeName

`func (o *ModelConnection) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelConnection) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelConnection) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *ModelConnection) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


