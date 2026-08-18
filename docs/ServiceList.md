# ServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Hostname** | **string** |  | [readonly] 
**Status** | [**ServiceStatusEnum**](ServiceStatusEnum.md) |  | [readonly] 
**Price** | **string** | Euro without TVA | [readonly] 
**NextInvoice** | **string** |  | [readonly] 
**Created** | **string** |  | [readonly] 
**BillingCycle** | **string** |  | [readonly] 
**AutoPayment** | **bool** |  | [readonly] 
**Company** | [**NullableServiceCompany**](ServiceCompany.md) |  | [readonly] 

## Methods

### NewServiceList

`func NewServiceList(id int32, hostname string, status ServiceStatusEnum, price string, nextInvoice string, created string, billingCycle string, autoPayment bool, company NullableServiceCompany, ) *ServiceList`

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

`func (o *ServiceList) GetStatus() ServiceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceList) GetStatusOk() (*ServiceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceList) SetStatus(v ServiceStatusEnum)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *ServiceList) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ServiceList) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ServiceList) SetPrice(v string)`

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

`func (o *ServiceList) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceList) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceList) SetCreated(v string)`

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

`func (o *ServiceList) GetAutoPayment() bool`

GetAutoPayment returns the AutoPayment field if non-nil, zero value otherwise.

### GetAutoPaymentOk

`func (o *ServiceList) GetAutoPaymentOk() (*bool, bool)`

GetAutoPaymentOk returns a tuple with the AutoPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayment

`func (o *ServiceList) SetAutoPayment(v bool)`

SetAutoPayment sets AutoPayment field to given value.


### GetCompany

`func (o *ServiceList) GetCompany() ServiceCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ServiceList) GetCompanyOk() (*ServiceCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ServiceList) SetCompany(v ServiceCompany)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *ServiceList) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ServiceList) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


