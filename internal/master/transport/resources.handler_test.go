package transport_test

import (
	"fmt"
	"testing"

	"github.com/harleywinston/x-manager/internal/master/models"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		resource        models.Resources
		responseMessage string
	}{
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			responseMessage: "oiajsfi",
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			responseMessage: "oiajsfi",
		},
		{
			resource: models.Resources{
				CloudflareDomains: "long-tooth-0da8.harleywinston19935891.workers.dev, long-tooth-0da8.harleywinston19935891.workers.dev",
				ServerIp:          "167.235.27.246",
				Domains:           "alireza-baneshi.ir, somerandom.aslfk",
				BrdigeDomain:      "test.darkube.ir",
			},
			responseMessage: "oiajsfi",
		},
	}

	for _, test := range tests {
		fmt.Println(test.responseMessage)
	}
}
