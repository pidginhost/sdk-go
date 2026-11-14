# GatewayBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Network** | **int32** | An attached private network&#39;s id. | 
**Address** | **string** |  | [readonly] 
**Port** | **int32** |  | [readonly] 
**Weight** | **int32** |  | [readonly] 
**Enabled** | **bool** |  | [readonly] 
**Health** | [**GatewayBackendHealthEnum**](GatewayBackendHealthEnum.md) |  | [readonly] 
**HealthChanged** | **NullableString** |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 

## Methods

### NewGatewayBackend

`func NewGatewayBackend(id int32, network int32, address string, port int32, weight int32, enabled bool, health GatewayBackendHealthEnum, healthChanged NullableString, changedRevision int32, pending bool, ) *GatewayBackend`

NewGatewayBackend instantiates a new GatewayBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayBackendWithDefaults

`func NewGatewayBackendWithDefaults() *GatewayBackend`

NewGatewayBackendWithDefaults instantiates a new GatewayBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayBackend) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayBackend) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayBackend) SetId(v int32)`

SetId sets Id field to given value.


### GetNetwork

`func (o *GatewayBackend) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayBackend) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayBackend) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *GatewayBackend) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GatewayBackend) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GatewayBackend) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPort

`func (o *GatewayBackend) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayBackend) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayBackend) SetPort(v int32)`

SetPort sets Port field to given value.


### GetWeight

`func (o *GatewayBackend) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GatewayBackend) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GatewayBackend) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetEnabled

`func (o *GatewayBackend) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayBackend) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayBackend) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHealth

`func (o *GatewayBackend) GetHealth() GatewayBackendHealthEnum`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GatewayBackend) GetHealthOk() (*GatewayBackendHealthEnum, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GatewayBackend) SetHealth(v GatewayBackendHealthEnum)`

SetHealth sets Health field to given value.


### GetHealthChanged

`func (o *GatewayBackend) GetHealthChanged() string`

GetHealthChanged returns the HealthChanged field if non-nil, zero value otherwise.

### GetHealthChangedOk

`func (o *GatewayBackend) GetHealthChangedOk() (*string, bool)`

GetHealthChangedOk returns a tuple with the HealthChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthChanged

`func (o *GatewayBackend) SetHealthChanged(v string)`

SetHealthChanged sets HealthChanged field to given value.


### SetHealthChangedNil

`func (o *GatewayBackend) SetHealthChangedNil(b bool)`

 SetHealthChangedNil sets the value for HealthChanged to be an explicit nil

### UnsetHealthChanged
`func (o *GatewayBackend) UnsetHealthChanged()`

UnsetHealthChanged ensures that no value is present for HealthChanged, not even an explicit nil
### GetChangedRevision

`func (o *GatewayBackend) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayBackend) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayBackend) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayBackend) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayBackend) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayBackend) SetPending(v bool)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


