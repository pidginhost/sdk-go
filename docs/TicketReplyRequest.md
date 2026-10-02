# TicketReplyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Attachment** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewTicketReplyRequest

`func NewTicketReplyRequest(message string, ) *TicketReplyRequest`

NewTicketReplyRequest instantiates a new TicketReplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketReplyRequestWithDefaults

`func NewTicketReplyRequestWithDefaults() *TicketReplyRequest`

NewTicketReplyRequestWithDefaults instantiates a new TicketReplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TicketReplyRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TicketReplyRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TicketReplyRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAttachment

`func (o *TicketReplyRequest) GetAttachment() *os.File`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TicketReplyRequest) GetAttachmentOk() (**os.File, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TicketReplyRequest) SetAttachment(v *os.File)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TicketReplyRequest) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


