# ControlsThreatIntelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPostureControlsHash** | **string** |  | 
**CloudPostureControlsUrl** | **string** |  | 
**IgnoredAlertRuleIds** | **[]string** |  | 
**InternalIps** | **[]string** |  | 
**MalwareScannerRulesHash** | **string** |  | 
**MalwareScannerRulesUrl** | **string** |  | 
**NetworkAlertRulesUrl** | **string** |  | 
**RulesHash** | **string** |  | 
**SecretScannerRulesHash** | **string** |  | 
**SecretScannerRulesUrl** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewControlsThreatIntelInfo

`func NewControlsThreatIntelInfo(cloudPostureControlsHash string, cloudPostureControlsUrl string, ignoredAlertRuleIds []string, internalIps []string, malwareScannerRulesHash string, malwareScannerRulesUrl string, networkAlertRulesUrl string, rulesHash string, secretScannerRulesHash string, secretScannerRulesUrl string, updatedAt int32, ) *ControlsThreatIntelInfo`

NewControlsThreatIntelInfo instantiates a new ControlsThreatIntelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsThreatIntelInfoWithDefaults

`func NewControlsThreatIntelInfoWithDefaults() *ControlsThreatIntelInfo`

NewControlsThreatIntelInfoWithDefaults instantiates a new ControlsThreatIntelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPostureControlsHash

`func (o *ControlsThreatIntelInfo) GetCloudPostureControlsHash() string`

GetCloudPostureControlsHash returns the CloudPostureControlsHash field if non-nil, zero value otherwise.

### GetCloudPostureControlsHashOk

`func (o *ControlsThreatIntelInfo) GetCloudPostureControlsHashOk() (*string, bool)`

GetCloudPostureControlsHashOk returns a tuple with the CloudPostureControlsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPostureControlsHash

`func (o *ControlsThreatIntelInfo) SetCloudPostureControlsHash(v string)`

SetCloudPostureControlsHash sets CloudPostureControlsHash field to given value.


### GetCloudPostureControlsUrl

`func (o *ControlsThreatIntelInfo) GetCloudPostureControlsUrl() string`

GetCloudPostureControlsUrl returns the CloudPostureControlsUrl field if non-nil, zero value otherwise.

### GetCloudPostureControlsUrlOk

`func (o *ControlsThreatIntelInfo) GetCloudPostureControlsUrlOk() (*string, bool)`

GetCloudPostureControlsUrlOk returns a tuple with the CloudPostureControlsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPostureControlsUrl

`func (o *ControlsThreatIntelInfo) SetCloudPostureControlsUrl(v string)`

SetCloudPostureControlsUrl sets CloudPostureControlsUrl field to given value.


### GetIgnoredAlertRuleIds

`func (o *ControlsThreatIntelInfo) GetIgnoredAlertRuleIds() []string`

GetIgnoredAlertRuleIds returns the IgnoredAlertRuleIds field if non-nil, zero value otherwise.

### GetIgnoredAlertRuleIdsOk

`func (o *ControlsThreatIntelInfo) GetIgnoredAlertRuleIdsOk() (*[]string, bool)`

GetIgnoredAlertRuleIdsOk returns a tuple with the IgnoredAlertRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredAlertRuleIds

`func (o *ControlsThreatIntelInfo) SetIgnoredAlertRuleIds(v []string)`

SetIgnoredAlertRuleIds sets IgnoredAlertRuleIds field to given value.


### SetIgnoredAlertRuleIdsNil

`func (o *ControlsThreatIntelInfo) SetIgnoredAlertRuleIdsNil(b bool)`

 SetIgnoredAlertRuleIdsNil sets the value for IgnoredAlertRuleIds to be an explicit nil

### UnsetIgnoredAlertRuleIds
`func (o *ControlsThreatIntelInfo) UnsetIgnoredAlertRuleIds()`

UnsetIgnoredAlertRuleIds ensures that no value is present for IgnoredAlertRuleIds, not even an explicit nil
### GetInternalIps

`func (o *ControlsThreatIntelInfo) GetInternalIps() []string`

GetInternalIps returns the InternalIps field if non-nil, zero value otherwise.

### GetInternalIpsOk

`func (o *ControlsThreatIntelInfo) GetInternalIpsOk() (*[]string, bool)`

GetInternalIpsOk returns a tuple with the InternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIps

`func (o *ControlsThreatIntelInfo) SetInternalIps(v []string)`

SetInternalIps sets InternalIps field to given value.


### SetInternalIpsNil

