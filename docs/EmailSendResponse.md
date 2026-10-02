# EmailSendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **interface{}** |  | 
**QueuedAt** | **string** |  | 
**Status** | [**EmailSendResponseStatusEnum**](EmailSendResponseStatusEnum.md) |  | 

## Methods

### NewEmailSendResponse

`func NewEmailSendResponse(messageId interface{}, queuedAt string, status EmailSendResponseStatusEnum, ) *EmailSendResponse`

NewEmailSendResponse instantiates a new EmailSendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendResponseWithDefaults

`func NewEmailSendResponseWithDefaults() *EmailSendResponse`

NewEmailSendResponseWithDefaults instantiates a new EmailSendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *EmailSendResponse) GetMessageId() interface{}`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EmailSendResponse) GetMessageIdOk() (*interface{}, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EmailSendResponse) SetMessageId(v interface{})`

SetMessageId sets MessageId field to given value.


### SetMessageIdNil

`func (o *EmailSendResponse) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *EmailSendResponse) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetQueuedAt

`func (o *EmailSendResponse) GetQueuedAt() string`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *EmailSendResponse) GetQueuedAtOk() (*string, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *EmailSendResponse) SetQueuedAt(v string)`

SetQueuedAt sets QueuedAt field to given value.


### GetStatus

`func (o *EmailSendResponse) GetStatus() EmailSendResponseStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailSendResponse) GetStatusOk() (*EmailSendResponseStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailSendResponse) SetStatus(v EmailSendResponseStatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


