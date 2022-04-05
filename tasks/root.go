package tasks

import (
	"fmt"

	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

var (
	SMTP_EMAIL_FROM = &sendinblue.SendSmtpEmailReplyTo{
		Email: "support@trackflow.io",
		Name:  "TrackFlow",
	}
	SignUpConfirmationEmailTemplateId = int64(2)
)

type tasks struct{}

func Task() *tasks {
	return &tasks{}
}

func buildTaskQueuePath(projectID, locationID, queueID string) string {
	return fmt.Sprintf(
		"projects/%s/locations/%s/queues/%s",
		projectID,
		locationID,
		queueID,
	)
}
