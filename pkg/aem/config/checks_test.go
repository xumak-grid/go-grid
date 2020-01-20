package config

import (
	"testing"
)

func TestCheckUser(t *testing.T) {
	client := &Client{}

	userValid := NewUser("", "")
	userValid.With = map[string]interface{}{"xx": "ss"}
	userValid.ACLS = &ACL{}

	userInvalid := NewUser("", "")
	userInvalid.With = map[string]interface{}{"xx": "ss"}

	userInvalid2 := NewUser("", "")
	userInvalid2.ACLS = &ACL{}

	table := []struct {
		name        string
		user        *User
		returnError bool
	}{
		{"user without 'With' and 'ACLS'", NewUser("dummy", "dummy"), true},
		{"user only with 'With'", userInvalid, true},
		{"user only with 'ACLS'", userInvalid2, true},
		{"user with 'With' and 'ACLS'", userValid, false},
	}

	for _, i := range table {
		client.RegisterUser(i.user)
		err := checkRequiredValues(client)
		if err == nil && i.returnError {
			t.Errorf("%v should return an error", i.name)
		}
		if err != nil && !i.returnError {
			t.Errorf("%v should not return an error, got: %v", i.name, err.Error())
		}
		// for the next iteration
		unregisterClient(client)
	}
}

func TestCheckOSGI(t *testing.T) {
	client := &Client{}

	OSGIValid := NewOSGI("dummy", "dummy")
	OSGIValid.With = map[string]interface{}{"xx": "ss"}

	table := []struct {
		name        string
		osgi        *OSGI
		returnError bool
	}{
		{"OSGI without 'With'", NewOSGI("dummy", "dummy"), true},
		{"OSGI with 'With'", OSGIValid, false},
	}

	for _, i := range table {
		client.RegisterOSGI(i.osgi)
		err := checkRequiredValues(client)
		if err == nil && i.returnError {
			t.Errorf("%v should return an error", i.name)
		}
		if err != nil && !i.returnError {
			t.Errorf("%v should not return an error, got: %v", i.name, err.Error())
		}
		// for the next iteration
		unregisterClient(client)
	}
}

func TestCheckAgent(t *testing.T) {
	client := &Client{}

	AgentValid := NewAgentAuthor("", PolicyCreate)
	AgentValid.With = map[string]interface{}{"xx": "ss"}
	AgentValid2 := NewAgentPublish("", PolicyUpdate)
	AgentValid2.With = map[string]interface{}{"xx": "ss"}

	table := []struct {
		name        string
		agent       *Agent
		returnError bool
	}{
		{"agent without 'With' PolicyCreate", NewAgentAuthor("", PolicyCreate), true},
		{"agent without 'With' PolicyUpdate", NewAgentPublish("", PolicyUpdate), true},
		{"agent valid PolicyCreate", AgentValid, false},
		{"agent valid PolicyUpdate'", AgentValid2, false},
	}

	for _, i := range table {
		client.RegisterAgent(i.agent)
		err := checkRequiredValues(client)
		if err == nil && i.returnError {
			t.Errorf("%v should return an error", i.name)
		}
		if err != nil && !i.returnError {
			t.Errorf("%v should not return an error, got: %v", i.name, err.Error())
		}
		// for the next iteration
		unregisterClient(client)
	}
}

func unregisterClient(c *Client) *Client {
	c.Agents = nil
	c.Configs = nil
	c.Users = nil
	return c
}
