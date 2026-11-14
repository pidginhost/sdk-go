# PaginatedGatewayLBListenerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]GatewayLBListener**](GatewayLBListener.md) |  | 

## Methods

### NewPaginatedGatewayLBListenerList

`func NewPaginatedGatewayLBListenerList(count int32, results []GatewayLBListener, ) *PaginatedGatewayLBListenerList`

NewPaginatedGatewayLBListenerList instantiates a new PaginatedGatewayLBListenerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGatewayLBListenerListWithDefaults

`func NewPaginatedGatewayLBListenerListWithDefaults() *PaginatedGatewayLBListenerList`

NewPaginatedGatewayLBListenerListWithDefaults instantiates a new PaginatedGatewayLBListenerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedGatewayLBListenerList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedGatewayLBListenerList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedGatewayLBListenerList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedGatewayLBListenerList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedGatewayLBListenerList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedGatewayLBListenerList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedGatewayLBListenerList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedGatewayLBListenerList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedGatewayLBListenerList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedGatewayLBListenerList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedGatewayLBListenerList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedGatewayLBListenerList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedGatewayLBListenerList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedGatewayLBListenerList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedGatewayLBListenerList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedGatewayLBListenerList) GetResults() []GatewayLBListener`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGatewayLBListenerList) GetResultsOk() (*[]GatewayLBListener, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGatewayLBListenerList) SetResults(v []GatewayLBListener)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


