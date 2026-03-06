# TicketList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Subject** | **string** |  | [readonly] 
**Department** | **string** |  | [readonly] 
**Priority** | [**Priority3cdEnum**](Priority3cdEnum.md) |  | [readonly] 
**Status** | [**StatusEf2Enum**](StatusEf2Enum.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewTicketList

`func NewTicketList(id int32, subject string, department string, priority Priority3cdEnum, status StatusEf2Enum, created time.Time, updated time.Time, ) *TicketList`

NewTicketList instantiates a new TicketList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketListWithDefaults

`func NewTicketListWithDefaults() *TicketList`

NewTicketListWithDefaults instantiates a new TicketList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketList) SetId(v int32)`

SetId sets Id field to given value.


### GetSubject

`func (o *TicketList) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketList) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketList) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDepartment

`func (o *TicketList) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TicketList) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TicketList) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetPriority

`func (o *TicketList) GetPriority() Priority3cdEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketList) GetPriorityOk() (*Priority3cdEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketList) SetPriority(v Priority3cdEnum)`

SetPriority sets Priority field to given value.


### GetStatus

`func (o *TicketList) GetStatus() StatusEf2Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketList) GetStatusOk() (*StatusEf2Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketList) SetStatus(v StatusEf2Enum)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *TicketList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TicketList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TicketList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *TicketList) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TicketList) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TicketList) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


