package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		resource models.Resources
		response *consts.CustomError
	}{
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.ADD_SUCCESS,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.ADD_SUCCESS,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "localhost",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_IP_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_IP_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "somerandom",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test",
			},
			response: consts.INVALID_DOMAIN_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "",
			},
			response: consts.INVALID_DOMAIN_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			response: consts.INVALID_DOMAIN_ERROR,
		},
	}

	HTTPClient := &http.Client{}
	for _, test := range tests {
		t.Run(test.response.Message, func(t *testing.T) {
			body, err := json.Marshal(test.resource)
			if err != nil {
				t.Error(err.Error())
			}

			req, err := http.NewRequest(
				"POST",
				"http://localhost:3000/resource",
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
					"message didn't math: %s, %s\n%s",
					respData.Message,
					test.response.Message,
					respData.Detail,
				)
			}
		})
	}
}
