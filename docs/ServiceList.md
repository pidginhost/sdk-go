# ServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Hostname** | **string** |  | [readonly] 
**Status** | [**Status63aEnum**](Status63aEnum.md) |  | [readonly] 
**Price** | **float64** | Euro without TVA | [readonly] 
**NextInvoice** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**BillingCycle** | **string** |  | [readonly] 
**AutoPayment** | **string** |  | [readonly] 
**Company** | **string** |  | [readonly] 

## Methods

### NewServiceList

`func NewServiceList(id int32, hostname string, status Status63aEnum, price float64, nextInvoice string, created time.Time, billingCycle string, autoPayment string, company string, ) *ServiceList`

NewServiceList instantiates a new ServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListWithDefaults

`func NewServiceListWithDefaults() *ServiceList`

NewServiceListWithDefaults instantiates a new ServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceList) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *ServiceList) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServiceList) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServiceList) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *ServiceList) GetStatus() Status63aEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceList) GetStatusOk() (*Status63aEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceList) SetStatus(v Status63aEnum)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *ServiceList) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ServiceList) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ServiceList) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetNextInvoice

`func (o *ServiceList) GetNextInvoice() string`

GetNextInvoice returns the NextInvoice field if non-nil, zero value otherwise.

### GetNextInvoiceOk

`func (o *ServiceList) GetNextInvoiceOk() (*string, bool)`

GetNextInvoiceOk returns a tuple with the NextInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoice

`func (o *ServiceList) SetNextInvoice(v string)`

SetNextInvoice sets NextInvoice field to given value.


### GetCreated

`func (o *ServiceList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetBillingCycle

`func (o *ServiceList) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *ServiceList) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *ServiceList) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetAutoPayment

`func (o *ServiceList) GetAutoPayment() string`

GetAutoPayment returns the AutoPayment field if non-nil, zero value otherwise.

### GetAutoPaymentOk

`func (o *ServiceList) GetAutoPaymentOk() (*string, bool)`

GetAutoPaymentOk returns a tuple with the AutoPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayment

`func (o *ServiceList) SetAutoPayment(v string)`

SetAutoPayment sets AutoPayment field to given value.


### GetCompany

`func (o *ServiceList) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ServiceList) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ServiceList) SetCompany(v string)`

SetCompany sets Company field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


