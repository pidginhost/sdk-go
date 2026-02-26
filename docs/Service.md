# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Company** | Pointer to **NullableInt32** |  | [optional] 
**BillingCycle** | **int32** |  | [readonly] 
**Hostname** | **string** |  | [readonly] 
**Price** | **float64** | Euro without TVA | [readonly] 
**Status** | [**ServiceStatusEnum**](ServiceStatusEnum.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**Terminated** | **NullableString** |  | [readonly] 
**NextInvoice** | **string** |  | [readonly] 
**PrimaryService** | **NullableInt32** |  | [readonly] 

## Methods

### NewService

`func NewService(id int32, billingCycle int32, hostname string, price float64, status ServiceStatusEnum, created time.Time, modified time.Time, terminated NullableString, nextInvoice string, primaryService NullableInt32, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.


### GetCompany

`func (o *Service) GetCompany() int32`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Service) GetCompanyOk() (*int32, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Service) SetCompany(v int32)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Service) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Service) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Service) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetBillingCycle

`func (o *Service) GetBillingCycle() int32`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *Service) GetBillingCycleOk() (*int32, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *Service) SetBillingCycle(v int32)`

SetBillingCycle sets BillingCycle field to given value.


### GetHostname

`func (o *Service) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Service) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Service) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPrice

`func (o *Service) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Service) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Service) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetStatus

`func (o *Service) GetStatus() ServiceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*ServiceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v ServiceStatusEnum)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *Service) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Service) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Service) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Service) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Service) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Service) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetTerminated

`func (o *Service) GetTerminated() string`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *Service) GetTerminatedOk() (*string, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *Service) SetTerminated(v string)`

SetTerminated sets Terminated field to given value.


### SetTerminatedNil

`func (o *Service) SetTerminatedNil(b bool)`

 SetTerminatedNil sets the value for Terminated to be an explicit nil

### UnsetTerminated
`func (o *Service) UnsetTerminated()`

UnsetTerminated ensures that no value is present for Terminated, not even an explicit nil
### GetNextInvoice

`func (o *Service) GetNextInvoice() string`

GetNextInvoice returns the NextInvoice field if non-nil, zero value otherwise.

### GetNextInvoiceOk

`func (o *Service) GetNextInvoiceOk() (*string, bool)`

GetNextInvoiceOk returns a tuple with the NextInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoice

`func (o *Service) SetNextInvoice(v string)`

SetNextInvoice sets NextInvoice field to given value.


### GetPrimaryService

`func (o *Service) GetPrimaryService() int32`

GetPrimaryService returns the PrimaryService field if non-nil, zero value otherwise.

### GetPrimaryServiceOk

`func (o *Service) GetPrimaryServiceOk() (*int32, bool)`

GetPrimaryServiceOk returns a tuple with the PrimaryService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryService

`func (o *Service) SetPrimaryService(v int32)`

SetPrimaryService sets PrimaryService field to given value.


### SetPrimaryServiceNil

`func (o *Service) SetPrimaryServiceNil(b bool)`

 SetPrimaryServiceNil sets the value for PrimaryService to be an explicit nil

### UnsetPrimaryService
`func (o *Service) UnsetPrimaryService()`

UnsetPrimaryService ensures that no value is present for PrimaryService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


