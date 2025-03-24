package internal

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/viper"
)

type APIRequest struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func SaveRequest(name, url, method string, headers map[string]string, body string) {
	viper.Set(name, APIRequest{URL: url, Method: method, Headers: headers, Body: body})
	viper.WriteConfig()
}

func LoadRequest(name string) (APIRequest, bool) {
	var req APIRequest
	if viper.Get(name) == nil {
		return req, false
	}
	viper.UnmarshalKey(name, &req)
	return req, true
}

func DeleteRequest(name string) bool {
	if viper.Get(name) == nil {
		return false
	}
	viper.Set(name, nil)
	viper.WriteConfig()
	return true
}

// ListRequests prints all saved API requests
func ListRequests() {
	requests := make(map[string]APIRequest)
	viper.Unmarshal(&requests)

	if len(requests) == 0 {
		fmt.Println("📂 No saved requests found!")
		return
	}

	fmt.Println("\n📌 Saved API Requests:")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintln(w, "Name\tMethod\tURL")

	for name, req := range requests {
		fmt.Fprintf(w, "%s\t%s\t%s\n", name, req.Method, req.URL)
	}

	w.Flush()
}
