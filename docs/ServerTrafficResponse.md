# ServerTrafficResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **int32** |  | 
**Month** | **int32** |  | 
**AsOf** | **NullableString** |  | 
**Status** | **string** |  | 
**BytesIn** | **int32** |  | 
**BytesOut** | **int32** |  | 
**BytesTotal** | **int32** |  | 
**IncludedTb** | **NullableInt32** |  | 
**UsedUnits** | **int32** |  | 
**UsedTb** | **string** | Usage rounded up to six decimal places; bytes_total is exact. | 
**BillableBytes** | **int32** | Of bytes_total, the part that may be charged. | 
**BillableUnits** | **int32** |  | 
**BillableTb** | **string** | Billable usage rounded up to six decimal places; billable_bytes is exact. | 
**BillableFrom** | **NullableString** | First fully billable day, when one date describes the usage. May fall after the reported month. Null when all usage is billable or streams have different boundaries; use billable_bytes for the billable total. | 
**ChargedTb** | **int32** |  | 
**ChargedAmount** | **string** |  | 
**PricePerTb** | **NullableString** |  | 
**Currency** | **string** |  | 
**UnitBytes** | **int32** |  | 
**RemainingBytes** | **NullableInt32** |  | 
**LastSampleAt** | **NullableString** |  | 
**Daily** | **[]map[string]interface{}** |  | 

## Methods

### NewServerTrafficResponse

`func NewServerTrafficResponse(year int32, month int32, asOf NullableString, status string, bytesIn int32, bytesOut int32, bytesTotal int32, includedTb NullableInt32, usedUnits int32, usedTb string, billableBytes int32, billableUnits int32, billableTb string, billableFrom NullableString, chargedTb int32, chargedAmount string, pricePerTb NullableString, currency string, unitBytes int32, remainingBytes NullableInt32, lastSampleAt NullableString, daily []map[string]interface{}, ) *ServerTrafficResponse`

NewServerTrafficResponse instantiates a new ServerTrafficResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTrafficResponseWithDefaults

`func NewServerTrafficResponseWithDefaults() *ServerTrafficResponse`

NewServerTrafficResponseWithDefaults instantiates a new ServerTrafficResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *ServerTrafficResponse) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ServerTrafficResponse) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ServerTrafficResponse) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *ServerTrafficResponse) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ServerTrafficResponse) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ServerTrafficResponse) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetAsOf

`func (o *ServerTrafficResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ServerTrafficResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ServerTrafficResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.


### SetAsOfNil

`func (o *ServerTrafficResponse) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *ServerTrafficResponse) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetStatus

`func (o *ServerTrafficResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerTrafficResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerTrafficResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBytesIn

`func (o *ServerTrafficResponse) GetBytesIn() int32`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *ServerTrafficResponse) GetBytesInOk() (*int32, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *ServerTrafficResponse) SetBytesIn(v int32)`

SetBytesIn sets BytesIn field to given value.


### GetBytesOut

