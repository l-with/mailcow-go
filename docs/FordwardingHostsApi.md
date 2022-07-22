# \FordwardingHostsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddForwardHost**](FordwardingHostsApi.md#AddForwardHost) | **Post** /api/v1/add/fwdhost | Add Forward Host
[**DeleteForwardHost**](FordwardingHostsApi.md#DeleteForwardHost) | **Post** /api/v1/delete/fwdhost | Delete Forward Host
[**GetForwardingHosts**](FordwardingHostsApi.md#GetForwardingHosts) | **Get** /api/v1/get/fwdhost/all | Get Forwarding Hosts



## AddForwardHost

> CreateTimeLimitedAlias200Response AddForwardHost(ctx).AddForwardHostRequest(addForwardHostRequest).Execute()

Add Forward Host



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
    addForwardHostRequest := *openapiclient.NewAddForwardHostRequest() // AddForwardHostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FordwardingHostsApi.AddForwardHost(context.Background()).AddForwardHostRequest(addForwardHostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FordwardingHostsApi.AddForwardHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddForwardHost`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `FordwardingHostsApi.AddForwardHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddForwardHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addForwardHostRequest** | [**AddForwardHostRequest**](AddForwardHostRequest.md) |  | 

### Return type

[**CreateTimeLimitedAlias200Response**](CreateTimeLimitedAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForwardHost

> CreateTimeLimitedAlias200Response DeleteForwardHost(ctx).DeleteForwardHostRequest(deleteForwardHostRequest).Execute()

Delete Forward Host



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
    deleteForwardHostRequest := *openapiclient.NewDeleteForwardHostRequest() // DeleteForwardHostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FordwardingHostsApi.DeleteForwardHost(context.Background()).DeleteForwardHostRequest(deleteForwardHostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FordwardingHostsApi.DeleteForwardHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteForwardHost`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `FordwardingHostsApi.DeleteForwardHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForwardHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteForwardHostRequest** | [**DeleteForwardHostRequest**](DeleteForwardHostRequest.md) |  | 

### Return type

[**CreateTimeLimitedAlias200Response**](CreateTimeLimitedAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForwardingHosts

> GetForwardingHosts(ctx).Execute()

Get Forwarding Hosts



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FordwardingHostsApi.GetForwardingHosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FordwardingHostsApi.GetForwardingHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetForwardingHostsRequest struct via the builder pattern


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

