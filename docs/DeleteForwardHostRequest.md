# DeleteForwardHostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | contains the ip of the fowarding host you want to delete | [optional] 

## Methods

### NewDeleteForwardHostRequest

`func NewDeleteForwardHostRequest() *DeleteForwardHostRequest`

NewDeleteForwardHostRequest instantiates a new DeleteForwardHostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteForwardHostRequestWithDefaults

`func NewDeleteForwardHostRequestWithDefaults() *DeleteForwardHostRequest`

NewDeleteForwardHostRequestWithDefaults instantiates a new DeleteForwardHostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DeleteForwardHostRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeleteForwardHostRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeleteForwardHostRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeleteForwardHostRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


