package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-anfitriones-servicio-stadistic/models"
)

func Pg_Insert_OrderDetails(orderdetails []models.Pg_Element) error {

	//Instanciando los valores
	idelement_pg, idbusiness_pg, idorder_pg, name_pg, idcarta_pg, url_pg, description_pg, typemoney_pg, unitprice_pg, quantity_pg, discount_pg := []int{}, []int{}, []int64{}, []string{}, []int{}, []string{}, []string{}, []int{}, []float64{}, []int{}, []float32{}

	for _, od := range orderdetails {
		idelement_pg = append(idelement_pg, od.IDElement)
		idbusiness_pg = append(idbusiness_pg, od.IDBusiness)
		idorder_pg = append(idorder_pg, od.IDOrder)
		name_pg = append(name_pg, od.NameE)
		idcarta_pg = append(idcarta_pg, od.IDCarta)
		url_pg = append(url_pg, od.URLPhoto)
		description_pg = append(description_pg, od.Description)
		typemoney_pg = append(typemoney_pg, od.TypeMoney)
		unitprice_pg = append(unitprice_pg, od.UnitPrice)
		quantity_pg = append(quantity_pg, od.Quantity)
		discount_pg = append(discount_pg, od.Discount)
	}

	//Enviado los datos a la base de datos
	db := models.Conectar_Pg_DB()

	query := `INSERT INTO OrderDetails(idelement,idorder,idbusiness,idcarta,unitprice,quantity,discount,namee,descriptione,typemoney,urle) (select * from unnest($1::int[], $2::bigint[],$3::int[],$4::int[],$5::decimal(8,2)[],$6::int[],$7::decimal(8,2)[],$8::varchar(100)[],$9::varchar(250)[],$10::int[],$11::varchar(230)[]))`
	if _, err := db.Exec(context.Background(), query, idelement_pg, idorder_pg, idbusiness_pg, idcarta_pg, unitprice_pg, quantity_pg, discount_pg, name_pg, description_pg, typemoney_pg, url_pg); err != nil {
		return err
	}

	return nil
}
