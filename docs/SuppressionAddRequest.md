# SuppressionAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewSuppressionAddRequest

`func NewSuppressionAddRequest(address string, ) *SuppressionAddRequest`

NewSuppressionAddRequest instantiates a new SuppressionAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionAddRequestWithDefaults

`func NewSuppressionAddRequestWithDefaults() *SuppressionAddRequest`

NewSuppressionAddRequestWithDefaults instantiates a new SuppressionAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SuppressionAddRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SuppressionAddRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SuppressionAddRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDetail

`func (o *SuppressionAddRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SuppressionAddRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SuppressionAddRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SuppressionAddRequest) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


