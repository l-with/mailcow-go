# AddForwardHostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterSpam** | Pointer to **float32** | 1 to enable spam filter, 0 to disable spam filter | [optional] 
**Hostname** | Pointer to **string** | contains the hostname you want to add | [optional] 

## Methods

### NewAddForwardHostRequest

`func NewAddForwardHostRequest() *AddForwardHostRequest`

NewAddForwardHostRequest instantiates a new AddForwardHostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddForwardHostRequestWithDefaults

`func NewAddForwardHostRequestWithDefaults() *AddForwardHostRequest`

NewAddForwardHostRequestWithDefaults instantiates a new AddForwardHostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterSpam

`func (o *AddForwardHostRequest) GetFilterSpam() float32`

GetFilterSpam returns the FilterSpam field if non-nil, zero value otherwise.

### GetFilterSpamOk

`func (o *AddForwardHostRequest) GetFilterSpamOk() (*float32, bool)`

GetFilterSpamOk returns a tuple with the FilterSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSpam

`func (o *AddForwardHostRequest) SetFilterSpam(v float32)`

SetFilterSpam sets FilterSpam field to given value.

### HasFilterSpam

`func (o *AddForwardHostRequest) HasFilterSpam() bool`

HasFilterSpam returns a boolean if a field has been set.

### GetHostname

`func (o *AddForwardHostRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddForwardHostRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddForwardHostRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AddForwardHostRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


