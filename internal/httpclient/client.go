package httpclient

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	OER_URL = `https://openexchangerates.org/api/latest.json?app_id=%s&base=USD&prettyprint=true&show_alternative=false`
)

const (
	ENV_OER_APP_ID = `OER_APP_ID`
)

var (
	onceHTTPCLient sync.Once
	client         *http.Client
)

func Client() *http.Client {

	onceHTTPCLient.Do(func() {

		client = &http.Client{
			Timeout: 5 * time.Second,
		}

	})

	return client
}

func OERReq() (*http.Request, error) {
	apiKey := os.Getenv(ENV_OER_APP_ID)
	if apiKey == "" {
		return nil, fmt.Errorf("Environment variable \"%s\" it not set", ENV_OER_APP_ID)
	}
	oerURL := fmt.Sprintf(OER_URL, apiKey)
	return http.NewRequest(http.MethodGet, oerURL, nil)
}
