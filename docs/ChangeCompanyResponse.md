# ChangeCompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Company** | **map[string]interface{}** |  | 

## Methods

### NewChangeCompanyResponse

`func NewChangeCompanyResponse(message string, company map[string]interface{}, ) *ChangeCompanyResponse`

NewChangeCompanyResponse instantiates a new ChangeCompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeCompanyResponseWithDefaults

`func NewChangeCompanyResponseWithDefaults() *ChangeCompanyResponse`

NewChangeCompanyResponseWithDefaults instantiates a new ChangeCompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ChangeCompanyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChangeCompanyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChangeCompanyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCompany

`func (o *ChangeCompanyResponse) GetCompany() map[string]interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ChangeCompanyResponse) GetCompanyOk() (*map[string]interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ChangeCompanyResponse) SetCompany(v map[string]interface{})`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *ChangeCompanyResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ChangeCompanyResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


