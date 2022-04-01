package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_Export_OrdersByElements() ([]models.Pg_Export_ByElement, []int64, int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Variable a exportar
	var ids []int64

	//Creamos un contador para evitar tener que enviar datos vacios a traves de la cola
	quantity := 0

	db := models.Conectar_Pg_DB()
	q := "SELECT od.idorder,od.idelement, COUNT(od.idelement),CONCAT(om.schedule->>'daterequired',' ',om.schedule->>'starttime')::timestamp FROM orderdetails AS od JOIN ordermade AS om ON od.idorder=om.idorder WHERE om.isexportedtoinventory=false GROUP BY od.idorder,od.idelement,CONCAT(om.schedule->>'daterequired',' ',om.schedule->>'starttime')::timestamp"
	rows, error_shown := db.Query(ctx, q)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListExpByElement []models.Pg_Export_ByElement

	if error_shown != nil {
		return oListExpByElement, ids, quantity, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oExportByElement := models.Pg_Export_ByElement{}
		var idorder int64
		rows.Scan(&idorder, &oExportByElement.IdElement, &oExportByElement.Quantity, &oExportByElement.Datetime)
		if oExportByElement.IdElement > 0 {
			quantity = quantity + 1
		}
		oListExpByElement = append(oListExpByElement, oExportByElement)
		ids = append(ids, idorder)
	}

	if quantity >= 1 {
		/*---------------------------MQTT---------------------------*/
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize_exportstadistic(oListExpByElement)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "anfitrion/ordersperelement", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

		/*----------------------------------------------------------*/
	}

	//Si todo esta bien
	return oListExpByElement, ids, quantity, nil
}

//SERIALIZADORA
func serialize_exportstadistic(serialize_export_tocpy []models.Pg_Export_ByElement) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_export_tocpy)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
