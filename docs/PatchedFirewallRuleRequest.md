# PatchedFirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**FirewallRuleDirectionEnum**](FirewallRuleDirectionEnum.md) |  | [optional] 
**Action** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Sport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Destination** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Dport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedFirewallRuleRequest

`func NewPatchedFirewallRuleRequest() *PatchedFirewallRuleRequest`

NewPatchedFirewallRuleRequest instantiates a new PatchedFirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFirewallRuleRequestWithDefaults

`func NewPatchedFirewallRuleRequestWithDefaults() *PatchedFirewallRuleRequest`

NewPatchedFirewallRuleRequestWithDefaults instantiates a new PatchedFirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *PatchedFirewallRuleRequest) GetDirection() FirewallRuleDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PatchedFirewallRuleRequest) GetDirectionOk() (*FirewallRuleDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PatchedFirewallRuleRequest) SetDirection(v FirewallRuleDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PatchedFirewallRuleRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAction

`func (o *PatchedFirewallRuleRequest) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedFirewallRuleRequest) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedFirewallRuleRequest) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedFirewallRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedFirewallRuleRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedFirewallRuleRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedFirewallRuleRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedFirewallRuleRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *PatchedFirewallRuleRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedFirewallRuleRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedFirewallRuleRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedFirewallRuleRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *PatchedFirewallRuleRequest) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *PatchedFirewallRuleRequest) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *PatchedFirewallRuleRequest) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *PatchedFirewallRuleRequest) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *PatchedFirewallRuleRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PatchedFirewallRuleRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PatchedFirewallRuleRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PatchedFirewallRuleRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *PatchedFirewallRuleRequest) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *PatchedFirewallRuleRequest) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *PatchedFirewallRuleRequest) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *PatchedFirewallRuleRequest) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedFirewallRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedFirewallRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedFirewallRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedFirewallRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedFirewallRuleRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedFirewallRuleRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedFirewallRuleRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedFirewallRuleRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


