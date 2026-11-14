# GatewayFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Position** | **int32** |  | [readonly] 
**Action** | [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [readonly] 
**Source** | [**GatewayFirewallSource**](GatewayFirewallSource.md) |  | [readonly] 
**Destination** | [**GatewayFirewallDestination**](GatewayFirewallDestination.md) |  | [readonly] 
**Protocol** | [**GatewayFirewallProtocolEnum**](GatewayFirewallProtocolEnum.md) |  | [readonly] 
**Ports** | [**[]GatewayPortRange**](GatewayPortRange.md) |  | [readonly] 
**Enabled** | **bool** |  | [readonly] 
**Comment** | **string** |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 

## Methods

### NewGatewayFirewallRule

`func NewGatewayFirewallRule(id int32, position int32, action FwPolicyOutEnum, source GatewayFirewallSource, destination GatewayFirewallDestination, protocol GatewayFirewallProtocolEnum, ports []GatewayPortRange, enabled bool, comment string, changedRevision int32, pending bool, ) *GatewayFirewallRule`

NewGatewayFirewallRule instantiates a new GatewayFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayFirewallRuleWithDefaults

`func NewGatewayFirewallRuleWithDefaults() *GatewayFirewallRule`

NewGatewayFirewallRuleWithDefaults instantiates a new GatewayFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayFirewallRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayFirewallRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayFirewallRule) SetId(v int32)`

SetId sets Id field to given value.


### GetPosition

`func (o *GatewayFirewallRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GatewayFirewallRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GatewayFirewallRule) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetAction

`func (o *GatewayFirewallRule) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GatewayFirewallRule) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GatewayFirewallRule) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.


### GetSource

`func (o *GatewayFirewallRule) GetSource() GatewayFirewallSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GatewayFirewallRule) GetSourceOk() (*GatewayFirewallSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GatewayFirewallRule) SetSource(v GatewayFirewallSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *GatewayFirewallRule) GetDestination() GatewayFirewallDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GatewayFirewallRule) GetDestinationOk() (*GatewayFirewallDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GatewayFirewallRule) SetDestination(v GatewayFirewallDestination)`

SetDestination sets Destination field to given value.


### GetProtocol

`func (o *GatewayFirewallRule) GetProtocol() GatewayFirewallProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayFirewallRule) GetProtocolOk() (*GatewayFirewallProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayFirewallRule) SetProtocol(v GatewayFirewallProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetPorts

`func (o *GatewayFirewallRule) GetPorts() []GatewayPortRange`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GatewayFirewallRule) GetPortsOk() (*[]GatewayPortRange, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GatewayFirewallRule) SetPorts(v []GatewayPortRange)`

SetPorts sets Ports field to given value.


### GetEnabled

`func (o *GatewayFirewallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayFirewallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayFirewallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetComment

`func (o *GatewayFirewallRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GatewayFirewallRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GatewayFirewallRule) SetComment(v string)`

SetComment sets Comment field to given value.


### GetChangedRevision

`func (o *GatewayFirewallRule) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayFirewallRule) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayFirewallRule) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayFirewallRule) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayFirewallRule) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayFirewallRule) SetPending(v bool)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


