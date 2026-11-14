# GatewayForwardWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**GatewayForwardProtocolEnum**](GatewayForwardProtocolEnum.md) |  | 
**PublicPorts** | [**GatewayPortRangeRequest**](GatewayPortRangeRequest.md) |  | 
**Target** | [**GatewayForwardTargetRequest**](GatewayForwardTargetRequest.md) |  | 
**AllowedSources** | Pointer to **[]string** | IPv4 networks; empty means any source. | [optional] 
**PreserveClientIp** | Pointer to **bool** | Only false is accepted for now. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayForwardWriteRequest

`func NewGatewayForwardWriteRequest(protocol GatewayForwardProtocolEnum, publicPorts GatewayPortRangeRequest, target GatewayForwardTargetRequest, ) *GatewayForwardWriteRequest`

NewGatewayForwardWriteRequest instantiates a new GatewayForwardWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayForwardWriteRequestWithDefaults

`func NewGatewayForwardWriteRequestWithDefaults() *GatewayForwardWriteRequest`

NewGatewayForwardWriteRequestWithDefaults instantiates a new GatewayForwardWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *GatewayForwardWriteRequest) GetProtocol() GatewayForwardProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayForwardWriteRequest) GetProtocolOk() (*GatewayForwardProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayForwardWriteRequest) SetProtocol(v GatewayForwardProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetPublicPorts

`func (o *GatewayForwardWriteRequest) GetPublicPorts() GatewayPortRangeRequest`

GetPublicPorts returns the PublicPorts field if non-nil, zero value otherwise.

### GetPublicPortsOk

`func (o *GatewayForwardWriteRequest) GetPublicPortsOk() (*GatewayPortRangeRequest, bool)`

GetPublicPortsOk returns a tuple with the PublicPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPorts

`func (o *GatewayForwardWriteRequest) SetPublicPorts(v GatewayPortRangeRequest)`

SetPublicPorts sets PublicPorts field to given value.


### GetTarget

`func (o *GatewayForwardWriteRequest) GetTarget() GatewayForwardTargetRequest`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GatewayForwardWriteRequest) GetTargetOk() (*GatewayForwardTargetRequest, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GatewayForwardWriteRequest) SetTarget(v GatewayForwardTargetRequest)`

SetTarget sets Target field to given value.


### GetAllowedSources

`func (o *GatewayForwardWriteRequest) GetAllowedSources() []string`

GetAllowedSources returns the AllowedSources field if non-nil, zero value otherwise.

### GetAllowedSourcesOk

`func (o *GatewayForwardWriteRequest) GetAllowedSourcesOk() (*[]string, bool)`

GetAllowedSourcesOk returns a tuple with the AllowedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSources

`func (o *GatewayForwardWriteRequest) SetAllowedSources(v []string)`

SetAllowedSources sets AllowedSources field to given value.

### HasAllowedSources

`func (o *GatewayForwardWriteRequest) HasAllowedSources() bool`

HasAllowedSources returns a boolean if a field has been set.

### GetPreserveClientIp

`func (o *GatewayForwardWriteRequest) GetPreserveClientIp() bool`

GetPreserveClientIp returns the PreserveClientIp field if non-nil, zero value otherwise.

### GetPreserveClientIpOk

`func (o *GatewayForwardWriteRequest) GetPreserveClientIpOk() (*bool, bool)`

GetPreserveClientIpOk returns a tuple with the PreserveClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveClientIp

`func (o *GatewayForwardWriteRequest) SetPreserveClientIp(v bool)`

SetPreserveClientIp sets PreserveClientIp field to given value.

### HasPreserveClientIp

`func (o *GatewayForwardWriteRequest) HasPreserveClientIp() bool`

HasPreserveClientIp returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayForwardWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayForwardWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayForwardWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayForwardWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComment

`func (o *GatewayForwardWriteRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GatewayForwardWriteRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GatewayForwardWriteRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GatewayForwardWriteRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


