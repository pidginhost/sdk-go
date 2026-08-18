# IsoBootRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso** | Pointer to **string** | Catalog ISO slug. Omit it to use the default rescue image. | [optional] 

## Methods

### NewIsoBootRequest

`func NewIsoBootRequest() *IsoBootRequest`

NewIsoBootRequest instantiates a new IsoBootRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsoBootRequestWithDefaults

`func NewIsoBootRequestWithDefaults() *IsoBootRequest`

NewIsoBootRequestWithDefaults instantiates a new IsoBootRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso

`func (o *IsoBootRequest) GetIso() string`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *IsoBootRequest) GetIsoOk() (*string, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *IsoBootRequest) SetIso(v string)`

SetIso sets Iso field to given value.

### HasIso

`func (o *IsoBootRequest) HasIso() bool`

HasIso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


