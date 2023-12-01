# ControlsProcessTracerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitoredprocessevents** | Pointer to [**[]ControlsProcessEventEntry**](ControlsProcessEventEntry.md) |  | [optional] 
**NodeId** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewControlsProcessTracerConfig

`func NewControlsProcessTracerConfig(nodeId string, updatedAt int32, ) *ControlsProcessTracerConfig`

NewControlsProcessTracerConfig instantiates a new ControlsProcessTracerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsProcessTracerConfigWithDefaults

`func NewControlsProcessTracerConfigWithDefaults() *ControlsProcessTracerConfig`

NewControlsProcessTracerConfigWithDefaults instantiates a new ControlsProcessTracerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoredprocessevents

`func (o *ControlsProcessTracerConfig) GetMonitoredprocessevents() []ControlsProcessEventEntry`

GetMonitoredprocessevents returns the Monitoredprocessevents field if non-nil, zero value otherwise.

### GetMonitoredprocesseventsOk

`func (o *ControlsProcessTracerConfig) GetMonitoredprocesseventsOk() (*[]ControlsProcessEventEntry, bool)`

GetMonitoredprocesseventsOk returns a tuple with the Monitoredprocessevents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredprocessevents

`func (o *ControlsProcessTracerConfig) SetMonitoredprocessevents(v []ControlsProcessEventEntry)`

SetMonitoredprocessevents sets Monitoredprocessevents field to given value.

### HasMonitoredprocessevents

`func (o *ControlsProcessTracerConfig) HasMonitoredprocessevents() bool`

HasMonitoredprocessevents returns a boolean if a field has been set.

### SetMonitoredprocesseventsNil

`func (o *ControlsProcessTracerConfig) SetMonitoredprocesseventsNil(b bool)`

 SetMonitoredprocesseventsNil sets the value for Monitoredprocessevents to be an explicit nil

### UnsetMonitoredprocessevents
`func (o *ControlsProcessTracerConfig) UnsetMonitoredprocessevents()`

UnsetMonitoredprocessevents ensures that no value is present for Monitoredprocessevents, not even an explicit nil
### GetNodeId

`func (o *ControlsProcessTracerConfig) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ControlsProcessTracerConfig) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ControlsProcessTracerConfig) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUpdatedAt

`func (o *ControlsProcessTracerConfig) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsProcessTracerConfig) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsProcessTracerConfig) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


