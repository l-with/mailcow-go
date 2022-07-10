# EditDomainAdminUserRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is the domain admin active or not | [optional] 
**UsernameNew** | Pointer to **string** | the username of the domain admin, change this to change the username | [optional] 
**Domains** | Pointer to **[]string** | a list of all domains managed by this domain admin | [optional] 
**Password** | Pointer to **string** | the new domain admin user password | [optional] 
**Password2** | Pointer to **string** | the new domain admin user password for confirmation | [optional] 

## Methods

### NewEditDomainAdminUserRequestAttr

`func NewEditDomainAdminUserRequestAttr() *EditDomainAdminUserRequestAttr`

NewEditDomainAdminUserRequestAttr instantiates a new EditDomainAdminUserRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDomainAdminUserRequestAttrWithDefaults

`func NewEditDomainAdminUserRequestAttrWithDefaults() *EditDomainAdminUserRequestAttr`

NewEditDomainAdminUserRequestAttrWithDefaults instantiates a new EditDomainAdminUserRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *EditDomainAdminUserRequestAttr) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EditDomainAdminUserRequestAttr) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EditDomainAdminUserRequestAttr) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EditDomainAdminUserRequestAttr) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUsernameNew

`func (o *EditDomainAdminUserRequestAttr) GetUsernameNew() string`

GetUsernameNew returns the UsernameNew field if non-nil, zero value otherwise.

### GetUsernameNewOk

`func (o *EditDomainAdminUserRequestAttr) GetUsernameNewOk() (*string, bool)`

GetUsernameNewOk returns a tuple with the UsernameNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameNew

`func (o *EditDomainAdminUserRequestAttr) SetUsernameNew(v string)`

SetUsernameNew sets UsernameNew field to given value.

### HasUsernameNew

`func (o *EditDomainAdminUserRequestAttr) HasUsernameNew() bool`

HasUsernameNew returns a boolean if a field has been set.

### GetDomains

`func (o *EditDomainAdminUserRequestAttr) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *EditDomainAdminUserRequestAttr) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *EditDomainAdminUserRequestAttr) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *EditDomainAdminUserRequestAttr) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetPassword

`func (o *EditDomainAdminUserRequestAttr) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EditDomainAdminUserRequestAttr) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EditDomainAdminUserRequestAttr) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EditDomainAdminUserRequestAttr) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassword2

`func (o *EditDomainAdminUserRequestAttr) GetPassword2() string`

GetPassword2 returns the Password2 field if non-nil, zero value otherwise.

### GetPassword2Ok

`func (o *EditDomainAdminUserRequestAttr) GetPassword2Ok() (*string, bool)`

GetPassword2Ok returns a tuple with the Password2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword2

`func (o *EditDomainAdminUserRequestAttr) SetPassword2(v string)`

SetPassword2 sets Password2 field to given value.

### HasPassword2

`func (o *EditDomainAdminUserRequestAttr) HasPassword2() bool`

HasPassword2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


