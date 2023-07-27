# ControlsAgentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemTracerConfig** | [**ControlsFilesystemTracerConfig**](ControlsFilesystemTracerConfig.md) |  | 
**Name** | **string** |  | 
**NetworkTracerConfig** | [**ControlsNetworkTracerConfig**](ControlsNetworkTracerConfig.md) |  | 
**PolicyTracerConfig** | [**ControlsPolicyFilterConfig**](ControlsPolicyFilterConfig.md) |  | 

## Methods

### NewControlsAgentConfig

`func NewControlsAgentConfig(filesystemTracerConfig ControlsFilesystemTracerConfig, name string, networkTracerConfig ControlsNetworkTracerConfig, policyTracerConfig ControlsPolicyFilterConfig, ) *ControlsAgentConfig`

NewControlsAgentConfig instantiates a new ControlsAgentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsAgentConfigWithDefaults

`func NewControlsAgentConfigWithDefaults() *ControlsAgentConfig`

NewControlsAgentConfigWithDefaults instantiates a new ControlsAgentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemTracerConfig

`func (o *ControlsAgentConfig) GetFilesystemTracerConfig() ControlsFilesystemTracerConfig`

GetFilesystemTracerConfig returns the FilesystemTracerConfig field if non-nil, zero value otherwise.

### GetFilesystemTracerConfigOk

`func (o *ControlsAgentConfig) GetFilesystemTracerConfigOk() (*ControlsFilesystemTracerConfig, bool)`

GetFilesystemTracerConfigOk returns a tuple with the FilesystemTracerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemTracerConfig

`func (o *ControlsAgentConfig) SetFilesystemTracerConfig(v ControlsFilesystemTracerConfig)`

SetFilesystemTracerConfig sets FilesystemTracerConfig field to given value.


### GetName

`func (o *ControlsAgentConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsAgentConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsAgentConfig) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkTracerConfig

`func (o *ControlsAgentConfig) GetNetworkTracerConfig() ControlsNetworkTracerConfig`

GetNetworkTracerConfig returns the NetworkTracerConfig field if non-nil, zero value otherwise.

### GetNetworkTracerConfigOk

`func (o *ControlsAgentConfig) GetNetworkTracerConfigOk() (*ControlsNetworkTracerConfig, bool)`

GetNetworkTracerConfigOk returns a tuple with the NetworkTracerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTracerConfig

`func (o *ControlsAgentConfig) SetNetworkTracerConfig(v ControlsNetworkTracerConfig)`

SetNetworkTracerConfig sets NetworkTracerConfig field to given value.


### GetPolicyTracerConfig

`func (o *ControlsAgentConfig) GetPolicyTracerConfig() ControlsPolicyFilterConfig`

GetPolicyTracerConfig returns the PolicyTracerConfig field if non-nil, zero value otherwise.

### GetPolicyTracerConfigOk

`func (o *ControlsAgentConfig) GetPolicyTracerConfigOk() (*ControlsPolicyFilterConfig, bool)`

GetPolicyTracerConfigOk returns a tuple with the PolicyTracerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTracerConfig

`func (o *ControlsAgentConfig) SetPolicyTracerConfig(v ControlsPolicyFilterConfig)`

SetPolicyTracerConfig sets PolicyTracerConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


