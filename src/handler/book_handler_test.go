package handler

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
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
			bookHandler.handleCreateBook(ctx)
			responseRecorder, _ := ctx.Writer.(*httptest.ResponseRecorder)
			assert.Equal(t, tc.expectedResponse.StatusCode, responseRecorder.Code)

			var response httpkit.Response
			rr := ctx.Writer.(*httptest.ResponseRecorder).Result().Body
			defer rr.Close()

			err := json.NewDecoder(rr).Decode(&response)
			require.NoError(t, err)
			require.Equal(t, tc.expectedResponse.Verdict, response.Verdict)
		})
	}
}

func TestBookHandler_List(t *testing.T) {
	t.Parallel()

	ctx := &httpkit.RequestContext{
		Request: httptest.NewRequest("GET", "/books?per_page=5", nil),
		Writer:  httptest.NewRecorder(),
	}
	bookHandler.handleListBooks(ctx)
	responseRecorder, _ := ctx.Writer.(*httptest.ResponseRecorder)
	assert.Equal(t, 200, responseRecorder.Code)

	rr := ctx.Writer.(*httptest.ResponseRecorder).Result().Body
	defer rr.Close()

	body, _ := container.CreateMapFromReader(rr)
	log.Println("data", body["data"])
	data, ok := body["data"].(map[string]interface{})
	require.True(t, ok)
	items, ok := data["items"].([]interface{})
	log.Println(items)
	require.True(t, ok)
	require.NotZero(t, len(items))
}
