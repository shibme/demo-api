package main

import (
	"os"

	"github.com/shibme/demo-api/api"
)

func getEnv(envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		panic(envName + " environment variable is not set")
	}
	return env
}

func main() {
	api.Serve(getEnv("DEMO_API_TOKEN"), getEnv("DEMO_API_MESSAGE"), 8888)
}
