package api

import (
	"air/azuredevops"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobRequest struct {
	PAT          string                 `json:"pat"`
	Organization string                 `json:"organization"`
	Project      string                 `json:"project"`
	DefinitionID int                    `json:"definition_id"`
	Parameters   map[string]interface{} `json:"parameters"`
}

func TriggerJobHandler(c *gin.Context) {
	var jobReq JobRequest

	if err := c.ShouldBindJSON(&jobReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, err := azuredevops.TriggerJob(jobReq.PAT, jobReq.Organization, jobReq.Project, jobReq.DefinitionID, jobReq.Parameters)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Job triggered successfully"})
}
