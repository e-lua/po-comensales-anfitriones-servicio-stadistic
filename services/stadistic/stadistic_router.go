package order

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	"github.com/labstack/echo/v4"
)

var StadisticRouter_pg *stadisticRouter_pg

type stadisticRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

/*----------------------IMPORTING DATA ORDERS----------------------*/

func (sr *stadisticRouter_pg) Import_OrderMade(order_tocopy []models.Pg_Order_ToCopy) {

	//Enviamos los datos importados a registrar
	error_order_tocopy := Import_OrderMade_Service(order_tocopy)
	if error_order_tocopy != nil {
		log.Fatal(error_order_tocopy)
	}
}

func (sr *stadisticRouter_pg) Import_OrderDetails(order_details []models.Pg_Element) {

	//Enviamos los datos importados a registrar
	error_order_details := Import_OrderDetails_Service(order_details)
	if error_order_details != nil {
		log.Fatal(error_order_details)
	}
}

/*----------------------GET STADISTICS----------------------*/

func (sr *stadisticRouter_pg) Get_ComensalStadistic_All(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha inicial y final
	start_date := c.Request().URL.Query().Get("start_date")
	end_date := c.Request().URL.Query().Get("end_date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Get_ComensalStadistic_All_Service(start_date, end_date, data_idcomensal)
	results := Response_StadisticComensal{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
