# ControlsNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**CountLimit** | **int32** |  | 
**DurationCountLimitSec** | **int32** |  | 
**DurationSec** | **int32** |  | 
**Matcher** | [**ControlsPolicyAlertMatcher**](ControlsPolicyAlertMatcher.md) |  | 
**PolicyId** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewControlsNetworkPolicy

`func NewControlsNetworkPolicy(action string, countLimit int32, durationCountLimitSec int32, durationSec int32, matcher ControlsPolicyAlertMatcher, policyId string, updatedAt int32, ) *ControlsNetworkPolicy`

NewControlsNetworkPolicy instantiates a new ControlsNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsNetworkPolicyWithDefaults

`func NewControlsNetworkPolicyWithDefaults() *ControlsNetworkPolicy`

NewControlsNetworkPolicyWithDefaults instantiates a new ControlsNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ControlsNetworkPolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ControlsNetworkPolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ControlsNetworkPolicy) SetAction(v string)`

SetAction sets Action field to given value.


### GetCountLimit

`func (o *ControlsNetworkPolicy) GetCountLimit() int32`

GetCountLimit returns the CountLimit field if non-nil, zero value otherwise.

### GetCountLimitOk

`func (o *ControlsNetworkPolicy) GetCountLimitOk() (*int32, bool)`

GetCountLimitOk returns a tuple with the CountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimit

`func (o *ControlsNetworkPolicy) SetCountLimit(v int32)`

SetCountLimit sets CountLimit field to given value.


### GetDurationCountLimitSec

`func (o *ControlsNetworkPolicy) GetDurationCountLimitSec() int32`

GetDurationCountLimitSec returns the DurationCountLimitSec field if non-nil, zero value otherwise.

### GetDurationCountLimitSecOk

`func (o *ControlsNetworkPolicy) GetDurationCountLimitSecOk() (*int32, bool)`

GetDurationCountLimitSecOk returns a tuple with the DurationCountLimitSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCountLimitSec

`func (o *ControlsNetworkPolicy) SetDurationCountLimitSec(v int32)`

SetDurationCountLimitSec sets DurationCountLimitSec field to given value.


### GetDurationSec

`func (o *ControlsNetworkPolicy) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *ControlsNetworkPolicy) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *ControlsNetworkPolicy) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.


### GetMatcher

`func (o *ControlsNetworkPolicy) GetMatcher() ControlsPolicyAlertMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *ControlsNetworkPolicy) GetMatcherOk() (*ControlsPolicyAlertMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *ControlsNetworkPolicy) SetMatcher(v ControlsPolicyAlertMatcher)`

SetMatcher sets Matcher field to given value.


### GetPolicyId

`func (o *ControlsNetworkPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ControlsNetworkPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ControlsNetworkPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetUpdatedAt

`func (o *ControlsNetworkPolicy) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsNetworkPolicy) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsNetworkPolicy) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


