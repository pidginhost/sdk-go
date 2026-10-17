# MobileOAuthSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorEmail** | **string** |  | 
**AccountEmail** | **string** |  | 
**Role** | **string** |  | 
**TokenId** | **NullableInt32** |  | 
**TokenScope** | **string** |  | 
**Currency** | **string** |  | 
**Capabilities** | **map[string]bool** |  | 

## Methods

### NewMobileOAuthSession

`func NewMobileOAuthSession(actorEmail string, accountEmail string, role string, tokenId NullableInt32, tokenScope string, currency string, capabilities map[string]bool, ) *MobileOAuthSession`

NewMobileOAuthSession instantiates a new MobileOAuthSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileOAuthSessionWithDefaults

`func NewMobileOAuthSessionWithDefaults() *MobileOAuthSession`

NewMobileOAuthSessionWithDefaults instantiates a new MobileOAuthSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorEmail

`func (o *MobileOAuthSession) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *MobileOAuthSession) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *MobileOAuthSession) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.


### GetAccountEmail

`func (o *MobileOAuthSession) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *MobileOAuthSession) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *MobileOAuthSession) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.


### GetRole

`func (o *MobileOAuthSession) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MobileOAuthSession) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MobileOAuthSession) SetRole(v string)`

SetRole sets Role field to given value.


### GetTokenId

`func (o *MobileOAuthSession) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MobileOAuthSession) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MobileOAuthSession) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *MobileOAuthSession) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *MobileOAuthSession) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetTokenScope

`func (o *MobileOAuthSession) GetTokenScope() string`

GetTokenScope returns the TokenScope field if non-nil, zero value otherwise.

### GetTokenScopeOk

`func (o *MobileOAuthSession) GetTokenScopeOk() (*string, bool)`

GetTokenScopeOk returns a tuple with the TokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScope

`func (o *MobileOAuthSession) SetTokenScope(v string)`

SetTokenScope sets TokenScope field to given value.


### GetCurrency

`func (o *MobileOAuthSession) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MobileOAuthSession) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MobileOAuthSession) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCapabilities

`func (o *MobileOAuthSession) GetCapabilities() map[string]bool`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MobileOAuthSession) GetCapabilitiesOk() (*map[string]bool, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MobileOAuthSession) SetCapabilities(v map[string]bool)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


