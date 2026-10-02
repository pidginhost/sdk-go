# EmailMessageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]EmailMessageSummary**](EmailMessageSummary.md) |  | 
**Count** | **int32** |  | 
**Page** | **int32** |  | 
**PerPage** | **int32** |  | 

## Methods

### NewEmailMessageList

`func NewEmailMessageList(results []EmailMessageSummary, count int32, page int32, perPage int32, ) *EmailMessageList`

NewEmailMessageList instantiates a new EmailMessageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageListWithDefaults

`func NewEmailMessageListWithDefaults() *EmailMessageList`

NewEmailMessageListWithDefaults instantiates a new EmailMessageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *EmailMessageList) GetResults() []EmailMessageSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EmailMessageList) GetResultsOk() (*[]EmailMessageSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EmailMessageList) SetResults(v []EmailMessageSummary)`

SetResults sets Results field to given value.


### GetCount

`func (o *EmailMessageList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EmailMessageList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EmailMessageList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *EmailMessageList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EmailMessageList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EmailMessageList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *EmailMessageList) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *EmailMessageList) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *EmailMessageList) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


