# PatchedK8sPortForwardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalIp** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to [**ProtocolEnum**](ProtocolEnum.md) |  | [optional] 

## Methods

### NewPatchedK8sPortForwardRequest

`func NewPatchedK8sPortForwardRequest() *PatchedK8sPortForwardRequest`

NewPatchedK8sPortForwardRequest instantiates a new PatchedK8sPortForwardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedK8sPortForwardRequestWithDefaults

`func NewPatchedK8sPortForwardRequestWithDefaults() *PatchedK8sPortForwardRequest`

NewPatchedK8sPortForwardRequestWithDefaults instantiates a new PatchedK8sPortForwardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalIp

`func (o *PatchedK8sPortForwardRequest) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *PatchedK8sPortForwardRequest) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *PatchedK8sPortForwardRequest) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *PatchedK8sPortForwardRequest) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetPort

`func (o *PatchedK8sPortForwardRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedK8sPortForwardRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedK8sPortForwardRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedK8sPortForwardRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedK8sPortForwardRequest) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedK8sPortForwardRequest) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedK8sPortForwardRequest) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedK8sPortForwardRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


