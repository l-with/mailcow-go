# CreateSyncJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **string** | your local mailcow mailbox | [optional] 
**Host1** | Pointer to **string** | the smtp server where mails should be synced from | [optional] 
**Port1** | Pointer to **string** | the smtp port of the target mail server | [optional] 
**Password** | Pointer to **string** | the password of the mailbox | [optional] 
**Enc1** | Pointer to **string** | the encryption method used to connect to the mailserver | [optional] 
**MinsInternal** | Pointer to **float32** | the interval in which messages should be syned | [optional] 
**Subfolder2** | Pointer to **string** | sync into subfolder on destination (empty &#x3D; do not use subfolder) | [optional] 
**Maxage** | Pointer to **float32** | only sync messages up to this age in days | [optional] 
**Maxbytespersecond** | Pointer to **float32** | max speed transfer limit for the sync | [optional] 
**Timeout1** | Pointer to **float32** | timeout for connection to remote host | [optional] 
**Timeout2** | Pointer to **float32** | timeout for connection to local host | [optional] 
**Exclude** | Pointer to **string** | exclude objects (regex) | [optional] 
**CustomParams** | Pointer to **string** | custom parameters | [optional] 
**Delete2duplicates** | Pointer to **bool** | delete duplicates on destination (--delete2duplicates) | [optional] 
**Delete1** | Pointer to **bool** | delete from source when completed (--delete1) | [optional] 
**Delete2** | Pointer to **bool** | delete messages on destination that are not on source (--delete2) | [optional] 
**Automap** | Pointer to **bool** | try to automap folders (\&quot;Sent items\&quot;, \&quot;Sent\&quot; &#x3D;&gt; \&quot;Sent\&quot; etc.) (--automap) | [optional] 
**Skipcrossduplicates** | Pointer to **bool** | skip duplicate messages across folders (first come, first serve) (--skipcrossduplicates) | [optional] 
**Subscribeall** | Pointer to **bool** | subscribe all folders (--subscribeall) | [optional] 
**Active** | Pointer to **bool** | enables or disables the sync job | [optional] 

## Methods

### NewCreateSyncJobRequest

`func NewCreateSyncJobRequest() *CreateSyncJobRequest`

NewCreateSyncJobRequest instantiates a new CreateSyncJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSyncJobRequestWithDefaults

`func NewCreateSyncJobRequestWithDefaults() *CreateSyncJobRequest`

NewCreateSyncJobRequestWithDefaults instantiates a new CreateSyncJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *CreateSyncJobRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateSyncJobRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateSyncJobRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateSyncJobRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetHost1

`func (o *CreateSyncJobRequest) GetHost1() string`

GetHost1 returns the Host1 field if non-nil, zero value otherwise.

### GetHost1Ok

`func (o *CreateSyncJobRequest) GetHost1Ok() (*string, bool)`

GetHost1Ok returns a tuple with the Host1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost1

`func (o *CreateSyncJobRequest) SetHost1(v string)`

SetHost1 sets Host1 field to given value.

### HasHost1

`func (o *CreateSyncJobRequest) HasHost1() bool`

HasHost1 returns a boolean if a field has been set.

### GetPort1

`func (o *CreateSyncJobRequest) GetPort1() string`

GetPort1 returns the Port1 field if non-nil, zero value otherwise.

### GetPort1Ok

`func (o *CreateSyncJobRequest) GetPort1Ok() (*string, bool)`

GetPort1Ok returns a tuple with the Port1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort1

`func (o *CreateSyncJobRequest) SetPort1(v string)`

SetPort1 sets Port1 field to given value.

### HasPort1

`func (o *CreateSyncJobRequest) HasPort1() bool`

HasPort1 returns a boolean if a field has been set.

### GetPassword

`func (o *CreateSyncJobRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSyncJobRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSyncJobRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSyncJobRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEnc1

`func (o *CreateSyncJobRequest) GetEnc1() string`

GetEnc1 returns the Enc1 field if non-nil, zero value otherwise.

### GetEnc1Ok

`func (o *CreateSyncJobRequest) GetEnc1Ok() (*string, bool)`

GetEnc1Ok returns a tuple with the Enc1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnc1

`func (o *CreateSyncJobRequest) SetEnc1(v string)`

SetEnc1 sets Enc1 field to given value.

### HasEnc1

`func (o *CreateSyncJobRequest) HasEnc1() bool`

HasEnc1 returns a boolean if a field has been set.

### GetMinsInternal

`func (o *CreateSyncJobRequest) GetMinsInternal() float32`

GetMinsInternal returns the MinsInternal field if non-nil, zero value otherwise.

### GetMinsInternalOk

`func (o *CreateSyncJobRequest) GetMinsInternalOk() (*float32, bool)`

GetMinsInternalOk returns a tuple with the MinsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinsInternal

`func (o *CreateSyncJobRequest) SetMinsInternal(v float32)`

SetMinsInternal sets MinsInternal field to given value.

### HasMinsInternal

`func (o *CreateSyncJobRequest) HasMinsInternal() bool`

HasMinsInternal returns a boolean if a field has been set.

### GetSubfolder2

`func (o *CreateSyncJobRequest) GetSubfolder2() string`

GetSubfolder2 returns the Subfolder2 field if non-nil, zero value otherwise.

### GetSubfolder2Ok

`func (o *CreateSyncJobRequest) GetSubfolder2Ok() (*string, bool)`

GetSubfolder2Ok returns a tuple with the Subfolder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolder2

`func (o *CreateSyncJobRequest) SetSubfolder2(v string)`

SetSubfolder2 sets Subfolder2 field to given value.

### HasSubfolder2

`func (o *CreateSyncJobRequest) HasSubfolder2() bool`

HasSubfolder2 returns a boolean if a field has been set.

### GetMaxage

`func (o *CreateSyncJobRequest) GetMaxage() float32`

GetMaxage returns the Maxage field if non-nil, zero value otherwise.

### GetMaxageOk

`func (o *CreateSyncJobRequest) GetMaxageOk() (*float32, bool)`

GetMaxageOk returns a tuple with the Maxage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxage

`func (o *CreateSyncJobRequest) SetMaxage(v float32)`

SetMaxage sets Maxage field to given value.

### HasMaxage

`func (o *CreateSyncJobRequest) HasMaxage() bool`

HasMaxage returns a boolean if a field has been set.

### GetMaxbytespersecond

`func (o *CreateSyncJobRequest) GetMaxbytespersecond() float32`

GetMaxbytespersecond returns the Maxbytespersecond field if non-nil, zero value otherwise.

### GetMaxbytespersecondOk

`func (o *CreateSyncJobRequest) GetMaxbytespersecondOk() (*float32, bool)`

GetMaxbytespersecondOk returns a tuple with the Maxbytespersecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbytespersecond

`func (o *CreateSyncJobRequest) SetMaxbytespersecond(v float32)`

SetMaxbytespersecond sets Maxbytespersecond field to given value.

### HasMaxbytespersecond

`func (o *CreateSyncJobRequest) HasMaxbytespersecond() bool`

HasMaxbytespersecond returns a boolean if a field has been set.

### GetTimeout1

`func (o *CreateSyncJobRequest) GetTimeout1() float32`

GetTimeout1 returns the Timeout1 field if non-nil, zero value otherwise.

### GetTimeout1Ok

`func (o *CreateSyncJobRequest) GetTimeout1Ok() (*float32, bool)`

GetTimeout1Ok returns a tuple with the Timeout1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout1

`func (o *CreateSyncJobRequest) SetTimeout1(v float32)`

SetTimeout1 sets Timeout1 field to given value.

### HasTimeout1

`func (o *CreateSyncJobRequest) HasTimeout1() bool`

HasTimeout1 returns a boolean if a field has been set.

### GetTimeout2

`func (o *CreateSyncJobRequest) GetTimeout2() float32`

GetTimeout2 returns the Timeout2 field if non-nil, zero value otherwise.

### GetTimeout2Ok

`func (o *CreateSyncJobRequest) GetTimeout2Ok() (*float32, bool)`

GetTimeout2Ok returns a tuple with the Timeout2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout2

`func (o *CreateSyncJobRequest) SetTimeout2(v float32)`

SetTimeout2 sets Timeout2 field to given value.

### HasTimeout2

`func (o *CreateSyncJobRequest) HasTimeout2() bool`

HasTimeout2 returns a boolean if a field has been set.

### GetExclude

`func (o *CreateSyncJobRequest) GetExclude() string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CreateSyncJobRequest) GetExcludeOk() (*string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CreateSyncJobRequest) SetExclude(v string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CreateSyncJobRequest) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetCustomParams

`func (o *CreateSyncJobRequest) GetCustomParams() string`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *CreateSyncJobRequest) GetCustomParamsOk() (*string, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *CreateSyncJobRequest) SetCustomParams(v string)`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *CreateSyncJobRequest) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetDelete2duplicates

`func (o *CreateSyncJobRequest) GetDelete2duplicates() bool`

GetDelete2duplicates returns the Delete2duplicates field if non-nil, zero value otherwise.

### GetDelete2duplicatesOk

`func (o *CreateSyncJobRequest) GetDelete2duplicatesOk() (*bool, bool)`

GetDelete2duplicatesOk returns a tuple with the Delete2duplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete2duplicates

`func (o *CreateSyncJobRequest) SetDelete2duplicates(v bool)`

SetDelete2duplicates sets Delete2duplicates field to given value.

### HasDelete2duplicates

`func (o *CreateSyncJobRequest) HasDelete2duplicates() bool`

HasDelete2duplicates returns a boolean if a field has been set.

### GetDelete1

`func (o *CreateSyncJobRequest) GetDelete1() bool`

GetDelete1 returns the Delete1 field if non-nil, zero value otherwise.

### GetDelete1Ok

`func (o *CreateSyncJobRequest) GetDelete1Ok() (*bool, bool)`

GetDelete1Ok returns a tuple with the Delete1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete1

`func (o *CreateSyncJobRequest) SetDelete1(v bool)`

SetDelete1 sets Delete1 field to given value.

### HasDelete1

`func (o *CreateSyncJobRequest) HasDelete1() bool`

HasDelete1 returns a boolean if a field has been set.

### GetDelete2

`func (o *CreateSyncJobRequest) GetDelete2() bool`

GetDelete2 returns the Delete2 field if non-nil, zero value otherwise.

### GetDelete2Ok

`func (o *CreateSyncJobRequest) GetDelete2Ok() (*bool, bool)`

GetDelete2Ok returns a tuple with the Delete2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete2

`func (o *CreateSyncJobRequest) SetDelete2(v bool)`

SetDelete2 sets Delete2 field to given value.

### HasDelete2

`func (o *CreateSyncJobRequest) HasDelete2() bool`

HasDelete2 returns a boolean if a field has been set.

### GetAutomap

`func (o *CreateSyncJobRequest) GetAutomap() bool`

GetAutomap returns the Automap field if non-nil, zero value otherwise.

### GetAutomapOk

`func (o *CreateSyncJobRequest) GetAutomapOk() (*bool, bool)`

GetAutomapOk returns a tuple with the Automap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomap

`func (o *CreateSyncJobRequest) SetAutomap(v bool)`

SetAutomap sets Automap field to given value.

### HasAutomap

`func (o *CreateSyncJobRequest) HasAutomap() bool`

HasAutomap returns a boolean if a field has been set.

### GetSkipcrossduplicates

`func (o *CreateSyncJobRequest) GetSkipcrossduplicates() bool`

GetSkipcrossduplicates returns the Skipcrossduplicates field if non-nil, zero value otherwise.

### GetSkipcrossduplicatesOk

`func (o *CreateSyncJobRequest) GetSkipcrossduplicatesOk() (*bool, bool)`

GetSkipcrossduplicatesOk returns a tuple with the Skipcrossduplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipcrossduplicates

`func (o *CreateSyncJobRequest) SetSkipcrossduplicates(v bool)`

SetSkipcrossduplicates sets Skipcrossduplicates field to given value.

### HasSkipcrossduplicates

`func (o *CreateSyncJobRequest) HasSkipcrossduplicates() bool`

HasSkipcrossduplicates returns a boolean if a field has been set.

### GetSubscribeall

`func (o *CreateSyncJobRequest) GetSubscribeall() bool`

GetSubscribeall returns the Subscribeall field if non-nil, zero value otherwise.

### GetSubscribeallOk

`func (o *CreateSyncJobRequest) GetSubscribeallOk() (*bool, bool)`

GetSubscribeallOk returns a tuple with the Subscribeall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeall

`func (o *CreateSyncJobRequest) SetSubscribeall(v bool)`

SetSubscribeall sets Subscribeall field to given value.

### HasSubscribeall

`func (o *CreateSyncJobRequest) HasSubscribeall() bool`

HasSubscribeall returns a boolean if a field has been set.

### GetActive

`func (o *CreateSyncJobRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateSyncJobRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateSyncJobRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateSyncJobRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


