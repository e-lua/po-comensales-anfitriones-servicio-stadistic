package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Stadistic_OrdersByElements(input_idelement int) ([]interface{}, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "select json_build_object(extract(isodow FROM (om.schedule->>'daterequired')::timestamp),COUNT(idelement)) from orderdetails AS od JOIN ordermade AS om ON od.idorder=om.idorder where od.idelement=$1 GROUP BY om.schedule->>'daterequired'"
	rows, error_shown := db.Query(ctx, q, input_idelement)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListInterfaces []interface{}

	if error_shown != nil {
		return oListInterfaces, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oInterface interface{}
		rows.Scan(&oInterface)

		oListInterfaces = append(oListInterfaces, oInterface)
	}
	//Si todo esta bien
	return oListInterfaces, nil
}
