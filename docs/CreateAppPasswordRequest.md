# CreateAppPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is alias active or not | [optional] 
**Username** | Pointer to **string** | mailbox for which the app password should be created | [optional] 
**AppName** | Pointer to **string** | name of your app password | [optional] 
**AppPasswd** | Pointer to **string** | your app password | [optional] 
**AppPasswd2** | Pointer to **string** | your app password | [optional] 

## Methods

### NewCreateAppPasswordRequest

`func NewCreateAppPasswordRequest() *CreateAppPasswordRequest`

NewCreateAppPasswordRequest instantiates a new CreateAppPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppPasswordRequestWithDefaults

`func NewCreateAppPasswordRequestWithDefaults() *CreateAppPasswordRequest`

NewCreateAppPasswordRequestWithDefaults instantiates a new CreateAppPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateAppPasswordRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateAppPasswordRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateAppPasswordRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateAppPasswordRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUsername

`func (o *CreateAppPasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAppPasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAppPasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateAppPasswordRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAppName

`func (o *CreateAppPasswordRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *CreateAppPasswordRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *CreateAppPasswordRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *CreateAppPasswordRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppPasswd

`func (o *CreateAppPasswordRequest) GetAppPasswd() string`

GetAppPasswd returns the AppPasswd field if non-nil, zero value otherwise.

### GetAppPasswdOk

`func (o *CreateAppPasswordRequest) GetAppPasswdOk() (*string, bool)`

GetAppPasswdOk returns a tuple with the AppPasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPasswd

`func (o *CreateAppPasswordRequest) SetAppPasswd(v string)`

SetAppPasswd sets AppPasswd field to given value.

### HasAppPasswd

`func (o *CreateAppPasswordRequest) HasAppPasswd() bool`

HasAppPasswd returns a boolean if a field has been set.

### GetAppPasswd2

`func (o *CreateAppPasswordRequest) GetAppPasswd2() string`

GetAppPasswd2 returns the AppPasswd2 field if non-nil, zero value otherwise.

### GetAppPasswd2Ok

`func (o *CreateAppPasswordRequest) GetAppPasswd2Ok() (*string, bool)`

GetAppPasswd2Ok returns a tuple with the AppPasswd2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPasswd2

`func (o *CreateAppPasswordRequest) SetAppPasswd2(v string)`

SetAppPasswd2 sets AppPasswd2 field to given value.

### HasAppPasswd2

`func (o *CreateAppPasswordRequest) HasAppPasswd2() bool`

HasAppPasswd2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


