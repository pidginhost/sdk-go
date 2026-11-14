# GatewayForward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Protocol** | [**GatewayForwardProtocolEnum**](GatewayForwardProtocolEnum.md) |  | [readonly] 
**PublicPorts** | [**GatewayPortRange**](GatewayPortRange.md) |  | [readonly] 
**AllowedSources** | **interface{}** |  | [readonly] 
**Target** | [**GatewayForwardTarget**](GatewayForwardTarget.md) |  | [readonly] 
**PreserveClientIp** | **bool** | Always false: the target sees the gateway&#39;s address. | [readonly] 
**Enabled** | **bool** |  | [readonly] 
**Comment** | **string** |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 

## Methods

### NewGatewayForward

`func NewGatewayForward(id int32, protocol GatewayForwardProtocolEnum, publicPorts GatewayPortRange, allowedSources interface{}, target GatewayForwardTarget, preserveClientIp bool, enabled bool, comment string, changedRevision int32, pending bool, ) *GatewayForward`

NewGatewayForward instantiates a new GatewayForward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayForwardWithDefaults

`func NewGatewayForwardWithDefaults() *GatewayForward`

NewGatewayForwardWithDefaults instantiates a new GatewayForward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayForward) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayForward) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayForward) SetId(v int32)`

SetId sets Id field to given value.


### GetProtocol

`func (o *GatewayForward) GetProtocol() GatewayForwardProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayForward) GetProtocolOk() (*GatewayForwardProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayForward) SetProtocol(v GatewayForwardProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetPublicPorts

`func (o *GatewayForward) GetPublicPorts() GatewayPortRange`

GetPublicPorts returns the PublicPorts field if non-nil, zero value otherwise.

### GetPublicPortsOk

`func (o *GatewayForward) GetPublicPortsOk() (*GatewayPortRange, bool)`

GetPublicPortsOk returns a tuple with the PublicPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPorts

`func (o *GatewayForward) SetPublicPorts(v GatewayPortRange)`

SetPublicPorts sets PublicPorts field to given value.


### GetAllowedSources

`func (o *GatewayForward) GetAllowedSources() interface{}`

GetAllowedSources returns the AllowedSources field if non-nil, zero value otherwise.

### GetAllowedSourcesOk

`func (o *GatewayForward) GetAllowedSourcesOk() (*interface{}, bool)`

GetAllowedSourcesOk returns a tuple with the AllowedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSources

`func (o *GatewayForward) SetAllowedSources(v interface{})`

SetAllowedSources sets AllowedSources field to given value.


### SetAllowedSourcesNil

`func (o *GatewayForward) SetAllowedSourcesNil(b bool)`

 SetAllowedSourcesNil sets the value for AllowedSources to be an explicit nil

### UnsetAllowedSources
`func (o *GatewayForward) UnsetAllowedSources()`

UnsetAllowedSources ensures that no value is present for AllowedSources, not even an explicit nil
### GetTarget

`func (o *GatewayForward) GetTarget() GatewayForwardTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GatewayForward) GetTargetOk() (*GatewayForwardTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GatewayForward) SetTarget(v GatewayForwardTarget)`

SetTarget sets Target field to given value.


### GetPreserveClientIp

`func (o *GatewayForward) GetPreserveClientIp() bool`

GetPreserveClientIp returns the PreserveClientIp field if non-nil, zero value otherwise.

### GetPreserveClientIpOk

`func (o *GatewayForward) GetPreserveClientIpOk() (*bool, bool)`

GetPreserveClientIpOk returns a tuple with the PreserveClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveClientIp

`func (o *GatewayForward) SetPreserveClientIp(v bool)`

SetPreserveClientIp sets PreserveClientIp field to given value.


### GetEnabled

`func (o *GatewayForward) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayForward) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayForward) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetComment

`func (o *GatewayForward) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GatewayForward) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GatewayForward) SetComment(v string)`

SetComment sets Comment field to given value.


### GetChangedRevision

`func (o *GatewayForward) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayForward) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayForward) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayForward) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayForward) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayForward) SetPending(v bool)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


