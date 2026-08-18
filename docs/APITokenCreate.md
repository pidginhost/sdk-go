# APITokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Scope** | Pointer to [**ScopeEnum**](ScopeEnum.md) |  | [optional] 
**Key** | **string** |  | [readonly] 
**Created** | **string** |  | [readonly] 
**Account** | **NullableString** |  | [readonly] 
**MembershipStatus** | **NullableString** |  | [readonly] 

## Methods

### NewAPITokenCreate

`func NewAPITokenCreate(id int32, name string, key string, created string, account NullableString, membershipStatus NullableString, ) *APITokenCreate`

NewAPITokenCreate instantiates a new APITokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITokenCreateWithDefaults

`func NewAPITokenCreateWithDefaults() *APITokenCreate`

NewAPITokenCreateWithDefaults instantiates a new APITokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITokenCreate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITokenCreate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITokenCreate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *APITokenCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APITokenCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APITokenCreate) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *APITokenCreate) GetScope() ScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *APITokenCreate) GetScopeOk() (*ScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *APITokenCreate) SetScope(v ScopeEnum)`

SetScope sets Scope field to given value.

### HasScope

`func (o *APITokenCreate) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetKey

`func (o *APITokenCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *APITokenCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *APITokenCreate) SetKey(v string)`

SetKey sets Key field to given value.


### GetCreated

`func (o *APITokenCreate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *APITokenCreate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *APITokenCreate) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetAccount

`func (o *APITokenCreate) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *APITokenCreate) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *APITokenCreate) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *APITokenCreate) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *APITokenCreate) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetMembershipStatus

`func (o *APITokenCreate) GetMembershipStatus() string`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *APITokenCreate) GetMembershipStatusOk() (*string, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *APITokenCreate) SetMembershipStatus(v string)`

SetMembershipStatus sets MembershipStatus field to given value.


### SetMembershipStatusNil

`func (o *APITokenCreate) SetMembershipStatusNil(b bool)`

 SetMembershipStatusNil sets the value for MembershipStatus to be an explicit nil

### UnsetMembershipStatus
`func (o *APITokenCreate) UnsetMembershipStatus()`

UnsetMembershipStatus ensures that no value is present for MembershipStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


