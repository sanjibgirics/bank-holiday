package pkg

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8082/", nil)
	if err != nil {
		panic(err)
	}

	c := http.Client{}
	resp, err := c.Do(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	// Verify the response data
	responseBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(responseBody), "Welcome to UK bank holiday application")
}

// Verifying client is able to call the API and getting response without errors.

func TestGetHolidaysEnglandAndWalesNoBunting(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8082/getHolidaysEnglandAndWalesNoBunting", nil)
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, req)

	c := http.Client{}
	resp, err := c.Do(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	// Commenting below as verification might fail, due to change in source data as don't have
	// control on the  API from which we are fetching the data.

	// Verify response data
	// responseBody, err := io.ReadAll(resp.Body)
	// assert.Nil(t, err)
	// assert.Equal(t, "[{\"title\":\"Good Friday\",\"date\":\"2018-03-30\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2019-04-19\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2020-04-10\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Easter Monday\",\"date\":\"2020-04-13\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2021-04-02\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2022-04-15\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Bank Holiday for the State Funeral of Queen Elizabeth II\",\"date\":\"2022-09-19\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2023-04-07\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2024-03-29\",\"notes\":\"\",\"bunting\":false},{\"title\":\"Good Friday\",\"date\":\"2025-04-18\",\"notes\":\"\",\"bunting\":false}]\n", string(responseBody)
}

func TestGetHolidays2023(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8082/getHolidays2023", nil)
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, req)

	c := http.Client{}
	resp, err := c.Do(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestGetHolidays2024TitleAndDate(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8082/getHolidays2024TitleAndDate", nil)
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, req)

	c := http.Client{}
	resp, err := c.Do(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
