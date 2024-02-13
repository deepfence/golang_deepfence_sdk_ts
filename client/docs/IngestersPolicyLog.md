# IngestersPolicyLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | **string** |  | 
**ConfigId** | **string** |  | 
**CreatedAt** | **int32** |  | 
**Direction** | **string** |  | 
**HostName** | **string** |  | 
**Incident** | **string** |  | 
**LocalIp** | **string** |  | 
**LocalPort** | **int32** |  | 
**PolicyIndex** | **int32** |  | 
**RemoteIp** | **string** |  | 
**RemotePort** | **int32** |  | 
**Severity** | **string** |  | 

## Methods

### NewIngestersPolicyLog

`func NewIngestersPolicyLog(alertId string, configId string, createdAt int32, direction string, hostName string, incident string, localIp string, localPort int32, policyIndex int32, remoteIp string, remotePort int32, severity string, ) *IngestersPolicyLog`

NewIngestersPolicyLog instantiates a new IngestersPolicyLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestersPolicyLogWithDefaults

`func NewIngestersPolicyLogWithDefaults() *IngestersPolicyLog`

NewIngestersPolicyLogWithDefaults instantiates a new IngestersPolicyLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *IngestersPolicyLog) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *IngestersPolicyLog) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *IngestersPolicyLog) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.


### GetConfigId

`func (o *IngestersPolicyLog) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *IngestersPolicyLog) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *IngestersPolicyLog) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetCreatedAt

`func (o *IngestersPolicyLog) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngestersPolicyLog) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngestersPolicyLog) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDirection

`func (o *IngestersPolicyLog) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IngestersPolicyLog) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IngestersPolicyLog) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetHostName

`func (o *IngestersPolicyLog) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *IngestersPolicyLog) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *IngestersPolicyLog) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetIncident

`func (o *IngestersPolicyLog) GetIncident() string`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *IngestersPolicyLog) GetIncidentOk() (*string, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *IngestersPolicyLog) SetIncident(v string)`

SetIncident sets Incident field to given value.


### GetLocalIp

`func (o *IngestersPolicyLog) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *IngestersPolicyLog) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *IngestersPolicyLog) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetLocalPort

`func (o *IngestersPolicyLog) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *IngestersPolicyLog) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *IngestersPolicyLog) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.


### GetPolicyIndex

`func (o *IngestersPolicyLog) GetPolicyIndex() int32`

GetPolicyIndex returns the PolicyIndex field if non-nil, zero value otherwise.

### GetPolicyIndexOk

`func (o *IngestersPolicyLog) GetPolicyIndexOk() (*int32, bool)`

GetPolicyIndexOk returns a tuple with the PolicyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIndex

`func (o *IngestersPolicyLog) SetPolicyIndex(v int32)`

SetPolicyIndex sets PolicyIndex field to given value.


### GetRemoteIp

`func (o *IngestersPolicyLog) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *IngestersPolicyLog) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *IngestersPolicyLog) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemotePort

`func (o *IngestersPolicyLog) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *IngestersPolicyLog) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *IngestersPolicyLog) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.


### GetSeverity

`func (o *IngestersPolicyLog) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IngestersPolicyLog) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IngestersPolicyLog) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


