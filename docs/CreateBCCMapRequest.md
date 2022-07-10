# CreateBCCMapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **float32** | 1 for a active user account 0 for a disabled user account | [optional] 
**BccDest** | Pointer to **string** | the email address where all mails should be send to | [optional] 
**LocalDest** | Pointer to **string** | the domain which emails should be forwarded | [optional] 
**Type** | Pointer to **string** | the type of bcc map can be &#x60;sender&#x60; or &#x60;recipient&#x60; | [optional] 

## Methods

### NewCreateBCCMapRequest

`func NewCreateBCCMapRequest() *CreateBCCMapRequest`

NewCreateBCCMapRequest instantiates a new CreateBCCMapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBCCMapRequestWithDefaults

`func NewCreateBCCMapRequestWithDefaults() *CreateBCCMapRequest`

NewCreateBCCMapRequestWithDefaults instantiates a new CreateBCCMapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateBCCMapRequest) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateBCCMapRequest) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateBCCMapRequest) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateBCCMapRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBccDest

`func (o *CreateBCCMapRequest) GetBccDest() string`

GetBccDest returns the BccDest field if non-nil, zero value otherwise.

### GetBccDestOk

`func (o *CreateBCCMapRequest) GetBccDestOk() (*string, bool)`

GetBccDestOk returns a tuple with the BccDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccDest

`func (o *CreateBCCMapRequest) SetBccDest(v string)`

SetBccDest sets BccDest field to given value.

### HasBccDest

`func (o *CreateBCCMapRequest) HasBccDest() bool`

HasBccDest returns a boolean if a field has been set.

### GetLocalDest

`func (o *CreateBCCMapRequest) GetLocalDest() string`

GetLocalDest returns the LocalDest field if non-nil, zero value otherwise.

### GetLocalDestOk

`func (o *CreateBCCMapRequest) GetLocalDestOk() (*string, bool)`

GetLocalDestOk returns a tuple with the LocalDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDest

`func (o *CreateBCCMapRequest) SetLocalDest(v string)`

SetLocalDest sets LocalDest field to given value.

### HasLocalDest

`func (o *CreateBCCMapRequest) HasLocalDest() bool`

HasLocalDest returns a boolean if a field has been set.

### GetType

`func (o *CreateBCCMapRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBCCMapRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBCCMapRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateBCCMapRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


