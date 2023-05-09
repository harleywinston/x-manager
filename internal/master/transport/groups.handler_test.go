package transport_test

import (
	"net/http"
	"testing"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type groupsTest struct {
	*testing.T
}

func TestGroups(t *testing.T) {
	tests := &groupsTest{t}

	t.Run("Test ADD group", tests.TestADD)
}

func (tg *groupsTest) TestADD(t *testing.T) {
	tests := []httpTestType{
		{
			url:    "http://localhost:3000/group",
			method: http.MethodGet,
			body: models.Groups{
				ResourcesID: 1,
				Mode:        "Direct",
			},
			response: consts.ADD_SUCCESS,
		},
		{
			url:    "http://localhost:3000/group",
			method: http.MethodGet,
			body: models.Groups{
				ResourcesID: 1,
				Mode:        "InDirect",
			},
			response: consts.ADD_SUCCESS,
		},
		{
			url:    "http://localhost:3000/group",
			method: http.MethodGet,
			body: models.Groups{
				ResourcesID: 999999,
				Mode:        "Direct",
			},
			response: consts.RECOURSE_ID_NOT_VALID_ERROR,
		},
		{
			url:    "http://localhost:3000/group",
			method: http.MethodGet,
			body: models.Groups{
				ResourcesID: 1,
				Mode:        "oaijsdfoi",
			},
			response: consts.INVALID_GROUP_MODE_ERROR,
		},
	}

	for _, test := range tests {
		runHTTPSubtests(t, test)
	}
}
