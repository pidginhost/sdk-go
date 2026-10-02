# FirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | [**FirewallRuleDirectionEnum**](FirewallRuleDirectionEnum.md) |  | 
**Action** | [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | 
**Protocol** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Sport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Destination** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Dport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 

## Methods

### NewFirewallRuleRequest

`func NewFirewallRuleRequest(direction FirewallRuleDirectionEnum, action FwPolicyOutEnum, ) *FirewallRuleRequest`

NewFirewallRuleRequest instantiates a new FirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleRequestWithDefaults

`func NewFirewallRuleRequestWithDefaults() *FirewallRuleRequest`

NewFirewallRuleRequestWithDefaults instantiates a new FirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *FirewallRuleRequest) GetDirection() FirewallRuleDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRuleRequest) GetDirectionOk() (*FirewallRuleDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRuleRequest) SetDirection(v FirewallRuleDirectionEnum)`

SetDirection sets Direction field to given value.


### GetAction

`func (o *FirewallRuleRequest) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRuleRequest) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRuleRequest) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.


### GetProtocol

`func (o *FirewallRuleRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRuleRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRuleRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FirewallRuleRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *FirewallRuleRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FirewallRuleRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FirewallRuleRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FirewallRuleRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *FirewallRuleRequest) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *FirewallRuleRequest) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *FirewallRuleRequest) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *FirewallRuleRequest) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *FirewallRuleRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FirewallRuleRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FirewallRuleRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *FirewallRuleRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *FirewallRuleRequest) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *FirewallRuleRequest) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *FirewallRuleRequest) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *FirewallRuleRequest) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetEnabled

`func (o *FirewallRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FirewallRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FirewallRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FirewallRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallRuleRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallRuleRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallRuleRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallRuleRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


