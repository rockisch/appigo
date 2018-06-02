package appigo

import (
	"encoding/json"
	"fmt"
)

// mapToJSON takes a map[string]string and turns it into a *[]byte element in JSON format.
func mapToJSON(data map[string]string) *[]byte {
	capabilities := make(map[string]map[string]string)
	capabilities["desiredCapabilities"] = data

	jsonCaps, err := json.Marshal(capabilities)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON Sent:\t" + string(jsonCaps))

	return &jsonCaps
}

// jsonToMap takes a *[]byte and turns it into a map[string]string element.
func jsonToMap(body *[]byte) map[string]*json.RawMessage {
	fmt.Println("JSON Recieved:\t" + string(*body))
	var sessionJSON map[string]*json.RawMessage
	err := json.Unmarshal(*body, &sessionJSON)
	if err != nil {
		panic(err)
	}

	return sessionJSON
}
