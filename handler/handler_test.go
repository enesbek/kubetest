package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testUser := User{
		Name:    "Ahmet",
		Surname: "Mehmet",
	}

	// Assertions
	if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body, err := ioutil.ReadAll(rec.Body)
		if err != nil {
			fmt.Println(err)
		}
		actual := User{}
		err = json.Unmarshal(body, &actual)
	
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, testUser, actual)
	}
}
