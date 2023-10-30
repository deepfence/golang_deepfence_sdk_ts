# ModelFilesystemAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anomaly** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Tactics** | Pointer to **[]string** |  | [optional] 
**Techniques** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewModelFilesystemAlertRule

`func NewModelFilesystemAlertRule() *ModelFilesystemAlertRule`

NewModelFilesystemAlertRule instantiates a new ModelFilesystemAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFilesystemAlertRuleWithDefaults

`func NewModelFilesystemAlertRuleWithDefaults() *ModelFilesystemAlertRule`

NewModelFilesystemAlertRuleWithDefaults instantiates a new ModelFilesystemAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomaly

`func (o *ModelFilesystemAlertRule) GetAnomaly() string`

GetAnomaly returns the Anomaly field if non-nil, zero value otherwise.

### GetAnomalyOk

`func (o *ModelFilesystemAlertRule) GetAnomalyOk() (*string, bool)`

GetAnomalyOk returns a tuple with the Anomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomaly

`func (o *ModelFilesystemAlertRule) SetAnomaly(v string)`

SetAnomaly sets Anomaly field to given value.

### HasAnomaly

`func (o *ModelFilesystemAlertRule) HasAnomaly() bool`

HasAnomaly returns a boolean if a field has been set.

### GetCategory

`func (o *ModelFilesystemAlertRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelFilesystemAlertRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelFilesystemAlertRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelFilesystemAlertRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEventType

`func (o *ModelFilesystemAlertRule) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelFilesystemAlertRule) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelFilesystemAlertRule) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ModelFilesystemAlertRule) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetResourceType

`func (o *ModelFilesystemAlertRule) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelFilesystemAlertRule) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelFilesystemAlertRule) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ModelFilesystemAlertRule) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRuleId

`func (o *ModelFilesystemAlertRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelFilesystemAlertRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelFilesystemAlertRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ModelFilesystemAlertRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSeverity

`func (o *ModelFilesystemAlertRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelFilesystemAlertRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelFilesystemAlertRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ModelFilesystemAlertRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *ModelFilesystemAlertRule) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelFilesystemAlertRule) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelFilesystemAlertRule) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ModelFilesystemAlertRule) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTactics

`func (o *ModelFilesystemAlertRule) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelFilesystemAlertRule) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelFilesystemAlertRule) SetTactics(v []string)`

SetTactics sets Tactics field to given value.

### HasTactics

`func (o *ModelFilesystemAlertRule) HasTactics() bool`

HasTactics returns a boolean if a field has been set.

### SetTacticsNil

`func (o *ModelFilesystemAlertRule) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelFilesystemAlertRule) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
### GetTechniques

`func (o *ModelFilesystemAlertRule) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelFilesystemAlertRule) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelFilesystemAlertRule) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.

### HasTechniques

`func (o *ModelFilesystemAlertRule) HasTechniques() bool`

HasTechniques returns a boolean if a field has been set.

### SetTechniquesNil

`func (o *ModelFilesystemAlertRule) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelFilesystemAlertRule) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil
### GetType

`func (o *ModelFilesystemAlertRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelFilesystemAlertRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelFilesystemAlertRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelFilesystemAlertRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


