# ControlsFilesystemTracerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** |  | 
**UpdatedAt** | **int32** |  | 
**Watchedentries** | [**[]ControlsMonitoredFilesConfig**](ControlsMonitoredFilesConfig.md) |  | 

## Methods

### NewControlsFilesystemTracerConfig

`func NewControlsFilesystemTracerConfig(nodeId string, updatedAt int32, watchedentries []ControlsMonitoredFilesConfig, ) *ControlsFilesystemTracerConfig`

NewControlsFilesystemTracerConfig instantiates a new ControlsFilesystemTracerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsFilesystemTracerConfigWithDefaults

`func NewControlsFilesystemTracerConfigWithDefaults() *ControlsFilesystemTracerConfig`

NewControlsFilesystemTracerConfigWithDefaults instantiates a new ControlsFilesystemTracerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *ControlsFilesystemTracerConfig) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ControlsFilesystemTracerConfig) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ControlsFilesystemTracerConfig) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUpdatedAt

`func (o *ControlsFilesystemTracerConfig) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsFilesystemTracerConfig) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsFilesystemTracerConfig) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWatchedentries

`func (o *ControlsFilesystemTracerConfig) GetWatchedentries() []ControlsMonitoredFilesConfig`

GetWatchedentries returns the Watchedentries field if non-nil, zero value otherwise.

### GetWatchedentriesOk

`func (o *ControlsFilesystemTracerConfig) GetWatchedentriesOk() (*[]ControlsMonitoredFilesConfig, bool)`

GetWatchedentriesOk returns a tuple with the Watchedentries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchedentries

`func (o *ControlsFilesystemTracerConfig) SetWatchedentries(v []ControlsMonitoredFilesConfig)`

SetWatchedentries sets Watchedentries field to given value.


### SetWatchedentriesNil

`func (o *ControlsFilesystemTracerConfig) SetWatchedentriesNil(b bool)`

 SetWatchedentriesNil sets the value for Watchedentries to be an explicit nil

### UnsetWatchedentries
`func (o *ControlsFilesystemTracerConfig) UnsetWatchedentries()`

UnsetWatchedentries ensures that no value is present for Watchedentries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


