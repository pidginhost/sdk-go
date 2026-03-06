# \SupportAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupportDepartmentsList**](SupportAPI.md#SupportDepartmentsList) | **Get** /api/support/departments/ | 
[**SupportDepartmentsList2**](SupportAPI.md#SupportDepartmentsList2) | **Get** /api/v1/support/departments/ | 
[**SupportTicketsCloseCreate**](SupportAPI.md#SupportTicketsCloseCreate) | **Post** /api/support/tickets/{id}/close/ | 
[**SupportTicketsCloseCreate2**](SupportAPI.md#SupportTicketsCloseCreate2) | **Post** /api/v1/support/tickets/{id}/close/ | 
[**SupportTicketsCreate**](SupportAPI.md#SupportTicketsCreate) | **Post** /api/support/tickets/ | 
[**SupportTicketsCreate2**](SupportAPI.md#SupportTicketsCreate2) | **Post** /api/v1/support/tickets/ | 
[**SupportTicketsList**](SupportAPI.md#SupportTicketsList) | **Get** /api/support/tickets/ | 
[**SupportTicketsList2**](SupportAPI.md#SupportTicketsList2) | **Get** /api/v1/support/tickets/ | 
[**SupportTicketsMessagesAttachmentRetrieve**](SupportAPI.md#SupportTicketsMessagesAttachmentRetrieve) | **Get** /api/support/tickets/{id}/messages/{message_id}/attachment/ | 
[**SupportTicketsMessagesAttachmentRetrieve2**](SupportAPI.md#SupportTicketsMessagesAttachmentRetrieve2) | **Get** /api/v1/support/tickets/{id}/messages/{message_id}/attachment/ | 
[**SupportTicketsReopenCreate**](SupportAPI.md#SupportTicketsReopenCreate) | **Post** /api/support/tickets/{id}/reopen/ | 
[**SupportTicketsReopenCreate2**](SupportAPI.md#SupportTicketsReopenCreate2) | **Post** /api/v1/support/tickets/{id}/reopen/ | 
[**SupportTicketsReplyCreate**](SupportAPI.md#SupportTicketsReplyCreate) | **Post** /api/support/tickets/{id}/reply/ | 
[**SupportTicketsReplyCreate2**](SupportAPI.md#SupportTicketsReplyCreate2) | **Post** /api/v1/support/tickets/{id}/reply/ | 
[**SupportTicketsRetrieve**](SupportAPI.md#SupportTicketsRetrieve) | **Get** /api/support/tickets/{id}/ | 
[**SupportTicketsRetrieve2**](SupportAPI.md#SupportTicketsRetrieve2) | **Get** /api/v1/support/tickets/{id}/ | 



## SupportDepartmentsList

> []Department SupportDepartmentsList(ctx).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportDepartmentsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportDepartmentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportDepartmentsList`: []Department
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportDepartmentsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSupportDepartmentsListRequest struct via the builder pattern


### Return type

[**[]Department**](Department.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportDepartmentsList2

> []Department SupportDepartmentsList2(ctx).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportDepartmentsList2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportDepartmentsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportDepartmentsList2`: []Department
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportDepartmentsList2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSupportDepartmentsList2Request struct via the builder pattern


### Return type

[**[]Department**](Department.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsCloseCreate

> TicketCloseResponse SupportTicketsCloseCreate(ctx, id).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsCloseCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsCloseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsCloseCreate`: TicketCloseResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsCloseCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsCloseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketCloseResponse**](TicketCloseResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsCloseCreate2

> TicketCloseResponse SupportTicketsCloseCreate2(ctx, id).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsCloseCreate2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsCloseCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsCloseCreate2`: TicketCloseResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsCloseCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsCloseCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketCloseResponse**](TicketCloseResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsCreate

> TicketDetail SupportTicketsCreate(ctx).TicketCreate(ticketCreate).Execute()





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
	ticketCreate := *openapiclient.NewTicketCreate("Subject_example", int32(123), "Message_example") // TicketCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.SupportTicketsCreate(context.Background()).TicketCreate(ticketCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsCreate`: TicketDetail
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ticketCreate** | [**TicketCreate**](TicketCreate.md) |  | 

### Return type

[**TicketDetail**](TicketDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsCreate2

> TicketDetail SupportTicketsCreate2(ctx).TicketCreate(ticketCreate).Execute()





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
	ticketCreate := *openapiclient.NewTicketCreate("Subject_example", int32(123), "Message_example") // TicketCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.SupportTicketsCreate2(context.Background()).TicketCreate(ticketCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsCreate2`: TicketDetail
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsCreate2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ticketCreate** | [**TicketCreate**](TicketCreate.md) |  | 

### Return type

[**TicketDetail**](TicketDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsList

> PaginatedTicketListList SupportTicketsList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsList`: PaginatedTicketListList
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedTicketListList**](PaginatedTicketListList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsList2

> PaginatedTicketListList SupportTicketsList2(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsList2`: PaginatedTicketListList
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedTicketListList**](PaginatedTicketListList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsMessagesAttachmentRetrieve

> *os.File SupportTicketsMessagesAttachmentRetrieve(ctx, id, messageId).Execute()





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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.SupportTicketsMessagesAttachmentRetrieve(context.Background(), id, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsMessagesAttachmentRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsMessagesAttachmentRetrieve`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsMessagesAttachmentRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsMessagesAttachmentRetrieveRequest struct via the builder pattern


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


## SupportTicketsMessagesAttachmentRetrieve2

> *os.File SupportTicketsMessagesAttachmentRetrieve2(ctx, id, messageId).Execute()





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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.SupportTicketsMessagesAttachmentRetrieve2(context.Background(), id, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsMessagesAttachmentRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsMessagesAttachmentRetrieve2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsMessagesAttachmentRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsMessagesAttachmentRetrieve2Request struct via the builder pattern


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


## SupportTicketsReopenCreate

> TicketReopenResponse SupportTicketsReopenCreate(ctx, id).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsReopenCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsReopenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsReopenCreate`: TicketReopenResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsReopenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsReopenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketReopenResponse**](TicketReopenResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsReopenCreate2

> TicketReopenResponse SupportTicketsReopenCreate2(ctx, id).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsReopenCreate2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsReopenCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsReopenCreate2`: TicketReopenResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsReopenCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsReopenCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketReopenResponse**](TicketReopenResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsReplyCreate

> TicketReplyResponse SupportTicketsReplyCreate(ctx, id).TicketReply(ticketReply).Execute()





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
	ticketReply := *openapiclient.NewTicketReply("Message_example") // TicketReply | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.SupportTicketsReplyCreate(context.Background(), id).TicketReply(ticketReply).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsReplyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsReplyCreate`: TicketReplyResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsReplyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsReplyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ticketReply** | [**TicketReply**](TicketReply.md) |  | 

### Return type

[**TicketReplyResponse**](TicketReplyResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsReplyCreate2

> TicketReplyResponse SupportTicketsReplyCreate2(ctx, id).TicketReply(ticketReply).Execute()





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
	ticketReply := *openapiclient.NewTicketReply("Message_example") // TicketReply | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportAPI.SupportTicketsReplyCreate2(context.Background(), id).TicketReply(ticketReply).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsReplyCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsReplyCreate2`: TicketReplyResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsReplyCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsReplyCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ticketReply** | [**TicketReply**](TicketReply.md) |  | 

### Return type

[**TicketReplyResponse**](TicketReplyResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsRetrieve

> TicketDetail SupportTicketsRetrieve(ctx, id).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsRetrieve`: TicketDetail
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketDetail**](TicketDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportTicketsRetrieve2

> TicketDetail SupportTicketsRetrieve2(ctx, id).Execute()





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
	resp, r, err := apiClient.SupportAPI.SupportTicketsRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.SupportTicketsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupportTicketsRetrieve2`: TicketDetail
	fmt.Fprintf(os.Stdout, "Response from `SupportAPI.SupportTicketsRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportTicketsRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketDetail**](TicketDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

