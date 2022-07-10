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
	var informationlegal []interface{}

	//Instanciando los valores
	idorders_pg, data_registered_pg, fourcode_pg, idstatus_pg, datelisto_pg, datefinish_pg, dateporfinalizar_pg, schedule_pg, informationbusiness_pg, addressbusiness_pg, informationcomensal_pg, addresscomensal_pg, note_pg, service_pg, payment_pg, _, ismade_by_comensal_pg, _, profit_pg := []int64{}, []time.Time{}, []string{}, []int{}, []string{}, []string{}, []string{}, []models.Pg_Schedule{}, []models.Pg_Information_Business{}, []models.Pg_Address_Business{}, []models.Pg_Information_Comensal{}, []models.Pg_Address_Comensal{}, []string{}, []models.Pg_Service{}, []models.Pg_Payment{}, []models.Pg_Data_Rejected{}, []bool{}, []bool{}, []float32{}

	/*-------------------------------------------------------------------------------------------------------------------------*/

	var insumos_od []interface{}

	//Instanciando los valores
	idelement_od, idbusiness_od, idorder_od, name_od, idcarta_od, url_od, description_od, typemoney_od, unitprice_od, quantity_od, discount_od, category_od, typefood_od, idcategory_od, costos_od, iva_od := []int{}, []int{}, []int64{}, []string{}, []int{}, []string{}, []string{}, []int{}, []float64{}, []int{}, []float32{}, []string{}, []string{}, []int{}, []float64{}, []float64{}

	/*-------------------------------------------------------------------------------------------------------------------------*/

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
		information_workers = append(information_workers, om.Information_Worker)
		ismade_by_comensal_pg = append(ismade_by_comensal_pg, om.Ismadebycomensal)
		informationlegal = append(informationlegal, om.LegalInfo)
		profit_pg = append(profit_pg, om.EstimatedProfit)

		for _, od := range om.Elements {
			idelement_od = append(idelement_od, od.IDElement)
			idbusiness_od = append(idbusiness_od, od.IDBusiness)
			idorder_od = append(idorder_od, om.IDOrder)
			name_od = append(name_od, od.NameE)
			idcarta_od = append(idcarta_od, od.IDCarta)
			url_od = append(url_od, od.URLPhoto)
			description_od = append(description_od, od.Description)
			typemoney_od = append(typemoney_od, od.TypeMoney)
			unitprice_od = append(unitprice_od, od.UnitPrice)
			quantity_od = append(quantity_od, od.Quantity)
			discount_od = append(discount_od, od.Discount)
			category_od = append(category_od, od.Category)
			typefood_od = append(typefood_od, od.TypeFood)
			idcategory_od = append(idcategory_od, od.IdCategory)
			insumos_od = append(insumos_od, od.Insumos)
			costos_od = append(costos_od, od.Costo)
			iva_od = append(iva_od, od.IVA)
		}

	}

	/*-------------------------------------------------------------------------------------------------------------------------*/

	//Enviado los datos a la base de datos
	db := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db.Begin(ctx)
	if error_tx != nil {
		return error_tx
	}

	//ADD ORDERMADE
	query := `INSERT INTO ordermade(idOrder,dateRegistered,fourCode,idStatus,datelisto,datefinish,dateporfinalizar,schedule,informationBusiness,addressBusiness,informationComensal,addressComensal,note,service,payment,informationWorker,ismadebycomensal,informationLegal,estimatedmargin) (select * from 
	unnest($1::bigint[], $2::timestamp[],$3::int[],$4::int[],$5::varchar(35)[],$6::varchar(35)[],$7::varchar(35)[],$8::jsonb[],$9::jsonb[],$10::jsonb[],$11::jsonb[],$12::jsonb[],$13::varchar(200)[],$14::jsonb[],$15::jsonb[],$16::jsonb[],$17::bool[],$18::jsonb[],$19::decimal(10,2)[]))`
	if _, err := tx.Exec(ctx, query, idorders_pg, data_registered_pg, fourcode_pg, idstatus_pg, datelisto_pg, datefinish_pg, dateporfinalizar_pg, schedule_pg, informationbusiness_pg, addressbusiness_pg, informationcomensal_pg, addresscomensal_pg, note_pg, service_pg, payment_pg, information_workers, ismade_by_comensal_pg, informationlegal, profit_pg); err != nil {
		tx.Rollback(ctx)
		return err
	}

	//ADD ORDERDETAILS
	query_od := `INSERT INTO OrderDetails(idelement,idorder,idbusiness,idcarta,unitprice,quantity,discount,namee,descriptione,typemoney,urle,category,typefood,idcategory,insumos,costo,iva) (select * from unnest($1::int[], $2::bigint[],$3::int[],$4::int[],$5::decimal(8,2)[],$6::int[],$7::decimal(8,2)[],$8::varchar(100)[],$9::varchar(250)[],$10::int[],$11::varchar(230)[],$12::varchar(100)[],$13::varchar(100)[],$14::int[],$15::jsonb[],$16::decimal(10,2)[],$17::decimal(10,2)[]))`
	if _, err_od := tx.Exec(ctx, query_od, idelement_od, idorder_od, idbusiness_od, idcarta_od, unitprice_od, quantity_od, discount_od, name_od, description_od, typemoney_od, url_od, category_od, typefood_od, idcategory_od, insumos_od, costos_od, iva_od); err_od != nil {
		tx.Rollback(ctx)
		return err_od
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
