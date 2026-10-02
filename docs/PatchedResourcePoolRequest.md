# PatchedResourcePoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSize** | Pointer to **int32** |  | [optional] 
**LocalDataLossAccepted** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPatchedResourcePoolRequest

`func NewPatchedResourcePoolRequest() *PatchedResourcePoolRequest`

NewPatchedResourcePoolRequest instantiates a new PatchedResourcePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedResourcePoolRequestWithDefaults

`func NewPatchedResourcePoolRequestWithDefaults() *PatchedResourcePoolRequest`

NewPatchedResourcePoolRequestWithDefaults instantiates a new PatchedResourcePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSize

`func (o *PatchedResourcePoolRequest) GetNewSize() int32`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *PatchedResourcePoolRequest) GetNewSizeOk() (*int32, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *PatchedResourcePoolRequest) SetNewSize(v int32)`

SetNewSize sets NewSize field to given value.

### HasNewSize

`func (o *PatchedResourcePoolRequest) HasNewSize() bool`

HasNewSize returns a boolean if a field has been set.

### GetLocalDataLossAccepted

`func (o *PatchedResourcePoolRequest) GetLocalDataLossAccepted() bool`

GetLocalDataLossAccepted returns the LocalDataLossAccepted field if non-nil, zero value otherwise.

### GetLocalDataLossAcceptedOk

`func (o *PatchedResourcePoolRequest) GetLocalDataLossAcceptedOk() (*bool, bool)`

GetLocalDataLossAcceptedOk returns a tuple with the LocalDataLossAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataLossAccepted

`func (o *PatchedResourcePoolRequest) SetLocalDataLossAccepted(v bool)`

SetLocalDataLossAccepted sets LocalDataLossAccepted field to given value.

### HasLocalDataLossAccepted

`func (o *PatchedResourcePoolRequest) HasLocalDataLossAccepted() bool`

HasLocalDataLossAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


