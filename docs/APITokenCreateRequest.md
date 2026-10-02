# APITokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Scope** | Pointer to [**ScopeEnum**](ScopeEnum.md) |  | [optional] 

## Methods

### NewAPITokenCreateRequest

`func NewAPITokenCreateRequest(name string, ) *APITokenCreateRequest`

NewAPITokenCreateRequest instantiates a new APITokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITokenCreateRequestWithDefaults

`func NewAPITokenCreateRequestWithDefaults() *APITokenCreateRequest`

NewAPITokenCreateRequestWithDefaults instantiates a new APITokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *APITokenCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APITokenCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APITokenCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *APITokenCreateRequest) GetScope() ScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *APITokenCreateRequest) GetScopeOk() (*ScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *APITokenCreateRequest) SetScope(v ScopeEnum)`

SetScope sets Scope field to given value.

### HasScope

`func (o *APITokenCreateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


