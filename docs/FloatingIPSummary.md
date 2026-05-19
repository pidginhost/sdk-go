# FloatingIPSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Version** | [**VersionEnum**](VersionEnum.md) |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Label** | **string** |  | [readonly] 
**ReverseDns** | **string** |  | [readonly] 

## Methods

### NewFloatingIPSummary

`func NewFloatingIPSummary(id int32, version VersionEnum, address string, label string, reverseDns string, ) *FloatingIPSummary`

NewFloatingIPSummary instantiates a new FloatingIPSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIPSummaryWithDefaults

`func NewFloatingIPSummaryWithDefaults() *FloatingIPSummary`

NewFloatingIPSummaryWithDefaults instantiates a new FloatingIPSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FloatingIPSummary) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatingIPSummary) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatingIPSummary) SetId(v int32)`

SetId sets Id field to given value.


### GetVersion

`func (o *FloatingIPSummary) GetVersion() VersionEnum`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FloatingIPSummary) GetVersionOk() (*VersionEnum, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FloatingIPSummary) SetVersion(v VersionEnum)`

SetVersion sets Version field to given value.


### GetAddress

`func (o *FloatingIPSummary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FloatingIPSummary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FloatingIPSummary) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLabel

`func (o *FloatingIPSummary) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FloatingIPSummary) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FloatingIPSummary) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetReverseDns

`func (o *FloatingIPSummary) GetReverseDns() string`

GetReverseDns returns the ReverseDns field if non-nil, zero value otherwise.

### GetReverseDnsOk

`func (o *FloatingIPSummary) GetReverseDnsOk() (*string, bool)`

GetReverseDnsOk returns a tuple with the ReverseDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDns

`func (o *FloatingIPSummary) SetReverseDns(v string)`

SetReverseDns sets ReverseDns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


