package notification

import "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"

//NOTIFY

type Response_Notify_Stadistic struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      []models.Pg_ToNotify `json:"data"`
}
