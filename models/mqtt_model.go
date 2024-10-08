package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mqtt_UpdateName struct {
	Name       string `json:"name"`
	IdComensal int    `json:"idcomensal"`
}

type Mqtt_Stock struct {
	Price        float64   `json:"price"`
	IdProvider   string    `json:"idprovider"`
	TimeZone     string    `json:"timezone"`
	CreatedDate  time.Time `json:"createdDate"`
	Quantity     int       `json:"quantity"`
	ProviderName string    `json:"providername"`
}

type Mqtt_Insumo_Elements struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Quantity int                `json:"quantity"`
}

type Mqtt_Element_Order struct {
	NameE       string                 `json:"f1"`
	Category    string                 `json:"f2"`
	TypeFood    string                 `json:"f3"`
	URLPhoto    string                 `json:"f4"`
	Description string                 `json:"f5"`
	Insumos     []Mqtt_Insumo_Elements `json:"f6"`
	UnitPrice   float64                `json:"f7"`
	Discount    float32                `json:"f8"`
	Costo       float32                `json:"f9"`
	IVA         float32                `json:"f10"`
}

type Mqtt_Order struct {
	DateRegistered       string                  `json:"dateregistered"`
	IDOrder              int64                   `json:"id"`
	FourCode             string                  `json:"fourcode"`
	Schedule             Pg_Schedule             `json:"schedule"`
	Information_Business Pg_Information_Business `json:"informationbusiness"`
	Address_Busines      Pg_Address_Business     `json:"addressbusiness"`
	Information_Comensal Pg_Information_Comensal `json:"informationcomensal"`
	Address_Comensal     Pg_Address_Comensal     `json:"addresscomensal"`
	Note                 string                  `json:"note"`
	Service              Pg_Service              `json:"service"`
	Payment              Pg_Payment              `json:"payment"`
	DataRejected         Pg_Data_Rejected        `json:"datarejected"`
	Elements             []Mqtt_Element_Order    `json:"elements"`
	EstimatedProfit      float32                 `json:"estimatedprofit"`
	TotalPrice           float32                 `json:"totalprice"`
	Information_Worker   Pg_Information_Worker   `json:"informationworker"`
	Ismadebycomensal     bool                    `json:"ismadebycomensal"`
}

type Mqtt_Request_Order struct {
	IDBusiness int          `json:"idbusiness"`
	Email      string       `json:"email"`
	Orders     []Mqtt_Order `json:"orders"`
}
