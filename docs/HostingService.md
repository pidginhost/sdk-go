# HostingService

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
**PackageName** | **string** |  | [readonly] 
**NodeUrl** | **string** |  | [readonly] 
**Username** | **string** |  | [readonly] 

## Methods

### NewHostingService

`func NewHostingService(id int32, hostname string, status Status63aEnum, price float64, nextInvoice string, created time.Time, billingCycle string, packageName string, nodeUrl string, username string, ) *HostingService`

NewHostingService instantiates a new HostingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingServiceWithDefaults

`func NewHostingServiceWithDefaults() *HostingService`

NewHostingServiceWithDefaults instantiates a new HostingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostingService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostingService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostingService) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *HostingService) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostingService) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HostingService) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *HostingService) GetStatus() Status63aEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostingService) GetStatusOk() (*Status63aEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostingService) SetStatus(v Status63aEnum)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *HostingService) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *HostingService) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *HostingService) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetNextInvoice

`func (o *HostingService) GetNextInvoice() string`

GetNextInvoice returns the NextInvoice field if non-nil, zero value otherwise.

### GetNextInvoiceOk

`func (o *HostingService) GetNextInvoiceOk() (*string, bool)`

GetNextInvoiceOk returns a tuple with the NextInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoice

`func (o *HostingService) SetNextInvoice(v string)`

SetNextInvoice sets NextInvoice field to given value.


### GetCreated

`func (o *HostingService) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HostingService) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HostingService) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetBillingCycle

`func (o *HostingService) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *HostingService) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *HostingService) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetPackageName

`func (o *HostingService) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *HostingService) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *HostingService) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetNodeUrl

`func (o *HostingService) GetNodeUrl() string`

GetNodeUrl returns the NodeUrl field if non-nil, zero value otherwise.

### GetNodeUrlOk

`func (o *HostingService) GetNodeUrlOk() (*string, bool)`

GetNodeUrlOk returns a tuple with the NodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUrl

`func (o *HostingService) SetNodeUrl(v string)`

SetNodeUrl sets NodeUrl field to given value.


### GetUsername

`func (o *HostingService) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HostingService) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HostingService) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


