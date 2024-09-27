# ModelNetworkAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | 
**Description** | **string** |  | 
**Masked** | **bool** |  | 
**NodeId** | **string** |  | 
**Payload** | **string** |  | 
**RuleId** | **string** |  | 
**Severity** | **string** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Techniques** | **[]string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelNetworkAlertRule

`func NewModelNetworkAlertRule(category string, description string, masked bool, nodeId string, payload string, ruleId string, severity string, summary string, tactics []string, techniques []string, updatedAt int32, ) *ModelNetworkAlertRule`

NewModelNetworkAlertRule instantiates a new ModelNetworkAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkAlertRuleWithDefaults

`func NewModelNetworkAlertRuleWithDefaults() *ModelNetworkAlertRule`

NewModelNetworkAlertRuleWithDefaults instantiates a new ModelNetworkAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ModelNetworkAlertRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelNetworkAlertRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelNetworkAlertRule) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *ModelNetworkAlertRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelNetworkAlertRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelNetworkAlertRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMasked

`func (o *ModelNetworkAlertRule) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelNetworkAlertRule) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelNetworkAlertRule) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetNodeId

`func (o *ModelNetworkAlertRule) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelNetworkAlertRule) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelNetworkAlertRule) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetPayload

`func (o *ModelNetworkAlertRule) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ModelNetworkAlertRule) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ModelNetworkAlertRule) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetRuleId

`func (o *ModelNetworkAlertRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelNetworkAlertRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelNetworkAlertRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetSeverity

`func (o *ModelNetworkAlertRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelNetworkAlertRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelNetworkAlertRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSummary

`func (o *ModelNetworkAlertRule) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelNetworkAlertRule) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelNetworkAlertRule) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTactics

`func (o *ModelNetworkAlertRule) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelNetworkAlertRule) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelNetworkAlertRule) SetTactics(v []string)`

SetTactics sets Tactics field to given value.


### SetTacticsNil

`func (o *ModelNetworkAlertRule) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelNetworkAlertRule) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
### GetTechniques

`func (o *ModelNetworkAlertRule) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelNetworkAlertRule) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelNetworkAlertRule) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.


### SetTechniquesNil

`func (o *ModelNetworkAlertRule) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelNetworkAlertRule) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil
### GetUpdatedAt

`func (o *ModelNetworkAlertRule) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelNetworkAlertRule) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelNetworkAlertRule) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


