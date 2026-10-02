# K8sPortForwardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalIp** | **string** |  | 
**Port** | **int32** |  | 
**Protocol** | [**ProtocolEnum**](ProtocolEnum.md) |  | 

## Methods

### NewK8sPortForwardRequest

`func NewK8sPortForwardRequest(internalIp string, port int32, protocol ProtocolEnum, ) *K8sPortForwardRequest`

NewK8sPortForwardRequest instantiates a new K8sPortForwardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sPortForwardRequestWithDefaults

`func NewK8sPortForwardRequestWithDefaults() *K8sPortForwardRequest`

NewK8sPortForwardRequestWithDefaults instantiates a new K8sPortForwardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalIp

`func (o *K8sPortForwardRequest) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *K8sPortForwardRequest) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *K8sPortForwardRequest) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.


### GetPort

`func (o *K8sPortForwardRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *K8sPortForwardRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *K8sPortForwardRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *K8sPortForwardRequest) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *K8sPortForwardRequest) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *K8sPortForwardRequest) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


