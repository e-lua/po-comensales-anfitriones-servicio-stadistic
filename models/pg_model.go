package models

import "time"

type Pg_Order struct {
	IDOrder              int64                   `json:"id"`
	DateRegistered       string                  `json:"dateregistered"`
	FourCode             int                     `json:"fourcode"`
	IdStatus             int                     `json:"idstatus"`
	Schedule             Pg_Schedule             `json:"schedule"`
	Information_Business Pg_Information_Business `json:"informationbusiness"`
	Address_Busines      Pg_Address_Business     `json:"addressbusiness"`
	Information_Comensal Pg_Information_Comensal `json:"informationcomensal"`
	Address_Comensal     Pg_Address_Comensal     `json:"addresscomensal"`
	Note                 string                  `json:"note"`
	Service              Pg_Service              `json:"service"`
	Payment              Pg_Payment              `json:"payment"`
	DataRejected         Pg_Data_Rejected        `json:"datarejected"`
	Elements             []Pg_Element            `json:"elements"`
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
	IDBusiness int    `json:"idbusiness"`
	Name       string `json:"name"`
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
	IDComensal   int    `json:"idcomensal"`
	Name         string `json:"name"`
	PhoneContact string `json:"phonecontact"`
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
	IDElement   int     `json:"idelement"`
	IDBusiness  int     `json:"idbusiness"`
	IDOrder     int64   `json:"idorder"`
	NameE       string  `json:"name"`
	Category    string  `json:"category"`
	Typefood    string  `json:"typefood"`
	IDCarta     int     `json:"idcarta"`
	URLPhoto    string  `json:"url"`
	Description string  `json:"description"`
	TypeMoney   int     `json:"typemoney"`
	UnitPrice   float64 `json:"unitprice"`
	Quantity    int     `json:"quantity"`
	Discount    float32 `json:"discount"`
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
	Note                 string                  `json:"note"`
	Service              Pg_Service              `json:"service"`
	Payment              Pg_Payment              `json:"payment"`
	DataRejected         Pg_Data_Rejected        `json:"datarejected"`
}

type Pg_Stadistic_Comensal struct {
	Outgoing        float32     `json:"outgoing"`
	Orders          int         `json:"orders"`
	Outgoing_Week   interface{} `json:"outgoingbyweek"`
	Orders_typefood interface{} `json:"ordersbytypefood"`
}

type Pg_Stadistic_Anfitrion_Orders struct {
	Orders            int         `json:"total"`
	Orders_by_week    interface{} `json:"ordersbyweek"`
	Orders_by_day     interface{} `json:"ordersbyday"`
	Orders_by_service interface{} `json:"ordersbyservice"`
	Orders_by_payment interface{} `json:"ordersbypayment"`
	Elements          interface{} `json:"elements"`
}

type Pg_Stadistic_Anfitrion_Incoming struct {
	Incoming            interface{} `json:"total"`
	Incoming_by_week    interface{} `json:"incomingbyweek"`
	Incoming_by_day     interface{} `json:"incomingbyday"`
	Incoming_by_service interface{} `json:"incomingbyservice"`
	Incoming_by_payment interface{} `json:"incomingbypayment"`
	Elements            interface{} `json:"elements"`
}
