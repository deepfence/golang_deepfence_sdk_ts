# ModelCloudNodeAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CloudNetworkTracerStatus** | Pointer to [**ModelPluginStatus**](ModelPluginStatus.md) |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CompliancePercentage** | Pointer to **float32** |  | [optional] 
**HostNodeId** | Pointer to **string** |  | [optional] 
**LastScanId** | Pointer to **string** |  | [optional] 
**LastScanStatus** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**RefreshMessage** | Pointer to **string** |  | [optional] 
**RefreshMetadata** | Pointer to **string** |  | [optional] 
**RefreshStatus** | Pointer to **string** |  | [optional] 
**RefreshStatusMap** | Pointer to **map[string]int32** |  | [optional] 
**ScanStatusMap** | Pointer to **map[string]int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCloudNodeAccountInfo

`func NewModelCloudNodeAccountInfo() *ModelCloudNodeAccountInfo`

NewModelCloudNodeAccountInfo instantiates a new ModelCloudNodeAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudNodeAccountInfoWithDefaults

`func NewModelCloudNodeAccountInfoWithDefaults() *ModelCloudNodeAccountInfo`

NewModelCloudNodeAccountInfoWithDefaults instantiates a new ModelCloudNodeAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *ModelCloudNodeAccountInfo) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ModelCloudNodeAccountInfo) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ModelCloudNodeAccountInfo) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ModelCloudNodeAccountInfo) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *ModelCloudNodeAccountInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ModelCloudNodeAccountInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ModelCloudNodeAccountInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ModelCloudNodeAccountInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCloudNetworkTracerStatus

`func (o *ModelCloudNodeAccountInfo) GetCloudNetworkTracerStatus() ModelPluginStatus`

GetCloudNetworkTracerStatus returns the CloudNetworkTracerStatus field if non-nil, zero value otherwise.

### GetCloudNetworkTracerStatusOk

`func (o *ModelCloudNodeAccountInfo) GetCloudNetworkTracerStatusOk() (*ModelPluginStatus, bool)`

GetCloudNetworkTracerStatusOk returns a tuple with the CloudNetworkTracerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkTracerStatus

`func (o *ModelCloudNodeAccountInfo) SetCloudNetworkTracerStatus(v ModelPluginStatus)`

SetCloudNetworkTracerStatus sets CloudNetworkTracerStatus field to given value.

### HasCloudNetworkTracerStatus

`func (o *ModelCloudNodeAccountInfo) HasCloudNetworkTracerStatus() bool`

HasCloudNetworkTracerStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ModelCloudNodeAccountInfo) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudNodeAccountInfo) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudNodeAccountInfo) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ModelCloudNodeAccountInfo) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompliancePercentage

`func (o *ModelCloudNodeAccountInfo) GetCompliancePercentage() float32`

GetCompliancePercentage returns the CompliancePercentage field if non-nil, zero value otherwise.

### GetCompliancePercentageOk

`func (o *ModelCloudNodeAccountInfo) GetCompliancePercentageOk() (*float32, bool)`

GetCompliancePercentageOk returns a tuple with the CompliancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancePercentage

`func (o *ModelCloudNodeAccountInfo) SetCompliancePercentage(v float32)`

SetCompliancePercentage sets CompliancePercentage field to given value.

### HasCompliancePercentage

`func (o *ModelCloudNodeAccountInfo) HasCompliancePercentage() bool`

HasCompliancePercentage returns a boolean if a field has been set.

### GetHostNodeId

`func (o *ModelCloudNodeAccountInfo) GetHostNodeId() string`

GetHostNodeId returns the HostNodeId field if non-nil, zero value otherwise.

### GetHostNodeIdOk

`func (o *ModelCloudNodeAccountInfo) GetHostNodeIdOk() (*string, bool)`

GetHostNodeIdOk returns a tuple with the HostNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNodeId

`func (o *ModelCloudNodeAccountInfo) SetHostNodeId(v string)`

SetHostNodeId sets HostNodeId field to given value.

### HasHostNodeId

`func (o *ModelCloudNodeAccountInfo) HasHostNodeId() bool`

HasHostNodeId returns a boolean if a field has been set.

### GetLastScanId

`func (o *ModelCloudNodeAccountInfo) GetLastScanId() string`

GetLastScanId returns the LastScanId field if non-nil, zero value otherwise.

### GetLastScanIdOk

`func (o *ModelCloudNodeAccountInfo) GetLastScanIdOk() (*string, bool)`

GetLastScanIdOk returns a tuple with the LastScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanId

`func (o *ModelCloudNodeAccountInfo) SetLastScanId(v string)`

SetLastScanId sets LastScanId field to given value.

### HasLastScanId

`func (o *ModelCloudNodeAccountInfo) HasLastScanId() bool`

HasLastScanId returns a boolean if a field has been set.

### GetLastScanStatus

`func (o *ModelCloudNodeAccountInfo) GetLastScanStatus() string`

GetLastScanStatus returns the LastScanStatus field if non-nil, zero value otherwise.

### GetLastScanStatusOk

`func (o *ModelCloudNodeAccountInfo) GetLastScanStatusOk() (*string, bool)`

GetLastScanStatusOk returns a tuple with the LastScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanStatus

`func (o *ModelCloudNodeAccountInfo) SetLastScanStatus(v string)`

SetLastScanStatus sets LastScanStatus field to given value.

### HasLastScanStatus

