package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Update_Exported(idorder_made []int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE ordermade SET isexportedtoinventory=true FROM (select * from  unnest($1::bigint[])) as ex(idrder) WHERE idorder=ex.idrder`
	if _, err := db.Exec(ctx, query, idorder_made); err != nil {
		return err
	}

	return nil
}
