# ResourcePoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSize** | Pointer to **int32** |  | [optional] 
**LocalDataLossAccepted** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewResourcePoolRequest

`func NewResourcePoolRequest() *ResourcePoolRequest`

NewResourcePoolRequest instantiates a new ResourcePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolRequestWithDefaults

`func NewResourcePoolRequestWithDefaults() *ResourcePoolRequest`

NewResourcePoolRequestWithDefaults instantiates a new ResourcePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSize

`func (o *ResourcePoolRequest) GetNewSize() int32`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *ResourcePoolRequest) GetNewSizeOk() (*int32, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *ResourcePoolRequest) SetNewSize(v int32)`

SetNewSize sets NewSize field to given value.

### HasNewSize

`func (o *ResourcePoolRequest) HasNewSize() bool`

HasNewSize returns a boolean if a field has been set.

### GetLocalDataLossAccepted

`func (o *ResourcePoolRequest) GetLocalDataLossAccepted() bool`

GetLocalDataLossAccepted returns the LocalDataLossAccepted field if non-nil, zero value otherwise.

### GetLocalDataLossAcceptedOk

`func (o *ResourcePoolRequest) GetLocalDataLossAcceptedOk() (*bool, bool)`

GetLocalDataLossAcceptedOk returns a tuple with the LocalDataLossAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataLossAccepted

`func (o *ResourcePoolRequest) SetLocalDataLossAccepted(v bool)`

SetLocalDataLossAccepted sets LocalDataLossAccepted field to given value.

### HasLocalDataLossAccepted

`func (o *ResourcePoolRequest) HasLocalDataLossAccepted() bool`

HasLocalDataLossAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


