# GatewayInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | **int32** |  | 
**Name** | **string** |  | 
**State** | **string** | up, down or unknown. | 
**Addresses** | **[]string** |  | 

## Methods

### NewGatewayInterface

`func NewGatewayInterface(slot int32, name string, state string, addresses []string, ) *GatewayInterface`

NewGatewayInterface instantiates a new GatewayInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInterfaceWithDefaults

`func NewGatewayInterfaceWithDefaults() *GatewayInterface`

NewGatewayInterfaceWithDefaults instantiates a new GatewayInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *GatewayInterface) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *GatewayInterface) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *GatewayInterface) SetSlot(v int32)`

SetSlot sets Slot field to given value.


### GetName

`func (o *GatewayInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayInterface) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *GatewayInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GatewayInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GatewayInterface) SetState(v string)`

SetState sets State field to given value.


### GetAddresses

`func (o *GatewayInterface) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GatewayInterface) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GatewayInterface) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


