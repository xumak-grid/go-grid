package sns

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	MessageStructureRaw = "raw"
)

// Notify represents a SNS notification
type Notify struct {
	Subject  string                 `json:"-"`
	Messages map[string]interface{} `json:"messages"`
	TopicArn string                 `json:"-"`
	Session  *session.Session       `json:"-"`
}

// AddNotification adds notifications message to the Notify Messages map
// this addition does not override the messages already saved
func (n *Notify) AddNotification(name string, message interface{}) error {
	if n.Messages == nil {
		n.Messages = map[string]interface{}{
			name: message,
		}
		return nil
	}

	// check if the key already exist
	if _, exist := n.Messages[name]; exist {
		return fmt.Errorf("the key %v already exist", name)
	}
	n.Messages[name] = message
	return nil
}

// Publish publishes the messages with the values provided
// the Messages are marshaled to get a JSON string value
func (n *Notify) Publish() error {

	err := validate(n)
	if err != nil {
		return err
	}

	msgEnconded, err := customMarshal(n.Messages)
	if err != nil {
		return err
	}
	message := string(msgEnconded)

	svc := sns.New(n.Session)
	topic := &sns.PublishInput{
		TopicArn:         aws.String(n.TopicArn),
		Message:          &message,
		MessageStructure: aws.String(MessageStructureRaw),
		Subject:          aws.String(n.Subject),
	}
	_, err = svc.Publish(topic)
	if err != nil {
		return err
	}
	return nil
}

// validates if the values of Notify are present
func validate(n *Notify) error {
	if n == nil {
		return errors.New("notify is nil")
	}
	if len(n.Messages) == 0 {
		return errors.New("no messages to publish")
	}
	if n.TopicArn == "" {
		return errors.New("topicArn is empty")
	}
	if n.Subject == "" {
		return errors.New("subject is empty")
	}
	if n.Session == nil {
		return errors.New("session is nil")
	}
	return nil
}

// customMarshal adds indentation and ignores escapeHTML to the encoded message
// ignore html scape allows to add urls in the message
func customMarshal(m interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(m)
	return buffer.Bytes(), err
}
