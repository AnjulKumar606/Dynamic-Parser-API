package Utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"gopkg.in/yaml.v2"
)

func ParseJSONWithYAMLKeys(jsonData []byte, yamlConfig []byte) map[string]interface{} {
	var config struct {
		Keys []string `yaml:"keys"`
	}

	if err := yaml.Unmarshal(yamlConfig, &config); err != nil {
		log.Fatalf("failed to parse YAML: %v", err)
	}

	var parsedJSON map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsedJSON); err != nil {
		log.Fatalf("failed to parse JSON: %v", err)
	}

	result := make(map[string]interface{})
	for _, key := range config.Keys {
		if val, ok := parsedJSON[key]; ok {
			result[key] = val
		}
	}

	return result
}

func ParseJSONByKey(jsonData []byte, format string) string {
	jsonParsed, err := gabs.ParseJSON([]byte(jsonData))

	if err != nil {
		panic(err)
	}

	var value string
	var ok bool
	var finalValue string


	//comment
	res1 := strings.Split(format, ",")
	fmt.Print("String value")
	fmt.Print(res1)
	for index, itr := range res1 {
		fmt.Print(index, " : ", itr, "\n")
		finalValue = strings.TrimSpace(itr)
	}


	fmt.Print("finalValue : ", finalValue, "\n")
	fmt.Print("jsonParsed : ", jsonParsed, "\n")
	fmt.Print("jsonData : ", jsonData, "\n")

	value, ok = jsonParsed.Path(finalValue).Data().(string)

	
	// value == 10.0, ok == true


	fmt.Print("value : ", value, "\n")
	fmt.Print("Test value", ok)
	return value
}
