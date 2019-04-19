package routes

import (
	"test/weekday/controllers"

	"github.com/labstack/echo"
)

type Routers struct {
	Ec *echo.Echo
}

func (rts Routers) GetRouter() {
	rts.Ec.GET("/", controllers.GetCalculateWeekdayPage)
	rts.Ec.POST("/calWeekday", controllers.CalculateWeekdayOfDate)
}
