package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/harleywinston/x-manager/internal/master/consts"
)

type httpTestType struct {
	method   string
	url      string
	body     interface{}
	response *consts.CustomError
}

func runHTTPSubtests(t *testing.T, test httpTestType) {
	HTTPClient := &http.Client{}
	t.Run(test.response.Message, func(t *testing.T) {
		body, err := json.Marshal(test.body)
		if err != nil {
			t.Error(err.Error())
		}

		req, err := http.NewRequest(
			test.method,
			test.url,
			bytes.NewBuffer(body),
		)
		if err != nil {
			t.Error(err.Error())
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := HTTPClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}

		if resp.StatusCode != test.response.Code {
			t.Errorf(
				"code didn't match recieved: %d, expected: %d",
				resp.StatusCode,
				test.response.Code,
			)
		}

		type respDataType struct {
			Message string `json:"message"`
			Detail  string `json:"detail"`
		}

		var respData respDataType
		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			t.Error(err.Error())
		}

		if respData.Message != test.response.Message {
			t.Errorf(
				"message didn't match: %s, %s\n%s",
				respData.Message,
				test.response.Message,
				respData.Detail,
			)
		}
	})
}

func TestTransport(t *testing.T) {
	resourceTests := &ResourceTests{t}
	t.Run("Test ADD resource", resourceTests.TestAdd)
	t.Run("Test Delete resource", resourceTests.TestDelete)

	groupTests := &groupsTest{t}
	t.Run("Test ADD group", groupTests.TestADD)
}
