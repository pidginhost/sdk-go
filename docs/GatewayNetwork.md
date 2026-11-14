# GatewayNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **int32** |  | [readonly] 
**Role** | [**GatewayNetworkRoleEnum**](GatewayNetworkRoleEnum.md) |  | [readonly] 
**Snat** | **bool** |  | [readonly] 
**State** | [**GatewayNetworkStateEnum**](GatewayNetworkStateEnum.md) |  | [readonly] 
**Slot** | **int32** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 
**Reason** | **string** | Why the last network change failed, or empty. | [readonly] 

## Methods

### NewGatewayNetwork

`func NewGatewayNetwork(network int32, role GatewayNetworkRoleEnum, snat bool, state GatewayNetworkStateEnum, slot int32, address string, changedRevision int32, pending bool, reason string, ) *GatewayNetwork`

NewGatewayNetwork instantiates a new GatewayNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayNetworkWithDefaults

`func NewGatewayNetworkWithDefaults() *GatewayNetwork`

NewGatewayNetworkWithDefaults instantiates a new GatewayNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GatewayNetwork) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayNetwork) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayNetwork) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetRole

`func (o *GatewayNetwork) GetRole() GatewayNetworkRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GatewayNetwork) GetRoleOk() (*GatewayNetworkRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GatewayNetwork) SetRole(v GatewayNetworkRoleEnum)`

SetRole sets Role field to given value.


### GetSnat

`func (o *GatewayNetwork) GetSnat() bool`

GetSnat returns the Snat field if non-nil, zero value otherwise.

### GetSnatOk

`func (o *GatewayNetwork) GetSnatOk() (*bool, bool)`

GetSnatOk returns a tuple with the Snat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnat

`func (o *GatewayNetwork) SetSnat(v bool)`

SetSnat sets Snat field to given value.


### GetState

`func (o *GatewayNetwork) GetState() GatewayNetworkStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GatewayNetwork) GetStateOk() (*GatewayNetworkStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GatewayNetwork) SetState(v GatewayNetworkStateEnum)`

SetState sets State field to given value.


### GetSlot

`func (o *GatewayNetwork) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *GatewayNetwork) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *GatewayNetwork) SetSlot(v int32)`

SetSlot sets Slot field to given value.


### GetAddress

`func (o *GatewayNetwork) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GatewayNetwork) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GatewayNetwork) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChangedRevision

`func (o *GatewayNetwork) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayNetwork) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayNetwork) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayNetwork) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayNetwork) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayNetwork) SetPending(v bool)`

SetPending sets Pending field to given value.


### GetReason

`func (o *GatewayNetwork) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GatewayNetwork) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GatewayNetwork) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


