# \QuarantineApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMailsInQuarantine**](QuarantineApi.md#DeleteMailsInQuarantine) | **Post** /api/v1/delete/qitem | Delete mails in Quarantine
[**GetMailsInQuarantine**](QuarantineApi.md#GetMailsInQuarantine) | **Get** /api/v1/get/quarantine/all | Get mails in Quarantine



## DeleteMailsInQuarantine

> CreateTimeLimitedAlias200Response DeleteMailsInQuarantine(ctx).DeleteMailsInQuarantineRequest(deleteMailsInQuarantineRequest).Execute()

Delete mails in Quarantine



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
    deleteMailsInQuarantineRequest := *openapiclient.NewDeleteMailsInQuarantineRequest() // DeleteMailsInQuarantineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuarantineApi.DeleteMailsInQuarantine(context.Background()).DeleteMailsInQuarantineRequest(deleteMailsInQuarantineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuarantineApi.DeleteMailsInQuarantine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMailsInQuarantine`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `QuarantineApi.DeleteMailsInQuarantine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailsInQuarantineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMailsInQuarantineRequest** | [**DeleteMailsInQuarantineRequest**](DeleteMailsInQuarantineRequest.md) |  | 

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


## GetMailsInQuarantine

> GetMailsInQuarantine(ctx).Execute()

Get mails in Quarantine



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
    resp, r, err := apiClient.QuarantineApi.GetMailsInQuarantine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuarantineApi.GetMailsInQuarantine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailsInQuarantineRequest struct via the builder pattern


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

