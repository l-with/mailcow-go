# CreateRecipientMapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **float32** | 1 for a active user account 0 for a disabled user account | [optional] 
**RecipientMapNew** | Pointer to **string** | the email address that should receive the forwarded emails | [optional] 
**RecipientMapOld** | Pointer to **string** | the email address which emails should be forwarded | [optional] 

## Methods

### NewCreateRecipientMapRequest

`func NewCreateRecipientMapRequest() *CreateRecipientMapRequest`

NewCreateRecipientMapRequest instantiates a new CreateRecipientMapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecipientMapRequestWithDefaults

`func NewCreateRecipientMapRequestWithDefaults() *CreateRecipientMapRequest`

NewCreateRecipientMapRequestWithDefaults instantiates a new CreateRecipientMapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateRecipientMapRequest) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateRecipientMapRequest) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateRecipientMapRequest) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateRecipientMapRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRecipientMapNew

`func (o *CreateRecipientMapRequest) GetRecipientMapNew() string`

GetRecipientMapNew returns the RecipientMapNew field if non-nil, zero value otherwise.

### GetRecipientMapNewOk

`func (o *CreateRecipientMapRequest) GetRecipientMapNewOk() (*string, bool)`

GetRecipientMapNewOk returns a tuple with the RecipientMapNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMapNew

`func (o *CreateRecipientMapRequest) SetRecipientMapNew(v string)`

SetRecipientMapNew sets RecipientMapNew field to given value.

### HasRecipientMapNew

`func (o *CreateRecipientMapRequest) HasRecipientMapNew() bool`

HasRecipientMapNew returns a boolean if a field has been set.

### GetRecipientMapOld

`func (o *CreateRecipientMapRequest) GetRecipientMapOld() string`

GetRecipientMapOld returns the RecipientMapOld field if non-nil, zero value otherwise.

### GetRecipientMapOldOk

`func (o *CreateRecipientMapRequest) GetRecipientMapOldOk() (*string, bool)`

GetRecipientMapOldOk returns a tuple with the RecipientMapOld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMapOld

`func (o *CreateRecipientMapRequest) SetRecipientMapOld(v string)`

SetRecipientMapOld sets RecipientMapOld field to given value.

### HasRecipientMapOld

`func (o *CreateRecipientMapRequest) HasRecipientMapOld() bool`

HasRecipientMapOld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


