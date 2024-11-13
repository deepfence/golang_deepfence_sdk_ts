package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"time"

	openapi "github.com/deepfence/golang_deepfence_sdk/client"
	rhttp "github.com/hashicorp/go-retryablehttp"
)

func IsConsoleAgent(url string) bool {
	return len(os.Getenv("DEEPFENCE_CONSOLE_AGENT")) > 0
}

func GetConsoleApiToken(console, port string) (string, error) {
	// setup http client
	rhc := rhttp.NewClient()
	rhc.HTTPClient.Timeout = 10 * time.Second
	rhc.RetryMax = 3
	rhc.RetryWaitMin = 1 * time.Second
	rhc.RetryWaitMax = 10 * time.Second

	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		RootCAs:            x509.NewCertPool(),
		InsecureSkipVerify: true,
	}
	tr.DisableKeepAlives = false
	tr.Proxy = http.ProxyFromEnvironment

	rhc.HTTPClient = &http.Client{Transport: tr}

	servers := openapi.ServerConfigurations{
		{
			URL:         fmt.Sprintf("http://%s:%s", console, port),
			Description: "deepfence_server_internal",
		},
	}

	cfg := openapi.NewConfiguration()
	cfg.HTTPClient = rhc.StandardClient()
	cfg.Servers = servers
	client := openapi.NewAPIClient(cfg)

	token, resp, err := client.InternalAPI.GetConsoleApiTokenExecute(openapi.ApiGetConsoleApiTokenRequest{})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(resp.Status)
	}

	return token.GetApiToken(), nil
}
