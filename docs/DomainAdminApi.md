# \DomainAdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainAdminUser**](DomainAdminApi.md#CreateDomainAdminUser) | **Post** /api/v1/add/domain-admin | Create Domain Admin user
[**DeleteDomainAdmin**](DomainAdminApi.md#DeleteDomainAdmin) | **Post** /api/v1/delete/domain-admin | Delete Domain Admin
[**EditDomainAdminACL**](DomainAdminApi.md#EditDomainAdminACL) | **Post** /api/v1/edit/da-acl | Edit Domain Admin ACL
[**EditDomainAdminUser**](DomainAdminApi.md#EditDomainAdminUser) | **Post** /api/v1/edit/domain-admin | Edit Domain Admin user
[**GetDomainAdmins**](DomainAdminApi.md#GetDomainAdmins) | **Get** /api/v1/get/domain-admin/all | Get Domain Admins



## CreateDomainAdminUser

> CreateAlias200Response CreateDomainAdminUser(ctx).CreateDomainAdminUserRequest(createDomainAdminUserRequest).Execute()

Create Domain Admin user



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
    createDomainAdminUserRequest := *openapiclient.NewCreateDomainAdminUserRequest() // CreateDomainAdminUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAdminApi.CreateDomainAdminUser(context.Background()).CreateDomainAdminUserRequest(createDomainAdminUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminApi.CreateDomainAdminUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainAdminUser`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAdminApi.CreateDomainAdminUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainAdminUserRequest** | [**CreateDomainAdminUserRequest**](CreateDomainAdminUserRequest.md) |  | 

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


## DeleteDomainAdmin

> CreateAlias200Response DeleteDomainAdmin(ctx).DeleteDomainAdminRequest(deleteDomainAdminRequest).Execute()

Delete Domain Admin



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
    deleteDomainAdminRequest := *openapiclient.NewDeleteDomainAdminRequest() // DeleteDomainAdminRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAdminApi.DeleteDomainAdmin(context.Background()).DeleteDomainAdminRequest(deleteDomainAdminRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminApi.DeleteDomainAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomainAdmin`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAdminApi.DeleteDomainAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDomainAdminRequest** | [**DeleteDomainAdminRequest**](DeleteDomainAdminRequest.md) |  | 

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


## EditDomainAdminACL

> CreateAlias200Response EditDomainAdminACL(ctx).EditDomainAdminACLRequest(editDomainAdminACLRequest).Execute()

Edit Domain Admin ACL



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
    editDomainAdminACLRequest := *openapiclient.NewEditDomainAdminACLRequest() // EditDomainAdminACLRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAdminApi.EditDomainAdminACL(context.Background()).EditDomainAdminACLRequest(editDomainAdminACLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminApi.EditDomainAdminACL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDomainAdminACL`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAdminApi.EditDomainAdminACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDomainAdminACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDomainAdminACLRequest** | [**EditDomainAdminACLRequest**](EditDomainAdminACLRequest.md) |  | 

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


## EditDomainAdminUser

> EditDomainAdminUser200Response EditDomainAdminUser(ctx).EditDomainAdminUserRequest(editDomainAdminUserRequest).Execute()

Edit Domain Admin user



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
    editDomainAdminUserRequest := *openapiclient.NewEditDomainAdminUserRequest() // EditDomainAdminUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAdminApi.EditDomainAdminUser(context.Background()).EditDomainAdminUserRequest(editDomainAdminUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminApi.EditDomainAdminUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDomainAdminUser`: EditDomainAdminUser200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAdminApi.EditDomainAdminUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDomainAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDomainAdminUserRequest** | [**EditDomainAdminUserRequest**](EditDomainAdminUserRequest.md) |  | 

### Return type

[**EditDomainAdminUser200Response**](EditDomainAdminUser200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainAdmins

> GetDomainAdmins(ctx).Execute()

Get Domain Admins



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
    resp, r, err := apiClient.DomainAdminApi.GetDomainAdmins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminApi.GetDomainAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainAdminsRequest struct via the builder pattern


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

