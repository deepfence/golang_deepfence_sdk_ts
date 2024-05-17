# ModelNetworkAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base64Payload** | **string** |  | 
**Category** | **string** |  | 
**ContainerName** | **string** |  | 
**Count** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**DestinationIp** | **string** |  | 
**DestinationPort** | **int32** |  | 
**Direction** | **string** |  | 
**Encrypted** | **bool** |  | 
**EventType** | **string** |  | 
**Headers** | **string** |  | 
**HostName** | **string** |  | 
**HttpType** | **string** |  | 
**KubernetesClusterId** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**Masked** | **bool** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**PodName** | **string** |  | 
**Protocol** | **string** |  | 
**References** | **string** |  | 
**Request** | Pointer to **string** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**RuleId** | **string** |  | 
**Severity** | **string** |  | 
**SourceIp** | **string** |  | 
**SourcePort** | **int32** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Tags** | **string** |  | 
**Techniques** | **[]string** |  | 
**UpdatedAt** | **int32** |  | 
**Url** | **string** |  | 

## Methods

### NewModelNetworkAlert

`func NewModelNetworkAlert(base64Payload string, category string, containerName string, count int32, createdAt int32, destinationIp string, destinationPort int32, direction string, encrypted bool, eventType string, headers string, hostName string, httpType string, kubernetesClusterId string, kubernetesClusterName string, masked bool, nodeId string, nodeType string, podName string, protocol string, references string, ruleId string, severity string, sourceIp string, sourcePort int32, summary string, tactics []string, tags string, techniques []string, updatedAt int32, url string, ) *ModelNetworkAlert`

NewModelNetworkAlert instantiates a new ModelNetworkAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkAlertWithDefaults

`func NewModelNetworkAlertWithDefaults() *ModelNetworkAlert`

NewModelNetworkAlertWithDefaults instantiates a new ModelNetworkAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase64Payload

`func (o *ModelNetworkAlert) GetBase64Payload() string`

GetBase64Payload returns the Base64Payload field if non-nil, zero value otherwise.

### GetBase64PayloadOk

`func (o *ModelNetworkAlert) GetBase64PayloadOk() (*string, bool)`

GetBase64PayloadOk returns a tuple with the Base64Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Payload

`func (o *ModelNetworkAlert) SetBase64Payload(v string)`

SetBase64Payload sets Base64Payload field to given value.


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

`func (o *ModelNetworkAlert) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *ModelNetworkAlert) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *ModelNetworkAlert) SetEncrypted(v bool)`

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


### GetHeaders

`func (o *ModelNetworkAlert) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ModelNetworkAlert) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ModelNetworkAlert) SetHeaders(v string)`

SetHeaders sets Headers field to given value.


### GetHostName

`func (o *ModelNetworkAlert) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelNetworkAlert) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelNetworkAlert) SetHostName(v string)`

SetHostName sets HostName field to given value.


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


### GetKubernetesClusterId

`func (o *ModelNetworkAlert) GetKubernetesClusterId() string`

GetKubernetesClusterId returns the KubernetesClusterId field if non-nil, zero value otherwise.

### GetKubernetesClusterIdOk

`func (o *ModelNetworkAlert) GetKubernetesClusterIdOk() (*string, bool)`

GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterId

`func (o *ModelNetworkAlert) SetKubernetesClusterId(v string)`

SetKubernetesClusterId sets KubernetesClusterId field to given value.


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


### GetProtocol

`func (o *ModelNetworkAlert) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelNetworkAlert) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelNetworkAlert) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetReferences

`func (o *ModelNetworkAlert) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ModelNetworkAlert) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ModelNetworkAlert) SetReferences(v string)`

SetReferences sets References field to given value.


### GetRequest

`func (o *ModelNetworkAlert) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ModelNetworkAlert) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ModelNetworkAlert) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ModelNetworkAlert) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *ModelNetworkAlert) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ModelNetworkAlert) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ModelNetworkAlert) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ModelNetworkAlert) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetRuleId

`func (o *ModelNetworkAlert) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelNetworkAlert) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelNetworkAlert) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


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


### GetTactics

`func (o *ModelNetworkAlert) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ModelNetworkAlert) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ModelNetworkAlert) SetTactics(v []string)`

SetTactics sets Tactics field to given value.


### SetTacticsNil

`func (o *ModelNetworkAlert) SetTacticsNil(b bool)`

 SetTacticsNil sets the value for Tactics to be an explicit nil

### UnsetTactics
`func (o *ModelNetworkAlert) UnsetTactics()`

UnsetTactics ensures that no value is present for Tactics, not even an explicit nil
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


### GetTechniques

`func (o *ModelNetworkAlert) GetTechniques() []string`

GetTechniques returns the Techniques field if non-nil, zero value otherwise.

### GetTechniquesOk

`func (o *ModelNetworkAlert) GetTechniquesOk() (*[]string, bool)`

GetTechniquesOk returns a tuple with the Techniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniques

`func (o *ModelNetworkAlert) SetTechniques(v []string)`

SetTechniques sets Techniques field to given value.


### SetTechniquesNil

`func (o *ModelNetworkAlert) SetTechniquesNil(b bool)`

 SetTechniquesNil sets the value for Techniques to be an explicit nil

### UnsetTechniques
`func (o *ModelNetworkAlert) UnsetTechniques()`

UnsetTechniques ensures that no value is present for Techniques, not even an explicit nil
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


### GetUrl

`func (o *ModelNetworkAlert) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelNetworkAlert) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelNetworkAlert) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


