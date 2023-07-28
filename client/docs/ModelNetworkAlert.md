# ModelNetworkAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProto** | **string** |  | 
**Category** | **string** |  | 
**Classtype** | **string** |  | 
**ContainerName** | **string** |  | 
**Count** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**Description** | **string** |  | 
**DestinationIp** | **string** |  | 
**DestinationPort** | **int32** |  | 
**Direction** | **string** |  | 
**Encrypted** | **string** |  | 
**EventType** | **string** |  | 
**Geoip** | Pointer to [**ModelNetworkAlertGeoip**](ModelNetworkAlertGeoip.md) |  | [optional] 
**Hostname** | **string** |  | 
**Http** | Pointer to [**ModelNetworkAlertHttp**](ModelNetworkAlertHttp.md) |  | [optional] 
**HttpType** | **string** |  | 
**Intent** | **string** |  | 
**Internal** | **string** |  | 
**IpReputation** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**Masked** | **bool** |  | 
**Matched** | **string** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**PodName** | **string** |  | 
**ResourceType** | **string** |  | 
**RuleId** | **int32** |  | 
**ScopeId** | **string** |  | 
**Severity** | **string** |  | 
**SeverityAnomalyContainerName** | **string** |  | 
**SeverityScore** | **float32** |  | 
**SignatureId** | **int32** |  | 
**SourceIp** | **string** |  | 
**SourcePort** | **int32** |  | 
**Summary** | **string** |  | 
**Tags** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelNetworkAlert

`func NewModelNetworkAlert(appProto string, category string, classtype string, containerName string, count int32, createdAt int32, description string, destinationIp string, destinationPort int32, direction string, encrypted string, eventType string, hostname string, httpType string, intent string, internal string, ipReputation string, kubernetesClusterName string, masked bool, matched string, nodeId string, nodeType string, podName string, resourceType string, ruleId int32, scopeId string, severity string, severityAnomalyContainerName string, severityScore float32, signatureId int32, sourceIp string, sourcePort int32, summary string, tags string, updatedAt int32, ) *ModelNetworkAlert`

NewModelNetworkAlert instantiates a new ModelNetworkAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkAlertWithDefaults

`func NewModelNetworkAlertWithDefaults() *ModelNetworkAlert`

NewModelNetworkAlertWithDefaults instantiates a new ModelNetworkAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProto

`func (o *ModelNetworkAlert) GetAppProto() string`

GetAppProto returns the AppProto field if non-nil, zero value otherwise.

### GetAppProtoOk

`func (o *ModelNetworkAlert) GetAppProtoOk() (*string, bool)`

GetAppProtoOk returns a tuple with the AppProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProto

`func (o *ModelNetworkAlert) SetAppProto(v string)`

SetAppProto sets AppProto field to given value.


### GetCategory

`func (o *ModelNetworkAlert) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelNetworkAlert) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelNetworkAlert) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetClasstype

`func (o *ModelNetworkAlert) GetClasstype() string`

GetClasstype returns the Classtype field if non-nil, zero value otherwise.

### GetClasstypeOk

`func (o *ModelNetworkAlert) GetClasstypeOk() (*string, bool)`

GetClasstypeOk returns a tuple with the Classtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasstype

`func (o *ModelNetworkAlert) SetClasstype(v string)`

SetClasstype sets Classtype field to given value.


### GetContainerName

`func (o *ModelNetworkAlert) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ModelNetworkAlert) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ModelNetworkAlert) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetCount

`func (o *ModelNetworkAlert) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelNetworkAlert) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelNetworkAlert) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCreatedAt

`func (o *ModelNetworkAlert) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelNetworkAlert) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelNetworkAlert) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *ModelNetworkAlert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelNetworkAlert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelNetworkAlert) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIp

`func (o *ModelNetworkAlert) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *ModelNetworkAlert) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *ModelNetworkAlert) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.


### GetDestinationPort

`func (o *ModelNetworkAlert) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *ModelNetworkAlert) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *ModelNetworkAlert) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.


### GetDirection

`func (o *ModelNetworkAlert) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModelNetworkAlert) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModelNetworkAlert) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEncrypted

`func (o *ModelNetworkAlert) GetEncrypted() string`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *ModelNetworkAlert) GetEncryptedOk() (*string, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *ModelNetworkAlert) SetEncrypted(v string)`

SetEncrypted sets Encrypted field to given value.


### GetEventType

`func (o *ModelNetworkAlert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ModelNetworkAlert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ModelNetworkAlert) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetGeoip

`func (o *ModelNetworkAlert) GetGeoip() ModelNetworkAlertGeoip`

GetGeoip returns the Geoip field if non-nil, zero value otherwise.

### GetGeoipOk

`func (o *ModelNetworkAlert) GetGeoipOk() (*ModelNetworkAlertGeoip, bool)`

GetGeoipOk returns a tuple with the Geoip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoip

`func (o *ModelNetworkAlert) SetGeoip(v ModelNetworkAlertGeoip)`

SetGeoip sets Geoip field to given value.

### HasGeoip

`func (o *ModelNetworkAlert) HasGeoip() bool`

HasGeoip returns a boolean if a field has been set.

### GetHostname

`func (o *ModelNetworkAlert) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ModelNetworkAlert) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ModelNetworkAlert) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHttp

