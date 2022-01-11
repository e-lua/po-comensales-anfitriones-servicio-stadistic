package repositories

import (
	"context"
	"strconv"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find_Stadistic_Orders(date_init string, date_end string, idbusiness int) (models.Pg_Stadistic_Anfitrion_Orders, error) {

	var stadistic_anfitrion_order models.Pg_Stadistic_Anfitrion_Orders

	db := models.Conectar_Pg_DB()
	q := "SELECT (SELECT COUNT(idorder) as quantity FROM  ordermade AS om WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date),(select json_build_object('week',json_agg(w))from (SELECT EXTRACT(ISODOW FROM(schedule->>'daterequired')::date) as dayofweek,COUNT(idorder) as quantity FROM  ordermade AS om WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY schedule->>'daterequired') as w),(select json_build_object('days',json_agg(d))from (SELECT date_trunc('day',(schedule->>'daterequired')::date)::date AS day,count(idorder) as quantity FROM ordermade WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY schedule->>'daterequired' HAVING count(idorder)>0) as d),(select json_build_object('services',json_agg(s)) from ( SELECT om.service->'idservice' as idservice ,COUNT(om.service->>'idservice') as quantity FROM  ordermade AS om WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY om.service->'idservice') as s),(select json_build_object('payments',json_agg(p)) from ( SELECT om.payment->'idpayment' as idpayment,om.payment->'url' as url ,COUNT(om.payment->>'idpayment') as quantity FROM  ordermade AS om WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY om.payment->'idpayment',om.payment) as p),(select json_build_object('elements',json_agg(e)) from ( SELECT od.namee as name,SUM(od.quantity) as quantity FROM  ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY od.namee  ORDER BY quantity DESC LIMIT 10) as e)"
	error_shown := db.QueryRow(context.Background(), q, strconv.Itoa(idbusiness), date_init, date_end).Scan(&stadistic_anfitrion_order.Orders, &stadistic_anfitrion_order.Orders_by_week, &stadistic_anfitrion_order.Orders_by_day, &stadistic_anfitrion_order.Orders_by_service, &stadistic_anfitrion_order.Orders_by_payment, &stadistic_anfitrion_order.Elements)

	if error_shown != nil {

		return stadistic_anfitrion_order, error_shown
	}
	//Si todo esta bien
	return stadistic_anfitrion_order, nil
}
