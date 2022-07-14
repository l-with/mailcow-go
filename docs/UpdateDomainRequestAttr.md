# UpdateDomainRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is domain active or not | [optional] 
**Aliases** | Pointer to **float32** | limit count of aliases associated with this domain | [optional] 
**Backupmx** | Pointer to **bool** | relay domain or not | [optional] 
**Defquota** | Pointer to **float32** | predefined mailbox quota in &#x60;add mailbox&#x60; form | [optional] 
**Description** | Pointer to **string** | Description of domain | [optional] 
**Gal** | Pointer to **bool** | is domain global address list active or not, it enables shared contacts accross domain in SOGo webmail | [optional] 
**Mailboxes** | Pointer to **float32** | limit count of mailboxes associated with this domain | [optional] 
**Maxquota** | Pointer to **float32** | maximum quota per mailbox | [optional] 
**Quota** | Pointer to **float32** | maximum quota for this domain (for all mailboxes in sum) | [optional] 
**RelayAllRecipients** | Pointer to **bool** | if not, them you have to create \&quot;dummy\&quot; mailbox for each address to relay | [optional] 
**RelayUnknownOnly** | Pointer to **bool** | Relay non-existing mailboxes only. Existing mailboxes will be delivered locally. | [optional] 
**Relayhost** | Pointer to **float32** | id of relayhost | [optional] 
**RlFrame** | Pointer to **string** |  | [optional] 
**RlValue** | Pointer to **float32** | rate limit value | [optional] 
**Tags** | Pointer to **[]string** | tags for this Domain | [optional] 

## Methods

### NewUpdateDomainRequestAttr

`func NewUpdateDomainRequestAttr() *UpdateDomainRequestAttr`

NewUpdateDomainRequestAttr instantiates a new UpdateDomainRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainRequestAttrWithDefaults

`func NewUpdateDomainRequestAttrWithDefaults() *UpdateDomainRequestAttr`

NewUpdateDomainRequestAttrWithDefaults instantiates a new UpdateDomainRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateDomainRequestAttr) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateDomainRequestAttr) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateDomainRequestAttr) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateDomainRequestAttr) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAliases

`func (o *UpdateDomainRequestAttr) GetAliases() float32`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *UpdateDomainRequestAttr) GetAliasesOk() (*float32, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *UpdateDomainRequestAttr) SetAliases(v float32)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *UpdateDomainRequestAttr) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBackupmx

`func (o *UpdateDomainRequestAttr) GetBackupmx() bool`

GetBackupmx returns the Backupmx field if non-nil, zero value otherwise.

### GetBackupmxOk

`func (o *UpdateDomainRequestAttr) GetBackupmxOk() (*bool, bool)`

GetBackupmxOk returns a tuple with the Backupmx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupmx

`func (o *UpdateDomainRequestAttr) SetBackupmx(v bool)`

SetBackupmx sets Backupmx field to given value.

### HasBackupmx

`func (o *UpdateDomainRequestAttr) HasBackupmx() bool`

HasBackupmx returns a boolean if a field has been set.

### GetDefquota

`func (o *UpdateDomainRequestAttr) GetDefquota() float32`

GetDefquota returns the Defquota field if non-nil, zero value otherwise.

### GetDefquotaOk

`func (o *UpdateDomainRequestAttr) GetDefquotaOk() (*float32, bool)`

GetDefquotaOk returns a tuple with the Defquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefquota

`func (o *UpdateDomainRequestAttr) SetDefquota(v float32)`

SetDefquota sets Defquota field to given value.

### HasDefquota

`func (o *UpdateDomainRequestAttr) HasDefquota() bool`

HasDefquota returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDomainRequestAttr) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDomainRequestAttr) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDomainRequestAttr) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDomainRequestAttr) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGal

`func (o *UpdateDomainRequestAttr) GetGal() bool`

GetGal returns the Gal field if non-nil, zero value otherwise.

### GetGalOk

`func (o *UpdateDomainRequestAttr) GetGalOk() (*bool, bool)`

GetGalOk returns a tuple with the Gal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGal

`func (o *UpdateDomainRequestAttr) SetGal(v bool)`

SetGal sets Gal field to given value.

### HasGal

`func (o *UpdateDomainRequestAttr) HasGal() bool`

HasGal returns a boolean if a field has been set.

### GetMailboxes

