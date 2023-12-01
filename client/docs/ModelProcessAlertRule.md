# ModelProcessAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anomaly** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Tactics** | Pointer to **[]string** |  | [optional] 
**Techniques** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelProcessAlertRule

`func NewModelProcessAlertRule() *ModelProcessAlertRule`

NewModelProcessAlertRule instantiates a new ModelProcessAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelProcessAlertRuleWithDefaults

`func NewModelProcessAlertRuleWithDefaults() *ModelProcessAlertRule`

NewModelProcessAlertRuleWithDefaults instantiates a new ModelProcessAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomaly

`func (o *ModelProcessAlertRule) GetAnomaly() string`

GetAnomaly returns the Anomaly field if non-nil, zero value otherwise.

### GetAnomalyOk

`func (o *ModelProcessAlertRule) GetAnomalyOk() (*string, bool)`

GetAnomalyOk returns a tuple with the Anomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomaly

`func (o *ModelProcessAlertRule) SetAnomaly(v string)`

SetAnomaly sets Anomaly field to given value.

### HasAnomaly

`func (o *ModelProcessAlertRule) HasAnomaly() bool`

HasAnomaly returns a boolean if a field has been set.

### GetCategory

`func (o *ModelProcessAlertRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelProcessAlertRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelProcessAlertRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelProcessAlertRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEventType

`func (o *ModelProcessAlertRule) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelProcessAlertRule) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelProcessAlertRule) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ModelProcessAlertRule) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetRuleId

`func (o *ModelProcessAlertRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelProcessAlertRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelProcessAlertRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ModelProcessAlertRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSeverity

`func (o *ModelProcessAlertRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelProcessAlertRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelProcessAlertRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ModelProcessAlertRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *ModelProcessAlertRule) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelProcessAlertRule) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelProcessAlertRule) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ModelProcessAlertRule) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTactics

`func (o *ModelProcessAlertRule) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelProcessAlertRule) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelProcessAlertRule) SetTactics(v []string)`

SetTactics sets Tactics field to given value.

### HasTactics

`func (o *ModelProcessAlertRule) HasTactics() bool`

HasTactics returns a boolean if a field has been set.

### SetTacticsNil

`func (o *ModelProcessAlertRule) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelProcessAlertRule) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
### GetTechniques

`func (o *ModelProcessAlertRule) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelProcessAlertRule) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelProcessAlertRule) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.

### HasTechniques

`func (o *ModelProcessAlertRule) HasTechniques() bool`

HasTechniques returns a boolean if a field has been set.

### SetTechniquesNil

`func (o *ModelProcessAlertRule) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelProcessAlertRule) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


