package config_test

import (
	"fmt"
	"os"

	aem "github.com/xumak-grid/go-grid/pkg/aem/config"
)

func Example() {
	var c aem.Client
	var out *aem.Response
	var err error

	// 1. Change an user password
	c = aem.Client{}
	user := aem.NewUserChangePassword("admin", "12345")
	c.RegisterUser(user)

	out, err = c.Do("localhost", "4502", "admin", "admin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, user := range out.Data.Users {
		if user.Ok != "" {
			fmt.Println(user.Ok)
		} else {
			fmt.Println(user.Error)
		}
	}

	// 2. Create a new user
	user = aem.NewUser("author", "123")
	user.With = map[string]interface{}{
		"profile/aboutMe": "Used the Replication Agents to send everything under /content/*",
	}
	user.ACLS = &aem.ACL{
		Deny: map[string]interface{}{
			"/": []string{"C", "R", "U", "D", "X", "R*", "U*"},
		},
		Allow: map[string]interface{}{
			"/content": []string{"R"},
		},
	}
	c.RegisterUser(user)

	out, err = c.Do("localhost", "4502", "admin", "12345")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, user := range out.Data.Users {
		if user.Ok != "" {
			fmt.Println(user.Ok)
		} else {
			fmt.Println(user.Error)
		}
	}

	// 3. OSGI Configuraition
	c = aem.Client{}
	osgi := aem.NewOSGI("com.xumak.jcrsyncr.engine.SyncControllerImpl6.config", "/apps/system/config")
	osgi.With = map[string]interface{}{
		"jcrsync.definitions":    []string{"/BedrocK/XCQB/Demo/Source/myCompany/CQFiles/myCompany;/apps/grid"},
		"jcrsync.establish.sync": true,
		"foo":    "bar",
		"number": 123,
	}
	c.RegisterOSGI(osgi)

	out, err = c.Do("localhost", "4502", "admin", "12345")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, osgi := range out.Data.Configs {
		if osgi.Ok != "" {
			fmt.Println(osgi.Ok)
		} else {
			fmt.Println(osgi.Error)
		}
	}

	// 4. Replication Agent Configuration
	agent := aem.NewAgentPublish("toProdPublish1", aem.PolicyCreate)
	agent.With = map[string]interface{}{
		"jcr:title":         "Replication Agent for ProdPublish1",
		"enabled":           true,
		"userId":            "allContentSender",
		"host":              "prod.publish1.customer.xumak.cloud",
		"port":              5501,
		"transportUser":     "allContentReceiver",
		"transportPassword": "zzhV4gaDuxF9HUcdDaJhajjA",
		"currentVersion":    "832-040312048-1pou4poiup4oi123u4pio13up4",
	}
	c.RegisterAgent(agent)

	out, err = c.Do("localhost", "4502", "admin", "12345")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, agent := range out.Data.Agents {
		if agent["ok"] != nil {
			fmt.Println(agent["ok"])
		} else {
			fmt.Println(agent["error"])
		}
	}

	// 5. Replication Agent Show configuration
	agent1 := aem.NewAgentPublish("toProdPublish1", aem.PolicyShow)
	c.RegisterAgent(agent1)

	out, err = c.Do("localhost", "4502", "admin", "12345")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, agent := range out.Data.Agents {
		for i, val := range agent {
			fmt.Printf("%v - %v\n", i, val)
		}
	}
}
