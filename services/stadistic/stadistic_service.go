package order

import (

	//REPOSITORIES

	"log"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	order_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/order"
	stadistic_anfitrion_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/stadistic-anfitrion"
	stadistic_comensal_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/stadistic-comensal"
	stadistic_element_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/stadistic_elements"
)

func Import_NewNameComensal_Service(input_name models.Mqtt_UpdateName) error {

	error_add_ordermades := order_repository.Pg_Update_NameComensal(input_name)
	if error_add_ordermades != nil {
		log.Fatal(error_add_ordermades)
	}

	return nil
}

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

func Get_AnfitrionStadistic_Orders_Service(date_init string, date_end string, idbusiness int) (int, bool, string, models.Pg_Stadistic_Anfitrion_Orders) {

	//Enviamos los datos a la BD
	orders, error_add_order := stadistic_anfitrion_repository.Pg_Find_Stadistic_Orders(date_init, date_end, idbusiness)
	if error_add_order != nil {
		return 500, true, "Error interno en el servidor al buscar las ordenes, detalle: " + error_add_order.Error(), orders
	}

	return 200, false, "", orders
}

func Get_AnfitrionStadistic_Incoming_Service(date_init string, date_end string, idbusiness int) (int, bool, string, models.Pg_Stadistic_Anfitrion_Incoming) {

	//Enviamos los datos a la BD
	incoming, error_add_order := stadistic_anfitrion_repository.Pg_Find_Stadistic_Incoming(date_init, date_end, idbusiness)
	if error_add_order != nil {
		return 500, true, "Error interno en el servidor al buscar las ordenes, detalle: " + error_add_order.Error(), incoming
	}

	return 200, false, "", incoming
}

func Get_AnfitrionStadistic_Comensales_Service(idbusiness int, limit int, offset int) (int, bool, string, models.Pg_ComensalesByAnfitrion) {

	//Enviamos los datos a la BD
	comensals, error_add_order := stadistic_anfitrion_repository.Pg_Find_ComensalesByAnfitrion(idbusiness, limit, offset)
	if error_add_order != nil {
		return 500, true, "Error interno en el servidor al buscar los comensales, detalle: " + error_add_order.Error(), comensals
	}

	return 200, false, "", comensals
}

func Get_ElementStadistic_ByDay_Service(input_idelement int) (int, bool, string, []interface{}) {

	//Enviamos los datos a la BD
	elements_perday, error_add_order := stadistic_element_repository.Pg_Stadistic_OrdersByElements(input_idelement)
	if error_add_order != nil {
		return 500, true, "Error interno en el servidor al buscar las estadisticas de elementos por dia, detalle: " + error_add_order.Error(), elements_perday
	}

	return 200, false, "", elements_perday
}
