package breschema_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/remiges-tech/alya/wscutils"
	crux "github.com/remiges-tech/crux/matching-engine"
	breschema "github.com/remiges-tech/crux/server/BRESchema"
	"github.com/remiges-tech/crux/testutils"
	"github.com/stretchr/testify/require"
)

func TestBRESchemaUpdate(t *testing.T) {
	testCases := schemaUpdateTestcase()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Setting up buffer
			payload := bytes.NewBuffer(testutils.MarshalJson(tc.RequestPayload))

			res := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPut, "/breschemaupdate", payload)
			require.NoError(t, err)

			r.ServeHTTP(res, req)

			require.Equal(t, tc.ExpectedHttpCode, res.Code)
			if tc.ExpectedResult != nil {
				jsonData := testutils.MarshalJson(tc.ExpectedResult)
				require.JSONEq(t, string(jsonData), res.Body.String())
			} else {
				jsonData, err := testutils.ReadJsonFromFile(tc.TestJsonFile)
				require.NoError(t, err)
				require.JSONEq(t, string(jsonData), res.Body.String())
			}
		})
	}

}

func schemaUpdateTestcase() []testutils.TestCasesStruct {

	schemaNewTestcase := []testutils.TestCasesStruct{

		{
			Name: "err- standard validation",
			RequestPayload: wscutils.Request{
				Data: breschema.BREUpdateSchemaRequest{},
			},

			ExpectedHttpCode: http.StatusBadRequest,
			TestJsonFile:     "./data/schema_update_validation_error.json",
		},
		{
			Name: "Success- Update schema",
			RequestPayload: wscutils.Request{
				Data: breschema.BRESchemaNewReq{
					Slice: 12,
					App:   "startmf",
					Class: "custonboarding",
					PatternSchema: []breschema.PatternSchema{
						{
							Attr:      "step",
							EnumVals:  []string{"aof", "nomauth", "kycvalid"},
							ValType:   "enum",
							LongDesc:  "",
							ShortDesc: "",
						},
						{
							Attr:      "stepfailed",
							ValType:   "bool",
							LongDesc:  "",
							ShortDesc: "",
						},
						{
							Attr:      "mode",
							EnumVals:  []string{"demat", "physical"},
							ValType:   "enum",
							LongDesc:  "",
							ShortDesc: "",
						},
					},
					ActionSchema: crux.ActionSchema_t{
						Tasks:      []string{"getcustdetails", "aof", "dpandbankaccvalid", "kycvalid"},
						Properties: []string{"nextstep", "done"},
					},
				},
			},

			ExpectedHttpCode: http.StatusOK,
			ExpectedResult: &wscutils.Response{
				Status:   wscutils.SuccessStatus,
				Data:     nil,
				Messages: nil,
			},
		},
	}
	return schemaNewTestcase
}
