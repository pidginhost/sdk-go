# PatchedLBFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
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
**Created** | Pointer to **string** |  | [optional] [readonly] 
**Updated** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedLBFirewallRule

`func NewPatchedLBFirewallRule() *PatchedLBFirewallRule`

NewPatchedLBFirewallRule instantiates a new PatchedLBFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLBFirewallRuleWithDefaults

`func NewPatchedLBFirewallRuleWithDefaults() *PatchedLBFirewallRule`

NewPatchedLBFirewallRuleWithDefaults instantiates a new PatchedLBFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedLBFirewallRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedLBFirewallRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedLBFirewallRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedLBFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *PatchedLBFirewallRule) GetDirection() LBFirewallRuleDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PatchedLBFirewallRule) GetDirectionOk() (*LBFirewallRuleDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PatchedLBFirewallRule) SetDirection(v LBFirewallRuleDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PatchedLBFirewallRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAction

`func (o *PatchedLBFirewallRule) GetAction() LBFirewallRuleActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedLBFirewallRule) GetActionOk() (*LBFirewallRuleActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedLBFirewallRule) SetAction(v LBFirewallRuleActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedLBFirewallRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedLBFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedLBFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedLBFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedLBFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *PatchedLBFirewallRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedLBFirewallRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedLBFirewallRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedLBFirewallRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *PatchedLBFirewallRule) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *PatchedLBFirewallRule) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *PatchedLBFirewallRule) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *PatchedLBFirewallRule) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *PatchedLBFirewallRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PatchedLBFirewallRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PatchedLBFirewallRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PatchedLBFirewallRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *PatchedLBFirewallRule) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *PatchedLBFirewallRule) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *PatchedLBFirewallRule) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *PatchedLBFirewallRule) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetComment

`func (o *PatchedLBFirewallRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchedLBFirewallRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchedLBFirewallRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchedLBFirewallRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedLBFirewallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedLBFirewallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedLBFirewallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedLBFirewallRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedLBFirewallRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedLBFirewallRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedLBFirewallRule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedLBFirewallRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedLBFirewallRule) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedLBFirewallRule) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedLBFirewallRule) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedLBFirewallRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PatchedLBFirewallRule) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PatchedLBFirewallRule) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PatchedLBFirewallRule) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PatchedLBFirewallRule) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


