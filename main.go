package main

import (
	"os"

	"github.com/shibme/demo-api/api"
)

func main() {
	apiToken := os.Getenv("DEMO_API_TOKEN")
	if apiToken == "" {
		panic("DEMO_API_TOKEN environment variable is not set")
	}
	api.Serve(apiToken, 8080)
}
