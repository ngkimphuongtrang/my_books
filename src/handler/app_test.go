package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApp_CreateBook(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		reqBody            string
		expectedStatusCode int
	}{
		{
			name:               "success",
			reqBody:            `{"name":"Who moved my cheese?","author":"B"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "missing parameters",
			reqBody:            `{"name":"Who moved my cheese?"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("POST", "/books", strings.NewReader(tc.reqBody))
			rr := httptest.NewRecorder()

			testApp.mux.ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}

func TestApp_ListBooks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		reqParam           url.Values
		expectedStatusCode int
	}{
		{
			name:               "success_empty_param",
			reqParam:           url.Values{},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "success_with_params",
			reqParam:           url.Values{"name": []string{"Who moved my cheese?"}},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "ignore_invalid_param",
			reqParam:           url.Values{"invalid_param": []string{"Who moved my cheese?"}},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid_value",
			reqParam:           url.Values{"per_page": []string{"0"}},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/books?"+tc.reqParam.Encode(), nil)
			rr := httptest.NewRecorder()

			testApp.mux.ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}

func TestApp_CreateRead(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		reqBody            string
		expectedStatusCode int
	}{
		{
			name:               "success",
			reqBody:            `{"book_id": 1, "source": "hard_copy", "finished_date": "2023-02-01"}`,
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("POST", "/reads", strings.NewReader(tc.reqBody))
			rr := httptest.NewRecorder()

			testApp.mux.ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}
