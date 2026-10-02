# AttachVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vm** | **int32** | Server ID | 

## Methods

### NewAttachVolumeRequest

`func NewAttachVolumeRequest(vm int32, ) *AttachVolumeRequest`

NewAttachVolumeRequest instantiates a new AttachVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachVolumeRequestWithDefaults

`func NewAttachVolumeRequestWithDefaults() *AttachVolumeRequest`

NewAttachVolumeRequestWithDefaults instantiates a new AttachVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVm

`func (o *AttachVolumeRequest) GetVm() int32`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *AttachVolumeRequest) GetVmOk() (*int32, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *AttachVolumeRequest) SetVm(v int32)`

SetVm sets Vm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


