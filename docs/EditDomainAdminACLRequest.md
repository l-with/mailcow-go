# EditDomainAdminACLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **map[string]interface{}** | contains the domain admin username you want to edit | [optional] 
**Attr** | Pointer to [**EditDomainAdminACLRequestAttr**](EditDomainAdminACLRequestAttr.md) |  | [optional] 

## Methods

### NewEditDomainAdminACLRequest

`func NewEditDomainAdminACLRequest() *EditDomainAdminACLRequest`

NewEditDomainAdminACLRequest instantiates a new EditDomainAdminACLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDomainAdminACLRequestWithDefaults

`func NewEditDomainAdminACLRequestWithDefaults() *EditDomainAdminACLRequest`

NewEditDomainAdminACLRequestWithDefaults instantiates a new EditDomainAdminACLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EditDomainAdminACLRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EditDomainAdminACLRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EditDomainAdminACLRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *EditDomainAdminACLRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAttr

`func (o *EditDomainAdminACLRequest) GetAttr() EditDomainAdminACLRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *EditDomainAdminACLRequest) GetAttrOk() (*EditDomainAdminACLRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *EditDomainAdminACLRequest) SetAttr(v EditDomainAdminACLRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *EditDomainAdminACLRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


