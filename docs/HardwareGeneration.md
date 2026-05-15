# HardwareGeneration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CpuLabel** | Pointer to **string** |  | [optional] 
**PriceMultiplier** | Pointer to **float64** | Multiplier applied to base MeteredProduct price (1.00 &#x3D; no change, 1.50 &#x3D; +50%) | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IconClass** | Pointer to **string** |  | [optional] 
**FreeTierEligible** | Pointer to **bool** | Whether free tier VM plans are available on this generation | [optional] 

## Methods

### NewHardwareGeneration

`func NewHardwareGeneration(slug string, name string, ) *HardwareGeneration`

NewHardwareGeneration instantiates a new HardwareGeneration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareGenerationWithDefaults

`func NewHardwareGenerationWithDefaults() *HardwareGeneration`

NewHardwareGenerationWithDefaults instantiates a new HardwareGeneration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *HardwareGeneration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *HardwareGeneration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *HardwareGeneration) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *HardwareGeneration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HardwareGeneration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HardwareGeneration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HardwareGeneration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HardwareGeneration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HardwareGeneration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HardwareGeneration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCpuLabel

`func (o *HardwareGeneration) GetCpuLabel() string`

GetCpuLabel returns the CpuLabel field if non-nil, zero value otherwise.

### GetCpuLabelOk

`func (o *HardwareGeneration) GetCpuLabelOk() (*string, bool)`

GetCpuLabelOk returns a tuple with the CpuLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLabel

`func (o *HardwareGeneration) SetCpuLabel(v string)`

SetCpuLabel sets CpuLabel field to given value.

### HasCpuLabel

`func (o *HardwareGeneration) HasCpuLabel() bool`

HasCpuLabel returns a boolean if a field has been set.

### GetPriceMultiplier

`func (o *HardwareGeneration) GetPriceMultiplier() float64`

GetPriceMultiplier returns the PriceMultiplier field if non-nil, zero value otherwise.

### GetPriceMultiplierOk

`func (o *HardwareGeneration) GetPriceMultiplierOk() (*float64, bool)`

GetPriceMultiplierOk returns a tuple with the PriceMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMultiplier

`func (o *HardwareGeneration) SetPriceMultiplier(v float64)`

SetPriceMultiplier sets PriceMultiplier field to given value.

### HasPriceMultiplier

`func (o *HardwareGeneration) HasPriceMultiplier() bool`

HasPriceMultiplier returns a boolean if a field has been set.

### GetIsDefault

`func (o *HardwareGeneration) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *HardwareGeneration) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *HardwareGeneration) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *HardwareGeneration) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIconClass

`func (o *HardwareGeneration) GetIconClass() string`

GetIconClass returns the IconClass field if non-nil, zero value otherwise.

### GetIconClassOk

`func (o *HardwareGeneration) GetIconClassOk() (*string, bool)`

GetIconClassOk returns a tuple with the IconClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconClass

`func (o *HardwareGeneration) SetIconClass(v string)`

SetIconClass sets IconClass field to given value.

### HasIconClass

`func (o *HardwareGeneration) HasIconClass() bool`

HasIconClass returns a boolean if a field has been set.

### GetFreeTierEligible

`func (o *HardwareGeneration) GetFreeTierEligible() bool`

GetFreeTierEligible returns the FreeTierEligible field if non-nil, zero value otherwise.

### GetFreeTierEligibleOk

`func (o *HardwareGeneration) GetFreeTierEligibleOk() (*bool, bool)`

GetFreeTierEligibleOk returns a tuple with the FreeTierEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTierEligible

`func (o *HardwareGeneration) SetFreeTierEligible(v bool)`

SetFreeTierEligible sets FreeTierEligible field to given value.

### HasFreeTierEligible

`func (o *HardwareGeneration) HasFreeTierEligible() bool`

HasFreeTierEligible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


