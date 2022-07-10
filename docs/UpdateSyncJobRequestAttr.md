# UpdateSyncJobRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Is sync job active | [optional] 
**Automap** | Pointer to **bool** | Try to automap folders (\&quot;Sent items\&quot;, \&quot;Sent\&quot; &#x3D;&gt; \&quot;Sent\&quot; etc.) | [optional] 
**CustomParams** | Pointer to **string** | Custom parameters passed to imapsync command | [optional] 
**Delete1** | Pointer to **bool** | Delete from source when completed | [optional] 
**Delete2** | Pointer to **bool** | Delete messages on destination that are not on source | [optional] 
**Delete2duplicates** | Pointer to **bool** | Delete duplicates on destination | [optional] 
**Enc1** | Pointer to **string** | Encryption | [optional] 
**Exclude** | Pointer to **string** | Exclude objects (regex) | [optional] 
**Host1** | Pointer to **string** | Hostname | [optional] 
**Maxage** | Pointer to **float32** | Maximum age of messages in days that will be polled from remote (0 &#x3D; ignore age) | [optional] 
**Maxbytespersecond** | Pointer to **float32** | Max. bytes per second (0 &#x3D; unlimited) | [optional] 
**MinsInterval** | Pointer to **float32** | Interval (min) | [optional] 
**Password1** | Pointer to **string** | Password | [optional] 
**Port1** | Pointer to **string** | Port | [optional] 
**Skipcrossduplicates** | Pointer to **bool** | Skip duplicate messages across folders (first come, first serve) | [optional] 
**Subfolder2** | Pointer to **string** | Sync into subfolder on destination (empty &#x3D; do not use subfolder) | [optional] 
**Subscribeall** | Pointer to **bool** | Subscribe all folders | [optional] 
**Timeout1** | Pointer to **float32** | Timeout for connection to remote host | [optional] 
**Timeout2** | Pointer to **float32** | Timeout for connection to local host | [optional] 
**User1** | Pointer to **string** | Username | [optional] 

## Methods

### NewUpdateSyncJobRequestAttr

`func NewUpdateSyncJobRequestAttr() *UpdateSyncJobRequestAttr`

NewUpdateSyncJobRequestAttr instantiates a new UpdateSyncJobRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSyncJobRequestAttrWithDefaults

`func NewUpdateSyncJobRequestAttrWithDefaults() *UpdateSyncJobRequestAttr`

NewUpdateSyncJobRequestAttrWithDefaults instantiates a new UpdateSyncJobRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateSyncJobRequestAttr) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSyncJobRequestAttr) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSyncJobRequestAttr) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateSyncJobRequestAttr) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAutomap

`func (o *UpdateSyncJobRequestAttr) GetAutomap() bool`

GetAutomap returns the Automap field if non-nil, zero value otherwise.

### GetAutomapOk

`func (o *UpdateSyncJobRequestAttr) GetAutomapOk() (*bool, bool)`

GetAutomapOk returns a tuple with the Automap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomap

`func (o *UpdateSyncJobRequestAttr) SetAutomap(v bool)`

SetAutomap sets Automap field to given value.

### HasAutomap

`func (o *UpdateSyncJobRequestAttr) HasAutomap() bool`

HasAutomap returns a boolean if a field has been set.

### GetCustomParams

`func (o *UpdateSyncJobRequestAttr) GetCustomParams() string`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *UpdateSyncJobRequestAttr) GetCustomParamsOk() (*string, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *UpdateSyncJobRequestAttr) SetCustomParams(v string)`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *UpdateSyncJobRequestAttr) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetDelete1

`func (o *UpdateSyncJobRequestAttr) GetDelete1() bool`

GetDelete1 returns the Delete1 field if non-nil, zero value otherwise.

### GetDelete1Ok

`func (o *UpdateSyncJobRequestAttr) GetDelete1Ok() (*bool, bool)`

GetDelete1Ok returns a tuple with the Delete1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete1

`func (o *UpdateSyncJobRequestAttr) SetDelete1(v bool)`

SetDelete1 sets Delete1 field to given value.

### HasDelete1

`func (o *UpdateSyncJobRequestAttr) HasDelete1() bool`

HasDelete1 returns a boolean if a field has been set.

### GetDelete2

`func (o *UpdateSyncJobRequestAttr) GetDelete2() bool`

GetDelete2 returns the Delete2 field if non-nil, zero value otherwise.

### GetDelete2Ok

`func (o *UpdateSyncJobRequestAttr) GetDelete2Ok() (*bool, bool)`

GetDelete2Ok returns a tuple with the Delete2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete2

`func (o *UpdateSyncJobRequestAttr) SetDelete2(v bool)`

SetDelete2 sets Delete2 field to given value.

### HasDelete2

`func (o *UpdateSyncJobRequestAttr) HasDelete2() bool`

HasDelete2 returns a boolean if a field has been set.

### GetDelete2duplicates

`func (o *UpdateSyncJobRequestAttr) GetDelete2duplicates() bool`

GetDelete2duplicates returns the Delete2duplicates field if non-nil, zero value otherwise.

### GetDelete2duplicatesOk

`func (o *UpdateSyncJobRequestAttr) GetDelete2duplicatesOk() (*bool, bool)`

GetDelete2duplicatesOk returns a tuple with the Delete2duplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete2duplicates

`func (o *UpdateSyncJobRequestAttr) SetDelete2duplicates(v bool)`

SetDelete2duplicates sets Delete2duplicates field to given value.

### HasDelete2duplicates

`func (o *UpdateSyncJobRequestAttr) HasDelete2duplicates() bool`

HasDelete2duplicates returns a boolean if a field has been set.

### GetEnc1

`func (o *UpdateSyncJobRequestAttr) GetEnc1() string`

GetEnc1 returns the Enc1 field if non-nil, zero value otherwise.

### GetEnc1Ok

`func (o *UpdateSyncJobRequestAttr) GetEnc1Ok() (*string, bool)`

GetEnc1Ok returns a tuple with the Enc1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnc1

`func (o *UpdateSyncJobRequestAttr) SetEnc1(v string)`

SetEnc1 sets Enc1 field to given value.

### HasEnc1

`func (o *UpdateSyncJobRequestAttr) HasEnc1() bool`

HasEnc1 returns a boolean if a field has been set.

### GetExclude

`func (o *UpdateSyncJobRequestAttr) GetExclude() string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *UpdateSyncJobRequestAttr) GetExcludeOk() (*string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *UpdateSyncJobRequestAttr) SetExclude(v string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *UpdateSyncJobRequestAttr) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetHost1

