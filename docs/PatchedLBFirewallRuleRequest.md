# PatchedLBFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**LBFirewallRuleDirectionEnum**](LBFirewallRuleDirectionEnum.md) |  | [optional] 
**Action** | Pointer to [**LBFirewallRuleActionEnum**](LBFirewallRuleActionEnum.md) |  | [optional] 
**Protocol** | Pointer to **string** | tcp, udp, icmp, etc. | [optional] 
**Source** | Pointer to **string** | IP address or CIDR | [optional] 
**Sport** | Pointer to **string** | Port or range (e.g., 1024-65535) | [optional] 
**Destination** | Pointer to **string** | IP address or CIDR | [optional] 
**Dport** | Pointer to **string** | Port or range (e.g., 80, 8000-9000) | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** | Rule order (lower &#x3D; higher priority) | [optional] 

## Methods

### NewPatchedLBFirewallRuleRequest

`func NewPatchedLBFirewallRuleRequest() *PatchedLBFirewallRuleRequest`

NewPatchedLBFirewallRuleRequest instantiates a new PatchedLBFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLBFirewallRuleRequestWithDefaults

`func NewPatchedLBFirewallRuleRequestWithDefaults() *PatchedLBFirewallRuleRequest`

NewPatchedLBFirewallRuleRequestWithDefaults instantiates a new PatchedLBFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *PatchedLBFirewallRuleRequest) GetDirection() LBFirewallRuleDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PatchedLBFirewallRuleRequest) GetDirectionOk() (*LBFirewallRuleDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PatchedLBFirewallRuleRequest) SetDirection(v LBFirewallRuleDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PatchedLBFirewallRuleRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAction

`func (o *PatchedLBFirewallRuleRequest) GetAction() LBFirewallRuleActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedLBFirewallRuleRequest) GetActionOk() (*LBFirewallRuleActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedLBFirewallRuleRequest) SetAction(v LBFirewallRuleActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedLBFirewallRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedLBFirewallRuleRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedLBFirewallRuleRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedLBFirewallRuleRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedLBFirewallRuleRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *PatchedLBFirewallRuleRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedLBFirewallRuleRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedLBFirewallRuleRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedLBFirewallRuleRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *PatchedLBFirewallRuleRequest) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *PatchedLBFirewallRuleRequest) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *PatchedLBFirewallRuleRequest) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *PatchedLBFirewallRuleRequest) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *PatchedLBFirewallRuleRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PatchedLBFirewallRuleRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PatchedLBFirewallRuleRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PatchedLBFirewallRuleRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *PatchedLBFirewallRuleRequest) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *PatchedLBFirewallRuleRequest) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *PatchedLBFirewallRuleRequest) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *PatchedLBFirewallRuleRequest) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetComment

`func (o *PatchedLBFirewallRuleRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchedLBFirewallRuleRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchedLBFirewallRuleRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchedLBFirewallRuleRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedLBFirewallRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedLBFirewallRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedLBFirewallRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedLBFirewallRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedLBFirewallRuleRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedLBFirewallRuleRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedLBFirewallRuleRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedLBFirewallRuleRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


