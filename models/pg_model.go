package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pg_Information_Legal struct {
	IDWorker      int    `json:"idworker"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Type          int    `json:"type"`
	Description   string `json:"description"`
	DateLegalized string `json:"datelegalized"`
}

type Pg_Schedule struct {
	IDSchedule        int    `json:"idschedule"`
	IDCarta           int    `json:"idcarta"`
	DateRequired      string `json:"daterequired"`
	TimeStartRequired string `json:"starttime"`
	TimeEndRequired   string `json:"endtime"`
	TimeZone          string `json:"timezone"`
}

type Pg_Information_Business struct {
	IDBusiness      int     `json:"idbusiness"`
	Name            string  `json:"name"`
	Legalidentity   string  `json:"legalidentity"`
	Typesuscription int     `json:"typesuscription"`
	Fee             float64 `json:"fee"`
}

type Pg_Address_Business struct {
	FullAddres string  `json:"fulladdress"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	State      string  `json:"state"`
	City       string  `json:"city"`
	Reference  string  `json:"reference"`
	PostalCode int     `json:"postalcode"`
}

type Pg_Information_Comensal struct {
	IDComensal    int    `json:"idcomensal"`
	Name          string `json:"name"`
	PhoneContact  string `json:"phonecontact"`
	Legalidentity string `json:"legalidentity"`
}

type Pg_Address_Comensal struct {
	Name       string  `json:"name"`
	FullAddres string  `json:"fulladdress"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Reference  string  `json:"reference"`
	State      string  `json:"state"`
	City       string  `json:"city"`
	PostalCode int     `json:"postalcode"`
}

type Pg_Service struct {
	IDService int     `json:"idservice"`
	Typemoney int     `json:"typemoney"`
	Price     float32 `json:"price"`
}

type Pg_Information_Worker struct {
	IDWorker     int    `json:"idworker"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	CreatedOrder string `json:"createdorder"`
}

type Pg_Payment struct {
	IDPayment    int    `json:"idpayment"`
	Name         string `json:"name"`
	UrlPhoto     string `json:"url"`
	HasNumber    bool   `json:"hasnumber"`
	PhoneContact string `json:"phonecontact"`
}

type Pg_Data_Rejected struct {
	MadeByComensal bool   `json:"madebycomensal"`
	Date           string `json:"date"`
	Mesage         string `json:"message"`
}

type Pg_Element struct {
	IDElement   int              `json:"idelement"`
	IDBusiness  int              `json:"idbusiness"`
	IDOrder     int64            `json:"idorder"`
	NameE       string           `json:"name"`
	IdCategory  int              `json:"idcategory"`
	Category    string           `json:"category"`
	Typefood    string           `json:"typefood"`
	IDCarta     int              `json:"idcarta"`
	URLPhoto    string           `json:"url"`
	Description string           `json:"description"`
	TypeMoney   int              `json:"typemoney"`
	UnitPrice   float64          `json:"unitprice"`
	Quantity    int              `json:"quantity"`
	Discount    float32          `json:"discount"`
	Additionals []Pg_Additionals `json:"additionals"`
	Insumos     []Pg_Insumo      `json:"insumos"`
	Costos      float64          `json:"costos"`
	IVA         float32          `json:"iva"`
}

type Pg_Items struct {
	IDItem   string  `json:"id"`
	Name     string  `json:"name"`
	IsInsumo bool    `json:"isinsumo"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type Pg_Additionals struct {
	IDSubElement string     `json:"id"`
	Name         string     `json:"name"`
	MaxSelect    int        `json:"maxselect"`
	IsMandatory  bool       `json:"ismandatory"`
	Items        []Pg_Items `json:"items"`
}

type Pg_Insumo struct {
	Insumo   Mo_Insumo_Response `json:"insumo"`
	Quantity int                `json:"quantity"`
}

type Pg_Comensales struct {
	IdComensal int    `json:"idcomensal"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Orders     int    `json:"orders"`
}

type Mo_Insumo_Response struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Available      bool               `json:"available"`
	SendToDelete   time.Time          `json:"sendtodelete"`
}

type Mo_Stock struct {
	Price        float64   `json:"price"`
	IdProvider   string    `json:"idprovider"`
	TimeZone     string    `json:"timezone"`
	CreatedDate  time.Time `json:"createdDate"`
	Quantity     int       `json:"quantity"`
	ProviderName string    `json:"providername"`
}

type Pg_ToElement_Mqtt struct {
	IDElement []int `json:"idElement"`
	IDCarta   []int `json:"idCarta"`
	Quantity  []int `json:"Quantity"`
}

type Pg_ToSchedule_Mqtt struct {
	IDSchedule int `json:"idSchedule"`
	IDCarta    int `json:"idCarta"`
	Quantity   int `json:"Quantity"`
}