`func (o *ServerTrafficResponse) GetBytesOut() int32`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *ServerTrafficResponse) GetBytesOutOk() (*int32, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *ServerTrafficResponse) SetBytesOut(v int32)`

SetBytesOut sets BytesOut field to given value.


### GetBytesTotal

`func (o *ServerTrafficResponse) GetBytesTotal() int32`

GetBytesTotal returns the BytesTotal field if non-nil, zero value otherwise.

### GetBytesTotalOk

`func (o *ServerTrafficResponse) GetBytesTotalOk() (*int32, bool)`

GetBytesTotalOk returns a tuple with the BytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTotal

`func (o *ServerTrafficResponse) SetBytesTotal(v int32)`

SetBytesTotal sets BytesTotal field to given value.


### GetIncludedTb

`func (o *ServerTrafficResponse) GetIncludedTb() int32`

GetIncludedTb returns the IncludedTb field if non-nil, zero value otherwise.

### GetIncludedTbOk

`func (o *ServerTrafficResponse) GetIncludedTbOk() (*int32, bool)`

GetIncludedTbOk returns a tuple with the IncludedTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTb

`func (o *ServerTrafficResponse) SetIncludedTb(v int32)`

SetIncludedTb sets IncludedTb field to given value.


### SetIncludedTbNil

`func (o *ServerTrafficResponse) SetIncludedTbNil(b bool)`

 SetIncludedTbNil sets the value for IncludedTb to be an explicit nil

### UnsetIncludedTb
`func (o *ServerTrafficResponse) UnsetIncludedTb()`

UnsetIncludedTb ensures that no value is present for IncludedTb, not even an explicit nil
### GetUsedUnits

`func (o *ServerTrafficResponse) GetUsedUnits() int32`

GetUsedUnits returns the UsedUnits field if non-nil, zero value otherwise.

### GetUsedUnitsOk

`func (o *ServerTrafficResponse) GetUsedUnitsOk() (*int32, bool)`

GetUsedUnitsOk returns a tuple with the UsedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedUnits

`func (o *ServerTrafficResponse) SetUsedUnits(v int32)`

SetUsedUnits sets UsedUnits field to given value.


### GetUsedTb

`func (o *ServerTrafficResponse) GetUsedTb() string`

GetUsedTb returns the UsedTb field if non-nil, zero value otherwise.

### GetUsedTbOk

`func (o *ServerTrafficResponse) GetUsedTbOk() (*string, bool)`

GetUsedTbOk returns a tuple with the UsedTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTb

`func (o *ServerTrafficResponse) SetUsedTb(v string)`

SetUsedTb sets UsedTb field to given value.


### GetBillableBytes

`func (o *ServerTrafficResponse) GetBillableBytes() int32`

GetBillableBytes returns the BillableBytes field if non-nil, zero value otherwise.

### GetBillableBytesOk

`func (o *ServerTrafficResponse) GetBillableBytesOk() (*int32, bool)`

GetBillableBytesOk returns a tuple with the BillableBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableBytes

`func (o *ServerTrafficResponse) SetBillableBytes(v int32)`

SetBillableBytes sets BillableBytes field to given value.


### GetBillableUnits

`func (o *ServerTrafficResponse) GetBillableUnits() int32`

GetBillableUnits returns the BillableUnits field if non-nil, zero value otherwise.

### GetBillableUnitsOk

`func (o *ServerTrafficResponse) GetBillableUnitsOk() (*int32, bool)`

GetBillableUnitsOk returns a tuple with the BillableUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableUnits

`func (o *ServerTrafficResponse) SetBillableUnits(v int32)`

SetBillableUnits sets BillableUnits field to given value.


### GetBillableTb

`func (o *ServerTrafficResponse) GetBillableTb() string`

GetBillableTb returns the BillableTb field if non-nil, zero value otherwise.

### GetBillableTbOk

`func (o *ServerTrafficResponse) GetBillableTbOk() (*string, bool)`

GetBillableTbOk returns a tuple with the BillableTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTb

`func (o *ServerTrafficResponse) SetBillableTb(v string)`

SetBillableTb sets BillableTb field to given value.


### GetBillableFrom

`func (o *ServerTrafficResponse) GetBillableFrom() string`

GetBillableFrom returns the BillableFrom field if non-nil, zero value otherwise.

### GetBillableFromOk

`func (o *ServerTrafficResponse) GetBillableFromOk() (*string, bool)`

GetBillableFromOk returns a tuple with the BillableFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableFrom

`func (o *ServerTrafficResponse) SetBillableFrom(v string)`

SetBillableFrom sets BillableFrom field to given value.


### SetBillableFromNil

`func (o *ServerTrafficResponse) SetBillableFromNil(b bool)`

 SetBillableFromNil sets the value for BillableFrom to be an explicit nil

### UnsetBillableFrom
`func (o *ServerTrafficResponse) UnsetBillableFrom()`

UnsetBillableFrom ensures that no value is present for BillableFrom, not even an explicit nil
### GetChargedTb

`func (o *ServerTrafficResponse) GetChargedTb() int32`

GetChargedTb returns the ChargedTb field if non-nil, zero value otherwise.

### GetChargedTbOk

`func (o *ServerTrafficResponse) GetChargedTbOk() (*int32, bool)`

GetChargedTbOk returns a tuple with the ChargedTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedTb

`func (o *ServerTrafficResponse) SetChargedTb(v int32)`

SetChargedTb sets ChargedTb field to given value.


### GetChargedAmount

`func (o *ServerTrafficResponse) GetChargedAmount() string`

GetChargedAmount returns the ChargedAmount field if non-nil, zero value otherwise.

### GetChargedAmountOk

`func (o *ServerTrafficResponse) GetChargedAmountOk() (*string, bool)`

GetChargedAmountOk returns a tuple with the ChargedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedAmount

`func (o *ServerTrafficResponse) SetChargedAmount(v string)`

SetChargedAmount sets ChargedAmount field to given value.


### GetPricePerTb

`func (o *ServerTrafficResponse) GetPricePerTb() string`

GetPricePerTb returns the PricePerTb field if non-nil, zero value otherwise.

### GetPricePerTbOk

`func (o *ServerTrafficResponse) GetPricePerTbOk() (*string, bool)`

GetPricePerTbOk returns a tuple with the PricePerTb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerTb

`func (o *ServerTrafficResponse) SetPricePerTb(v string)`

SetPricePerTb sets PricePerTb field to given value.


### SetPricePerTbNil

`func (o *ServerTrafficResponse) SetPricePerTbNil(b bool)`

 SetPricePerTbNil sets the value for PricePerTb to be an explicit nil

### UnsetPricePerTb
`func (o *ServerTrafficResponse) UnsetPricePerTb()`

UnsetPricePerTb ensures that no value is present for PricePerTb, not even an explicit nil
### GetCurrency

`func (o *ServerTrafficResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ServerTrafficResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ServerTrafficResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitBytes

`func (o *ServerTrafficResponse) GetUnitBytes() int32`

GetUnitBytes returns the UnitBytes field if non-nil, zero value otherwise.

### GetUnitBytesOk

`func (o *ServerTrafficResponse) GetUnitBytesOk() (*int32, bool)`

GetUnitBytesOk returns a tuple with the UnitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitBytes

`func (o *ServerTrafficResponse) SetUnitBytes(v int32)`

SetUnitBytes sets UnitBytes field to given value.


### GetRemainingBytes

`func (o *ServerTrafficResponse) GetRemainingBytes() int32`

GetRemainingBytes returns the RemainingBytes field if non-nil, zero value otherwise.

### GetRemainingBytesOk

`func (o *ServerTrafficResponse) GetRemainingBytesOk() (*int32, bool)`

GetRemainingBytesOk returns a tuple with the RemainingBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBytes

`func (o *ServerTrafficResponse) SetRemainingBytes(v int32)`

SetRemainingBytes sets RemainingBytes field to given value.


### SetRemainingBytesNil

`func (o *ServerTrafficResponse) SetRemainingBytesNil(b bool)`

 SetRemainingBytesNil sets the value for RemainingBytes to be an explicit nil

### UnsetRemainingBytes
`func (o *ServerTrafficResponse) UnsetRemainingBytes()`

UnsetRemainingBytes ensures that no value is present for RemainingBytes, not even an explicit nil
### GetLastSampleAt

`func (o *ServerTrafficResponse) GetLastSampleAt() string`

GetLastSampleAt returns the LastSampleAt field if non-nil, zero value otherwise.

### GetLastSampleAtOk

`func (o *ServerTrafficResponse) GetLastSampleAtOk() (*string, bool)`

GetLastSampleAtOk returns a tuple with the LastSampleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSampleAt

`func (o *ServerTrafficResponse) SetLastSampleAt(v string)`

SetLastSampleAt sets LastSampleAt field to given value.


### SetLastSampleAtNil

`func (o *ServerTrafficResponse) SetLastSampleAtNil(b bool)`

 SetLastSampleAtNil sets the value for LastSampleAt to be an explicit nil

### UnsetLastSampleAt
`func (o *ServerTrafficResponse) UnsetLastSampleAt()`

UnsetLastSampleAt ensures that no value is present for LastSampleAt, not even an explicit nil
### GetDaily

`func (o *ServerTrafficResponse) GetDaily() []map[string]interface{}`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *ServerTrafficResponse) GetDailyOk() (*[]map[string]interface{}, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *ServerTrafficResponse) SetDaily(v []map[string]interface{})`

SetDaily sets Daily field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


