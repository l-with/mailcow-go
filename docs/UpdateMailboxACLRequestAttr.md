# UpdateMailboxACLRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAcl** | Pointer to **map[string]interface{}** | contains a list of active user acls | [optional] 

## Methods

### NewUpdateMailboxACLRequestAttr

`func NewUpdateMailboxACLRequestAttr() *UpdateMailboxACLRequestAttr`

NewUpdateMailboxACLRequestAttr instantiates a new UpdateMailboxACLRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxACLRequestAttrWithDefaults

`func NewUpdateMailboxACLRequestAttrWithDefaults() *UpdateMailboxACLRequestAttr`

NewUpdateMailboxACLRequestAttrWithDefaults instantiates a new UpdateMailboxACLRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAcl

`func (o *UpdateMailboxACLRequestAttr) GetUserAcl() map[string]interface{}`

GetUserAcl returns the UserAcl field if non-nil, zero value otherwise.

### GetUserAclOk

`func (o *UpdateMailboxACLRequestAttr) GetUserAclOk() (*map[string]interface{}, bool)`

GetUserAclOk returns a tuple with the UserAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAcl

`func (o *UpdateMailboxACLRequestAttr) SetUserAcl(v map[string]interface{})`

SetUserAcl sets UserAcl field to given value.

### HasUserAcl

`func (o *UpdateMailboxACLRequestAttr) HasUserAcl() bool`

HasUserAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


