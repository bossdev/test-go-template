package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"test/weekday/structures"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type CaseReault struct {
	Day     int
	Month   int
	Year    int
	Weekday string
}

func caseResults() []CaseReault {
	// Month (March = 1, April = 2 ... December = 10 and january = 11, February = 12)
	// day , month , year , weekday of date
	dataCase := []CaseReault{
		{1, 11, 1900, "Monday"},
		{12, 1, 1903, "Thursday"},
		{29, 12, 1908, "Saturday"},
		{16, 5, 1910, "Saturday"},
		{3, 7, 1940, "Tuesday"},
		{8, 8, 1962, "Monday"},
		{27, 10, 1975, "Saturday"},
		{20, 11, 1978, "Friday"},
		{12, 2, 1981, "Sunday"},
		{19, 12, 1988, "Friday"},
		{4, 3, 2001, "Friday"},
		{30, 6, 2004, "Monday"},
		{21, 9, 2007, "Wednesday"},
		{13, 8, 2010, "Wednesday"},
		{5, 4, 2012, "Tuesday"},
		{29, 12, 2016, "Monday"},
		{17, 9, 2017, "Friday"},
		{31, 10, 2018, "Monday"},
		{30, 10, 2019, "Monday"},
		{16, 11, 2019, "Wednesday"},
	}
	return dataCase
}

func TestCalculateWeekdayOfDate(t *testing.T) {
	caseResults := caseResults()
	e := echo.New()

	for _, value := range caseResults {
		mockDataDate := structures.CalculateWeekday{
			value.Day, value.Month, value.Year,
		}
		m, _ := json.Marshal(mockDataDate)

		// Set request body
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(m)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resRecorder := httptest.NewRecorder()
		echoContext := e.NewContext(req, resRecorder)

		// Call function
		CalculateWeekdayOfDate(echoContext)
		dataBody := make(map[string]interface{})
		json.NewDecoder(resRecorder.Body).Decode(&dataBody)
		assert.Equal(t, value.Weekday, dataBody["message"])
	}
}
