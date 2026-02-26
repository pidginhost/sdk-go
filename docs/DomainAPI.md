# \DomainAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainDomainCheckAvailabilityCreate**](DomainAPI.md#DomainDomainCheckAvailabilityCreate) | **Post** /api/domain/domain/check-availability/ | 
[**DomainDomainCheckAvailabilityCreate2**](DomainAPI.md#DomainDomainCheckAvailabilityCreate2) | **Post** /api/v1/domain/domain/check-availability/ | 
[**DomainDomainCreate**](DomainAPI.md#DomainDomainCreate) | **Post** /api/domain/domain/ | 
[**DomainDomainCreate2**](DomainAPI.md#DomainDomainCreate2) | **Post** /api/v1/domain/domain/ | 
[**DomainDomainList**](DomainAPI.md#DomainDomainList) | **Get** /api/domain/domain/ | 
[**DomainDomainList2**](DomainAPI.md#DomainDomainList2) | **Get** /api/v1/domain/domain/ | 
[**DomainDomainPartialUpdate**](DomainAPI.md#DomainDomainPartialUpdate) | **Patch** /api/domain/domain/{domain}/ | 
[**DomainDomainPartialUpdate2**](DomainAPI.md#DomainDomainPartialUpdate2) | **Patch** /api/v1/domain/domain/{domain}/ | 
[**DomainDomainRenewCreate**](DomainAPI.md#DomainDomainRenewCreate) | **Post** /api/domain/domain/{domain}/renew/ | 
[**DomainDomainRenewCreate2**](DomainAPI.md#DomainDomainRenewCreate2) | **Post** /api/v1/domain/domain/{domain}/renew/ | 
[**DomainDomainRetrieve**](DomainAPI.md#DomainDomainRetrieve) | **Get** /api/domain/domain/{domain}/ | 
[**DomainDomainRetrieve2**](DomainAPI.md#DomainDomainRetrieve2) | **Get** /api/v1/domain/domain/{domain}/ | 
[**DomainDomainTransferRoDomainCreate**](DomainAPI.md#DomainDomainTransferRoDomainCreate) | **Post** /api/domain/domain/transfer-ro-domain/ | 
[**DomainDomainTransferRoDomainCreate2**](DomainAPI.md#DomainDomainTransferRoDomainCreate2) | **Post** /api/v1/domain/domain/transfer-ro-domain/ | 
[**DomainDomainUpdate**](DomainAPI.md#DomainDomainUpdate) | **Put** /api/domain/domain/{domain}/ | 
[**DomainDomainUpdate2**](DomainAPI.md#DomainDomainUpdate2) | **Put** /api/v1/domain/domain/{domain}/ | 
[**DomainRegistrantsCreate**](DomainAPI.md#DomainRegistrantsCreate) | **Post** /api/domain/registrants/ | 
[**DomainRegistrantsCreate2**](DomainAPI.md#DomainRegistrantsCreate2) | **Post** /api/v1/domain/registrants/ | 
[**DomainRegistrantsDestroy**](DomainAPI.md#DomainRegistrantsDestroy) | **Delete** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsDestroy2**](DomainAPI.md#DomainRegistrantsDestroy2) | **Delete** /api/v1/domain/registrants/{id}/ | 
[**DomainRegistrantsList**](DomainAPI.md#DomainRegistrantsList) | **Get** /api/domain/registrants/ | 
[**DomainRegistrantsList2**](DomainAPI.md#DomainRegistrantsList2) | **Get** /api/v1/domain/registrants/ | 
[**DomainRegistrantsPartialUpdate**](DomainAPI.md#DomainRegistrantsPartialUpdate) | **Patch** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsPartialUpdate2**](DomainAPI.md#DomainRegistrantsPartialUpdate2) | **Patch** /api/v1/domain/registrants/{id}/ | 
[**DomainRegistrantsRetrieve**](DomainAPI.md#DomainRegistrantsRetrieve) | **Get** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsRetrieve2**](DomainAPI.md#DomainRegistrantsRetrieve2) | **Get** /api/v1/domain/registrants/{id}/ | 
[**DomainRegistrantsUpdate**](DomainAPI.md#DomainRegistrantsUpdate) | **Put** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsUpdate2**](DomainAPI.md#DomainRegistrantsUpdate2) | **Put** /api/v1/domain/registrants/{id}/ | 
[**DomainTldList**](DomainAPI.md#DomainTldList) | **Get** /api/domain/tld/ | 
[**DomainTldList2**](DomainAPI.md#DomainTldList2) | **Get** /api/v1/domain/tld/ | 
[**DomainTldRetrieve**](DomainAPI.md#DomainTldRetrieve) | **Get** /api/domain/tld/{id}/ | 
[**DomainTldRetrieve2**](DomainAPI.md#DomainTldRetrieve2) | **Get** /api/v1/domain/tld/{id}/ | 



## DomainDomainCheckAvailabilityCreate

> CheckAvailability DomainDomainCheckAvailabilityCreate(ctx).CheckAvailability(checkAvailability).Execute()





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
	checkAvailability := *openapiclient.NewCheckAvailability("Domain_example") // CheckAvailability | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainCheckAvailabilityCreate(context.Background()).CheckAvailability(checkAvailability).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainCheckAvailabilityCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainCheckAvailabilityCreate`: CheckAvailability
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainCheckAvailabilityCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainCheckAvailabilityCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAvailability** | [**CheckAvailability**](CheckAvailability.md) |  | 

### Return type

[**CheckAvailability**](CheckAvailability.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainCheckAvailabilityCreate2

> CheckAvailability DomainDomainCheckAvailabilityCreate2(ctx).CheckAvailability(checkAvailability).Execute()





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
	checkAvailability := *openapiclient.NewCheckAvailability("Domain_example") // CheckAvailability | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainCheckAvailabilityCreate2(context.Background()).CheckAvailability(checkAvailability).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainCheckAvailabilityCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainCheckAvailabilityCreate2`: CheckAvailability
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainCheckAvailabilityCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainCheckAvailabilityCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAvailability** | [**CheckAvailability**](CheckAvailability.md) |  | 

### Return type

[**CheckAvailability**](CheckAvailability.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainCreate

> DomainCreate DomainDomainCreate(ctx).DomainCreate(domainCreate).Execute()





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
	domainCreate := *openapiclient.NewDomainCreate("Domain_example") // DomainCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainCreate(context.Background()).DomainCreate(domainCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainCreate`: DomainCreate
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainCreate** | [**DomainCreate**](DomainCreate.md) |  | 

### Return type

[**DomainCreate**](DomainCreate.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainCreate2

> DomainCreate DomainDomainCreate2(ctx).DomainCreate(domainCreate).Execute()





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
	domainCreate := *openapiclient.NewDomainCreate("Domain_example") // DomainCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainCreate2(context.Background()).DomainCreate(domainCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainCreate2`: DomainCreate
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainCreate** | [**DomainCreate**](DomainCreate.md) |  | 

### Return type

[**DomainCreate**](DomainCreate.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainList

> PaginatedDomainList DomainDomainList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainDomainList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainList`: PaginatedDomainList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDomainList**](PaginatedDomainList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainList2

> PaginatedDomainList DomainDomainList2(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainDomainList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainList2`: PaginatedDomainList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDomainList**](PaginatedDomainList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainPartialUpdate

> Domain DomainDomainPartialUpdate(ctx, domain).PatchedDomain(patchedDomain).Execute()





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
	domain := "domain_example" // string | 
	patchedDomain := *openapiclient.NewPatchedDomain() // PatchedDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainPartialUpdate(context.Background(), domain).PatchedDomain(patchedDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainPartialUpdate`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDomain** | [**PatchedDomain**](PatchedDomain.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainPartialUpdate2

> Domain DomainDomainPartialUpdate2(ctx, domain).PatchedDomain(patchedDomain).Execute()





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
	domain := "domain_example" // string | 
	patchedDomain := *openapiclient.NewPatchedDomain() // PatchedDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainPartialUpdate2(context.Background(), domain).PatchedDomain(patchedDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainPartialUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainPartialUpdate2`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainPartialUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainPartialUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDomain** | [**PatchedDomain**](PatchedDomain.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainRenewCreate

> RenewDomain DomainDomainRenewCreate(ctx, domain).RenewDomain(renewDomain).Execute()





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
	domain := "domain_example" // string | 
	renewDomain := *openapiclient.NewRenewDomain(int32(123)) // RenewDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainRenewCreate(context.Background(), domain).RenewDomain(renewDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainRenewCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainRenewCreate`: RenewDomain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainRenewCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainRenewCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewDomain** | [**RenewDomain**](RenewDomain.md) |  | 

### Return type

[**RenewDomain**](RenewDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainRenewCreate2

> RenewDomain DomainDomainRenewCreate2(ctx, domain).RenewDomain(renewDomain).Execute()





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
	domain := "domain_example" // string | 
	renewDomain := *openapiclient.NewRenewDomain(int32(123)) // RenewDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainRenewCreate2(context.Background(), domain).RenewDomain(renewDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainRenewCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainRenewCreate2`: RenewDomain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainRenewCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainRenewCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewDomain** | [**RenewDomain**](RenewDomain.md) |  | 

### Return type

[**RenewDomain**](RenewDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainRetrieve

> Domain DomainDomainRetrieve(ctx, domain).Execute()





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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainRetrieve(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainRetrieve`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainRetrieve2

> Domain DomainDomainRetrieve2(ctx, domain).Execute()





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
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainRetrieve2(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainRetrieve2`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainTransferRoDomainCreate

> TransferRoDomain DomainDomainTransferRoDomainCreate(ctx).TransferRoDomain(transferRoDomain).Execute()





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
	transferRoDomain := *openapiclient.NewTransferRoDomain("Domain_example", "AuthCode_example") // TransferRoDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainTransferRoDomainCreate(context.Background()).TransferRoDomain(transferRoDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainTransferRoDomainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainTransferRoDomainCreate`: TransferRoDomain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainTransferRoDomainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainTransferRoDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferRoDomain** | [**TransferRoDomain**](TransferRoDomain.md) |  | 

### Return type

[**TransferRoDomain**](TransferRoDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainTransferRoDomainCreate2

> TransferRoDomain DomainDomainTransferRoDomainCreate2(ctx).TransferRoDomain(transferRoDomain).Execute()





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
	transferRoDomain := *openapiclient.NewTransferRoDomain("Domain_example", "AuthCode_example") // TransferRoDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainTransferRoDomainCreate2(context.Background()).TransferRoDomain(transferRoDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainTransferRoDomainCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainTransferRoDomainCreate2`: TransferRoDomain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainTransferRoDomainCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainTransferRoDomainCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferRoDomain** | [**TransferRoDomain**](TransferRoDomain.md) |  | 

### Return type

[**TransferRoDomain**](TransferRoDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainUpdate

> Domain DomainDomainUpdate(ctx, domain).Domain2(domain2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	domain := "domain_example" // string | 
	domain2 := *openapiclient.NewDomain(int32(123), "Domain_example", *openapiclient.NewTLD(int32(123), "Tld_example", "Price_example", "Registrar_example"), false, time.Now(), time.Now(), *openapiclient.NewService(int32(123), int32(123), "Hostname_example", "Price_example", openapiclient.ServiceStatusEnum("pending"), time.Now(), time.Now(), time.Now(), time.Now(), NullableInt32(123)), "IdnaName_example", int32(123), "ServiceStatus_example", interface{}(123)) // Domain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainUpdate(context.Background(), domain).Domain2(domain2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainUpdate`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain2** | [**Domain**](Domain.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainUpdate2

> Domain DomainDomainUpdate2(ctx, domain).Domain2(domain2).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	domain := "domain_example" // string | 
	domain2 := *openapiclient.NewDomain(int32(123), "Domain_example", *openapiclient.NewTLD(int32(123), "Tld_example", "Price_example", "Registrar_example"), false, time.Now(), time.Now(), *openapiclient.NewService(int32(123), int32(123), "Hostname_example", "Price_example", openapiclient.ServiceStatusEnum("pending"), time.Now(), time.Now(), time.Now(), time.Now(), NullableInt32(123)), "IdnaName_example", int32(123), "ServiceStatus_example", interface{}(123)) // Domain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainUpdate2(context.Background(), domain).Domain2(domain2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainUpdate2`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain2** | [**Domain**](Domain.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsCreate

> DomainRegistrant DomainRegistrantsCreate(ctx).DomainRegistrant(domainRegistrant).Execute()





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
	domainRegistrant := *openapiclient.NewDomainRegistrant(int32(123), "FirstName_example", "LastName_example", "Address_example", "City_example", "Region_example", "PostalCode_example", openapiclient.CountryEnum("AF"), "Email_example", "Phone_example") // DomainRegistrant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsCreate(context.Background()).DomainRegistrant(domainRegistrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsCreate`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainRegistrant** | [**DomainRegistrant**](DomainRegistrant.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsCreate2

> DomainRegistrant DomainRegistrantsCreate2(ctx).DomainRegistrant(domainRegistrant).Execute()





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
	domainRegistrant := *openapiclient.NewDomainRegistrant(int32(123), "FirstName_example", "LastName_example", "Address_example", "City_example", "Region_example", "PostalCode_example", openapiclient.CountryEnum("AF"), "Email_example", "Phone_example") // DomainRegistrant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsCreate2(context.Background()).DomainRegistrant(domainRegistrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsCreate2`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainRegistrant** | [**DomainRegistrant**](DomainRegistrant.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsDestroy

> DomainRegistrantsDestroy(ctx, id).Execute()





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
	r, err := apiClient.DomainAPI.DomainRegistrantsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsDestroy2

> DomainRegistrantsDestroy2(ctx, id).Execute()





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
	r, err := apiClient.DomainAPI.DomainRegistrantsDestroy2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsDestroy2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsDestroy2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsList

> PaginatedDomainRegistrantList DomainRegistrantsList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsList`: PaginatedDomainRegistrantList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDomainRegistrantList**](PaginatedDomainRegistrantList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsList2

> PaginatedDomainRegistrantList DomainRegistrantsList2(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsList2`: PaginatedDomainRegistrantList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDomainRegistrantList**](PaginatedDomainRegistrantList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsPartialUpdate

> DomainRegistrant DomainRegistrantsPartialUpdate(ctx, id).PatchedDomainRegistrant(patchedDomainRegistrant).Execute()





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
	patchedDomainRegistrant := *openapiclient.NewPatchedDomainRegistrant() // PatchedDomainRegistrant |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsPartialUpdate(context.Background(), id).PatchedDomainRegistrant(patchedDomainRegistrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsPartialUpdate`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDomainRegistrant** | [**PatchedDomainRegistrant**](PatchedDomainRegistrant.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsPartialUpdate2

> DomainRegistrant DomainRegistrantsPartialUpdate2(ctx, id).PatchedDomainRegistrant(patchedDomainRegistrant).Execute()





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
	patchedDomainRegistrant := *openapiclient.NewPatchedDomainRegistrant() // PatchedDomainRegistrant |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsPartialUpdate2(context.Background(), id).PatchedDomainRegistrant(patchedDomainRegistrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsPartialUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsPartialUpdate2`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsPartialUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsPartialUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDomainRegistrant** | [**PatchedDomainRegistrant**](PatchedDomainRegistrant.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsRetrieve

> DomainRegistrant DomainRegistrantsRetrieve(ctx, id).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsRetrieve`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsRetrieve2

> DomainRegistrant DomainRegistrantsRetrieve2(ctx, id).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsRetrieve2`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsUpdate

> DomainRegistrant DomainRegistrantsUpdate(ctx, id).DomainRegistrant(domainRegistrant).Execute()





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
	domainRegistrant := *openapiclient.NewDomainRegistrant(int32(123), "FirstName_example", "LastName_example", "Address_example", "City_example", "Region_example", "PostalCode_example", openapiclient.CountryEnum("AF"), "Email_example", "Phone_example") // DomainRegistrant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsUpdate(context.Background(), id).DomainRegistrant(domainRegistrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsUpdate`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainRegistrant** | [**DomainRegistrant**](DomainRegistrant.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsUpdate2

> DomainRegistrant DomainRegistrantsUpdate2(ctx, id).DomainRegistrant(domainRegistrant).Execute()





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
	domainRegistrant := *openapiclient.NewDomainRegistrant(int32(123), "FirstName_example", "LastName_example", "Address_example", "City_example", "Region_example", "PostalCode_example", openapiclient.CountryEnum("AF"), "Email_example", "Phone_example") // DomainRegistrant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsUpdate2(context.Background(), id).DomainRegistrant(domainRegistrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRegistrantsUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainRegistrantsUpdate2`: DomainRegistrant
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRegistrantsUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainRegistrantsUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainRegistrant** | [**DomainRegistrant**](DomainRegistrant.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainTldList

> PaginatedTLDList DomainTldList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainTldList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainTldList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainTldList`: PaginatedTLDList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainTldList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainTldListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedTLDList**](PaginatedTLDList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainTldList2

> PaginatedTLDList DomainTldList2(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainTldList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainTldList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainTldList2`: PaginatedTLDList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainTldList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainTldList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedTLDList**](PaginatedTLDList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainTldRetrieve

> TLD DomainTldRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this top level domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainTldRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainTldRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainTldRetrieve`: TLD
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainTldRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this top level domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainTldRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TLD**](TLD.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainTldRetrieve2

> TLD DomainTldRetrieve2(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this top level domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainTldRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainTldRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainTldRetrieve2`: TLD
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainTldRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this top level domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainTldRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TLD**](TLD.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

