package driver

import (
	"encoding/json"

	"github.com/rockisch/appigo/jsonutils"
	"github.com/rockisch/appigo/requester"
)

type Element struct {
	Driver *Driver
	ID     string
}

func (d *Driver) FindElement(elName string, elBy string) *Element {
	reqBody := map[string]string{
		"using": elBy,
		"value": elName,
	}

	appiumReq := &requester.AppiumRequest{
		"POST",
		reqBody,
		"/wd/hub/session/" + d.SessionID + "/element",
	}

	res := requester.DoAppiumRequest(appiumReq, d.driverClient, "")

	if res.StatusCode != 200 {
		statusCodeErrorHandler(res.StatusCode, 404,
			"driver: the driver was unable to find an element on the screen using the specified arguments")
		statusCodeErrorHandler(res.StatusCode, 400,
			"driver: an invalid argument was passed to the findElement function")
	}

	mapBody := jsonutils.JSONToMap(res.Body)
	value := map[string]string{}

	err := json.Unmarshal(*mapBody["value"], &value)
	if err != nil {
		panic(err)
	}

	return &Element{d, value["ELEMENT"]}
}

func (el *Element) Click() {
	appiumReq := &requester.AppiumRequest{
		"POST",
		nil,
		"/wd/hub/session" + el.Driver.SessionID + "element" + el.ID + "/click",
	}

	res := requester.DoAppiumRequest(appiumReq, el.Driver.driverClient, "")

	if res.StatusCode != 200 {
		panic("ERROR IN ELEMENT CLICK")
	}
}
