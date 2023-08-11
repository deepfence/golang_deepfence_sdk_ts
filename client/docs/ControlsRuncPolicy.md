# ControlsRuncPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**CountLimit** | **int32** |  | 
**DurationCountLimitSec** | **int32** |  | 
**Matcher** | [**ControlsPolicyAlertMatcher**](ControlsPolicyAlertMatcher.md) |  | 
**NodeType** | **string** |  | 
**PolicyId** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewControlsRuncPolicy

`func NewControlsRuncPolicy(action string, countLimit int32, durationCountLimitSec int32, matcher ControlsPolicyAlertMatcher, nodeType string, policyId string, updatedAt int32, ) *ControlsRuncPolicy`

NewControlsRuncPolicy instantiates a new ControlsRuncPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsRuncPolicyWithDefaults

`func NewControlsRuncPolicyWithDefaults() *ControlsRuncPolicy`

NewControlsRuncPolicyWithDefaults instantiates a new ControlsRuncPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ControlsRuncPolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ControlsRuncPolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ControlsRuncPolicy) SetAction(v string)`

SetAction sets Action field to given value.


### GetCountLimit

`func (o *ControlsRuncPolicy) GetCountLimit() int32`

GetCountLimit returns the CountLimit field if non-nil, zero value otherwise.

### GetCountLimitOk

`func (o *ControlsRuncPolicy) GetCountLimitOk() (*int32, bool)`

GetCountLimitOk returns a tuple with the CountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimit

`func (o *ControlsRuncPolicy) SetCountLimit(v int32)`

SetCountLimit sets CountLimit field to given value.


### GetDurationCountLimitSec

`func (o *ControlsRuncPolicy) GetDurationCountLimitSec() int32`

GetDurationCountLimitSec returns the DurationCountLimitSec field if non-nil, zero value otherwise.

### GetDurationCountLimitSecOk

`func (o *ControlsRuncPolicy) GetDurationCountLimitSecOk() (*int32, bool)`

GetDurationCountLimitSecOk returns a tuple with the DurationCountLimitSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCountLimitSec

`func (o *ControlsRuncPolicy) SetDurationCountLimitSec(v int32)`

SetDurationCountLimitSec sets DurationCountLimitSec field to given value.


### GetMatcher

`func (o *ControlsRuncPolicy) GetMatcher() ControlsPolicyAlertMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *ControlsRuncPolicy) GetMatcherOk() (*ControlsPolicyAlertMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *ControlsRuncPolicy) SetMatcher(v ControlsPolicyAlertMatcher)`

SetMatcher sets Matcher field to given value.


### GetNodeType

`func (o *ControlsRuncPolicy) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ControlsRuncPolicy) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ControlsRuncPolicy) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetPolicyId

`func (o *ControlsRuncPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ControlsRuncPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ControlsRuncPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetUpdatedAt

`func (o *ControlsRuncPolicy) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsRuncPolicy) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsRuncPolicy) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


