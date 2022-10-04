package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find_ToExportFee() ([]models.Pg_ToExportFee, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var stadistic_tonotify_all []models.Pg_ToExportFee

	db := models.Conectar_Pg_DB()
	q := "SELECT om.informationbusiness->'idbusiness',COUNT(om.idorder),SUM(od.unitprice*od.quantity) FROM ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE isexportedtofee=false AND ismadebycomensal=true GROUP BY om.informationbusiness->'idbusiness'"
	rows, error_shown := db.Query(ctx, q)

	if error_shown != nil {
		return stadistic_tonotify_all, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var stadistic_toexportfee models.Pg_ToExportFee
		rows.Scan(&stadistic_toexportfee.IDBusiness, &stadistic_toexportfee.Orders, &stadistic_toexportfee.Amount)

		stadistic_tonotify_all = append(stadistic_tonotify_all, stadistic_toexportfee)
	}

	return stadistic_tonotify_all, nil
}
