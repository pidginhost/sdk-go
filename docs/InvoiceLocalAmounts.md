# InvoiceLocalAmounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**Subtotal** | **string** |  | 
**VatValue** | **string** |  | 
**Total** | **string** |  | 
**Primary** | **bool** | True for a Romanian beneficiary (the rule that picks the PDF language); show RON first. | 

## Methods

### NewInvoiceLocalAmounts

`func NewInvoiceLocalAmounts(currency CurrencyEnum, subtotal string, vatValue string, total string, primary bool, ) *InvoiceLocalAmounts`

NewInvoiceLocalAmounts instantiates a new InvoiceLocalAmounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLocalAmountsWithDefaults

`func NewInvoiceLocalAmountsWithDefaults() *InvoiceLocalAmounts`

NewInvoiceLocalAmountsWithDefaults instantiates a new InvoiceLocalAmounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InvoiceLocalAmounts) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLocalAmounts) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLocalAmounts) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetSubtotal

`func (o *InvoiceLocalAmounts) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceLocalAmounts) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceLocalAmounts) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetVatValue

`func (o *InvoiceLocalAmounts) GetVatValue() string`

GetVatValue returns the VatValue field if non-nil, zero value otherwise.

### GetVatValueOk

`func (o *InvoiceLocalAmounts) GetVatValueOk() (*string, bool)`

GetVatValueOk returns a tuple with the VatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatValue

`func (o *InvoiceLocalAmounts) SetVatValue(v string)`

SetVatValue sets VatValue field to given value.


### GetTotal

`func (o *InvoiceLocalAmounts) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceLocalAmounts) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceLocalAmounts) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetPrimary

`func (o *InvoiceLocalAmounts) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *InvoiceLocalAmounts) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *InvoiceLocalAmounts) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


