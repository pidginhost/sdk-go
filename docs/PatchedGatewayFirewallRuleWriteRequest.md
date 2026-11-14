# PatchedGatewayFirewallRuleWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 
**Source** | Pointer to [**GatewayFirewallSourceRequest**](GatewayFirewallSourceRequest.md) |  | [optional] 
**Destination** | Pointer to [**GatewayFirewallDestinationRequest**](GatewayFirewallDestinationRequest.md) |  | [optional] 
**Protocol** | Pointer to [**GatewayFirewallProtocolEnum**](GatewayFirewallProtocolEnum.md) |  | [optional] 
**Ports** | Pointer to [**[]GatewayPortRangeRequest**](GatewayPortRangeRequest.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedGatewayFirewallRuleWriteRequest

`func NewPatchedGatewayFirewallRuleWriteRequest() *PatchedGatewayFirewallRuleWriteRequest`

NewPatchedGatewayFirewallRuleWriteRequest instantiates a new PatchedGatewayFirewallRuleWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGatewayFirewallRuleWriteRequestWithDefaults

`func NewPatchedGatewayFirewallRuleWriteRequestWithDefaults() *PatchedGatewayFirewallRuleWriteRequest`

NewPatchedGatewayFirewallRuleWriteRequestWithDefaults instantiates a new PatchedGatewayFirewallRuleWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSource

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetSource() GatewayFirewallSourceRequest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetSourceOk() (*GatewayFirewallSourceRequest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetSource(v GatewayFirewallSourceRequest)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetDestination() GatewayFirewallDestinationRequest`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetDestinationOk() (*GatewayFirewallDestinationRequest, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetDestination(v GatewayFirewallDestinationRequest)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetProtocol() GatewayFirewallProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetProtocolOk() (*GatewayFirewallProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetProtocol(v GatewayFirewallProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetPorts() []GatewayPortRangeRequest`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetPortsOk() (*[]GatewayPortRangeRequest, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetPorts(v []GatewayPortRangeRequest)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComment

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchedGatewayFirewallRuleWriteRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchedGatewayFirewallRuleWriteRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchedGatewayFirewallRuleWriteRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


