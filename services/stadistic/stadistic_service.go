package order

import (

	//REPOSITORIES

	"log"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	order_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/order"
	stadistic_comensal_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/stadistic-comensal"
)

func Import_OrderMade_Service(order_mades []models.Pg_Order_ToCopy) error {

	error_add_ordermades := order_repository.Pg_Insert_OrderMade(order_mades)
	if error_add_ordermades != nil {
		log.Fatal(error_add_ordermades)
	}

	return nil
}

func Import_OrderDetails_Service(order_details []models.Pg_Element) error {

	error_add_orderdetails := order_repository.Pg_Insert_OrderDetails(order_details)
	if error_add_orderdetails != nil {
		log.Fatal(error_add_orderdetails)
	}

	return nil
}

func Get_ComensalStadistic_All_Service(date_init string, date_end string, idcomensal int) (int, bool, string, models.Pg_Stadistic_Comensal) {

	//Enviamos los datos a la BD
	orders_comensal, error_add_order := stadistic_comensal_repository.Pg_Find_Stadistic(date_init, date_end, idcomensal)
	if error_add_order != nil {
		return 500, true, "Error interno en el servidor al buscar las ordenes, detalle: " + error_add_order.Error(), orders_comensal
	}

	return 200, false, "", orders_comensal
}
