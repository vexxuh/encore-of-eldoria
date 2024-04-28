package novel_ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type novelResponseBody struct {
	Output string `json:"output"`
}

func GenerateAiText(input string) (string, error) {
	client := http.DefaultClient

	requestBody := buildTextAiRequest(input)
	// Create an HTTP GET request
	req, err := http.NewRequest("POST", "https://api.novelai.net/ai/generate", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", errors.New("could not build request")
	}

	// Set any headers if required
	req.Header.Set("Authorization", "Bearer "+os.Getenv("NOVEL_AI_KEY"))
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("could not request the novel api for ai/textgen")
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("could read response body")
	}

	// Parse the JSON response body into an instance of MyResponse struct
	var novelResp novelResponseBody
	if err := json.Unmarshal(responseBody, &novelResp); err != nil {
		return "", errors.New("could not parse response body")
	}

	return novelResp.Output, nil
}
