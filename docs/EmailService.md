# EmailService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tier** | **string** |  | [readonly] 
**Status** | [**StatusA57Enum**](StatusA57Enum.md) |  | [readonly] 
**SandboxMode** | **bool** |  | [readonly] 
**AutoSuspended** | **bool** |  | [readonly] 
**AutoSuspendReason** | **string** |  | [readonly] 
**MsgsSent24h** | **int32** |  | [readonly] 
**MsgsSent30d** | **int32** |  | [readonly] 
**BounceRatePct** | **float64** |  | [readonly] 
**ComplaintRatePct** | **float64** |  | [readonly] 
**DedicatedIpAddon** | **bool** |  | [readonly] 
**QuotaMonthly** | **string** |  | [readonly] 
**PriceMonthlyEur** | **string** |  | [readonly] 

## Methods

### NewEmailService

`func NewEmailService(id int32, tier string, status StatusA57Enum, sandboxMode bool, autoSuspended bool, autoSuspendReason string, msgsSent24h int32, msgsSent30d int32, bounceRatePct float64, complaintRatePct float64, dedicatedIpAddon bool, quotaMonthly string, priceMonthlyEur string, ) *EmailService`

NewEmailService instantiates a new EmailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailServiceWithDefaults

`func NewEmailServiceWithDefaults() *EmailService`

NewEmailServiceWithDefaults instantiates a new EmailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailService) SetId(v int32)`

SetId sets Id field to given value.


### GetTier

`func (o *EmailService) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *EmailService) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *EmailService) SetTier(v string)`

SetTier sets Tier field to given value.


### GetStatus

`func (o *EmailService) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailService) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailService) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.


### GetSandboxMode

`func (o *EmailService) GetSandboxMode() bool`

GetSandboxMode returns the SandboxMode field if non-nil, zero value otherwise.

### GetSandboxModeOk

`func (o *EmailService) GetSandboxModeOk() (*bool, bool)`

GetSandboxModeOk returns a tuple with the SandboxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxMode

`func (o *EmailService) SetSandboxMode(v bool)`

SetSandboxMode sets SandboxMode field to given value.


### GetAutoSuspended

`func (o *EmailService) GetAutoSuspended() bool`

GetAutoSuspended returns the AutoSuspended field if non-nil, zero value otherwise.

### GetAutoSuspendedOk

`func (o *EmailService) GetAutoSuspendedOk() (*bool, bool)`

GetAutoSuspendedOk returns a tuple with the AutoSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspended

`func (o *EmailService) SetAutoSuspended(v bool)`

SetAutoSuspended sets AutoSuspended field to given value.


### GetAutoSuspendReason

`func (o *EmailService) GetAutoSuspendReason() string`

GetAutoSuspendReason returns the AutoSuspendReason field if non-nil, zero value otherwise.

### GetAutoSuspendReasonOk

`func (o *EmailService) GetAutoSuspendReasonOk() (*string, bool)`

GetAutoSuspendReasonOk returns a tuple with the AutoSuspendReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspendReason

`func (o *EmailService) SetAutoSuspendReason(v string)`

SetAutoSuspendReason sets AutoSuspendReason field to given value.


### GetMsgsSent24h

`func (o *EmailService) GetMsgsSent24h() int32`

GetMsgsSent24h returns the MsgsSent24h field if non-nil, zero value otherwise.

### GetMsgsSent24hOk

`func (o *EmailService) GetMsgsSent24hOk() (*int32, bool)`

GetMsgsSent24hOk returns a tuple with the MsgsSent24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsSent24h

`func (o *EmailService) SetMsgsSent24h(v int32)`

SetMsgsSent24h sets MsgsSent24h field to given value.


### GetMsgsSent30d

`func (o *EmailService) GetMsgsSent30d() int32`

GetMsgsSent30d returns the MsgsSent30d field if non-nil, zero value otherwise.

### GetMsgsSent30dOk

`func (o *EmailService) GetMsgsSent30dOk() (*int32, bool)`

GetMsgsSent30dOk returns a tuple with the MsgsSent30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsSent30d

`func (o *EmailService) SetMsgsSent30d(v int32)`

SetMsgsSent30d sets MsgsSent30d field to given value.


### GetBounceRatePct

`func (o *EmailService) GetBounceRatePct() float64`

GetBounceRatePct returns the BounceRatePct field if non-nil, zero value otherwise.

### GetBounceRatePctOk

`func (o *EmailService) GetBounceRatePctOk() (*float64, bool)`

GetBounceRatePctOk returns a tuple with the BounceRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRatePct

`func (o *EmailService) SetBounceRatePct(v float64)`

SetBounceRatePct sets BounceRatePct field to given value.


### GetComplaintRatePct

`func (o *EmailService) GetComplaintRatePct() float64`

GetComplaintRatePct returns the ComplaintRatePct field if non-nil, zero value otherwise.

### GetComplaintRatePctOk

`func (o *EmailService) GetComplaintRatePctOk() (*float64, bool)`

GetComplaintRatePctOk returns a tuple with the ComplaintRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaintRatePct

`func (o *EmailService) SetComplaintRatePct(v float64)`

SetComplaintRatePct sets ComplaintRatePct field to given value.


### GetDedicatedIpAddon

`func (o *EmailService) GetDedicatedIpAddon() bool`

GetDedicatedIpAddon returns the DedicatedIpAddon field if non-nil, zero value otherwise.

### GetDedicatedIpAddonOk

`func (o *EmailService) GetDedicatedIpAddonOk() (*bool, bool)`

GetDedicatedIpAddonOk returns a tuple with the DedicatedIpAddon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedIpAddon

`func (o *EmailService) SetDedicatedIpAddon(v bool)`

SetDedicatedIpAddon sets DedicatedIpAddon field to given value.


### GetQuotaMonthly

`func (o *EmailService) GetQuotaMonthly() string`

GetQuotaMonthly returns the QuotaMonthly field if non-nil, zero value otherwise.

### GetQuotaMonthlyOk

`func (o *EmailService) GetQuotaMonthlyOk() (*string, bool)`

GetQuotaMonthlyOk returns a tuple with the QuotaMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMonthly

`func (o *EmailService) SetQuotaMonthly(v string)`

SetQuotaMonthly sets QuotaMonthly field to given value.


### GetPriceMonthlyEur

`func (o *EmailService) GetPriceMonthlyEur() string`

GetPriceMonthlyEur returns the PriceMonthlyEur field if non-nil, zero value otherwise.

### GetPriceMonthlyEurOk

`func (o *EmailService) GetPriceMonthlyEurOk() (*string, bool)`

GetPriceMonthlyEurOk returns a tuple with the PriceMonthlyEur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthlyEur

`func (o *EmailService) SetPriceMonthlyEur(v string)`

SetPriceMonthlyEur sets PriceMonthlyEur field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


