# CreateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is domain active or not | [optional] 
**Aliases** | Pointer to **float32** | limit count of aliases associated with this domain | [optional] 
**Backupmx** | Pointer to **bool** | relay domain or not | [optional] 
**Defquota** | Pointer to **float32** | predefined mailbox quota in &#x60;add mailbox&#x60; form | [optional] 
**Description** | Pointer to **string** | Description of domain | [optional] 
**Domain** | Pointer to **string** | Fully qualified domain name | [optional] 
**Gal** | Pointer to **bool** | is domain global address list active or not, it enables shared contacts accross domain in SOGo webmail | [optional] 
**Mailboxes** | Pointer to **float32** | limit count of mailboxes associated with this domain | [optional] 
**Maxquota** | Pointer to **float32** | maximum quota per mailbox | [optional] 
**Quota** | Pointer to **float32** | maximum quota for this domain (for all mailboxes in sum) | [optional] 
**RestartSogo** | Pointer to **float32** | restart SOGo to activate the domain in SOGo | [optional] 
**RelayAllRecipients** | Pointer to **bool** | if not, them you have to create \&quot;dummy\&quot; mailbox for each address to relay | [optional] 
**RelayUnknownOnly** | Pointer to **bool** | Relay non-existing mailboxes only. Existing mailboxes will be delivered locally. | [optional] 
**RlFrame** | Pointer to **string** |  | [optional] 
**RlValue** | Pointer to **float32** | rate limit value | [optional] 
**Tags** | Pointer to **[]string** | tags for this Domain | [optional] 

## Methods

### NewCreateDomainRequest

`func NewCreateDomainRequest() *CreateDomainRequest`

NewCreateDomainRequest instantiates a new CreateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainRequestWithDefaults

`func NewCreateDomainRequestWithDefaults() *CreateDomainRequest`

NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateDomainRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateDomainRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateDomainRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateDomainRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAliases

`func (o *CreateDomainRequest) GetAliases() float32`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CreateDomainRequest) GetAliasesOk() (*float32, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CreateDomainRequest) SetAliases(v float32)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CreateDomainRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBackupmx

`func (o *CreateDomainRequest) GetBackupmx() bool`

GetBackupmx returns the Backupmx field if non-nil, zero value otherwise.

### GetBackupmxOk

`func (o *CreateDomainRequest) GetBackupmxOk() (*bool, bool)`

GetBackupmxOk returns a tuple with the Backupmx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupmx

`func (o *CreateDomainRequest) SetBackupmx(v bool)`

SetBackupmx sets Backupmx field to given value.

### HasBackupmx

`func (o *CreateDomainRequest) HasBackupmx() bool`

HasBackupmx returns a boolean if a field has been set.

### GetDefquota

`func (o *CreateDomainRequest) GetDefquota() float32`

GetDefquota returns the Defquota field if non-nil, zero value otherwise.

### GetDefquotaOk

`func (o *CreateDomainRequest) GetDefquotaOk() (*float32, bool)`

GetDefquotaOk returns a tuple with the Defquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefquota

`func (o *CreateDomainRequest) SetDefquota(v float32)`

SetDefquota sets Defquota field to given value.

### HasDefquota

`func (o *CreateDomainRequest) HasDefquota() bool`

HasDefquota returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDomainRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDomainRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDomainRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDomainRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *CreateDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateDomainRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGal

`func (o *CreateDomainRequest) GetGal() bool`

GetGal returns the Gal field if non-nil, zero value otherwise.

### GetGalOk

`func (o *CreateDomainRequest) GetGalOk() (*bool, bool)`

GetGalOk returns a tuple with the Gal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGal

`func (o *CreateDomainRequest) SetGal(v bool)`

SetGal sets Gal field to given value.

### HasGal

`func (o *CreateDomainRequest) HasGal() bool`

HasGal returns a boolean if a field has been set.

### GetMailboxes

