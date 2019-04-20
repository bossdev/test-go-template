package controllers

import (
	"log"
	"net/http"
	"strconv"
	"test/weekday/structures"
	"time"

	"github.com/labstack/echo"
)

func GetCalculateWeekdayPage(ec echo.Context) error {
	// Month (March = 1, April = 2 ... December = 10 and january = 11, February = 12)
	months := []map[string]interface{}{
		{"num": 11, "month": "January", "amountDay": 31},
		{"num": 12, "month": "February", "amountDay": 28},
		{"num": 1, "month": "March", "amountDay": 31},
		{"num": 2, "month": "April", "amountDay": 30},
		{"num": 3, "month": "May", "amountDay": 31},
		{"num": 4, "month": "June", "amountDay": 30},
		{"num": 5, "month": "July", "amountDay": 31},
		{"num": 6, "month": "August", "amountDay": 31},
		{"num": 7, "month": "September", "amountDay": 30},
		{"num": 8, "month": "October", "amountDay": 31},
		{"num": 9, "month": "November", "amountDay": 30},
		{"num": 10, "month": "December", "amountDay": 31},
	}

	// Get current year
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now()
	currentDate := now.In(loc)

	rangeYear := map[string]int{
		"start": 1900,
		"to":    currentDate.Year(),
	}
	var years []int
	for i := rangeYear["start"]; i <= rangeYear["to"]; i++ {
		years = append(years, i)
	}

	return ec.Render(http.StatusOK, "weekday.html", map[string]interface{}{
		"months":      months,
		"years":       years,
		"currentYear": years[len(years)-1],
	})
}

func CalculateWeekdayOfDate(ec echo.Context) error {
	weekDay := map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}

	// Bind request body to struct
	calWeekday := new(structures.CalculateWeekday)
	if err := ec.Bind(calWeekday); err != nil {
		log.Println(err)
	}

	// Get front and back unit year Ex: 1900 (front is 19 and back is 00)
	strYear := strconv.Itoa(calWeekday.Year)
	frontStrYear := strYear[:2]
	backStrYear := strYear[2:]
	frontUnitYear, _ := strconv.Atoi(frontStrYear)
	backUnitYear, _ := strconv.Atoi(backStrYear)
	if calWeekday.Month >= 11 {
		backUnitYear--
	}
	set1 := ((2.6 * float64(calWeekday.Month)) - 0.2)
	set2 := backUnitYear / 4
	set3 := frontUnitYear / 4
	num := (calWeekday.Day + int(set1) - (2 * frontUnitYear) + backUnitYear + int(set2) + int(set3)) % 7
	if num < 0 {
		num += 7
	}

	res := map[string]interface{}{
		"status":  "success",
		"message": weekDay[num],
	}

	return ec.JSON(http.StatusOK, res)
}
