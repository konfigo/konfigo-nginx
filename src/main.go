package main

import (
	"io"
	"net/http"
	"fmt"
	"os"
	"strconv"
	"encoding/json"
	"reflect"
	"time"
)

var previousData map[string]interface{}

func main() {
	apiEndpoint := os.Getenv("KONFIGO_API_ENDPOINT")
	apiKey := os.Getenv("KONFIGO_API_KEY")
	path := os.Getenv("KONFIGO_PATH")
	intervalSeconds, err := strconv.Atoi(os.Getenv("KONFIGO_INTERVAL"))
	if err != nil {
		// Handle error, default to 10 seconds if parsing fails
		intervalSeconds = 10
	}
	interval := time.Duration(intervalSeconds) * time.Second

	for {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", apiEndpoint, path), nil)
		if err != nil {
			// Log error and continue
			fmt.Printf("Error creating request: %s\n", err)
			time.Sleep(interval)
			continue
		}
		
		req.Header.Set("authorization", fmt.Sprintf("%s", apiKey))
		client := &http.Client{}
		resp, err := client.Do(req)
		
		if err != nil {
			fmt.Printf("Error making request: %s\n", err)

			// Log error and continue
			time.Sleep(interval)
			continue
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			// Log error and continue
			time.Sleep(interval)
			continue
		}

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			// Log error and continue
			time.Sleep(interval)
			continue
		}

		if previousData == nil || !reflect.DeepEqual(data, previousData) {
			// Log the data to console
			fmt.Println("Data:", data)

			previousData = data
		}

		time.Sleep(interval)
	}
}
