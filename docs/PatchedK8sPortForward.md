# PatchedK8sPortForward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**InternalIp** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to [**ProtocolEnum**](ProtocolEnum.md) |  | [optional] 

## Methods

### NewPatchedK8sPortForward

`func NewPatchedK8sPortForward() *PatchedK8sPortForward`

NewPatchedK8sPortForward instantiates a new PatchedK8sPortForward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedK8sPortForwardWithDefaults

`func NewPatchedK8sPortForwardWithDefaults() *PatchedK8sPortForward`

NewPatchedK8sPortForwardWithDefaults instantiates a new PatchedK8sPortForward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedK8sPortForward) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedK8sPortForward) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedK8sPortForward) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedK8sPortForward) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalIp

`func (o *PatchedK8sPortForward) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *PatchedK8sPortForward) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *PatchedK8sPortForward) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *PatchedK8sPortForward) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetPort

`func (o *PatchedK8sPortForward) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedK8sPortForward) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedK8sPortForward) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedK8sPortForward) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedK8sPortForward) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedK8sPortForward) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedK8sPortForward) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedK8sPortForward) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


