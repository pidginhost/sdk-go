# ConnectedVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Hostname** | **string** |  | 
**Ip** | **string** |  | 

## Methods

### NewConnectedVM

`func NewConnectedVM(id int32, hostname string, ip string, ) *ConnectedVM`

NewConnectedVM instantiates a new ConnectedVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedVMWithDefaults

`func NewConnectedVMWithDefaults() *ConnectedVM`

NewConnectedVMWithDefaults instantiates a new ConnectedVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectedVM) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectedVM) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectedVM) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *ConnectedVM) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ConnectedVM) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ConnectedVM) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIp

`func (o *ConnectedVM) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ConnectedVM) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ConnectedVM) SetIp(v string)`

SetIp sets Ip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


