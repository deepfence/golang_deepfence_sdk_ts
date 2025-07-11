# GraphThreatNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsCount** | **int32** |  | 
**AttackPath** | **[][]string** |  | 
**CloudComplianceCount** | **int32** |  | 
**CloudWarnAlarmCount** | **int32** |  | 
**ComplianceCount** | **int32** |  | 
**Count** | **int32** |  | 
**ExploitableSecretsCount** | **int32** |  | 
**ExploitableVulnerabilitiesCount** | **int32** |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**NodeType** | **string** |  | 
**Nodes** | [**map[string]GraphNodeInfo**](GraphNodeInfo.md) |  | 
**SecretsCount** | **int32** |  | 
**VulnerabilityCount** | **int32** |  | 
**WarnAlarmCount** | **int32** |  | 

## Methods

### NewGraphThreatNodeInfo

`func NewGraphThreatNodeInfo(alertsCount int32, attackPath [][]string, cloudComplianceCount int32, cloudWarnAlarmCount int32, complianceCount int32, count int32, exploitableSecretsCount int32, exploitableVulnerabilitiesCount int32, id string, label string, nodeType string, nodes map[string]GraphNodeInfo, secretsCount int32, vulnerabilityCount int32, warnAlarmCount int32, ) *GraphThreatNodeInfo`

NewGraphThreatNodeInfo instantiates a new GraphThreatNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphThreatNodeInfoWithDefaults

`func NewGraphThreatNodeInfoWithDefaults() *GraphThreatNodeInfo`

NewGraphThreatNodeInfoWithDefaults instantiates a new GraphThreatNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsCount

`func (o *GraphThreatNodeInfo) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *GraphThreatNodeInfo) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *GraphThreatNodeInfo) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetAttackPath

`func (o *GraphThreatNodeInfo) GetAttackPath() [][]string`

GetAttackPath returns the AttackPath field if non-nil, zero value otherwise.

### GetAttackPathOk

`func (o *GraphThreatNodeInfo) GetAttackPathOk() (*[][]string, bool)`

GetAttackPathOk returns a tuple with the AttackPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPath

`func (o *GraphThreatNodeInfo) SetAttackPath(v [][]string)`

SetAttackPath sets AttackPath field to given value.


### SetAttackPathNil

`func (o *GraphThreatNodeInfo) SetAttackPathNil(b bool)`

 SetAttackPathNil sets the value for AttackPath to be an explicit nil

### UnsetAttackPath
`func (o *GraphThreatNodeInfo) UnsetAttackPath()`

UnsetAttackPath ensures that no value is present for AttackPath, not even an explicit nil
### GetCloudComplianceCount

`func (o *GraphThreatNodeInfo) GetCloudComplianceCount() int32`

GetCloudComplianceCount returns the CloudComplianceCount field if non-nil, zero value otherwise.

### GetCloudComplianceCountOk

`func (o *GraphThreatNodeInfo) GetCloudComplianceCountOk() (*int32, bool)`

GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceCount

`func (o *GraphThreatNodeInfo) SetCloudComplianceCount(v int32)`

SetCloudComplianceCount sets CloudComplianceCount field to given value.


### GetCloudWarnAlarmCount

`func (o *GraphThreatNodeInfo) GetCloudWarnAlarmCount() int32`

GetCloudWarnAlarmCount returns the CloudWarnAlarmCount field if non-nil, zero value otherwise.

### GetCloudWarnAlarmCountOk

`func (o *GraphThreatNodeInfo) GetCloudWarnAlarmCountOk() (*int32, bool)`

GetCloudWarnAlarmCountOk returns a tuple with the CloudWarnAlarmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWarnAlarmCount

`func (o *GraphThreatNodeInfo) SetCloudWarnAlarmCount(v int32)`

SetCloudWarnAlarmCount sets CloudWarnAlarmCount field to given value.


### GetComplianceCount

`func (o *GraphThreatNodeInfo) GetComplianceCount() int32`

GetComplianceCount returns the ComplianceCount field if non-nil, zero value otherwise.

### GetComplianceCountOk

`func (o *GraphThreatNodeInfo) GetComplianceCountOk() (*int32, bool)`

GetComplianceCountOk returns a tuple with the ComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCount

`func (o *GraphThreatNodeInfo) SetComplianceCount(v int32)`

SetComplianceCount sets ComplianceCount field to given value.


### GetCount

`func (o *GraphThreatNodeInfo) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GraphThreatNodeInfo) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GraphThreatNodeInfo) SetCount(v int32)`

SetCount sets Count field to given value.


### GetExploitableSecretsCount

`func (o *GraphThreatNodeInfo) GetExploitableSecretsCount() int32`

GetExploitableSecretsCount returns the ExploitableSecretsCount field if non-nil, zero value otherwise.

### GetExploitableSecretsCountOk

`func (o *GraphThreatNodeInfo) GetExploitableSecretsCountOk() (*int32, bool)`

GetExploitableSecretsCountOk returns a tuple with the ExploitableSecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitableSecretsCount

`func (o *GraphThreatNodeInfo) SetExploitableSecretsCount(v int32)`

SetExploitableSecretsCount sets ExploitableSecretsCount field to given value.


### GetExploitableVulnerabilitiesCount

`func (o *GraphThreatNodeInfo) GetExploitableVulnerabilitiesCount() int32`

GetExploitableVulnerabilitiesCount returns the ExploitableVulnerabilitiesCount field if non-nil, zero value otherwise.

### GetExploitableVulnerabilitiesCountOk

`func (o *GraphThreatNodeInfo) GetExploitableVulnerabilitiesCountOk() (*int32, bool)`

GetExploitableVulnerabilitiesCountOk returns a tuple with the ExploitableVulnerabilitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitableVulnerabilitiesCount

`func (o *GraphThreatNodeInfo) SetExploitableVulnerabilitiesCount(v int32)`

SetExploitableVulnerabilitiesCount sets ExploitableVulnerabilitiesCount field to given value.


### GetId

`func (o *GraphThreatNodeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphThreatNodeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphThreatNodeInfo) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *GraphThreatNodeInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GraphThreatNodeInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GraphThreatNodeInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNodeType

