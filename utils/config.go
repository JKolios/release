package utils

import (
	"encoding/json"
	"io/ioutil"
)

//Configuration : Application configuration Template
type Configuration struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
	BaseAsset string `json:"baseAsset"`
}

// ParseJSONFile Reads the app's configuration from a JSON file at filename
func ParseJSONFile(filename string) (Configuration, error) {

	var confObject Configuration
	confFile, err := ioutil.ReadFile(filename)

	err = json.Unmarshal(confFile, &confObject)
	return confObject, err
}
