package api

import (
	"fmt"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/go-resty/resty/v2"
	"log"
)

const CONVERTOR_URL_TEMPLATE = "%s/v1/public/convertor"

type FCApiClient struct {
	host string
}

func NewFCApiClient(host string) FCApiClient {
	return FCApiClient{host: host}
}

func (c FCApiClient) CallPublicConvertor(payload string) {
	// For now, convertor supports only JSON to schema conversion
	convertorPayload := map[string]string{
		"from": "json",
		"to":   "schema",
		"code": payload,
	}
	result := c.callPublicAPI(convertorPayload)
	fmt.Println(result)
}

func (c FCApiClient) callPublicAPI(payload any) any {
	client := resty.New()

	var response any
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&response).
		Post(fmt.Sprintf(CONVERTOR_URL_TEMPLATE, utils.GetFCHost()))
	if err != nil {
		log.Fatalf("Error making fucioncatalyst server HTTP request: %v", err)
	}

	if resp.IsError() {
		log.Fatalf("Error making fucioncatalyst server request. Server returned status: %d.\nError: %s",
			resp.StatusCode(), resp.String())
	}

	return response
}
