# InvoiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**NumberProforma** | **string** |  | [readonly] 
**NumberFiscal** | **string** |  | [readonly] 
**Status** | [**Status03cEnum**](Status03cEnum.md) |  | [readonly] 
**Subtotal** | **float64** |  | [readonly] 
**VatValue** | **float64** |  | [readonly] 
**VatPercentage** | **int32** |  | [readonly] 
**Total** | **float64** |  | [readonly] 
**InvoiceDate** | **string** |  | [readonly] 
**DueDate** | **NullableString** |  | [readonly] 
**PaymentDate** | **NullableTime** |  | [readonly] 

## Methods

### NewInvoiceList

`func NewInvoiceList(id int32, numberProforma string, numberFiscal string, status Status03cEnum, subtotal float64, vatValue float64, vatPercentage int32, total float64, invoiceDate string, dueDate NullableString, paymentDate NullableTime, ) *InvoiceList`

NewInvoiceList instantiates a new InvoiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceListWithDefaults

`func NewInvoiceListWithDefaults() *InvoiceList`

NewInvoiceListWithDefaults instantiates a new InvoiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceList) SetId(v int32)`

SetId sets Id field to given value.


### GetNumberProforma

`func (o *InvoiceList) GetNumberProforma() string`

GetNumberProforma returns the NumberProforma field if non-nil, zero value otherwise.

### GetNumberProformaOk

`func (o *InvoiceList) GetNumberProformaOk() (*string, bool)`

GetNumberProformaOk returns a tuple with the NumberProforma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberProforma

`func (o *InvoiceList) SetNumberProforma(v string)`

SetNumberProforma sets NumberProforma field to given value.


### GetNumberFiscal

`func (o *InvoiceList) GetNumberFiscal() string`

GetNumberFiscal returns the NumberFiscal field if non-nil, zero value otherwise.

### GetNumberFiscalOk

`func (o *InvoiceList) GetNumberFiscalOk() (*string, bool)`

GetNumberFiscalOk returns a tuple with the NumberFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFiscal

`func (o *InvoiceList) SetNumberFiscal(v string)`

SetNumberFiscal sets NumberFiscal field to given value.


### GetStatus

`func (o *InvoiceList) GetStatus() Status03cEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceList) GetStatusOk() (*Status03cEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceList) SetStatus(v Status03cEnum)`

SetStatus sets Status field to given value.


### GetSubtotal

`func (o *InvoiceList) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceList) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceList) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetVatValue

`func (o *InvoiceList) GetVatValue() float64`

GetVatValue returns the VatValue field if non-nil, zero value otherwise.

### GetVatValueOk

`func (o *InvoiceList) GetVatValueOk() (*float64, bool)`

GetVatValueOk returns a tuple with the VatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatValue

`func (o *InvoiceList) SetVatValue(v float64)`

SetVatValue sets VatValue field to given value.


### GetVatPercentage

`func (o *InvoiceList) GetVatPercentage() int32`

GetVatPercentage returns the VatPercentage field if non-nil, zero value otherwise.

### GetVatPercentageOk

`func (o *InvoiceList) GetVatPercentageOk() (*int32, bool)`

GetVatPercentageOk returns a tuple with the VatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercentage

`func (o *InvoiceList) SetVatPercentage(v int32)`

SetVatPercentage sets VatPercentage field to given value.


### GetTotal

`func (o *InvoiceList) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceList) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceList) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetInvoiceDate

`func (o *InvoiceList) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoiceList) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoiceList) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetDueDate

`func (o *InvoiceList) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceList) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceList) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### SetDueDateNil

`func (o *InvoiceList) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceList) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaymentDate

`func (o *InvoiceList) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceList) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceList) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.


### SetPaymentDateNil

`func (o *InvoiceList) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *InvoiceList) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


