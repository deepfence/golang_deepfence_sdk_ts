# ModelAgentInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** |  | 
**CloudScannerId** | [**ModelAgentID**](ModelAgentID.md) |  | 
**RegionIds** | [**[]ModelRegionIDs**](ModelRegionIDs.md) |  | 

## Methods

### NewModelAgentInstall

`func NewModelAgentInstall(cloudProvider string, cloudScannerId ModelAgentID, regionIds []ModelRegionIDs, ) *ModelAgentInstall`

NewModelAgentInstall instantiates a new ModelAgentInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAgentInstallWithDefaults

`func NewModelAgentInstallWithDefaults() *ModelAgentInstall`

NewModelAgentInstallWithDefaults instantiates a new ModelAgentInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ModelAgentInstall) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelAgentInstall) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelAgentInstall) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCloudScannerId

`func (o *ModelAgentInstall) GetCloudScannerId() ModelAgentID`

GetCloudScannerId returns the CloudScannerId field if non-nil, zero value otherwise.

### GetCloudScannerIdOk

`func (o *ModelAgentInstall) GetCloudScannerIdOk() (*ModelAgentID, bool)`

GetCloudScannerIdOk returns a tuple with the CloudScannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudScannerId

`func (o *ModelAgentInstall) SetCloudScannerId(v ModelAgentID)`

SetCloudScannerId sets CloudScannerId field to given value.


### GetRegionIds

`func (o *ModelAgentInstall) GetRegionIds() []ModelRegionIDs`

GetRegionIds returns the RegionIds field if non-nil, zero value otherwise.

### GetRegionIdsOk

`func (o *ModelAgentInstall) GetRegionIdsOk() (*[]ModelRegionIDs, bool)`

GetRegionIdsOk returns a tuple with the RegionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIds

`func (o *ModelAgentInstall) SetRegionIds(v []ModelRegionIDs)`

SetRegionIds sets RegionIds field to given value.


### SetRegionIdsNil

`func (o *ModelAgentInstall) SetRegionIdsNil(b bool)`

 SetRegionIdsNil sets the value for RegionIds to be an explicit nil

### UnsetRegionIds
`func (o *ModelAgentInstall) UnsetRegionIds()`

UnsetRegionIds ensures that no value is present for RegionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


