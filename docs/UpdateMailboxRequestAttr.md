# UpdateMailboxRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is mailbox active or not | [optional] 
**ForcePwUpdate** | Pointer to **bool** | force user to change password on next login | [optional] 
**Name** | Pointer to **string** | Full name of the mailbox user | [optional] 
**Password2** | Pointer to **string** | new mailbox password for confirmation | [optional] 
**Password** | Pointer to **string** | new mailbox password | [optional] 
**Quota** | Pointer to **float32** | mailbox quota | [optional] 
**SenderAcl** | Pointer to **map[string]interface{}** | list of allowed send from addresses | [optional] 
**SogoAccess** | Pointer to **bool** | is access to SOGo webmail active or not | [optional] 

## Methods

### NewUpdateMailboxRequestAttr

`func NewUpdateMailboxRequestAttr() *UpdateMailboxRequestAttr`

NewUpdateMailboxRequestAttr instantiates a new UpdateMailboxRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxRequestAttrWithDefaults

`func NewUpdateMailboxRequestAttrWithDefaults() *UpdateMailboxRequestAttr`

NewUpdateMailboxRequestAttrWithDefaults instantiates a new UpdateMailboxRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateMailboxRequestAttr) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateMailboxRequestAttr) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateMailboxRequestAttr) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateMailboxRequestAttr) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetForcePwUpdate

`func (o *UpdateMailboxRequestAttr) GetForcePwUpdate() bool`

GetForcePwUpdate returns the ForcePwUpdate field if non-nil, zero value otherwise.

### GetForcePwUpdateOk

`func (o *UpdateMailboxRequestAttr) GetForcePwUpdateOk() (*bool, bool)`

GetForcePwUpdateOk returns a tuple with the ForcePwUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePwUpdate

`func (o *UpdateMailboxRequestAttr) SetForcePwUpdate(v bool)`

SetForcePwUpdate sets ForcePwUpdate field to given value.

### HasForcePwUpdate

`func (o *UpdateMailboxRequestAttr) HasForcePwUpdate() bool`

HasForcePwUpdate returns a boolean if a field has been set.

### GetName

`func (o *UpdateMailboxRequestAttr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMailboxRequestAttr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMailboxRequestAttr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMailboxRequestAttr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword2

`func (o *UpdateMailboxRequestAttr) GetPassword2() string`

GetPassword2 returns the Password2 field if non-nil, zero value otherwise.

### GetPassword2Ok

`func (o *UpdateMailboxRequestAttr) GetPassword2Ok() (*string, bool)`

GetPassword2Ok returns a tuple with the Password2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword2

`func (o *UpdateMailboxRequestAttr) SetPassword2(v string)`

SetPassword2 sets Password2 field to given value.

### HasPassword2

`func (o *UpdateMailboxRequestAttr) HasPassword2() bool`

HasPassword2 returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateMailboxRequestAttr) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMailboxRequestAttr) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMailboxRequestAttr) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMailboxRequestAttr) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetQuota

`func (o *UpdateMailboxRequestAttr) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateMailboxRequestAttr) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateMailboxRequestAttr) SetQuota(v float32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateMailboxRequestAttr) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetSenderAcl

`func (o *UpdateMailboxRequestAttr) GetSenderAcl() map[string]interface{}`

GetSenderAcl returns the SenderAcl field if non-nil, zero value otherwise.

### GetSenderAclOk

`func (o *UpdateMailboxRequestAttr) GetSenderAclOk() (*map[string]interface{}, bool)`

GetSenderAclOk returns a tuple with the SenderAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAcl

`func (o *UpdateMailboxRequestAttr) SetSenderAcl(v map[string]interface{})`

SetSenderAcl sets SenderAcl field to given value.

### HasSenderAcl

`func (o *UpdateMailboxRequestAttr) HasSenderAcl() bool`

HasSenderAcl returns a boolean if a field has been set.

### GetSogoAccess

`func (o *UpdateMailboxRequestAttr) GetSogoAccess() bool`

GetSogoAccess returns the SogoAccess field if non-nil, zero value otherwise.

### GetSogoAccessOk

`func (o *UpdateMailboxRequestAttr) GetSogoAccessOk() (*bool, bool)`

GetSogoAccessOk returns a tuple with the SogoAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSogoAccess

`func (o *UpdateMailboxRequestAttr) SetSogoAccess(v bool)`

SetSogoAccess sets SogoAccess field to given value.

### HasSogoAccess

`func (o *UpdateMailboxRequestAttr) HasSogoAccess() bool`

HasSogoAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


