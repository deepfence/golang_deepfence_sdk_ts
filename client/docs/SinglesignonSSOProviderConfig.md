# SinglesignonSSOProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**SsoProviderType** | **string** |  | 

## Methods

### NewSinglesignonSSOProviderConfig

`func NewSinglesignonSSOProviderConfig(clientId string, clientSecret string, ssoProviderType string, ) *SinglesignonSSOProviderConfig`

NewSinglesignonSSOProviderConfig instantiates a new SinglesignonSSOProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglesignonSSOProviderConfigWithDefaults

`func NewSinglesignonSSOProviderConfigWithDefaults() *SinglesignonSSOProviderConfig`

NewSinglesignonSSOProviderConfigWithDefaults instantiates a new SinglesignonSSOProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SinglesignonSSOProviderConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SinglesignonSSOProviderConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SinglesignonSSOProviderConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *SinglesignonSSOProviderConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SinglesignonSSOProviderConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SinglesignonSSOProviderConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetIssuerUrl

`func (o *SinglesignonSSOProviderConfig) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *SinglesignonSSOProviderConfig) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *SinglesignonSSOProviderConfig) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *SinglesignonSSOProviderConfig) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetSsoProviderType

`func (o *SinglesignonSSOProviderConfig) GetSsoProviderType() string`

GetSsoProviderType returns the SsoProviderType field if non-nil, zero value otherwise.

### GetSsoProviderTypeOk

`func (o *SinglesignonSSOProviderConfig) GetSsoProviderTypeOk() (*string, bool)`

GetSsoProviderTypeOk returns a tuple with the SsoProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoProviderType

`func (o *SinglesignonSSOProviderConfig) SetSsoProviderType(v string)`

SetSsoProviderType sets SsoProviderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


