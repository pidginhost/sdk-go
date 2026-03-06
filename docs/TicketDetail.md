# TicketDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Subject** | **string** |  | [readonly] 
**Department** | [**Department**](Department.md) |  | [readonly] 
**Priority** | [**Priority3cdEnum**](Priority3cdEnum.md) |  | [readonly] 
**Status** | [**StatusEf2Enum**](StatusEf2Enum.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Updated** | **time.Time** |  | [readonly] 
**Messages** | **string** |  | [readonly] 

## Methods

### NewTicketDetail

`func NewTicketDetail(id int32, subject string, department Department, priority Priority3cdEnum, status StatusEf2Enum, created time.Time, updated time.Time, messages string, ) *TicketDetail`

NewTicketDetail instantiates a new TicketDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketDetailWithDefaults

`func NewTicketDetailWithDefaults() *TicketDetail`

NewTicketDetailWithDefaults instantiates a new TicketDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetSubject

`func (o *TicketDetail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketDetail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketDetail) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDepartment

`func (o *TicketDetail) GetDepartment() Department`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TicketDetail) GetDepartmentOk() (*Department, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TicketDetail) SetDepartment(v Department)`

SetDepartment sets Department field to given value.


### GetPriority

`func (o *TicketDetail) GetPriority() Priority3cdEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketDetail) GetPriorityOk() (*Priority3cdEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketDetail) SetPriority(v Priority3cdEnum)`

SetPriority sets Priority field to given value.


### GetStatus

`func (o *TicketDetail) GetStatus() StatusEf2Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketDetail) GetStatusOk() (*StatusEf2Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketDetail) SetStatus(v StatusEf2Enum)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *TicketDetail) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TicketDetail) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TicketDetail) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *TicketDetail) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TicketDetail) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TicketDetail) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetMessages

`func (o *TicketDetail) GetMessages() string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *TicketDetail) GetMessagesOk() (*string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *TicketDetail) SetMessages(v string)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


