package repositories

import (
	"context"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find_Stadistic(date_init string, date_end string, idcomensal int) (models.Pg_Stadistic_Comensal, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var stadistic_comensal models.Pg_Stadistic_Comensal

	db := models.Conectar_Pg_DB()
	q := `SELECT
	(select concat(to_char((concat(datefinish,'+',((schedule->>'timezone')::integer*-1)::varchar(3))::timestamp with time zone)::date, 'DD/MM/YYYY'),' ',to_char((concat(datefinish,'+',((schedule->>'timezone')::integer*-1)::varchar(3))::timestamp with time zone)::time,'HH12:MI AM')) as lastdateorder FROM ordermade WHERE informationcomensal->>'idcomensal'=$3 ORDER BY concat(datefinish,'+',((schedule->>'timezone')::integer*-1)::varchar(3))::timestamp with time zone DESC LIMIT 1),
	(SELECT SUM(totalsales)+SUM((service->>'price')::decimal(8,2)) FROM ORDERMADE WHERE (schedule->>'daterequired')::date BETWEEN $1::date AND $2::date AND informationcomensal->>'idcomensal'=$3),
	(SELECT COUNT(idorder) as quantity FROM  ordermade WHERE (schedule->>'daterequired')::date BETWEEN $1::date AND $2::date AND informationcomensal->>'idcomensal'=$3),
	(SELECT json_build_object('week',json_agg(w)) FROM (SELECT out1.dayofweek,out1.outgoing1+out2.outgoing2 as outgoing FROM (SELECT EXTRACT(ISODOW FROM(schedule->>'daterequired')::date) as dayofweek,SUM(totalsales) as outgoing1 FROM ordermade WHERE schedule->>'daterequired')::date BETWEEN $1::date AND $2::date AND informationcomensal->>'idcomensal'=$3 GROUP BY schedule->>'daterequired') AS out1 JOIN (SELECT EXTRACT(ISODOW FROM(schedule->>'daterequired')::date) as dayofweek,SUM((service->>'price')::decimal(8,2)) as outgoing2 FROM  ordermade WHERE schedule->>'daterequired')::date BETWEEN $1::date AND $2::date AND informationcomensal->>'idcomensal'=$3 GROUP BY schedule->>'daterequired') AS out2 on out1.dayofweek=out2.dayofweek) as w),
	(select json_build_object('typefood',json_agg(ca)) from (SELECT typefood as typefood,COUNT(idelement) as quantity FROM orderdetail WHERE (schedule->>'daterequired')::date BETWEEN $1::date AND $2::date AND informationcomensal->>'idcomensal'=$3 AND typefood<>'' GROUP BY typefood) as ca)
	`
	error_shown := db.QueryRow(ctx, q, date_init, date_end, strconv.Itoa(idcomensal)).Scan(&stadistic_comensal.Orders_lastdateorder, &stadistic_comensal.Outgoing, &stadistic_comensal.Orders, &stadistic_comensal.Outgoing_Week, &stadistic_comensal.Orders_typefood)

	if error_shown != nil {

		return stadistic_comensal, error_shown
	}
	//Si todo esta bien
	return stadistic_comensal, nil
}
