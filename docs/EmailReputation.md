# EmailReputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BounceRatePct** | **string** |  | 
**ComplaintRatePct** | **string** |  | 
**MsgsSent24h** | **int32** |  | 
**MsgsSent30d** | **int32** |  | 

## Methods

### NewEmailReputation

`func NewEmailReputation(bounceRatePct string, complaintRatePct string, msgsSent24h int32, msgsSent30d int32, ) *EmailReputation`

NewEmailReputation instantiates a new EmailReputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailReputationWithDefaults

`func NewEmailReputationWithDefaults() *EmailReputation`

NewEmailReputationWithDefaults instantiates a new EmailReputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounceRatePct

`func (o *EmailReputation) GetBounceRatePct() string`

GetBounceRatePct returns the BounceRatePct field if non-nil, zero value otherwise.

### GetBounceRatePctOk

`func (o *EmailReputation) GetBounceRatePctOk() (*string, bool)`

GetBounceRatePctOk returns a tuple with the BounceRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRatePct

`func (o *EmailReputation) SetBounceRatePct(v string)`

SetBounceRatePct sets BounceRatePct field to given value.


### GetComplaintRatePct

`func (o *EmailReputation) GetComplaintRatePct() string`

GetComplaintRatePct returns the ComplaintRatePct field if non-nil, zero value otherwise.

### GetComplaintRatePctOk

`func (o *EmailReputation) GetComplaintRatePctOk() (*string, bool)`

GetComplaintRatePctOk returns a tuple with the ComplaintRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaintRatePct

`func (o *EmailReputation) SetComplaintRatePct(v string)`

SetComplaintRatePct sets ComplaintRatePct field to given value.


### GetMsgsSent24h

`func (o *EmailReputation) GetMsgsSent24h() int32`

GetMsgsSent24h returns the MsgsSent24h field if non-nil, zero value otherwise.

### GetMsgsSent24hOk

`func (o *EmailReputation) GetMsgsSent24hOk() (*int32, bool)`

GetMsgsSent24hOk returns a tuple with the MsgsSent24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsSent24h

`func (o *EmailReputation) SetMsgsSent24h(v int32)`

SetMsgsSent24h sets MsgsSent24h field to given value.


### GetMsgsSent30d

`func (o *EmailReputation) GetMsgsSent30d() int32`

GetMsgsSent30d returns the MsgsSent30d field if non-nil, zero value otherwise.

### GetMsgsSent30dOk

`func (o *EmailReputation) GetMsgsSent30dOk() (*int32, bool)`

GetMsgsSent30dOk returns a tuple with the MsgsSent30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsSent30d

`func (o *EmailReputation) SetMsgsSent30d(v int32)`

SetMsgsSent30d sets MsgsSent30d field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


