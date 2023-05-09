package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
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

type ResourceTests struct {
	*testing.T
}

func TestResources(t *testing.T) {
	tests := &ResourceTests{t}

	tests.Run("Test ADD resource", tests.TestAdd)
}

func (tr *ResourceTests) TestAdd(t *testing.T) {
	tests := []httpTestType{
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.ADD_SUCCESS,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.DUPLICATE_RECORD_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "localhost",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_IP_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_IP_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "somerandom",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test",
			},
			response: consts.INVALID_DOMAIN_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "",
			},
			response: consts.INVALID_DOMAIN_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
		{
			body: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
			method:   http.MethodPost,
			url:      "http://localhost:3000/resource",
		},
	}

	for _, test := range tests {
		runHTTPSubtests(t, test)
	}
}

//
// func TestGet(t *testing.T) {
// 	tests := []testType{
// 		{},
// 	}
// }
