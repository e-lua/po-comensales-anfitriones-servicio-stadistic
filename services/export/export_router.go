package export

import (
	"log"
)

var ExportRouter_pg *exportRouter_pg

type exportRouter_pg struct {
}

func (er *exportRouter_pg) Export_Stadistic() {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Export_Stadistic_Service()
	log.Println(status, boolerror, dataerror, data)
}
