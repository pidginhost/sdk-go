# APITokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Scope** | [**ScopeEnum**](ScopeEnum.md) |  | [readonly] 
**KeyPrefix** | **string** |  | [readonly] 
**Created** | **string** |  | [readonly] 
**LastUsed** | **NullableString** |  | [readonly] 
**RequestCount** | **int32** |  | [readonly] 

## Methods

### NewAPITokenList

`func NewAPITokenList(id int32, name string, scope ScopeEnum, keyPrefix string, created string, lastUsed NullableString, requestCount int32, ) *APITokenList`

NewAPITokenList instantiates a new APITokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITokenListWithDefaults

`func NewAPITokenListWithDefaults() *APITokenList`

NewAPITokenListWithDefaults instantiates a new APITokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITokenList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITokenList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITokenList) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *APITokenList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APITokenList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APITokenList) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *APITokenList) GetScope() ScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *APITokenList) GetScopeOk() (*ScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *APITokenList) SetScope(v ScopeEnum)`

SetScope sets Scope field to given value.


### GetKeyPrefix

`func (o *APITokenList) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *APITokenList) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *APITokenList) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### GetCreated

`func (o *APITokenList) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *APITokenList) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *APITokenList) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUsed

`func (o *APITokenList) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *APITokenList) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *APITokenList) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.


### SetLastUsedNil

`func (o *APITokenList) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *APITokenList) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetRequestCount

`func (o *APITokenList) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *APITokenList) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *APITokenList) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


