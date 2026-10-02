# TicketCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** |  | 
**Department** | **int32** |  | 
**Priority** | Pointer to [**TicketCreatePriorityEnum**](TicketCreatePriorityEnum.md) |  | [optional] [default to TICKETCREATEPRIORITYENUM_LOW]
**ServiceId** | Pointer to **NullableInt32** |  | [optional] 
**Message** | **string** |  | 
**Attachment** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewTicketCreateRequest

`func NewTicketCreateRequest(subject string, department int32, message string, ) *TicketCreateRequest`

NewTicketCreateRequest instantiates a new TicketCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCreateRequestWithDefaults

`func NewTicketCreateRequestWithDefaults() *TicketCreateRequest`

NewTicketCreateRequestWithDefaults instantiates a new TicketCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *TicketCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDepartment

`func (o *TicketCreateRequest) GetDepartment() int32`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TicketCreateRequest) GetDepartmentOk() (*int32, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TicketCreateRequest) SetDepartment(v int32)`

SetDepartment sets Department field to given value.


### GetPriority

`func (o *TicketCreateRequest) GetPriority() TicketCreatePriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketCreateRequest) GetPriorityOk() (*TicketCreatePriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketCreateRequest) SetPriority(v TicketCreatePriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TicketCreateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetServiceId

`func (o *TicketCreateRequest) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TicketCreateRequest) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TicketCreateRequest) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *TicketCreateRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *TicketCreateRequest) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *TicketCreateRequest) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetMessage

`func (o *TicketCreateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TicketCreateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TicketCreateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAttachment

`func (o *TicketCreateRequest) GetAttachment() *os.File`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TicketCreateRequest) GetAttachmentOk() (**os.File, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TicketCreateRequest) SetAttachment(v *os.File)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TicketCreateRequest) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


