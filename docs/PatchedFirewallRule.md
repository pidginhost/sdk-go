# PatchedFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Direction** | Pointer to [**DirectionEnum**](DirectionEnum.md) |  | [optional] 
**Action** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Sport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Destination** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Dport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedFirewallRule

`func NewPatchedFirewallRule() *PatchedFirewallRule`

NewPatchedFirewallRule instantiates a new PatchedFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFirewallRuleWithDefaults

`func NewPatchedFirewallRuleWithDefaults() *PatchedFirewallRule`

NewPatchedFirewallRuleWithDefaults instantiates a new PatchedFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedFirewallRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedFirewallRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedFirewallRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *PatchedFirewallRule) GetDirection() DirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PatchedFirewallRule) GetDirectionOk() (*DirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PatchedFirewallRule) SetDirection(v DirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PatchedFirewallRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAction

`func (o *PatchedFirewallRule) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedFirewallRule) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedFirewallRule) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedFirewallRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *PatchedFirewallRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedFirewallRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedFirewallRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedFirewallRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *PatchedFirewallRule) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *PatchedFirewallRule) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *PatchedFirewallRule) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *PatchedFirewallRule) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *PatchedFirewallRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PatchedFirewallRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PatchedFirewallRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PatchedFirewallRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *PatchedFirewallRule) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *PatchedFirewallRule) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *PatchedFirewallRule) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *PatchedFirewallRule) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedFirewallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedFirewallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedFirewallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedFirewallRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedFirewallRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedFirewallRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedFirewallRule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedFirewallRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetHasError

`func (o *PatchedFirewallRule) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *PatchedFirewallRule) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *PatchedFirewallRule) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *PatchedFirewallRule) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PatchedFirewallRule) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedFirewallRule) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedFirewallRule) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedFirewallRule) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


