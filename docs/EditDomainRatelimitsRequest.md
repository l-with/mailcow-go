# EditDomainRatelimitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**EditMailboxRatelimitsRequestAttr**](EditMailboxRatelimitsRequestAttr.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | contains list of domains you want to edit the ratelimit of | [optional] 

## Methods

### NewEditDomainRatelimitsRequest

`func NewEditDomainRatelimitsRequest() *EditDomainRatelimitsRequest`

NewEditDomainRatelimitsRequest instantiates a new EditDomainRatelimitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDomainRatelimitsRequestWithDefaults

`func NewEditDomainRatelimitsRequestWithDefaults() *EditDomainRatelimitsRequest`

NewEditDomainRatelimitsRequestWithDefaults instantiates a new EditDomainRatelimitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *EditDomainRatelimitsRequest) GetAttr() EditMailboxRatelimitsRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *EditDomainRatelimitsRequest) GetAttrOk() (*EditMailboxRatelimitsRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *EditDomainRatelimitsRequest) SetAttr(v EditMailboxRatelimitsRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *EditDomainRatelimitsRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *EditDomainRatelimitsRequest) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EditDomainRatelimitsRequest) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EditDomainRatelimitsRequest) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *EditDomainRatelimitsRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


