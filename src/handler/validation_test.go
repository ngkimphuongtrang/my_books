package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/httpkit"
)

func TestValidation_validateListParameters(t *testing.T) {
	v := NewValidation()

	tests := []struct {
		name            string
		page            string
		perPage         string
		expectedPage    int
		expectedPerPage int
		expectedErr     bool
	}{
		{
			name:            "valid parameters",
			page:            "1",
			perPage:         "10",
			expectedPage:    1,
			expectedPerPage: 10,
			expectedErr:     false,
		},
		{
			name:            "invalid page",
			page:            "abc",
			perPage:         "10",
			expectedPage:    0,
			expectedPerPage: 0,
			expectedErr:     true,
		},
		{
			name:            "empty page",
			expectedPage:    1,
			expectedPerPage: 30,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := &httpkit.RequestContext{
				Request: httptest.NewRequest("GET", "/?page="+tc.page+"&per_page="+tc.perPage, nil),
				Writer:  httptest.NewRecorder(),
			}

			page, perPage, ok := v.validateListParameters(ctx)
			if tc.expectedErr {
				require.False(t, ok)
				return
			}
			require.Equal(t, tc.expectedPage, page)
			require.Equal(t, tc.expectedPerPage, perPage)
		})
	}
}
