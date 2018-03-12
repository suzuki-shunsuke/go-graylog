package graylog

import (
	"fmt"
	"net/url"
	"path"
)

// Endpoints represents each API's endpoint URLs.
type Endpoints struct {
	Endpoint       *url.URL
	Roles          string
	Users          string
	Inputs         string
	IndexSets      string
	Streams        string
	EnabledStreams string
}

func NewEndpoints(endpoint string) (*Endpoints, error) {
	base, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &Endpoints{
		Roles:          getEndpoint(*base, "/roles"),
		Users:          getEndpoint(*base, "/users"),
		Inputs:         getEndpoint(*base, "/system/inputs"),
		IndexSets:      getEndpoint(*base, "/system/indices/index_sets"),
		Streams:        getEndpoint(*base, "/streams"),
		EnabledStreams: getEndpoint(*base, "/streams/enabled"),
		Endpoint:       base,
	}, nil
}

func getEndpoint(u url.URL, p string) string {
	u.Path = path.Join(u.Path, p)
	return u.String()
}

// SetDefaultIndexSet returns SetDefaultIndexSet API's endpoint url.
func (endpoint *Endpoints) SetDefaultIndexSet(id string) string {
	return fmt.Sprintf("%s/%s/default", endpoint.IndexSets, id)
}

// IndexSetsStats returns all IndexSets stats API's endpoint url.
func (endpoint *Endpoints) IndexSetsStats() string {
	return fmt.Sprintf("%s/stats", endpoint.IndexSets)
}

// IndexSetsStats returns an IndexSet stats API's endpoint url.
func (endpoint *Endpoints) IndexSetStats(id string) string {
	return fmt.Sprintf("%s/%s/stats", endpoint.IndexSets, id)
}

// IndexSet returns an IndexSet API's endpoint url.
func (endpoint *Endpoints) IndexSet(id string) string {
	return fmt.Sprintf("%s/%s", endpoint.IndexSets, id)
}

// Input returns an Input API's endpoint url.
func (endpoint *Endpoints) Input(id string) string {
	return fmt.Sprintf("%s/%s", endpoint.Inputs, id)
}

// Input returns a User API's endpoint url.
func (endpoint *Endpoints) User(name string) string {
	return fmt.Sprintf("%s/%s", endpoint.Users, name)
}

// Input returns a Role API's endpoint url.
func (endpoint *Endpoints) Role(name string) string {
	return fmt.Sprintf("%s/%s", endpoint.Roles, name)
}

// Stream returns a Stream API's endpoint url.
func (endpoint *Endpoints) Stream(id string) string {
	return fmt.Sprintf("%s/%s", endpoint.Streams, id)
}

// StreamRules returns Stream Rules API's endpoint url.
func (endpoint *Endpoints) StreamRules(streamID string) string {
	// /streams/{streamid}/rules
	return fmt.Sprintf("%s/%s/rules", endpoint.Streams, streamID)
}

// StreamRuleTypes returns Stream Rule Types API's endpoint url.
func (endpoint *Endpoints) StreamRuleTypes(streamID string) string {
	// /streams/{streamid}/rules/types
	return fmt.Sprintf("%s/%s/rules/types", endpoint.Streams, streamID)
}

// StreamRule returns a Stream Rule API's endpoint url.
func (endpoint *Endpoints) StreamRule(streamID, streamRuleID string) string {
	// /streams/{streamid}/rules/{streamRuleID}
	return fmt.Sprintf(
		"%s/%s/rules/%s", endpoint.Streams, streamID, streamRuleID)
}

// PauseStream returns PauseStream API's endpoint url.
func (endpoint *Endpoints) PauseStream(id string) string {
	return fmt.Sprintf("%s/%s/pause", endpoint.Streams, id)
}

// ResumeStream returns ResumeStream API's endpoint url.
func (endpoint *Endpoints) ResumeStream(id string) string {
	return fmt.Sprintf("%s/%s/resume", endpoint.Streams, id)
}

// RoleMembers returns given role's member endpoint url.
func (endpoint *Endpoints) RoleMembers(name string) string {
	return fmt.Sprintf("%s/%s/members", endpoint.Roles, name)
}

// RoleMember returns given role member endpoint url.
func (endpoint *Endpoints) RoleMember(userName, roleName string) string {
	return fmt.Sprintf("%s/%s/members/%s", endpoint.Roles, roleName, userName)
}
