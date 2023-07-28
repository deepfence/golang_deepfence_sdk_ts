# ControlsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationSec** | **int32** |  | 
**EventType** | **string** |  | 
**Hash** | **string** |  | 
**Matcher** | **map[string]interface{}** |  | 

## Methods

### NewControlsPolicy

`func NewControlsPolicy(durationSec int32, eventType string, hash string, matcher map[string]interface{}, ) *ControlsPolicy`

NewControlsPolicy instantiates a new ControlsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPolicyWithDefaults

`func NewControlsPolicyWithDefaults() *ControlsPolicy`

NewControlsPolicyWithDefaults instantiates a new ControlsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetHash

`func (o *ControlsPolicy) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ControlsPolicy) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ControlsPolicy) SetHash(v string)`

SetHash sets Hash field to given value.


### GetMatcher

`func (o *ControlsPolicy) GetMatcher() map[string]interface{}`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *ControlsPolicy) GetMatcherOk() (*map[string]interface{}, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *ControlsPolicy) SetMatcher(v map[string]interface{})`

SetMatcher sets Matcher field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


