# UpdateAliasRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is alias active or not | [optional] 
**Address** | Pointer to **string** | alias address, for catchall use \&quot;@domain.tld\&quot; | [optional] 
**Goto** | Pointer to **string** | destination address, comma separated | [optional] 
**GotoHam** | Pointer to **bool** | learn as ham | [optional] 
**GotoNull** | Pointer to **bool** | silently ignore | [optional] 
**GotoSpam** | Pointer to **bool** | learn as spam | [optional] 
**PrivateComment** | Pointer to **string** |  | [optional] 
**PublicComment** | Pointer to **string** |  | [optional] 
**SogoVisible** | Pointer to **bool** | toggle visibility as selectable sender in SOGo | [optional] 

## Methods

### NewUpdateAliasRequestAttr

`func NewUpdateAliasRequestAttr() *UpdateAliasRequestAttr`

NewUpdateAliasRequestAttr instantiates a new UpdateAliasRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAliasRequestAttrWithDefaults

`func NewUpdateAliasRequestAttrWithDefaults() *UpdateAliasRequestAttr`

NewUpdateAliasRequestAttrWithDefaults instantiates a new UpdateAliasRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateAliasRequestAttr) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateAliasRequestAttr) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateAliasRequestAttr) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateAliasRequestAttr) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateAliasRequestAttr) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateAliasRequestAttr) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateAliasRequestAttr) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateAliasRequestAttr) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGoto

`func (o *UpdateAliasRequestAttr) GetGoto() string`

GetGoto returns the Goto field if non-nil, zero value otherwise.

### GetGotoOk

`func (o *UpdateAliasRequestAttr) GetGotoOk() (*string, bool)`

GetGotoOk returns a tuple with the Goto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoto

`func (o *UpdateAliasRequestAttr) SetGoto(v string)`

SetGoto sets Goto field to given value.

### HasGoto

`func (o *UpdateAliasRequestAttr) HasGoto() bool`

HasGoto returns a boolean if a field has been set.

### GetGotoHam

`func (o *UpdateAliasRequestAttr) GetGotoHam() bool`

GetGotoHam returns the GotoHam field if non-nil, zero value otherwise.

### GetGotoHamOk

`func (o *UpdateAliasRequestAttr) GetGotoHamOk() (*bool, bool)`

GetGotoHamOk returns a tuple with the GotoHam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotoHam

`func (o *UpdateAliasRequestAttr) SetGotoHam(v bool)`

SetGotoHam sets GotoHam field to given value.

### HasGotoHam

`func (o *UpdateAliasRequestAttr) HasGotoHam() bool`

HasGotoHam returns a boolean if a field has been set.

### GetGotoNull

`func (o *UpdateAliasRequestAttr) GetGotoNull() bool`

GetGotoNull returns the GotoNull field if non-nil, zero value otherwise.

### GetGotoNullOk

`func (o *UpdateAliasRequestAttr) GetGotoNullOk() (*bool, bool)`

GetGotoNullOk returns a tuple with the GotoNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotoNull

`func (o *UpdateAliasRequestAttr) SetGotoNull(v bool)`

SetGotoNull sets GotoNull field to given value.

### HasGotoNull

`func (o *UpdateAliasRequestAttr) HasGotoNull() bool`

HasGotoNull returns a boolean if a field has been set.

### GetGotoSpam

`func (o *UpdateAliasRequestAttr) GetGotoSpam() bool`

GetGotoSpam returns the GotoSpam field if non-nil, zero value otherwise.

### GetGotoSpamOk

`func (o *UpdateAliasRequestAttr) GetGotoSpamOk() (*bool, bool)`

GetGotoSpamOk returns a tuple with the GotoSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotoSpam

`func (o *UpdateAliasRequestAttr) SetGotoSpam(v bool)`

SetGotoSpam sets GotoSpam field to given value.

### HasGotoSpam

`func (o *UpdateAliasRequestAttr) HasGotoSpam() bool`

HasGotoSpam returns a boolean if a field has been set.

### GetPrivateComment

`func (o *UpdateAliasRequestAttr) GetPrivateComment() string`

GetPrivateComment returns the PrivateComment field if non-nil, zero value otherwise.

### GetPrivateCommentOk

`func (o *UpdateAliasRequestAttr) GetPrivateCommentOk() (*string, bool)`

GetPrivateCommentOk returns a tuple with the PrivateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateComment

`func (o *UpdateAliasRequestAttr) SetPrivateComment(v string)`

SetPrivateComment sets PrivateComment field to given value.

### HasPrivateComment

`func (o *UpdateAliasRequestAttr) HasPrivateComment() bool`

HasPrivateComment returns a boolean if a field has been set.

### GetPublicComment

`func (o *UpdateAliasRequestAttr) GetPublicComment() string`

GetPublicComment returns the PublicComment field if non-nil, zero value otherwise.

### GetPublicCommentOk

`func (o *UpdateAliasRequestAttr) GetPublicCommentOk() (*string, bool)`

GetPublicCommentOk returns a tuple with the PublicComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicComment

`func (o *UpdateAliasRequestAttr) SetPublicComment(v string)`

SetPublicComment sets PublicComment field to given value.

### HasPublicComment

`func (o *UpdateAliasRequestAttr) HasPublicComment() bool`

HasPublicComment returns a boolean if a field has been set.

### GetSogoVisible

`func (o *UpdateAliasRequestAttr) GetSogoVisible() bool`

GetSogoVisible returns the SogoVisible field if non-nil, zero value otherwise.

### GetSogoVisibleOk

`func (o *UpdateAliasRequestAttr) GetSogoVisibleOk() (*bool, bool)`

GetSogoVisibleOk returns a tuple with the SogoVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSogoVisible

`func (o *UpdateAliasRequestAttr) SetSogoVisible(v bool)`

SetSogoVisible sets SogoVisible field to given value.

### HasSogoVisible

`func (o *UpdateAliasRequestAttr) HasSogoVisible() bool`

HasSogoVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


