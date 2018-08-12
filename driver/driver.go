package driver

import (
	"encoding/json"

	"github.com/rockisch/appigo/jsonutils"

	"github.com/rockisch/appigo/client"
	"github.com/rockisch/appigo/requester"
)

// Driver object containing some data related to your session
type Driver struct {
	driverClient       *client.Client
	driverCapabilities map[string]string
	SessionID          string
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

// Init tries to start a appium session with the url and capabilities stored in the driver.
func (d *Driver) Init() {
	appiumReq := &requester.AppiumRequest{
		"POST",
		d.driverCapabilities,
		"/wd/hub/session",
	}

	resp := requester.DoAppiumRequest(appiumReq, d.driverClient, "desiredCapabilities")

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"appigo: unable to create session. please, check if the specified capabilities are corret",
	)

	mapBody := jsonutils.JSONToMap(resp.Body)

	err := json.Unmarshal(*mapBody["sessionId"], &d.SessionID)
	if err != nil {
		panic(err)
	}
}

// Close closes the session stored in the driver. It's always good practice to defer "driver.Close()"
// as soon as you cal "driver.Init()"
func (d *Driver) Close() {
	appiumReq := &requester.AppiumRequest{
		"DELETE",
		nil,
		"/wd/hub/session/" + d.SessionID,
	}

	resp := requester.DoAppiumRequest(appiumReq, d.driverClient, "")

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"appigo: unable to close session",
	)
}
