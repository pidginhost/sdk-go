# FirewallRulesSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Status** | [**FirewallRulesSetStatusEnum**](FirewallRulesSetStatusEnum.md) |  | [readonly] 
**Rules** | [**[]FirewallRule**](FirewallRule.md) |  | [readonly] 
**ReadOnly** | **bool** | used with free tier vm | [readonly] 

## Methods

### NewFirewallRulesSet

`func NewFirewallRulesSet(id int32, name string, status FirewallRulesSetStatusEnum, rules []FirewallRule, readOnly bool, ) *FirewallRulesSet`

NewFirewallRulesSet instantiates a new FirewallRulesSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRulesSetWithDefaults

`func NewFirewallRulesSetWithDefaults() *FirewallRulesSet`

NewFirewallRulesSetWithDefaults instantiates a new FirewallRulesSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRulesSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRulesSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRulesSet) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FirewallRulesSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRulesSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRulesSet) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *FirewallRulesSet) GetStatus() FirewallRulesSetStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRulesSet) GetStatusOk() (*FirewallRulesSetStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRulesSet) SetStatus(v FirewallRulesSetStatusEnum)`

SetStatus sets Status field to given value.


### GetRules

`func (o *FirewallRulesSet) GetRules() []FirewallRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallRulesSet) GetRulesOk() (*[]FirewallRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallRulesSet) SetRules(v []FirewallRule)`

SetRules sets Rules field to given value.


### GetReadOnly

`func (o *FirewallRulesSet) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FirewallRulesSet) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FirewallRulesSet) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


