# ModelMitreAttackMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Tactic** | **string** |  | 
**TechniqueSummary** | [**map[string]ModelMitreTechniqueSummary**](ModelMitreTechniqueSummary.md) |  | 

## Methods

### NewModelMitreAttackMatrix

`func NewModelMitreAttackMatrix(count int32, tactic string, techniqueSummary map[string]ModelMitreTechniqueSummary, ) *ModelMitreAttackMatrix`

NewModelMitreAttackMatrix instantiates a new ModelMitreAttackMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMitreAttackMatrixWithDefaults

`func NewModelMitreAttackMatrixWithDefaults() *ModelMitreAttackMatrix`

NewModelMitreAttackMatrixWithDefaults instantiates a new ModelMitreAttackMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ModelMitreAttackMatrix) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelMitreAttackMatrix) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelMitreAttackMatrix) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTactic

`func (o *ModelMitreAttackMatrix) GetTactic() string`

GetTactic returns the Tactic field if non-nil, zero value otherwise.

### GetTacticOk

`func (o *ModelMitreAttackMatrix) GetTacticOk() (*string, bool)`

GetTacticOk returns a tuple with the Tactic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactic

`func (o *ModelMitreAttackMatrix) SetTactic(v string)`

SetTactic sets Tactic field to given value.


### GetTechniqueSummary

`func (o *ModelMitreAttackMatrix) GetTechniqueSummary() map[string]ModelMitreTechniqueSummary`

GetTechniqueSummary returns the TechniqueSummary field if non-nil, zero value otherwise.

### GetTechniqueSummaryOk

`func (o *ModelMitreAttackMatrix) GetTechniqueSummaryOk() (*map[string]ModelMitreTechniqueSummary, bool)`

GetTechniqueSummaryOk returns a tuple with the TechniqueSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueSummary

`func (o *ModelMitreAttackMatrix) SetTechniqueSummary(v map[string]ModelMitreTechniqueSummary)`

SetTechniqueSummary sets TechniqueSummary field to given value.


### SetTechniqueSummaryNil

`func (o *ModelMitreAttackMatrix) SetTechniqueSummaryNil(b bool)`

 SetTechniqueSummaryNil sets the value for TechniqueSummary to be an explicit nil

### UnsetTechniqueSummary
`func (o *ModelMitreAttackMatrix) UnsetTechniqueSummary()`

UnsetTechniqueSummary ensures that no value is present for TechniqueSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