type Pg_Order_ToCopy struct {
	IDOrder              int64                   `json:"id"`
	DateRegistered       time.Time               `json:"dateregistered"`
	FourCode             int                     `json:"fourcode"`
	IdStatus             int                     `json:"idstatus"`
	DateListo            string                  `json:"datelisto"`
	DateFinish           string                  `json:"datefinish"`
	DatePorFInalizar     string                  `json:"dateporfinalizar"`
	Schedule             Pg_Schedule             `json:"schedule"`
	Information_Business Pg_Information_Business `json:"informationbusiness"`
	Address_Busines      Pg_Address_Business     `json:"addressbusiness"`
	Information_Comensal Pg_Information_Comensal `json:"informationcomensal"`
	Address_Comensal     Pg_Address_Comensal     `json:"addresscomensal"`
	Information_Worker   Pg_Information_Worker   `json:"informationworker"`
	Ismadebycomensal     bool                    `json:"ismadebycomensal"`
	Note                 string                  `json:"note"`
	Service              Pg_Service              `json:"service"`
	Payment              Pg_Payment              `json:"payment"`
	Elements             []V2_Pg_Element         `json:"elements"`
	DataRejected         Pg_Data_Rejected        `json:"datarejected"`
	LegalInfo            Pg_Information_Legal    `json:"informationlegal"`
	IsLegal              bool                    `json:"islegal"`
	EstimatedProfit      float32                 `json:"profitmargin"`
	TotalDiscount        float32                 `json:"totaldiscount"`
	TotalIGV             float32                 `json:"totaligv"`
	TotalSales           float32                 `json:"totalsales"`
	Typemoney            int                     `json:"typemoney"`
	IsMadeByWeb          bool                    `json:"ismadebyweb"`
}

type V2_Pg_Element struct {
	IDElement   int                     `json:"idelement"`
	IDBusiness  int                     `json:"idbusiness"`
	IDCarta     int                     `json:"idcarta"`
	NameE       string                  `json:"name"`
	IdCategory  int                     `json:"idcategory"`
	Category    string                  `json:"category"`
	TypeFood    string                  `json:"typefood"`
	URLPhoto    string                  `json:"url"`
	Description string                  `json:"description"`
	TypeMoney   int                     `json:"typemoney"`
	UnitPrice   float64                 `json:"unitprice"`
	Quantity    int                     `json:"quantity"`
	Discount    float64                 `json:"discount"`
	Latitude    float64                 `json:"latitude"`
	Longitude   float64                 `json:"longitude"`
	Insumos     []Pg_Mo_Insumo_Elements `json:"insumos"`
	Additionals []Pg_Additionals        `json:"additionals"`
	Costo       float64                 `json:"costo"`
	IVA         float64                 `json:"iva"`
}

type Pg_Mo_Insumo_Elements struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Quantity       int                `json:"quantity"`
}

type Pg_Stadistic_Comensal struct {
	Orders_lastdateorder interface{} `json:"lastdateorder"`
	Outgoing             float32     `json:"outgoing"`
	Orders               int         `json:"orders"`
	Outgoing_Week        interface{} `json:"outgoingbyweek"`
	Orders_typefood      interface{} `json:"ordersbytypefood"`
}

type Pg_Stadistic_Anfitrion_Orders struct {
	Orders_lastdateorder interface{} `json:"lastdateorder"`
	Orders               int         `json:"total"`
	Orders_by_week       interface{} `json:"ordersbyweek"`
	Orders_by_day        interface{} `json:"ordersbyday"`
	Orders_by_service    interface{} `json:"ordersbyservice"`
	Orders_by_payment    interface{} `json:"ordersbypayment"`
	Elements             interface{} `json:"elements"`
}

type Pg_Stadistic_Anfitrion_Incoming struct {
	Incoming_lastdateorder string      `json:"lastdateorder"`
	Incoming               interface{} `json:"total"`
	Incoming_by_week       interface{} `json:"incomingbyweek"`
	Incoming_by_day        interface{} `json:"incomingbyday"`
	Incoming_by_service    interface{} `json:"incomingbyservice"`
	Incoming_by_payment    interface{} `json:"incomingbypayment"`
	Elements               interface{} `json:"elements"`
}

type Pg_Export_ByElement struct {
	IdElement   int     `json:"idelement"`
	Quantity    int     `json:"quantity"`
	Datetime    string  `json:"datetime"`
	TotalAmount float32 `json:"totalamount"`
	TotalCost   float32 `json:"totalcost"`
}

type Pg_ToNotify struct {
	IDBusiness    int     `json:"idbusiness"`
	Orders        int     `json:"orders"`
	GrossIncoming float64 `json:"grossincoming"`
	NetIncoming   float64 `json:"netincoming"`
	NetUtility    float64 `json:"netutility"`
}

type Pg_ToExportFee struct {
	IDBusiness  int     `json:"idbusiness"`
	TotalOrders int     `json:"totalorders"`
	TotalAmount float64 `json:"totalamount"`
}

type Pg_Movement struct {
	IdMovement     int       `json:"idmovement"`
	Dateregistered time.Time `json:"dateregistered"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Type           int       `json:"type"`
	Timezone       string    `json:"timezone"`
}
