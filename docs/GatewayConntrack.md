# GatewayConntrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **NullableInt32** |  | 
**Max** | **NullableInt32** |  | 
**InsertFailures** | **NullableString** | An unsigned 64-bit decimal string. | 

## Methods

### NewGatewayConntrack

`func NewGatewayConntrack(count NullableInt32, max NullableInt32, insertFailures NullableString, ) *GatewayConntrack`

NewGatewayConntrack instantiates a new GatewayConntrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayConntrackWithDefaults

`func NewGatewayConntrackWithDefaults() *GatewayConntrack`

NewGatewayConntrackWithDefaults instantiates a new GatewayConntrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GatewayConntrack) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GatewayConntrack) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GatewayConntrack) SetCount(v int32)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *GatewayConntrack) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *GatewayConntrack) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetMax

`func (o *GatewayConntrack) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GatewayConntrack) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GatewayConntrack) SetMax(v int32)`

SetMax sets Max field to given value.


### SetMaxNil

`func (o *GatewayConntrack) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *GatewayConntrack) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetInsertFailures

`func (o *GatewayConntrack) GetInsertFailures() string`

GetInsertFailures returns the InsertFailures field if non-nil, zero value otherwise.

### GetInsertFailuresOk

`func (o *GatewayConntrack) GetInsertFailuresOk() (*string, bool)`

GetInsertFailuresOk returns a tuple with the InsertFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertFailures

`func (o *GatewayConntrack) SetInsertFailures(v string)`

SetInsertFailures sets InsertFailures field to given value.


### SetInsertFailuresNil

`func (o *GatewayConntrack) SetInsertFailuresNil(b bool)`

 SetInsertFailuresNil sets the value for InsertFailures to be an explicit nil

### UnsetInsertFailures
`func (o *GatewayConntrack) UnsetInsertFailures()`

UnsetInsertFailures ensures that no value is present for InsertFailures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


