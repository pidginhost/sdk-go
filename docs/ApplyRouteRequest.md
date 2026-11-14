# ApplyRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | **[]int32** | The servers to re-render and reboot; each one listed reboots. | 

## Methods

### NewApplyRouteRequest

`func NewApplyRouteRequest(servers []int32, ) *ApplyRouteRequest`

NewApplyRouteRequest instantiates a new ApplyRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyRouteRequestWithDefaults

`func NewApplyRouteRequestWithDefaults() *ApplyRouteRequest`

NewApplyRouteRequestWithDefaults instantiates a new ApplyRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *ApplyRouteRequest) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ApplyRouteRequest) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ApplyRouteRequest) SetServers(v []int32)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