`func (o *UpdateSyncJobRequestAttr) GetHost1() string`

GetHost1 returns the Host1 field if non-nil, zero value otherwise.

### GetHost1Ok

`func (o *UpdateSyncJobRequestAttr) GetHost1Ok() (*string, bool)`

GetHost1Ok returns a tuple with the Host1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost1

`func (o *UpdateSyncJobRequestAttr) SetHost1(v string)`

SetHost1 sets Host1 field to given value.

### HasHost1

`func (o *UpdateSyncJobRequestAttr) HasHost1() bool`

HasHost1 returns a boolean if a field has been set.

### GetMaxage

`func (o *UpdateSyncJobRequestAttr) GetMaxage() float32`

GetMaxage returns the Maxage field if non-nil, zero value otherwise.

### GetMaxageOk

`func (o *UpdateSyncJobRequestAttr) GetMaxageOk() (*float32, bool)`

GetMaxageOk returns a tuple with the Maxage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxage

`func (o *UpdateSyncJobRequestAttr) SetMaxage(v float32)`

SetMaxage sets Maxage field to given value.

### HasMaxage

`func (o *UpdateSyncJobRequestAttr) HasMaxage() bool`

HasMaxage returns a boolean if a field has been set.

### GetMaxbytespersecond

`func (o *UpdateSyncJobRequestAttr) GetMaxbytespersecond() float32`

GetMaxbytespersecond returns the Maxbytespersecond field if non-nil, zero value otherwise.

### GetMaxbytespersecondOk

`func (o *UpdateSyncJobRequestAttr) GetMaxbytespersecondOk() (*float32, bool)`

GetMaxbytespersecondOk returns a tuple with the Maxbytespersecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbytespersecond

`func (o *UpdateSyncJobRequestAttr) SetMaxbytespersecond(v float32)`

SetMaxbytespersecond sets Maxbytespersecond field to given value.

### HasMaxbytespersecond

`func (o *UpdateSyncJobRequestAttr) HasMaxbytespersecond() bool`

HasMaxbytespersecond returns a boolean if a field has been set.

### GetMinsInterval

`func (o *UpdateSyncJobRequestAttr) GetMinsInterval() float32`

GetMinsInterval returns the MinsInterval field if non-nil, zero value otherwise.

### GetMinsIntervalOk

`func (o *UpdateSyncJobRequestAttr) GetMinsIntervalOk() (*float32, bool)`

GetMinsIntervalOk returns a tuple with the MinsInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinsInterval

`func (o *UpdateSyncJobRequestAttr) SetMinsInterval(v float32)`

SetMinsInterval sets MinsInterval field to given value.

### HasMinsInterval

`func (o *UpdateSyncJobRequestAttr) HasMinsInterval() bool`

HasMinsInterval returns a boolean if a field has been set.

### GetPassword1

`func (o *UpdateSyncJobRequestAttr) GetPassword1() string`

GetPassword1 returns the Password1 field if non-nil, zero value otherwise.

### GetPassword1Ok

`func (o *UpdateSyncJobRequestAttr) GetPassword1Ok() (*string, bool)`

GetPassword1Ok returns a tuple with the Password1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword1

`func (o *UpdateSyncJobRequestAttr) SetPassword1(v string)`

SetPassword1 sets Password1 field to given value.

### HasPassword1

`func (o *UpdateSyncJobRequestAttr) HasPassword1() bool`

