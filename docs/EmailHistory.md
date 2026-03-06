# EmailHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Subject** | **string** |  | [readonly] 
**Date** | **time.Time** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Read** | **bool** |  | [readonly] 

## Methods

### NewEmailHistory

`func NewEmailHistory(id int32, subject string, date time.Time, address string, read bool, ) *EmailHistory`

NewEmailHistory instantiates a new EmailHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailHistoryWithDefaults

`func NewEmailHistoryWithDefaults() *EmailHistory`

NewEmailHistoryWithDefaults instantiates a new EmailHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailHistory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailHistory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailHistory) SetId(v int32)`

SetId sets Id field to given value.


### GetSubject

`func (o *EmailHistory) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailHistory) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailHistory) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDate

`func (o *EmailHistory) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EmailHistory) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EmailHistory) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetAddress

`func (o *EmailHistory) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailHistory) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailHistory) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetRead

`func (o *EmailHistory) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *EmailHistory) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *EmailHistory) SetRead(v bool)`

SetRead sets Read field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


