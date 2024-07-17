# ModelAPIEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** |  | [optional] 
**CloudRegion** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**HostIp** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**SchemaInfo** | Pointer to **string** |  | [optional] 
**SourceHosts** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelAPIEndpoint

`func NewModelAPIEndpoint() *ModelAPIEndpoint`

NewModelAPIEndpoint instantiates a new ModelAPIEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIEndpointWithDefaults

`func NewModelAPIEndpointWithDefaults() *ModelAPIEndpoint`

NewModelAPIEndpointWithDefaults instantiates a new ModelAPIEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ModelAPIEndpoint) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelAPIEndpoint) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelAPIEndpoint) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ModelAPIEndpoint) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCloudRegion

`func (o *ModelAPIEndpoint) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *ModelAPIEndpoint) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *ModelAPIEndpoint) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.

### HasCloudRegion

`func (o *ModelAPIEndpoint) HasCloudRegion() bool`

HasCloudRegion returns a boolean if a field has been set.

### GetCloudType

`func (o *ModelAPIEndpoint) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ModelAPIEndpoint) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ModelAPIEndpoint) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ModelAPIEndpoint) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelAPIEndpoint) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelAPIEndpoint) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelAPIEndpoint) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelAPIEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *ModelAPIEndpoint) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModelAPIEndpoint) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModelAPIEndpoint) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ModelAPIEndpoint) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetHost

`func (o *ModelAPIEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ModelAPIEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ModelAPIEndpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ModelAPIEndpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostIp

`func (o *ModelAPIEndpoint) GetHostIp() string`

GetHostIp returns the HostIp field if non-nil, zero value otherwise.

### GetHostIpOk

`func (o *ModelAPIEndpoint) GetHostIpOk() (*string, bool)`

GetHostIpOk returns a tuple with the HostIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIp

`func (o *ModelAPIEndpoint) SetHostIp(v string)`

SetHostIp sets HostIp field to given value.

### HasHostIp

`func (o *ModelAPIEndpoint) HasHostIp() bool`

HasHostIp returns a boolean if a field has been set.

### GetMethod

`func (o *ModelAPIEndpoint) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ModelAPIEndpoint) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ModelAPIEndpoint) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ModelAPIEndpoint) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *ModelAPIEndpoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelAPIEndpoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelAPIEndpoint) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelAPIEndpoint) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *ModelAPIEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModelAPIEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModelAPIEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModelAPIEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSchemaInfo

`func (o *ModelAPIEndpoint) GetSchemaInfo() string`

GetSchemaInfo returns the SchemaInfo field if non-nil, zero value otherwise.

### GetSchemaInfoOk

`func (o *ModelAPIEndpoint) GetSchemaInfoOk() (*string, bool)`

GetSchemaInfoOk returns a tuple with the SchemaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfo

`func (o *ModelAPIEndpoint) SetSchemaInfo(v string)`

SetSchemaInfo sets SchemaInfo field to given value.

### HasSchemaInfo

`func (o *ModelAPIEndpoint) HasSchemaInfo() bool`

HasSchemaInfo returns a boolean if a field has been set.

### GetSourceHosts

`func (o *ModelAPIEndpoint) GetSourceHosts() []string`

GetSourceHosts returns the SourceHosts field if non-nil, zero value otherwise.

### GetSourceHostsOk

`func (o *ModelAPIEndpoint) GetSourceHostsOk() (*[]string, bool)`

GetSourceHostsOk returns a tuple with the SourceHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHosts

`func (o *ModelAPIEndpoint) SetSourceHosts(v []string)`

SetSourceHosts sets SourceHosts field to given value.

### HasSourceHosts

`func (o *ModelAPIEndpoint) HasSourceHosts() bool`

HasSourceHosts returns a boolean if a field has been set.

### SetSourceHostsNil

`func (o *ModelAPIEndpoint) SetSourceHostsNil(b bool)`

 SetSourceHostsNil sets the value for SourceHosts to be an explicit nil

### UnsetSourceHosts
`func (o *ModelAPIEndpoint) UnsetSourceHosts()`

UnsetSourceHosts ensures that no value is present for SourceHosts, not even an explicit nil
### GetUpdatedAt

`func (o *ModelAPIEndpoint) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelAPIEndpoint) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelAPIEndpoint) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelAPIEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


