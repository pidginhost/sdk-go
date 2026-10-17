# DashboardBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** |  | 
**UnpaidCount** | **int32** |  | 
**UnpaidTotal** | **string** |  | 
**Invoices** | [**[]InvoiceList**](InvoiceList.md) |  | 

## Methods

### NewDashboardBilling

`func NewDashboardBilling(balance string, unpaidCount int32, unpaidTotal string, invoices []InvoiceList, ) *DashboardBilling`

NewDashboardBilling instantiates a new DashboardBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardBillingWithDefaults

`func NewDashboardBillingWithDefaults() *DashboardBilling`

NewDashboardBillingWithDefaults instantiates a new DashboardBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *DashboardBilling) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *DashboardBilling) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *DashboardBilling) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetUnpaidCount

`func (o *DashboardBilling) GetUnpaidCount() int32`

GetUnpaidCount returns the UnpaidCount field if non-nil, zero value otherwise.

### GetUnpaidCountOk

`func (o *DashboardBilling) GetUnpaidCountOk() (*int32, bool)`

GetUnpaidCountOk returns a tuple with the UnpaidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidCount

`func (o *DashboardBilling) SetUnpaidCount(v int32)`

SetUnpaidCount sets UnpaidCount field to given value.


### GetUnpaidTotal

`func (o *DashboardBilling) GetUnpaidTotal() string`

GetUnpaidTotal returns the UnpaidTotal field if non-nil, zero value otherwise.

### GetUnpaidTotalOk

`func (o *DashboardBilling) GetUnpaidTotalOk() (*string, bool)`

GetUnpaidTotalOk returns a tuple with the UnpaidTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidTotal

`func (o *DashboardBilling) SetUnpaidTotal(v string)`

SetUnpaidTotal sets UnpaidTotal field to given value.


### GetInvoices

`func (o *DashboardBilling) GetInvoices() []InvoiceList`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *DashboardBilling) GetInvoicesOk() (*[]InvoiceList, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *DashboardBilling) SetInvoices(v []InvoiceList)`

SetInvoices sets Invoices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


