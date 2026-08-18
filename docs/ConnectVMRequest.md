# ConnectVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **int32** | VM resource PK or hostname | 

## Methods

### NewConnectVMRequest

`func NewConnectVMRequest(server int32, ) *ConnectVMRequest`

NewConnectVMRequest instantiates a new ConnectVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectVMRequestWithDefaults

`func NewConnectVMRequestWithDefaults() *ConnectVMRequest`

NewConnectVMRequestWithDefaults instantiates a new ConnectVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *ConnectVMRequest) GetServer() int32`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConnectVMRequest) GetServerOk() (*int32, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConnectVMRequest) SetServer(v int32)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


