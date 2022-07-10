# UpdateMailboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**UpdateMailboxRequestAttr**](UpdateMailboxRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains list of mailboxes you want update | [optional] 

## Methods

### NewUpdateMailboxRequest

`func NewUpdateMailboxRequest() *UpdateMailboxRequest`

NewUpdateMailboxRequest instantiates a new UpdateMailboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxRequestWithDefaults

`func NewUpdateMailboxRequestWithDefaults() *UpdateMailboxRequest`

NewUpdateMailboxRequestWithDefaults instantiates a new UpdateMailboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *UpdateMailboxRequest) GetAttr() UpdateMailboxRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *UpdateMailboxRequest) GetAttrOk() (*UpdateMailboxRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *UpdateMailboxRequest) SetAttr(v UpdateMailboxRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *UpdateMailboxRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *UpdateMailboxRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateMailboxRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateMailboxRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateMailboxRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


