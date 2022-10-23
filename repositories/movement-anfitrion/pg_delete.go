package movement

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Delete(idmovement int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE Movement SET isdeleted=true WHERE idmovement=$1 AND idbusiness=$2`
	if _, err_connect := db.Exec(ctx, query, idmovement, idbusiness); err_connect != nil {
		return err_connect
	}

	return nil
}
