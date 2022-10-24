package movement

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Find(idbusiness int, limit int, datestart string, datefinish string) ([]models.Pg_Movement, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "SELECT idmovement,dateregistered,description,amount,type,timezone FROM ordermade WHERE idbusiness=$1 AND dateregistered::date BETWEEN $2::date AND $3::date ORDER BY dateregistered ASC LIMIT $4"
	rows, error_shown := db.Query(ctx, q, idbusiness, datestart, datefinish, limit)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListMovements []models.Pg_Movement

	//Instanciamos un contador
	counter := 0

	if error_shown != nil {
		return oListMovements, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oMovement models.Pg_Movement
		rows.Scan(&oMovement.IdMovement, &oMovement.Dateregistered, &oMovement.Description, &oMovement.Amount, &oMovement.Type, &oMovement.Timezone)
		oListMovements = append(oListMovements, oMovement)
		counter += 1
	}

	//Si todo esta bien
	return oListMovements, nil
}
