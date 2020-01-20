/*
Package config provides a Client to manage AEM configuration using the osgi-service-json-config service
To see all the configuration this package manage see the doc for the repo https://github.com/xumak-grid/osgi-service-json-config
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// SlingSuffix is the service path endpoint must start with /
	SlingSuffix = "/system/instanceConfiguration.servlet.html"
	DefaultHost = "localhost"

	PolicyCreate         = "CREATE"
	PolicyShow           = "SHOW"
	PolicyUpdate         = "UPDATE"
	PolicyDelete         = "DELETE"
	PolicyChangePassword = "CHANGE_PASSWORD"
	PolicyUpdateCreate   = "UPDATE_CREATE_POLICY"

	TypeOSGI    = "osgi"
	TypeIternal = "internal"
	TypeRoot    = "root"
	TypePublsih = "publish"
	TypeAuthor  = "author"
)

// Client represents an authenticated client to perfome some config in the AEM instance
type Client struct {
	Users   []*User  `json:"users,omitempty"`
	Configs []*OSGI  `json:"configs,omitempty"`
	Agents  []*Agent `json:"replicationAgents,omitempty"`
}

// RegisterUser adds one or more user config to the current client
func (c *Client) RegisterUser(users ...*User) {
	if c != nil {
		c.Users = append(c.Users, users...)
	}
}

// RegisterOSGI adds one or more OSGI config to the current client
func (c *Client) RegisterOSGI(configs ...*OSGI) {
	if c != nil {
		c.Configs = append(c.Configs, configs...)
	}
}

// RegisterAgent adds one or more replication config to the current client
func (c *Client) RegisterAgent(agents ...*Agent) {
	if c != nil {
		c.Agents = append(c.Agents, agents...)
	}
}

// Do applies the configuration for client using POST request
func (c *Client) Do(host, port, user, password string) (*Response, error) {

	if c == nil {
		return nil, errors.New("client is nil")
	}
	if len(c.Agents) == 0 && len(c.Configs) == 0 && len(c.Users) == 0 {
		return nil, errors.New("client with nothing to do")
	}
	if host == "" {
		host = DefaultHost
	}
	if user == "" || password == "" {
		return nil, errors.New("user and password are required")
	}
	url := ""
	if port == "" {
		url = fmt.Sprintf("http://%v%v", host, SlingSuffix)
	} else {
		url = fmt.Sprintf("http://%v:%v%v", host, port, SlingSuffix)
	}

	// Check if With and ACLS values exist in the configuration
	err := checkRequiredValues(c)
	if err != nil {
		return nil, err
	}

	// wrap the client to fix the service unnecessary key
	// Temporal fix, The service contains an unnecessary key at the beginning
	// This fix can be removed when the Grid service is fixed
	type Definition struct {
		Fix *Client `json:"fix,omitempty"`
	}
	def := new(Definition)
	def.Fix = c

	b, err := json.Marshal(def)
	if err != nil {
		return nil, errors.New("marshal error: " + err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(user, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("an error detected reading the response body: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error detected status code: %v error: %v", resp.Status, string(body))
	}
	outResponse := new(Response)
	err = json.Unmarshal(body, outResponse)
	if err != nil {
		return nil, err
	}
	outResponse.JSONOutput = string(body)
	return outResponse, nil
}
