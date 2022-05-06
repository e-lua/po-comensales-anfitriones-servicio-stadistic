package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	"github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
	export "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/services/export"
	export_file "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/services/exportfile"
	stadistic "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/services/stadistic"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go Consume_OrderMade()
	go Consume_OrderDetails()
	go Consume_UpdateName()
	go Export_Stadistc()

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*===========CARTA===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_stadistic := version_1.Group("/stadistic")
	router_stadistic.GET("/comensal", stadistic.StadisticRouter_pg.Get_ComensalStadistic_All)
	router_stadistic.GET("/anfitrion/order", stadistic.StadisticRouter_pg.Get_AnfitrionStadistic_Orders)
	router_stadistic.GET("/anfitrion/incoming", stadistic.StadisticRouter_pg.Get_AnfitrionStadistic_Incoming)
	router_stadistic.GET("/anfitrion/comensales", stadistic.StadisticRouter_pg.Get_AnfitrionStadistic_Comensales)
	router_stadistic.GET("/anfitrion/element/:idelement", stadistic.StadisticRouter_pg.Get_ElementStadistic_ByDay)
	router_stadistic.GET("/anfitrion/sendtoemail", export_file.ExportfileRouter_pg.ExportFile_Pedido)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "4320"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

func Consume_OrderMade() {
	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("comensal/ordermade", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop := make(chan bool)

	go func() {
		for {
			time.Sleep(10 * time.Minute)
			for d := range msgs {
				var order_tocopy []models.Pg_Order_ToCopy
				buf := bytes.NewBuffer(d.Body)
				decoder := json.NewDecoder(buf)
				err_consume := decoder.Decode(&order_tocopy)
				if err_consume != nil {
					log.Fatal("Error decoding")
				}
				stadistic.StadisticRouter_pg.Import_OrderMade(order_tocopy)
			}
		}
	}()

	<-noStop
}

func Consume_OrderDetails() {
	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("comensal/orderdetails", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop2 := make(chan bool)

	go func() {
		for {
			time.Sleep(12 * time.Minute)
			for d := range msgs {
				var order_tdetails []models.Pg_Element
				buf := bytes.NewBuffer(d.Body)
				decoder := json.NewDecoder(buf)
				err_consume := decoder.Decode(&order_tdetails)
				if err_consume != nil {
					log.Fatal("Error decoding")
				}
				stadistic.StadisticRouter_pg.Import_OrderDetails(order_tdetails)
			}
		}
	}()

	<-noStop2
}

func Consume_UpdateName() {
	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("comensal/changed_name", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop3 := make(chan bool)

	go func() {
		for {
			time.Sleep(12 * time.Minute)
			for d := range msgs {
				var input_name models.Mqtt_UpdateName
				buf := bytes.NewBuffer(d.Body)
				decoder := json.NewDecoder(buf)
				err_consume := decoder.Decode(&input_name)
				if err_consume != nil {
					log.Fatal("Error decoding")
				}
				stadistic.StadisticRouter_pg.Import_NewNameComensal(input_name)
			}
		}
	}()

	<-noStop3
}

func Export_Stadistc() {
	for {
		time.Sleep(15 * time.Minute)
		export.ExportRouter_pg.Export_Stadistic()
	}
}
