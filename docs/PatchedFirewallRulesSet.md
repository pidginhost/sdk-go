# PatchedFirewallRulesSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**FirewallRulesSetStatusEnum**](FirewallRulesSetStatusEnum.md) |  | [optional] [readonly] 
**Rules** | Pointer to [**[]FirewallRule**](FirewallRule.md) |  | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | used with free tier vm | [optional] [readonly] 

## Methods

### NewPatchedFirewallRulesSet

`func NewPatchedFirewallRulesSet() *PatchedFirewallRulesSet`

NewPatchedFirewallRulesSet instantiates a new PatchedFirewallRulesSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFirewallRulesSetWithDefaults

`func NewPatchedFirewallRulesSetWithDefaults() *PatchedFirewallRulesSet`

NewPatchedFirewallRulesSetWithDefaults instantiates a new PatchedFirewallRulesSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedFirewallRulesSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedFirewallRulesSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedFirewallRulesSet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedFirewallRulesSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedFirewallRulesSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFirewallRulesSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFirewallRulesSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFirewallRulesSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedFirewallRulesSet) GetStatus() FirewallRulesSetStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedFirewallRulesSet) GetStatusOk() (*FirewallRulesSetStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedFirewallRulesSet) SetStatus(v FirewallRulesSetStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedFirewallRulesSet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRules

`func (o *PatchedFirewallRulesSet) GetRules() []FirewallRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PatchedFirewallRulesSet) GetRulesOk() (*[]FirewallRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PatchedFirewallRulesSet) SetRules(v []FirewallRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PatchedFirewallRulesSet) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetReadOnly

`func (o *PatchedFirewallRulesSet) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PatchedFirewallRulesSet) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PatchedFirewallRulesSet) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PatchedFirewallRulesSet) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


