# FreeDNSDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | [readonly] 
**Active** | **bool** |  | [readonly] [default to true]
**Source** | **string** |  | [readonly] 

## Methods

### NewFreeDNSDomain

`func NewFreeDNSDomain(domain string, active bool, source string, ) *FreeDNSDomain`

NewFreeDNSDomain instantiates a new FreeDNSDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeDNSDomainWithDefaults

`func NewFreeDNSDomainWithDefaults() *FreeDNSDomain`

NewFreeDNSDomainWithDefaults instantiates a new FreeDNSDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *FreeDNSDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *FreeDNSDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *FreeDNSDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetActive

`func (o *FreeDNSDomain) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FreeDNSDomain) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FreeDNSDomain) SetActive(v bool)`

SetActive sets Active field to given value.


### GetSource

`func (o *FreeDNSDomain) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FreeDNSDomain) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FreeDNSDomain) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


