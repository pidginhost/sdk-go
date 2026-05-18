# DeleteRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **int32** | Line number of the DNS record to delete. | 

## Methods

### NewDeleteRecord

`func NewDeleteRecord(line int32, ) *DeleteRecord`

NewDeleteRecord instantiates a new DeleteRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRecordWithDefaults

`func NewDeleteRecordWithDefaults() *DeleteRecord`

NewDeleteRecordWithDefaults instantiates a new DeleteRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *DeleteRecord) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DeleteRecord) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DeleteRecord) SetLine(v int32)`

SetLine sets Line field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


