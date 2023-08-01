# SinglesignonGetSingleSignOnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**[]SinglesignonSSOResponse**](SinglesignonSSOResponse.md) |  | 
**Instructions** | [**SinglesignonSSOConfigurationInstructions**](SinglesignonSSOConfigurationInstructions.md) |  | 

## Methods

### NewSinglesignonGetSingleSignOnResponse

`func NewSinglesignonGetSingleSignOnResponse(config []SinglesignonSSOResponse, instructions SinglesignonSSOConfigurationInstructions, ) *SinglesignonGetSingleSignOnResponse`

NewSinglesignonGetSingleSignOnResponse instantiates a new SinglesignonGetSingleSignOnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglesignonGetSingleSignOnResponseWithDefaults

`func NewSinglesignonGetSingleSignOnResponseWithDefaults() *SinglesignonGetSingleSignOnResponse`

NewSinglesignonGetSingleSignOnResponseWithDefaults instantiates a new SinglesignonGetSingleSignOnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SinglesignonGetSingleSignOnResponse) GetConfig() []SinglesignonSSOResponse`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SinglesignonGetSingleSignOnResponse) GetConfigOk() (*[]SinglesignonSSOResponse, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SinglesignonGetSingleSignOnResponse) SetConfig(v []SinglesignonSSOResponse)`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *SinglesignonGetSingleSignOnResponse) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *SinglesignonGetSingleSignOnResponse) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetInstructions

`func (o *SinglesignonGetSingleSignOnResponse) GetInstructions() SinglesignonSSOConfigurationInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *SinglesignonGetSingleSignOnResponse) GetInstructionsOk() (*SinglesignonSSOConfigurationInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *SinglesignonGetSingleSignOnResponse) SetInstructions(v SinglesignonSSOConfigurationInstructions)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


