# AccountSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorEmail** | **string** |  | 
**AccountEmail** | **string** |  | 
**Role** | **string** |  | 
**TokenId** | **NullableInt32** |  | 
**TokenScope** | **NullableString** |  | 
**Capabilities** | **map[string]bool** |  | 
**Currency** | **string** |  | 

## Methods

### NewAccountSessionResponse

`func NewAccountSessionResponse(actorEmail string, accountEmail string, role string, tokenId NullableInt32, tokenScope NullableString, capabilities map[string]bool, currency string, ) *AccountSessionResponse`

NewAccountSessionResponse instantiates a new AccountSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSessionResponseWithDefaults

`func NewAccountSessionResponseWithDefaults() *AccountSessionResponse`

NewAccountSessionResponseWithDefaults instantiates a new AccountSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorEmail

`func (o *AccountSessionResponse) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *AccountSessionResponse) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *AccountSessionResponse) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.


### GetAccountEmail

`func (o *AccountSessionResponse) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *AccountSessionResponse) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *AccountSessionResponse) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.


### GetRole

`func (o *AccountSessionResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccountSessionResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccountSessionResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetTokenId

`func (o *AccountSessionResponse) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AccountSessionResponse) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AccountSessionResponse) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *AccountSessionResponse) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *AccountSessionResponse) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetTokenScope

`func (o *AccountSessionResponse) GetTokenScope() string`

GetTokenScope returns the TokenScope field if non-nil, zero value otherwise.

### GetTokenScopeOk

`func (o *AccountSessionResponse) GetTokenScopeOk() (*string, bool)`

GetTokenScopeOk returns a tuple with the TokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScope

`func (o *AccountSessionResponse) SetTokenScope(v string)`

SetTokenScope sets TokenScope field to given value.


### SetTokenScopeNil

`func (o *AccountSessionResponse) SetTokenScopeNil(b bool)`

 SetTokenScopeNil sets the value for TokenScope to be an explicit nil

### UnsetTokenScope
`func (o *AccountSessionResponse) UnsetTokenScope()`

UnsetTokenScope ensures that no value is present for TokenScope, not even an explicit nil
### GetCapabilities

`func (o *AccountSessionResponse) GetCapabilities() map[string]bool`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountSessionResponse) GetCapabilitiesOk() (*map[string]bool, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountSessionResponse) SetCapabilities(v map[string]bool)`

SetCapabilities sets Capabilities field to given value.


### GetCurrency

`func (o *AccountSessionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountSessionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountSessionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


