package api

import (
	"air/azuredevops"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Organization string `json:"organization"`
	Project      string `json:"project"`
	PipelineID   int    `json:"pipeline_id"`
}

func TriggerJobHandler(c *gin.Context) {
	var jobReq JobRequest

	if err := c.ShouldBindJSON(&jobReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, err := azuredevops.TriggerJob(jobReq.Username, jobReq.Password, jobReq.Organization, jobReq.Project, jobReq.PipelineID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Job triggered successfully"})
}
