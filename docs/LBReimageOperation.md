# LBReimageOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** |  | 
**State** | [**LBReimageOperationStateEnum**](LBReimageOperationStateEnum.md) |  | 
**RollbackUntil** | **NullableString** |  | 

## Methods

### NewLBReimageOperation

`func NewLBReimageOperation(operationId string, state LBReimageOperationStateEnum, rollbackUntil NullableString, ) *LBReimageOperation`

NewLBReimageOperation instantiates a new LBReimageOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBReimageOperationWithDefaults

`func NewLBReimageOperationWithDefaults() *LBReimageOperation`

NewLBReimageOperationWithDefaults instantiates a new LBReimageOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *LBReimageOperation) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *LBReimageOperation) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *LBReimageOperation) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetState

`func (o *LBReimageOperation) GetState() LBReimageOperationStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LBReimageOperation) GetStateOk() (*LBReimageOperationStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LBReimageOperation) SetState(v LBReimageOperationStateEnum)`

SetState sets State field to given value.


### GetRollbackUntil

`func (o *LBReimageOperation) GetRollbackUntil() string`

GetRollbackUntil returns the RollbackUntil field if non-nil, zero value otherwise.

### GetRollbackUntilOk

`func (o *LBReimageOperation) GetRollbackUntilOk() (*string, bool)`

GetRollbackUntilOk returns a tuple with the RollbackUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackUntil

`func (o *LBReimageOperation) SetRollbackUntil(v string)`

SetRollbackUntil sets RollbackUntil field to given value.


### SetRollbackUntilNil

`func (o *LBReimageOperation) SetRollbackUntilNil(b bool)`

 SetRollbackUntilNil sets the value for RollbackUntil to be an explicit nil

### UnsetRollbackUntil
`func (o *LBReimageOperation) UnsetRollbackUntil()`

UnsetRollbackUntil ensures that no value is present for RollbackUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


