package repositories

import (
	"context"
	"strconv"
	"time"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Insert_OrderMade(ordermades []models.Pg_Order_ToCopy) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var information_workers []interface{}

	//Instanciando los valores
	idorders_pg, data_registered_pg, fourcode_pg, idstatus_pg, datelisto_pg, datefinish_pg, dateporfinalizar_pg, schedule_pg, informationbusiness_pg, addressbusiness_pg, informationcomensal_pg, addresscomensal_pg, note_pg, service_pg, payment_pg, datarejected_pg, ismade_by_comensal_pg := []int64{}, []time.Time{}, []string{}, []int{}, []string{}, []string{}, []string{}, []models.Pg_Schedule{}, []models.Pg_Information_Business{}, []models.Pg_Address_Business{}, []models.Pg_Information_Comensal{}, []models.Pg_Address_Comensal{}, []string{}, []models.Pg_Service{}, []models.Pg_Payment{}, []models.Pg_Data_Rejected{}, []bool{}

	for _, om := range ordermades {
		idorders_pg = append(idorders_pg, om.IDOrder)
		data_registered_pg = append(data_registered_pg, om.DateRegistered)
		fourcode_pg = append(fourcode_pg, strconv.Itoa(om.FourCode))
		idstatus_pg = append(idstatus_pg, om.IdStatus)
		datelisto_pg = append(datelisto_pg, om.DateListo)
		datefinish_pg = append(datefinish_pg, om.DateFinish)
		dateporfinalizar_pg = append(dateporfinalizar_pg, om.DatePorFInalizar)
		schedule_pg = append(schedule_pg, om.Schedule)
		informationbusiness_pg = append(informationbusiness_pg, om.Information_Business)
		addressbusiness_pg = append(addressbusiness_pg, om.Address_Busines)
		informationcomensal_pg = append(informationcomensal_pg, om.Information_Comensal)
		addresscomensal_pg = append(addresscomensal_pg, om.Address_Comensal)
		note_pg = append(note_pg, om.Note)
		service_pg = append(service_pg, om.Service)
		payment_pg = append(payment_pg, om.Payment)
		datarejected_pg = append(datarejected_pg, om.DataRejected)
		information_workers = append(information_workers, om.Information_Worker)
		ismade_by_comensal_pg = append(ismade_by_comensal_pg, om.Ismadebycomensal)
	}

	//Enviado los datos a la base de datos
	db := models.Conectar_Pg_DB()

	query := `INSERT INTO ordermade(idOrder,dateRegistered,fourCode,idStatus,datelisto,datefinish,dateporfinalizar,schedule,informationBusiness,addressBusiness,informationComensal,addressComensal,note,service,payment,datarejected) (select * from unnest($1::bigint[], $2::timestamp[],$3::int[],$4::int[],$5::varchar(35)[],$6::varchar(35)[],$7::varchar(35)[],$8::jsonb[],$9::jsonb[],$10::jsonb[],$11::jsonb[],$12::jsonb[],$13::varchar(200)[],$14::jsonb[],$15::jsonb[],$16::jsonb[],$17::jsonb[],$18::bool[]))`
	if _, err := db.Exec(ctx, query, idorders_pg, data_registered_pg, fourcode_pg, idstatus_pg, datelisto_pg, datefinish_pg, dateporfinalizar_pg, schedule_pg, informationbusiness_pg, addressbusiness_pg, informationcomensal_pg, addresscomensal_pg, note_pg, service_pg, payment_pg, datarejected_pg, information_workers, ismade_by_comensal_pg); err != nil {
		return err
	}

	return nil
}
