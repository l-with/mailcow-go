# \Fail2BanApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditFail2Ban**](Fail2BanApi.md#EditFail2Ban) | **Post** /api/v1/edit/fail2ban | Edit Fail2Ban
[**GetFail2BanConfig**](Fail2BanApi.md#GetFail2BanConfig) | **Get** /api/v1/get/fail2ban | Get Fail2Ban Config



## EditFail2Ban

> CreateAlias200ResponseInner EditFail2Ban(ctx).EditFail2BanRequest(editFail2BanRequest).Execute()

Edit Fail2Ban



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
    editFail2BanRequest := *openapiclient.NewEditFail2BanRequest() // EditFail2BanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Fail2BanApi.EditFail2Ban(context.Background()).EditFail2BanRequest(editFail2BanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Fail2BanApi.EditFail2Ban``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditFail2Ban`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `Fail2BanApi.EditFail2Ban`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditFail2BanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editFail2BanRequest** | [**EditFail2BanRequest**](EditFail2BanRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFail2BanConfig

> GetFail2BanConfig(ctx).Execute()

Get Fail2Ban Config



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
    resp, r, err := apiClient.Fail2BanApi.GetFail2BanConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Fail2BanApi.GetFail2BanConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFail2BanConfigRequest struct via the builder pattern


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