`func (o *GraphThreatNodeInfo) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *GraphThreatNodeInfo) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *GraphThreatNodeInfo) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetNodes

`func (o *GraphThreatNodeInfo) GetNodes() map[string]GraphNodeInfo`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GraphThreatNodeInfo) GetNodesOk() (*map[string]GraphNodeInfo, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GraphThreatNodeInfo) SetNodes(v map[string]GraphNodeInfo)`

SetNodes sets Nodes field to given value.


### SetNodesNil

`func (o *GraphThreatNodeInfo) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *GraphThreatNodeInfo) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetSecretsCount

`func (o *GraphThreatNodeInfo) GetSecretsCount() int32`

GetSecretsCount returns the SecretsCount field if non-nil, zero value otherwise.

### GetSecretsCountOk

`func (o *GraphThreatNodeInfo) GetSecretsCountOk() (*int32, bool)`

GetSecretsCountOk returns a tuple with the SecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsCount

`func (o *GraphThreatNodeInfo) SetSecretsCount(v int32)`

SetSecretsCount sets SecretsCount field to given value.


### GetVulnerabilityCount

`func (o *GraphThreatNodeInfo) GetVulnerabilityCount() int32`

GetVulnerabilityCount returns the VulnerabilityCount field if non-nil, zero value otherwise.

### GetVulnerabilityCountOk

`func (o *GraphThreatNodeInfo) GetVulnerabilityCountOk() (*int32, bool)`

GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCount

`func (o *GraphThreatNodeInfo) SetVulnerabilityCount(v int32)`

SetVulnerabilityCount sets VulnerabilityCount field to given value.


### GetWarnAlarmCount

`func (o *GraphThreatNodeInfo) GetWarnAlarmCount() int32`

GetWarnAlarmCount returns the WarnAlarmCount field if non-nil, zero value otherwise.

### GetWarnAlarmCountOk

`func (o *GraphThreatNodeInfo) GetWarnAlarmCountOk() (*int32, bool)`

GetWarnAlarmCountOk returns a tuple with the WarnAlarmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnAlarmCount

`func (o *GraphThreatNodeInfo) SetWarnAlarmCount(v int32)`

SetWarnAlarmCount sets WarnAlarmCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


