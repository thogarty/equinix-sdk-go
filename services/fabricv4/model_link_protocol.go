/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// LinkProtocol - Connection link protocol Configuration
type LinkProtocol struct {
	LinkProtocolDot1q     *LinkProtocolDot1q
	LinkProtocolEvpnVxlan *LinkProtocolEvpnVxlan
	LinkProtocolQinq      *LinkProtocolQinq
	LinkProtocolUntagged  *LinkProtocolUntagged
	LinkProtocolVxlan     *LinkProtocolVxlan
}

// LinkProtocolDot1qAsLinkProtocol is a convenience function that returns LinkProtocolDot1q wrapped in LinkProtocol
func LinkProtocolDot1qAsLinkProtocol(v *LinkProtocolDot1q) LinkProtocol {
	return LinkProtocol{
		LinkProtocolDot1q: v,
	}
}

// LinkProtocolEvpnVxlanAsLinkProtocol is a convenience function that returns LinkProtocolEvpnVxlan wrapped in LinkProtocol
func LinkProtocolEvpnVxlanAsLinkProtocol(v *LinkProtocolEvpnVxlan) LinkProtocol {
	return LinkProtocol{
		LinkProtocolEvpnVxlan: v,
	}
}

// LinkProtocolQinqAsLinkProtocol is a convenience function that returns LinkProtocolQinq wrapped in LinkProtocol
func LinkProtocolQinqAsLinkProtocol(v *LinkProtocolQinq) LinkProtocol {
	return LinkProtocol{
		LinkProtocolQinq: v,
	}
}

// LinkProtocolUntaggedAsLinkProtocol is a convenience function that returns LinkProtocolUntagged wrapped in LinkProtocol
func LinkProtocolUntaggedAsLinkProtocol(v *LinkProtocolUntagged) LinkProtocol {
	return LinkProtocol{
		LinkProtocolUntagged: v,
	}
}

// LinkProtocolVxlanAsLinkProtocol is a convenience function that returns LinkProtocolVxlan wrapped in LinkProtocol
func LinkProtocolVxlanAsLinkProtocol(v *LinkProtocolVxlan) LinkProtocol {
	return LinkProtocol{
		LinkProtocolVxlan: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LinkProtocol) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LinkProtocolDot1q
	err = newStrictDecoder(data).Decode(&dst.LinkProtocolDot1q)
	if err == nil {
		jsonLinkProtocolDot1q, _ := json.Marshal(dst.LinkProtocolDot1q)
		if string(jsonLinkProtocolDot1q) == "{}" { // empty struct
			dst.LinkProtocolDot1q = nil
		} else {
			match++
		}
	} else {
		dst.LinkProtocolDot1q = nil
	}

	// try to unmarshal data into LinkProtocolEvpnVxlan
	err = newStrictDecoder(data).Decode(&dst.LinkProtocolEvpnVxlan)
	if err == nil {
		jsonLinkProtocolEvpnVxlan, _ := json.Marshal(dst.LinkProtocolEvpnVxlan)
		if string(jsonLinkProtocolEvpnVxlan) == "{}" { // empty struct
			dst.LinkProtocolEvpnVxlan = nil
		} else {
			match++
		}
	} else {
		dst.LinkProtocolEvpnVxlan = nil
	}

	// try to unmarshal data into LinkProtocolQinq
	err = newStrictDecoder(data).Decode(&dst.LinkProtocolQinq)
	if err == nil {
		jsonLinkProtocolQinq, _ := json.Marshal(dst.LinkProtocolQinq)
		if string(jsonLinkProtocolQinq) == "{}" { // empty struct
			dst.LinkProtocolQinq = nil
		} else {
			match++
		}
	} else {
		dst.LinkProtocolQinq = nil
	}

	// try to unmarshal data into LinkProtocolUntagged
	err = newStrictDecoder(data).Decode(&dst.LinkProtocolUntagged)
	if err == nil {
		jsonLinkProtocolUntagged, _ := json.Marshal(dst.LinkProtocolUntagged)
		if string(jsonLinkProtocolUntagged) == "{}" { // empty struct
			dst.LinkProtocolUntagged = nil
		} else {
			match++
		}
	} else {
		dst.LinkProtocolUntagged = nil
	}

	// try to unmarshal data into LinkProtocolVxlan
	err = newStrictDecoder(data).Decode(&dst.LinkProtocolVxlan)
	if err == nil {
		jsonLinkProtocolVxlan, _ := json.Marshal(dst.LinkProtocolVxlan)
		if string(jsonLinkProtocolVxlan) == "{}" { // empty struct
			dst.LinkProtocolVxlan = nil
		} else {
			match++
		}
	} else {
		dst.LinkProtocolVxlan = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LinkProtocolDot1q = nil
		dst.LinkProtocolEvpnVxlan = nil
		dst.LinkProtocolQinq = nil
		dst.LinkProtocolUntagged = nil
		dst.LinkProtocolVxlan = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LinkProtocol)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LinkProtocol)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LinkProtocol) MarshalJSON() ([]byte, error) {
	if src.LinkProtocolDot1q != nil {
		return json.Marshal(&src.LinkProtocolDot1q)
	}

	if src.LinkProtocolEvpnVxlan != nil {
		return json.Marshal(&src.LinkProtocolEvpnVxlan)
	}

	if src.LinkProtocolQinq != nil {
		return json.Marshal(&src.LinkProtocolQinq)
	}

	if src.LinkProtocolUntagged != nil {
		return json.Marshal(&src.LinkProtocolUntagged)
	}

	if src.LinkProtocolVxlan != nil {
		return json.Marshal(&src.LinkProtocolVxlan)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LinkProtocol) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LinkProtocolDot1q != nil {
		return obj.LinkProtocolDot1q
	}

	if obj.LinkProtocolEvpnVxlan != nil {
		return obj.LinkProtocolEvpnVxlan
	}

	if obj.LinkProtocolQinq != nil {
		return obj.LinkProtocolQinq
	}

	if obj.LinkProtocolUntagged != nil {
		return obj.LinkProtocolUntagged
	}

	if obj.LinkProtocolVxlan != nil {
		return obj.LinkProtocolVxlan
	}

	// all schemas are nil
	return nil
}

type NullableLinkProtocol struct {
	value *LinkProtocol
	isSet bool
}

func (v NullableLinkProtocol) Get() *LinkProtocol {
	return v.value
}

func (v *NullableLinkProtocol) Set(val *LinkProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkProtocol(val *LinkProtocol) *NullableLinkProtocol {
	return &NullableLinkProtocol{value: val, isSet: true}
}

func (v NullableLinkProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
