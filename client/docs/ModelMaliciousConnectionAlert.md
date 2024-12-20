# ModelMaliciousConnectionAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIp** | **string** |  | 
**Direction** | **string** |  | 
**EventType** | **string** |  | 
**HostName** | **string** |  | 
**NodeId** | **string** |  | 
**Severity** | **string** |  | 
**SourceIp** | **string** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Techniques** | **[]string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelMaliciousConnectionAlert

`func NewModelMaliciousConnectionAlert(destinationIp string, direction string, eventType string, hostName string, nodeId string, severity string, sourceIp string, summary string, tactics []string, techniques []string, updatedAt int32, ) *ModelMaliciousConnectionAlert`

NewModelMaliciousConnectionAlert instantiates a new ModelMaliciousConnectionAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMaliciousConnectionAlertWithDefaults

`func NewModelMaliciousConnectionAlertWithDefaults() *ModelMaliciousConnectionAlert`

NewModelMaliciousConnectionAlertWithDefaults instantiates a new ModelMaliciousConnectionAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIp

`func (o *ModelMaliciousConnectionAlert) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *ModelMaliciousConnectionAlert) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *ModelMaliciousConnectionAlert) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.


### GetDirection

`func (o *ModelMaliciousConnectionAlert) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModelMaliciousConnectionAlert) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModelMaliciousConnectionAlert) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEventType

`func (o *ModelMaliciousConnectionAlert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelMaliciousConnectionAlert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelMaliciousConnectionAlert) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetHostName

`func (o *ModelMaliciousConnectionAlert) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelMaliciousConnectionAlert) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelMaliciousConnectionAlert) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetNodeId

`func (o *ModelMaliciousConnectionAlert) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelMaliciousConnectionAlert) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelMaliciousConnectionAlert) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetSeverity

`func (o *ModelMaliciousConnectionAlert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelMaliciousConnectionAlert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelMaliciousConnectionAlert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSourceIp

`func (o *ModelMaliciousConnectionAlert) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *ModelMaliciousConnectionAlert) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *ModelMaliciousConnectionAlert) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.


### GetSummary

`func (o *ModelMaliciousConnectionAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelMaliciousConnectionAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelMaliciousConnectionAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTactics

`func (o *ModelMaliciousConnectionAlert) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelMaliciousConnectionAlert) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelMaliciousConnectionAlert) SetTactics(v []string)`

SetTactics sets Tactics field to given value.


### SetTacticsNil

`func (o *ModelMaliciousConnectionAlert) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelMaliciousConnectionAlert) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
### GetTechniques

`func (o *ModelMaliciousConnectionAlert) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelMaliciousConnectionAlert) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelMaliciousConnectionAlert) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.


### SetTechniquesNil

`func (o *ModelMaliciousConnectionAlert) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelMaliciousConnectionAlert) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil
### GetUpdatedAt

`func (o *ModelMaliciousConnectionAlert) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelMaliciousConnectionAlert) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelMaliciousConnectionAlert) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


