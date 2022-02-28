package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPeople(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getAllPeople(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var person []*models.Person
		json.Unmarshal(rec.Body.Bytes(), &person)
		assert.Equal(t, 5, len(person))
	}
}

func TestGetPersonByName(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people?first_name=John&last_name=Doe", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getAllPeople(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var person []*models.Person
		json.Unmarshal(rec.Body.Bytes(), &person)
		assert.Equal(t, 2, len(person))
	}
}

func TestGetPersonByPhone(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("phone_number", "+1 (800) 555-1414")
	req := httptest.NewRequest(http.MethodGet, "/people?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getAllPeople(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var person []*models.Person
		json.Unmarshal(rec.Body.Bytes(), &person)
		assert.Equal(t, 1, len(person))
	}
}

func TestGetPersonByPhoneEmptyParam(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people?phone_number=", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getAllPeople(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var person []*models.Person
		json.Unmarshal(rec.Body.Bytes(), &person)
		assert.Equal(t, 0, len(person))
	}
}

func TestGetPersonByNameParamEmpty(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people?first_name=&last_name=", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getAllPeople(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var person []*models.Person
		json.Unmarshal(rec.Body.Bytes(), &person)
		assert.Equal(t, 0, len(person))
	}
}

func TestGetPersonByIDInvalidID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("12345")

	if assert.NoError(t, getPersonByID(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetPersonByIDNotFound(t *testing.T) {
	e := echo.New()
	uuid := uuid.NewV4()
	req := httptest.NewRequest(http.MethodGet, "/people/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(uuid.String())

	if assert.NoError(t, getPersonByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestGetPersonByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("81eb745b-3aae-400b-959f-748fcafafd81")

	expected := &models.Person{
		ID:          uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81")),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1212",
	}

	if assert.NoError(t, getPersonByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var person *models.Person
		json.Unmarshal(rec.Body.Bytes(), &person)
		assert.Equal(t, expected, person)
	}
}
