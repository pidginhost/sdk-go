# TicketMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Date** | **string** |  | [readonly] 
**Message** | **string** |  | [readonly] 
**AuthorName** | **string** |  | [readonly] 
**HasAttachment** | **string** |  | [readonly] 
**AttachmentFilename** | **string** |  | [readonly] 

## Methods

### NewTicketMessage

`func NewTicketMessage(id int32, date string, message string, authorName string, hasAttachment string, attachmentFilename string, ) *TicketMessage`

NewTicketMessage instantiates a new TicketMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketMessageWithDefaults

`func NewTicketMessageWithDefaults() *TicketMessage`

NewTicketMessageWithDefaults instantiates a new TicketMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetDate

`func (o *TicketMessage) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TicketMessage) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TicketMessage) SetDate(v string)`

SetDate sets Date field to given value.


### GetMessage

`func (o *TicketMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TicketMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TicketMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthorName

`func (o *TicketMessage) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *TicketMessage) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *TicketMessage) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetHasAttachment

`func (o *TicketMessage) GetHasAttachment() string`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *TicketMessage) GetHasAttachmentOk() (*string, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *TicketMessage) SetHasAttachment(v string)`

SetHasAttachment sets HasAttachment field to given value.


### GetAttachmentFilename

`func (o *TicketMessage) GetAttachmentFilename() string`

GetAttachmentFilename returns the AttachmentFilename field if non-nil, zero value otherwise.

### GetAttachmentFilenameOk

`func (o *TicketMessage) GetAttachmentFilenameOk() (*string, bool)`

GetAttachmentFilenameOk returns a tuple with the AttachmentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentFilename

`func (o *TicketMessage) SetAttachmentFilename(v string)`

SetAttachmentFilename sets AttachmentFilename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


