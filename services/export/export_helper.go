package export

import "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"

type Response_ToExportFee struct {
	Error     bool                    `json:"error"`
	DataError string                  `json:"dataError"`
	Data      []models.Pg_ToExportFee `json:"data"`
}
