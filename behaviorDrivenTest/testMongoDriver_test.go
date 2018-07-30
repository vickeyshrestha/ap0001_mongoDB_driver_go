package behaviorDrivenTest

import (
	"github.com/DATA-DOG/godog"
	"net/http/httptest"
	"net/http"
	"fmt"
	"ap0001_mongo_engine/internal/healthCheck"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type apiFeature struct{
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(interface{}){
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendRequestUsingThe(requestMethod, uri string) (err error) {
	req, err := http.NewRequest(requestMethod, uri, nil)
	if err != nil {
		return
	}
	defer func() {
		switch t:= recover().(type){
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch uri{
	case "/health":
		healthCheck.HealthCheckHandler(a.resp, req)
	default:
		err = fmt.Errorf("Bad URI", uri)
	}
	return
}

func (a *apiFeature) iShouldBeGettingAsExpected(httpCode int) error {
	if httpCode != a.resp.Code{
		return fmt.Errorf("The expected http code is wrong. Expected %d, Returned %d", httpCode, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) aJSONResponseWithAnd(applicationName, healthStatus string) error {
	var jsonResponseForHealth healthCheck.HealthEndpoint
	body, _ := ioutil.ReadAll(a.resp.Body)
	json.Unmarshal(body, &jsonResponseForHealth)
	if !strings.EqualFold(string(jsonResponseForHealth.Application), applicationName){
		return fmt.Errorf("Expected application name: %s, Actual: %s", applicationName, string(jsonResponseForHealth.Application))
	}
	if !strings.EqualFold(string(jsonResponseForHealth.HealthStatus), healthStatus){
		return fmt.Errorf("Expected health status of application name: %s, Actual: %s",healthStatus, string(jsonResponseForHealth.HealthStatus))
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}
	s.BeforeScenario(api.resetResponse)
	s.Step(`^I send "(GET|POST|PUT|DELETE)" request using the "([^"]*)"$`, api.iSendRequestUsingThe)
	s.Step(`^I should be getting (\d+) as expected$`, api.iShouldBeGettingAsExpected)
	s.Step(`^a JSON response with "([^"]*)" and "([^"]*)"$`,api.aJSONResponseWithAnd)
}
