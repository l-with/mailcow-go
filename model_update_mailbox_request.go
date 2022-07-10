/*
mailcow API

mailcow is complete e-mailing solution with advanced antispam, antivirus, nice UI and API.  In order to use this API you have to create a API key and add your IP address to the whitelist of allowed IPs this can be done by logging into the Mailcow UI using your admin account, then go to Configuration > Access > Edit administrator details > API. There you will find a collapsed API menu.  There are two types of API keys   - The read only key can only be used for all get endpoints   - The read write key can be used for all endpoints

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateMailboxRequest struct for UpdateMailboxRequest
type UpdateMailboxRequest struct {
	Attr *UpdateMailboxRequestAttr `json:"attr,omitempty"`
	// contains list of mailboxes you want update
	Items map[string]interface{} `json:"items,omitempty"`
}

// NewUpdateMailboxRequest instantiates a new UpdateMailboxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMailboxRequest() *UpdateMailboxRequest {
	this := UpdateMailboxRequest{}
	return &this
}

// NewUpdateMailboxRequestWithDefaults instantiates a new UpdateMailboxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMailboxRequestWithDefaults() *UpdateMailboxRequest {
	this := UpdateMailboxRequest{}
	return &this
}

// GetAttr returns the Attr field value if set, zero value otherwise.
func (o *UpdateMailboxRequest) GetAttr() UpdateMailboxRequestAttr {
	if o == nil || o.Attr == nil {
		var ret UpdateMailboxRequestAttr
		return ret
	}
	return *o.Attr
}

// GetAttrOk returns a tuple with the Attr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequest) GetAttrOk() (*UpdateMailboxRequestAttr, bool) {
	if o == nil || o.Attr == nil {
		return nil, false
	}
	return o.Attr, true
}

// HasAttr returns a boolean if a field has been set.
func (o *UpdateMailboxRequest) HasAttr() bool {
	if o != nil && o.Attr != nil {
		return true
	}

	return false
}

// SetAttr gets a reference to the given UpdateMailboxRequestAttr and assigns it to the Attr field.
func (o *UpdateMailboxRequest) SetAttr(v UpdateMailboxRequestAttr) {
	o.Attr = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UpdateMailboxRequest) GetItems() map[string]interface{} {
	if o == nil || o.Items == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequest) GetItemsOk() (map[string]interface{}, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UpdateMailboxRequest) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string]interface{} and assigns it to the Items field.
func (o *UpdateMailboxRequest) SetItems(v map[string]interface{}) {
	o.Items = v
}

func (o UpdateMailboxRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attr != nil {
		toSerialize["attr"] = o.Attr
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateMailboxRequest struct {
	value *UpdateMailboxRequest
	isSet bool
}

func (v NullableUpdateMailboxRequest) Get() *UpdateMailboxRequest {
	return v.value
}

func (v *NullableUpdateMailboxRequest) Set(val *UpdateMailboxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMailboxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMailboxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMailboxRequest(val *UpdateMailboxRequest) *NullableUpdateMailboxRequest {
	return &NullableUpdateMailboxRequest{value: val, isSet: true}
}

func (v NullableUpdateMailboxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMailboxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
