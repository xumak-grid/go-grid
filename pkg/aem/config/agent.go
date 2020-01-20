package config

// Agent represent a replication agent in the client
type Agent struct {
	Environment string                 `json:"type"`
	Name        string                 `json:"name,omitempty"`
	Policy      string                 `json:"policy"`
	With        map[string]interface{} `json:"with,omitempty"`
	// The follwing keys are used to the response
	Ok    string   `json:"ok,omitempty"`
	Error []string `json:"error,omitempty"`
}

// agent return a new Agent with the configuration
func agent(name, environment, policy string) *Agent {
	return &Agent{
		Environment: environment,
		Name:        name,
		Policy:      policy,
	}
}

// NewAgentPublish returns a replication agent with publish config
// for the policies use the constants defined in this package
func NewAgentPublish(name, policy string) *Agent {
	return agent(name, TypePublsih, policy)
}

// NewAgentAuthor returns a replication agent with author config
// for the policies use the constants defined in this package
func NewAgentAuthor(name, policy string) *Agent {
	return agent(name, TypeAuthor, policy)
}

// NewAgentShowAll returns a replication agent with SHOW config
func NewAgentShowAll(environment string) *Agent {
	return agent("", environment, PolicyShow)
}
