package repositories

import (
	"context"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find_Stadistic_Incoming(date_init string, date_end string, idbusiness int) (models.Pg_Stadistic_Anfitrion_Incoming, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var stadistic_anfitrion_incoming models.Pg_Stadistic_Anfitrion_Incoming

	db := models.Conectar_Pg_DB()
	q := "SELECT (select concat(to_char((concat(datefinish,'+',((schedule->>'timezone')::integer*-1)::varchar(3))::timestamp with time zone)::date, 'DD/MM/YYYY'),' ',to_char((concat(datefinish,'+',((schedule->>'timezone')::integer*-1)::varchar(3))::timestamp with time zone)::time,'HH12:MI AM')) as lastdateorder FROM ordermade WHERE informationbusiness->>'idbusiness'=$1 ORDER BY concat(datefinish,'+',((schedule->>'timezone')::integer*-1)::varchar(3))::timestamp with time zone DESC LIMIT 1),(select json_build_object('total',json_agg(i))from (SELECT SUM(od.quantity *(od.unitprice-od.discount)) as incoming, od.typemoney as typemoney FROM  ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY od.typemoney) as i),(select json_build_object('week',json_agg(w)) from (SELECT EXTRACT(ISODOW FROM(schedule->>'daterequired')::date) as dayofweek,SUM(od.quantity *(od.unitprice-od.discount)) as incoming, od.typemoney AS typemoney FROM  ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY schedule->>'daterequired',od.typemoney) as w),(select json_build_object('days',json_agg(d)) from (SELECT date_trunc('day',(schedule->>'daterequired')::date)::date AS day,SUM(od.quantity *(od.unitprice-od.discount)) as incoming, od.typemoney as typemoney FROM ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY schedule->>'daterequired',od.typemoney) as d),(select json_build_object('services',json_agg(s))from ( SELECT om.service->'idservice' as idservice ,SUM(od.quantity *(od.unitprice-od.discount)) as incoming,od.typemoney as typemoney FROM  ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY om.service->'idservice',od.typemoney) as s),(select json_build_object('payments',json_agg(p)) from ( SELECT om.payment->'idpayment' as idpayment,om.payment->'url' as url,SUM(od.quantity *(od.unitprice-od.discount)) as incoming,od.typemoney as typemoney FROM  ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY om.payment->'idpayment',od.typemoney,om.payment) as p),(select json_build_object('elements',json_agg(e))from (SELECT od.namee as name,SUM(od.quantity *(od.unitprice-od.discount)) as incoming,od.typemoney as typemoney FROM  ordermade AS om JOIN orderdetails AS od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2::date AND $3::date GROUP BY od.namee,od.typemoney  ORDER BY incoming DESC LIMIT 10) as e)"
	error_shown := db.QueryRow(ctx, q, strconv.Itoa(idbusiness), date_init, date_end).Scan(&stadistic_anfitrion_incoming.Incoming_lastdateorder, &stadistic_anfitrion_incoming.Incoming, &stadistic_anfitrion_incoming.Incoming_by_week, &stadistic_anfitrion_incoming.Incoming_by_day, &stadistic_anfitrion_incoming.Incoming_by_service, &stadistic_anfitrion_incoming.Incoming_by_payment, &stadistic_anfitrion_incoming.Elements)

	if error_shown != nil {

		return stadistic_anfitrion_incoming, error_shown
	}
	//Si todo esta bien
	return stadistic_anfitrion_incoming, nil
}