`func (o *UpdateDomainRequestAttr) GetMailboxes() float32`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *UpdateDomainRequestAttr) GetMailboxesOk() (*float32, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *UpdateDomainRequestAttr) SetMailboxes(v float32)`

SetMailboxes sets Mailboxes field to given value.

### HasMailboxes

`func (o *UpdateDomainRequestAttr) HasMailboxes() bool`

HasMailboxes returns a boolean if a field has been set.

### GetMaxquota

`func (o *UpdateDomainRequestAttr) GetMaxquota() float32`

GetMaxquota returns the Maxquota field if non-nil, zero value otherwise.

### GetMaxquotaOk

`func (o *UpdateDomainRequestAttr) GetMaxquotaOk() (*float32, bool)`

GetMaxquotaOk returns a tuple with the Maxquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxquota

`func (o *UpdateDomainRequestAttr) SetMaxquota(v float32)`

SetMaxquota sets Maxquota field to given value.

### HasMaxquota

`func (o *UpdateDomainRequestAttr) HasMaxquota() bool`

HasMaxquota returns a boolean if a field has been set.

### GetQuota

`func (o *UpdateDomainRequestAttr) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateDomainRequestAttr) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateDomainRequestAttr) SetQuota(v float32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateDomainRequestAttr) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetRelayAllRecipients

`func (o *UpdateDomainRequestAttr) GetRelayAllRecipients() bool`

GetRelayAllRecipients returns the RelayAllRecipients field if non-nil, zero value otherwise.

### GetRelayAllRecipientsOk

`func (o *UpdateDomainRequestAttr) GetRelayAllRecipientsOk() (*bool, bool)`

GetRelayAllRecipientsOk returns a tuple with the RelayAllRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAllRecipients

`func (o *UpdateDomainRequestAttr) SetRelayAllRecipients(v bool)`

SetRelayAllRecipients sets RelayAllRecipients field to given value.

### HasRelayAllRecipients

`func (o *UpdateDomainRequestAttr) HasRelayAllRecipients() bool`

HasRelayAllRecipients returns a boolean if a field has been set.

### GetRelayUnknownOnly

`func (o *UpdateDomainRequestAttr) GetRelayUnknownOnly() bool`

GetRelayUnknownOnly returns the RelayUnknownOnly field if non-nil, zero value otherwise.

### GetRelayUnknownOnlyOk

`func (o *UpdateDomainRequestAttr) GetRelayUnknownOnlyOk() (*bool, bool)`

GetRelayUnknownOnlyOk returns a tuple with the RelayUnknownOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUnknownOnly

`func (o *UpdateDomainRequestAttr) SetRelayUnknownOnly(v bool)`

SetRelayUnknownOnly sets RelayUnknownOnly field to given value.

### HasRelayUnknownOnly

`func (o *UpdateDomainRequestAttr) HasRelayUnknownOnly() bool`

HasRelayUnknownOnly returns a boolean if a field has been set.

### GetRelayhost

`func (o *UpdateDomainRequestAttr) GetRelayhost() float32`

GetRelayhost returns the Relayhost field if non-nil, zero value otherwise.

### GetRelayhostOk

`func (o *UpdateDomainRequestAttr) GetRelayhostOk() (*float32, bool)`

GetRelayhostOk returns a tuple with the Relayhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayhost

`func (o *UpdateDomainRequestAttr) SetRelayhost(v float32)`

SetRelayhost sets Relayhost field to given value.

### HasRelayhost

`func (o *UpdateDomainRequestAttr) HasRelayhost() bool`

HasRelayhost returns a boolean if a field has been set.

### GetRlFrame

`func (o *UpdateDomainRequestAttr) GetRlFrame() string`

GetRlFrame returns the RlFrame field if non-nil, zero value otherwise.

### GetRlFrameOk

`func (o *UpdateDomainRequestAttr) GetRlFrameOk() (*string, bool)`

GetRlFrameOk returns a tuple with the RlFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlFrame

`func (o *UpdateDomainRequestAttr) SetRlFrame(v string)`

SetRlFrame sets RlFrame field to given value.

### HasRlFrame

`func (o *UpdateDomainRequestAttr) HasRlFrame() bool`

HasRlFrame returns a boolean if a field has been set.

### GetRlValue

`func (o *UpdateDomainRequestAttr) GetRlValue() float32`

GetRlValue returns the RlValue field if non-nil, zero value otherwise.

### GetRlValueOk

`func (o *UpdateDomainRequestAttr) GetRlValueOk() (*float32, bool)`

GetRlValueOk returns a tuple with the RlValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlValue

`func (o *UpdateDomainRequestAttr) SetRlValue(v float32)`

SetRlValue sets RlValue field to given value.

### HasRlValue

`func (o *UpdateDomainRequestAttr) HasRlValue() bool`

HasRlValue returns a boolean if a field has been set.

### GetTags

`func (o *UpdateDomainRequestAttr) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDomainRequestAttr) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDomainRequestAttr) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDomainRequestAttr) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


