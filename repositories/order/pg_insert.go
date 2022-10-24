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
	var insumo_element_pg []interface{}

	//Instanciando los valores
	idorders_pg, data_registered_pg, fourcode_pg, idstatus_pg, datelisto_pg, datefinish_pg, dateporfinalizar_pg, schedule_pg, informationbusiness_pg, addressbusiness_pg, informationcomensal_pg, addresscomensal_pg, note_pg, service_pg, payment_pg, _, ismade_by_comensal_pg, _, profit_pg, totaldiscount_pg, totaliva_pg, totalsales_pg, typemoney_pg, ismadebyweb_pg, idorder_element_pg, totaldiscount_element_pg, totaliva_element_pg, totalsales_element_pg, iva_element_pg, ismadebyweb_element_pg, latitude_pg, longitude_pg, idbusiness_element_pg, category_element_pg, name_element_pg, typefood_element_pg, totalcost_element_pg, totalquantity_element_pg, typemoney_element_pg, idcategory_element_pg, totalcost_pg, totalquantity_pg, islegal_pg, daterequired_element_pg, hour_element_pg, totalfee_pg := []int64{}, []time.Time{}, []string{}, []int{}, []string{}, []string{}, []string{}, []models.Pg_Schedule{}, []models.Pg_Information_Business{}, []models.Pg_Address_Business{}, []models.Pg_Information_Comensal{}, []models.Pg_Address_Comensal{}, []string{}, []models.Pg_Service{}, []models.Pg_Payment{}, []models.Pg_Data_Rejected{}, []bool{}, []bool{}, []float32{}, []float64{}, []float64{}, []float64{}, []int{}, []bool{}, []int64{}, []float64{}, []float64{}, []float64{}, []float64{}, []bool{}, []float64{}, []float64{}, []int{}, []string{}, []string{}, []string{}, []float64{}, []int{}, []int{}, []int{}, []float64{}, []int{}, []bool{}, []string{}, []string{}, []float64{}

	/*-------------------------------------------------------------------------------------------------------------------------*/

	//Instanciando los valores
	idelement_od := []int{}

	/*-------------------------------------------------------------------------------------------------------------------------*/

	for _, om := range ordermades {

		idorders_pg = append(idorders_pg, om.IDOrder)
		data_registered_pg = append(data_registered_pg, om.DateRegistered)
		fourcode_pg = append(fourcode_pg, strconv.Itoa(om.FourCode))
		idstatus_pg = append(idstatus_pg, om.IdStatus)

		if !om.DataRejected.MadeByComensal {
			datelisto_pg = append(datelisto_pg, om.DateRegistered.Format("2006-01-02 15:04:05"))
			datefinish_pg = append(datefinish_pg, om.DateRegistered.Format("2006-01-02 15:04:05"))
			dateporfinalizar_pg = append(dateporfinalizar_pg, om.DateRegistered.Format("2006-01-02 15:04:05"))
		} else {
			datelisto_pg = append(datelisto_pg, om.DateListo)
			datefinish_pg = append(datefinish_pg, om.DateFinish)
			dateporfinalizar_pg = append(dateporfinalizar_pg, om.DatePorFInalizar)
		}

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

		var totaldiscount float64
		totaldiscount = 0

		var totalsales float64
		totalsales = 0

		var totaliva float64
		totaliva = 0

		var totalcost float64
		totalcost = 0

		var totalquantity int
		totalquantity = 0

		var typemoney int

		for _, od := range om.Elements {

			var montoAdicionales float64
			montoAdicionales = 0

			for _, aditional := range od.Additionals {
				for _, item := range aditional.Items {
					montoAdicionales = montoAdicionales + (float64(item.Quantity) * float64(item.Price))
				}
			}

			totalamount_element := (od.UnitPrice * float64(od.Quantity)) - float64(od.Discount) + montoAdicionales
			totaliva_element := totalamount_element * od.IVA

			/**/
			totaldiscount = totaldiscount + od.Discount
			totalsales = totalsales + totalamount_element
			totaliva = totaliva + totaliva_element
			totalcost = totalcost + (od.Costo * float64(od.Quantity))
			typemoney = od.TypeMoney
			totalquantity = totalquantity + od.Quantity
			/**/

			idorder_element_pg = append(idorder_element_pg, om.IDOrder)
			idelement_od = append(idelement_od, od.IDElement)
			totalsales_element_pg = append(totalsales_element_pg, totalamount_element)
			totaliva_element_pg = append(totaliva_element_pg, totaliva_element)
			iva_element_pg = append(iva_element_pg, od.IVA)
			totaldiscount_element_pg = append(totaldiscount_element_pg, od.Discount)
			ismadebyweb_element_pg = append(ismadebyweb_element_pg, om.IsMadeByWeb)
			latitude_pg = append(latitude_pg, od.Latitude)
			longitude_pg = append(longitude_pg, od.Longitude)
			idbusiness_element_pg = append(idbusiness_element_pg, om.Information_Business.IDBusiness)
			category_element_pg = append(category_element_pg, od.Category)
			name_element_pg = append(name_element_pg, od.NameE)
			typemoney_element_pg = append(typemoney_element_pg, od.TypeMoney)
			typefood_element_pg = append(typefood_element_pg, od.TypeFood)
			idcategory_element_pg = append(idcategory_element_pg, od.IdCategory)
			insumo_element_pg = append(insumo_element_pg, od.Insumos)
			totalcost_element_pg = append(totalcost_element_pg, od.Costo*float64(od.Quantity))
			totalquantity_element_pg = append(totalquantity_element_pg, od.Quantity)
			daterequired_element_pg = append(daterequired_element_pg, om.Schedule.DateRequired)
			hour_element_pg = append(hour_element_pg, om.Schedule.TimeStartRequired)
		}

		totaldiscount_pg = append(totaldiscount_pg, totaldiscount)
		totalsales_pg = append(totalsales_pg, totalsales)
		totaliva_pg = append(totaliva_pg, totaliva)
		ismadebyweb_pg = append(ismadebyweb_pg, om.IsMadeByWeb)
		typemoney_pg = append(typemoney_pg, typemoney)
		totalcost_pg = append(totalcost_pg, totalcost)
		totalquantity_pg = append(totalquantity_pg, totalquantity)
		totalfee_pg = append(totalfee_pg, float64(totalquantity)*om.Information_Business.Fee)
		islegal_pg = append(islegal_pg, om.IsLegal)
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
	query := `INSERT INTO ordermade(idOrder,dateRegistered,fourCode,idStatus,datelisto,datefinish,dateporfinalizar,schedule,informationBusiness,addressBusiness,informationComensal,addressComensal,note,service,payment,informationWorker,ismadebycomensal,informationLegal,estimatedmargin,totaldiscount,totaliva,totalsales,typemoney,ismadebyweb,islegal,totalcost,totalquantity,totalfee) (select * from 
	unnest($1::bigint[], $2::timestamp[],$3::int[],$4::int[],$5::varchar(35)[],$6::varchar(35)[],$7::varchar(35)[],$8::jsonb[],$9::jsonb[],$10::jsonb[],$11::jsonb[],$12::jsonb[],$13::varchar(200)[],$14::jsonb[],$15::jsonb[],$16::jsonb[],$17::bool[],$18::jsonb[],$19::decimal(10,2)[],$20::decimal(10,2)[],$21::decimal(10,2)[],$22::decimal(10,2)[],$23::integer[],$24::boolean[],$25::boolean[],$26::decimal(10,2)[],$27::int[],$28::decimal(10,2)[]))`
	if _, err := tx.Exec(ctx, query, idorders_pg, data_registered_pg, fourcode_pg, idstatus_pg, datelisto_pg, datefinish_pg, dateporfinalizar_pg, schedule_pg, informationbusiness_pg, addressbusiness_pg, informationcomensal_pg, addresscomensal_pg, note_pg, service_pg, payment_pg, information_workers, ismade_by_comensal_pg, informationlegal, profit_pg, totaldiscount_pg, totaliva_pg, totalsales_pg, typemoney_pg, ismadebyweb_pg, islegal_pg, totalcost_pg, totalquantity_pg, totalfee_pg); err != nil {
		tx.Rollback(ctx)
		return err
	}

	//ADD ORDERDETAILS
	/*query_od := `INSERT INTO OrderDetails(idelement,idorder,idbusiness,idcarta,unitprice,quantity,discount,namee,descriptione,typemoney,urle,category,typefood,idcategory,insumos,costo,iva) (select * from unnest($1::int[], $2::bigint[],$3::int[],$4::int[],$5::decimal(8,2)[],$6::int[],$7::decimal(8,2)[],$8::varchar(100)[],$9::varchar(250)[],$10::int[],$11::varchar(230)[],$12::varchar(100)[],$13::varchar(100)[],$14::int[],$15::jsonb[],$16::decimal(10,2)[],$17::decimal(10,2)[]))`
	if _, err_od := tx.Exec(ctx, query_od, idelement_od, idorder_od, idbusiness_od, idcarta_od, unitprice_od, quantity_od, discount_od, name_od, description_od, typemoney_od, url_od, category_od, typefood_od, idcategory_od, insumos_od, costos_od, iva_od); err_od != nil {
		tx.Rollback(ctx)
		return err_od
	}*/

	query_od := `INSERT INTO OrderDetail(idelement,totalsales,totaliva,iva,totaldiscount,ismadebyweb,latitude,longitude,idbusiness,idorder,category,namee,typefood,totalcost,totalquantity,typemoney,idcategory,insumos,dateregistered,hour) (select * from unnest($1::int[],$2::decimal(10,2)[],$3::decimal(10,2)[],$4::decimal(10,2)[],$5::decimal(10,2)[],$6::bool[],$7::real[],$8::real[],$9::int[],$10::bigint[],$11::varchar(100)[],$12::varchar(100)[],$13::varchar(100)[],$14::decimal(10,2)[],$15::int[],$16::int[],$17::int[],$18::jsonb[],$19::varchar(50)[],$20::varchar(50)[]))`
	if _, err_od := tx.Exec(ctx, query_od, idelement_od, totalsales_element_pg, totaliva_element_pg, iva_element_pg, totaldiscount_element_pg, ismadebyweb_element_pg, latitude_pg, longitude_pg, idbusiness_element_pg, idorder_element_pg, category_element_pg, name_element_pg, typefood_element_pg, totalcost_element_pg, totalquantity_element_pg, typemoney_element_pg, idcategory_element_pg, insumo_element_pg, daterequired_element_pg, hour_element_pg); err_od != nil {
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