`func (o *ControlsThreatIntelInfo) SetInternalIpsNil(b bool)`

 SetInternalIpsNil sets the value for InternalIps to be an explicit nil

### UnsetInternalIps
`func (o *ControlsThreatIntelInfo) UnsetInternalIps()`

UnsetInternalIps ensures that no value is present for InternalIps, not even an explicit nil
### GetMalwareScannerRulesHash

`func (o *ControlsThreatIntelInfo) GetMalwareScannerRulesHash() string`

GetMalwareScannerRulesHash returns the MalwareScannerRulesHash field if non-nil, zero value otherwise.

### GetMalwareScannerRulesHashOk

`func (o *ControlsThreatIntelInfo) GetMalwareScannerRulesHashOk() (*string, bool)`

GetMalwareScannerRulesHashOk returns a tuple with the MalwareScannerRulesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareScannerRulesHash

`func (o *ControlsThreatIntelInfo) SetMalwareScannerRulesHash(v string)`

SetMalwareScannerRulesHash sets MalwareScannerRulesHash field to given value.


### GetMalwareScannerRulesUrl

`func (o *ControlsThreatIntelInfo) GetMalwareScannerRulesUrl() string`

GetMalwareScannerRulesUrl returns the MalwareScannerRulesUrl field if non-nil, zero value otherwise.

### GetMalwareScannerRulesUrlOk

`func (o *ControlsThreatIntelInfo) GetMalwareScannerRulesUrlOk() (*string, bool)`

GetMalwareScannerRulesUrlOk returns a tuple with the MalwareScannerRulesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareScannerRulesUrl

`func (o *ControlsThreatIntelInfo) SetMalwareScannerRulesUrl(v string)`

SetMalwareScannerRulesUrl sets MalwareScannerRulesUrl field to given value.


### GetNetworkAlertRulesUrl

`func (o *ControlsThreatIntelInfo) GetNetworkAlertRulesUrl() string`

GetNetworkAlertRulesUrl returns the NetworkAlertRulesUrl field if non-nil, zero value otherwise.

### GetNetworkAlertRulesUrlOk

`func (o *ControlsThreatIntelInfo) GetNetworkAlertRulesUrlOk() (*string, bool)`

GetNetworkAlertRulesUrlOk returns a tuple with the NetworkAlertRulesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAlertRulesUrl

`func (o *ControlsThreatIntelInfo) SetNetworkAlertRulesUrl(v string)`

SetNetworkAlertRulesUrl sets NetworkAlertRulesUrl field to given value.


### GetRulesHash

`func (o *ControlsThreatIntelInfo) GetRulesHash() string`

GetRulesHash returns the RulesHash field if non-nil, zero value otherwise.

### GetRulesHashOk

`func (o *ControlsThreatIntelInfo) GetRulesHashOk() (*string, bool)`

GetRulesHashOk returns a tuple with the RulesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesHash

`func (o *ControlsThreatIntelInfo) SetRulesHash(v string)`

SetRulesHash sets RulesHash field to given value.


### GetSecretScannerRulesHash

`func (o *ControlsThreatIntelInfo) GetSecretScannerRulesHash() string`

GetSecretScannerRulesHash returns the SecretScannerRulesHash field if non-nil, zero value otherwise.

### GetSecretScannerRulesHashOk

`func (o *ControlsThreatIntelInfo) GetSecretScannerRulesHashOk() (*string, bool)`

GetSecretScannerRulesHashOk returns a tuple with the SecretScannerRulesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScannerRulesHash

`func (o *ControlsThreatIntelInfo) SetSecretScannerRulesHash(v string)`

SetSecretScannerRulesHash sets SecretScannerRulesHash field to given value.


### GetSecretScannerRulesUrl

`func (o *ControlsThreatIntelInfo) GetSecretScannerRulesUrl() string`

GetSecretScannerRulesUrl returns the SecretScannerRulesUrl field if non-nil, zero value otherwise.

### GetSecretScannerRulesUrlOk

`func (o *ControlsThreatIntelInfo) GetSecretScannerRulesUrlOk() (*string, bool)`

GetSecretScannerRulesUrlOk returns a tuple with the SecretScannerRulesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScannerRulesUrl

`func (o *ControlsThreatIntelInfo) SetSecretScannerRulesUrl(v string)`

SetSecretScannerRulesUrl sets SecretScannerRulesUrl field to given value.


### GetUpdatedAt

`func (o *ControlsThreatIntelInfo) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsThreatIntelInfo) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsThreatIntelInfo) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


