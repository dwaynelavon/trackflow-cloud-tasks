package handlers

import (
	"net/http"
	"os"

	"github.com/dwaynelavon/weissach/trackflow-cloud-tasks/tasks"
	"github.com/gin-gonic/gin"
)

/**
 * CompleteSignUpHandler responds to requests for sign up completion
 */
func CompleteSignUpHandler(c *gin.Context) {
	if c.FullPath() != CompleteSignUpPath {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Resource Not Found"})
		return
	}

	var json tasks.SignUpConfirmationEmailBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectID := os.Getenv("GCLOUD_PROJECT")
	locationID := os.Getenv("GCLOUD_LOCATION")
	queueID := os.Getenv("GCLOUD_TASKS_SIGN_UP_CONFIRMATION_QUEUE")

	tasks := tasks.Task()
	_, err := tasks.CompleteSignUp(projectID, locationID, queueID, json.To)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully enqueued signup tasks"})
}
