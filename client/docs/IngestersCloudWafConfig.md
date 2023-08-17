# IngestersCloudWafConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsWafArn** | [**[]IngestersAWSWafARN**](IngestersAWSWafARN.md) |  | 
**CloudProvider** | **string** |  | 

## Methods

### NewIngestersCloudWafConfig

`func NewIngestersCloudWafConfig(awsWafArn []IngestersAWSWafARN, cloudProvider string, ) *IngestersCloudWafConfig`

NewIngestersCloudWafConfig instantiates a new IngestersCloudWafConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestersCloudWafConfigWithDefaults

`func NewIngestersCloudWafConfigWithDefaults() *IngestersCloudWafConfig`

NewIngestersCloudWafConfigWithDefaults instantiates a new IngestersCloudWafConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsWafArn

`func (o *IngestersCloudWafConfig) GetAwsWafArn() []IngestersAWSWafARN`

GetAwsWafArn returns the AwsWafArn field if non-nil, zero value otherwise.

### GetAwsWafArnOk

`func (o *IngestersCloudWafConfig) GetAwsWafArnOk() (*[]IngestersAWSWafARN, bool)`

GetAwsWafArnOk returns a tuple with the AwsWafArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsWafArn

`func (o *IngestersCloudWafConfig) SetAwsWafArn(v []IngestersAWSWafARN)`

SetAwsWafArn sets AwsWafArn field to given value.


### SetAwsWafArnNil

`func (o *IngestersCloudWafConfig) SetAwsWafArnNil(b bool)`

 SetAwsWafArnNil sets the value for AwsWafArn to be an explicit nil

### UnsetAwsWafArn
`func (o *IngestersCloudWafConfig) UnsetAwsWafArn()`

UnsetAwsWafArn ensures that no value is present for AwsWafArn, not even an explicit nil
### GetCloudProvider

`func (o *IngestersCloudWafConfig) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *IngestersCloudWafConfig) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *IngestersCloudWafConfig) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


