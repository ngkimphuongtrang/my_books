package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/httpkit"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBookHandler_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		requestBody      string
		expectedResponse httpkit.Response
	}{
		{
			name:        "success",
			requestBody: `{"name":"Who moved my cheese?","author":"A"}`,
			expectedResponse: httpkit.Response{
				StatusCode: 200,
				Verdict:    "success",
			},
		},
		{
			name:        "missing_parameters",
			requestBody: `{"name":"Who moved my cheese?"}`,
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "missing_parameters",
			},
		},
		{
			name:        "success_with_extra_param",
			requestBody: `{"name":"Who moved my cheese2?", "author":"X",date":"2023-02-01"}`,
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "invalid_parameters",
			},
		},
		{
			name:        "corrupted_body",
			requestBody: `{"name":"Who moved my cheese2?",}`,
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "invalid_parameters",
			},
		},
		{
			name:        "duplicated_book",
			requestBody: `{"name":"duplicated", "author":"duplicated"}`,
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "existed_record",
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := &httpkit.RequestContext{
				Request: httptest.NewRequest("POST", "/books", strings.NewReader(tc.requestBody)),
				Writer:  httptest.NewRecorder(),
			}
			testApp.handleCreateBook(ctx)
			responseRecorder := ctx.Writer.(*httptest.ResponseRecorder)
			assert.Equal(t, tc.expectedResponse.StatusCode, responseRecorder.Code)

			var response httpkit.Response
			rr := ctx.Writer.(*httptest.ResponseRecorder).Result().Body
			err := json.NewDecoder(rr).Decode(&response)
			require.NoError(t, err)
			require.Equal(t, tc.expectedResponse.Verdict, response.Verdict)
		})
	}
}