HasPassword1 returns a boolean if a field has been set.

### GetPort1

`func (o *UpdateSyncJobRequestAttr) GetPort1() string`

GetPort1 returns the Port1 field if non-nil, zero value otherwise.

### GetPort1Ok

`func (o *UpdateSyncJobRequestAttr) GetPort1Ok() (*string, bool)`

GetPort1Ok returns a tuple with the Port1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort1

`func (o *UpdateSyncJobRequestAttr) SetPort1(v string)`

SetPort1 sets Port1 field to given value.

### HasPort1

`func (o *UpdateSyncJobRequestAttr) HasPort1() bool`

HasPort1 returns a boolean if a field has been set.

### GetSkipcrossduplicates

`func (o *UpdateSyncJobRequestAttr) GetSkipcrossduplicates() bool`

GetSkipcrossduplicates returns the Skipcrossduplicates field if non-nil, zero value otherwise.

### GetSkipcrossduplicatesOk

`func (o *UpdateSyncJobRequestAttr) GetSkipcrossduplicatesOk() (*bool, bool)`

GetSkipcrossduplicatesOk returns a tuple with the Skipcrossduplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipcrossduplicates

`func (o *UpdateSyncJobRequestAttr) SetSkipcrossduplicates(v bool)`

SetSkipcrossduplicates sets Skipcrossduplicates field to given value.

### HasSkipcrossduplicates

`func (o *UpdateSyncJobRequestAttr) HasSkipcrossduplicates() bool`

HasSkipcrossduplicates returns a boolean if a field has been set.

### GetSubfolder2

`func (o *UpdateSyncJobRequestAttr) GetSubfolder2() string`

GetSubfolder2 returns the Subfolder2 field if non-nil, zero value otherwise.

### GetSubfolder2Ok

`func (o *UpdateSyncJobRequestAttr) GetSubfolder2Ok() (*string, bool)`

GetSubfolder2Ok returns a tuple with the Subfolder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolder2

`func (o *UpdateSyncJobRequestAttr) SetSubfolder2(v string)`

SetSubfolder2 sets Subfolder2 field to given value.

### HasSubfolder2

`func (o *UpdateSyncJobRequestAttr) HasSubfolder2() bool`

HasSubfolder2 returns a boolean if a field has been set.

### GetSubscribeall

`func (o *UpdateSyncJobRequestAttr) GetSubscribeall() bool`

GetSubscribeall returns the Subscribeall field if non-nil, zero value otherwise.

### GetSubscribeallOk

`func (o *UpdateSyncJobRequestAttr) GetSubscribeallOk() (*bool, bool)`

GetSubscribeallOk returns a tuple with the Subscribeall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeall

`func (o *UpdateSyncJobRequestAttr) SetSubscribeall(v bool)`

SetSubscribeall sets Subscribeall field to given value.

### HasSubscribeall

`func (o *UpdateSyncJobRequestAttr) HasSubscribeall() bool`

HasSubscribeall returns a boolean if a field has been set.

### GetTimeout1

`func (o *UpdateSyncJobRequestAttr) GetTimeout1() float32`

GetTimeout1 returns the Timeout1 field if non-nil, zero value otherwise.

### GetTimeout1Ok

`func (o *UpdateSyncJobRequestAttr) GetTimeout1Ok() (*float32, bool)`

GetTimeout1Ok returns a tuple with the Timeout1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout1

`func (o *UpdateSyncJobRequestAttr) SetTimeout1(v float32)`

SetTimeout1 sets Timeout1 field to given value.

### HasTimeout1

`func (o *UpdateSyncJobRequestAttr) HasTimeout1() bool`

HasTimeout1 returns a boolean if a field has been set.

### GetTimeout2

`func (o *UpdateSyncJobRequestAttr) GetTimeout2() float32`

GetTimeout2 returns the Timeout2 field if non-nil, zero value otherwise.

### GetTimeout2Ok

`func (o *UpdateSyncJobRequestAttr) GetTimeout2Ok() (*float32, bool)`

GetTimeout2Ok returns a tuple with the Timeout2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout2

`func (o *UpdateSyncJobRequestAttr) SetTimeout2(v float32)`

SetTimeout2 sets Timeout2 field to given value.

### HasTimeout2

`func (o *UpdateSyncJobRequestAttr) HasTimeout2() bool`

HasTimeout2 returns a boolean if a field has been set.

### GetUser1

`func (o *UpdateSyncJobRequestAttr) GetUser1() string`

GetUser1 returns the User1 field if non-nil, zero value otherwise.

### GetUser1Ok

`func (o *UpdateSyncJobRequestAttr) GetUser1Ok() (*string, bool)`

GetUser1Ok returns a tuple with the User1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser1

`func (o *UpdateSyncJobRequestAttr) SetUser1(v string)`

SetUser1 sets User1 field to given value.

### HasUser1

`func (o *UpdateSyncJobRequestAttr) HasUser1() bool`

HasUser1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


