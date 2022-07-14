# UpdateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to [**UpdateDomainRequestAttr**](UpdateDomainRequestAttr.md) |  | [optional] 
**Items** | Pointer to **[]string** | contains list of domain names you want update | [optional] 

## Methods

### NewUpdateDomainRequest

`func NewUpdateDomainRequest() *UpdateDomainRequest`

NewUpdateDomainRequest instantiates a new UpdateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainRequestWithDefaults

`func NewUpdateDomainRequestWithDefaults() *UpdateDomainRequest`

NewUpdateDomainRequestWithDefaults instantiates a new UpdateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *UpdateDomainRequest) GetAttr() UpdateDomainRequestAttr`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *UpdateDomainRequest) GetAttrOk() (*UpdateDomainRequestAttr, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *UpdateDomainRequest) SetAttr(v UpdateDomainRequestAttr)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *UpdateDomainRequest) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetItems

`func (o *UpdateDomainRequest) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateDomainRequest) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateDomainRequest) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateDomainRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


