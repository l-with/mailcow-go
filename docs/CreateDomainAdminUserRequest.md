# CreateDomainAdminUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **float32** | 1 for a active user account 0 for a disabled user account | [optional] 
**Domains** | Pointer to **string** | the domains the user should be a admin of | [optional] 
**Password** | Pointer to **string** | domain admin user password | [optional] 
**Password2** | Pointer to **string** | domain admin user password | [optional] 
**Username** | Pointer to **string** | the username for the admin user | [optional] 

## Methods

### NewCreateDomainAdminUserRequest

`func NewCreateDomainAdminUserRequest() *CreateDomainAdminUserRequest`

NewCreateDomainAdminUserRequest instantiates a new CreateDomainAdminUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainAdminUserRequestWithDefaults

`func NewCreateDomainAdminUserRequestWithDefaults() *CreateDomainAdminUserRequest`

NewCreateDomainAdminUserRequestWithDefaults instantiates a new CreateDomainAdminUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateDomainAdminUserRequest) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateDomainAdminUserRequest) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateDomainAdminUserRequest) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateDomainAdminUserRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDomains

`func (o *CreateDomainAdminUserRequest) GetDomains() string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateDomainAdminUserRequest) GetDomainsOk() (*string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateDomainAdminUserRequest) SetDomains(v string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateDomainAdminUserRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetPassword

`func (o *CreateDomainAdminUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDomainAdminUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDomainAdminUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateDomainAdminUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassword2

`func (o *CreateDomainAdminUserRequest) GetPassword2() string`

GetPassword2 returns the Password2 field if non-nil, zero value otherwise.

### GetPassword2Ok

`func (o *CreateDomainAdminUserRequest) GetPassword2Ok() (*string, bool)`

GetPassword2Ok returns a tuple with the Password2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword2

`func (o *CreateDomainAdminUserRequest) SetPassword2(v string)`

SetPassword2 sets Password2 field to given value.

### HasPassword2

`func (o *CreateDomainAdminUserRequest) HasPassword2() bool`

HasPassword2 returns a boolean if a field has been set.

### GetUsername

`func (o *CreateDomainAdminUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDomainAdminUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDomainAdminUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateDomainAdminUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


