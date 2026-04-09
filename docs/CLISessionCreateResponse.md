# CLISessionCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**VerificationUrl** | **string** |  | 
**ExpiresAt** | **string** |  | 
**PollInterval** | **int32** | Recommended polling interval in seconds | 

## Methods

### NewCLISessionCreateResponse

`func NewCLISessionCreateResponse(sessionId string, verificationUrl string, expiresAt string, pollInterval int32, ) *CLISessionCreateResponse`

NewCLISessionCreateResponse instantiates a new CLISessionCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCLISessionCreateResponseWithDefaults

`func NewCLISessionCreateResponseWithDefaults() *CLISessionCreateResponse`

NewCLISessionCreateResponseWithDefaults instantiates a new CLISessionCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *CLISessionCreateResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CLISessionCreateResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CLISessionCreateResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetVerificationUrl

`func (o *CLISessionCreateResponse) GetVerificationUrl() string`

GetVerificationUrl returns the VerificationUrl field if non-nil, zero value otherwise.

### GetVerificationUrlOk

`func (o *CLISessionCreateResponse) GetVerificationUrlOk() (*string, bool)`

GetVerificationUrlOk returns a tuple with the VerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUrl

`func (o *CLISessionCreateResponse) SetVerificationUrl(v string)`

SetVerificationUrl sets VerificationUrl field to given value.


### GetExpiresAt

`func (o *CLISessionCreateResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CLISessionCreateResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CLISessionCreateResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetPollInterval

`func (o *CLISessionCreateResponse) GetPollInterval() int32`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *CLISessionCreateResponse) GetPollIntervalOk() (*int32, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *CLISessionCreateResponse) SetPollInterval(v int32)`

SetPollInterval sets PollInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


