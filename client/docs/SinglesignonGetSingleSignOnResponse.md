# SinglesignonGetSingleSignOnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **int32** |  | 
**Instructions** | [**[]SinglesignonSSOConfigurationInstructions**](SinglesignonSSOConfigurationInstructions.md) |  | 
**IssuerUrl** | **string** |  | 
**Label** | **string** |  | 
**SsoProviderType** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSinglesignonGetSingleSignOnResponse

`func NewSinglesignonGetSingleSignOnResponse(clientId string, createdAt time.Time, id int32, instructions []SinglesignonSSOConfigurationInstructions, issuerUrl string, label string, ssoProviderType string, updatedAt time.Time, ) *SinglesignonGetSingleSignOnResponse`

NewSinglesignonGetSingleSignOnResponse instantiates a new SinglesignonGetSingleSignOnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglesignonGetSingleSignOnResponseWithDefaults

`func NewSinglesignonGetSingleSignOnResponseWithDefaults() *SinglesignonGetSingleSignOnResponse`

NewSinglesignonGetSingleSignOnResponseWithDefaults instantiates a new SinglesignonGetSingleSignOnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SinglesignonGetSingleSignOnResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SinglesignonGetSingleSignOnResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SinglesignonGetSingleSignOnResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreatedAt

`func (o *SinglesignonGetSingleSignOnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SinglesignonGetSingleSignOnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SinglesignonGetSingleSignOnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *SinglesignonGetSingleSignOnResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SinglesignonGetSingleSignOnResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SinglesignonGetSingleSignOnResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetInstructions

`func (o *SinglesignonGetSingleSignOnResponse) GetInstructions() []SinglesignonSSOConfigurationInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *SinglesignonGetSingleSignOnResponse) GetInstructionsOk() (*[]SinglesignonSSOConfigurationInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *SinglesignonGetSingleSignOnResponse) SetInstructions(v []SinglesignonSSOConfigurationInstructions)`

SetInstructions sets Instructions field to given value.


### SetInstructionsNil

`func (o *SinglesignonGetSingleSignOnResponse) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *SinglesignonGetSingleSignOnResponse) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetIssuerUrl

`func (o *SinglesignonGetSingleSignOnResponse) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *SinglesignonGetSingleSignOnResponse) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *SinglesignonGetSingleSignOnResponse) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetLabel

`func (o *SinglesignonGetSingleSignOnResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SinglesignonGetSingleSignOnResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SinglesignonGetSingleSignOnResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSsoProviderType

`func (o *SinglesignonGetSingleSignOnResponse) GetSsoProviderType() string`

GetSsoProviderType returns the SsoProviderType field if non-nil, zero value otherwise.

### GetSsoProviderTypeOk

`func (o *SinglesignonGetSingleSignOnResponse) GetSsoProviderTypeOk() (*string, bool)`

GetSsoProviderTypeOk returns a tuple with the SsoProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoProviderType

`func (o *SinglesignonGetSingleSignOnResponse) SetSsoProviderType(v string)`

SetSsoProviderType sets SsoProviderType field to given value.


### GetUpdatedAt

`func (o *SinglesignonGetSingleSignOnResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SinglesignonGetSingleSignOnResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SinglesignonGetSingleSignOnResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