`func (o *ModelNetworkAlert) GetHttp() ModelNetworkAlertHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ModelNetworkAlert) GetHttpOk() (*ModelNetworkAlertHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ModelNetworkAlert) SetHttp(v ModelNetworkAlertHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ModelNetworkAlert) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetHttpType

`func (o *ModelNetworkAlert) GetHttpType() string`

GetHttpType returns the HttpType field if non-nil, zero value otherwise.

### GetHttpTypeOk

`func (o *ModelNetworkAlert) GetHttpTypeOk() (*string, bool)`

GetHttpTypeOk returns a tuple with the HttpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpType

`func (o *ModelNetworkAlert) SetHttpType(v string)`

SetHttpType sets HttpType field to given value.


### GetIntent

`func (o *ModelNetworkAlert) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *ModelNetworkAlert) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *ModelNetworkAlert) SetIntent(v string)`

SetIntent sets Intent field to given value.


### GetInternal

`func (o *ModelNetworkAlert) GetInternal() string`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ModelNetworkAlert) GetInternalOk() (*string, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ModelNetworkAlert) SetInternal(v string)`

SetInternal sets Internal field to given value.


### GetIpReputation

`func (o *ModelNetworkAlert) GetIpReputation() string`

GetIpReputation returns the IpReputation field if non-nil, zero value otherwise.

### GetIpReputationOk

`func (o *ModelNetworkAlert) GetIpReputationOk() (*string, bool)`

GetIpReputationOk returns a tuple with the IpReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReputation

`func (o *ModelNetworkAlert) SetIpReputation(v string)`

SetIpReputation sets IpReputation field to given value.


### GetKubernetesClusterName

`func (o *ModelNetworkAlert) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ModelNetworkAlert) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ModelNetworkAlert) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.


### GetMasked

`func (o *ModelNetworkAlert) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelNetworkAlert) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelNetworkAlert) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetMatched

`func (o *ModelNetworkAlert) GetMatched() string`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *ModelNetworkAlert) GetMatchedOk() (*string, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *ModelNetworkAlert) SetMatched(v string)`

SetMatched sets Matched field to given value.


### GetNodeId

`func (o *ModelNetworkAlert) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelNetworkAlert) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelNetworkAlert) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeType

`func (o *ModelNetworkAlert) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelNetworkAlert) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelNetworkAlert) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetPodName

`func (o *ModelNetworkAlert) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ModelNetworkAlert) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ModelNetworkAlert) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetResourceType

`func (o *ModelNetworkAlert) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelNetworkAlert) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelNetworkAlert) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetRuleId

`func (o *ModelNetworkAlert) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelNetworkAlert) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelNetworkAlert) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetScopeId

`func (o *ModelNetworkAlert) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *ModelNetworkAlert) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *ModelNetworkAlert) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetSeverity

`func (o *ModelNetworkAlert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelNetworkAlert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelNetworkAlert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSeverityAnomalyContainerName

`func (o *ModelNetworkAlert) GetSeverityAnomalyContainerName() string`

GetSeverityAnomalyContainerName returns the SeverityAnomalyContainerName field if non-nil, zero value otherwise.

### GetSeverityAnomalyContainerNameOk

`func (o *ModelNetworkAlert) GetSeverityAnomalyContainerNameOk() (*string, bool)`

GetSeverityAnomalyContainerNameOk returns a tuple with the SeverityAnomalyContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityAnomalyContainerName

`func (o *ModelNetworkAlert) SetSeverityAnomalyContainerName(v string)`

SetSeverityAnomalyContainerName sets SeverityAnomalyContainerName field to given value.


### GetSeverityScore

`func (o *ModelNetworkAlert) GetSeverityScore() float32`

GetSeverityScore returns the SeverityScore field if non-nil, zero value otherwise.

### GetSeverityScoreOk

`func (o *ModelNetworkAlert) GetSeverityScoreOk() (*float32, bool)`

GetSeverityScoreOk returns a tuple with the SeverityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityScore

`func (o *ModelNetworkAlert) SetSeverityScore(v float32)`

SetSeverityScore sets SeverityScore field to given value.


### GetSignatureId

`func (o *ModelNetworkAlert) GetSignatureId() int32`

GetSignatureId returns the SignatureId field if non-nil, zero value otherwise.

### GetSignatureIdOk

`func (o *ModelNetworkAlert) GetSignatureIdOk() (*int32, bool)`

GetSignatureIdOk returns a tuple with the SignatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureId

`func (o *ModelNetworkAlert) SetSignatureId(v int32)`

SetSignatureId sets SignatureId field to given value.


### GetSourceIp

`func (o *ModelNetworkAlert) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *ModelNetworkAlert) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *ModelNetworkAlert) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.


### GetSourcePort

`func (o *ModelNetworkAlert) GetSourcePort() int32`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *ModelNetworkAlert) GetSourcePortOk() (*int32, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *ModelNetworkAlert) SetSourcePort(v int32)`

SetSourcePort sets SourcePort field to given value.


### GetSummary

`func (o *ModelNetworkAlert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelNetworkAlert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelNetworkAlert) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTags

`func (o *ModelNetworkAlert) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelNetworkAlert) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelNetworkAlert) SetTags(v string)`

SetTags sets Tags field to given value.


### GetUpdatedAt

`func (o *ModelNetworkAlert) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelNetworkAlert) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelNetworkAlert) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


