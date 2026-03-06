# \BillingAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingDepositsCreate**](BillingAPI.md#BillingDepositsCreate) | **Post** /api/billing/deposits/ | 
[**BillingDepositsCreate2**](BillingAPI.md#BillingDepositsCreate2) | **Post** /api/v1/billing/deposits/ | 
[**BillingDepositsList**](BillingAPI.md#BillingDepositsList) | **Get** /api/billing/deposits/ | 
[**BillingDepositsList2**](BillingAPI.md#BillingDepositsList2) | **Get** /api/v1/billing/deposits/ | 
[**BillingDepositsRetrieve**](BillingAPI.md#BillingDepositsRetrieve) | **Get** /api/billing/deposits/{id}/ | 
[**BillingDepositsRetrieve2**](BillingAPI.md#BillingDepositsRetrieve2) | **Get** /api/v1/billing/deposits/{id}/ | 
[**BillingFundsList**](BillingAPI.md#BillingFundsList) | **Get** /api/billing/funds/ | 
[**BillingFundsList2**](BillingAPI.md#BillingFundsList2) | **Get** /api/v1/billing/funds/ | 
[**BillingFundsLogList**](BillingAPI.md#BillingFundsLogList) | **Get** /api/billing/funds-log/ | 
[**BillingFundsLogList2**](BillingAPI.md#BillingFundsLogList2) | **Get** /api/v1/billing/funds-log/ | 
[**BillingFundsNotificationSettingsCreate**](BillingAPI.md#BillingFundsNotificationSettingsCreate) | **Post** /api/billing/funds/notification-settings/ | 
[**BillingFundsNotificationSettingsCreate2**](BillingAPI.md#BillingFundsNotificationSettingsCreate2) | **Post** /api/v1/billing/funds/notification-settings/ | 
[**BillingInvoicesList**](BillingAPI.md#BillingInvoicesList) | **Get** /api/billing/invoices/ | 
[**BillingInvoicesList2**](BillingAPI.md#BillingInvoicesList2) | **Get** /api/v1/billing/invoices/ | 
[**BillingInvoicesPayWithFundsCreate**](BillingAPI.md#BillingInvoicesPayWithFundsCreate) | **Post** /api/billing/invoices/{id}/pay-with-funds/ | 
[**BillingInvoicesPayWithFundsCreate2**](BillingAPI.md#BillingInvoicesPayWithFundsCreate2) | **Post** /api/v1/billing/invoices/{id}/pay-with-funds/ | 
[**BillingInvoicesPdfRetrieve**](BillingAPI.md#BillingInvoicesPdfRetrieve) | **Get** /api/billing/invoices/{id}/pdf/ | 
[**BillingInvoicesPdfRetrieve2**](BillingAPI.md#BillingInvoicesPdfRetrieve2) | **Get** /api/v1/billing/invoices/{id}/pdf/ | 
[**BillingInvoicesRetrieve**](BillingAPI.md#BillingInvoicesRetrieve) | **Get** /api/billing/invoices/{id}/ | 
[**BillingInvoicesRetrieve2**](BillingAPI.md#BillingInvoicesRetrieve2) | **Get** /api/v1/billing/invoices/{id}/ | 
[**BillingServicesCancelCreate**](BillingAPI.md#BillingServicesCancelCreate) | **Post** /api/billing/services/{id}/cancel/ | 
[**BillingServicesCancelCreate2**](BillingAPI.md#BillingServicesCancelCreate2) | **Post** /api/v1/billing/services/{id}/cancel/ | 
[**BillingServicesChangeBillingCycleCreate**](BillingAPI.md#BillingServicesChangeBillingCycleCreate) | **Post** /api/billing/services/{id}/change-billing-cycle/ | 
[**BillingServicesChangeBillingCycleCreate2**](BillingAPI.md#BillingServicesChangeBillingCycleCreate2) | **Post** /api/v1/billing/services/{id}/change-billing-cycle/ | 
[**BillingServicesChangeCompanyCreate**](BillingAPI.md#BillingServicesChangeCompanyCreate) | **Post** /api/billing/services/{id}/change-company/ | 
[**BillingServicesChangeCompanyCreate2**](BillingAPI.md#BillingServicesChangeCompanyCreate2) | **Post** /api/v1/billing/services/{id}/change-company/ | 
[**BillingServicesList**](BillingAPI.md#BillingServicesList) | **Get** /api/billing/services/ | 
[**BillingServicesList2**](BillingAPI.md#BillingServicesList2) | **Get** /api/v1/billing/services/ | 
[**BillingServicesRetrieve**](BillingAPI.md#BillingServicesRetrieve) | **Get** /api/billing/services/{id}/ | 
[**BillingServicesRetrieve2**](BillingAPI.md#BillingServicesRetrieve2) | **Get** /api/v1/billing/services/{id}/ | 
[**BillingServicesToggleAutoPaymentCreate**](BillingAPI.md#BillingServicesToggleAutoPaymentCreate) | **Post** /api/billing/services/{id}/toggle-auto-payment/ | 
[**BillingServicesToggleAutoPaymentCreate2**](BillingAPI.md#BillingServicesToggleAutoPaymentCreate2) | **Post** /api/v1/billing/services/{id}/toggle-auto-payment/ | 
[**BillingSubscriptionsList**](BillingAPI.md#BillingSubscriptionsList) | **Get** /api/billing/subscriptions/ | 
[**BillingSubscriptionsList2**](BillingAPI.md#BillingSubscriptionsList2) | **Get** /api/v1/billing/subscriptions/ | 
[**BillingSubscriptionsRetrieve**](BillingAPI.md#BillingSubscriptionsRetrieve) | **Get** /api/billing/subscriptions/{id}/ | 
[**BillingSubscriptionsRetrieve2**](BillingAPI.md#BillingSubscriptionsRetrieve2) | **Get** /api/v1/billing/subscriptions/{id}/ | 



## BillingDepositsCreate

> Deposit BillingDepositsCreate(ctx).DepositCreate(depositCreate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	depositCreate := *openapiclient.NewDepositCreate(int32(123)) // DepositCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingDepositsCreate(context.Background()).DepositCreate(depositCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingDepositsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingDepositsCreate`: Deposit
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingDepositsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingDepositsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositCreate** | [**DepositCreate**](DepositCreate.md) |  | 

### Return type

[**Deposit**](Deposit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingDepositsCreate2

> Deposit BillingDepositsCreate2(ctx).DepositCreate(depositCreate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	depositCreate := *openapiclient.NewDepositCreate(int32(123)) // DepositCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingDepositsCreate2(context.Background()).DepositCreate(depositCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingDepositsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingDepositsCreate2`: Deposit
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingDepositsCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingDepositsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositCreate** | [**DepositCreate**](DepositCreate.md) |  | 

### Return type

[**Deposit**](Deposit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingDepositsList

> PaginatedDepositList BillingDepositsList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingDepositsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingDepositsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingDepositsList`: PaginatedDepositList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingDepositsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingDepositsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDepositList**](PaginatedDepositList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingDepositsList2

> PaginatedDepositList BillingDepositsList2(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingDepositsList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingDepositsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingDepositsList2`: PaginatedDepositList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingDepositsList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingDepositsList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDepositList**](PaginatedDepositList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingDepositsRetrieve

> Deposit BillingDepositsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingDepositsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingDepositsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingDepositsRetrieve`: Deposit
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingDepositsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingDepositsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deposit**](Deposit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingDepositsRetrieve2

> Deposit BillingDepositsRetrieve2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingDepositsRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingDepositsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingDepositsRetrieve2`: Deposit
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingDepositsRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingDepositsRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deposit**](Deposit.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingFundsList

> []FundsBalanceResponse BillingFundsList(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingFundsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingFundsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingFundsList`: []FundsBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingFundsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingFundsListRequest struct via the builder pattern


### Return type

[**[]FundsBalanceResponse**](FundsBalanceResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingFundsList2

> []FundsBalanceResponse BillingFundsList2(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingFundsList2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingFundsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingFundsList2`: []FundsBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingFundsList2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingFundsList2Request struct via the builder pattern


### Return type

[**[]FundsBalanceResponse**](FundsBalanceResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingFundsLogList

> PaginatedFundsLogList BillingFundsLogList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingFundsLogList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingFundsLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingFundsLogList`: PaginatedFundsLogList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingFundsLogList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingFundsLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedFundsLogList**](PaginatedFundsLogList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingFundsLogList2

> PaginatedFundsLogList BillingFundsLogList2(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingFundsLogList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingFundsLogList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingFundsLogList2`: PaginatedFundsLogList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingFundsLogList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingFundsLogList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedFundsLogList**](PaginatedFundsLogList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingFundsNotificationSettingsCreate

> NotificationSettingsResponse BillingFundsNotificationSettingsCreate(ctx).LowBalanceSettings(lowBalanceSettings).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	lowBalanceSettings := *openapiclient.NewLowBalanceSettings(openapiclient.ThresholdTypeEnum("disabled")) // LowBalanceSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingFundsNotificationSettingsCreate(context.Background()).LowBalanceSettings(lowBalanceSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingFundsNotificationSettingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingFundsNotificationSettingsCreate`: NotificationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingFundsNotificationSettingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingFundsNotificationSettingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lowBalanceSettings** | [**LowBalanceSettings**](LowBalanceSettings.md) |  | 

### Return type

[**NotificationSettingsResponse**](NotificationSettingsResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingFundsNotificationSettingsCreate2

> NotificationSettingsResponse BillingFundsNotificationSettingsCreate2(ctx).LowBalanceSettings(lowBalanceSettings).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	lowBalanceSettings := *openapiclient.NewLowBalanceSettings(openapiclient.ThresholdTypeEnum("disabled")) // LowBalanceSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingFundsNotificationSettingsCreate2(context.Background()).LowBalanceSettings(lowBalanceSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingFundsNotificationSettingsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingFundsNotificationSettingsCreate2`: NotificationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingFundsNotificationSettingsCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingFundsNotificationSettingsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lowBalanceSettings** | [**LowBalanceSettings**](LowBalanceSettings.md) |  | 

### Return type

[**NotificationSettingsResponse**](NotificationSettingsResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesList

> PaginatedInvoiceListList BillingInvoicesList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesList`: PaginatedInvoiceListList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedInvoiceListList**](PaginatedInvoiceListList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesList2

> PaginatedInvoiceListList BillingInvoicesList2(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesList2`: PaginatedInvoiceListList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedInvoiceListList**](PaginatedInvoiceListList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesPayWithFundsCreate

> PayWithFundsResponse BillingInvoicesPayWithFundsCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesPayWithFundsCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesPayWithFundsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesPayWithFundsCreate`: PayWithFundsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesPayWithFundsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesPayWithFundsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayWithFundsResponse**](PayWithFundsResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesPayWithFundsCreate2

> PayWithFundsResponse BillingInvoicesPayWithFundsCreate2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesPayWithFundsCreate2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesPayWithFundsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesPayWithFundsCreate2`: PayWithFundsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesPayWithFundsCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesPayWithFundsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayWithFundsResponse**](PayWithFundsResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesPdfRetrieve

> *os.File BillingInvoicesPdfRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesPdfRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesPdfRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesPdfRetrieve`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesPdfRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesPdfRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesPdfRetrieve2

> *os.File BillingInvoicesPdfRetrieve2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesPdfRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesPdfRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesPdfRetrieve2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesPdfRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesPdfRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesRetrieve

> InvoiceDetail BillingInvoicesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesRetrieve`: InvoiceDetail
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceDetail**](InvoiceDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingInvoicesRetrieve2

> InvoiceDetail BillingInvoicesRetrieve2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingInvoicesRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingInvoicesRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingInvoicesRetrieve2`: InvoiceDetail
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingInvoicesRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingInvoicesRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceDetail**](InvoiceDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesCancelCreate

> CancelServiceResponse BillingServicesCancelCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesCancelCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesCancelCreate`: CancelServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesCancelCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelServiceResponse**](CancelServiceResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesCancelCreate2

> CancelServiceResponse BillingServicesCancelCreate2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesCancelCreate2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesCancelCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesCancelCreate2`: CancelServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesCancelCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesCancelCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelServiceResponse**](CancelServiceResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesChangeBillingCycleCreate

> ChangeBillingCycleResponse BillingServicesChangeBillingCycleCreate(ctx, id).ChangeBillingCycle(changeBillingCycle).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	changeBillingCycle := *openapiclient.NewChangeBillingCycle(int32(123)) // ChangeBillingCycle | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesChangeBillingCycleCreate(context.Background(), id).ChangeBillingCycle(changeBillingCycle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesChangeBillingCycleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesChangeBillingCycleCreate`: ChangeBillingCycleResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesChangeBillingCycleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesChangeBillingCycleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeBillingCycle** | [**ChangeBillingCycle**](ChangeBillingCycle.md) |  | 

### Return type

[**ChangeBillingCycleResponse**](ChangeBillingCycleResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesChangeBillingCycleCreate2

> ChangeBillingCycleResponse BillingServicesChangeBillingCycleCreate2(ctx, id).ChangeBillingCycle(changeBillingCycle).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	changeBillingCycle := *openapiclient.NewChangeBillingCycle(int32(123)) // ChangeBillingCycle | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesChangeBillingCycleCreate2(context.Background(), id).ChangeBillingCycle(changeBillingCycle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesChangeBillingCycleCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesChangeBillingCycleCreate2`: ChangeBillingCycleResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesChangeBillingCycleCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesChangeBillingCycleCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeBillingCycle** | [**ChangeBillingCycle**](ChangeBillingCycle.md) |  | 

### Return type

[**ChangeBillingCycleResponse**](ChangeBillingCycleResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesChangeCompanyCreate

> ChangeCompanyResponse BillingServicesChangeCompanyCreate(ctx, id).ChangeCompany(changeCompany).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	changeCompany := *openapiclient.NewChangeCompany() // ChangeCompany |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesChangeCompanyCreate(context.Background(), id).ChangeCompany(changeCompany).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesChangeCompanyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesChangeCompanyCreate`: ChangeCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesChangeCompanyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesChangeCompanyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeCompany** | [**ChangeCompany**](ChangeCompany.md) |  | 

### Return type

[**ChangeCompanyResponse**](ChangeCompanyResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesChangeCompanyCreate2

> ChangeCompanyResponse BillingServicesChangeCompanyCreate2(ctx, id).ChangeCompany(changeCompany).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	changeCompany := *openapiclient.NewChangeCompany() // ChangeCompany |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesChangeCompanyCreate2(context.Background(), id).ChangeCompany(changeCompany).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesChangeCompanyCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesChangeCompanyCreate2`: ChangeCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesChangeCompanyCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesChangeCompanyCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeCompany** | [**ChangeCompany**](ChangeCompany.md) |  | 

### Return type

[**ChangeCompanyResponse**](ChangeCompanyResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesList

> PaginatedServiceListList BillingServicesList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesList`: PaginatedServiceListList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedServiceListList**](PaginatedServiceListList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesList2

> PaginatedServiceListList BillingServicesList2(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesList2`: PaginatedServiceListList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedServiceListList**](PaginatedServiceListList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesRetrieve

> ServiceList BillingServicesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesRetrieve`: ServiceList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceList**](ServiceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesRetrieve2

> ServiceList BillingServicesRetrieve2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesRetrieve2`: ServiceList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceList**](ServiceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesToggleAutoPaymentCreate

> ToggleAutoPaymentResponse BillingServicesToggleAutoPaymentCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesToggleAutoPaymentCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesToggleAutoPaymentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesToggleAutoPaymentCreate`: ToggleAutoPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesToggleAutoPaymentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesToggleAutoPaymentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToggleAutoPaymentResponse**](ToggleAutoPaymentResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingServicesToggleAutoPaymentCreate2

> ToggleAutoPaymentResponse BillingServicesToggleAutoPaymentCreate2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingServicesToggleAutoPaymentCreate2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingServicesToggleAutoPaymentCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingServicesToggleAutoPaymentCreate2`: ToggleAutoPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingServicesToggleAutoPaymentCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingServicesToggleAutoPaymentCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToggleAutoPaymentResponse**](ToggleAutoPaymentResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingSubscriptionsList

> PaginatedSubscriptionList BillingSubscriptionsList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingSubscriptionsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingSubscriptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingSubscriptionsList`: PaginatedSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingSubscriptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingSubscriptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSubscriptionList**](PaginatedSubscriptionList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingSubscriptionsList2

> PaginatedSubscriptionList BillingSubscriptionsList2(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingSubscriptionsList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingSubscriptionsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingSubscriptionsList2`: PaginatedSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingSubscriptionsList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingSubscriptionsList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSubscriptionList**](PaginatedSubscriptionList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingSubscriptionsRetrieve

> Subscription BillingSubscriptionsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingSubscriptionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingSubscriptionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingSubscriptionsRetrieve`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingSubscriptionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingSubscriptionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingSubscriptionsRetrieve2

> Subscription BillingSubscriptionsRetrieve2(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingSubscriptionsRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingSubscriptionsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingSubscriptionsRetrieve2`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingSubscriptionsRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingSubscriptionsRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