`func (o *ModelCloudNodeAccountInfo) HasLastScanStatus() bool`

HasLastScanStatus returns a boolean if a field has been set.

### GetNodeId

`func (o *ModelCloudNodeAccountInfo) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudNodeAccountInfo) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudNodeAccountInfo) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelCloudNodeAccountInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeName

`func (o *ModelCloudNodeAccountInfo) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelCloudNodeAccountInfo) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelCloudNodeAccountInfo) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *ModelCloudNodeAccountInfo) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetRefreshMessage

`func (o *ModelCloudNodeAccountInfo) GetRefreshMessage() string`

GetRefreshMessage returns the RefreshMessage field if non-nil, zero value otherwise.

### GetRefreshMessageOk

`func (o *ModelCloudNodeAccountInfo) GetRefreshMessageOk() (*string, bool)`

GetRefreshMessageOk returns a tuple with the RefreshMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshMessage

`func (o *ModelCloudNodeAccountInfo) SetRefreshMessage(v string)`

SetRefreshMessage sets RefreshMessage field to given value.

### HasRefreshMessage

`func (o *ModelCloudNodeAccountInfo) HasRefreshMessage() bool`

HasRefreshMessage returns a boolean if a field has been set.

### GetRefreshMetadata

`func (o *ModelCloudNodeAccountInfo) GetRefreshMetadata() string`

GetRefreshMetadata returns the RefreshMetadata field if non-nil, zero value otherwise.

### GetRefreshMetadataOk

`func (o *ModelCloudNodeAccountInfo) GetRefreshMetadataOk() (*string, bool)`

GetRefreshMetadataOk returns a tuple with the RefreshMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshMetadata

`func (o *ModelCloudNodeAccountInfo) SetRefreshMetadata(v string)`

SetRefreshMetadata sets RefreshMetadata field to given value.

### HasRefreshMetadata

`func (o *ModelCloudNodeAccountInfo) HasRefreshMetadata() bool`

HasRefreshMetadata returns a boolean if a field has been set.

### GetRefreshStatus

`func (o *ModelCloudNodeAccountInfo) GetRefreshStatus() string`

GetRefreshStatus returns the RefreshStatus field if non-nil, zero value otherwise.

### GetRefreshStatusOk

`func (o *ModelCloudNodeAccountInfo) GetRefreshStatusOk() (*string, bool)`

GetRefreshStatusOk returns a tuple with the RefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatus

`func (o *ModelCloudNodeAccountInfo) SetRefreshStatus(v string)`

SetRefreshStatus sets RefreshStatus field to given value.

### HasRefreshStatus

`func (o *ModelCloudNodeAccountInfo) HasRefreshStatus() bool`

HasRefreshStatus returns a boolean if a field has been set.

### GetRefreshStatusMap

`func (o *ModelCloudNodeAccountInfo) GetRefreshStatusMap() map[string]int32`

GetRefreshStatusMap returns the RefreshStatusMap field if non-nil, zero value otherwise.

### GetRefreshStatusMapOk

`func (o *ModelCloudNodeAccountInfo) GetRefreshStatusMapOk() (*map[string]int32, bool)`

GetRefreshStatusMapOk returns a tuple with the RefreshStatusMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatusMap

`func (o *ModelCloudNodeAccountInfo) SetRefreshStatusMap(v map[string]int32)`

SetRefreshStatusMap sets RefreshStatusMap field to given value.

### HasRefreshStatusMap

`func (o *ModelCloudNodeAccountInfo) HasRefreshStatusMap() bool`

HasRefreshStatusMap returns a boolean if a field has been set.

### SetRefreshStatusMapNil

`func (o *ModelCloudNodeAccountInfo) SetRefreshStatusMapNil(b bool)`

 SetRefreshStatusMapNil sets the value for RefreshStatusMap to be an explicit nil

### UnsetRefreshStatusMap
`func (o *ModelCloudNodeAccountInfo) UnsetRefreshStatusMap()`

UnsetRefreshStatusMap ensures that no value is present for RefreshStatusMap, not even an explicit nil
### GetScanStatusMap

`func (o *ModelCloudNodeAccountInfo) GetScanStatusMap() map[string]int32`

GetScanStatusMap returns the ScanStatusMap field if non-nil, zero value otherwise.

### GetScanStatusMapOk

`func (o *ModelCloudNodeAccountInfo) GetScanStatusMapOk() (*map[string]int32, bool)`

GetScanStatusMapOk returns a tuple with the ScanStatusMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatusMap

`func (o *ModelCloudNodeAccountInfo) SetScanStatusMap(v map[string]int32)`

SetScanStatusMap sets ScanStatusMap field to given value.

### HasScanStatusMap

`func (o *ModelCloudNodeAccountInfo) HasScanStatusMap() bool`

HasScanStatusMap returns a boolean if a field has been set.

### SetScanStatusMapNil

`func (o *ModelCloudNodeAccountInfo) SetScanStatusMapNil(b bool)`

 SetScanStatusMapNil sets the value for ScanStatusMap to be an explicit nil

### UnsetScanStatusMap
`func (o *ModelCloudNodeAccountInfo) UnsetScanStatusMap()`

UnsetScanStatusMap ensures that no value is present for ScanStatusMap, not even an explicit nil
### GetVersion

`func (o *ModelCloudNodeAccountInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelCloudNodeAccountInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelCloudNodeAccountInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelCloudNodeAccountInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


