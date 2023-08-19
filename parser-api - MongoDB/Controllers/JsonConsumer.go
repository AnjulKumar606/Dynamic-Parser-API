package Controllers

import (
	"appdirs/cns-parser/Models"
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
)

// GetConsumer ... Get all Consumer
func WebHook(c *gin.Context) {
	id := c.Params.ByName("id")
	var consumer Models.Consumer

	err := Models.GetConsumerByID(&consumer, id)
	if err != nil {
		fmt.Printf("Error ------>")
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {

			//parse the request data by gabs
			jsonData, err := gabs.ParseJSON([]byte(jsonData))
			if err != nil {
				panic(err)
			}

			//function to take any type of data
			pathing := func(path, field string) string {
				//storing the value
				var value string
				//pathing
				Path := jsonData.Path(path)
				if Path != nil {
					Data := Path.Data()
					switch age := Data.(type) {
					//cases for each data type
					//changing each data type to the suitable data type i.e. string
					case float64:
						value = fmt.Sprintf("%f", age)
					case string:
						value = age
					case bool:
						value = strconv.FormatBool(age)
					default:
						fmt.Println("Path value is not a recognized data type")
					}
				} else {
					fmt.Println("Path not found for field :-", field)
				}
				return value
			}

			//making new json
			finalparsedData := gabs.New()

			//appending the fields
			finalparsedData.SetP(consumer.ApplicationName, "application_name")
			finalparsedData.SetP(pathing(consumer.Title, "title"), "title")
			finalparsedData.SetP(pathing(consumer.Url, "url"), "url")
			finalparsedData.SetP(pathing(consumer.Body, "body"), "body")
			finalparsedData.SetP(pathing(consumer.Image, "image"), "image")
			finalparsedData.SetP(pathing(consumer.Extra, "extra"), "extra")

			//indent
			outputparsedData := finalparsedData.StringIndent("", "  ")

			var data Models.Consumer

			// Unmarshal the JSON string into the MyData struct
			err2 := json.Unmarshal([]byte(outputparsedData), &data)
			if err2 != nil {
				fmt.Println("Error:", err2)
				return
			}

			c.JSON(http.StatusOK, gin.H{"parsedData": data, "consumer": consumer})
		}
	}
}

// GetConsumer ... Get all Consumer
func PullHook(c *gin.Context) {
	id := c.Params.ByName("id")
	// id := "64d23004b54f8edd47f893d3"
	var consumer Models.Consumer

	//fetch api
	apiURL := "https://64d22d1ff8d60b17436195a9.mockapi.io/event"

	err := Models.GetConsumerByID(&consumer, id)
	if err != nil {
		fmt.Printf("Error ------>")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//sending GET request
		response, err := http.Get(apiURL)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		defer response.Body.Close()

		//fetching the data from the api
		jsonData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		} else {
			//put the JSON array into SLICE
			var data []map[string]interface{}
			err2 := json.Unmarshal([]byte(jsonData), &data)
			if err2 != nil {
				fmt.Println("Error:", err)
				return
			}

			//data struct for output
			var parsedData []json.RawMessage

			//iterate through every slice element
			for _, item := range data {
				// Convert the map to JSON data
				jsonBytes, err := json.Marshal(item)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				// Parse the JSON data using gabs
				jsonData, err := gabs.ParseJSON([]byte(jsonBytes))
				if err != nil {
					panic(err)
				} else {

					//function to take any type of data
					pathing := func(path, field string) string {
						//storing the value
						var value string
						//pathing
						Path := jsonData.Path(path)
						if Path != nil {
							Data := Path.Data()
							switch age := Data.(type) {
							//cases for each data type
							//changing each data type to the suitable data type i.e. string
							case float64:
								value = fmt.Sprintf("%f", age)
							case string:
								value = age
							case bool:
								value = strconv.FormatBool(age)
							default:
								fmt.Println("Path value is not a recognized data type")
							}
						} else {
							fmt.Println("Path not found for field :-", field)
						}
						return value
					}

					//making new json
					finalparsedData := gabs.New()

					//appending the fields
					finalparsedData.SetP(consumer.ApplicationName, "application_name")
					finalparsedData.SetP(pathing(consumer.Title, "title"), "title")
					finalparsedData.SetP(pathing(consumer.Url, "url"), "url")
					finalparsedData.SetP(pathing(consumer.Body, "body"), "body")
					finalparsedData.SetP(pathing(consumer.Image, "image"), "image")
					finalparsedData.SetP(pathing(consumer.Extra, "extra"), "extra")

					outputparsedData := finalparsedData.StringIndent("", "  ")

					var data Models.Consumer

					// Unmarshal the JSON string into the MyData struct
					err := json.Unmarshal([]byte(outputparsedData), &data)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					// Marshal the JSON objects to byte slices
					jsonBytes, err := json.Marshal(data)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					//append the parsed data into the final data for output
					parsedData = append(parsedData, jsonBytes)
				}
			}
			//output
			c.JSON(http.StatusOK, gin.H{"parsedData": parsedData, "consumer": consumer})
		}
	}
}
