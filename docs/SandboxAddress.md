# SandboxAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Address** | **string** |  | 
**VerifiedAt** | **NullableString** |  | [readonly] 
**CreatedAt** | **string** |  | [readonly] 

## Methods

### NewSandboxAddress

`func NewSandboxAddress(id int32, address string, verifiedAt NullableString, createdAt string, ) *SandboxAddress`

NewSandboxAddress instantiates a new SandboxAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxAddressWithDefaults

`func NewSandboxAddressWithDefaults() *SandboxAddress`

NewSandboxAddressWithDefaults instantiates a new SandboxAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SandboxAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *SandboxAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SandboxAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SandboxAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVerifiedAt

`func (o *SandboxAddress) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *SandboxAddress) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *SandboxAddress) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.


### SetVerifiedAtNil

`func (o *SandboxAddress) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *SandboxAddress) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetCreatedAt

`func (o *SandboxAddress) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxAddress) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxAddress) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


