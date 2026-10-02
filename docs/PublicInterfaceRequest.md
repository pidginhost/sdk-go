# PublicInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FwRulesSet** | Pointer to **NullableString** | ID or slug | [optional] 
**FwPolicyIn** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 
**FwPolicyOut** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] 

## Methods

### NewPublicInterfaceRequest

`func NewPublicInterfaceRequest() *PublicInterfaceRequest`

NewPublicInterfaceRequest instantiates a new PublicInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicInterfaceRequestWithDefaults

`func NewPublicInterfaceRequestWithDefaults() *PublicInterfaceRequest`

NewPublicInterfaceRequestWithDefaults instantiates a new PublicInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFwRulesSet

`func (o *PublicInterfaceRequest) GetFwRulesSet() string`

GetFwRulesSet returns the FwRulesSet field if non-nil, zero value otherwise.

### GetFwRulesSetOk

`func (o *PublicInterfaceRequest) GetFwRulesSetOk() (*string, bool)`

GetFwRulesSetOk returns a tuple with the FwRulesSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRulesSet

`func (o *PublicInterfaceRequest) SetFwRulesSet(v string)`

SetFwRulesSet sets FwRulesSet field to given value.

### HasFwRulesSet

`func (o *PublicInterfaceRequest) HasFwRulesSet() bool`

HasFwRulesSet returns a boolean if a field has been set.

### SetFwRulesSetNil

`func (o *PublicInterfaceRequest) SetFwRulesSetNil(b bool)`

 SetFwRulesSetNil sets the value for FwRulesSet to be an explicit nil

### UnsetFwRulesSet
`func (o *PublicInterfaceRequest) UnsetFwRulesSet()`

UnsetFwRulesSet ensures that no value is present for FwRulesSet, not even an explicit nil
### GetFwPolicyIn

`func (o *PublicInterfaceRequest) GetFwPolicyIn() FwPolicyOutEnum`

GetFwPolicyIn returns the FwPolicyIn field if non-nil, zero value otherwise.

### GetFwPolicyInOk

`func (o *PublicInterfaceRequest) GetFwPolicyInOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyInOk returns a tuple with the FwPolicyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyIn

`func (o *PublicInterfaceRequest) SetFwPolicyIn(v FwPolicyOutEnum)`

SetFwPolicyIn sets FwPolicyIn field to given value.

### HasFwPolicyIn

`func (o *PublicInterfaceRequest) HasFwPolicyIn() bool`

HasFwPolicyIn returns a boolean if a field has been set.

### GetFwPolicyOut

`func (o *PublicInterfaceRequest) GetFwPolicyOut() FwPolicyOutEnum`

GetFwPolicyOut returns the FwPolicyOut field if non-nil, zero value otherwise.

### GetFwPolicyOutOk

`func (o *PublicInterfaceRequest) GetFwPolicyOutOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyOutOk returns a tuple with the FwPolicyOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyOut

`func (o *PublicInterfaceRequest) SetFwPolicyOut(v FwPolicyOutEnum)`

SetFwPolicyOut sets FwPolicyOut field to given value.

### HasFwPolicyOut

`func (o *PublicInterfaceRequest) HasFwPolicyOut() bool`

HasFwPolicyOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


