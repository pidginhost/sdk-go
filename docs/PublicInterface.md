# PublicInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | **string** |  | [readonly] 
**Ipv4** | **string** |  | [readonly] 
**Ipv6** | **string** |  | [readonly] 
**FwRulesSet** | Pointer to **NullableString** | ID or slug | [optional] 
**FwPolicyIn** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 
**FwPolicyOut** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 

## Methods

### NewPublicInterface

`func NewPublicInterface(interface_ string, ipv4 string, ipv6 string, ) *PublicInterface`

NewPublicInterface instantiates a new PublicInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicInterfaceWithDefaults

`func NewPublicInterfaceWithDefaults() *PublicInterface`

NewPublicInterfaceWithDefaults instantiates a new PublicInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *PublicInterface) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PublicInterface) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PublicInterface) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetIpv4

`func (o *PublicInterface) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PublicInterface) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PublicInterface) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *PublicInterface) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *PublicInterface) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *PublicInterface) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.


### GetFwRulesSet

`func (o *PublicInterface) GetFwRulesSet() string`

GetFwRulesSet returns the FwRulesSet field if non-nil, zero value otherwise.

### GetFwRulesSetOk

`func (o *PublicInterface) GetFwRulesSetOk() (*string, bool)`

GetFwRulesSetOk returns a tuple with the FwRulesSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRulesSet

`func (o *PublicInterface) SetFwRulesSet(v string)`

SetFwRulesSet sets FwRulesSet field to given value.

### HasFwRulesSet

`func (o *PublicInterface) HasFwRulesSet() bool`

HasFwRulesSet returns a boolean if a field has been set.

### SetFwRulesSetNil

`func (o *PublicInterface) SetFwRulesSetNil(b bool)`

 SetFwRulesSetNil sets the value for FwRulesSet to be an explicit nil

### UnsetFwRulesSet
`func (o *PublicInterface) UnsetFwRulesSet()`

UnsetFwRulesSet ensures that no value is present for FwRulesSet, not even an explicit nil
### GetFwPolicyIn

`func (o *PublicInterface) GetFwPolicyIn() FwPolicyOutEnum`

GetFwPolicyIn returns the FwPolicyIn field if non-nil, zero value otherwise.

### GetFwPolicyInOk

`func (o *PublicInterface) GetFwPolicyInOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyInOk returns a tuple with the FwPolicyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyIn

`func (o *PublicInterface) SetFwPolicyIn(v FwPolicyOutEnum)`

SetFwPolicyIn sets FwPolicyIn field to given value.

### HasFwPolicyIn

`func (o *PublicInterface) HasFwPolicyIn() bool`

HasFwPolicyIn returns a boolean if a field has been set.

### GetFwPolicyOut

`func (o *PublicInterface) GetFwPolicyOut() FwPolicyOutEnum`

GetFwPolicyOut returns the FwPolicyOut field if non-nil, zero value otherwise.

### GetFwPolicyOutOk

`func (o *PublicInterface) GetFwPolicyOutOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyOutOk returns a tuple with the FwPolicyOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyOut

`func (o *PublicInterface) SetFwPolicyOut(v FwPolicyOutEnum)`

SetFwPolicyOut sets FwPolicyOut field to given value.

### HasFwPolicyOut

`func (o *PublicInterface) HasFwPolicyOut() bool`

HasFwPolicyOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


