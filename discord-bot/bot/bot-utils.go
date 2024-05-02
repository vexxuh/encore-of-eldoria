package discordbot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type BackendProcessorParams struct {
	Msg string `json:"msg"`
}

// backendGameProcessorRequest sends request to the backend game server then recieves the response of data!
func backendGameProcessorRequest(apiEndpoint, userData string) (*BackendResponse, error) {
	requestData, err := backendRequest(apiEndpoint, &BackendProcessorParams{Msg: userData})
	if err != nil {
		return nil, err
	}

	return requestData, nil
}

func backendRequest(endpoint string, backendParams *BackendProcessorParams) (*BackendResponse, error) {
	url := "http://127.0.0.1:4000/" + endpoint

	// Convert request data to JSON
	requestDataJSON, err := json.Marshal(backendParams)
	if err != nil {
		return nil, err
	}

	// Send POST request to the server
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestDataJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response JSON into ResponseStruct
	var responseData BackendResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, err
	}

	return &BackendResponse{}, nil
}
