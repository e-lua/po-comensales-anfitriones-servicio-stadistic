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

type Response_StadisticAnfitrion_Order struct {
	Error     bool                                 `json:"error"`
	DataError string                               `json:"dataError"`
	Data      models.Pg_Stadistic_Anfitrion_Orders `json:"data"`
}

type Response_StadisticAnfitrion_Incoming struct {
	Error     bool                                   `json:"error"`
	DataError string                                 `json:"dataError"`
	Data      models.Pg_Stadistic_Anfitrion_Incoming `json:"data"`
}

type Response_StadisticAnfitrion_Comensal struct {
	Error     bool                   `json:"error"`
	DataError string                 `json:"dataError"`
	Data      []models.Pg_Comensales `json:"data"`
}

type Response_StadisticElements_ByDay struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      []interface{} `json:"data"`
}

type ResponseJWT_B struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT_B  `json:"data"`
}

type JWT_B struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}
