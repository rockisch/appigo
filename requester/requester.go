package requester

import (
	"github.com/rockisch/appigo/client"
	"github.com/rockisch/appigo/jsonutils"
)

type AppiumRequest struct {
	Method  string
	BodyMap map[string]string
	Path    string
}

func DoAppiumRequest(appiumReq *AppiumRequest, c *client.Client, name string) *client.Response {
	resp, err := c.MakeRequest(
		appiumReq.Method,
		jsonutils.StringMapToJSON(appiumReq.BodyMap, name),
		appiumReq.Path,
	)

	if err != nil {
		panic(err)
	}

	return &resp
}
