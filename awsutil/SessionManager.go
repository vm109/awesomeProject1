package awsutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateNewSession(region, profile string) (sess *session.Session, err error) {
	if region == "" {
		sess, err = session.NewSession()
	} else if region != "" && profile != "" {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String(region),
			},
			Profile: profile,
			SharedConfigState: session.SharedConfigEnable,
		})
	} else if region != "" {
		sess, err = session.NewSession(
			&aws.Config{
				Region: aws.String(region),
			})
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to Create AWS Session %v", err))
		return nil, err
	}

	return session.Must(sess, err), nil
}
