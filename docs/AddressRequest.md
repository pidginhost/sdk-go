# AddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | [**CountryEnum**](CountryEnum.md) |  | 
**City** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**RegionFk** | Pointer to **NullableInt32** |  | [optional] 
**LocalityFk** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAddressRequest

`func NewAddressRequest(country CountryEnum, ) *AddressRequest`

NewAddressRequest instantiates a new AddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRequestWithDefaults

`func NewAddressRequestWithDefaults() *AddressRequest`

NewAddressRequestWithDefaults instantiates a new AddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *AddressRequest) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressRequest) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressRequest) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.


### GetCity

`func (o *AddressRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *AddressRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddressRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZipcode

`func (o *AddressRequest) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *AddressRequest) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *AddressRequest) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *AddressRequest) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetAddress

`func (o *AddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRegionFk

`func (o *AddressRequest) GetRegionFk() int32`

GetRegionFk returns the RegionFk field if non-nil, zero value otherwise.

### GetRegionFkOk

`func (o *AddressRequest) GetRegionFkOk() (*int32, bool)`

GetRegionFkOk returns a tuple with the RegionFk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionFk

`func (o *AddressRequest) SetRegionFk(v int32)`

SetRegionFk sets RegionFk field to given value.

### HasRegionFk

`func (o *AddressRequest) HasRegionFk() bool`

HasRegionFk returns a boolean if a field has been set.

### SetRegionFkNil

`func (o *AddressRequest) SetRegionFkNil(b bool)`

 SetRegionFkNil sets the value for RegionFk to be an explicit nil

### UnsetRegionFk
`func (o *AddressRequest) UnsetRegionFk()`

UnsetRegionFk ensures that no value is present for RegionFk, not even an explicit nil
### GetLocalityFk

`func (o *AddressRequest) GetLocalityFk() int32`

GetLocalityFk returns the LocalityFk field if non-nil, zero value otherwise.

### GetLocalityFkOk

`func (o *AddressRequest) GetLocalityFkOk() (*int32, bool)`

GetLocalityFkOk returns a tuple with the LocalityFk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityFk

`func (o *AddressRequest) SetLocalityFk(v int32)`

SetLocalityFk sets LocalityFk field to given value.

### HasLocalityFk

`func (o *AddressRequest) HasLocalityFk() bool`

HasLocalityFk returns a boolean if a field has been set.

### SetLocalityFkNil

`func (o *AddressRequest) SetLocalityFkNil(b bool)`

 SetLocalityFkNil sets the value for LocalityFk to be an explicit nil

### UnsetLocalityFk
`func (o *AddressRequest) UnsetLocalityFk()`

UnsetLocalityFk ensures that no value is present for LocalityFk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