`func (o *CreateDomainRequest) GetMailboxes() float32`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *CreateDomainRequest) GetMailboxesOk() (*float32, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *CreateDomainRequest) SetMailboxes(v float32)`

SetMailboxes sets Mailboxes field to given value.

### HasMailboxes

`func (o *CreateDomainRequest) HasMailboxes() bool`

HasMailboxes returns a boolean if a field has been set.

### GetMaxquota

`func (o *CreateDomainRequest) GetMaxquota() float32`

GetMaxquota returns the Maxquota field if non-nil, zero value otherwise.

### GetMaxquotaOk

`func (o *CreateDomainRequest) GetMaxquotaOk() (*float32, bool)`

GetMaxquotaOk returns a tuple with the Maxquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxquota

`func (o *CreateDomainRequest) SetMaxquota(v float32)`

SetMaxquota sets Maxquota field to given value.

### HasMaxquota

`func (o *CreateDomainRequest) HasMaxquota() bool`

HasMaxquota returns a boolean if a field has been set.

### GetQuota

`func (o *CreateDomainRequest) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateDomainRequest) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateDomainRequest) SetQuota(v float32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateDomainRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetRestartSogo

`func (o *CreateDomainRequest) GetRestartSogo() float32`

GetRestartSogo returns the RestartSogo field if non-nil, zero value otherwise.

### GetRestartSogoOk

`func (o *CreateDomainRequest) GetRestartSogoOk() (*float32, bool)`

GetRestartSogoOk returns a tuple with the RestartSogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSogo

`func (o *CreateDomainRequest) SetRestartSogo(v float32)`

SetRestartSogo sets RestartSogo field to given value.

### HasRestartSogo

`func (o *CreateDomainRequest) HasRestartSogo() bool`

HasRestartSogo returns a boolean if a field has been set.

### GetRelayAllRecipients

`func (o *CreateDomainRequest) GetRelayAllRecipients() bool`

GetRelayAllRecipients returns the RelayAllRecipients field if non-nil, zero value otherwise.

### GetRelayAllRecipientsOk

`func (o *CreateDomainRequest) GetRelayAllRecipientsOk() (*bool, bool)`

GetRelayAllRecipientsOk returns a tuple with the RelayAllRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAllRecipients

`func (o *CreateDomainRequest) SetRelayAllRecipients(v bool)`

SetRelayAllRecipients sets RelayAllRecipients field to given value.

### HasRelayAllRecipients

`func (o *CreateDomainRequest) HasRelayAllRecipients() bool`

HasRelayAllRecipients returns a boolean if a field has been set.

### GetRelayUnknownOnly

`func (o *CreateDomainRequest) GetRelayUnknownOnly() bool`

GetRelayUnknownOnly returns the RelayUnknownOnly field if non-nil, zero value otherwise.

### GetRelayUnknownOnlyOk

`func (o *CreateDomainRequest) GetRelayUnknownOnlyOk() (*bool, bool)`

GetRelayUnknownOnlyOk returns a tuple with the RelayUnknownOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUnknownOnly

`func (o *CreateDomainRequest) SetRelayUnknownOnly(v bool)`

SetRelayUnknownOnly sets RelayUnknownOnly field to given value.

### HasRelayUnknownOnly

`func (o *CreateDomainRequest) HasRelayUnknownOnly() bool`

HasRelayUnknownOnly returns a boolean if a field has been set.

### GetRlFrame

`func (o *CreateDomainRequest) GetRlFrame() string`

GetRlFrame returns the RlFrame field if non-nil, zero value otherwise.

### GetRlFrameOk

`func (o *CreateDomainRequest) GetRlFrameOk() (*string, bool)`

GetRlFrameOk returns a tuple with the RlFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlFrame

`func (o *CreateDomainRequest) SetRlFrame(v string)`

SetRlFrame sets RlFrame field to given value.

### HasRlFrame

`func (o *CreateDomainRequest) HasRlFrame() bool`

HasRlFrame returns a boolean if a field has been set.

### GetRlValue

`func (o *CreateDomainRequest) GetRlValue() float32`

GetRlValue returns the RlValue field if non-nil, zero value otherwise.

### GetRlValueOk

`func (o *CreateDomainRequest) GetRlValueOk() (*float32, bool)`

GetRlValueOk returns a tuple with the RlValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlValue

`func (o *CreateDomainRequest) SetRlValue(v float32)`

SetRlValue sets RlValue field to given value.

### HasRlValue

`func (o *CreateDomainRequest) HasRlValue() bool`

HasRlValue returns a boolean if a field has been set.

### GetTags

`func (o *CreateDomainRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDomainRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDomainRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDomainRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


