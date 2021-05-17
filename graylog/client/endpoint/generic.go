package endpoint

import "strings"

// Generic returns the URL of a generic endpoint without encoding
func (ep *Endpoints) Generic(endpoint string) string {
	// /{endpoint}

	endpoint = strings.TrimLeft(endpoint, "/")
	return ep.root + "/" + endpoint
}
