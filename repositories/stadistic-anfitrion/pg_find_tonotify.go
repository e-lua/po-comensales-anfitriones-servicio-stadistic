package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find_ToNotify(date string) ([]models.Pg_ToNotify, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var stadistic_tonotify_all []models.Pg_ToNotify

	db := models.Conectar_Pg_DB()
	q := "SELECT (informationbusiness->'idbusiness')::bigint,COUNT(idorder),SUM(totalsales),SUM(totalsales-totaliva-totalfee),SUM(totalsales-totaliva-totalfee-totalcost) FROM ordermade WHERE dateregistered::date=$1::date GROUP BY informationbusiness->'idbusiness'"
	rows, error_shown := db.Query(ctx, q, date)

	if error_shown != nil {
		return stadistic_tonotify_all, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var stadistic_tonotify models.Pg_ToNotify
		rows.Scan(&stadistic_tonotify.IDBusiness, &stadistic_tonotify.Orders, &stadistic_tonotify.GrossIncoming, &stadistic_tonotify.NetIncoming, &stadistic_tonotify.NetUtility)

		stadistic_tonotify_all = append(stadistic_tonotify_all, stadistic_tonotify)
	}

	return stadistic_tonotify_all, nil
}
