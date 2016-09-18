package util

import (
	"crypto/tls"
	"http"
	"io/ioutil"
	"strings"

	"github.com/op/go-logging"
)

type HttpAction struct {
	Method      string `yaml:"method"`
	Url         string `yaml:"url"`
	Body        string `yaml:"body"`
	Template    string `yaml:"template"`
	Accept      string `yaml:"accept"`
	ContentType string `yaml:"contentType"`
	Title       string `yaml:"title"`
	StoreCookie string `yaml:"storeCookie"`
}

var logger = logging.MustGetLogger("goreka")

func buildHttpRequest(action HttpAction) *http.Request {

	var request *http.Request
	var err error
	if action.Body != "" {
		reader := strings.NewReader(action.Body)
		request, err = http.NewRequest(action.Method, action.Url, reader)
	} else if action.Template != "" {
		reader := strings.NewReader(action.Template)
		request, err = http.NewRequest(action.Method, action.Url, reader)
	} else {
		request, err = http.NewRequest(action.Method, action.Url, nil)
	}
	if err != nil {
		logger.Fatal(err)
	}

	// Add headers
	request.Header.Add("Accept", action.Accept)
	if action.ContentType != "" {
		request.Header.Add("Content-Type", action.ContentType)
	}
	return request
}

func DoHttp(action HttpAction) ([]byte, bool) {

	request := buildHttpRequest(httpAction)

	var DefaultTransport http.RoundTripper = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	response, err := DefaultTransport.RoundTrip(request)
	if err != nil {
		logger.Errorf("HTTP request failed: %s", err)
		return nil, false
	} else {
		defer response.Body.Close()
		if content, err := ioutil.ReadAll(response.Body); err == nil {
			return content, true
		} else {
			logger.Errorf("Error extracting response content: %s", err)
			return nil, false
		}
	}
	return nil, false
}
