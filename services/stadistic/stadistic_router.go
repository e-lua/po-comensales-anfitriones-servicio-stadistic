package order

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func GetJWT_Anfitrion(jwt string, service int, module int, epic int, endpoint int) (int, bool, string, int, int) {
	//Obtenemos los datos del auth
	respuesta_b, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/trylogin?jwt=" + jwt + "&service=" + strconv.Itoa(service) + "&module=" + strconv.Itoa(module) + "&epic=" + strconv.Itoa(epic) + "&endpoint=" + strconv.Itoa(endpoint))
	var get_respuesta_b ResponseJWT_B
	error_decode_respuesta_b := json.NewDecoder(respuesta_b.Body).Decode(&get_respuesta_b)
	if error_decode_respuesta_b != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta_b.Error(), 0, 0
	}
	return 200, false, "", get_respuesta_b.Data.IdBusiness, get_respuesta_b.Data.IdRol
}

/*----------------------IMPORTING DATA ORDERS----------------------*/

func (sr *stadisticRouter_pg) Import_OrderMade(order_tocopy []models.Pg_Order_ToCopy) {

	//Enviamos los datos importados a registrar
	error_order_tocopy := Import_OrderMade_Service(order_tocopy)
	if error_order_tocopy != nil {
		log.Fatal(error_order_tocopy)
	}
}

/*
func (sr *stadisticRouter_pg) Import_OrderDetails(order_details []models.Pg_Element) {

	//Enviamos los datos importados a registrar
	error_order_details := Import_OrderDetails_Service(order_details)
	if error_order_details != nil {
		log.Fatal(error_order_details)
	}
}
*/
func (sr *stadisticRouter_pg) Import_NewNameComensal(name_comensal models.Mqtt_UpdateName) {

	//Enviamos los datos importados a registrar
	error_order_details := Import_NewNameComensal_Service(name_comensal)
	if error_order_details != nil {
		log.Fatal(error_order_details)
	}
}

/*----------------------GET STADISTICS - COMENSAL----------------------*/

func (sr *stadisticRouter_pg) Get_ComensalStadistic_All(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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

/*----------------------GET STADISTICS - ANFITRION----------------------*/

func (sr *stadisticRouter_pg) Get_AnfitrionStadistic_Orders(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, rol := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if rol != 1 {
		results := Response{Error: true, DataError: "Este rol no esta permitido para visualizar las estadísticas", Data: ""}
		return c.JSON(403, results)
	}

	//Recibimos la fecha inicial y final
	start_date := c.Request().URL.Query().Get("start_date")
	end_date := c.Request().URL.Query().Get("end_date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Get_AnfitrionStadistic_Orders_Service(start_date, end_date, data_idbusiness)
	results := Response_StadisticAnfitrion_Order{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (sr *stadisticRouter_pg) Get_AnfitrionStadistic_Incoming(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, rol := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if rol != 1 {
		results := Response{Error: true, DataError: "Este rol no esta permitido para visualizar las estadísticas", Data: ""}
		return c.JSON(403, results)
	}

	//Recibimos la fecha inicial y final
	start_date := c.Request().URL.Query().Get("start_date")
	end_date := c.Request().URL.Query().Get("end_date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Get_AnfitrionStadistic_Incoming_Service(start_date, end_date, data_idbusiness)
	results := Response_StadisticAnfitrion_Incoming{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (sr *stadisticRouter_pg) Get_AnfitrionStadistic_Comensales(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, rol := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if rol != 1 {
		results := Response{Error: true, DataError: "Este rol no esta permitido para visualizar las estadísticas", Data: ""}
		return c.JSON(403, results)
	}

	//Recibimos el id de la proveedor
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Recibimos el id de la proveedor
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Get_AnfitrionStadistic_Comensales_Service(data_idbusiness, limit_int, offset_int)
	results := Response_StadisticAnfitrion_Comensal{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (sr *stadisticRouter_pg) Get_ElementStadistic_ByDay(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness, rol := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"), 2, 2, 1, 3)
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}
	if rol != 1 {
		results := Response{Error: true, DataError: "Este rol no esta permitido para visualizar las estadísticas", Data: ""}
		return c.JSON(403, results)
	}

	//Recibimos el limit
	idelement := c.Param("idelement")
	idelement_int, _ := strconv.Atoi(idelement)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Get_ElementStadistic_ByDay_Service(idelement_int)
	results := Response_StadisticElements_ByDay{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
