package sns_test

import (
	"fmt"

	"github.com/xumak-grid/go-grid/pkg/notifications/sns"
)

type Message struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
}

func Example() {
	session, err := sns.NewSessionWithCredentials("XXXX", "XXXX", "us-east-1")
	if err != nil {
		fmt.Println("session error:", err)
		return
	}

	notify := sns.Notify{
		Subject:  "Bedrock deployment",
		TopicArn: "arn:aws:sns:us-east-1:281327226678:toolbelt_notification",
		Session:  session,
	}

	type Message struct {
		Message string `json:"message"`
	}

	t := Message{"Toolbelt deployment success "}
	notify.AddNotification("toolbelt", t)

	k := Message{"k8s deplotment success"}
	notify.AddNotification("k8s", k)

	// publish all the messages
	err = notify.Publish()
	if err != nil {
		fmt.Println(err.Error())
	}
}
