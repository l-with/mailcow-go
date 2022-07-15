# DeleteDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **map[string]interface{}** | contains list of domains you want to delete | [optional] 

## Methods

### NewDeleteDomainRequest

`func NewDeleteDomainRequest() *DeleteDomainRequest`

NewDeleteDomainRequest instantiates a new DeleteDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDomainRequestWithDefaults

`func NewDeleteDomainRequestWithDefaults() *DeleteDomainRequest`

NewDeleteDomainRequestWithDefaults instantiates a new DeleteDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeleteDomainRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeleteDomainRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeleteDomainRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *DeleteDomainRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


