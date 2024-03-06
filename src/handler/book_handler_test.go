package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/trangnkp/my_books/src/httpkit"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBookHandler_Create(t *testing.T) {
	t.Parallel()

	requestBody := `{"name":"Who moved my cheese?","author":"A"}`
	ctx := &httpkit.RequestContext{
		Request: httptest.NewRequest("POST", "/books", strings.NewReader(requestBody)),
		Writer:  httptest.NewRecorder(),
	}
	testApp.handleCreateBook(ctx)
	assert.Equal(t, http.StatusOK, ctx.Writer.(*httptest.ResponseRecorder).Code)
}
