package driver

import (
	"encoding/json"
	"fmt"

	"github.com/rockisch/appigo/client"
)

// Driver object containing some data related to your session
type Driver struct {
	driverClient       *client.Client
	driverCapabilities map[string]string
	sessionID          string
}

// appiumRequest stores data to do requests to your appium server with "doAppiumRequest"
type appiumRequest struct {
	method  string
	bodyMap map[string]string
	path    string
}

// CreateDriver takes the create a Driver with the specified URL and and a map of
// capabilities, returning the driver after that.
func CreateDriver(url string, capabilities map[string]string) *Driver {
	newDriver := &Driver{
		client.CreateClient(url),
		capabilities,
		"",
	}

	return newDriver
}

// doAppiumRequest does a request to the appium server with the specified driver and data (appiumRequest).
func doAppiumRequest(d *Driver, appiumReq *appiumRequest) client.Response {
	resp, err := d.driverClient.MakeRequest(
		appiumReq.method,
		*mapToJSON(appiumReq.bodyMap),
		appiumReq.path,
	)

	if err != nil {
		panic(err)
	}

	return resp
}

// manageErrorStatusCode is used to handle error codes
func statusCodeErrorHandler(respStatusCode int, errStatusCode int, errString string) {
	if respStatusCode == errStatusCode {
		var err error
		if errString != "" {
			err = fmt.Errorf(errString)
		} else {
			err = fmt.Errorf("appigo: unexpected error occured, recieved status code %d", respStatusCode)
		}
		panic(err)
	}
}

// Init tries to start a appium session with the url and capabilities stored in the driver.
func (d *Driver) Init() {
	appiumReq := &appiumRequest{
		"POST",
		d.driverCapabilities,
		"/wd/hub/session",
	}

	resp := doAppiumRequest(d, appiumReq)

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"appigo: unable to create session. please, check if the specified capabilities are corret",
	)

	mapBody := jsonToMap(resp.Body)

	err := json.Unmarshal(*mapBody["sessionId"], &d.sessionID)
	if err != nil {
		panic(err)
	}
}

// Close closes the session stored in the driver. It's always good practice to defer "driver.Close()"
// as soon as you cal "driver.Init()"
func (d *Driver) Close() {
	appiumReq := &appiumRequest{
		"DELETE",
		nil,
		"/wd/hub/session/" + d.sessionID,
	}

	resp := doAppiumRequest(d, appiumReq)

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"appigo: unable to close session",
	)
}
