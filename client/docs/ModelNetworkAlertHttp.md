# ModelNetworkAlertHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | **string** |  | 
**Hostname** | **string** |  | 
**HttpContentType** | **interface{}** |  | 
**HttpUserAgent** | **string** |  | 
**Length** | **interface{}** |  | 
**LocalPort** | **int32** |  | 
**Protocol** | **int32** |  | 
**RequestMethod** | **string** |  | 
**RequestPath** | **string** |  | 
**RequestPayload** | **string** |  | 
**RequestPrintablePayload** | **string** |  | 
**ResponsePayload** | **interface{}** |  | 
**ResponsePrintablePayload** | **interface{}** |  | 
**Status** | **interface{}** |  | 
**Url** | **string** |  | 

## Methods

### NewModelNetworkAlertHttp

`func NewModelNetworkAlertHttp(headers string, hostname string, httpContentType interface{}, httpUserAgent string, length interface{}, localPort int32, protocol int32, requestMethod string, requestPath string, requestPayload string, requestPrintablePayload string, responsePayload interface{}, responsePrintablePayload interface{}, status interface{}, url string, ) *ModelNetworkAlertHttp`

NewModelNetworkAlertHttp instantiates a new ModelNetworkAlertHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkAlertHttpWithDefaults

`func NewModelNetworkAlertHttpWithDefaults() *ModelNetworkAlertHttp`

NewModelNetworkAlertHttpWithDefaults instantiates a new ModelNetworkAlertHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *ModelNetworkAlertHttp) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ModelNetworkAlertHttp) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ModelNetworkAlertHttp) SetHeaders(v string)`

SetHeaders sets Headers field to given value.


### GetHostname

`func (o *ModelNetworkAlertHttp) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ModelNetworkAlertHttp) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ModelNetworkAlertHttp) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHttpContentType

`func (o *ModelNetworkAlertHttp) GetHttpContentType() interface{}`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *ModelNetworkAlertHttp) GetHttpContentTypeOk() (*interface{}, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *ModelNetworkAlertHttp) SetHttpContentType(v interface{})`

SetHttpContentType sets HttpContentType field to given value.


### SetHttpContentTypeNil

`func (o *ModelNetworkAlertHttp) SetHttpContentTypeNil(b bool)`

 SetHttpContentTypeNil sets the value for HttpContentType to be an explicit nil

### UnsetHttpContentType
`func (o *ModelNetworkAlertHttp) UnsetHttpContentType()`

UnsetHttpContentType ensures that no value is present for HttpContentType, not even an explicit nil
### GetHttpUserAgent

`func (o *ModelNetworkAlertHttp) GetHttpUserAgent() string`

GetHttpUserAgent returns the HttpUserAgent field if non-nil, zero value otherwise.

### GetHttpUserAgentOk

`func (o *ModelNetworkAlertHttp) GetHttpUserAgentOk() (*string, bool)`

GetHttpUserAgentOk returns a tuple with the HttpUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUserAgent

`func (o *ModelNetworkAlertHttp) SetHttpUserAgent(v string)`

SetHttpUserAgent sets HttpUserAgent field to given value.


### GetLength

`func (o *ModelNetworkAlertHttp) GetLength() interface{}`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ModelNetworkAlertHttp) GetLengthOk() (*interface{}, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ModelNetworkAlertHttp) SetLength(v interface{})`

SetLength sets Length field to given value.


### SetLengthNil

`func (o *ModelNetworkAlertHttp) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *ModelNetworkAlertHttp) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLocalPort

`func (o *ModelNetworkAlertHttp) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *ModelNetworkAlertHttp) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *ModelNetworkAlertHttp) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.


### GetProtocol

`func (o *ModelNetworkAlertHttp) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelNetworkAlertHttp) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelNetworkAlertHttp) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### GetRequestMethod

`func (o *ModelNetworkAlertHttp) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *ModelNetworkAlertHttp) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *ModelNetworkAlertHttp) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.


