package export

import (
	"log"

	"github.com/labstack/echo/v4"
)

var ExportRouter_pg *exportRouter_pg

type exportRouter_pg struct {
}

func (er *exportRouter_pg) Export_Stadistic() {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Export_Stadistic_Service()
	log.Println(status, boolerror, dataerror, data)
}

func (er *exportRouter_pg) Export_ToFee(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Export_ToFee_Service()
	results := Response_ToExportFee{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
