# ModelSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExploitabilityScore** | Pointer to **int32** |  | [optional] 
**FullFilename** | **string** |  | 
**Level** | **string** |  | 
**Masked** | **bool** |  | 
**MatchedContent** | **string** |  | 
**MaxExploitabilityScore** | Pointer to **int32** |  | [optional] 
**NodeId** | **string** |  | 
**Resources** | Pointer to [**[]ModelBasicNode**](ModelBasicNode.md) |  | [optional] 
**RuleId** | **string** |  | 
**Score** | **float32** |  | 
**StartingIndex** | **int32** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelSecret

`func NewModelSecret(fullFilename string, level string, masked bool, matchedContent string, nodeId string, ruleId string, score float32, startingIndex int32, updatedAt int32, ) *ModelSecret`

NewModelSecret instantiates a new ModelSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretWithDefaults

`func NewModelSecretWithDefaults() *ModelSecret`

NewModelSecretWithDefaults instantiates a new ModelSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExploitabilityScore

`func (o *ModelSecret) GetExploitabilityScore() int32`

GetExploitabilityScore returns the ExploitabilityScore field if non-nil, zero value otherwise.

### GetExploitabilityScoreOk

`func (o *ModelSecret) GetExploitabilityScoreOk() (*int32, bool)`

GetExploitabilityScoreOk returns a tuple with the ExploitabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitabilityScore

`func (o *ModelSecret) SetExploitabilityScore(v int32)`

SetExploitabilityScore sets ExploitabilityScore field to given value.

### HasExploitabilityScore

`func (o *ModelSecret) HasExploitabilityScore() bool`

HasExploitabilityScore returns a boolean if a field has been set.

### GetFullFilename

`func (o *ModelSecret) GetFullFilename() string`

GetFullFilename returns the FullFilename field if non-nil, zero value otherwise.

### GetFullFilenameOk

`func (o *ModelSecret) GetFullFilenameOk() (*string, bool)`

GetFullFilenameOk returns a tuple with the FullFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullFilename

`func (o *ModelSecret) SetFullFilename(v string)`

SetFullFilename sets FullFilename field to given value.


### GetLevel

`func (o *ModelSecret) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ModelSecret) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ModelSecret) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMasked

`func (o *ModelSecret) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelSecret) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelSecret) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetMatchedContent

`func (o *ModelSecret) GetMatchedContent() string`

GetMatchedContent returns the MatchedContent field if non-nil, zero value otherwise.

### GetMatchedContentOk

`func (o *ModelSecret) GetMatchedContentOk() (*string, bool)`

GetMatchedContentOk returns a tuple with the MatchedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedContent

`func (o *ModelSecret) SetMatchedContent(v string)`

SetMatchedContent sets MatchedContent field to given value.


### GetMaxExploitabilityScore

`func (o *ModelSecret) GetMaxExploitabilityScore() int32`

GetMaxExploitabilityScore returns the MaxExploitabilityScore field if non-nil, zero value otherwise.

### GetMaxExploitabilityScoreOk

`func (o *ModelSecret) GetMaxExploitabilityScoreOk() (*int32, bool)`

GetMaxExploitabilityScoreOk returns a tuple with the MaxExploitabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExploitabilityScore

`func (o *ModelSecret) SetMaxExploitabilityScore(v int32)`

SetMaxExploitabilityScore sets MaxExploitabilityScore field to given value.

### HasMaxExploitabilityScore

`func (o *ModelSecret) HasMaxExploitabilityScore() bool`

HasMaxExploitabilityScore returns a boolean if a field has been set.

### GetNodeId

`func (o *ModelSecret) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelSecret) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelSecret) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetResources

`func (o *ModelSecret) GetResources() []ModelBasicNode`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ModelSecret) GetResourcesOk() (*[]ModelBasicNode, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ModelSecret) SetResources(v []ModelBasicNode)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ModelSecret) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *ModelSecret) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *ModelSecret) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetRuleId

`func (o *ModelSecret) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelSecret) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelSecret) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetScore

`func (o *ModelSecret) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ModelSecret) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ModelSecret) SetScore(v float32)`

SetScore sets Score field to given value.


### GetStartingIndex

`func (o *ModelSecret) GetStartingIndex() int32`

GetStartingIndex returns the StartingIndex field if non-nil, zero value otherwise.

### GetStartingIndexOk

`func (o *ModelSecret) GetStartingIndexOk() (*int32, bool)`

GetStartingIndexOk returns a tuple with the StartingIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingIndex

`func (o *ModelSecret) SetStartingIndex(v int32)`

SetStartingIndex sets StartingIndex field to given value.


### GetUpdatedAt

`func (o *ModelSecret) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelSecret) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelSecret) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


