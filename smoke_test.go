package http_smoke_tests_test

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTTP Smoke tests", func() {
	var (
		responeCode       string
		responseBodyRegex string
		resp              *http.Response
	)

	const defaultRequestMethod = "GET"

	BeforeEach(func() {
		responeCode = loadVar("RESPONSE_CODE")
		responseBodyRegex = loadVar("RESPONSE_BODY_REGEX")
		headersJSON := os.Getenv("HEADERS")
		if headersJSON == "" {
			headersJSON = "{}"
		}

		requestMethod := os.Getenv("REQUEST_METHOD")
		if requestMethod == "" {
			requestMethod = defaultRequestMethod
		}

		url := loadVar("URL")

		var headers map[string]string
		if err := json.Unmarshal([]byte(headersJSON), &headers); err != nil {
			defer GinkgoRecover()
			Fail(fmt.Sprintf(
				"Loading headers returned a JSON unmarshal error: %s\n provided JSON was: %s",
				err.Error(), headersJSON),
			)
		}

		req, err := http.NewRequest(requestMethod, url, nil)
		Ω(err).Should(BeNil())
		for key, value := range headers {
			if strings.ToLower(key) == "host" {
				req.Host = value
			}
			req.Header.Set(key, value)
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipSSLVerification()},
		}
		client := &http.Client{Transport: tr}
		resp, err = client.Do(req)
		if err != nil {
			defer GinkgoRecover()
			Fail(fmt.Sprintf("HTTP Request failed: %s", err.Error()))
		}
	})

	It("returns the desired response code and body", func() {
		Ω(strconv.Itoa(resp.StatusCode)).Should(Equal(responeCode))
		Ω(readRespBody(resp)).Should(MatchRegexp(responseBodyRegex))
	})
})

func loadVar(variableName string) string {
	variable := os.Getenv(variableName)
	if variable == "" {
		defer GinkgoRecover()
		Fail(fmt.Sprintf("Variable %s is required but not set", variableName))
	}
	return variable
}

func skipSSLVerification() bool {
	skipSSL := os.Getenv("SKIP_SSL_VERIFICATION")
	if skipSSL == "true" {
		return true
	}
	return false
}

func readRespBody(resp *http.Response) string {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		defer GinkgoRecover()
		Fail(fmt.Sprintf("Reading response body errored: %s", err.Error()))
	}
	defer resp.Body.Close()
	return string(data)
}
