/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the AuthTokenUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenUser{}

// AuthTokenUser struct for AuthTokenUser
type AuthTokenUser struct {
	AvatarThumbUrl        *string                `json:"avatar_thumb_url,omitempty"`
	AvatarUrl             *string                `json:"avatar_url,omitempty"`
	CreatedAt             *time.Time             `json:"created_at,omitempty"`
	Customdata            map[string]interface{} `json:"customdata,omitempty"`
	DefaultOrganizationId *string                `json:"default_organization_id,omitempty"`
	DefaultProjectId      *string                `json:"default_project_id,omitempty"`
	Email                 *string                `json:"email,omitempty"`
	Emails                []Href                 `json:"emails,omitempty"`
	FirstName             *string                `json:"first_name,omitempty"`
	FraudScore            *string                `json:"fraud_score,omitempty"`
	FullName              *string                `json:"full_name,omitempty"`
	Href                  *string                `json:"href,omitempty"`
	Id                    *string                `json:"id,omitempty"`
	LastLoginAt           *time.Time             `json:"last_login_at,omitempty"`
	LastName              *string                `json:"last_name,omitempty"`
	MaxOrganizations      *int32                 `json:"max_organizations,omitempty"`
	MaxProjects           *int32                 `json:"max_projects,omitempty"`
	PhoneNumber           *string                `json:"phone_number,omitempty"`
	ShortId               *string                `json:"short_id,omitempty"`
	Timezone              *string                `json:"timezone,omitempty"`
	TwoFactorAuth         *string                `json:"two_factor_auth,omitempty"`
	UpdatedAt             *time.Time             `json:"updated_at,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AuthTokenUser AuthTokenUser

// NewAuthTokenUser instantiates a new AuthTokenUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenUser() *AuthTokenUser {
	this := AuthTokenUser{}
	return &this
}

// NewAuthTokenUserWithDefaults instantiates a new AuthTokenUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenUserWithDefaults() *AuthTokenUser {
	this := AuthTokenUser{}
	return &this
}

// GetAvatarThumbUrl returns the AvatarThumbUrl field value if set, zero value otherwise.
func (o *AuthTokenUser) GetAvatarThumbUrl() string {
	if o == nil || IsNil(o.AvatarThumbUrl) {
		var ret string
		return ret
	}
	return *o.AvatarThumbUrl
}

// GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetAvatarThumbUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarThumbUrl) {
		return nil, false
	}
	return o.AvatarThumbUrl, true
}

// HasAvatarThumbUrl returns a boolean if a field has been set.
func (o *AuthTokenUser) HasAvatarThumbUrl() bool {
	if o != nil && !IsNil(o.AvatarThumbUrl) {
		return true
	}

	return false
}

// SetAvatarThumbUrl gets a reference to the given string and assigns it to the AvatarThumbUrl field.
func (o *AuthTokenUser) SetAvatarThumbUrl(v string) {
	o.AvatarThumbUrl = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *AuthTokenUser) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *AuthTokenUser) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *AuthTokenUser) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuthTokenUser) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuthTokenUser) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuthTokenUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *AuthTokenUser) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *AuthTokenUser) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *AuthTokenUser) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDefaultOrganizationId returns the DefaultOrganizationId field value if set, zero value otherwise.
func (o *AuthTokenUser) GetDefaultOrganizationId() string {
	if o == nil || IsNil(o.DefaultOrganizationId) {
		var ret string
		return ret
	}
	return *o.DefaultOrganizationId
}

// GetDefaultOrganizationIdOk returns a tuple with the DefaultOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetDefaultOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultOrganizationId) {
		return nil, false
	}
	return o.DefaultOrganizationId, true
}

// HasDefaultOrganizationId returns a boolean if a field has been set.
func (o *AuthTokenUser) HasDefaultOrganizationId() bool {
	if o != nil && !IsNil(o.DefaultOrganizationId) {
		return true
	}

	return false
}

// SetDefaultOrganizationId gets a reference to the given string and assigns it to the DefaultOrganizationId field.
func (o *AuthTokenUser) SetDefaultOrganizationId(v string) {
	o.DefaultOrganizationId = &v
}

// GetDefaultProjectId returns the DefaultProjectId field value if set, zero value otherwise.
func (o *AuthTokenUser) GetDefaultProjectId() string {
	if o == nil || IsNil(o.DefaultProjectId) {
		var ret string
		return ret
	}
	return *o.DefaultProjectId
}

// GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetDefaultProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultProjectId) {
		return nil, false
	}
	return o.DefaultProjectId, true
}

// HasDefaultProjectId returns a boolean if a field has been set.
func (o *AuthTokenUser) HasDefaultProjectId() bool {
	if o != nil && !IsNil(o.DefaultProjectId) {
		return true
	}

	return false
}

// SetDefaultProjectId gets a reference to the given string and assigns it to the DefaultProjectId field.
func (o *AuthTokenUser) SetDefaultProjectId(v string) {
	o.DefaultProjectId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthTokenUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthTokenUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthTokenUser) SetEmail(v string) {
	o.Email = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *AuthTokenUser) GetEmails() []Href {
	if o == nil || IsNil(o.Emails) {
		var ret []Href
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetEmailsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *AuthTokenUser) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []Href and assigns it to the Emails field.
func (o *AuthTokenUser) SetEmails(v []Href) {
	o.Emails = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AuthTokenUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AuthTokenUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AuthTokenUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFraudScore returns the FraudScore field value if set, zero value otherwise.
func (o *AuthTokenUser) GetFraudScore() string {
	if o == nil || IsNil(o.FraudScore) {
		var ret string
		return ret
	}
	return *o.FraudScore
}

// GetFraudScoreOk returns a tuple with the FraudScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetFraudScoreOk() (*string, bool) {
	if o == nil || IsNil(o.FraudScore) {
		return nil, false
	}
	return o.FraudScore, true
}

// HasFraudScore returns a boolean if a field has been set.
func (o *AuthTokenUser) HasFraudScore() bool {
	if o != nil && !IsNil(o.FraudScore) {
		return true
	}

	return false
}

// SetFraudScore gets a reference to the given string and assigns it to the FraudScore field.
func (o *AuthTokenUser) SetFraudScore(v string) {
	o.FraudScore = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *AuthTokenUser) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *AuthTokenUser) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *AuthTokenUser) SetFullName(v string) {
	o.FullName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AuthTokenUser) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AuthTokenUser) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AuthTokenUser) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthTokenUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthTokenUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthTokenUser) SetId(v string) {
	o.Id = &v
}

// GetLastLoginAt returns the LastLoginAt field value if set, zero value otherwise.
func (o *AuthTokenUser) GetLastLoginAt() time.Time {
	if o == nil || IsNil(o.LastLoginAt) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginAt
}

// GetLastLoginAtOk returns a tuple with the LastLoginAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetLastLoginAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLoginAt) {
		return nil, false
	}
	return o.LastLoginAt, true
}

// HasLastLoginAt returns a boolean if a field has been set.
func (o *AuthTokenUser) HasLastLoginAt() bool {
	if o != nil && !IsNil(o.LastLoginAt) {
		return true
	}

	return false
}

// SetLastLoginAt gets a reference to the given time.Time and assigns it to the LastLoginAt field.
func (o *AuthTokenUser) SetLastLoginAt(v time.Time) {
	o.LastLoginAt = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AuthTokenUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AuthTokenUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AuthTokenUser) SetLastName(v string) {
	o.LastName = &v
}

// GetMaxOrganizations returns the MaxOrganizations field value if set, zero value otherwise.
func (o *AuthTokenUser) GetMaxOrganizations() int32 {
	if o == nil || IsNil(o.MaxOrganizations) {
		var ret int32
		return ret
	}
	return *o.MaxOrganizations
}

// GetMaxOrganizationsOk returns a tuple with the MaxOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetMaxOrganizationsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxOrganizations) {
		return nil, false
	}
	return o.MaxOrganizations, true
}

// HasMaxOrganizations returns a boolean if a field has been set.
func (o *AuthTokenUser) HasMaxOrganizations() bool {
	if o != nil && !IsNil(o.MaxOrganizations) {
		return true
	}

	return false
}

// SetMaxOrganizations gets a reference to the given int32 and assigns it to the MaxOrganizations field.
func (o *AuthTokenUser) SetMaxOrganizations(v int32) {
	o.MaxOrganizations = &v
}

// GetMaxProjects returns the MaxProjects field value if set, zero value otherwise.
func (o *AuthTokenUser) GetMaxProjects() int32 {
	if o == nil || IsNil(o.MaxProjects) {
		var ret int32
		return ret
	}
	return *o.MaxProjects
}

// GetMaxProjectsOk returns a tuple with the MaxProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetMaxProjectsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxProjects) {
		return nil, false
	}
	return o.MaxProjects, true
}

// HasMaxProjects returns a boolean if a field has been set.
func (o *AuthTokenUser) HasMaxProjects() bool {
	if o != nil && !IsNil(o.MaxProjects) {
		return true
	}

	return false
}

// SetMaxProjects gets a reference to the given int32 and assigns it to the MaxProjects field.
func (o *AuthTokenUser) SetMaxProjects(v int32) {
	o.MaxProjects = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AuthTokenUser) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AuthTokenUser) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AuthTokenUser) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetShortId returns the ShortId field value if set, zero value otherwise.
func (o *AuthTokenUser) GetShortId() string {
	if o == nil || IsNil(o.ShortId) {
		var ret string
		return ret
	}
	return *o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetShortIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShortId) {
		return nil, false
	}
	return o.ShortId, true
}

// HasShortId returns a boolean if a field has been set.
func (o *AuthTokenUser) HasShortId() bool {
	if o != nil && !IsNil(o.ShortId) {
		return true
	}

	return false
}

// SetShortId gets a reference to the given string and assigns it to the ShortId field.
func (o *AuthTokenUser) SetShortId(v string) {
	o.ShortId = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AuthTokenUser) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AuthTokenUser) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AuthTokenUser) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTwoFactorAuth returns the TwoFactorAuth field value if set, zero value otherwise.
func (o *AuthTokenUser) GetTwoFactorAuth() string {
	if o == nil || IsNil(o.TwoFactorAuth) {
		var ret string
		return ret
	}
	return *o.TwoFactorAuth
}

// GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetTwoFactorAuthOk() (*string, bool) {
	if o == nil || IsNil(o.TwoFactorAuth) {
		return nil, false
	}
	return o.TwoFactorAuth, true
}

// HasTwoFactorAuth returns a boolean if a field has been set.
func (o *AuthTokenUser) HasTwoFactorAuth() bool {
	if o != nil && !IsNil(o.TwoFactorAuth) {
		return true
	}

	return false
}

// SetTwoFactorAuth gets a reference to the given string and assigns it to the TwoFactorAuth field.
func (o *AuthTokenUser) SetTwoFactorAuth(v string) {
	o.TwoFactorAuth = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuthTokenUser) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenUser) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuthTokenUser) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AuthTokenUser) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o AuthTokenUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvatarThumbUrl) {
		toSerialize["avatar_thumb_url"] = o.AvatarThumbUrl
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.DefaultOrganizationId) {
		toSerialize["default_organization_id"] = o.DefaultOrganizationId
	}
	if !IsNil(o.DefaultProjectId) {
		toSerialize["default_project_id"] = o.DefaultProjectId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.FraudScore) {
		toSerialize["fraud_score"] = o.FraudScore
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastLoginAt) {
		toSerialize["last_login_at"] = o.LastLoginAt
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.MaxOrganizations) {
		toSerialize["max_organizations"] = o.MaxOrganizations
	}
	if !IsNil(o.MaxProjects) {
		toSerialize["max_projects"] = o.MaxProjects
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.ShortId) {
		toSerialize["short_id"] = o.ShortId
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.TwoFactorAuth) {
		toSerialize["two_factor_auth"] = o.TwoFactorAuth
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthTokenUser) UnmarshalJSON(bytes []byte) (err error) {
	varAuthTokenUser := _AuthTokenUser{}

	if err = json.Unmarshal(bytes, &varAuthTokenUser); err == nil {
		*o = AuthTokenUser(varAuthTokenUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "avatar_thumb_url")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "default_organization_id")
		delete(additionalProperties, "default_project_id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "fraud_score")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_login_at")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "max_organizations")
		delete(additionalProperties, "max_projects")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "short_id")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "two_factor_auth")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthTokenUser struct {
	value *AuthTokenUser
	isSet bool
}

func (v NullableAuthTokenUser) Get() *AuthTokenUser {
	return v.value
}

func (v *NullableAuthTokenUser) Set(val *AuthTokenUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenUser(val *AuthTokenUser) *NullableAuthTokenUser {
	return &NullableAuthTokenUser{value: val, isSet: true}
}

func (v NullableAuthTokenUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