### GetRequestPath

`func (o *ModelNetworkAlertHttp) GetRequestPath() string`

GetRequestPath returns the RequestPath field if non-nil, zero value otherwise.

### GetRequestPathOk

`func (o *ModelNetworkAlertHttp) GetRequestPathOk() (*string, bool)`

GetRequestPathOk returns a tuple with the RequestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPath

`func (o *ModelNetworkAlertHttp) SetRequestPath(v string)`

SetRequestPath sets RequestPath field to given value.


### GetRequestPayload

`func (o *ModelNetworkAlertHttp) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *ModelNetworkAlertHttp) GetRequestPayloadOk() (*string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *ModelNetworkAlertHttp) SetRequestPayload(v string)`

SetRequestPayload sets RequestPayload field to given value.


### GetRequestPrintablePayload

`func (o *ModelNetworkAlertHttp) GetRequestPrintablePayload() string`

GetRequestPrintablePayload returns the RequestPrintablePayload field if non-nil, zero value otherwise.

### GetRequestPrintablePayloadOk

`func (o *ModelNetworkAlertHttp) GetRequestPrintablePayloadOk() (*string, bool)`

GetRequestPrintablePayloadOk returns a tuple with the RequestPrintablePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPrintablePayload

`func (o *ModelNetworkAlertHttp) SetRequestPrintablePayload(v string)`

SetRequestPrintablePayload sets RequestPrintablePayload field to given value.


### GetResponsePayload

`func (o *ModelNetworkAlertHttp) GetResponsePayload() interface{}`

GetResponsePayload returns the ResponsePayload field if non-nil, zero value otherwise.

### GetResponsePayloadOk

`func (o *ModelNetworkAlertHttp) GetResponsePayloadOk() (*interface{}, bool)`

GetResponsePayloadOk returns a tuple with the ResponsePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePayload

`func (o *ModelNetworkAlertHttp) SetResponsePayload(v interface{})`

SetResponsePayload sets ResponsePayload field to given value.


### SetResponsePayloadNil

`func (o *ModelNetworkAlertHttp) SetResponsePayloadNil(b bool)`

 SetResponsePayloadNil sets the value for ResponsePayload to be an explicit nil

### UnsetResponsePayload
`func (o *ModelNetworkAlertHttp) UnsetResponsePayload()`

UnsetResponsePayload ensures that no value is present for ResponsePayload, not even an explicit nil
### GetResponsePrintablePayload

`func (o *ModelNetworkAlertHttp) GetResponsePrintablePayload() interface{}`

GetResponsePrintablePayload returns the ResponsePrintablePayload field if non-nil, zero value otherwise.

### GetResponsePrintablePayloadOk

`func (o *ModelNetworkAlertHttp) GetResponsePrintablePayloadOk() (*interface{}, bool)`

GetResponsePrintablePayloadOk returns a tuple with the ResponsePrintablePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePrintablePayload

`func (o *ModelNetworkAlertHttp) SetResponsePrintablePayload(v interface{})`

SetResponsePrintablePayload sets ResponsePrintablePayload field to given value.


### SetResponsePrintablePayloadNil

`func (o *ModelNetworkAlertHttp) SetResponsePrintablePayloadNil(b bool)`

 SetResponsePrintablePayloadNil sets the value for ResponsePrintablePayload to be an explicit nil

### UnsetResponsePrintablePayload
`func (o *ModelNetworkAlertHttp) UnsetResponsePrintablePayload()`

UnsetResponsePrintablePayload ensures that no value is present for ResponsePrintablePayload, not even an explicit nil
### GetStatus

`func (o *ModelNetworkAlertHttp) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelNetworkAlertHttp) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelNetworkAlertHttp) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ModelNetworkAlertHttp) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelNetworkAlertHttp) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *ModelNetworkAlertHttp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelNetworkAlertHttp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelNetworkAlertHttp) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


