# ResourceBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** |  | 
**Currency** | **string** |  | 
**TaxIncluded** | **bool** |  | 
**CycleMonths** | **int32** |  | 
**CycleLabel** | **string** |  | 
**NextInvoice** | **NullableString** |  | 
**Profile** | [**NullableResourceBillingProfile**](ResourceBillingProfile.md) |  | 
**Metered** | **bool** |  | 
**Included** | **bool** |  | 
**HourlyAmount** | Pointer to **string** |  | [optional] 
**MonthlyTotal** | Pointer to **string** |  | [optional] 
**Components** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewResourceBilling

`func NewResourceBilling(amount string, currency string, taxIncluded bool, cycleMonths int32, cycleLabel string, nextInvoice NullableString, profile NullableResourceBillingProfile, metered bool, included bool, ) *ResourceBilling`

NewResourceBilling instantiates a new ResourceBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceBillingWithDefaults

`func NewResourceBillingWithDefaults() *ResourceBilling`

NewResourceBillingWithDefaults instantiates a new ResourceBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResourceBilling) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResourceBilling) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResourceBilling) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *ResourceBilling) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResourceBilling) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResourceBilling) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTaxIncluded

`func (o *ResourceBilling) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *ResourceBilling) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *ResourceBilling) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.


### GetCycleMonths

`func (o *ResourceBilling) GetCycleMonths() int32`

GetCycleMonths returns the CycleMonths field if non-nil, zero value otherwise.

### GetCycleMonthsOk

`func (o *ResourceBilling) GetCycleMonthsOk() (*int32, bool)`

GetCycleMonthsOk returns a tuple with the CycleMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleMonths

`func (o *ResourceBilling) SetCycleMonths(v int32)`

SetCycleMonths sets CycleMonths field to given value.


### GetCycleLabel

`func (o *ResourceBilling) GetCycleLabel() string`

GetCycleLabel returns the CycleLabel field if non-nil, zero value otherwise.

### GetCycleLabelOk

`func (o *ResourceBilling) GetCycleLabelOk() (*string, bool)`

GetCycleLabelOk returns a tuple with the CycleLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLabel

`func (o *ResourceBilling) SetCycleLabel(v string)`

SetCycleLabel sets CycleLabel field to given value.


### GetNextInvoice

`func (o *ResourceBilling) GetNextInvoice() string`

GetNextInvoice returns the NextInvoice field if non-nil, zero value otherwise.

### GetNextInvoiceOk

`func (o *ResourceBilling) GetNextInvoiceOk() (*string, bool)`

GetNextInvoiceOk returns a tuple with the NextInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoice

`func (o *ResourceBilling) SetNextInvoice(v string)`

SetNextInvoice sets NextInvoice field to given value.


### SetNextInvoiceNil

`func (o *ResourceBilling) SetNextInvoiceNil(b bool)`

 SetNextInvoiceNil sets the value for NextInvoice to be an explicit nil

### UnsetNextInvoice
`func (o *ResourceBilling) UnsetNextInvoice()`

UnsetNextInvoice ensures that no value is present for NextInvoice, not even an explicit nil
### GetProfile

`func (o *ResourceBilling) GetProfile() ResourceBillingProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResourceBilling) GetProfileOk() (*ResourceBillingProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResourceBilling) SetProfile(v ResourceBillingProfile)`

SetProfile sets Profile field to given value.


### SetProfileNil

`func (o *ResourceBilling) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ResourceBilling) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetMetered

`func (o *ResourceBilling) GetMetered() bool`

GetMetered returns the Metered field if non-nil, zero value otherwise.

### GetMeteredOk

`func (o *ResourceBilling) GetMeteredOk() (*bool, bool)`

GetMeteredOk returns a tuple with the Metered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetered

`func (o *ResourceBilling) SetMetered(v bool)`

SetMetered sets Metered field to given value.


### GetIncluded

`func (o *ResourceBilling) GetIncluded() bool`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ResourceBilling) GetIncludedOk() (*bool, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ResourceBilling) SetIncluded(v bool)`

SetIncluded sets Included field to given value.


### GetHourlyAmount

`func (o *ResourceBilling) GetHourlyAmount() string`

GetHourlyAmount returns the HourlyAmount field if non-nil, zero value otherwise.

### GetHourlyAmountOk

`func (o *ResourceBilling) GetHourlyAmountOk() (*string, bool)`

GetHourlyAmountOk returns a tuple with the HourlyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyAmount

`func (o *ResourceBilling) SetHourlyAmount(v string)`

SetHourlyAmount sets HourlyAmount field to given value.

### HasHourlyAmount

`func (o *ResourceBilling) HasHourlyAmount() bool`

HasHourlyAmount returns a boolean if a field has been set.

### GetMonthlyTotal

`func (o *ResourceBilling) GetMonthlyTotal() string`

GetMonthlyTotal returns the MonthlyTotal field if non-nil, zero value otherwise.

### GetMonthlyTotalOk

`func (o *ResourceBilling) GetMonthlyTotalOk() (*string, bool)`

GetMonthlyTotalOk returns a tuple with the MonthlyTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyTotal

`func (o *ResourceBilling) SetMonthlyTotal(v string)`

SetMonthlyTotal sets MonthlyTotal field to given value.

### HasMonthlyTotal

`func (o *ResourceBilling) HasMonthlyTotal() bool`

HasMonthlyTotal returns a boolean if a field has been set.

### GetComponents

`func (o *ResourceBilling) GetComponents() map[string]string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ResourceBilling) GetComponentsOk() (*map[string]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ResourceBilling) SetComponents(v map[string]string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ResourceBilling) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


