# UpdatePushoverSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**UpdatePushoverSettingsRequestAttr**](UpdatePushoverSettingsRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains list of mailboxes you want to delete | [optional] 

## Methods

### NewUpdatePushoverSettingsRequest

`func NewUpdatePushoverSettingsRequest() *UpdatePushoverSettingsRequest`

NewUpdatePushoverSettingsRequest instantiates a new UpdatePushoverSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePushoverSettingsRequestWithDefaults

`func NewUpdatePushoverSettingsRequestWithDefaults() *UpdatePushoverSettingsRequest`

NewUpdatePushoverSettingsRequestWithDefaults instantiates a new UpdatePushoverSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *UpdatePushoverSettingsRequest) GetAttr() UpdatePushoverSettingsRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *UpdatePushoverSettingsRequest) GetAttrOk() (*UpdatePushoverSettingsRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *UpdatePushoverSettingsRequest) SetAttr(v UpdatePushoverSettingsRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *UpdatePushoverSettingsRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *UpdatePushoverSettingsRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdatePushoverSettingsRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdatePushoverSettingsRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdatePushoverSettingsRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


