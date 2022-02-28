package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	"github.com/sirupsen/logrus"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

var log = logrus.New()

// getAllPeople returns people records
func getAllPeople(c echo.Context) error {
	params := c.QueryParams()
	firstName := params.Get("first_name")
	lastName := params.Get("last_name")
	phone := params.Get("phone_number")
	people := []*models.Person{}

	if len(params) == 0 {
		people = models.AllPeople()
		return c.JSON(http.StatusOK, people)
	}

	if len(firstName) != 0 && len(lastName) != 0 {
		people = models.FindPeopleByName(firstName, lastName)
		return c.JSON(http.StatusOK, people)
	}

	if len(phone) != 0 {
		people = models.FindPeopleByPhoneNumber(phone)
		return c.JSON(http.StatusOK, people)
	}

	return c.JSON(http.StatusOK, people)
}

// getPersonByID returns person record by id
func getPersonByID(c echo.Context) error {
	id := c.Param("id")
	personID, err := uuid.FromString(id)

	// Invalid Person ID
	if err != nil {
		log.Errorf("Invalid Person ID %s", id)
		return c.NoContent(http.StatusBadRequest)
	}

	person, err := models.FindPersonByID(personID)
	if err != nil {
		var notFoundErr *models.NotFoundError
		if errors.As(err, &notFoundErr) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, person)
}
