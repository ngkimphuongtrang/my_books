package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/internal/container"
	"github.com/trangnkp/my_books/src/internal/httpkit"
	"github.com/trangnkp/my_books/src/model"
)

func TestReadHandler_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		needSeedBook     bool
		requestBody      container.Map
		expectedResponse httpkit.Response
	}{
		{
			name:         "success",
			needSeedBook: true,
			requestBody:  container.Map{"source": "hard_copy", "language": "VI", "finished_date": time.Now().Format("2006-01-02")},
			expectedResponse: httpkit.Response{
				StatusCode: 200,
				Verdict:    "success",
				Message:    "read is created successfully",
			},
		},
		{
			name:         "success_with_empty_language",
			needSeedBook: true,
			requestBody:  container.Map{"source": "hard_copy", "finished_date": time.Now().Format("2006-01-02")},
			expectedResponse: httpkit.Response{
				StatusCode: 200,
				Verdict:    "success",
				Message:    "read is created successfully",
			},
		},
		{
			name:        "invalid_finished_date",
			requestBody: container.Map{"source": "", "language": "", "finished_date": time.Now().Format(time.RFC3339)},
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "invalid_parameters",
				Message:    "parameters are invalid",
			},
		},
		{
			name:        "missing_params",
			requestBody: container.Map{"source": "", "language": "C", "finished_date": time.Now().Format("2006-01-02"), "book_id": 1},
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "missing_parameters",
				Message:    "some required parameters are missing",
			},
		},
		{
			name:        "missing_params_book_id_and_book_name",
			requestBody: container.Map{"source": "", "language": "C", "finished_date": time.Now().Format("2006-01-02")},
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "missing_parameters",
				Message:    "some required parameters are missing",
			},
		},
		{
			name:        "invalid_source_read",
			requestBody: container.Map{"source": "library", "language": "VI", "finished_date": time.Now().Format("2006-01-02"), "book_id": 1},
			expectedResponse: httpkit.Response{
				StatusCode: 400,
				Verdict:    "invalid_parameters",
				Message:    "source is invalid",
			},
		},
		{
			name:        "book_not_found",
			requestBody: container.Map{"source": "hard_copy", "finished_date": time.Now().Format("2006-01-02"), "book_id": 100},
			expectedResponse: httpkit.Response{
				StatusCode: 404,
				Verdict:    "record_not_found",
				Message:    "book_id not found",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.needSeedBook {
				//TODO: restrict not create empty string book in database
				randomBook := &model.Book{
					Name:   tc.name,
					Author: tc.name,
				}
				err := testApp.stores.BookStore.Create(context.Background(), randomBook)
				require.NoError(t, err)
				tc.requestBody["book_id"] = randomBook.ID
			}
			requestBody, err := tc.requestBody.JSON()
			require.NoError(t, err)
			ctx := &httpkit.RequestContext{
				Request: httptest.NewRequest("POST", "/reads", strings.NewReader(requestBody)),
				Writer:  httptest.NewRecorder(),
			}
			readHandler.handleCreateRead(ctx)
			responseRecorder, _ := ctx.Writer.(*httptest.ResponseRecorder)

			var response httpkit.Response
			rr := ctx.Writer.(*httptest.ResponseRecorder).Result().Body
			defer rr.Close()

			err = json.NewDecoder(rr).Decode(&response)
			log.Println(response.Message, response.Data)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResponse.StatusCode, responseRecorder.Code)
			require.Equal(t, tc.expectedResponse.Verdict, response.Verdict)
			require.Equal(t, tc.expectedResponse.Message, response.Message)
		})
	}
}

func TestReadHandler_List(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name             string
		needSeedRead     bool
		expectedResponse httpkit.Response
	}{
		{
			name:         "success",
			needSeedRead: true,
			expectedResponse: httpkit.Response{
				StatusCode: 200,
				Verdict:    "success",
				Message:    "list reads successfully",
			},
		},
		{
			name:         "no_read",
			needSeedRead: false,
			expectedResponse: httpkit.Response{
				StatusCode: 200,
				Verdict:    "success",
				Message:    "list reads successfully",
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.needSeedRead {
				randomRead := &model.Read{
					BookID:       1,
					Source:       "hard_copy",
					FinishedDate: model.NewDate(2024, 03, 24),
				}
				err := testApp.stores.ReadStore.Create(context.Background(), randomRead)
				require.NoError(t, err)
			}
			ctx := &httpkit.RequestContext{
				Request: httptest.NewRequest("GET", "/reads", nil),
				Writer:  httptest.NewRecorder(),
			}
			readHandler.handleListReads(ctx)
			responseRecorder, _ := ctx.Writer.(*httptest.ResponseRecorder)
			var response httpkit.Response
			rr := ctx.Writer.(*httptest.ResponseRecorder).Result().Body
			defer rr.Close()
			err := json.NewDecoder(rr).Decode(&response)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResponse.StatusCode, responseRecorder.Code)
			require.Equal(t, tc.expectedResponse.Verdict, response.Verdict)
			require.Equal(t, tc.expectedResponse.Message, response.Message)
		})
	}
}
