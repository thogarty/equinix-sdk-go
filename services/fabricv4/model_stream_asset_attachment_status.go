/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// StreamAssetAttachmentStatus asset status
type StreamAssetAttachmentStatus string

// List of StreamAsset_attachmentStatus
const (
	STREAMASSETATTACHMENTSTATUS_ATTACHING StreamAssetAttachmentStatus = "ATTACHING"
	STREAMASSETATTACHMENTSTATUS_ATTACHED  StreamAssetAttachmentStatus = "ATTACHED"
	STREAMASSETATTACHMENTSTATUS_DETACHED  StreamAssetAttachmentStatus = "DETACHED"
	STREAMASSETATTACHMENTSTATUS_DETACHING StreamAssetAttachmentStatus = "DETACHING"
	STREAMASSETATTACHMENTSTATUS_FAILED    StreamAssetAttachmentStatus = "FAILED"
)

// All allowed values of StreamAssetAttachmentStatus enum
var AllowedStreamAssetAttachmentStatusEnumValues = []StreamAssetAttachmentStatus{
	"ATTACHING",
	"ATTACHED",
	"DETACHED",
	"DETACHING",
	"FAILED",
}

func (v *StreamAssetAttachmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StreamAssetAttachmentStatus(value)
	for _, existing := range AllowedStreamAssetAttachmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StreamAssetAttachmentStatus", value)
}

// NewStreamAssetAttachmentStatusFromValue returns a pointer to a valid StreamAssetAttachmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStreamAssetAttachmentStatusFromValue(v string) (*StreamAssetAttachmentStatus, error) {
	ev := StreamAssetAttachmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StreamAssetAttachmentStatus: valid values are %v", v, AllowedStreamAssetAttachmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StreamAssetAttachmentStatus) IsValid() bool {
	for _, existing := range AllowedStreamAssetAttachmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StreamAsset_attachmentStatus value
func (v StreamAssetAttachmentStatus) Ptr() *StreamAssetAttachmentStatus {
	return &v
}

type NullableStreamAssetAttachmentStatus struct {
	value *StreamAssetAttachmentStatus
	isSet bool
}

func (v NullableStreamAssetAttachmentStatus) Get() *StreamAssetAttachmentStatus {
	return v.value
}

func (v *NullableStreamAssetAttachmentStatus) Set(val *StreamAssetAttachmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamAssetAttachmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamAssetAttachmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamAssetAttachmentStatus(val *StreamAssetAttachmentStatus) *NullableStreamAssetAttachmentStatus {
	return &NullableStreamAssetAttachmentStatus{value: val, isSet: true}
}

func (v NullableStreamAssetAttachmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamAssetAttachmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}