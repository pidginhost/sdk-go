# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Status** | [**SubscriptionStatusEnum**](SubscriptionStatusEnum.md) |  | [readonly] 
**Service** | **int32** |  | [readonly] 
**ServiceHostname** | **string** |  | [readonly] 
**Subtotal** | **float64** |  | [readonly] 
**VatValue** | **float64** |  | [readonly] 
**Total** | **float64** |  | [readonly] 
**CreationDate** | **time.Time** |  | [readonly] 

## Methods

### NewSubscription

`func NewSubscription(id int32, status SubscriptionStatusEnum, service int32, serviceHostname string, subtotal float64, vatValue float64, total float64, creationDate time.Time, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *Subscription) GetStatus() SubscriptionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*SubscriptionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v SubscriptionStatusEnum)`

SetStatus sets Status field to given value.


### GetService

`func (o *Subscription) GetService() int32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Subscription) GetServiceOk() (*int32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Subscription) SetService(v int32)`

SetService sets Service field to given value.


### GetServiceHostname

`func (o *Subscription) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *Subscription) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *Subscription) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.


### GetSubtotal

`func (o *Subscription) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *Subscription) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *Subscription) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetVatValue

`func (o *Subscription) GetVatValue() float64`

GetVatValue returns the VatValue field if non-nil, zero value otherwise.

### GetVatValueOk

`func (o *Subscription) GetVatValueOk() (*float64, bool)`

GetVatValueOk returns a tuple with the VatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatValue

`func (o *Subscription) SetVatValue(v float64)`

SetVatValue sets VatValue field to given value.


### GetTotal

`func (o *Subscription) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Subscription) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Subscription) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetCreationDate

`func (o *Subscription) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Subscription) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Subscription) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


