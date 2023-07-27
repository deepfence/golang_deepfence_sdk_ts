# ControlsFilesystemTracerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**MonitoredFiles** | [**[]ControlsMonitoredFilesConfig**](ControlsMonitoredFilesConfig.md) |  | 
**ProcessEvents** | [**[]ControlsProcessEventConfig**](ControlsProcessEventConfig.md) |  | 

## Methods

### NewControlsFilesystemTracerConfig

`func NewControlsFilesystemTracerConfig(hash string, monitoredFiles []ControlsMonitoredFilesConfig, processEvents []ControlsProcessEventConfig, ) *ControlsFilesystemTracerConfig`

NewControlsFilesystemTracerConfig instantiates a new ControlsFilesystemTracerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsFilesystemTracerConfigWithDefaults

`func NewControlsFilesystemTracerConfigWithDefaults() *ControlsFilesystemTracerConfig`

NewControlsFilesystemTracerConfigWithDefaults instantiates a new ControlsFilesystemTracerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ControlsFilesystemTracerConfig) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ControlsFilesystemTracerConfig) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ControlsFilesystemTracerConfig) SetHash(v string)`

SetHash sets Hash field to given value.


### GetMonitoredFiles

`func (o *ControlsFilesystemTracerConfig) GetMonitoredFiles() []ControlsMonitoredFilesConfig`

GetMonitoredFiles returns the MonitoredFiles field if non-nil, zero value otherwise.

### GetMonitoredFilesOk

`func (o *ControlsFilesystemTracerConfig) GetMonitoredFilesOk() (*[]ControlsMonitoredFilesConfig, bool)`

GetMonitoredFilesOk returns a tuple with the MonitoredFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredFiles

`func (o *ControlsFilesystemTracerConfig) SetMonitoredFiles(v []ControlsMonitoredFilesConfig)`

SetMonitoredFiles sets MonitoredFiles field to given value.


### SetMonitoredFilesNil

`func (o *ControlsFilesystemTracerConfig) SetMonitoredFilesNil(b bool)`

 SetMonitoredFilesNil sets the value for MonitoredFiles to be an explicit nil

### UnsetMonitoredFiles
`func (o *ControlsFilesystemTracerConfig) UnsetMonitoredFiles()`

UnsetMonitoredFiles ensures that no value is present for MonitoredFiles, not even an explicit nil
### GetProcessEvents

`func (o *ControlsFilesystemTracerConfig) GetProcessEvents() []ControlsProcessEventConfig`

GetProcessEvents returns the ProcessEvents field if non-nil, zero value otherwise.

### GetProcessEventsOk

`func (o *ControlsFilesystemTracerConfig) GetProcessEventsOk() (*[]ControlsProcessEventConfig, bool)`

GetProcessEventsOk returns a tuple with the ProcessEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessEvents

`func (o *ControlsFilesystemTracerConfig) SetProcessEvents(v []ControlsProcessEventConfig)`

SetProcessEvents sets ProcessEvents field to given value.


### SetProcessEventsNil

`func (o *ControlsFilesystemTracerConfig) SetProcessEventsNil(b bool)`

 SetProcessEventsNil sets the value for ProcessEvents to be an explicit nil

### UnsetProcessEvents
`func (o *ControlsFilesystemTracerConfig) UnsetProcessEvents()`

UnsetProcessEvents ensures that no value is present for ProcessEvents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


