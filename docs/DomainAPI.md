# \DomainAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainDomainCancelCreate**](DomainAPI.md#DomainDomainCancelCreate) | **Post** /api/domain/domain/{domain}/cancel/ | 
[**DomainDomainCheckAvailabilityCreate**](DomainAPI.md#DomainDomainCheckAvailabilityCreate) | **Post** /api/domain/domain/check-availability/ | 
[**DomainDomainContactsCreate**](DomainAPI.md#DomainDomainContactsCreate) | **Post** /api/domain/domain/{domain}/contacts/ | 
[**DomainDomainCreate**](DomainAPI.md#DomainDomainCreate) | **Post** /api/domain/domain/ | 
[**DomainDomainDnsCreate**](DomainAPI.md#DomainDomainDnsCreate) | **Post** /api/domain/domain/{domain}/dns/ | 
[**DomainDomainDnsDestroy**](DomainAPI.md#DomainDomainDnsDestroy) | **Delete** /api/domain/domain/{domain}/dns/{name}/ | 
[**DomainDomainDnsList**](DomainAPI.md#DomainDomainDnsList) | **Get** /api/domain/domain/{domain}/dns/ | 
[**DomainDomainList**](DomainAPI.md#DomainDomainList) | **Get** /api/domain/domain/ | 
[**DomainDomainNameserversCreate**](DomainAPI.md#DomainDomainNameserversCreate) | **Post** /api/domain/domain/{domain}/nameservers/ | 
[**DomainDomainPartialUpdate**](DomainAPI.md#DomainDomainPartialUpdate) | **Patch** /api/domain/domain/{domain}/ | 
[**DomainDomainRenewCreate**](DomainAPI.md#DomainDomainRenewCreate) | **Post** /api/domain/domain/{domain}/renew/ | 
[**DomainDomainRetrieve**](DomainAPI.md#DomainDomainRetrieve) | **Get** /api/domain/domain/{domain}/ | 
[**DomainDomainTransferRoDomainCreate**](DomainAPI.md#DomainDomainTransferRoDomainCreate) | **Post** /api/domain/domain/transfer-ro-domain/ | 
[**DomainDomainUpdate**](DomainAPI.md#DomainDomainUpdate) | **Put** /api/domain/domain/{domain}/ | 
[**DomainRegistrantsCreate**](DomainAPI.md#DomainRegistrantsCreate) | **Post** /api/domain/registrants/ | 
[**DomainRegistrantsDestroy**](DomainAPI.md#DomainRegistrantsDestroy) | **Delete** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsList**](DomainAPI.md#DomainRegistrantsList) | **Get** /api/domain/registrants/ | 
[**DomainRegistrantsPartialUpdate**](DomainAPI.md#DomainRegistrantsPartialUpdate) | **Patch** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsRetrieve**](DomainAPI.md#DomainRegistrantsRetrieve) | **Get** /api/domain/registrants/{id}/ | 
[**DomainRegistrantsUpdate**](DomainAPI.md#DomainRegistrantsUpdate) | **Put** /api/domain/registrants/{id}/ | 
[**DomainTldList**](DomainAPI.md#DomainTldList) | **Get** /api/domain/tld/ | 
[**DomainTldRetrieve**](DomainAPI.md#DomainTldRetrieve) | **Get** /api/domain/tld/{id}/ | 



## DomainDomainCancelCreate

> DomainCancelResponse DomainDomainCancelCreate(ctx, domain).Execute()





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
	resp, r, err := apiClient.DomainAPI.DomainDomainCancelCreate(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainCancelCreate`: DomainCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainCancelCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainCancelResponse**](DomainCancelResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainCheckAvailabilityCreate

> CheckAvailability DomainDomainCheckAvailabilityCreate(ctx).CheckAvailabilityRequest(checkAvailabilityRequest).Execute()





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
	checkAvailabilityRequest := *openapiclient.NewCheckAvailabilityRequest("Domain_example") // CheckAvailabilityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainCheckAvailabilityCreate(context.Background()).CheckAvailabilityRequest(checkAvailabilityRequest).Execute()
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
 **checkAvailabilityRequest** | [**CheckAvailabilityRequest**](CheckAvailabilityRequest.md) |  | 

### Return type

[**CheckAvailability**](CheckAvailability.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainContactsCreate

> ContactsUpdateResponse DomainDomainContactsCreate(ctx, domain).ContactsUpdateRequest(contactsUpdateRequest).Execute()





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
	contactsUpdateRequest := *openapiclient.NewContactsUpdateRequest(openapiclient.ContactTypeEnum("registrant"), int32(123)) // ContactsUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainContactsCreate(context.Background(), domain).ContactsUpdateRequest(contactsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainContactsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainContactsCreate`: ContactsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainContactsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainContactsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactsUpdateRequest** | [**ContactsUpdateRequest**](ContactsUpdateRequest.md) |  | 

### Return type

[**ContactsUpdateResponse**](ContactsUpdateResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainCreate

> DomainCreate DomainDomainCreate(ctx).DomainCreateRequest(domainCreateRequest).Execute()





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
	domainCreateRequest := *openapiclient.NewDomainCreateRequest("Domain_example") // DomainCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainCreate(context.Background()).DomainCreateRequest(domainCreateRequest).Execute()
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
 **domainCreateRequest** | [**DomainCreateRequest**](DomainCreateRequest.md) |  | 

### Return type

[**DomainCreate**](DomainCreate.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainDnsCreate

> DNSGlue DomainDomainDnsCreate(ctx, domain).DNSGlueRequest(dNSGlueRequest).Execute()





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
	dNSGlueRequest := *openapiclient.NewDNSGlueRequest("Name_example", "Ip_example") // DNSGlueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainDnsCreate(context.Background(), domain).DNSGlueRequest(dNSGlueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainDnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainDnsCreate`: DNSGlue
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainDnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainDnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dNSGlueRequest** | [**DNSGlueRequest**](DNSGlueRequest.md) |  | 

### Return type

[**DNSGlue**](DNSGlue.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainDnsDestroy

> DomainDomainDnsDestroy(ctx, domain, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainAPI.DomainDomainDnsDestroy(context.Background(), domain, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainDnsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainDnsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainDnsList

> PaginatedDNSGlueList DomainDomainDnsList(ctx, domain).Page(page).Execute()





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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainDnsList(context.Background(), domain).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainDnsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainDnsList`: PaginatedDNSGlueList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainDnsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainDnsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDNSGlueList**](PaginatedDNSGlueList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainNameserversCreate

> NameserversUpdateResponse DomainDomainNameserversCreate(ctx, domain).NameserversUpdateRequest(nameserversUpdateRequest).Execute()





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
	nameserversUpdateRequest := *openapiclient.NewNameserversUpdateRequest([]string{"Nameservers_example"}) // NameserversUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainNameserversCreate(context.Background(), domain).NameserversUpdateRequest(nameserversUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDomainNameserversCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDomainNameserversCreate`: NameserversUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainDomainNameserversCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDomainNameserversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nameserversUpdateRequest** | [**NameserversUpdateRequest**](NameserversUpdateRequest.md) |  | 

### Return type

[**NameserversUpdateResponse**](NameserversUpdateResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainPartialUpdate

> Domain DomainDomainPartialUpdate(ctx, domain).PatchedDomainRequest(patchedDomainRequest).Execute()





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
	patchedDomainRequest := *openapiclient.NewPatchedDomainRequest() // PatchedDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainPartialUpdate(context.Background(), domain).PatchedDomainRequest(patchedDomainRequest).Execute()
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

 **patchedDomainRequest** | [**PatchedDomainRequest**](PatchedDomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainRenewCreate

> RenewDomain DomainDomainRenewCreate(ctx, domain).RenewDomainRequest(renewDomainRequest).Execute()





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
	renewDomainRequest := *openapiclient.NewRenewDomainRequest(int32(123)) // RenewDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainRenewCreate(context.Background(), domain).RenewDomainRequest(renewDomainRequest).Execute()
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

 **renewDomainRequest** | [**RenewDomainRequest**](RenewDomainRequest.md) |  | 

### Return type

[**RenewDomain**](RenewDomain.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainTransferRoDomainCreate

> TransferRoDomain DomainDomainTransferRoDomainCreate(ctx).TransferRoDomainRequest(transferRoDomainRequest).Execute()





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
	transferRoDomainRequest := *openapiclient.NewTransferRoDomainRequest("Domain_example", "AuthCode_example") // TransferRoDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainTransferRoDomainCreate(context.Background()).TransferRoDomainRequest(transferRoDomainRequest).Execute()
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
 **transferRoDomainRequest** | [**TransferRoDomainRequest**](TransferRoDomainRequest.md) |  | 

### Return type

[**TransferRoDomain**](TransferRoDomain.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDomainUpdate

> Domain DomainDomainUpdate(ctx, domain).DomainRequest(domainRequest).Execute()





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
	domainRequest := *openapiclient.NewDomainRequest() // DomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainUpdate(context.Background(), domain).DomainRequest(domainRequest).Execute()
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

 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsCreate

> DomainRegistrant DomainRegistrantsCreate(ctx).DomainRegistrantRequest(domainRegistrantRequest).Execute()





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
	domainRegistrantRequest := *openapiclient.NewDomainRegistrantRequest("FirstName_example", "LastName_example", "Address_example", "City_example", "Region_example", "PostalCode_example", openapiclient.CountryEnum("AF"), "Email_example", "Phone_example") // DomainRegistrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsCreate(context.Background()).DomainRegistrantRequest(domainRegistrantRequest).Execute()
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
 **domainRegistrantRequest** | [**DomainRegistrantRequest**](DomainRegistrantRequest.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsPartialUpdate

> DomainRegistrant DomainRegistrantsPartialUpdate(ctx, id).PatchedDomainRegistrantRequest(patchedDomainRegistrantRequest).Execute()





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
	patchedDomainRegistrantRequest := *openapiclient.NewPatchedDomainRegistrantRequest() // PatchedDomainRegistrantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsPartialUpdate(context.Background(), id).PatchedDomainRegistrantRequest(patchedDomainRegistrantRequest).Execute()
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

 **patchedDomainRegistrantRequest** | [**PatchedDomainRegistrantRequest**](PatchedDomainRegistrantRequest.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRegistrantsUpdate

> DomainRegistrant DomainRegistrantsUpdate(ctx, id).DomainRegistrantRequest(domainRegistrantRequest).Execute()





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
	domainRegistrantRequest := *openapiclient.NewDomainRegistrantRequest("FirstName_example", "LastName_example", "Address_example", "City_example", "Region_example", "PostalCode_example", openapiclient.CountryEnum("AF"), "Email_example", "Phone_example") // DomainRegistrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainRegistrantsUpdate(context.Background(), id).DomainRegistrantRequest(domainRegistrantRequest).Execute()
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

 **domainRegistrantRequest** | [**DomainRegistrantRequest**](DomainRegistrantRequest.md) |  | 

### Return type

[**DomainRegistrant**](DomainRegistrant.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

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

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

