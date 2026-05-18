# InvoiceDetail

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
**PaymentDate** | **NullableString** |  | [readonly] 
**ProductInfo** | **interface{}** |  | [readonly] 
**ClientInfo** | **interface{}** |  | [readonly] 
**InvoiceInfo** | **interface{}** |  | [readonly] 
**PaymentMethod** | **string** |  | [readonly] 
**Services** | **string** |  | [readonly] 

## Methods

### NewInvoiceDetail

`func NewInvoiceDetail(id int32, numberProforma string, numberFiscal string, status Status03cEnum, subtotal float64, vatValue float64, vatPercentage int32, total float64, invoiceDate string, dueDate NullableString, paymentDate NullableString, productInfo interface{}, clientInfo interface{}, invoiceInfo interface{}, paymentMethod string, services string, ) *InvoiceDetail`

NewInvoiceDetail instantiates a new InvoiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailWithDefaults

`func NewInvoiceDetailWithDefaults() *InvoiceDetail`

NewInvoiceDetailWithDefaults instantiates a new InvoiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetNumberProforma

`func (o *InvoiceDetail) GetNumberProforma() string`

GetNumberProforma returns the NumberProforma field if non-nil, zero value otherwise.

### GetNumberProformaOk

`func (o *InvoiceDetail) GetNumberProformaOk() (*string, bool)`

GetNumberProformaOk returns a tuple with the NumberProforma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberProforma

`func (o *InvoiceDetail) SetNumberProforma(v string)`

SetNumberProforma sets NumberProforma field to given value.


### GetNumberFiscal

`func (o *InvoiceDetail) GetNumberFiscal() string`

GetNumberFiscal returns the NumberFiscal field if non-nil, zero value otherwise.

### GetNumberFiscalOk

`func (o *InvoiceDetail) GetNumberFiscalOk() (*string, bool)`

GetNumberFiscalOk returns a tuple with the NumberFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFiscal

`func (o *InvoiceDetail) SetNumberFiscal(v string)`

SetNumberFiscal sets NumberFiscal field to given value.


### GetStatus

`func (o *InvoiceDetail) GetStatus() Status03cEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceDetail) GetStatusOk() (*Status03cEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceDetail) SetStatus(v Status03cEnum)`

SetStatus sets Status field to given value.


### GetSubtotal

`func (o *InvoiceDetail) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceDetail) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceDetail) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetVatValue

`func (o *InvoiceDetail) GetVatValue() float64`

GetVatValue returns the VatValue field if non-nil, zero value otherwise.

### GetVatValueOk

`func (o *InvoiceDetail) GetVatValueOk() (*float64, bool)`

GetVatValueOk returns a tuple with the VatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatValue

`func (o *InvoiceDetail) SetVatValue(v float64)`

SetVatValue sets VatValue field to given value.


### GetVatPercentage

`func (o *InvoiceDetail) GetVatPercentage() int32`

GetVatPercentage returns the VatPercentage field if non-nil, zero value otherwise.

### GetVatPercentageOk

`func (o *InvoiceDetail) GetVatPercentageOk() (*int32, bool)`

GetVatPercentageOk returns a tuple with the VatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercentage

`func (o *InvoiceDetail) SetVatPercentage(v int32)`

SetVatPercentage sets VatPercentage field to given value.


### GetTotal

`func (o *InvoiceDetail) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceDetail) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceDetail) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetInvoiceDate

`func (o *InvoiceDetail) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoiceDetail) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoiceDetail) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetDueDate

`func (o *InvoiceDetail) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceDetail) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceDetail) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### SetDueDateNil

`func (o *InvoiceDetail) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceDetail) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaymentDate

`func (o *InvoiceDetail) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceDetail) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceDetail) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### SetPaymentDateNil

`func (o *InvoiceDetail) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *InvoiceDetail) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetProductInfo

`func (o *InvoiceDetail) GetProductInfo() interface{}`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *InvoiceDetail) GetProductInfoOk() (*interface{}, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *InvoiceDetail) SetProductInfo(v interface{})`

SetProductInfo sets ProductInfo field to given value.


### SetProductInfoNil

`func (o *InvoiceDetail) SetProductInfoNil(b bool)`

 SetProductInfoNil sets the value for ProductInfo to be an explicit nil

### UnsetProductInfo
`func (o *InvoiceDetail) UnsetProductInfo()`

UnsetProductInfo ensures that no value is present for ProductInfo, not even an explicit nil
### GetClientInfo

`func (o *InvoiceDetail) GetClientInfo() interface{}`

GetClientInfo returns the ClientInfo field if non-nil, zero value otherwise.

### GetClientInfoOk

`func (o *InvoiceDetail) GetClientInfoOk() (*interface{}, bool)`

GetClientInfoOk returns a tuple with the ClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInfo

`func (o *InvoiceDetail) SetClientInfo(v interface{})`

SetClientInfo sets ClientInfo field to given value.


### SetClientInfoNil

`func (o *InvoiceDetail) SetClientInfoNil(b bool)`

 SetClientInfoNil sets the value for ClientInfo to be an explicit nil

### UnsetClientInfo
`func (o *InvoiceDetail) UnsetClientInfo()`

UnsetClientInfo ensures that no value is present for ClientInfo, not even an explicit nil
### GetInvoiceInfo

`func (o *InvoiceDetail) GetInvoiceInfo() interface{}`

GetInvoiceInfo returns the InvoiceInfo field if non-nil, zero value otherwise.

### GetInvoiceInfoOk

`func (o *InvoiceDetail) GetInvoiceInfoOk() (*interface{}, bool)`

GetInvoiceInfoOk returns a tuple with the InvoiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceInfo

`func (o *InvoiceDetail) SetInvoiceInfo(v interface{})`

SetInvoiceInfo sets InvoiceInfo field to given value.


### SetInvoiceInfoNil

`func (o *InvoiceDetail) SetInvoiceInfoNil(b bool)`

 SetInvoiceInfoNil sets the value for InvoiceInfo to be an explicit nil

### UnsetInvoiceInfo
`func (o *InvoiceDetail) UnsetInvoiceInfo()`

UnsetInvoiceInfo ensures that no value is present for InvoiceInfo, not even an explicit nil
### GetPaymentMethod

`func (o *InvoiceDetail) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *InvoiceDetail) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *InvoiceDetail) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetServices

`func (o *InvoiceDetail) GetServices() string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InvoiceDetail) GetServicesOk() (*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InvoiceDetail) SetServices(v string)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


