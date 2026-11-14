# GatewayFirewallRuleWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | 
**Source** | [**GatewayFirewallSourceRequest**](GatewayFirewallSourceRequest.md) |  | 
**Destination** | [**GatewayFirewallDestinationRequest**](GatewayFirewallDestinationRequest.md) |  | 
**Protocol** | [**GatewayFirewallProtocolEnum**](GatewayFirewallProtocolEnum.md) |  | 
**Ports** | Pointer to [**[]GatewayPortRangeRequest**](GatewayPortRangeRequest.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayFirewallRuleWriteRequest

`func NewGatewayFirewallRuleWriteRequest(action FwPolicyOutEnum, source GatewayFirewallSourceRequest, destination GatewayFirewallDestinationRequest, protocol GatewayFirewallProtocolEnum, ) *GatewayFirewallRuleWriteRequest`

NewGatewayFirewallRuleWriteRequest instantiates a new GatewayFirewallRuleWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayFirewallRuleWriteRequestWithDefaults

`func NewGatewayFirewallRuleWriteRequestWithDefaults() *GatewayFirewallRuleWriteRequest`

NewGatewayFirewallRuleWriteRequestWithDefaults instantiates a new GatewayFirewallRuleWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GatewayFirewallRuleWriteRequest) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GatewayFirewallRuleWriteRequest) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GatewayFirewallRuleWriteRequest) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.


### GetSource

`func (o *GatewayFirewallRuleWriteRequest) GetSource() GatewayFirewallSourceRequest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GatewayFirewallRuleWriteRequest) GetSourceOk() (*GatewayFirewallSourceRequest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GatewayFirewallRuleWriteRequest) SetSource(v GatewayFirewallSourceRequest)`

SetSource sets Source field to given value.


### GetDestination

`func (o *GatewayFirewallRuleWriteRequest) GetDestination() GatewayFirewallDestinationRequest`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GatewayFirewallRuleWriteRequest) GetDestinationOk() (*GatewayFirewallDestinationRequest, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GatewayFirewallRuleWriteRequest) SetDestination(v GatewayFirewallDestinationRequest)`

SetDestination sets Destination field to given value.


### GetProtocol

`func (o *GatewayFirewallRuleWriteRequest) GetProtocol() GatewayFirewallProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayFirewallRuleWriteRequest) GetProtocolOk() (*GatewayFirewallProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayFirewallRuleWriteRequest) SetProtocol(v GatewayFirewallProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetPorts

`func (o *GatewayFirewallRuleWriteRequest) GetPorts() []GatewayPortRangeRequest`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GatewayFirewallRuleWriteRequest) GetPortsOk() (*[]GatewayPortRangeRequest, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GatewayFirewallRuleWriteRequest) SetPorts(v []GatewayPortRangeRequest)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GatewayFirewallRuleWriteRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayFirewallRuleWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayFirewallRuleWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayFirewallRuleWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayFirewallRuleWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComment

`func (o *GatewayFirewallRuleWriteRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GatewayFirewallRuleWriteRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GatewayFirewallRuleWriteRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GatewayFirewallRuleWriteRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


