# ModelAgentPluginsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemTracerStatus** | [**ModelPluginStatus**](ModelPluginStatus.md) |  | 
**NetworkFilterStatus** | [**ModelPluginStatus**](ModelPluginStatus.md) |  | 
**NetworkTracerStatus** | [**ModelPluginStatus**](ModelPluginStatus.md) |  | 

## Methods

### NewModelAgentPluginsStatus

`func NewModelAgentPluginsStatus(filesystemTracerStatus ModelPluginStatus, networkFilterStatus ModelPluginStatus, networkTracerStatus ModelPluginStatus, ) *ModelAgentPluginsStatus`

NewModelAgentPluginsStatus instantiates a new ModelAgentPluginsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAgentPluginsStatusWithDefaults

`func NewModelAgentPluginsStatusWithDefaults() *ModelAgentPluginsStatus`

NewModelAgentPluginsStatusWithDefaults instantiates a new ModelAgentPluginsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemTracerStatus

`func (o *ModelAgentPluginsStatus) GetFilesystemTracerStatus() ModelPluginStatus`

GetFilesystemTracerStatus returns the FilesystemTracerStatus field if non-nil, zero value otherwise.

### GetFilesystemTracerStatusOk

`func (o *ModelAgentPluginsStatus) GetFilesystemTracerStatusOk() (*ModelPluginStatus, bool)`

GetFilesystemTracerStatusOk returns a tuple with the FilesystemTracerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemTracerStatus

`func (o *ModelAgentPluginsStatus) SetFilesystemTracerStatus(v ModelPluginStatus)`

SetFilesystemTracerStatus sets FilesystemTracerStatus field to given value.


### GetNetworkFilterStatus

`func (o *ModelAgentPluginsStatus) GetNetworkFilterStatus() ModelPluginStatus`

GetNetworkFilterStatus returns the NetworkFilterStatus field if non-nil, zero value otherwise.

### GetNetworkFilterStatusOk

`func (o *ModelAgentPluginsStatus) GetNetworkFilterStatusOk() (*ModelPluginStatus, bool)`

GetNetworkFilterStatusOk returns a tuple with the NetworkFilterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilterStatus

`func (o *ModelAgentPluginsStatus) SetNetworkFilterStatus(v ModelPluginStatus)`

SetNetworkFilterStatus sets NetworkFilterStatus field to given value.


### GetNetworkTracerStatus

`func (o *ModelAgentPluginsStatus) GetNetworkTracerStatus() ModelPluginStatus`

GetNetworkTracerStatus returns the NetworkTracerStatus field if non-nil, zero value otherwise.

### GetNetworkTracerStatusOk

`func (o *ModelAgentPluginsStatus) GetNetworkTracerStatusOk() (*ModelPluginStatus, bool)`

GetNetworkTracerStatusOk returns a tuple with the NetworkTracerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTracerStatus

`func (o *ModelAgentPluginsStatus) SetNetworkTracerStatus(v ModelPluginStatus)`

SetNetworkTracerStatus sets NetworkTracerStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


