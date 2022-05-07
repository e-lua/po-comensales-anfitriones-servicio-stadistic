package exportfile

import (
	//REPOSITORIES

	"github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	export_file "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/export_to_file"
)

func ExportFile_Pedido_Service(order_data models.Mqtt_Request_Order, date_start string, date_end string) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_export := export_file.Pg_Orders_ToFile(order_data, date_start, date_end)
	if error_export != nil {
		return 500, true, "Error en el servidor interno al intentar exportar los insumos, detalles: " + error_export.Error(), ""
	}

	return 201, false, "", "Enviado correctamente"
}
