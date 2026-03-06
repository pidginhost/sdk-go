# LBFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
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
**Created** | **time.Time** |  | [readonly] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewLBFirewallRule

`func NewLBFirewallRule(id int32, created time.Time, updated time.Time, ) *LBFirewallRule`

NewLBFirewallRule instantiates a new LBFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBFirewallRuleWithDefaults

`func NewLBFirewallRuleWithDefaults() *LBFirewallRule`

NewLBFirewallRuleWithDefaults instantiates a new LBFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LBFirewallRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LBFirewallRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LBFirewallRule) SetId(v int32)`

SetId sets Id field to given value.


### GetDirection

`func (o *LBFirewallRule) GetDirection() LBFirewallRuleDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LBFirewallRule) GetDirectionOk() (*LBFirewallRuleDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LBFirewallRule) SetDirection(v LBFirewallRuleDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LBFirewallRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAction

`func (o *LBFirewallRule) GetAction() LBFirewallRuleActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LBFirewallRule) GetActionOk() (*LBFirewallRuleActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LBFirewallRule) SetAction(v LBFirewallRuleActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *LBFirewallRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProtocol

`func (o *LBFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LBFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LBFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LBFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *LBFirewallRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LBFirewallRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LBFirewallRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LBFirewallRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *LBFirewallRule) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *LBFirewallRule) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *LBFirewallRule) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *LBFirewallRule) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *LBFirewallRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LBFirewallRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LBFirewallRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *LBFirewallRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *LBFirewallRule) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *LBFirewallRule) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *LBFirewallRule) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *LBFirewallRule) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetComment

`func (o *LBFirewallRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LBFirewallRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LBFirewallRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LBFirewallRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnabled

`func (o *LBFirewallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LBFirewallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LBFirewallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LBFirewallRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *LBFirewallRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *LBFirewallRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *LBFirewallRule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *LBFirewallRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCreated

`func (o *LBFirewallRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LBFirewallRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LBFirewallRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *LBFirewallRule) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LBFirewallRule) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LBFirewallRule) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


