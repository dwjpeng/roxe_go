package roxe_test

import "os"

func getAPIURL() string {
	apiURL := os.Getenv("ROXE_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://mainnet.roxecanada.com"
}
