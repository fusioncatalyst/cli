package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fusioncatalyst/cli/contracts"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

const CONVERTOR_URL_TEMPLATE = "%s/v1/public/convertor"
const PROJECTS_URL_TEMPLATE = "%s/v1/protected/projects"
const SCHENAS_URL_TEMPLATE = "/v1/protected/projects/%s/schemas"

type FCApiClient struct {
	host string
}

func NewFCApiClient(host string) FCApiClient {
	return FCApiClient{host: host}
}

func (c FCApiClient) getAPIKeyIfItExists() string {
	key, exists := utils.GetFCAPIKey()
	if !exists {
		fmt.Println("FusionCatalyst API key not found. Please set the environment variable %s",
			utils.FUSIONCATALYST_API_KEY_ENV_VAR)
		os.Exit(1)
	}
	return key
}

func (c FCApiClient) CallPublicConvertor(payload string) contracts.PublicUtilConvertorResponse {
	// For now, convertor supports only JSON to schema conversion
	convertorPayload := map[string]string{
		"from": "json",
		"to":   "schema",
		"code": payload,
	}
	response := c.callPublicAPIPost(CONVERTOR_URL_TEMPLATE, convertorPayload)

	intermediateJSON, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error marshaling map to JSON: %v", err)
	}

	var specificData contracts.PublicUtilConvertorResponse
	err = json.Unmarshal(intermediateJSON, &specificData)
	if err != nil {
		log.Fatalf("Error unmarshaling into specific struct: %v", err)
	}

	return specificData
}

func (c FCApiClient) CallPrivateNewProject(projectName string) (*contracts.PrivateProjectsResponse, error) {
	projectPayload := map[string]string{
		"name": projectName,
	}
	response, err := c.callPrivateAPIPost(PROJECTS_URL_TEMPLATE, projectPayload)
	if err != nil {
		return nil, err
	}

	intermediateJSON, err := json.Marshal(response)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error marshaling map to JSON: %v", err))
	}

	var specificData contracts.PrivateProjectsResponse
	err = json.Unmarshal(intermediateJSON, &specificData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshaling into specific struct: %v", err))
	}

	return &specificData, nil
}

func (c FCApiClient) CallPrivateListProjects() (*[]contracts.PrivateProjectsResponse, error) {
	response, e := c.callPrivateAPIGet(PROJECTS_URL_TEMPLATE)
	if e != nil {
		return nil, e
	}

	intermediateJSON, err := json.Marshal(response)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error marshaling map to JSON: %v", err))
	}

	var specificData []contracts.PrivateProjectsResponse
	err = json.Unmarshal(intermediateJSON, &specificData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshaling into specific struct: %v", err))
	}

	return &specificData, nil
}

func (c FCApiClient) CallPrivateNewJSONSchema(schema string, schemaName string, projectID string) (*contracts.PrivateSchemaResponse, error) {
	payload := fmt.Sprintf(`{
		"name": "%s",
		"type": "jsonschema",
		"schema": %s
	}`, schemaName, utils.StringifyJSON(schema))

	urlTemplate := fmt.Sprintf(SCHENAS_URL_TEMPLATE, projectID)

	response, e := c.callPrivateAPIPost("%s"+urlTemplate, payload)
	if e != nil {
		return nil, e
	}

	intermediateJSON, err := json.Marshal(response)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error marshaling map to JSON: %v", err))
	}

	var schemaResponse contracts.PrivateSchemaResponse
	err = json.Unmarshal(intermediateJSON, &schemaResponse)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshaling into specific struct: %v", err))
	}

	return &schemaResponse, nil
}

func (c FCApiClient) callPublicAPIPost(url string, payload any) any {
	client := resty.New()

	var response any
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&response).
		Post(fmt.Sprintf(url, utils.GetFCHost()))
	if err != nil {
		fmt.Println("Error making fucioncatalyst server HTTP request: %v", err)
		os.Exit(1)
	}

	if resp.IsError() {
		fmt.Println("Error making fucioncatalyst server request. Server returned status: %d.\nError: %s",
			resp.StatusCode(), resp.String())
		os.Exit(1)
	}

	return response
}

func (c FCApiClient) callPrivateAPIGet(url string) (any, error) {
	client := resty.New()

	var response any
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&response).
		SetAuthToken(c.getAPIKeyIfItExists()).
		Get(fmt.Sprintf(url, utils.GetFCHost()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making fucioncatalyst server HTTP request: %v", err))
	}

	if resp.IsError() {
		return nil, errors.New(fmt.Sprintf("Error making fucioncatalyst server request. Server returned status: %d.\n Error: %s",
			resp.StatusCode(), resp.String()))
	}

	return response, nil
}

func (c FCApiClient) callPrivateAPIPost(url string, payload any) (any, error) {
	client := resty.New()

	var response any
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetAuthToken(c.getAPIKeyIfItExists()).
		SetResult(&response).
		Post(fmt.Sprintf(url, utils.GetFCHost()))
	if err != nil {
		e := errors.New(fmt.Sprintf("Error making fucioncatalyst server HTTP request: %v", err))
		return nil, e
	}

	if resp.IsError() {
		e := errors.New(fmt.Sprintf("Error making fucioncatalyst server request. Server returned status: %d.\n Error: %s",
			resp.StatusCode(), resp.String()))
		return nil, e

	}

	return response, nil
}
