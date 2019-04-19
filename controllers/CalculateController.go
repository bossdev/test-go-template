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
	// return ec.HTML(200, "<p>sssa</p>")
	// log.Println(filepath.Abs("assets/images/410116.jpg"))
	// abPath, _ := filepath.Abs("")
	// log.Println(abPath)
	// f, _ := os.Open("assets/images/410116.jpg")
	// log.Println(f)
	// log.Println(ec.Request().Host)
	// log.Println(ec.Request().URL.Path)
	// return ec.Stream(http.StatusOK, "image/png", f)

	// return ec.HTML(200, "<img src='../assets/images/410116.jpg' />")

	// Month february is 12 because this month has leap year calculate
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
		// "host": abPath,
		"months": months,
		"years":  years,
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
	log.Println(calWeekday.Day, calWeekday.Month, calWeekday.Year)

	// Get front and back unit year Ex: 1900 (19 and 00)
	strYear := strconv.Itoa(calWeekday.Year)
	frontStrYear := strYear[:2]
	backStrYear := strYear[2:]
	frontUnitYear, _ := strconv.Atoi(frontStrYear)
	backUnitYear, _ := strconv.Atoi(backStrYear)
	if calWeekday.Month >= 11 {
		backUnitYear--
	}

	set1 := (13*calWeekday.Month - 1) / 5
	set2 := backUnitYear / 4
	set3 := frontUnitYear / 4

	num := calWeekday.Day + int(set1) + backUnitYear + int(set2) + int(set3) - 2*20
	log.Println(num)
	log.Println(num % 7)

	return ec.JSON(200, weekDay[num%7])
}