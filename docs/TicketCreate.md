# TicketCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** |  | 
**Department** | **int32** |  | 
**Priority** | Pointer to [**TicketCreatePriorityEnum**](TicketCreatePriorityEnum.md) |  | [optional] [default to TICKETCREATEPRIORITYENUM_LOW]
**ServiceId** | Pointer to **NullableInt32** |  | [optional] 
**Message** | **string** |  | 
**Attachment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTicketCreate

`func NewTicketCreate(subject string, department int32, message string, ) *TicketCreate`

NewTicketCreate instantiates a new TicketCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCreateWithDefaults

`func NewTicketCreateWithDefaults() *TicketCreate`

NewTicketCreateWithDefaults instantiates a new TicketCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *TicketCreate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketCreate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketCreate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDepartment

`func (o *TicketCreate) GetDepartment() int32`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TicketCreate) GetDepartmentOk() (*int32, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TicketCreate) SetDepartment(v int32)`

SetDepartment sets Department field to given value.


### GetPriority

`func (o *TicketCreate) GetPriority() TicketCreatePriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketCreate) GetPriorityOk() (*TicketCreatePriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketCreate) SetPriority(v TicketCreatePriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TicketCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetServiceId

`func (o *TicketCreate) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TicketCreate) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TicketCreate) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *TicketCreate) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *TicketCreate) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *TicketCreate) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetMessage

`func (o *TicketCreate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TicketCreate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TicketCreate) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAttachment

`func (o *TicketCreate) GetAttachment() string`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TicketCreate) GetAttachmentOk() (*string, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TicketCreate) SetAttachment(v string)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TicketCreate) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### SetAttachmentNil

`func (o *TicketCreate) SetAttachmentNil(b bool)`

 SetAttachmentNil sets the value for Attachment to be an explicit nil

### UnsetAttachment
`func (o *TicketCreate) UnsetAttachment()`

UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


