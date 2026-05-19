# CLISessionPollResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**CLISessionPollResponseStatusEnum**](CLISessionPollResponseStatusEnum.md) |  | 
**TokenKey** | Pointer to **string** | Only present when status is approved | [optional] 
**TokenName** | Pointer to **string** |  | [optional] 
**TokenScope** | Pointer to **string** |  | [optional] 

## Methods

### NewCLISessionPollResponse

`func NewCLISessionPollResponse(status CLISessionPollResponseStatusEnum, ) *CLISessionPollResponse`

NewCLISessionPollResponse instantiates a new CLISessionPollResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCLISessionPollResponseWithDefaults

`func NewCLISessionPollResponseWithDefaults() *CLISessionPollResponse`

NewCLISessionPollResponseWithDefaults instantiates a new CLISessionPollResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CLISessionPollResponse) GetStatus() CLISessionPollResponseStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CLISessionPollResponse) GetStatusOk() (*CLISessionPollResponseStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CLISessionPollResponse) SetStatus(v CLISessionPollResponseStatusEnum)`

SetStatus sets Status field to given value.


### GetTokenKey

`func (o *CLISessionPollResponse) GetTokenKey() string`

GetTokenKey returns the TokenKey field if non-nil, zero value otherwise.

### GetTokenKeyOk

`func (o *CLISessionPollResponse) GetTokenKeyOk() (*string, bool)`

GetTokenKeyOk returns a tuple with the TokenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenKey

`func (o *CLISessionPollResponse) SetTokenKey(v string)`

SetTokenKey sets TokenKey field to given value.

### HasTokenKey

`func (o *CLISessionPollResponse) HasTokenKey() bool`

HasTokenKey returns a boolean if a field has been set.

### GetTokenName

`func (o *CLISessionPollResponse) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *CLISessionPollResponse) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *CLISessionPollResponse) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *CLISessionPollResponse) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTokenScope

`func (o *CLISessionPollResponse) GetTokenScope() string`

GetTokenScope returns the TokenScope field if non-nil, zero value otherwise.

### GetTokenScopeOk

`func (o *CLISessionPollResponse) GetTokenScopeOk() (*string, bool)`

GetTokenScopeOk returns a tuple with the TokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScope

`func (o *CLISessionPollResponse) SetTokenScope(v string)`

SetTokenScope sets TokenScope field to given value.

### HasTokenScope

`func (o *CLISessionPollResponse) HasTokenScope() bool`

HasTokenScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


