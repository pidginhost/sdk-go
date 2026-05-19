# APITokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Scope** | Pointer to [**ScopeEnum**](ScopeEnum.md) |  | [optional] 
**Key** | **string** |  | [readonly] 
**Created** | **string** |  | [readonly] 

## Methods

### NewAPITokenCreate

`func NewAPITokenCreate(id int32, name string, key string, created string, ) *APITokenCreate`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


