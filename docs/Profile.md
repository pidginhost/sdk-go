# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Funds** | **float64** |  | [readonly] 
**Phone** | **string** |  | 

## Methods

### NewProfile

`func NewProfile(firstName string, lastName string, funds float64, phone string, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *Profile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Profile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Profile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Profile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Profile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Profile) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFunds

`func (o *Profile) GetFunds() float64`

GetFunds returns the Funds field if non-nil, zero value otherwise.

### GetFundsOk

`func (o *Profile) GetFundsOk() (*float64, bool)`

GetFundsOk returns a tuple with the Funds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunds

`func (o *Profile) SetFunds(v float64)`

SetFunds sets Funds field to given value.


### GetPhone

`func (o *Profile) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Profile) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Profile) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


