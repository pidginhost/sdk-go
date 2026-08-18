# AttachIPv6Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv6** | **string** | ID or address of an IPv6 you own. | 
**Reboot** | Pointer to **bool** | Restart the server so the guest OS picks up the address. | [optional] [default to false]

## Methods

### NewAttachIPv6Request

`func NewAttachIPv6Request(ipv6 string, ) *AttachIPv6Request`

NewAttachIPv6Request instantiates a new AttachIPv6Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachIPv6RequestWithDefaults

`func NewAttachIPv6RequestWithDefaults() *AttachIPv6Request`

NewAttachIPv6RequestWithDefaults instantiates a new AttachIPv6Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv6

`func (o *AttachIPv6Request) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *AttachIPv6Request) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *AttachIPv6Request) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.


### GetReboot

`func (o *AttachIPv6Request) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *AttachIPv6Request) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *AttachIPv6Request) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *AttachIPv6Request) HasReboot() bool`

HasReboot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


