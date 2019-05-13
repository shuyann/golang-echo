package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*User{
		"foo@com": {"foo", "foo@com"},
	}
	userJSON = `{"name":"foo","email":"foo@com"}` + "\n"
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	// create request for test
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	// set request header
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// create recorder
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	// create request for test
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// create recorder
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// set context parameter(path,param,value)
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("foo@com")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
