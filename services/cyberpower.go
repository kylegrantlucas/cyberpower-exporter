package cyberpower

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/kylegrantlucas/cyberpower-exporter/models"
)

type Client struct {
	Host string
	Port string
}

func (c Client) Poll() (models.CyberpowerResponse, error) {
	cbpResp := models.CyberpowerResponse{}

	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", c.Host, c.Port),
		Path:   "agent/ppbe.js/init_status.js",
	}

	client := http.Client{}
	request, err := http.NewRequest("GET", url.String(), bytes.NewReader([]byte{}))
	if err != nil {
		return cbpResp, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return cbpResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cbpResp, err
	}

	fistParsedBody := strings.Trim(string(body), "var ppbeJsObj=")
	parsedBody := strings.Replace(fistParsedBody, ";", "", -1)

	err = json.Unmarshal([]byte(parsedBody), &cbpResp)
	if err != nil {
		return cbpResp, err
	}

	return cbpResp, err
}
