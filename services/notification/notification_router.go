package notification

import (
	"github.com/labstack/echo/v4"
)

var NotificationRouter_pg *notificationRouter_pg

type notificationRouter_pg struct {
}

/*----------------------------NOTIFICATION-----------------------------*/

func (nr *notificationRouter_pg) Notify_Stadistic(c echo.Context) error {

	date := c.Request().URL.Query().Get("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Notify_Stadistic_Service(date)
	results := Response_Notify_Stadistic{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
