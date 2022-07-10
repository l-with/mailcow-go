# CreateMailboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is mailbox active or not | [optional] 
**Domain** | Pointer to **string** | domain name | [optional] 
**LocalPart** | Pointer to **string** | left part of email address | [optional] 
**Name** | Pointer to **string** | Full name of the mailbox user | [optional] 
**Password2** | Pointer to **string** | mailbox password for confirmation | [optional] 
**Password** | Pointer to **string** | mailbox password | [optional] 
**Quota** | Pointer to **float32** | mailbox quota | [optional] 
**ForcePwUpdate** | Pointer to **bool** | forces the user to update its password on first login | [optional] 
**TlsEnforceIn** | Pointer to **bool** | force inbound email tls encryption | [optional] 
**TlsEnforceOut** | Pointer to **bool** | force oubound tmail tls encryption | [optional] 

## Methods

### NewCreateMailboxRequest

`func NewCreateMailboxRequest() *CreateMailboxRequest`

NewCreateMailboxRequest instantiates a new CreateMailboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMailboxRequestWithDefaults

`func NewCreateMailboxRequestWithDefaults() *CreateMailboxRequest`

NewCreateMailboxRequestWithDefaults instantiates a new CreateMailboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateMailboxRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateMailboxRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateMailboxRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateMailboxRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDomain

`func (o *CreateMailboxRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateMailboxRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateMailboxRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateMailboxRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLocalPart

`func (o *CreateMailboxRequest) GetLocalPart() string`

GetLocalPart returns the LocalPart field if non-nil, zero value otherwise.

### GetLocalPartOk

`func (o *CreateMailboxRequest) GetLocalPartOk() (*string, bool)`

GetLocalPartOk returns a tuple with the LocalPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPart

`func (o *CreateMailboxRequest) SetLocalPart(v string)`

SetLocalPart sets LocalPart field to given value.

### HasLocalPart

`func (o *CreateMailboxRequest) HasLocalPart() bool`

HasLocalPart returns a boolean if a field has been set.

### GetName

`func (o *CreateMailboxRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMailboxRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMailboxRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateMailboxRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword2

`func (o *CreateMailboxRequest) GetPassword2() string`

GetPassword2 returns the Password2 field if non-nil, zero value otherwise.

### GetPassword2Ok

`func (o *CreateMailboxRequest) GetPassword2Ok() (*string, bool)`

GetPassword2Ok returns a tuple with the Password2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword2

`func (o *CreateMailboxRequest) SetPassword2(v string)`

SetPassword2 sets Password2 field to given value.

### HasPassword2

`func (o *CreateMailboxRequest) HasPassword2() bool`

HasPassword2 returns a boolean if a field has been set.

### GetPassword

`func (o *CreateMailboxRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateMailboxRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateMailboxRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateMailboxRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetQuota

`func (o *CreateMailboxRequest) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateMailboxRequest) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateMailboxRequest) SetQuota(v float32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateMailboxRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetForcePwUpdate

`func (o *CreateMailboxRequest) GetForcePwUpdate() bool`

GetForcePwUpdate returns the ForcePwUpdate field if non-nil, zero value otherwise.

### GetForcePwUpdateOk

`func (o *CreateMailboxRequest) GetForcePwUpdateOk() (*bool, bool)`

GetForcePwUpdateOk returns a tuple with the ForcePwUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePwUpdate

`func (o *CreateMailboxRequest) SetForcePwUpdate(v bool)`

SetForcePwUpdate sets ForcePwUpdate field to given value.

### HasForcePwUpdate

`func (o *CreateMailboxRequest) HasForcePwUpdate() bool`

HasForcePwUpdate returns a boolean if a field has been set.

### GetTlsEnforceIn

`func (o *CreateMailboxRequest) GetTlsEnforceIn() bool`

GetTlsEnforceIn returns the TlsEnforceIn field if non-nil, zero value otherwise.

### GetTlsEnforceInOk

`func (o *CreateMailboxRequest) GetTlsEnforceInOk() (*bool, bool)`

GetTlsEnforceInOk returns a tuple with the TlsEnforceIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnforceIn

`func (o *CreateMailboxRequest) SetTlsEnforceIn(v bool)`

SetTlsEnforceIn sets TlsEnforceIn field to given value.

### HasTlsEnforceIn

`func (o *CreateMailboxRequest) HasTlsEnforceIn() bool`

HasTlsEnforceIn returns a boolean if a field has been set.

### GetTlsEnforceOut

`func (o *CreateMailboxRequest) GetTlsEnforceOut() bool`

GetTlsEnforceOut returns the TlsEnforceOut field if non-nil, zero value otherwise.

### GetTlsEnforceOutOk

`func (o *CreateMailboxRequest) GetTlsEnforceOutOk() (*bool, bool)`

GetTlsEnforceOutOk returns a tuple with the TlsEnforceOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnforceOut

`func (o *CreateMailboxRequest) SetTlsEnforceOut(v bool)`

SetTlsEnforceOut sets TlsEnforceOut field to given value.

### HasTlsEnforceOut

`func (o *CreateMailboxRequest) HasTlsEnforceOut() bool`

HasTlsEnforceOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


