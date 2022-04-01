package export

import (

	//REPOSITORIES

	"log"

	export_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/export"
)

func Export_Stadistic_Service() (int, bool, string, string) {

	_, idorders, quantity, error_export := export_repository.Pg_Export_OrdersByElements()
	if error_export != nil {
		return 500, true, "Error en al intentar exportar los datos, detalles:" + error_export.Error(), ""
	}

	if quantity > 0 {
		error_update := export_repository.Pg_Update_Exported(idorders)
		if error_update != nil {
			return 500, true, "Error en al intentar actualizar los datos, detalles:" + error_update.Error(), ""
		}
	}

	log.Println(idorders)

	return 200, false, "", "Exportacion correcta"
}
