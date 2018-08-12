package driver

import "github.com/rockisch/appigo/requester"

func (d *Driver) ImplicitWait(seconds int) {
	reqBody := map[string]string{
		"ms": string(seconds * 1000),
	}

	appiumReq := requester.AppiumRequest{
		"POST",
		reqBody,
		"/wd/hub/session/" + d.SessionID + "/timeouts/implicit_wait",
	}
	resp := requester.DoAppiumRequest(&appiumReq, d.driverClient, "")

	if resp.StatusCode != 200 {
		panic("IMPLICIT WAIT ERROR")
	}
}
