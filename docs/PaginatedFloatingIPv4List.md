# PaginatedFloatingIPv4List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]FloatingIPv4**](FloatingIPv4.md) |  | 

## Methods

### NewPaginatedFloatingIPv4List

`func NewPaginatedFloatingIPv4List(count int32, results []FloatingIPv4, ) *PaginatedFloatingIPv4List`

NewPaginatedFloatingIPv4List instantiates a new PaginatedFloatingIPv4List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedFloatingIPv4ListWithDefaults

`func NewPaginatedFloatingIPv4ListWithDefaults() *PaginatedFloatingIPv4List`

NewPaginatedFloatingIPv4ListWithDefaults instantiates a new PaginatedFloatingIPv4List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedFloatingIPv4List) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedFloatingIPv4List) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedFloatingIPv4List) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedFloatingIPv4List) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedFloatingIPv4List) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedFloatingIPv4List) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedFloatingIPv4List) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedFloatingIPv4List) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedFloatingIPv4List) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedFloatingIPv4List) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedFloatingIPv4List) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedFloatingIPv4List) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedFloatingIPv4List) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedFloatingIPv4List) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedFloatingIPv4List) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedFloatingIPv4List) GetResults() []FloatingIPv4`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedFloatingIPv4List) GetResultsOk() (*[]FloatingIPv4, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedFloatingIPv4List) SetResults(v []FloatingIPv4)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


