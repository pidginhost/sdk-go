# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Direction** | [**DirectionEnum**](DirectionEnum.md) |  | 
**Action** | [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | 
**Protocol** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Sport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Destination** | Pointer to **string** | single IP, range (20.34.101.207-201.3.9.99) or comma separated list | [optional] 
**Dport** | Pointer to **string** | numbers (0-65535), range (\&quot;\\d+:\\d+\&quot;, like \&quot;80:85\&quot;), comma separated list | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**HasError** | **bool** |  | [readonly] 
**ErrorMessage** | **string** |  | [readonly] 

## Methods

### NewFirewallRule

`func NewFirewallRule(id int32, direction DirectionEnum, action FwPolicyOutEnum, hasError bool, errorMessage string, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRule) SetId(v int32)`

SetId sets Id field to given value.


### GetDirection

`func (o *FirewallRule) GetDirection() DirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRule) GetDirectionOk() (*DirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRule) SetDirection(v DirectionEnum)`

SetDirection sets Direction field to given value.


### GetAction

`func (o *FirewallRule) GetAction() FwPolicyOutEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRule) GetActionOk() (*FwPolicyOutEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRule) SetAction(v FwPolicyOutEnum)`

SetAction sets Action field to given value.


### GetProtocol

`func (o *FirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *FirewallRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FirewallRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FirewallRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FirewallRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSport

`func (o *FirewallRule) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *FirewallRule) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *FirewallRule) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *FirewallRule) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetDestination

`func (o *FirewallRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FirewallRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FirewallRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *FirewallRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDport

`func (o *FirewallRule) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *FirewallRule) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *FirewallRule) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *FirewallRule) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetEnabled

`func (o *FirewallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FirewallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FirewallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FirewallRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *FirewallRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FirewallRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FirewallRule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FirewallRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetHasError

`func (o *FirewallRule) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *FirewallRule) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *FirewallRule) SetHasError(v bool)`

SetHasError sets HasError field to given value.


### GetErrorMessage

`func (o *FirewallRule) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FirewallRule) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FirewallRule) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


