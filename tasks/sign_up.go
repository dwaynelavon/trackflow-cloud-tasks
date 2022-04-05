package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"github.com/dwaynelavon/weissach/trackflow-cloud-tasks/common"
	"github.com/dwaynelavon/weissach/trackflow-cloud-tasks/email"
)

type SignUpConfirmationEmailBody struct {
	To string
}

func createSignUpEmailConfirmationTask(
	queuePath string,
	data []byte,
) *taskspb.CreateTaskRequest {
	return &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &taskspb.AppEngineHttpRequest{
					HttpMethod:  taskspb.HttpMethod_POST,
					RelativeUri: os.Getenv("SEND_EMAIL_TAKS_HANDLER_URI"),
					Body:        []byte(data),
				},
			},
		},
	}
}

/**
 * CompleteSignUp creates a new task to schedule Sign Up
 * completion tasks.
 */
func (tasks *tasks) CompleteSignUp(
	projectID,
	locationID,
	queueID,
	to string,
) (*taskspb.Task, error) {
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	queuePath := buildTaskQueuePath(projectID, locationID, queueID)
	message := &SignUpConfirmationEmailBody{
		To: to,
	}
	data, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %v", err)
	}

	emailReq := createSignUpEmailConfirmationTask(queuePath, data)
	// Cloud tasks is not supported in local development.
	if !common.IsProduction() {
		return nil, nil
	}

	emailTask, err := client.CreateTask(ctx, emailReq)
	if err != nil {
		return nil, fmt.Errorf(
			"cloudtasks.CreateTask createSignUpEmailConfirmationTask: %v",
			err,
		)
	}

	return emailTask, nil
}

func (tasks *tasks) SendSignUpConfirmationEmail(to string) (
	sendinblue.CreateSmtpEmail,
	error,
) {
	email := email.Email()
	emailService := email.Client.TransactionalEmailsApi
	emailResponse, _, err := emailService.SendTransacEmail(
		context.Background(),
		sendinblue.SendSmtpEmail{
			To: []sendinblue.SendSmtpEmailTo{
				{Email: to},
			},
			ReplyTo:    SMTP_EMAIL_FROM,
			TemplateId: SignUpConfirmationEmailTemplateId,
		},
	)

	return emailResponse, err
}
