# LBReimageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | [**NullableLBReimageOperation**](LBReimageOperation.md) |  | 
**Start** | [**NullableLBReimageStart**](LBReimageStart.md) |  | 

## Methods

### NewLBReimageStatus

`func NewLBReimageStatus(operation NullableLBReimageOperation, start NullableLBReimageStart, ) *LBReimageStatus`

NewLBReimageStatus instantiates a new LBReimageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBReimageStatusWithDefaults

`func NewLBReimageStatusWithDefaults() *LBReimageStatus`

NewLBReimageStatusWithDefaults instantiates a new LBReimageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *LBReimageStatus) GetOperation() LBReimageOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LBReimageStatus) GetOperationOk() (*LBReimageOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LBReimageStatus) SetOperation(v LBReimageOperation)`

SetOperation sets Operation field to given value.


### SetOperationNil

`func (o *LBReimageStatus) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *LBReimageStatus) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetStart

`func (o *LBReimageStatus) GetStart() LBReimageStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LBReimageStatus) GetStartOk() (*LBReimageStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LBReimageStatus) SetStart(v LBReimageStart)`

SetStart sets Start field to given value.


### SetStartNil

`func (o *LBReimageStatus) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *LBReimageStatus) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


