package repositories

import (
	"context"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find_ComensalesByAnfitrion(idbusiness int, limit int, offset int) ([]models.Pg_Comensales, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "SELECT (informationcomensal->>'idcomensal')::integer,informationcomensal->>'name',informationcomensal->>'phonecontact',COUNT(idorder)as quantity FROM ordermade WHERE informationbusiness->>'idbusiness'=$1 GROUP BY informationcomensal ORDER BY quantity DESC LIMIT $2 OFFSET $3"
	rows, error_shown := db.Query(ctx, q, strconv.Itoa(idbusiness), limit, offset)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oComensalesByAnf []models.Pg_Comensales

	//Instanciamos un contador
	counter := 0

	if error_shown != nil {
		return oComensalesByAnf, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oComensal models.Pg_Comensales
		rows.Scan(&oComensal.IdComensal, &oComensal.Name, &oComensal.Phone, &oComensal.Orders)
		oComensalesByAnf = append(oComensalesByAnf, oComensal)
		counter += 1
	}

	//Si todo esta bien
	return oComensalesByAnf, nil
}
