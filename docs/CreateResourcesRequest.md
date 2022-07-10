# CreateResourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **float32** | 1 for a active transport map 0 for a disabled transport map | [optional] 
**Description** | Pointer to **string** | a description of the resource | [optional] 
**Domain** | Pointer to **string** | the domain for which the resource should be | [optional] 
**Kind** | Pointer to **string** | the kind of recouse | [optional] 
**MultipleBookings** | Pointer to **string** |  | [optional] 
**MultipleBookingsCustom** | Pointer to **float32** | always empty | [optional] 
**MultipleBookingsSelect** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateResourcesRequest

`func NewCreateResourcesRequest() *CreateResourcesRequest`

NewCreateResourcesRequest instantiates a new CreateResourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourcesRequestWithDefaults

`func NewCreateResourcesRequestWithDefaults() *CreateResourcesRequest`

NewCreateResourcesRequestWithDefaults instantiates a new CreateResourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateResourcesRequest) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateResourcesRequest) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateResourcesRequest) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateResourcesRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *CreateResourcesRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourcesRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourcesRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateResourcesRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *CreateResourcesRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateResourcesRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateResourcesRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateResourcesRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetKind

`func (o *CreateResourcesRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateResourcesRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateResourcesRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateResourcesRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMultipleBookings

`func (o *CreateResourcesRequest) GetMultipleBookings() string`

GetMultipleBookings returns the MultipleBookings field if non-nil, zero value otherwise.

### GetMultipleBookingsOk

`func (o *CreateResourcesRequest) GetMultipleBookingsOk() (*string, bool)`

GetMultipleBookingsOk returns a tuple with the MultipleBookings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleBookings

`func (o *CreateResourcesRequest) SetMultipleBookings(v string)`

SetMultipleBookings sets MultipleBookings field to given value.

### HasMultipleBookings

`func (o *CreateResourcesRequest) HasMultipleBookings() bool`

HasMultipleBookings returns a boolean if a field has been set.

### GetMultipleBookingsCustom

`func (o *CreateResourcesRequest) GetMultipleBookingsCustom() float32`

GetMultipleBookingsCustom returns the MultipleBookingsCustom field if non-nil, zero value otherwise.

### GetMultipleBookingsCustomOk

`func (o *CreateResourcesRequest) GetMultipleBookingsCustomOk() (*float32, bool)`

GetMultipleBookingsCustomOk returns a tuple with the MultipleBookingsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleBookingsCustom

`func (o *CreateResourcesRequest) SetMultipleBookingsCustom(v float32)`

SetMultipleBookingsCustom sets MultipleBookingsCustom field to given value.

### HasMultipleBookingsCustom

`func (o *CreateResourcesRequest) HasMultipleBookingsCustom() bool`

HasMultipleBookingsCustom returns a boolean if a field has been set.

### GetMultipleBookingsSelect

`func (o *CreateResourcesRequest) GetMultipleBookingsSelect() string`

GetMultipleBookingsSelect returns the MultipleBookingsSelect field if non-nil, zero value otherwise.

### GetMultipleBookingsSelectOk

`func (o *CreateResourcesRequest) GetMultipleBookingsSelectOk() (*string, bool)`

GetMultipleBookingsSelectOk returns a tuple with the MultipleBookingsSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleBookingsSelect

`func (o *CreateResourcesRequest) SetMultipleBookingsSelect(v string)`

SetMultipleBookingsSelect sets MultipleBookingsSelect field to given value.

### HasMultipleBookingsSelect

`func (o *CreateResourcesRequest) HasMultipleBookingsSelect() bool`

HasMultipleBookingsSelect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


