# \RoutingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSenderDependentTransports**](RoutingApi.md#CreateSenderDependentTransports) | **Post** /api/v1/add/relayhost | Create Sender-Dependent Transports
[**CreateTransportMaps**](RoutingApi.md#CreateTransportMaps) | **Post** /api/v1/add/transport | Create Transport Maps
[**DeleteSenderDependentTransports**](RoutingApi.md#DeleteSenderDependentTransports) | **Post** /api/v1/delete/relayhost | Delete Sender-Dependent Transports
[**DeleteTransportMaps**](RoutingApi.md#DeleteTransportMaps) | **Post** /api/v1/delete/transport | Delete Transport Maps
[**GetSenderDependentTransports**](RoutingApi.md#GetSenderDependentTransports) | **Get** /api/v1/get/relayhost/{id} | Get Sender-Dependent Transports
[**GetTransportMaps**](RoutingApi.md#GetTransportMaps) | **Get** /api/v1/get/transport/{id} | Get Transport Maps



## CreateSenderDependentTransports

> CreateAlias200Response CreateSenderDependentTransports(ctx).CreateSenderDependentTransportsRequest(createSenderDependentTransportsRequest).Execute()

Create Sender-Dependent Transports



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
    createSenderDependentTransportsRequest := *openapiclient.NewCreateSenderDependentTransportsRequest() // CreateSenderDependentTransportsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateSenderDependentTransports(context.Background()).CreateSenderDependentTransportsRequest(createSenderDependentTransportsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateSenderDependentTransports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSenderDependentTransports`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateSenderDependentTransports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSenderDependentTransportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSenderDependentTransportsRequest** | [**CreateSenderDependentTransportsRequest**](CreateSenderDependentTransportsRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransportMaps

> CreateAlias200Response CreateTransportMaps(ctx).CreateTransportMapsRequest(createTransportMapsRequest).Execute()

Create Transport Maps



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
    createTransportMapsRequest := *openapiclient.NewCreateTransportMapsRequest() // CreateTransportMapsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateTransportMaps(context.Background()).CreateTransportMapsRequest(createTransportMapsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateTransportMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransportMaps`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateTransportMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransportMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransportMapsRequest** | [**CreateTransportMapsRequest**](CreateTransportMapsRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSenderDependentTransports

> CreateAlias200Response DeleteSenderDependentTransports(ctx).DeleteSenderDependentTransportsRequest(deleteSenderDependentTransportsRequest).Execute()

Delete Sender-Dependent Transports



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
    deleteSenderDependentTransportsRequest := *openapiclient.NewDeleteSenderDependentTransportsRequest() // DeleteSenderDependentTransportsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteSenderDependentTransports(context.Background()).DeleteSenderDependentTransportsRequest(deleteSenderDependentTransportsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteSenderDependentTransports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSenderDependentTransports`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.DeleteSenderDependentTransports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSenderDependentTransportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSenderDependentTransportsRequest** | [**DeleteSenderDependentTransportsRequest**](DeleteSenderDependentTransportsRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransportMaps

> CreateAlias200Response DeleteTransportMaps(ctx).DeleteTransportMapsRequest(deleteTransportMapsRequest).Execute()

Delete Transport Maps



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
    deleteTransportMapsRequest := *openapiclient.NewDeleteTransportMapsRequest() // DeleteTransportMapsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteTransportMaps(context.Background()).DeleteTransportMapsRequest(deleteTransportMapsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteTransportMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTransportMaps`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.DeleteTransportMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransportMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTransportMapsRequest** | [**DeleteTransportMapsRequest**](DeleteTransportMapsRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSenderDependentTransports

> GetSenderDependentTransports(ctx, id).XAPIKey(xAPIKey).Execute()

Get Sender-Dependent Transports



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
    resp, r, err := apiClient.RoutingApi.GetSenderDependentTransports(context.Background(), id).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetSenderDependentTransports``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSenderDependentTransportsRequest struct via the builder pattern


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


## GetTransportMaps

> GetTransportMaps(ctx, id).XAPIKey(xAPIKey).Execute()

Get Transport Maps



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
    resp, r, err := apiClient.RoutingApi.GetTransportMaps(context.Background(), id).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetTransportMaps``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetTransportMapsRequest struct via the builder pattern


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

