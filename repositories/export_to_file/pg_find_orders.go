package export_to_file

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_Orders_ToFile(order_data models.Mqtt_Request_Order, date_start string, date_end string) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	quantity := 0

	db := models.Conectar_Pg_DB()
	q := "SELECT (om.dateregistered)::varchar(80),om.idorder,om.fourcode,om.schedule,om.informationbusiness,om.addressbusiness,om.informationcomensal,om.addresscomensal,om.note,om.service,om.payment,om.datarejected,json_agg((od.namee,od.category,od.typefood,od.urle,od.descriptione,od.insumos,od.unitprice,od.discount,od.costo,od.iva)),SUM(quantity*(unitprice-costo)),SUM((quantity*unitprice)-discount)+(service->>'price')::decimal(8,2),informationworker,ismadebycomensal FROM ordermade as om JOIN orderdetails as od ON om.idorder=od.idorder WHERE informationbusiness->>'idbusiness'=$1 AND (schedule->>'daterequired')::date BETWEEN $2 AND $3 GROUP BY om.idorder,om.fourcode,om.idstatus,om.schedule,om.informationbusiness,om.addressbusiness,om.informationcomensal,om.addresscomensal,om.note,om.service,om.payment,om.datarejected,om.dateregistered"
	rows, error_shown := db.Query(ctx, q, strconv.Itoa(order_data.IDBusiness), date_start, date_end)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListOrder := []models.Mqtt_Order{}

	if error_shown != nil {

		return error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		oOrder := models.Mqtt_Order{}
		rows.Scan(&oOrder.DateRegistered, &oOrder.IDOrder, &oOrder.FourCode, &oOrder.Schedule, &oOrder.Information_Business, &oOrder.Address_Busines, &oOrder.Information_Comensal, &oOrder.Address_Comensal, &oOrder.Note, &oOrder.Service, &oOrder.Payment, &oOrder.DataRejected, &oOrder.Elements, &oOrder.EstimatedProfit, &oOrder.TotalPrice, &oOrder.Information_Worker, &oOrder.Ismadebycomensal)
		oListOrder = append(oListOrder, oOrder)
		if oOrder.IDOrder > 0 {
			quantity += 1
		}
	}

	order_data.Orders = oListOrder

	if quantity > 0 {

		/*----------------------------MQTT----------------------------*/

		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes_element, error_serializar_ele := serialize_pedidos(order_data)
		if error_serializar_ele != nil {
			log.Error(error_serializar_ele)
		}

		error_publish_2 := ch.Publish("", "anfitrion/pedido_to_file", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes_element,
			})
		if error_publish_2 != nil {
			log.Error(error_publish_2)
		}

	}

	//Si todo esta bien
	return nil

}

//SERIALIZADORA PEDIDO
func serialize_pedidos(order_data models.Mqtt_Request_Order) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(order_data)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
