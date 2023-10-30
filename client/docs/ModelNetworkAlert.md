# ModelNetworkAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProto** | **string** |  | 
**Category** | **string** |  | 
**ContainerName** | **string** |  | 
**Count** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**Description** | **string** |  | 
**DestinationIp** | **string** |  | 
**DestinationPort** | **int32** |  | 
**Direction** | **string** |  | 
**Encrypted** | **string** |  | 
**EventType** | **string** |  | 
**Geoip** | **string** |  | 
**Headers** | **string** |  | 
**HostName** | **string** |  | 
**HttpContentType** | **interface{}** |  | 
**HttpType** | **string** |  | 
**HttpUserAgent** | **string** |  | 
**Internal** | **string** |  | 
**IpReputation** | **string** |  | 
**KubernetesClusterId** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**Length** | **interface{}** |  | 
**LocalPort** | **int32** |  | 
**Masked** | **bool** |  | 
**Matched** | **string** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**PodName** | **string** |  | 
**Protocol** | **int32** |  | 
**RequestMethod** | **string** |  | 
**RequestPath** | **string** |  | 
**RequestPayload** | **string** |  | 
**RequestPrintablePayload** | **string** |  | 
**ResourceType** | **string** |  | 
**ResponsePayload** | **interface{}** |  | 
**ResponsePrintablePayload** | **interface{}** |  | 
**RuleId** | **string** |  | 
**Severity** | **string** |  | 
**SeverityScore** | **float32** |  | 
**SourceIp** | **string** |  | 
**SourcePort** | **int32** |  | 
**Status** | **interface{}** |  | 
**Summary** | **string** |  | 
**Tactics** | **[]string** |  | 
**Tags** | **string** |  | 
**Techniques** | **[]string** |  | 
**UpdatedAt** | **int32** |  | 
**Url** | **string** |  | 

## Methods

### NewModelNetworkAlert

`func NewModelNetworkAlert(appProto string, category string, containerName string, count int32, createdAt int32, description string, destinationIp string, destinationPort int32, direction string, encrypted string, eventType string, geoip string, headers string, hostName string, httpContentType interface{}, httpType string, httpUserAgent string, internal string, ipReputation string, kubernetesClusterId string, kubernetesClusterName string, length interface{}, localPort int32, masked bool, matched string, nodeId string, nodeType string, podName string, protocol int32, requestMethod string, requestPath string, requestPayload string, requestPrintablePayload string, resourceType string, responsePayload interface{}, responsePrintablePayload interface{}, ruleId string, severity string, severityScore float32, sourceIp string, sourcePort int32, status interface{}, summary string, tactics []string, tags string, techniques []string, updatedAt int32, url string, ) *ModelNetworkAlert`

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

`func (o *ModelNetworkAlert) GetGeoip() string`

GetGeoip returns the Geoip field if non-nil, zero value otherwise.

### GetGeoipOk

`func (o *ModelNetworkAlert) GetGeoipOk() (*string, bool)`

GetGeoipOk returns a tuple with the Geoip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoip

`func (o *ModelNetworkAlert) SetGeoip(v string)`

SetGeoip sets Geoip field to given value.


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


### GetHttpContentType

`func (o *ModelNetworkAlert) GetHttpContentType() interface{}`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *ModelNetworkAlert) GetHttpContentTypeOk() (*interface{}, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *ModelNetworkAlert) SetHttpContentType(v interface{})`

SetHttpContentType sets HttpContentType field to given value.


### SetHttpContentTypeNil

`func (o *ModelNetworkAlert) SetHttpContentTypeNil(b bool)`

 SetHttpContentTypeNil sets the value for HttpContentType to be an explicit nil

### UnsetHttpContentType
`func (o *ModelNetworkAlert) UnsetHttpContentType()`

UnsetHttpContentType ensures that no value is present for HttpContentType, not even an explicit nil
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


### GetHttpUserAgent

`func (o *ModelNetworkAlert) GetHttpUserAgent() string`

GetHttpUserAgent returns the HttpUserAgent field if non-nil, zero value otherwise.

### GetHttpUserAgentOk

`func (o *ModelNetworkAlert) GetHttpUserAgentOk() (*string, bool)`

GetHttpUserAgentOk returns a tuple with the HttpUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUserAgent

`func (o *ModelNetworkAlert) SetHttpUserAgent(v string)`

SetHttpUserAgent sets HttpUserAgent field to given value.


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


### GetLength

`func (o *ModelNetworkAlert) GetLength() interface{}`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ModelNetworkAlert) GetLengthOk() (*interface{}, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ModelNetworkAlert) SetLength(v interface{})`

