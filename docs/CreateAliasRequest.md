# CreateAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is alias active or not | [optional] 
**Address** | Pointer to **string** | alias address, for catchall use \&quot;@domain.tld\&quot; | [optional] 
**Goto** | Pointer to **string** | destination address, comma separated | [optional] 
**GotoHam** | Pointer to **bool** | learn as ham | [optional] 
**GotoNull** | Pointer to **bool** | silently ignore | [optional] 
**GotoSpam** | Pointer to **bool** | learn as spam | [optional] 
**SogoVisible** | Pointer to **bool** | toggle visibility as selectable sender in SOGo | [optional] 

## Methods

### NewCreateAliasRequest

`func NewCreateAliasRequest() *CreateAliasRequest`

NewCreateAliasRequest instantiates a new CreateAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAliasRequestWithDefaults

`func NewCreateAliasRequestWithDefaults() *CreateAliasRequest`

NewCreateAliasRequestWithDefaults instantiates a new CreateAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateAliasRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateAliasRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateAliasRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateAliasRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress

`func (o *CreateAliasRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateAliasRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateAliasRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateAliasRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGoto

`func (o *CreateAliasRequest) GetGoto() string`

GetGoto returns the Goto field if non-nil, zero value otherwise.

### GetGotoOk

`func (o *CreateAliasRequest) GetGotoOk() (*string, bool)`

GetGotoOk returns a tuple with the Goto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoto

`func (o *CreateAliasRequest) SetGoto(v string)`

SetGoto sets Goto field to given value.

### HasGoto

`func (o *CreateAliasRequest) HasGoto() bool`

HasGoto returns a boolean if a field has been set.

### GetGotoHam

`func (o *CreateAliasRequest) GetGotoHam() bool`

GetGotoHam returns the GotoHam field if non-nil, zero value otherwise.

### GetGotoHamOk

`func (o *CreateAliasRequest) GetGotoHamOk() (*bool, bool)`

GetGotoHamOk returns a tuple with the GotoHam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotoHam

`func (o *CreateAliasRequest) SetGotoHam(v bool)`

SetGotoHam sets GotoHam field to given value.

### HasGotoHam

`func (o *CreateAliasRequest) HasGotoHam() bool`

HasGotoHam returns a boolean if a field has been set.

### GetGotoNull

`func (o *CreateAliasRequest) GetGotoNull() bool`

GetGotoNull returns the GotoNull field if non-nil, zero value otherwise.

### GetGotoNullOk

`func (o *CreateAliasRequest) GetGotoNullOk() (*bool, bool)`

GetGotoNullOk returns a tuple with the GotoNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotoNull

`func (o *CreateAliasRequest) SetGotoNull(v bool)`

SetGotoNull sets GotoNull field to given value.

### HasGotoNull

`func (o *CreateAliasRequest) HasGotoNull() bool`

HasGotoNull returns a boolean if a field has been set.

### GetGotoSpam

`func (o *CreateAliasRequest) GetGotoSpam() bool`

GetGotoSpam returns the GotoSpam field if non-nil, zero value otherwise.

### GetGotoSpamOk

`func (o *CreateAliasRequest) GetGotoSpamOk() (*bool, bool)`

GetGotoSpamOk returns a tuple with the GotoSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotoSpam

`func (o *CreateAliasRequest) SetGotoSpam(v bool)`

SetGotoSpam sets GotoSpam field to given value.

### HasGotoSpam

`func (o *CreateAliasRequest) HasGotoSpam() bool`

HasGotoSpam returns a boolean if a field has been set.

### GetSogoVisible

`func (o *CreateAliasRequest) GetSogoVisible() bool`

GetSogoVisible returns the SogoVisible field if non-nil, zero value otherwise.

### GetSogoVisibleOk

`func (o *CreateAliasRequest) GetSogoVisibleOk() (*bool, bool)`

GetSogoVisibleOk returns a tuple with the SogoVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSogoVisible

`func (o *CreateAliasRequest) SetSogoVisible(v bool)`

SetSogoVisible sets SogoVisible field to given value.

### HasSogoVisible

`func (o *CreateAliasRequest) HasSogoVisible() bool`

HasSogoVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


