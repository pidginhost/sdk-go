# EmailStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **string** |  | 
**End** | **string** |  | 
**Totals** | [**StatsTotals**](StatsTotals.md) |  | 
**Days** | [**[]StatsDay**](StatsDay.md) |  | 
**Reputation** | [**EmailReputation**](EmailReputation.md) |  | 

## Methods

### NewEmailStats

`func NewEmailStats(start string, end string, totals StatsTotals, days []StatsDay, reputation EmailReputation, ) *EmailStats`

NewEmailStats instantiates a new EmailStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatsWithDefaults

`func NewEmailStatsWithDefaults() *EmailStats`

NewEmailStatsWithDefaults instantiates a new EmailStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *EmailStats) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EmailStats) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EmailStats) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *EmailStats) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EmailStats) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EmailStats) SetEnd(v string)`

SetEnd sets End field to given value.


### GetTotals

`func (o *EmailStats) GetTotals() StatsTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *EmailStats) GetTotalsOk() (*StatsTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *EmailStats) SetTotals(v StatsTotals)`

SetTotals sets Totals field to given value.


### GetDays

`func (o *EmailStats) GetDays() []StatsDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EmailStats) GetDaysOk() (*[]StatsDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EmailStats) SetDays(v []StatsDay)`

SetDays sets Days field to given value.


### GetReputation

`func (o *EmailStats) GetReputation() EmailReputation`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *EmailStats) GetReputationOk() (*EmailReputation, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *EmailStats) SetReputation(v EmailReputation)`

SetReputation sets Reputation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


