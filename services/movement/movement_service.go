package movement

import (

	//REPOSITORIES

	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	movement_anfitrion_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/movement-anfitrion"
)

func AddMovement_Service(idbusiness int, movement models.Pg_Movement) (int, bool, string, string) {

	timezone_int, _ := strconv.Atoi(movement.Timezone)
	movement.Dateregistered = movement.Dateregistered.Add(time.Hour * time.Duration(timezone_int))

	//Enviamos los datos a la BD
	error_add := movement_anfitrion_repository.Pg_Add(idbusiness, movement)
	if error_add != nil {
		return 500, true, "Error interno en el servidor al intentar registrar el movimiento, detalle: " + error_add.Error(), ""
	}

	return 200, false, "", "Movimiento agregado correctamente"
}

func UpdateMovement_Service(idbusiness int, idmovement int) (int, bool, string, string) {

	//Enviamos los datos a la BD
	error_update := movement_anfitrion_repository.Pg_Delete(idbusiness, idmovement)
	if error_update != nil {
		return 500, true, "Error interno en el servidor al intentar eliminar el movimiento, detalle: " + error_update.Error(), ""
	}

	return 200, false, "", "Movimiento eliminado correctamente"
}

func FindMovement_Service(idbusiness int, limit int, datestart string, dateend string) (int, bool, string, []models.Pg_Movement) {

	//Enviamos los datos a la BD
	movements, error_find := movement_anfitrion_repository.Pg_Find(idbusiness, limit, datestart, dateend)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar listar los movimientos, detalle: " + error_find.Error(), movements
	}

	return 200, false, "", movements
}
