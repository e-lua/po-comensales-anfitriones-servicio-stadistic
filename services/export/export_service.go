package export

import (

	//REPOSITORIES

	"github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
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

	return 200, false, "", "Exportacion correcta"
}

func Export_ToFee_Service() (int, bool, string, []models.Pg_ToExportFee) {

	data_to_export, error_export := export_repository.Pg_Find_ToExportFee()
	if error_export != nil {
		return 500, true, "Error en al intentar exportar los datos, detalles:" + error_export.Error(), data_to_export
	}

	error_update := export_repository.Pg_Update_ExportedFee()
	if error_update != nil {
		return 500, true, "Error en al intentar actualizar los datos, detalles:" + error_update.Error(), data_to_export
	}

	return 200, false, "", data_to_export
}
