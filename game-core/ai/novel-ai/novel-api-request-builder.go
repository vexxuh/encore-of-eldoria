package novel_ai

import "encoding/json"

const text_request_template = `{
"input": "INPUT",
"model": "euterpe-v2",
"parameters": {
  "use_string": true,
  "temperature": 1,
  "min_length": 10,
  "max_length": 30
}
}`

type aiTextRequest struct {
	Input      string              `json:"input"`
	Model      string              `json:"model"`
	Parameters aiRequestParameters `json:"parameters"`
}

type aiRequestParameters struct {
	UseString   bool `json:"use_string"`
	Temperature uint `json:"temperature"`
	MinLength   uint `json:"min_length"`
	MaxLength   uint `json:"max_length"`
}

func buildTextAiRequest(input string) []byte {
	var req = aiTextRequest{
		Input: input,
		Model: "euterpe-v2",
		Parameters: aiRequestParameters{
			UseString:   true,
			Temperature: 1,
			MinLength:   10,
			MaxLength:   30,
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		panic("System is not working correctly when processing text request a review of data should be seen too")
	}

	return reqBody
}
