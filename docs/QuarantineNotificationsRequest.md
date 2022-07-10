# QuarantineNotificationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**QuarantineNotificationsRequestAttr**](QuarantineNotificationsRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains list of mailboxes you want set qurantine notifications | [optional] 

## Methods

### NewQuarantineNotificationsRequest

`func NewQuarantineNotificationsRequest() *QuarantineNotificationsRequest`

NewQuarantineNotificationsRequest instantiates a new QuarantineNotificationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantineNotificationsRequestWithDefaults

`func NewQuarantineNotificationsRequestWithDefaults() *QuarantineNotificationsRequest`

NewQuarantineNotificationsRequestWithDefaults instantiates a new QuarantineNotificationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *QuarantineNotificationsRequest) GetAttr() QuarantineNotificationsRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *QuarantineNotificationsRequest) GetAttrOk() (*QuarantineNotificationsRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *QuarantineNotificationsRequest) SetAttr(v QuarantineNotificationsRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *QuarantineNotificationsRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *QuarantineNotificationsRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QuarantineNotificationsRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QuarantineNotificationsRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *QuarantineNotificationsRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


