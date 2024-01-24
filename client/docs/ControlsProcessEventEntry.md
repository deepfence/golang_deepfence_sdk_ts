# ControlsProcessEventEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**FailureSeverity** | **string** |  | 
**SkipCommList** | Pointer to **[]string** |  | [optional] 
**SkipPathList** | Pointer to **[]string** |  | [optional] 
**SkipUserList** | Pointer to **[]string** |  | [optional] 
**SuccessSeverity** | **string** |  | 

## Methods

### NewControlsProcessEventEntry

`func NewControlsProcessEventEntry(event string, failureSeverity string, successSeverity string, ) *ControlsProcessEventEntry`

NewControlsProcessEventEntry instantiates a new ControlsProcessEventEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsProcessEventEntryWithDefaults

`func NewControlsProcessEventEntryWithDefaults() *ControlsProcessEventEntry`

NewControlsProcessEventEntryWithDefaults instantiates a new ControlsProcessEventEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ControlsProcessEventEntry) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ControlsProcessEventEntry) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ControlsProcessEventEntry) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetFailureSeverity

`func (o *ControlsProcessEventEntry) GetFailureSeverity() string`

GetFailureSeverity returns the FailureSeverity field if non-nil, zero value otherwise.

### GetFailureSeverityOk

`func (o *ControlsProcessEventEntry) GetFailureSeverityOk() (*string, bool)`

GetFailureSeverityOk returns a tuple with the FailureSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureSeverity

`func (o *ControlsProcessEventEntry) SetFailureSeverity(v string)`

SetFailureSeverity sets FailureSeverity field to given value.


### GetSkipCommList

`func (o *ControlsProcessEventEntry) GetSkipCommList() []string`

GetSkipCommList returns the SkipCommList field if non-nil, zero value otherwise.

### GetSkipCommListOk

`func (o *ControlsProcessEventEntry) GetSkipCommListOk() (*[]string, bool)`

GetSkipCommListOk returns a tuple with the SkipCommList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCommList

`func (o *ControlsProcessEventEntry) SetSkipCommList(v []string)`

SetSkipCommList sets SkipCommList field to given value.

### HasSkipCommList

`func (o *ControlsProcessEventEntry) HasSkipCommList() bool`

HasSkipCommList returns a boolean if a field has been set.

### SetSkipCommListNil

`func (o *ControlsProcessEventEntry) SetSkipCommListNil(b bool)`

 SetSkipCommListNil sets the value for SkipCommList to be an explicit nil

### UnsetSkipCommList
`func (o *ControlsProcessEventEntry) UnsetSkipCommList()`

UnsetSkipCommList ensures that no value is present for SkipCommList, not even an explicit nil
### GetSkipPathList

`func (o *ControlsProcessEventEntry) GetSkipPathList() []string`

GetSkipPathList returns the SkipPathList field if non-nil, zero value otherwise.

### GetSkipPathListOk

`func (o *ControlsProcessEventEntry) GetSkipPathListOk() (*[]string, bool)`

GetSkipPathListOk returns a tuple with the SkipPathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPathList

`func (o *ControlsProcessEventEntry) SetSkipPathList(v []string)`

SetSkipPathList sets SkipPathList field to given value.

### HasSkipPathList

`func (o *ControlsProcessEventEntry) HasSkipPathList() bool`

HasSkipPathList returns a boolean if a field has been set.

### SetSkipPathListNil

`func (o *ControlsProcessEventEntry) SetSkipPathListNil(b bool)`

 SetSkipPathListNil sets the value for SkipPathList to be an explicit nil

### UnsetSkipPathList
`func (o *ControlsProcessEventEntry) UnsetSkipPathList()`

UnsetSkipPathList ensures that no value is present for SkipPathList, not even an explicit nil
### GetSkipUserList

`func (o *ControlsProcessEventEntry) GetSkipUserList() []string`

GetSkipUserList returns the SkipUserList field if non-nil, zero value otherwise.

### GetSkipUserListOk

`func (o *ControlsProcessEventEntry) GetSkipUserListOk() (*[]string, bool)`

GetSkipUserListOk returns a tuple with the SkipUserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUserList

`func (o *ControlsProcessEventEntry) SetSkipUserList(v []string)`

SetSkipUserList sets SkipUserList field to given value.

### HasSkipUserList

`func (o *ControlsProcessEventEntry) HasSkipUserList() bool`

HasSkipUserList returns a boolean if a field has been set.

### SetSkipUserListNil

`func (o *ControlsProcessEventEntry) SetSkipUserListNil(b bool)`

 SetSkipUserListNil sets the value for SkipUserList to be an explicit nil

### UnsetSkipUserList
`func (o *ControlsProcessEventEntry) UnsetSkipUserList()`

UnsetSkipUserList ensures that no value is present for SkipUserList, not even an explicit nil
### GetSuccessSeverity

`func (o *ControlsProcessEventEntry) GetSuccessSeverity() string`

GetSuccessSeverity returns the SuccessSeverity field if non-nil, zero value otherwise.

### GetSuccessSeverityOk

`func (o *ControlsProcessEventEntry) GetSuccessSeverityOk() (*string, bool)`

GetSuccessSeverityOk returns a tuple with the SuccessSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessSeverity

`func (o *ControlsProcessEventEntry) SetSuccessSeverity(v string)`

SetSuccessSeverity sets SuccessSeverity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


