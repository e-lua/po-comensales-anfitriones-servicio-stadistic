package movement

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	"github.com/labstack/echo/v4"
)

var Movement_pg *movement_pg

type movement_pg struct {
}

func GetJWT_Business(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT_Business
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

func (mv *movement_pg) AddMovement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Business(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}

	//Instanciamos la variable
	var movement models.Pg_Movement

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&movement)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar los datos de la orden, revise la estructura o los valores, detalles: " + err.Error(), Data: ""}
		return c.JSON(400, results)
	}

	if movement.Type < 0 || movement.Type > 1 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddMovement_Service(data_idbusiness, movement)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (mv *movement_pg) UpdateMovement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Business(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}

	//Recibimos el limit
	idmovement := c.Param("idmovement")
	idmovement_int, _ := strconv.Atoi(idmovement)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateMovement_Service(data_idbusiness, idmovement_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (mv *movement_pg) FindMovement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Business(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}

	//Recibimos la fecha inicial y final
	start_date := c.Request().URL.Query().Get("start_date")
	end_date := c.Request().URL.Query().Get("end_date")
	limit := c.Request().URL.Query().Get("limit")

	limit_int, _ := strconv.Atoi(limit)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindMovement_Service(data_idbusiness, limit_int, start_date, end_date)
	results := ResponseMovement{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
