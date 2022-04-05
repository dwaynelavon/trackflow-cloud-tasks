package handlers

import (
	"log"
	"net/http"
	"os"

	tfwTasks "github.com/dwaynelavon/weissach/trackflow-cloud-tasks/tasks"
	"github.com/gin-gonic/gin"
)

type sendEmailHeader struct {
	TaskName  string `header:"X-Appengine-Taskname"`
	QueueName string `header:"X-Appengine-Queuename"`
}

/**
 * SendEmailHandler responds to requests for sending email templates
 */
func SendEmailHandler(c *gin.Context) {
	if c.FullPath() != SendEmailTaskHandlerPath {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Resource Not Found"})
		return
	}

	var header sendEmailHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if header.TaskName == "" {
		// Use the presence of the X-Appengine-Taskname header to validate
		// the request comes from Cloud Tasks.
		log.Println(
			"Invalid Task: No X-Appengine-Taskname request header found",
		)
		c.JSON(http.StatusBadRequest, "Invalid headers")
		return
	}

	// Pull useful headers from Task request.
	signUpConfirmationQueueName := os.Getenv("GCLOUD_TASKS_SIGN_UP_CONFIRMATION_QUEUE")

	tasks := tfwTasks.Task()
	switch header.QueueName {
	case signUpConfirmationQueueName:
		var json tfwTasks.SignUpConfirmationEmailBody
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Sending Sign Up Confirmation Email to %v", json.To)
		emailResp, err := tasks.SendSignUpConfirmationEmail(json.To)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(), "msg": "Error sending sign up confirmation email",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": emailResp.MessageId})
	}
}
