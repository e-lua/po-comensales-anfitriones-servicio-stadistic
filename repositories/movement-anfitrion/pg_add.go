package movement

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Add(idbusiness int, movement models.Pg_Movement) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO Movement(idbusiness,dateregistered,description,amount,type,timezone) VALUES ($1,$2,$3,$4,$5,$6)`
	if _, err_connect := db.Exec(ctx, query, idbusiness, movement.Dateregistered, movement.Description, movement.Amount, movement.Type, movement.Timezone); err_connect != nil {
		return err_connect
	}

	return nil
}
