# EditDomainAdminUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**EditDomainAdminUserRequestAttr**](EditDomainAdminUserRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains the domain admin username you want to edit | [optional] 

## Methods

### NewEditDomainAdminUserRequest

`func NewEditDomainAdminUserRequest() *EditDomainAdminUserRequest`

NewEditDomainAdminUserRequest instantiates a new EditDomainAdminUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDomainAdminUserRequestWithDefaults

`func NewEditDomainAdminUserRequestWithDefaults() *EditDomainAdminUserRequest`

NewEditDomainAdminUserRequestWithDefaults instantiates a new EditDomainAdminUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *EditDomainAdminUserRequest) GetAttr() EditDomainAdminUserRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *EditDomainAdminUserRequest) GetAttrOk() (*EditDomainAdminUserRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *EditDomainAdminUserRequest) SetAttr(v EditDomainAdminUserRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *EditDomainAdminUserRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *EditDomainAdminUserRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EditDomainAdminUserRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EditDomainAdminUserRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *EditDomainAdminUserRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


