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

// DeleteSyncJobRequest struct for DeleteSyncJobRequest
type DeleteSyncJobRequest struct {
	// contains list of aliases you want to delete
	Items map[string]interface{} `json:"items,omitempty"`
}

// NewDeleteSyncJobRequest instantiates a new DeleteSyncJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSyncJobRequest() *DeleteSyncJobRequest {
	this := DeleteSyncJobRequest{}
	return &this
}

// NewDeleteSyncJobRequestWithDefaults instantiates a new DeleteSyncJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSyncJobRequestWithDefaults() *DeleteSyncJobRequest {
	this := DeleteSyncJobRequest{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DeleteSyncJobRequest) GetItems() map[string]interface{} {
	if o == nil || o.Items == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSyncJobRequest) GetItemsOk() (map[string]interface{}, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeleteSyncJobRequest) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string]interface{} and assigns it to the Items field.
func (o *DeleteSyncJobRequest) SetItems(v map[string]interface{}) {
	o.Items = v
}

func (o DeleteSyncJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSyncJobRequest struct {
	value *DeleteSyncJobRequest
	isSet bool
}

func (v NullableDeleteSyncJobRequest) Get() *DeleteSyncJobRequest {
	return v.value
}

func (v *NullableDeleteSyncJobRequest) Set(val *DeleteSyncJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSyncJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSyncJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSyncJobRequest(val *DeleteSyncJobRequest) *NullableDeleteSyncJobRequest {
	return &NullableDeleteSyncJobRequest{value: val, isSet: true}
}

func (v NullableDeleteSyncJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSyncJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
