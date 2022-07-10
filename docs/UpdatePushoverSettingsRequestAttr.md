# UpdatePushoverSettingsRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **float32** | Enables pushover 1 disable pushover 0 | [optional] 
**EvaluateXPrio** | Pointer to **float32** | Send the Push with High priority | [optional] 
**Key** | Pointer to **string** | Pushover key | [optional] 
**OnlyXPrio** | Pointer to **float32** | Only send push for prio mails | [optional] 
**Senders** | Pointer to **string** | Only send push for emails from these senders | [optional] 
**SendersRegex** | Pointer to **string** | Regex to match senders for which a push will be send | [optional] 
**Text** | Pointer to **string** | Custom push noficiation text | [optional] 
**Title** | Pointer to **string** | Push title | [optional] 
**Token** | Pointer to **string** | Pushover token | [optional] 

## Methods

### NewUpdatePushoverSettingsRequestAttr

`func NewUpdatePushoverSettingsRequestAttr() *UpdatePushoverSettingsRequestAttr`

NewUpdatePushoverSettingsRequestAttr instantiates a new UpdatePushoverSettingsRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePushoverSettingsRequestAttrWithDefaults

`func NewUpdatePushoverSettingsRequestAttrWithDefaults() *UpdatePushoverSettingsRequestAttr`

NewUpdatePushoverSettingsRequestAttrWithDefaults instantiates a new UpdatePushoverSettingsRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdatePushoverSettingsRequestAttr) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdatePushoverSettingsRequestAttr) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdatePushoverSettingsRequestAttr) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdatePushoverSettingsRequestAttr) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEvaluateXPrio

`func (o *UpdatePushoverSettingsRequestAttr) GetEvaluateXPrio() float32`

GetEvaluateXPrio returns the EvaluateXPrio field if non-nil, zero value otherwise.

### GetEvaluateXPrioOk

`func (o *UpdatePushoverSettingsRequestAttr) GetEvaluateXPrioOk() (*float32, bool)`

GetEvaluateXPrioOk returns a tuple with the EvaluateXPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateXPrio

`func (o *UpdatePushoverSettingsRequestAttr) SetEvaluateXPrio(v float32)`

SetEvaluateXPrio sets EvaluateXPrio field to given value.

### HasEvaluateXPrio

`func (o *UpdatePushoverSettingsRequestAttr) HasEvaluateXPrio() bool`

HasEvaluateXPrio returns a boolean if a field has been set.

### GetKey

`func (o *UpdatePushoverSettingsRequestAttr) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdatePushoverSettingsRequestAttr) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdatePushoverSettingsRequestAttr) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdatePushoverSettingsRequestAttr) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOnlyXPrio

`func (o *UpdatePushoverSettingsRequestAttr) GetOnlyXPrio() float32`

GetOnlyXPrio returns the OnlyXPrio field if non-nil, zero value otherwise.

### GetOnlyXPrioOk

`func (o *UpdatePushoverSettingsRequestAttr) GetOnlyXPrioOk() (*float32, bool)`

GetOnlyXPrioOk returns a tuple with the OnlyXPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyXPrio

`func (o *UpdatePushoverSettingsRequestAttr) SetOnlyXPrio(v float32)`

SetOnlyXPrio sets OnlyXPrio field to given value.

### HasOnlyXPrio

`func (o *UpdatePushoverSettingsRequestAttr) HasOnlyXPrio() bool`

HasOnlyXPrio returns a boolean if a field has been set.

### GetSenders

`func (o *UpdatePushoverSettingsRequestAttr) GetSenders() string`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *UpdatePushoverSettingsRequestAttr) GetSendersOk() (*string, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *UpdatePushoverSettingsRequestAttr) SetSenders(v string)`

SetSenders sets Senders field to given value.

### HasSenders

`func (o *UpdatePushoverSettingsRequestAttr) HasSenders() bool`

HasSenders returns a boolean if a field has been set.

### GetSendersRegex

`func (o *UpdatePushoverSettingsRequestAttr) GetSendersRegex() string`

GetSendersRegex returns the SendersRegex field if non-nil, zero value otherwise.

### GetSendersRegexOk

`func (o *UpdatePushoverSettingsRequestAttr) GetSendersRegexOk() (*string, bool)`

GetSendersRegexOk returns a tuple with the SendersRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendersRegex

`func (o *UpdatePushoverSettingsRequestAttr) SetSendersRegex(v string)`

SetSendersRegex sets SendersRegex field to given value.

### HasSendersRegex

`func (o *UpdatePushoverSettingsRequestAttr) HasSendersRegex() bool`

HasSendersRegex returns a boolean if a field has been set.

### GetText

`func (o *UpdatePushoverSettingsRequestAttr) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *UpdatePushoverSettingsRequestAttr) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *UpdatePushoverSettingsRequestAttr) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *UpdatePushoverSettingsRequestAttr) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *UpdatePushoverSettingsRequestAttr) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePushoverSettingsRequestAttr) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePushoverSettingsRequestAttr) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdatePushoverSettingsRequestAttr) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetToken

`func (o *UpdatePushoverSettingsRequestAttr) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdatePushoverSettingsRequestAttr) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdatePushoverSettingsRequestAttr) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdatePushoverSettingsRequestAttr) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


