# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Country** | [**CountryEnum**](CountryEnum.md) |  | 
**City** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**RegionFk** | Pointer to **NullableInt32** |  | [optional] 
**LocalityFk** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAddress

`func NewAddress(id int32, country CountryEnum, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v int32)`

SetId sets Id field to given value.


### GetCountry

`func (o *Address) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.


### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *Address) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Address) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Address) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Address) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZipcode

`func (o *Address) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *Address) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *Address) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *Address) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Address) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRegionFk

`func (o *Address) GetRegionFk() int32`

GetRegionFk returns the RegionFk field if non-nil, zero value otherwise.

### GetRegionFkOk

`func (o *Address) GetRegionFkOk() (*int32, bool)`

GetRegionFkOk returns a tuple with the RegionFk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionFk

`func (o *Address) SetRegionFk(v int32)`

SetRegionFk sets RegionFk field to given value.

### HasRegionFk

`func (o *Address) HasRegionFk() bool`

HasRegionFk returns a boolean if a field has been set.

### SetRegionFkNil

`func (o *Address) SetRegionFkNil(b bool)`

 SetRegionFkNil sets the value for RegionFk to be an explicit nil

### UnsetRegionFk
`func (o *Address) UnsetRegionFk()`

UnsetRegionFk ensures that no value is present for RegionFk, not even an explicit nil
### GetLocalityFk

`func (o *Address) GetLocalityFk() int32`

GetLocalityFk returns the LocalityFk field if non-nil, zero value otherwise.

### GetLocalityFkOk

`func (o *Address) GetLocalityFkOk() (*int32, bool)`

GetLocalityFkOk returns a tuple with the LocalityFk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityFk

`func (o *Address) SetLocalityFk(v int32)`

SetLocalityFk sets LocalityFk field to given value.

### HasLocalityFk

`func (o *Address) HasLocalityFk() bool`

HasLocalityFk returns a boolean if a field has been set.

### SetLocalityFkNil

`func (o *Address) SetLocalityFkNil(b bool)`

 SetLocalityFkNil sets the value for LocalityFk to be an explicit nil

### UnsetLocalityFk
`func (o *Address) UnsetLocalityFk()`

UnsetLocalityFk ensures that no value is present for LocalityFk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


