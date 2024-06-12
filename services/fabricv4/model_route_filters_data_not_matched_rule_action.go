/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// RouteFiltersDataNotMatchedRuleAction the model 'RouteFiltersDataNotMatchedRuleAction'
type RouteFiltersDataNotMatchedRuleAction string

// List of RouteFiltersData_notMatchedRuleAction
const (
	ROUTEFILTERSDATANOTMATCHEDRULEACTION_ALLOW RouteFiltersDataNotMatchedRuleAction = "ALLOW"
	ROUTEFILTERSDATANOTMATCHEDRULEACTION_DENY  RouteFiltersDataNotMatchedRuleAction = "DENY"
)

// All allowed values of RouteFiltersDataNotMatchedRuleAction enum
var AllowedRouteFiltersDataNotMatchedRuleActionEnumValues = []RouteFiltersDataNotMatchedRuleAction{
	"ALLOW",
	"DENY",
}

func (v *RouteFiltersDataNotMatchedRuleAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteFiltersDataNotMatchedRuleAction(value)
	for _, existing := range AllowedRouteFiltersDataNotMatchedRuleActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteFiltersDataNotMatchedRuleAction", value)
}

// NewRouteFiltersDataNotMatchedRuleActionFromValue returns a pointer to a valid RouteFiltersDataNotMatchedRuleAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteFiltersDataNotMatchedRuleActionFromValue(v string) (*RouteFiltersDataNotMatchedRuleAction, error) {
	ev := RouteFiltersDataNotMatchedRuleAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteFiltersDataNotMatchedRuleAction: valid values are %v", v, AllowedRouteFiltersDataNotMatchedRuleActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteFiltersDataNotMatchedRuleAction) IsValid() bool {
	for _, existing := range AllowedRouteFiltersDataNotMatchedRuleActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteFiltersData_notMatchedRuleAction value
func (v RouteFiltersDataNotMatchedRuleAction) Ptr() *RouteFiltersDataNotMatchedRuleAction {
	return &v
}

type NullableRouteFiltersDataNotMatchedRuleAction struct {
	value *RouteFiltersDataNotMatchedRuleAction
	isSet bool
}

func (v NullableRouteFiltersDataNotMatchedRuleAction) Get() *RouteFiltersDataNotMatchedRuleAction {
	return v.value
}

func (v *NullableRouteFiltersDataNotMatchedRuleAction) Set(val *RouteFiltersDataNotMatchedRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFiltersDataNotMatchedRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFiltersDataNotMatchedRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFiltersDataNotMatchedRuleAction(val *RouteFiltersDataNotMatchedRuleAction) *NullableRouteFiltersDataNotMatchedRuleAction {
	return &NullableRouteFiltersDataNotMatchedRuleAction{value: val, isSet: true}
}

func (v NullableRouteFiltersDataNotMatchedRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFiltersDataNotMatchedRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
