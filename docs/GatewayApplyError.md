# GatewayApplyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **NullableString** | The kind of the object the agent refused, or null. | 
**Id** | **NullableInt32** |  | 
**Message** | **string** |  | 
**Rollback** | **string** | not_needed, pending, succeeded or failed. | 

## Methods

### NewGatewayApplyError

`func NewGatewayApplyError(kind NullableString, id NullableInt32, message string, rollback string, ) *GatewayApplyError`

NewGatewayApplyError instantiates a new GatewayApplyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayApplyErrorWithDefaults

`func NewGatewayApplyErrorWithDefaults() *GatewayApplyError`

NewGatewayApplyErrorWithDefaults instantiates a new GatewayApplyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GatewayApplyError) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GatewayApplyError) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GatewayApplyError) SetKind(v string)`

SetKind sets Kind field to given value.


### SetKindNil

`func (o *GatewayApplyError) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *GatewayApplyError) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetId

`func (o *GatewayApplyError) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayApplyError) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayApplyError) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GatewayApplyError) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GatewayApplyError) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMessage

`func (o *GatewayApplyError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GatewayApplyError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GatewayApplyError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRollback

`func (o *GatewayApplyError) GetRollback() string`

GetRollback returns the Rollback field if non-nil, zero value otherwise.

### GetRollbackOk

`func (o *GatewayApplyError) GetRollbackOk() (*string, bool)`

GetRollbackOk returns a tuple with the Rollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollback

`func (o *GatewayApplyError) SetRollback(v string)`

SetRollback sets Rollback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


