package transport_test

import (
	"net/http"
	"testing"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type ResourceTests struct {
	*testing.T
}

//
// func TestResources(t *testing.T) {
// 	tests := &ResourceTests{t}
//
// 	tests.Run("Test ADD resource", tests.TestAdd)
// }

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

func (tr *ResourceTests) TestDelete(t *testing.T) {
	tests := []httpTestType{
		{
			body: models.Resources{
				ServerIp: "167.235.27.246",
				Domains:  "alireza-baneshi.ir, somerandom.aslfk",
			},
			response: consts.DELETE_SUCCESS,
			method:   http.MethodDelete,
			url:      "http://localhost:3000/resource",
		},
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
	}

	for _, test := range tests {
		runHTTPSubtests(t, test)
	}
}
