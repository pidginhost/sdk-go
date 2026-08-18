# TicketReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Attachment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTicketReply

`func NewTicketReply(message string, ) *TicketReply`

NewTicketReply instantiates a new TicketReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketReplyWithDefaults

`func NewTicketReplyWithDefaults() *TicketReply`

NewTicketReplyWithDefaults instantiates a new TicketReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TicketReply) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TicketReply) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TicketReply) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAttachment

`func (o *TicketReply) GetAttachment() string`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TicketReply) GetAttachmentOk() (*string, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TicketReply) SetAttachment(v string)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TicketReply) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### SetAttachmentNil

`func (o *TicketReply) SetAttachmentNil(b bool)`

 SetAttachmentNil sets the value for Attachment to be an explicit nil

### UnsetAttachment
`func (o *TicketReply) UnsetAttachment()`

UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


