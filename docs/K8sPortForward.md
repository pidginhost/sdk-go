# K8sPortForward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**InternalIp** | **string** |  | 
**Port** | **int32** |  | 
**Protocol** | [**ProtocolEnum**](ProtocolEnum.md) |  | 

## Methods

### NewK8sPortForward

`func NewK8sPortForward(id int32, internalIp string, port int32, protocol ProtocolEnum, ) *K8sPortForward`

NewK8sPortForward instantiates a new K8sPortForward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sPortForwardWithDefaults

`func NewK8sPortForwardWithDefaults() *K8sPortForward`

NewK8sPortForwardWithDefaults instantiates a new K8sPortForward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *K8sPortForward) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *K8sPortForward) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *K8sPortForward) SetId(v int32)`

SetId sets Id field to given value.


### GetInternalIp

`func (o *K8sPortForward) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *K8sPortForward) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *K8sPortForward) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.


### GetPort

`func (o *K8sPortForward) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *K8sPortForward) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *K8sPortForward) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *K8sPortForward) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *K8sPortForward) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *K8sPortForward) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


