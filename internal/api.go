package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/manifoldco/promptui"
)

// GetHeadersFromUser prompts the user to enter headers interactively
func GetHeadersFromUser() map[string]string {
	headers := make(map[string]string)

	addHeadersPrompt := promptui.Prompt{
		Label:     "Add Headers? (y/n)",
		IsConfirm: true,
	}
	addHeaders, _ := addHeadersPrompt.Run()
	if strings.ToLower(addHeaders) == "y" {
		for {
			headerKeyPrompt := promptui.Prompt{Label: "Header Key (or press Enter to skip)"}
			headerKey, _ := headerKeyPrompt.Run()
			if headerKey == "" {
				break
			}
			headerValuePrompt := promptui.Prompt{Label: fmt.Sprintf("Value for %s", headerKey)}
			headerValue, _ := headerValuePrompt.Run()
			headers[headerKey] = headerValue
		}
	}
	return headers
}

// ParseHeaders converts a JSON string into a `map[string]string`
func ParseHeaders(headersStr string) map[string]string {
	headers := make(map[string]string)

	if headersStr != "" {
		err := json.Unmarshal([]byte(headersStr), &headers)
		if err != nil {
			log.Fatalf("❌ Error parsing headers: %v", err)
		}
	}

	return headers
}

// GetBodyFromUser prompts the user to enter a request body if required
func GetBodyFromUser(method string) string {
	if method == "POST" || method == "PUT" {
		bodyPrompt := promptui.Prompt{Label: "Enter JSON Body (or press Enter to skip)"}
		body, _ := bodyPrompt.Run()
		return body
	}
	return ""
}

// SendRequest makes an HTTP request and prints the response
func SendRequest(url, method string, headers map[string]string, body []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body)) // Body is []byte
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print response
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("\n📩 Response:")
	fmt.Println(string(bodyBytes))
}
