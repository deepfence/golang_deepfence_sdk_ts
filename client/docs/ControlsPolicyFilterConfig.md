# ControlsPolicyFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**Name** | **string** |  | 
**Policies** | [**[]ControlsPolicy**](ControlsPolicy.md) |  | 

## Methods

### NewControlsPolicyFilterConfig

`func NewControlsPolicyFilterConfig(hash string, name string, policies []ControlsPolicy, ) *ControlsPolicyFilterConfig`

NewControlsPolicyFilterConfig instantiates a new ControlsPolicyFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsPolicyFilterConfigWithDefaults

`func NewControlsPolicyFilterConfigWithDefaults() *ControlsPolicyFilterConfig`

NewControlsPolicyFilterConfigWithDefaults instantiates a new ControlsPolicyFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetName

`func (o *ControlsPolicyFilterConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsPolicyFilterConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsPolicyFilterConfig) SetName(v string)`

SetName sets Name field to given value.


### GetPolicies

`func (o *ControlsPolicyFilterConfig) GetPolicies() []ControlsPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ControlsPolicyFilterConfig) GetPoliciesOk() (*[]ControlsPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ControlsPolicyFilterConfig) SetPolicies(v []ControlsPolicy)`

SetPolicies sets Policies field to given value.


### SetPoliciesNil

`func (o *ControlsPolicyFilterConfig) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *ControlsPolicyFilterConfig) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


