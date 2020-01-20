package sns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSessionWithProfile returns a new AWS session with the values provided
func NewSessionWithProfile(region, profile string) (*session.Session, error) {

	sess, err := session.NewSessionWithOptions(
		session.Options{
			Profile: profile,
			Config: aws.Config{
				Region: aws.String(region),
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// NewSessionWithCredentials returns a new AWS session with the values provided
func NewSessionWithCredentials(accessKeyID, secretAccessKey, region string) (*session.Session, error) {

	value := credentials.Value{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
	return sessionWithValue(value, region)
}

// NewSessionWitToken returns a new AWS session with the values provided
func NewSessionWitToken(id, secret, token, region string) (*session.Session, error) {

	value := credentials.Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    token,
	}
	return sessionWithValue(value, region)
}

func sessionWithValue(value credentials.Value, region string) (*session.Session, error) {
	creds := credentials.NewStaticCredentialsFromCreds(value)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}
