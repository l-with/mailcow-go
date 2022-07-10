# DeleteRecipientMapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **map[string]interface{}** | contains list of recipient maps you want to delete | [optional] 

## Methods

### NewDeleteRecipientMapRequest

`func NewDeleteRecipientMapRequest() *DeleteRecipientMapRequest`

NewDeleteRecipientMapRequest instantiates a new DeleteRecipientMapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRecipientMapRequestWithDefaults

`func NewDeleteRecipientMapRequestWithDefaults() *DeleteRecipientMapRequest`

NewDeleteRecipientMapRequestWithDefaults instantiates a new DeleteRecipientMapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeleteRecipientMapRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeleteRecipientMapRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeleteRecipientMapRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *DeleteRecipientMapRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


