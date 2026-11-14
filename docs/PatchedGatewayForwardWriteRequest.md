# PatchedGatewayForwardWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**GatewayForwardProtocolEnum**](GatewayForwardProtocolEnum.md) |  | [optional] 
**PublicPorts** | Pointer to [**GatewayPortRangeRequest**](GatewayPortRangeRequest.md) |  | [optional] 
**Target** | Pointer to [**GatewayForwardTargetRequest**](GatewayForwardTargetRequest.md) |  | [optional] 
**AllowedSources** | Pointer to **[]string** | IPv4 networks; empty means any source. | [optional] 
**PreserveClientIp** | Pointer to **bool** | Only false is accepted for now. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedGatewayForwardWriteRequest

`func NewPatchedGatewayForwardWriteRequest() *PatchedGatewayForwardWriteRequest`

NewPatchedGatewayForwardWriteRequest instantiates a new PatchedGatewayForwardWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGatewayForwardWriteRequestWithDefaults

`func NewPatchedGatewayForwardWriteRequestWithDefaults() *PatchedGatewayForwardWriteRequest`

NewPatchedGatewayForwardWriteRequestWithDefaults instantiates a new PatchedGatewayForwardWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *PatchedGatewayForwardWriteRequest) GetProtocol() GatewayForwardProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedGatewayForwardWriteRequest) GetProtocolOk() (*GatewayForwardProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedGatewayForwardWriteRequest) SetProtocol(v GatewayForwardProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedGatewayForwardWriteRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublicPorts

`func (o *PatchedGatewayForwardWriteRequest) GetPublicPorts() GatewayPortRangeRequest`

GetPublicPorts returns the PublicPorts field if non-nil, zero value otherwise.

### GetPublicPortsOk

`func (o *PatchedGatewayForwardWriteRequest) GetPublicPortsOk() (*GatewayPortRangeRequest, bool)`

GetPublicPortsOk returns a tuple with the PublicPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPorts

`func (o *PatchedGatewayForwardWriteRequest) SetPublicPorts(v GatewayPortRangeRequest)`

SetPublicPorts sets PublicPorts field to given value.

### HasPublicPorts

`func (o *PatchedGatewayForwardWriteRequest) HasPublicPorts() bool`

HasPublicPorts returns a boolean if a field has been set.

### GetTarget

`func (o *PatchedGatewayForwardWriteRequest) GetTarget() GatewayForwardTargetRequest`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchedGatewayForwardWriteRequest) GetTargetOk() (*GatewayForwardTargetRequest, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchedGatewayForwardWriteRequest) SetTarget(v GatewayForwardTargetRequest)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchedGatewayForwardWriteRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetAllowedSources

`func (o *PatchedGatewayForwardWriteRequest) GetAllowedSources() []string`

GetAllowedSources returns the AllowedSources field if non-nil, zero value otherwise.

### GetAllowedSourcesOk

`func (o *PatchedGatewayForwardWriteRequest) GetAllowedSourcesOk() (*[]string, bool)`

GetAllowedSourcesOk returns a tuple with the AllowedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSources

`func (o *PatchedGatewayForwardWriteRequest) SetAllowedSources(v []string)`

SetAllowedSources sets AllowedSources field to given value.

### HasAllowedSources

`func (o *PatchedGatewayForwardWriteRequest) HasAllowedSources() bool`

HasAllowedSources returns a boolean if a field has been set.

### GetPreserveClientIp

`func (o *PatchedGatewayForwardWriteRequest) GetPreserveClientIp() bool`

GetPreserveClientIp returns the PreserveClientIp field if non-nil, zero value otherwise.

### GetPreserveClientIpOk

`func (o *PatchedGatewayForwardWriteRequest) GetPreserveClientIpOk() (*bool, bool)`

GetPreserveClientIpOk returns a tuple with the PreserveClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveClientIp

`func (o *PatchedGatewayForwardWriteRequest) SetPreserveClientIp(v bool)`

SetPreserveClientIp sets PreserveClientIp field to given value.

### HasPreserveClientIp

`func (o *PatchedGatewayForwardWriteRequest) HasPreserveClientIp() bool`

HasPreserveClientIp returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedGatewayForwardWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedGatewayForwardWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedGatewayForwardWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedGatewayForwardWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComment

`func (o *PatchedGatewayForwardWriteRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchedGatewayForwardWriteRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchedGatewayForwardWriteRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchedGatewayForwardWriteRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


