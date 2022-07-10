# DeleteMailboxTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **map[string]interface{}** | contains list of mailboxes you want to delete | [optional] 

## Methods

### NewDeleteMailboxTagsRequest

`func NewDeleteMailboxTagsRequest() *DeleteMailboxTagsRequest`

NewDeleteMailboxTagsRequest instantiates a new DeleteMailboxTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMailboxTagsRequestWithDefaults

`func NewDeleteMailboxTagsRequestWithDefaults() *DeleteMailboxTagsRequest`

NewDeleteMailboxTagsRequestWithDefaults instantiates a new DeleteMailboxTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeleteMailboxTagsRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeleteMailboxTagsRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeleteMailboxTagsRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *DeleteMailboxTagsRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


