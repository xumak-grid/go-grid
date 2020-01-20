package config

import "errors"
import "fmt"

// checkRequiredValues Checks if 'with' and 'acls' values are present in the required policies
func checkRequiredValues(c *Client) error {

	// For OSGI conf 'with' value is required
	for _, osgi := range c.Configs {
		if len(osgi.With) == 0 {
			return errors.New("OSGI config requires 'with' value")
		}
	}
	// For User conf 'with' and 'acls' are required in CREATE and UPDATE policies
	for _, user := range c.Users {
		if (user.Policy == PolicyCreate || user.Policy == PolicyUpdate) && (len(user.With) == 0 || user.ACLS == nil) {
			return fmt.Errorf("keys 'with' and 'acls' are required in user %v for policy %v", user.ID, user.Policy)
		}
	}
	// For Replication Agent 'with' is required with CREATE and UPDATE policies
	for _, agent := range c.Agents {
		if (agent.Policy == PolicyCreate || agent.Policy == PolicyUpdate) && len(agent.With) == 0 {
			return fmt.Errorf("key 'with' is required in user %v with policy %v", agent.Name, agent.Policy)
		}
	}
	return nil

}
