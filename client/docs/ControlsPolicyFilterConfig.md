# ControlsPolicyFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**Hash** | **string** |  | 
**Matcher** | **map[string]interface{}** |  | 

## Methods

### NewControlsPolicyFilterConfig

`func NewControlsPolicyFilterConfig(eventType string, hash string, matcher map[string]interface{}, ) *ControlsPolicyFilterConfig`

NewControlsPolicyFilterConfig instantiates a new ControlsPolicyFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPolicyFilterConfigWithDefaults

`func NewControlsPolicyFilterConfigWithDefaults() *ControlsPolicyFilterConfig`

NewControlsPolicyFilterConfigWithDefaults instantiates a new ControlsPolicyFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ControlsPolicyFilterConfig) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ControlsPolicyFilterConfig) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ControlsPolicyFilterConfig) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetHash

`func (o *ControlsPolicyFilterConfig) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ControlsPolicyFilterConfig) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ControlsPolicyFilterConfig) SetHash(v string)`

SetHash sets Hash field to given value.


### GetMatcher

`func (o *ControlsPolicyFilterConfig) GetMatcher() map[string]interface{}`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *ControlsPolicyFilterConfig) GetMatcherOk() (*map[string]interface{}, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *ControlsPolicyFilterConfig) SetMatcher(v map[string]interface{})`

SetMatcher sets Matcher field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


