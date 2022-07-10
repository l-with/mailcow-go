# DeleteDomainAdminRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **map[string]interface{}** | contains list of usernames of the users you want to delete | [optional] 

## Methods

### NewDeleteDomainAdminRequest

`func NewDeleteDomainAdminRequest() *DeleteDomainAdminRequest`

NewDeleteDomainAdminRequest instantiates a new DeleteDomainAdminRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDomainAdminRequestWithDefaults

`func NewDeleteDomainAdminRequestWithDefaults() *DeleteDomainAdminRequest`

NewDeleteDomainAdminRequestWithDefaults instantiates a new DeleteDomainAdminRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeleteDomainAdminRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeleteDomainAdminRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeleteDomainAdminRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *DeleteDomainAdminRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