SetLength sets Length field to given value.


### SetLengthNil

`func (o *ModelNetworkAlert) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *ModelNetworkAlert) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLocalPort

`func (o *ModelNetworkAlert) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *ModelNetworkAlert) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *ModelNetworkAlert) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.


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


### GetProtocol

`func (o *ModelNetworkAlert) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelNetworkAlert) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelNetworkAlert) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### GetRequestMethod

`func (o *ModelNetworkAlert) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *ModelNetworkAlert) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *ModelNetworkAlert) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.


### GetRequestPath

`func (o *ModelNetworkAlert) GetRequestPath() string`

GetRequestPath returns the RequestPath field if non-nil, zero value otherwise.

### GetRequestPathOk

`func (o *ModelNetworkAlert) GetRequestPathOk() (*string, bool)`

GetRequestPathOk returns a tuple with the RequestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPath

`func (o *ModelNetworkAlert) SetRequestPath(v string)`

SetRequestPath sets RequestPath field to given value.


### GetRequestPayload

`func (o *ModelNetworkAlert) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *ModelNetworkAlert) GetRequestPayloadOk() (*string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *ModelNetworkAlert) SetRequestPayload(v string)`

SetRequestPayload sets RequestPayload field to given value.


### GetRequestPrintablePayload

`func (o *ModelNetworkAlert) GetRequestPrintablePayload() string`

GetRequestPrintablePayload returns the RequestPrintablePayload field if non-nil, zero value otherwise.

### GetRequestPrintablePayloadOk

`func (o *ModelNetworkAlert) GetRequestPrintablePayloadOk() (*string, bool)`

GetRequestPrintablePayloadOk returns a tuple with the RequestPrintablePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPrintablePayload

`func (o *ModelNetworkAlert) SetRequestPrintablePayload(v string)`

SetRequestPrintablePayload sets RequestPrintablePayload field to given value.


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


### GetResponsePayload

`func (o *ModelNetworkAlert) GetResponsePayload() interface{}`

GetResponsePayload returns the ResponsePayload field if non-nil, zero value otherwise.

### GetResponsePayloadOk

`func (o *ModelNetworkAlert) GetResponsePayloadOk() (*interface{}, bool)`

GetResponsePayloadOk returns a tuple with the ResponsePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePayload

`func (o *ModelNetworkAlert) SetResponsePayload(v interface{})`

SetResponsePayload sets ResponsePayload field to given value.


### SetResponsePayloadNil

`func (o *ModelNetworkAlert) SetResponsePayloadNil(b bool)`

 SetResponsePayloadNil sets the value for ResponsePayload to be an explicit nil

### UnsetResponsePayload
`func (o *ModelNetworkAlert) UnsetResponsePayload()`

UnsetResponsePayload ensures that no value is present for ResponsePayload, not even an explicit nil
### GetResponsePrintablePayload

`func (o *ModelNetworkAlert) GetResponsePrintablePayload() interface{}`

GetResponsePrintablePayload returns the ResponsePrintablePayload field if non-nil, zero value otherwise.

### GetResponsePrintablePayloadOk

`func (o *ModelNetworkAlert) GetResponsePrintablePayloadOk() (*interface{}, bool)`

GetResponsePrintablePayloadOk returns a tuple with the ResponsePrintablePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePrintablePayload

`func (o *ModelNetworkAlert) SetResponsePrintablePayload(v interface{})`

SetResponsePrintablePayload sets ResponsePrintablePayload field to given value.


### SetResponsePrintablePayloadNil

`func (o *ModelNetworkAlert) SetResponsePrintablePayloadNil(b bool)`

 SetResponsePrintablePayloadNil sets the value for ResponsePrintablePayload to be an explicit nil

### UnsetResponsePrintablePayload
`func (o *ModelNetworkAlert) UnsetResponsePrintablePayload()`

UnsetResponsePrintablePayload ensures that no value is present for ResponsePrintablePayload, not even an explicit nil
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


### GetStatus

`func (o *ModelNetworkAlert) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelNetworkAlert) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelNetworkAlert) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ModelNetworkAlert) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelNetworkAlert) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
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


