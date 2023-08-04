# ControlsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountLimit** | **int32** |  | 
**DurationCountLimitSec** | **int32** |  | 
**DurationSec** | **int32** |  | 
**EventType** | **string** |  | 
**Matcher** | [**ControlsPolicyAlertMatcher**](ControlsPolicyAlertMatcher.md) |  | 
**PolicyId** | **string** |  | 

## Methods

### NewControlsPolicy

`func NewControlsPolicy(countLimit int32, durationCountLimitSec int32, durationSec int32, eventType string, matcher ControlsPolicyAlertMatcher, policyId string, ) *ControlsPolicy`

NewControlsPolicy instantiates a new ControlsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPolicyWithDefaults

`func NewControlsPolicyWithDefaults() *ControlsPolicy`

NewControlsPolicyWithDefaults instantiates a new ControlsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountLimit

`func (o *ControlsPolicy) GetCountLimit() int32`

GetCountLimit returns the CountLimit field if non-nil, zero value otherwise.

### GetCountLimitOk

`func (o *ControlsPolicy) GetCountLimitOk() (*int32, bool)`

GetCountLimitOk returns a tuple with the CountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimit

`func (o *ControlsPolicy) SetCountLimit(v int32)`

SetCountLimit sets CountLimit field to given value.


### GetDurationCountLimitSec

`func (o *ControlsPolicy) GetDurationCountLimitSec() int32`

GetDurationCountLimitSec returns the DurationCountLimitSec field if non-nil, zero value otherwise.

### GetDurationCountLimitSecOk

`func (o *ControlsPolicy) GetDurationCountLimitSecOk() (*int32, bool)`

GetDurationCountLimitSecOk returns a tuple with the DurationCountLimitSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCountLimitSec

`func (o *ControlsPolicy) SetDurationCountLimitSec(v int32)`

SetDurationCountLimitSec sets DurationCountLimitSec field to given value.


### GetDurationSec

`func (o *ControlsPolicy) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *ControlsPolicy) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *ControlsPolicy) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.


### GetEventType

`func (o *ControlsPolicy) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ControlsPolicy) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ControlsPolicy) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMatcher

`func (o *ControlsPolicy) GetMatcher() ControlsPolicyAlertMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *ControlsPolicy) GetMatcherOk() (*ControlsPolicyAlertMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *ControlsPolicy) SetMatcher(v ControlsPolicyAlertMatcher)`

SetMatcher sets Matcher field to given value.


### GetPolicyId

`func (o *ControlsPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ControlsPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ControlsPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


