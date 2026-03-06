# \HostingAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostingHostingChangePasswordCreate**](HostingAPI.md#HostingHostingChangePasswordCreate) | **Post** /api/hosting/hosting/{id}/change-password/ | 
[**HostingHostingChangePasswordCreate2**](HostingAPI.md#HostingHostingChangePasswordCreate2) | **Post** /api/v1/hosting/hosting/{id}/change-password/ | 
[**HostingHostingList**](HostingAPI.md#HostingHostingList) | **Get** /api/hosting/hosting/ | 
[**HostingHostingList2**](HostingAPI.md#HostingHostingList2) | **Get** /api/v1/hosting/hosting/ | 
[**HostingHostingRetrieve**](HostingAPI.md#HostingHostingRetrieve) | **Get** /api/hosting/hosting/{id}/ | 
[**HostingHostingRetrieve2**](HostingAPI.md#HostingHostingRetrieve2) | **Get** /api/v1/hosting/hosting/{id}/ | 



## HostingHostingChangePasswordCreate

> HostingChangePasswordResponse HostingHostingChangePasswordCreate(ctx, id).ChangePassword(changePassword).Execute()





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
	changePassword := *openapiclient.NewChangePassword("Password_example") // ChangePassword | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostingAPI.HostingHostingChangePasswordCreate(context.Background(), id).ChangePassword(changePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.HostingHostingChangePasswordCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostingHostingChangePasswordCreate`: HostingChangePasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.HostingHostingChangePasswordCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostingHostingChangePasswordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changePassword** | [**ChangePassword**](ChangePassword.md) |  | 

### Return type

[**HostingChangePasswordResponse**](HostingChangePasswordResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostingHostingChangePasswordCreate2

> HostingChangePasswordResponse HostingHostingChangePasswordCreate2(ctx, id).ChangePassword(changePassword).Execute()





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
	changePassword := *openapiclient.NewChangePassword("Password_example") // ChangePassword | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostingAPI.HostingHostingChangePasswordCreate2(context.Background(), id).ChangePassword(changePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.HostingHostingChangePasswordCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostingHostingChangePasswordCreate2`: HostingChangePasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.HostingHostingChangePasswordCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostingHostingChangePasswordCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changePassword** | [**ChangePassword**](ChangePassword.md) |  | 

### Return type

[**HostingChangePasswordResponse**](HostingChangePasswordResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostingHostingList

> PaginatedHostingServiceList HostingHostingList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.HostingAPI.HostingHostingList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.HostingHostingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostingHostingList`: PaginatedHostingServiceList
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.HostingHostingList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostingHostingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedHostingServiceList**](PaginatedHostingServiceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostingHostingList2

> PaginatedHostingServiceList HostingHostingList2(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.HostingAPI.HostingHostingList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.HostingHostingList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostingHostingList2`: PaginatedHostingServiceList
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.HostingHostingList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostingHostingList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedHostingServiceList**](PaginatedHostingServiceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostingHostingRetrieve

> HostingService HostingHostingRetrieve(ctx, id).Execute()





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
	resp, r, err := apiClient.HostingAPI.HostingHostingRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.HostingHostingRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostingHostingRetrieve`: HostingService
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.HostingHostingRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostingHostingRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostingService**](HostingService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostingHostingRetrieve2

> HostingService HostingHostingRetrieve2(ctx, id).Execute()





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
	resp, r, err := apiClient.HostingAPI.HostingHostingRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.HostingHostingRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostingHostingRetrieve2`: HostingService
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.HostingHostingRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostingHostingRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostingService**](HostingService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

