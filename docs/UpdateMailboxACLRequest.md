# UpdateMailboxACLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**UpdateMailboxACLRequestAttr**](UpdateMailboxACLRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains list of mailboxes you want to delete | [optional] 

## Methods

### NewUpdateMailboxACLRequest

`func NewUpdateMailboxACLRequest() *UpdateMailboxACLRequest`

NewUpdateMailboxACLRequest instantiates a new UpdateMailboxACLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxACLRequestWithDefaults

`func NewUpdateMailboxACLRequestWithDefaults() *UpdateMailboxACLRequest`

NewUpdateMailboxACLRequestWithDefaults instantiates a new UpdateMailboxACLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *UpdateMailboxACLRequest) GetAttr() UpdateMailboxACLRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *UpdateMailboxACLRequest) GetAttrOk() (*UpdateMailboxACLRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *UpdateMailboxACLRequest) SetAttr(v UpdateMailboxACLRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *UpdateMailboxACLRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *UpdateMailboxACLRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateMailboxACLRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateMailboxACLRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateMailboxACLRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


