package notification

import (
	"github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	stadistic_anfitrion_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/stadistic-anfitrion"
)

/*----------------------------NOTIFICATION-----------------------------*/

func Notify_Stadistic_Service(date string) (int, bool, string, []models.Pg_ToNotify) {

	stadistic_notify_all, error_add_order := stadistic_anfitrion_repository.Pg_Find_ToNotify(date)
	if error_add_order != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los datos de las estadisticas, detalle: " + error_add_order.Error(), stadistic_notify_all
	}

	return 201, false, "", stadistic_notify_all
}
