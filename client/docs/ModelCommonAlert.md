# ModelCommonAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | 
**Count** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**EventType** | **string** |  | 
**Masked** | **bool** |  | 
**NodeId** | **string** |  | 
**Severity** | **string** |  | 
**Summary** | **string** |  | 
**Tactic** | **[]string** |  | 
**Technique** | **[]string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelCommonAlert

`func NewModelCommonAlert(category string, count int32, createdAt int32, eventType string, masked bool, nodeId string, severity string, summary string, tactic []string, technique []string, updatedAt int32, ) *ModelCommonAlert`

NewModelCommonAlert instantiates a new ModelCommonAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCommonAlertWithDefaults

`func NewModelCommonAlertWithDefaults() *ModelCommonAlert`

NewModelCommonAlertWithDefaults instantiates a new ModelCommonAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ModelCommonAlert) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelCommonAlert) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelCommonAlert) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCount

`func (o *ModelCommonAlert) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelCommonAlert) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelCommonAlert) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCreatedAt

`func (o *ModelCommonAlert) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelCommonAlert) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelCommonAlert) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetEventType

`func (o *ModelCommonAlert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelCommonAlert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelCommonAlert) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMasked

`func (o *ModelCommonAlert) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelCommonAlert) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelCommonAlert) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetNodeId

`func (o *ModelCommonAlert) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCommonAlert) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCommonAlert) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetSeverity

`func (o *ModelCommonAlert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelCommonAlert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelCommonAlert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSummary

`func (o *ModelCommonAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelCommonAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelCommonAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTactic

`func (o *ModelCommonAlert) GetTactic() []string`

GetTactic returns the Tactic field if non-nil, zero value otherwise.

### GetTacticOk

`func (o *ModelCommonAlert) GetTacticOk() (*[]string, bool)`

GetTacticOk returns a tuple with the Tactic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactic

`func (o *ModelCommonAlert) SetTactic(v []string)`

SetTactic sets Tactic field to given value.


### SetTacticNil

`func (o *ModelCommonAlert) SetTacticNil(b bool)`

 SetTacticNil sets the value for Tactic to be an explicit nil

### UnsetTactic
`func (o *ModelCommonAlert) UnsetTactic()`

UnsetTactic ensures that no value is present for Tactic, not even an explicit nil
### GetTechnique

`func (o *ModelCommonAlert) GetTechnique() []string`

GetTechnique returns the Technique field if non-nil, zero value otherwise.

### GetTechniqueOk

`func (o *ModelCommonAlert) GetTechniqueOk() (*[]string, bool)`

GetTechniqueOk returns a tuple with the Technique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnique

`func (o *ModelCommonAlert) SetTechnique(v []string)`

SetTechnique sets Technique field to given value.


### SetTechniqueNil

`func (o *ModelCommonAlert) SetTechniqueNil(b bool)`

 SetTechniqueNil sets the value for Technique to be an explicit nil

### UnsetTechnique
`func (o *ModelCommonAlert) UnsetTechnique()`

UnsetTechnique ensures that no value is present for Technique, not even an explicit nil
### GetUpdatedAt

`func (o *ModelCommonAlert) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelCommonAlert) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelCommonAlert) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


