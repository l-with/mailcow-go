# EditFail2BanRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backlist** | Pointer to **string** | the backlisted ips or hostnames separated by comma | [optional] 
**BanTime** | Pointer to **float32** | the time a ip should be banned | [optional] 
**MaxAttempts** | Pointer to **float32** | the maximum numbe of wrong logins before a ip is banned | [optional] 
**NetbanIpv4** | Pointer to **float32** | the networks mask to ban for ipv4 | [optional] 
**NetbanIpv6** | Pointer to **float32** | the networks mask to ban for ipv6 | [optional] 
**RetryWindow** | Pointer to **float32** | the maximum time in which a ip as to login with false credentials to be banned | [optional] 
**Whitelist** | Pointer to **string** | whitelisted ips or hostnames sepereated by comma | [optional] 

## Methods

### NewEditFail2BanRequestAttr

`func NewEditFail2BanRequestAttr() *EditFail2BanRequestAttr`

NewEditFail2BanRequestAttr instantiates a new EditFail2BanRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditFail2BanRequestAttrWithDefaults

`func NewEditFail2BanRequestAttrWithDefaults() *EditFail2BanRequestAttr`

NewEditFail2BanRequestAttrWithDefaults instantiates a new EditFail2BanRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacklist

`func (o *EditFail2BanRequestAttr) GetBacklist() string`

GetBacklist returns the Backlist field if non-nil, zero value otherwise.

### GetBacklistOk

`func (o *EditFail2BanRequestAttr) GetBacklistOk() (*string, bool)`

GetBacklistOk returns a tuple with the Backlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklist

`func (o *EditFail2BanRequestAttr) SetBacklist(v string)`

SetBacklist sets Backlist field to given value.

### HasBacklist

`func (o *EditFail2BanRequestAttr) HasBacklist() bool`

HasBacklist returns a boolean if a field has been set.

### GetBanTime

`func (o *EditFail2BanRequestAttr) GetBanTime() float32`

GetBanTime returns the BanTime field if non-nil, zero value otherwise.

### GetBanTimeOk

`func (o *EditFail2BanRequestAttr) GetBanTimeOk() (*float32, bool)`

GetBanTimeOk returns a tuple with the BanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanTime

`func (o *EditFail2BanRequestAttr) SetBanTime(v float32)`

SetBanTime sets BanTime field to given value.

### HasBanTime

`func (o *EditFail2BanRequestAttr) HasBanTime() bool`

HasBanTime returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *EditFail2BanRequestAttr) GetMaxAttempts() float32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *EditFail2BanRequestAttr) GetMaxAttemptsOk() (*float32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *EditFail2BanRequestAttr) SetMaxAttempts(v float32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *EditFail2BanRequestAttr) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetNetbanIpv4

`func (o *EditFail2BanRequestAttr) GetNetbanIpv4() float32`

GetNetbanIpv4 returns the NetbanIpv4 field if non-nil, zero value otherwise.

### GetNetbanIpv4Ok

`func (o *EditFail2BanRequestAttr) GetNetbanIpv4Ok() (*float32, bool)`

GetNetbanIpv4Ok returns a tuple with the NetbanIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanIpv4

`func (o *EditFail2BanRequestAttr) SetNetbanIpv4(v float32)`

SetNetbanIpv4 sets NetbanIpv4 field to given value.

### HasNetbanIpv4

`func (o *EditFail2BanRequestAttr) HasNetbanIpv4() bool`

HasNetbanIpv4 returns a boolean if a field has been set.

### GetNetbanIpv6

`func (o *EditFail2BanRequestAttr) GetNetbanIpv6() float32`

GetNetbanIpv6 returns the NetbanIpv6 field if non-nil, zero value otherwise.

### GetNetbanIpv6Ok

`func (o *EditFail2BanRequestAttr) GetNetbanIpv6Ok() (*float32, bool)`

GetNetbanIpv6Ok returns a tuple with the NetbanIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanIpv6

`func (o *EditFail2BanRequestAttr) SetNetbanIpv6(v float32)`

SetNetbanIpv6 sets NetbanIpv6 field to given value.

### HasNetbanIpv6

`func (o *EditFail2BanRequestAttr) HasNetbanIpv6() bool`

HasNetbanIpv6 returns a boolean if a field has been set.

### GetRetryWindow

`func (o *EditFail2BanRequestAttr) GetRetryWindow() float32`

GetRetryWindow returns the RetryWindow field if non-nil, zero value otherwise.

### GetRetryWindowOk

`func (o *EditFail2BanRequestAttr) GetRetryWindowOk() (*float32, bool)`

GetRetryWindowOk returns a tuple with the RetryWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWindow

`func (o *EditFail2BanRequestAttr) SetRetryWindow(v float32)`

SetRetryWindow sets RetryWindow field to given value.

### HasRetryWindow

`func (o *EditFail2BanRequestAttr) HasRetryWindow() bool`

HasRetryWindow returns a boolean if a field has been set.

### GetWhitelist

`func (o *EditFail2BanRequestAttr) GetWhitelist() string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *EditFail2BanRequestAttr) GetWhitelistOk() (*string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *EditFail2BanRequestAttr) SetWhitelist(v string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *EditFail2BanRequestAttr) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


