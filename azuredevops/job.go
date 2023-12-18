package azuredevops

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type PipelineRunRequest struct {
	Definition struct {
		ID int `json:"id"`
	} `json:"definition"`
	Parameters string `json:"parameters"`
}

func TriggerJob(username, password, organization, project string, pipelineId int, parameters string) (int, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/pipelines/%d/runs?api-version=6.0-preview.1", organization, project, pipelineId)

	// Create the request body
	reqBody := PipelineRunRequest{}
	reqBody.Definition.ID = pipelineId
	reqBody.Parameters = parameters
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	req.Header.Set("Authorization", "Basic "+auth)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("failed to trigger job")
	}

	return resp.StatusCode, nil
}
