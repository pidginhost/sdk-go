# AttachIPv4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | **string** | ID or address of an IPv4 you own. | 
**Reboot** | Pointer to **bool** | Restart the server so the guest OS picks up the address. | [optional] [default to false]

## Methods

### NewAttachIPv4Request

`func NewAttachIPv4Request(ipv4 string, ) *AttachIPv4Request`

NewAttachIPv4Request instantiates a new AttachIPv4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachIPv4RequestWithDefaults

`func NewAttachIPv4RequestWithDefaults() *AttachIPv4Request`

NewAttachIPv4RequestWithDefaults instantiates a new AttachIPv4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *AttachIPv4Request) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *AttachIPv4Request) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *AttachIPv4Request) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.


### GetReboot

`func (o *AttachIPv4Request) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *AttachIPv4Request) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *AttachIPv4Request) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *AttachIPv4Request) HasReboot() bool`

HasReboot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


