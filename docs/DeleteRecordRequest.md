# DeleteRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **int32** | Line number of the DNS record to delete. | 

## Methods

### NewDeleteRecordRequest

`func NewDeleteRecordRequest(line int32, ) *DeleteRecordRequest`

NewDeleteRecordRequest instantiates a new DeleteRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRecordRequestWithDefaults

`func NewDeleteRecordRequestWithDefaults() *DeleteRecordRequest`

NewDeleteRecordRequestWithDefaults instantiates a new DeleteRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *DeleteRecordRequest) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DeleteRecordRequest) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DeleteRecordRequest) SetLine(v int32)`

SetLine sets Line field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


