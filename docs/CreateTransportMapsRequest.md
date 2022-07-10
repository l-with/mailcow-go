# CreateTransportMapsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **float32** | 1 for a active transport map 0 for a disabled transport map | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**Nexthop** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | the password for the smtp user | [optional] 
**Username** | Pointer to **string** | the username used to authenticate | [optional] 

## Methods

### NewCreateTransportMapsRequest

`func NewCreateTransportMapsRequest() *CreateTransportMapsRequest`

NewCreateTransportMapsRequest instantiates a new CreateTransportMapsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportMapsRequestWithDefaults

`func NewCreateTransportMapsRequestWithDefaults() *CreateTransportMapsRequest`

NewCreateTransportMapsRequestWithDefaults instantiates a new CreateTransportMapsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateTransportMapsRequest) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateTransportMapsRequest) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateTransportMapsRequest) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateTransportMapsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDestination

`func (o *CreateTransportMapsRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreateTransportMapsRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreateTransportMapsRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CreateTransportMapsRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNexthop

`func (o *CreateTransportMapsRequest) GetNexthop() string`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *CreateTransportMapsRequest) GetNexthopOk() (*string, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *CreateTransportMapsRequest) SetNexthop(v string)`

SetNexthop sets Nexthop field to given value.

### HasNexthop

`func (o *CreateTransportMapsRequest) HasNexthop() bool`

HasNexthop returns a boolean if a field has been set.

### GetPassword

`func (o *CreateTransportMapsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateTransportMapsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateTransportMapsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateTransportMapsRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *CreateTransportMapsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateTransportMapsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateTransportMapsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateTransportMapsRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


