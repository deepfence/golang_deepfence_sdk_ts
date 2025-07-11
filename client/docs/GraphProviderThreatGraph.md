# GraphProviderThreatGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsCount** | **int32** |  | 
**CloudComplianceCount** | **int32** |  | 
**CloudWarnAlarmCount** | **int32** |  | 
**ComplianceCount** | **int32** |  | 
**ExploitableSecretsCount** | **int32** |  | 
**ExploitableVulnerabilitiesCount** | **int32** |  | 
**Resources** | [**[]GraphThreatNodeInfo**](GraphThreatNodeInfo.md) |  | 
**SecretsCount** | **int32** |  | 
**VulnerabilityCount** | **int32** |  | 
**WarnAlarmCount** | **int32** |  | 

## Methods

### NewGraphProviderThreatGraph

`func NewGraphProviderThreatGraph(alertsCount int32, cloudComplianceCount int32, cloudWarnAlarmCount int32, complianceCount int32, exploitableSecretsCount int32, exploitableVulnerabilitiesCount int32, resources []GraphThreatNodeInfo, secretsCount int32, vulnerabilityCount int32, warnAlarmCount int32, ) *GraphProviderThreatGraph`

NewGraphProviderThreatGraph instantiates a new GraphProviderThreatGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphProviderThreatGraphWithDefaults

`func NewGraphProviderThreatGraphWithDefaults() *GraphProviderThreatGraph`

NewGraphProviderThreatGraphWithDefaults instantiates a new GraphProviderThreatGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsCount

`func (o *GraphProviderThreatGraph) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *GraphProviderThreatGraph) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *GraphProviderThreatGraph) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetCloudComplianceCount

`func (o *GraphProviderThreatGraph) GetCloudComplianceCount() int32`

GetCloudComplianceCount returns the CloudComplianceCount field if non-nil, zero value otherwise.

### GetCloudComplianceCountOk

`func (o *GraphProviderThreatGraph) GetCloudComplianceCountOk() (*int32, bool)`

GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceCount

`func (o *GraphProviderThreatGraph) SetCloudComplianceCount(v int32)`

SetCloudComplianceCount sets CloudComplianceCount field to given value.


### GetCloudWarnAlarmCount

`func (o *GraphProviderThreatGraph) GetCloudWarnAlarmCount() int32`

GetCloudWarnAlarmCount returns the CloudWarnAlarmCount field if non-nil, zero value otherwise.

### GetCloudWarnAlarmCountOk

`func (o *GraphProviderThreatGraph) GetCloudWarnAlarmCountOk() (*int32, bool)`

GetCloudWarnAlarmCountOk returns a tuple with the CloudWarnAlarmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWarnAlarmCount

`func (o *GraphProviderThreatGraph) SetCloudWarnAlarmCount(v int32)`

SetCloudWarnAlarmCount sets CloudWarnAlarmCount field to given value.


### GetComplianceCount

`func (o *GraphProviderThreatGraph) GetComplianceCount() int32`

GetComplianceCount returns the ComplianceCount field if non-nil, zero value otherwise.

### GetComplianceCountOk

`func (o *GraphProviderThreatGraph) GetComplianceCountOk() (*int32, bool)`

GetComplianceCountOk returns a tuple with the ComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCount

`func (o *GraphProviderThreatGraph) SetComplianceCount(v int32)`

SetComplianceCount sets ComplianceCount field to given value.


### GetExploitableSecretsCount

`func (o *GraphProviderThreatGraph) GetExploitableSecretsCount() int32`

GetExploitableSecretsCount returns the ExploitableSecretsCount field if non-nil, zero value otherwise.

### GetExploitableSecretsCountOk

`func (o *GraphProviderThreatGraph) GetExploitableSecretsCountOk() (*int32, bool)`

GetExploitableSecretsCountOk returns a tuple with the ExploitableSecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitableSecretsCount

`func (o *GraphProviderThreatGraph) SetExploitableSecretsCount(v int32)`

SetExploitableSecretsCount sets ExploitableSecretsCount field to given value.


### GetExploitableVulnerabilitiesCount

`func (o *GraphProviderThreatGraph) GetExploitableVulnerabilitiesCount() int32`

GetExploitableVulnerabilitiesCount returns the ExploitableVulnerabilitiesCount field if non-nil, zero value otherwise.

### GetExploitableVulnerabilitiesCountOk

`func (o *GraphProviderThreatGraph) GetExploitableVulnerabilitiesCountOk() (*int32, bool)`

GetExploitableVulnerabilitiesCountOk returns a tuple with the ExploitableVulnerabilitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitableVulnerabilitiesCount

`func (o *GraphProviderThreatGraph) SetExploitableVulnerabilitiesCount(v int32)`

SetExploitableVulnerabilitiesCount sets ExploitableVulnerabilitiesCount field to given value.


### GetResources

`func (o *GraphProviderThreatGraph) GetResources() []GraphThreatNodeInfo`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GraphProviderThreatGraph) GetResourcesOk() (*[]GraphThreatNodeInfo, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GraphProviderThreatGraph) SetResources(v []GraphThreatNodeInfo)`

SetResources sets Resources field to given value.


### SetResourcesNil

`func (o *GraphProviderThreatGraph) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *GraphProviderThreatGraph) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetSecretsCount

`func (o *GraphProviderThreatGraph) GetSecretsCount() int32`

GetSecretsCount returns the SecretsCount field if non-nil, zero value otherwise.

### GetSecretsCountOk

`func (o *GraphProviderThreatGraph) GetSecretsCountOk() (*int32, bool)`

GetSecretsCountOk returns a tuple with the SecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsCount

`func (o *GraphProviderThreatGraph) SetSecretsCount(v int32)`

SetSecretsCount sets SecretsCount field to given value.


### GetVulnerabilityCount

`func (o *GraphProviderThreatGraph) GetVulnerabilityCount() int32`

GetVulnerabilityCount returns the VulnerabilityCount field if non-nil, zero value otherwise.

### GetVulnerabilityCountOk

`func (o *GraphProviderThreatGraph) GetVulnerabilityCountOk() (*int32, bool)`

GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCount

`func (o *GraphProviderThreatGraph) SetVulnerabilityCount(v int32)`

SetVulnerabilityCount sets VulnerabilityCount field to given value.


### GetWarnAlarmCount

`func (o *GraphProviderThreatGraph) GetWarnAlarmCount() int32`

GetWarnAlarmCount returns the WarnAlarmCount field if non-nil, zero value otherwise.

### GetWarnAlarmCountOk

`func (o *GraphProviderThreatGraph) GetWarnAlarmCountOk() (*int32, bool)`

GetWarnAlarmCountOk returns a tuple with the WarnAlarmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnAlarmCount

`func (o *GraphProviderThreatGraph) SetWarnAlarmCount(v int32)`

SetWarnAlarmCount sets WarnAlarmCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


