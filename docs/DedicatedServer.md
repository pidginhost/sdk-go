# DedicatedServer

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
**ServerStatus** | **string** |  | [readonly] 
**Ips** | **string** |  | [readonly] 
**OsName** | **string** |  | [readonly] 

## Methods

### NewDedicatedServer

`func NewDedicatedServer(id int32, hostname string, status Status63aEnum, price float64, nextInvoice string, created time.Time, billingCycle string, serverStatus string, ips string, osName string, ) *DedicatedServer`

NewDedicatedServer instantiates a new DedicatedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerWithDefaults

`func NewDedicatedServerWithDefaults() *DedicatedServer`

NewDedicatedServerWithDefaults instantiates a new DedicatedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DedicatedServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DedicatedServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DedicatedServer) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *DedicatedServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DedicatedServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DedicatedServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *DedicatedServer) GetStatus() Status63aEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DedicatedServer) GetStatusOk() (*Status63aEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DedicatedServer) SetStatus(v Status63aEnum)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *DedicatedServer) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DedicatedServer) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DedicatedServer) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetNextInvoice

`func (o *DedicatedServer) GetNextInvoice() string`

GetNextInvoice returns the NextInvoice field if non-nil, zero value otherwise.

### GetNextInvoiceOk

`func (o *DedicatedServer) GetNextInvoiceOk() (*string, bool)`

GetNextInvoiceOk returns a tuple with the NextInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoice

`func (o *DedicatedServer) SetNextInvoice(v string)`

SetNextInvoice sets NextInvoice field to given value.


### GetCreated

`func (o *DedicatedServer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DedicatedServer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DedicatedServer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetBillingCycle

`func (o *DedicatedServer) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DedicatedServer) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DedicatedServer) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetServerStatus

`func (o *DedicatedServer) GetServerStatus() string`

GetServerStatus returns the ServerStatus field if non-nil, zero value otherwise.

### GetServerStatusOk

`func (o *DedicatedServer) GetServerStatusOk() (*string, bool)`

GetServerStatusOk returns a tuple with the ServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatus

`func (o *DedicatedServer) SetServerStatus(v string)`

SetServerStatus sets ServerStatus field to given value.


### GetIps

`func (o *DedicatedServer) GetIps() string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *DedicatedServer) GetIpsOk() (*string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *DedicatedServer) SetIps(v string)`

SetIps sets Ips field to given value.


### GetOsName

`func (o *DedicatedServer) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *DedicatedServer) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *DedicatedServer) SetOsName(v string)`

SetOsName sets OsName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


