package config

// OSGI represents an OSGI configuration
type OSGI struct {
	Policy string                 `json:"policy"`
	Type   string                 `json:"type,omitempty"`
	ID     string                 `json:"id"`
	Path   string                 `json:"path"`
	With   map[string]interface{} `json:"with,omitempty"`
	// The following keys are used to the response
	Ok    string   `json:"ok,omitempty"`
	Error []string `json:"error,omitempty"`
}

// NewOSGI returns a basic OSGI configuration
// the ID is the full bundle name for example:
// com.xumak.jcrsyncr.engine.SyncControllerImpl6.config
// the path is the location of the configuration example /apps/system/config
func NewOSGI(id, path string) *OSGI {
	return &OSGI{
		Policy: PolicyUpdateCreate,
		Type:   TypeOSGI,
		ID:     id,
		Path:   path,
	}
}
