package order

import (

	//REPOSITORIES

	"log"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	order_repository "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/repositories/order"
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
