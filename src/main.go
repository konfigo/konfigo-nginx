package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"time"
)

var previousData map[string]interface{}

func main() {
	apiEndpoint := os.Getenv("KONFIGO_API_ENDPOINT")
	apiKey := os.Getenv("KONFIGO_API_KEY")
	path := os.Getenv("KONFIGO_PATH")
	intervalSeconds, err := strconv.Atoi(os.Getenv("KONFIGO_INTERVAL"))

	// Copy contents of /usr/share/nginx/html to temp folder
	tempDir, err := os.MkdirTemp("", "konfigo-")
	if err != nil {
		fmt.Printf("Error creating temp dir: %s\n", err)
		os.Exit(1)
	}
	cmd := fmt.Sprintf("cp -r /usr/share/nginx/html/* %s/", tempDir)
	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Error copying files: %s\n%s\n", err, output)
		os.Exit(1)
	}
	fmt.Printf("Temporary files are located at: %s\n", tempDir)

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
			rmcmd := "rm -rf /usr/share/nginx/html/*"
			output, err1 := exec.Command("sh", "-c", rmcmd).CombinedOutput()
			if err1 != nil {
				fmt.Printf("Error deleting files: %s\n%s\n", err1, output)
			}

			cmd := fmt.Sprintf("cp -r %s/* /usr/share/nginx/html/", tempDir)
			output, err2 := exec.Command("sh", "-c", cmd).CombinedOutput()
			if err2 != nil {
				fmt.Printf("Error copying files: %s\n%s\n", err2, output)
			}

			for k, v := range data {
				search := fmt.Sprintf("__%s__", k)
				cmd := fmt.Sprintf("find /usr/share/nginx/html/ -type f -print0 | xargs -0 sed -i 's/%s/%v/g'", search, v)
				output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
				if err != nil {
					fmt.Printf("Error replacing key %s: %s\n%s\n", k, err, output)
				}
			}

			previousData = data
		}

		time.Sleep(interval)
	}
}
