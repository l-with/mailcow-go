# UpdateAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**UpdateAliasRequestAttr**](UpdateAliasRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains list of aliases you want update | [optional] 

## Methods

### NewUpdateAliasRequest

`func NewUpdateAliasRequest() *UpdateAliasRequest`

NewUpdateAliasRequest instantiates a new UpdateAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAliasRequestWithDefaults

`func NewUpdateAliasRequestWithDefaults() *UpdateAliasRequest`

NewUpdateAliasRequestWithDefaults instantiates a new UpdateAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *UpdateAliasRequest) GetAttr() UpdateAliasRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *UpdateAliasRequest) GetAttrOk() (*UpdateAliasRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *UpdateAliasRequest) SetAttr(v UpdateAliasRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *UpdateAliasRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *UpdateAliasRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateAliasRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateAliasRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateAliasRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


