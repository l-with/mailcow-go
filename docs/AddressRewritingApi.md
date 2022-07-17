# \AddressRewritingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBCCMap**](AddressRewritingApi.md#CreateBCCMap) | **Post** /api/v1/add/bcc | Create BCC Map
[**CreateRecipientMap**](AddressRewritingApi.md#CreateRecipientMap) | **Post** /api/v1/add/recipient_map | Create Recipient Map
[**DeleteBCCMap**](AddressRewritingApi.md#DeleteBCCMap) | **Post** /api/v1/delete/bcc | Delete BCC Map
[**DeleteRecipientMap**](AddressRewritingApi.md#DeleteRecipientMap) | **Post** /api/v1/delete/recipient_map | Delete Recipient Map
[**GetBCCMap**](AddressRewritingApi.md#GetBCCMap) | **Get** /api/v1/get/bcc/{id} | Get BCC Map
[**GetRecipientMap**](AddressRewritingApi.md#GetRecipientMap) | **Get** /api/v1/get/recipient_map/{id} | Get Recipient Map



## CreateBCCMap

> CreateAlias200ResponseInner CreateBCCMap(ctx).CreateBCCMapRequest(createBCCMapRequest).Execute()

Create BCC Map



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBCCMapRequest := *openapiclient.NewCreateBCCMapRequest() // CreateBCCMapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressRewritingApi.CreateBCCMap(context.Background()).CreateBCCMapRequest(createBCCMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingApi.CreateBCCMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBCCMap`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AddressRewritingApi.CreateBCCMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBCCMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBCCMapRequest** | [**CreateBCCMapRequest**](CreateBCCMapRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRecipientMap

> CreateAlias200ResponseInner CreateRecipientMap(ctx).CreateRecipientMapRequest(createRecipientMapRequest).Execute()

Create Recipient Map



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createRecipientMapRequest := *openapiclient.NewCreateRecipientMapRequest() // CreateRecipientMapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressRewritingApi.CreateRecipientMap(context.Background()).CreateRecipientMapRequest(createRecipientMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingApi.CreateRecipientMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecipientMap`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AddressRewritingApi.CreateRecipientMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecipientMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecipientMapRequest** | [**CreateRecipientMapRequest**](CreateRecipientMapRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBCCMap

> CreateAlias200ResponseInner DeleteBCCMap(ctx).DeleteBCCMapRequest(deleteBCCMapRequest).Execute()

Delete BCC Map



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteBCCMapRequest := *openapiclient.NewDeleteBCCMapRequest() // DeleteBCCMapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressRewritingApi.DeleteBCCMap(context.Background()).DeleteBCCMapRequest(deleteBCCMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingApi.DeleteBCCMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBCCMap`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AddressRewritingApi.DeleteBCCMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBCCMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBCCMapRequest** | [**DeleteBCCMapRequest**](DeleteBCCMapRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecipientMap

> CreateAlias200ResponseInner DeleteRecipientMap(ctx).DeleteRecipientMapRequest(deleteRecipientMapRequest).Execute()

Delete Recipient Map



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteRecipientMapRequest := *openapiclient.NewDeleteRecipientMapRequest() // DeleteRecipientMapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressRewritingApi.DeleteRecipientMap(context.Background()).DeleteRecipientMapRequest(deleteRecipientMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingApi.DeleteRecipientMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRecipientMap`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AddressRewritingApi.DeleteRecipientMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecipientMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRecipientMapRequest** | [**DeleteRecipientMapRequest**](DeleteRecipientMapRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBCCMap

> GetBCCMap(ctx, id).XAPIKey(xAPIKey).Execute()

Get BCC Map



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "all" // string | id of entry you want to get
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressRewritingApi.GetBCCMap(context.Background(), id).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingApi.GetBCCMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of entry you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBCCMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** | e.g. api-key-string | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecipientMap

> GetRecipientMap(ctx, id).XAPIKey(xAPIKey).Execute()

Get Recipient Map



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "all" // string | id of entry you want to get
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressRewritingApi.GetRecipientMap(context.Background(), id).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingApi.GetRecipientMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of entry you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipientMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** | e.g. api-key-string | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

