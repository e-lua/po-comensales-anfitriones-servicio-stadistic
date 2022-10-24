package movement

import "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type JWT_Business struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type ResponseJWT_Business struct {
	Error     bool         `json:"error"`
	DataError string       `json:"dataError"`
	Data      JWT_Business `json:"data"`
}

type ResponseMovement struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      []models.Pg_Movement `json:"data"`
}
