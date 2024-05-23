# ModelEnableCloudTracerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | [**[]ModelAgentID**](ModelAgentID.md) |  | 
**AwsS3Bucket** | **string** |  | 

## Methods

### NewModelEnableCloudTracerReq

`func NewModelEnableCloudTracerReq(agentIds []ModelAgentID, awsS3Bucket string, ) *ModelEnableCloudTracerReq`

NewModelEnableCloudTracerReq instantiates a new ModelEnableCloudTracerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEnableCloudTracerReqWithDefaults

`func NewModelEnableCloudTracerReqWithDefaults() *ModelEnableCloudTracerReq`

NewModelEnableCloudTracerReqWithDefaults instantiates a new ModelEnableCloudTracerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *ModelEnableCloudTracerReq) GetAgentIds() []ModelAgentID`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *ModelEnableCloudTracerReq) GetAgentIdsOk() (*[]ModelAgentID, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *ModelEnableCloudTracerReq) SetAgentIds(v []ModelAgentID)`

SetAgentIds sets AgentIds field to given value.


### SetAgentIdsNil

`func (o *ModelEnableCloudTracerReq) SetAgentIdsNil(b bool)`

 SetAgentIdsNil sets the value for AgentIds to be an explicit nil

### UnsetAgentIds
`func (o *ModelEnableCloudTracerReq) UnsetAgentIds()`

UnsetAgentIds ensures that no value is present for AgentIds, not even an explicit nil
### GetAwsS3Bucket

`func (o *ModelEnableCloudTracerReq) GetAwsS3Bucket() string`

GetAwsS3Bucket returns the AwsS3Bucket field if non-nil, zero value otherwise.

### GetAwsS3BucketOk

`func (o *ModelEnableCloudTracerReq) GetAwsS3BucketOk() (*string, bool)`

GetAwsS3BucketOk returns a tuple with the AwsS3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3Bucket

`func (o *ModelEnableCloudTracerReq) SetAwsS3Bucket(v string)`

SetAwsS3Bucket sets AwsS3Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


