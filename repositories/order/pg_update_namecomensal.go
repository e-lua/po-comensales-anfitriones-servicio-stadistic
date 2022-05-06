package repositories

import (
	"context"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Update_NameComensal(inputname models.Mqtt_UpdateName) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UUPDATE ordermade SET informationcomensal = jsonb_set(informationcomensal, '{name}', '$1', false) WHERE informationcomensal->>'idcomensal'=$2;`
	if _, err := db.Exec(ctx, query, `"`+inputname.Name+`"`, strconv.Itoa(inputname.IdComensal)); err != nil {
		return err
	}

	return nil
}
