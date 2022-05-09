package exportfile

import (
	"encoding/json"
	"net/http"

	"github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	"github.com/labstack/echo/v4"
)

var ExportfileRouter_pg *exportfileRouter_pg

type exportfileRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

func (efr *exportfileRouter_pg) ExportFile_Pedido(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)

	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha inicial y final
	start_date := c.Request().URL.Query().Get("start_date")
	end_date := c.Request().URL.Query().Get("end_date")

	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/worker/email")
	respuesta.Header.Set("Authorization", c.Request().Header.Get("Authorization"))
	var get_respuesta Response
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}

	var order_data models.Mqtt_Request_Order
	order_data.IDBusiness = data_idbusiness
	order_data.Email = get_respuesta.Data

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := ExportFile_Pedido_Service(order_data, start_date, end_date)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
