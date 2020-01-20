package config

// Response represents the response of the service after Do
type Response struct {
	// JSONOutput the response in JSON format
	JSONOutput string         `json:"-"`
	Status     string         `json:"status"`
	Message    string         `json:"message"`
	Data       ClientResponse `json:"data,omitempty"`
}

// ClientResponse represents an authenticated client to perfome some config in the AEM instance
type ClientResponse struct {
	Users   []*User                  `json:"users,omitempty"`
	Configs []*OSGI                  `json:"configs,omitempty"`
	Agents  []map[string]interface{} `json:"replicationAgents,omitempty"`
}
