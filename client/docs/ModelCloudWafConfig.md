# ModelCloudWafConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsWafArn** | [**[]ModelAWSWafARN**](ModelAWSWafARN.md) |  | 
**CloudProvider** | **string** |  | 

## Methods

### NewModelCloudWafConfig

`func NewModelCloudWafConfig(awsWafArn []ModelAWSWafARN, cloudProvider string, ) *ModelCloudWafConfig`

NewModelCloudWafConfig instantiates a new ModelCloudWafConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudWafConfigWithDefaults

`func NewModelCloudWafConfigWithDefaults() *ModelCloudWafConfig`

NewModelCloudWafConfigWithDefaults instantiates a new ModelCloudWafConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsWafArn

`func (o *ModelCloudWafConfig) GetAwsWafArn() []ModelAWSWafARN`

GetAwsWafArn returns the AwsWafArn field if non-nil, zero value otherwise.

### GetAwsWafArnOk

`func (o *ModelCloudWafConfig) GetAwsWafArnOk() (*[]ModelAWSWafARN, bool)`

GetAwsWafArnOk returns a tuple with the AwsWafArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsWafArn

`func (o *ModelCloudWafConfig) SetAwsWafArn(v []ModelAWSWafARN)`

SetAwsWafArn sets AwsWafArn field to given value.


### SetAwsWafArnNil

`func (o *ModelCloudWafConfig) SetAwsWafArnNil(b bool)`

 SetAwsWafArnNil sets the value for AwsWafArn to be an explicit nil

### UnsetAwsWafArn
`func (o *ModelCloudWafConfig) UnsetAwsWafArn()`

UnsetAwsWafArn ensures that no value is present for AwsWafArn, not even an explicit nil
### GetCloudProvider

`func (o *ModelCloudWafConfig) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudWafConfig) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudWafConfig) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


