# \DomainAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainDomainCancelCreate**](DomainAPI.md#DomainDomainCancelCreate) | **Post** /api/domain/domain/{domain}/cancel/ | 
[**DomainDomainCheckAvailabilityCreate**](DomainAPI.md#DomainDomainCheckAvailabilityCreate) | **Post** /api/domain/domain/check-availability/ | 
[**DomainDomainContactsCreate**](DomainAPI.md#DomainDomainContactsCreate) | **Post** /api/domain/domain/{domain}/contacts/ | 
[**DomainDomainCreate**](DomainAPI.md#DomainDomainCreate) | **Post** /api/domain/domain/ | 
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

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## DomainDomainContactsCreate

> ContactsUpdateResponse DomainDomainContactsCreate(ctx, domain).ContactsUpdate(contactsUpdate).Execute()





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
	contactsUpdate := *openapiclient.NewContactsUpdate(openapiclient.ContactTypeEnum("registrant"), int32(123)) // ContactsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainContactsCreate(context.Background(), domain).ContactsUpdate(contactsUpdate).Execute()
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

 **contactsUpdate** | [**ContactsUpdate**](ContactsUpdate.md) |  | 

### Return type

[**ContactsUpdateResponse**](ContactsUpdateResponse.md)

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


## DomainDomainNameserversCreate

> NameserversUpdateResponse DomainDomainNameserversCreate(ctx, domain).NameserversUpdate(nameserversUpdate).Execute()





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
	nameserversUpdate := *openapiclient.NewNameserversUpdate([]string{"Nameservers_example"}) // NameserversUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DomainDomainNameserversCreate(context.Background(), domain).NameserversUpdate(nameserversUpdate).Execute()
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

 **nameserversUpdate** | [**NameserversUpdate**](NameserversUpdate.md) |  | 

### Return type

[**NameserversUpdateResponse**](NameserversUpdateResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	domain2 := *openapiclient.NewDomain(int32(123), "Domain_example", *openapiclient.NewTLD(int32(123), "Tld_example", "Price_example", "Registrar_example"), false, time.Now(), time.Now(), *openapiclient.NewService(int32(123), int32(123), "Hostname_example", "Price_example", openapiclient.Status63aEnum("pending"), "Created_example", "Modified_example", time.Now(), time.Now(), NullableInt32(123)), "IdnaName_example", int32(123), "ServiceStatus_example", interface{}(123)) // Domain |  (optional)

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

