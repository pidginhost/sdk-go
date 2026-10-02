# ContactsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactType** | [**ContactTypeEnum**](ContactTypeEnum.md) | Contact type to update  * &#x60;registrant&#x60; - registrant * &#x60;admin&#x60; - admin * &#x60;tech&#x60; - tech * &#x60;billing&#x60; - billing | 
**RegistrantId** | **int32** | ID of DomainRegistrant to use | 

## Methods

### NewContactsUpdateRequest

`func NewContactsUpdateRequest(contactType ContactTypeEnum, registrantId int32, ) *ContactsUpdateRequest`

NewContactsUpdateRequest instantiates a new ContactsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsUpdateRequestWithDefaults

`func NewContactsUpdateRequestWithDefaults() *ContactsUpdateRequest`

NewContactsUpdateRequestWithDefaults instantiates a new ContactsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactType

`func (o *ContactsUpdateRequest) GetContactType() ContactTypeEnum`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ContactsUpdateRequest) GetContactTypeOk() (*ContactTypeEnum, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ContactsUpdateRequest) SetContactType(v ContactTypeEnum)`

SetContactType sets ContactType field to given value.


### GetRegistrantId

`func (o *ContactsUpdateRequest) GetRegistrantId() int32`

GetRegistrantId returns the RegistrantId field if non-nil, zero value otherwise.

### GetRegistrantIdOk

`func (o *ContactsUpdateRequest) GetRegistrantIdOk() (*int32, bool)`

GetRegistrantIdOk returns a tuple with the RegistrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantId

`func (o *ContactsUpdateRequest) SetRegistrantId(v int32)`

SetRegistrantId sets RegistrantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


