# CreateSenderDependentTransportsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | the hostname of the smtp server with port | [optional] 
**Password** | Pointer to **string** | the password for the smtp user | [optional] 
**Username** | Pointer to **string** | the username used to authenticate | [optional] 

## Methods

### NewCreateSenderDependentTransportsRequest

`func NewCreateSenderDependentTransportsRequest() *CreateSenderDependentTransportsRequest`

NewCreateSenderDependentTransportsRequest instantiates a new CreateSenderDependentTransportsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSenderDependentTransportsRequestWithDefaults

`func NewCreateSenderDependentTransportsRequestWithDefaults() *CreateSenderDependentTransportsRequest`

NewCreateSenderDependentTransportsRequestWithDefaults instantiates a new CreateSenderDependentTransportsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *CreateSenderDependentTransportsRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateSenderDependentTransportsRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateSenderDependentTransportsRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateSenderDependentTransportsRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPassword

`func (o *CreateSenderDependentTransportsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSenderDependentTransportsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSenderDependentTransportsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSenderDependentTransportsRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *CreateSenderDependentTransportsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateSenderDependentTransportsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateSenderDependentTransportsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateSenderDependentTransportsRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


