# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Billing** | [**NullableDashboardBilling**](DashboardBilling.md) |  | 
**Resources** | [**NullableDashboardResources**](DashboardResources.md) |  | 

## Methods

### NewDashboard

`func NewDashboard(currency string, billing NullableDashboardBilling, resources NullableDashboardResources, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Dashboard) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Dashboard) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Dashboard) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBilling

`func (o *Dashboard) GetBilling() DashboardBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Dashboard) GetBillingOk() (*DashboardBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Dashboard) SetBilling(v DashboardBilling)`

SetBilling sets Billing field to given value.


### SetBillingNil

`func (o *Dashboard) SetBillingNil(b bool)`

 SetBillingNil sets the value for Billing to be an explicit nil

### UnsetBilling
`func (o *Dashboard) UnsetBilling()`

UnsetBilling ensures that no value is present for Billing, not even an explicit nil
### GetResources

`func (o *Dashboard) GetResources() DashboardResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Dashboard) GetResourcesOk() (*DashboardResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Dashboard) SetResources(v DashboardResources)`

SetResources sets Resources field to given value.


### SetResourcesNil

`func (o *Dashboard) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *Dashboard) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


