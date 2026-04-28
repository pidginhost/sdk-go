# ConnectedVMsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vms** | [**[]ConnectedVM**](ConnectedVM.md) |  | 

## Methods

### NewConnectedVMsResponse

`func NewConnectedVMsResponse(vms []ConnectedVM, ) *ConnectedVMsResponse`

NewConnectedVMsResponse instantiates a new ConnectedVMsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedVMsResponseWithDefaults

`func NewConnectedVMsResponseWithDefaults() *ConnectedVMsResponse`

NewConnectedVMsResponseWithDefaults instantiates a new ConnectedVMsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVms

`func (o *ConnectedVMsResponse) GetVms() []ConnectedVM`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *ConnectedVMsResponse) GetVmsOk() (*[]ConnectedVM, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *ConnectedVMsResponse) SetVms(v []ConnectedVM)`

SetVms sets Vms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


