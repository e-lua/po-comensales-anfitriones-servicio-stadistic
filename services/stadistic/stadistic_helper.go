package order

import models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	Phone      int    `json:"phone"`
	Country    int    `json:"country"`
	IDComensal int    ` json:"comensal"`
	Name       string ` json:"name"`
	LastName   string ` json:"lastName"`
}

type Response_OrderMade struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      []models.Pg_Order_ToCopy `json:"data"`
}

type Response_OrderDetails struct {
	Error     bool                `json:"error"`
	DataError string              `json:"dataError"`
	Data      []models.Pg_Element `json:"data"`
}

type Response_StadisticComensal struct {
	Error     bool                         `json:"error"`
	DataError string                       `json:"dataError"`
	Data      models.Pg_Stadistic_Comensal `json:"data"`
}
