package behaviorDrivenTest

import (
	"ap0001_mongo_engine/internal/healthCheck"
	"ap0001_mongo_engine/mocks"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendRequestUsingThe(requestMethod, uri string) (err error) {
	req, err := http.NewRequest(requestMethod, uri, nil)
	if err != nil {
		return
	}
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch uri {
	case "/health":
		mockInitialConfig := &mocks.InitialConfig{}
		mockInitialConfig.On("GetApplicationBinary").Return("0.0.1")
		mockInitialConfig.On("GetAppStartupTime").Return(time.Now())
		service, _ := healthCheck.NewHealthService(mockInitialConfig)
		service.HealthCheck(a.resp, req)
	default:
		err = fmt.Errorf("bad uri: %s", uri)
	}
	return
}

/*func (a *apiFeature) iSendRequestUsingThe(requestMethod, uri string) (err error) {
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
}*/

func (a *apiFeature) iShouldBeGettingAsExpected(httpCode int) error {
	if httpCode != a.resp.Code {
		return fmt.Errorf("expected http code is wrong. Expected %d, Returned %d", httpCode, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) aJSONResponseWithAnd(applicationName, healthStatus string) error {
	var jsonResponseForHealth healthCheck.HealthEndpoint
	body, _ := ioutil.ReadAll(a.resp.Body)
	_ = json.Unmarshal(body, &jsonResponseForHealth)
	if !strings.EqualFold(string(jsonResponseForHealth.Application), applicationName) {
		return fmt.Errorf("expected application name: %s, Actual: %s", applicationName, string(jsonResponseForHealth.Application))
	}
	if !strings.EqualFold(string(jsonResponseForHealth.HealthStatus), healthStatus) {
		return fmt.Errorf("expected health status of application name: %s, Actual: %s", healthStatus, string(jsonResponseForHealth.HealthStatus))
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}
	s.BeforeScenario(api.resetResponse)
	s.Step(`^I send "(GET|POST|PUT|DELETE)" request using the "([^"]*)"$`, api.iSendRequestUsingThe)
	s.Step(`^I should be getting (\d+) as expected$`, api.iShouldBeGettingAsExpected)
	s.Step(`^a JSON response with "([^"]*)" and "([^"]*)"$`, api.aJSONResponseWithAnd)
}
