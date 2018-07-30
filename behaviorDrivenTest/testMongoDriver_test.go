package behaviorDrivenTest

import (
	"github.com/DATA-DOG/godog"
	"net/http/httptest"
	"net/http"
	"fmt"
)

type apiFeature struct{
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(interface{}){
	a.resp = httptest.NewRecorder()
}

var responseBody []byte

func iSendRequestUsingThe(requestMethod, uri string) (err error) {
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

	}
	return
}

func iShouldBeGettingAsExpected(arg1 int) error {
	return godog.ErrPending
}

func aJSONResponseWithAnd(arg1, arg2 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request using the "([^"]*)"$`, iSendRequestUsingThe)
	s.Step(`^I should be getting (\d+) as expected$`, iShouldBeGettingAsExpected)
	s.Step(`^a JSON response with "([^"]*)" and "([^"]*)"$`, aJSONResponseWithAnd)
}
